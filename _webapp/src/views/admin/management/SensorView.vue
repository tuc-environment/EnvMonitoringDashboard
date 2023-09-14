<template>
  <div class="row align-items-stretch">
    <div class="col-md-4 my-2">
      <div class="card h-100">
        <div class="card-body">
          <div class="d-flex align-items-center">
            <div class="h4 my-0 me-2">站点信息</div>
            <button
              type="button"
              class="btn btn-sm btn-outline-primary"
              @click="gotoStationsView()"
            >
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
            <h4 class="my-0 me-2">上传数据</h4>
            <button class="btn btn-sm btn-outline-primary my-2" @click="handleTemplateDownload">
              下载数据模版
            </button>
          </div>

          <div class="input-group my-2">
            <input type="file" class="form-control" @change="selectCSVFile" />
            <button
              class="btn btn-outline-primary"
              type="button"
              :disabled="disableSubmitButton"
              @click="uploadCSVFile"
            >
              Submit
            </button>
          </div>
        </div>
      </div>
    </div>
    <div class="col-md-8 my-2">
      <div class="card h-100">
        <div class="card-body">
          <div class="row">
            <div class="col-md-6">
              <div class="d-flex justify-content-between">
                <div class="h4 mb-0">传感器列表</div>
                <button type="button" class="btn btn-primary" @click="showCreateSensorModal">
                  添加传感器
                </button>
              </div>
              <label class="text-secondary small mb-3">选择查看相应的传感器</label>

              <div style="height: 170px; overflow-y: auto">
                <div class="form-check mb-3 d-flex" v-for="sensor in allSensors" :key="sensor.id">
                  <input
                    class="form-check-input"
                    type="checkbox"
                    @change="onChangeSensor($event, sensor.id)"
                  />
                  <label class="form-check-label ms-1">
                    {{ getSensorDisplayText(sensor) }}
                  </label>
                  <div v-if="operatingSensorId == sensor.id" class="ms-2">
                    <span
                      class="spinner-grow spinner-grow-sm"
                      role="status"
                      aria-hidden="true"
                    ></span>
                    修改中...
                  </div>
                  <div v-else class="d-flex">
                    <button
                      type="button"
                      class="btn btn-outline-primary btn-sm mx-2"
                      @click="showSensorUpsertModal(sensor)"
                    >
                      编辑
                    </button>
                    <button
                      type="button"
                      class="btn btn-outline-danger btn-sm mx-2"
                      @click="onDeleteSensor(sensor.id)"
                    >
                      删除
                    </button>
                  </div>
                </div>
              </div>
            </div>
            <div class="col-md-6">
              <div class="h4 mb-0">时间</div>
              <label class="text-secondary small mb-3">选择时间范围</label>

              <div class="my-2">
                <div>开始时间</div>
                <Datepicker v-model="startDate" style="cursor: pointer" />
              </div>

              <div class="my-2">
                <div>结束时间</div>
                <Datepicker v-model="endDate" style="cursor: pointer" />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <upsert-sensor-modal
      ref="upsertSensorModal"
      :visible="isShowingUpsertSensorModal"
      @close="onUpsertSensorModalClosed"
      @did-upsert-sensor="didUpsertSensor"
    />
    <double-confirm-modal
      :visible="isShowingDeletionModal"
      :title="'删除传感器 ' + operatingSensorId"
      content="确定删除传感器？"
      @close="onDeleteSensorModalClosed"
      @on-confirm="onConfirmDeleteSensor"
    />
  </div>

  <div v-if="loading" class="mx-auto mt-5 text-center">
    <div class="spinner-border mx-auto" role="status">
      <span class="sr-only"></span>
    </div>
  </div>

  <div v-else-if="allDataRecords.length == 0" class="mx-auto mt-5 text-center">
    <label>无法查询到数据</label>
  </div>

  <div v-else class="row">
    <div class="col-md-12">
      <table class="table table-bordered align-middle">
        <thead class="table-dark">
          <tr>
            <th scope="col" width="30%">时间</th>
            <th scope="col" v-for="s in selectedSensors" :key="s.id">
              {{ s.id }}. {{ s.name }} ({{ positionName(s.position) }})
            </th>
            <th scope="col" width="20%">操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(dataRecords, idx) in allDataRecords" :key="idx">
            <td v-for="record in dataRecords" :key="record">{{ record }}</td>
            <td>
              <button type="button" class="btn btn-outline-primary btn-sm mx-2" disabled>
                编辑
              </button>
            </td>
          </tr>
        </tbody>
      </table>
      <table-paginator
        :offset="offsetVal"
        :limit="limit"
        :total="totalVal"
        @to-previous="toPreviousPage"
        @to-index="toOffset"
        @to-next="toNextPage"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import Datepicker from 'vue3-datepicker'
import httpclient, {
  type Station,
  type Sensor,
  SensorPosition,
  getPositionName,
  getSensorDisplayText
} from '@/httpclient'
import UpsertSensorModal from '@/components/modal/UpsertSensorModal.vue'
import { computed, nextTick, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import DoubleConfirmModal from '@/components/modal/DoubleConfirmModal.vue'
import TablePaginator from '@/components/TablePaginator.vue'
const route = useRoute()
const router = useRouter()
const loading = ref(true)
const silentLoadingRecords = ref(false)
const isShowingDeletionModal = ref(false)
const operatingSensorId = ref<number | undefined>(undefined)
const startDate = ref(new Date())
const endDate = ref(new Date())

const station = ref<Station | undefined>(undefined)
const allSensors = ref<Sensor[]>([])
const selectedSensorIDs = ref<number[]>([])
const allDataRecords = ref<string[][]>([])

const uploading = ref(false)
const file = ref<string | null>(null)
const isShowingUpsertSensorModal = ref(false)
const disableSubmitButton = computed(() => file.value == null)
const upsertSensorModal = ref<InstanceType<typeof UpsertSensorModal> | null>(null)
const stationID = computed(() => {
  const stationID = route.query.station_id
  return stationID ? parseInt(stationID.toString()) : 0
})

// record pagination
const limit = 10
const offsetVal = ref(0)
const totalVal = ref(0)

const selectedSensors = computed(() => {
  return allSensors.value
    .filter((sensor) => selectedSensorIDs.value.includes(sensor.id))
    .sort((a, b) => a.id - b.id)
})

const positionName = (position: SensorPosition | undefined): string => {
  return getPositionName(position)
}
const gotoStationsView = () => {
  router.push({ query: { view: 'Stations' } })
}
const viewSensorRecord = (sensorID: number) => {
  router.push({
    query: { view: 'Records', sensor_id: sensorID, station_id: stationID.value }
  })
}
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
  }
}
const onChangeSensor = (evt: any, id: number) => {
  nextTick(() => {
    if (evt.target.checked) {
      selectedSensorIDs.value.push(id)
    } else {
      selectedSensorIDs.value = selectedSensorIDs.value.filter((sensorID) => sensorID != id)
    }
    offsetVal.value = 0
    updateSensorRecords(false)
  })
}
const updateSensorRecords = async (showLoading: boolean) => {
  if (showLoading) {
    loading.value = true
  } else {
    silentLoadingRecords.value = true
  }
  const resp = await httpclient.getRecords({
    sensorIDs: selectedSensorIDs.value,
    startTime: startDate.value,
    endTime: endDate.value,
    offset: offsetVal.value,
    limit: limit
  })
  totalVal.value = resp?.total ?? 0
  const records = resp?.payload || []
  const recordsByTime = records.reduce((acc, record) => {
    const time = record?.time?.toString() || ''
    const sensorId = record?.sensor_id || 0
    const value = record?.value?.toString() || ''
    if (acc[time]) {
      acc[time][sensorId] = value
    } else {
      acc[time] = {}
      acc[time][sensorId] = value
    }
    return acc
  }, {} as { [time: string]: { [sensorId: number]: string } })
  console.log(recordsByTime)

  allDataRecords.value = Object.keys(recordsByTime).map((t) => [
    t,
    ...selectedSensors.value.map((s) => recordsByTime[t][s.id])
  ])
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
  allDataRecords.value = []

  startDate.value.setDate(startDate.value.getDate() - 1)
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
    } catch (_) {}
  }
  await refresh()
  operatingSensorId.value = undefined
  isShowingDeletionModal.value = false
}

const onDeleteSensorModalClosed = () => {
  isShowingDeletionModal.value = false
  operatingSensorId.value = undefined
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
