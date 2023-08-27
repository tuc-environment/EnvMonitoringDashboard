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
              <div class="h4 mb-0">传感器列表</div>
              <label class="text-secondary small mb-3">选择查看相应的传感器</label>

              <div style="height: 170px; overflow-y: auto">
                <div class="form-check" v-for="sensor in allSensors" :key="sensor.id">
                  <input
                    class="form-check-input"
                    type="checkbox"
                    @change="onChangeSensor($event, sensor.id)"
                  />
                  <label class="form-check-label">
                    {{ sensor.id }}. {{ sensor.name }} ({{ positionName(sensor.position) }})
                  </label>
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
    </div>
  </div>
</template>

<script lang="ts">
import Datepicker from 'vue3-datepicker'
import httpclient, {
  type Station,
  type Sensor,
  SensorPosition,
  getPositionName
} from '@/httpclient'

export default {
  components: {
    Datepicker
  },
  async created() {
    this.loading = true
    this.startDate.setDate(this.startDate.getDate() - 1)
    this.station = (await httpclient.getStations())?.payload
      ?.filter((station) => station.id == this.stationID)
      .at(0)

    const resps = await Promise.all([
      httpclient.getStations(),
      httpclient.getSensors(this.stationID)
    ])
    this.station = resps[0]?.payload.filter((station) => station.id == this.stationID).at(0)
    this.allSensors = resps[1]?.payload.sort((a, b) => a.id - b.id) || []
    this.selectedSensorIDs = []
    await this.updateSensorRecords()
    this.loading = false
  },
  data() {
    return {
      loading: true,
      startDate: new Date(),
      endDate: new Date(),
      station: undefined as Station | undefined,
      allSensors: [] as Sensor[],
      selectedSensorIDs: [] as number[],
      allDataRecords: [] as string[][],
      uploading: false,
      requesting: false,
      file: null as string | null
    }
  },
  methods: {
    positionName(position: SensorPosition | undefined): string {
      return getPositionName(position)
    },
    gotoStationsView() {
      this.$router.push({ query: { view: 'Stations' } })
    },
    viewSensorRecord(sensorID: number) {
      this.$router.push({
        query: { view: 'Records', sensor_id: sensorID, station_id: this.stationID }
      })
    },
    async handleTemplateDownload() {
      const csv = await httpclient.downloadTemplate()
      if (csv) {
        const anchor = document.createElement('a')
        anchor.href = 'data:text/csv;charset=utf-8,' + encodeURIComponent(csv)
        anchor.target = '_blank'
        anchor.download = 'template.csv'
        anchor.click()
      }
    },
    selectCSVFile(event: any) {
      this.file = event.target.files[0]
    },
    async uploadCSVFile() {
      if (this.file) {
        this.uploading = true
        const formData = new FormData()
        formData.append('file', this.file)
        try {
          const response = await httpclient.uploadCSVRecords(formData)
          console.log(response?.payload)
        } catch (error) {
          console.error(error)
        }
        this.uploading = false
      }
    },
    onChangeSensor(evt: any, id: number) {
      this.$nextTick(() => {
        if (evt.target.checked) {
          this.selectedSensorIDs.push(id)
        } else {
          this.selectedSensorIDs = this.selectedSensorIDs.filter((sensorID) => sensorID != id)
        }
        this.updateSensorRecords()
      })
    },
    async updateSensorRecords() {
      this.loading = true
      const resp = await httpclient.getRecords(
        this.selectedSensorIDs,
        this.startDate,
        this.endDate,
        0,
        100
      )

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

      this.allDataRecords = Object.keys(recordsByTime).map((t) => [
        t,
        ...this.selectedSensors.map((s) => recordsByTime[t][s.id])
      ])
      this.loading = false
    }
  },
  computed: {
    disableSubmitButton(): boolean {
      return this.file == null
    },
    stationID(): number {
      const stationID = this.$route.query.station_id
      return stationID ? parseInt(stationID.toString()) : 0
    },
    selectedSensors(): Sensor[] {
      return this.allSensors
        .filter((sensor) => this.selectedSensorIDs.includes(sensor.id))
        .sort((a, b) => a.id - b.id)
    }
  }
}
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
