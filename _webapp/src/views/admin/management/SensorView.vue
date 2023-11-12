<template>
  <div class="row align-items-stretch">
    <div class="col-md-3 my-2">
      <div class="card h-100">
        <div class="card-body">
          <div class="d-flex align-items-center">
            <div class="h5 my-0 me-2">站点信息</div>
            <button type="button" class="btn btn-sm btn-outline-primary" @click="gotoStationsView()">
              查看所有站点
            </button>
          </div>

          <table class="table my-3">
            <tr>
              <td width="20%">名称:</td>
              <td>{{ station?.name }}</td>
            </tr>
            <tr>
              <td>纬度:</td>
              <td>{{ station?.lat }}</td>
            </tr>
            <tr>
              <td>经度:</td>
              <td>{{ station?.lng }}</td>
            </tr>
            <tr>
              <td>海拔:</td>
              <td>{{ station?.altitude }}</td>
            </tr>
          </table>

          <div class="d-flex align-items-center">
            <h5 class="my-0 me-2">上传数据</h5>
            <button class="btn btn-sm btn-outline-primary my-2" @click="handleTemplateDownload">
              下载数据模版
            </button>
          </div>

          <div class="input-group my-2">
            <input type="file" class="form-control" @change="selectCSVFile" />
            <button class="btn btn-outline-primary" type="button" :disabled="disableSubmitButton" @click="uploadCSVFile">
              Submit
            </button>
          </div>
        </div>
      </div>
    </div>
    <div class="col-md-5 my-2">
      <div class="card h-100">
        <div class="card-body">
          <div class="row">
            <div class="col-md-6">
              <div class="d-flex align-items-center">
                <div class="h5 mb-0 me-2">传感器列表</div>
              </div>
              <div class="d-flex align-items-center my-2">
                <button type="button" style="margin-right: 8px" class="btn btn-sm btn-success"
                  @click="showCreateSensorModal">
                  + 添加传感器
                </button>
                <button type="button" class="btn btn-sm btn-success" @click="showLinkSensorModal">
                  @ 关联传感器
                </button>
              </div>
              <label class="text-secondary small mb-3">选择查看相应的传感器</label>

              <div style="height: 170px; overflow-y: auto">
                <div class="form-check mb-3 d-flex align-items-center" v-for="sensor in allSensors" :key="sensor.id">
                  <input class="form-check-input" style="margin-top: 0; flex-shrink: 0" type="checkbox"
                    @change="onChangeSensor($event, sensor.id)" />
                  <label class="form-check-label ms-1">
                    {{ getSensorDisplayText(sensor) }}
                  </label>
                  <div v-if="operatingSensorId == sensor.id" class="ms-auto small">
                    <span class="spinner-grow spinner-grow-sm me-1" role="status" aria-hidden="true"></span>
                    修改中...
                  </div>
                  <div v-else class="d-flex align-items-center ms-auto">
                    <i class="bi bi-exclamation-circle-fill text-primary mx-1" style="cursor: pointer"
                      @click="showSensorUpsertModal(sensor)"></i>
                    <i class="bi bi-x-circle-fill text-danger mx-1" style="cursor: pointer"
                      @click="onDeleteSensor(sensor.id)"></i>
                  </div>
                </div>
              </div>
            </div>
            <div class="col-md-6">
              <div class="h5 mb-0">时间范围</div>
              <div class="my-2">
                <div>开始时间</div>
                <Datepicker v-model="startDate" class="btn btn-sm btn-primary" style="cursor: pointer"
                  @closed="onDatePickerClosed" />
              </div>

              <div class="my-2">
                <div>结束时间</div>
                <Datepicker v-model="endDate" class="btn btn-sm btn-primary" style="cursor: pointer"
                  @closed="onDatePickerClosed" />
              </div>
              <div class="h6 mb-0">每页数据时间跨度(30mins * limit)</div>
              <div class="my-2">
                <select class="form-select" @input="onLimitChanged">
                  <option v-for="limitOption in [10, 20, 50, 100, 200]" :key="limitOption" :value="limitOption"
                    :selected="limitOption == limitVal">
                    {{ limitOption }}
                  </option>
                </select>
              </div>

              <button class="btn btn-sm btn-outline-primary my-2" :disabled="totalVal <= 0" @click="handleDownloadData">
                <div v-if="isDownloadData">
                  <span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>
                  下载中...
                </div>
                <div v-if="!isDownloadData">下载全部数据</div>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="col-md-4 my-2">
      <div class="card h-100">
        <div class="card-body">
          <line-chart :stations="station ? [station] : []" :sensors="selectedSensors" :records="displayedRecords"
            title="当前数据" default-text="请选择数据项" no-data-text="暂无数据" :show-default-text="false" :loading="loading" />
        </div>
      </div>
    </div>
    <upsert-sensor-modal ref="upsertSensorModal" :visible="isShowingUpsertSensorModal"
      :title="operatingSensorId ? '编辑传感器' : '添加传感器'" @close="onUpsertSensorModalClosed"
      @did-upsert-sensor="didUpsertSensor" />
    <link-sensors-modal ref="linkSensorModal" :visible="isShowingLinkSensorsModal" @close="onCloseLinkedSensors"
      @on-confirm="onConfirmLinkedSensors" />
    <double-confirm-modal :visible="isShowingDeletionModal" :title="'删除传感器 ' + operatingSensorId" content="确定删除传感器？"
      @close="onDeleteSensorModalClosed" @on-confirm="onConfirmDeleteSensor" />
  </div>

  <div v-if="loading" class="mx-auto mt-5 text-center">
    <div class="spinner-border mx-auto" role="status">
      <span class="sr-only"></span>
    </div>
  </div>

  <div v-else-if="!totalVal || totalVal == 0" class="mx-auto mt-5 text-center">
    <label>无法查询到数据</label>
  </div>

  <div v-else class="row">
    <div class="col-md-12">
      <div class="card">
        <table class="table card-table align-middle">
          <thead class="table-dark">
            <tr>
              <th scope="col" width="30%">时间</th>
              <th scope="col" v-for="s in selectedSensors" :key="s.id">
                {{ getSensorDisplayText(s) }}
              </th>
              <th scope="col" width="20%">操作</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="[date, sensorRecordMap] in recordsByTime" :key="date.toString()">
              <td>{{ formatDatetime(date) }}</td>
              <td v-for="sensor in selectedSensors" :key="sensor.id">
                <div v-if="rowEditingIndex != date || sensor.id == 0">
                  <div v-if="sensorRecordMap.get(sensor.id)?.value">
                    {{ sensorRecordMap.get(sensor.id)?.value }}
                  </div>
                  <div v-else class="bg-danger w-100 h-100">N.A.</div>
                </div>
                <div v-else>
                  <input type="number" :value="sensorRecordMap.get(sensor.id)?.value"
                    @input="(event) => onChangeRecordValue(event, date, sensor.id)" />
                </div>
              </td>
              <td>
                <div v-if="rowUpdateingIndex == date">
                  <span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>
                  修改中...
                </div>

                <button v-if="rowEditingIndex != date && rowUpdateingIndex != date" type="button"
                  class="btn btn-outline-primary btn-sm mx-2" @click="setEditing(date)">
                  编辑
                </button>

                <button v-if="rowEditingIndex == date && rowUpdateingIndex != date" type="button"
                  class="btn btn-outline-success btn-sm mx-2" @click="confirmEditing(date)">
                  确定
                </button>

                <button v-if="rowEditingIndex == date && rowUpdateingIndex != date" type="button"
                  class="btn btn-outline-secondary btn-sm mx-2" @click="cancelEditing(date)">
                  取消
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <table-paginator :offset="offsetVal" :limit="limitVal" :total="totalVal" @to-previous="toPreviousPage"
        @to-index="toOffset" @to-next="toNextPage" />

    </div>
  </div>
</template>

<script setup lang="ts">
import LineChart from '@/components/LineChart.vue'
import TablePaginator from '@/components/TablePaginator.vue'
import DoubleConfirmModal from '@/components/modal/DoubleConfirmModal.vue'
import UpsertSensorModal from '@/components/modal/UpsertSensorModal.vue'
import LinkSensorsModal from '@/components/modal/link-sensors-model/LinkSensorsModal.vue'
import httpclient, {
  getSensorDisplayText,
  type DataRecord,
  type Sensor,
  type Station
} from '@/http-client'
import { formatDatetime } from '@/utils/utils'
import moment from 'moment'
import { computed, nextTick, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Datepicker from 'vue3-datepicker'

const route = useRoute()
const router = useRouter()
const loading = ref(true)
const silentLoadingRecords = ref(false)
const isShowingDeletionModal = ref(false)
const operatingSensorId = ref<number | undefined>(undefined)
const startDateVal = new Date()
const endDateVal = new Date()
startDateVal.setHours(0, 0, 0, 0)
endDateVal.setHours(0, 0, 0, 0)
const startDate = ref(startDateVal)
const endDate = ref(endDateVal)
const displayedRecords = ref<DataRecord[]>([])

const station = ref<Station | undefined>(undefined)
const allSensors = ref<Sensor[]>([])
const selectedSensorIDs = ref<number[]>([])
const recordsByTime = ref<Map<Date, Map<number, DataRecord>> | undefined>(undefined)
const uploading = ref(false)
const file = ref<string | null>(null)
const isShowingUpsertSensorModal = ref(false)
const isShowingLinkSensorsModal = ref(false)
const disableSubmitButton = computed(() => file.value == null)
const upsertSensorModal = ref<InstanceType<typeof UpsertSensorModal> | null>(null)
const linkSensorModal = ref<InstanceType<typeof LinkSensorsModal> | null>(null)
const stationID = computed(() => {
  const stationID = route.query.station_id
  return stationID ? parseInt(stationID.toString()) : 0
})

// record pagination
const limitVal = ref(20)
const offsetVal = ref(0)
const totalVal = ref(0)

const selectedSensors = computed(() => {
  return allSensors.value
    .filter((sensor) => selectedSensorIDs.value.includes(sensor.id))
    .sort((a, b) => a.id - b.id)
})

// const positionName = (position: SensorPosition | undefined): string => {
//   return getPositionName(position)
// }
const gotoStationsView = () => {
  router.push({ query: { view: '站点管理' } })
}
// const viewSensorRecord = (sensorID: number) => {
//   router.push({
//     query: { view: 'Records', sensor_id: sensorID, station_id: stationID.value }
//   })
// }
const handleTemplateDownload = async () => {
  const csv = await httpclient.downloadTemplate()
  if (csv) {
    const anchor = document.createElement('a')
    anchor.href = 'data:text/csv;charset=utf-8,' + encodeURIComponent(csv)
    anchor.target = '_blank'
    anchor.download = 'template.csv'
    anchor.click()
  }
}
const selectCSVFile = (event: any) => {
  file.value = event.target.files[0]
}
const uploadCSVFile = async () => {
  if (file.value) {
    uploading.value = true
    const formData = new FormData()
    formData.append('file', file.value)
    try {
      const response = await httpclient.uploadCSVRecords(formData)
      console.log(response?.payload)
    } catch (error) {
      console.error(error)
    }
    uploading.value = false
    file.value = null
  }
  await updateSensorRecords(true)
}
const onChangeSensor = (evt: any, id: number) => {
  nextTick(() => {
    if (evt.target.checked) {
      selectedSensorIDs.value.push(id)
    } else {
      selectedSensorIDs.value = selectedSensorIDs.value.filter((sensorID) => sensorID != id)
    }
    offsetVal.value = 0
    updateSensorRecords(true)
  })
}
const updateSensorRecords = async (showLoading: boolean) => {
  if (showLoading) {
    loading.value = true
  } else {
    silentLoadingRecords.value = true
  }

  // clear editing cache
  rowEditingIndex.value = undefined
  rowUpdateingIndex.value = undefined
  recordsToBeUpdated = undefined
  recordsByTime.value = undefined
  displayedRecords.value = []

  const resp = await httpclient.getRecords({
    sensorIDs: selectedSensorIDs.value,
    startTime: startDate.value,
    endTime: endDate.value,
    offset: offsetVal.value,
    limit: limitVal.value
  })
  totalVal.value = resp?.total ?? 0
  const records = resp?.payload || []
  displayedRecords.value = records
  var result: Map<number, Map<number, DataRecord>> = new Map<number, Map<number, DataRecord>>()
  records.forEach((record) => {
    if (record.time) {
      const date = new Date(record.time).getTime()
      const sensorId = record?.sensor_id || 0
      var existingMap = result.get(date)
      if (existingMap) {
        existingMap.set(sensorId, record)
        result.set(date, existingMap)
      } else {
        const map: Map<number, DataRecord> = new Map<number, DataRecord>()
        map.set(sensorId, record)
        result.set(date, map)
      }
    }
  })
  const datetimes: number[] = Array.from(result.keys())

  var startDateTime: Date
  if (datetimes.length > 0) {
    startDateTime = new Date(datetimes[0])
  } else {
    startDateTime = startDate.value
  }
  console.log(startDateTime)
  const MS_PER_MINUTE = 60000
  for (let index = 0; index < limitVal.value; index++) {
    const temp = new Date(startDateTime.getTime() + 30 * index * MS_PER_MINUTE).getTime()
    const existingMap = result.get(temp)
    console.log(existingMap)
    if (!existingMap) {
      result.set(temp, new Map<number, DataRecord>())
    }
  }
  console.log(result)

  var sortedResult: Map<Date, Map<number, DataRecord>> = new Map<Date, Map<number, DataRecord>>()
  const fullDatetimes = Array.from(result.keys())
  fullDatetimes.sort((d1, d2) => d1 - d2)
  for (const date of fullDatetimes) {
    const data = result.get(date)
    if (data) {
      sortedResult.set(new Date(date), data)
    }
  }
  recordsByTime.value = sortedResult
  if (showLoading) {
    loading.value = false
  } else {
    silentLoadingRecords.value = true
  }
}

const refresh = async () => {
  loading.value = true
  station.value = undefined
  allSensors.value = []
  selectedSensorIDs.value = []

  endDate.value.setDate(startDate.value.getDate() + 1)
  station.value = (await httpclient.getStations())?.payload
    ?.filter((station) => station.id == stationID.value)
    .at(0)

  const resps = await Promise.all([
    httpclient.getStations(),
    httpclient.getSensors({ stationID: stationID.value })
  ])
  station.value = resps[0]?.payload.filter((station) => station.id == stationID.value).at(0)
  allSensors.value = resps[1]?.payload.sort((a, b) => a.id - b.id) || []
  selectedSensorIDs.value = []
  offsetVal.value = 0
  await updateSensorRecords(false)
  loading.value = false
}

// upsert sensor
const showCreateSensorModal = () => {
  upsertSensorModal.value?.setSensor(stationID.value)
  isShowingUpsertSensorModal.value = true
}

const onUpsertSensorModalClosed = () => {
  isShowingUpsertSensorModal.value = false
  operatingSensorId.value = undefined
}

const didUpsertSensor = () => {
  isShowingUpsertSensorModal.value = false
  operatingSensorId.value = undefined
  refresh()
}

const showSensorUpsertModal = (sensor: Sensor) => {
  operatingSensorId.value = sensor.id
  upsertSensorModal.value?.setSensor(stationID.value, sensor)
  isShowingUpsertSensorModal.value = true
}

const onDeleteSensor = (sensorId?: number) => {
  if (sensorId) {
    operatingSensorId.value = sensorId
    isShowingDeletionModal.value = true
  }
}

const onConfirmDeleteSensor = async () => {
  if (operatingSensorId.value) {
    try {
      await httpclient.deleteSensor(operatingSensorId.value)
    } catch (_) {
      // ignore
    }
  }
  await refresh()
  operatingSensorId.value = undefined
  isShowingDeletionModal.value = false
}

const onDeleteSensorModalClosed = () => {
  isShowingDeletionModal.value = false
  operatingSensorId.value = undefined
}

// link sensors

const showLinkSensorModal = () => {
  linkSensorModal.value?.setStation(station.value)
  isShowingLinkSensorsModal.value = true
}

const onConfirmLinkedSensors = async () => {
  isShowingLinkSensorsModal.value = false
  await refresh()
}

const onCloseLinkedSensors = () => {
  isShowingLinkSensorsModal.value = false
}

// record pagination

const toPreviousPage = (offset: number) => {
  offsetVal.value = offset
  updateSensorRecords(false)
}

const toNextPage = (offset: number) => {
  offsetVal.value = offset
  updateSensorRecords(false)
}

const toOffset = (offset: number) => {
  offsetVal.value = offset
  updateSensorRecords(false)
}

const onLimitChanged = (e: any) => {
  offsetVal.value = 0
  limitVal.value = e.target.value
  updateSensorRecords(false)
}

// date picker closed

const onDatePickerClosed = () => {
  nextTick(() => {
    offsetVal.value = 0
    updateSensorRecords(true)
  })
}

// edit records
const rowEditingIndex = ref<Date | undefined>(undefined)
const rowUpdateingIndex = ref<Date | undefined>(undefined)
var recordsToBeUpdated: Map<number, DataRecord> | undefined = undefined

const setEditing = (date: Date) => {
  rowEditingIndex.value = date
  if (recordsByTime.value) {
    recordsToBeUpdated = new Map(recordsByTime.value.get(date))
  }
}

const confirmEditing = async (date: Date) => {
  rowEditingIndex.value = undefined
  if (recordsByTime.value && recordsToBeUpdated) {
    recordsByTime.value.set(date, recordsToBeUpdated)
    rowUpdateingIndex.value = date
    const map = recordsToBeUpdated
    if (map) {
      const records = Array.from(map.values())
      const result = await httpclient.updateRecords(records)
      if (result?.code == 200) {
        recordsByTime.value.set(date, recordsToBeUpdated)
      }
      rowUpdateingIndex.value = undefined
    }
  }
}

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const cancelEditing = (date: Date) => {
  rowEditingIndex.value = undefined
}

const onChangeRecordValue = (e: any, date: Date, sensorId: number) => {
  const val = e.target.value
  console.log(val, date, sensorId)
  const existingRecord = recordsByTime.value?.get(date)?.get(sensorId)
  if (existingRecord && recordsToBeUpdated) {
    recordsToBeUpdated.set(sensorId, {
      ...existingRecord,
      value: +val
    })
  } else {
    if (!recordsToBeUpdated) {
      recordsToBeUpdated = new Map<number, DataRecord>()
    }
    recordsToBeUpdated.set(sensorId, {
      id: 0,
      sensor_id: sensorId,
      time: date,
      value: +val
    })
  }
}

const isDownloadData = ref(false)

const handleDownloadData = async () => {
  if (isDownloadData.value) return
  isDownloadData.value = true
  const resp = await httpclient.getRecords({
    sensorIDs: selectedSensorIDs.value,
    startTime: startDate.value,
    endTime: endDate.value,
    offset: 0,
    limit: totalVal.value
  })
  const records = resp?.payload || []
  var result: Map<Date, Map<number, DataRecord>> = new Map<Date, Map<number, DataRecord>>()
  records.forEach((record) => {
    const date = record.time
    if (date != null) {
      const sensorId = record?.sensor_id || 0
      var existingMap = result.get(date)
      if (existingMap) {
        existingMap.set(sensorId, record)
        result.set(date, existingMap)
      } else {
        const map: Map<number, DataRecord> = new Map<number, DataRecord>()
        map.set(sensorId, record)
        result.set(date, map)
      }
    }
  })

  // insert csv title row
  var csv = '数据时间,'
  const sensors = selectedSensors.value
  csv += sensors.map((sensor) => getSensorDisplayText(sensor)).join(',')
  csv += '\n'

  var times: Date[] = Array.from(result.keys()).sort(
    (d1, d2) => new Date(d1).getTime() - new Date(d2).getTime()
  )
  for (const time of times) {
    // 2023/4/1 15:30:00
    csv += moment(time).format('YYYY/M/D HH:mm:ss,')
    for (const [index, sensor] of sensors.entries()) {
      const data = result.get(time)?.get(sensor.id)?.value ?? ''
      csv += data
      if (index < sensors.length - 1) {
        csv += ','
      }
    }
    csv += '\n'
  }
  if (csv) {
    const anchor = document.createElement('a')
    anchor.href = 'data:text/csv;charset=utf-8,' + encodeURIComponent(csv)
    anchor.target = '_blank'
    anchor.download = `${startDate.value.toDateString()} - ${endDate.value.toDateString()}.csv`
    anchor.click()
  }
  isDownloadData.value = false
}

// setup

refresh()
</script>

<style scoped>
.tableFixHead {
  overflow: auto;
}

.tableFixHead thead th {
  position: sticky;
  top: 0;
  z-index: 1;
}
</style>
@/http-client
