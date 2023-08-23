<template>
  <div class="row">
    <div class="m-2 p-3 border rounded col-md-4">
      <h4>站点信息</h4>
      <table class="table">
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

      <button type="button" class="btn btn-sm btn-outline-primary" @click="gotoStationsView()">
        查看所有站点
      </button>
    </div>
    <div class="m-2 p-3 border rounded col-md-4">
      <h4>上传数据</h4>
      <button class="btn btn-sm btn-outline-primary my-2" @click="handleTemplateDownload">
        下载数据模版
      </button>

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

  <div class="row">
    <div class="col-md-3 tableFixHead" style="height: 400px">
      <table class="table align-middle">
        <thead class="table-dark">
          <tr>
            <th scope="col" width="10%">#</th>
            <th scope="col" width="30%">名称</th>
            <th scope="col" width="10%">位置</th>
            <th scope="col" width="10%">组</th>
            <th scope="col" width="30%">操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="sensor in allSensors" :key="sensor.id">
            <td>{{ sensor.id }}</td>
            <td>{{ sensor.name }}</td>
            <td>{{ sensor.position }}</td>
            <td>{{ sensor.group }}</td>
            <td>
              <button type="button" class="btn btn-outline-primary btn-sm mx-2" disabled>
                编辑
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="col-md-9">
      <table class="table table-bordered align-middle">
        <thead class="table-dark">
          <tr>
            <th scope="col" width="20%">时间</th>
            <th scope="col" width="70%">值</th>
            <th scope="col" width="10%">操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="dataRecord in allDataRecords" :key="dataRecord.id">
            <td>{{ dataRecord.time }}</td>
            <td>{{ dataRecord.value }}</td>
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
import httpclient, { type Station, type Sensor, type DataRecord } from '@/httpclient'

export default {
  async created() {
    this.station = (await httpclient.getStations())?.payload
      ?.filter((station) => station.id == this.stationID)
      .at(0)

    const resps = await Promise.all([
      httpclient.getStations(),
      httpclient.getSensors(this.stationID)
    ])
    this.station = resps[0]?.payload.filter((station) => station.id == this.stationID).at(0)
    this.allSensors = resps[1]?.payload || []
    const resp = await httpclient.getRecords(
      [8],
      new Date('2023-04-02T14:00:00+08:00'),
      new Date('2023-04-02T14:30:00+08:00'),
      0,
      100
    )
    this.allDataRecords = resp?.payload || []
  },
  data() {
    return {
      station: undefined as Station | undefined,
      allSensors: [] as Sensor[],
      allDataRecords: [] as DataRecord[],
      uploading: false,
      requesting: false,
      file: null as string | null
    }
  },
  methods: {
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
    }
  },
  computed: {
    disableSubmitButton(): boolean {
      return this.file == null
    },
    stationID(): number {
      const stationID = this.$route.query.station_id
      return stationID ? parseInt(stationID.toString()) : 0
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
