<template>
  <div class="p-3 border rounded my-2" style="max-width: 500px">
    <div class="d-flex flex-wrap">
      <div class="mx-2 my-2" style="width: 200px">
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
      </div>

      <div class="mx-2 my-2" style="width: 200px">
        <h4>传感器信息</h4>
        <table class="table">
          <tr>
            <td width="20%">名称:</td>
            <td>{{ sensor?.name }}</td>
          </tr>
          <tr>
            <td>位置:</td>
            <td>{{ sensor?.position }}</td>
          </tr>
          <tr>
            <td>单位:</td>
            <td>{{ sensor?.unit }}</td>
          </tr>
          <tr>
            <td>组:</td>
            <td>{{ sensor?.group }}</td>
          </tr>
        </table>
      </div>

      <div class="d-flex">
        <button
          type="button"
          class="btn btn-sm btn-outline-primary mx-2"
          @click="gotoStationsView()"
        >
          查看所有站点
        </button>
        <button
          type="button"
          class="btn btn-sm btn-outline-primary mx-2"
          @click="gotoSensorsView()"
        >
          查看所有传感器
        </button>
      </div>
    </div>
  </div>

  <table class="table table-bordered align-middle">
    <thead class="table-dark">
      <tr>
        <th scope="col" width="20%">时间</th>
        <th scope="col" width="70%">{{ sensor?.name }}</th>
        <th scope="col" width="10%">操作</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="dataRecord in allDataRecords" :key="dataRecord.id">
        <td>{{ dataRecord.time }}</td>
        <td>{{ dataRecord.value }} {{ sensor?.unit }}</td>
        <td>
          <button type="button" class="btn btn-outline-primary btn-sm mx-2" disabled>编辑</button>
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script lang="ts">
import httpclient, { type Station, type Sensor, type DataRecord } from '@/httpclient'

export default {
  async created() {
    const resps = await Promise.all([
      httpclient.getStations(),
      httpclient.getSensors(this.stationID)
    ])
    this.station = resps[0]?.payload.filter((station) => station.id == this.stationID).at(0)
    this.sensor = resps[1]?.payload.filter((sensor) => sensor.id == this.sensorID).at(0)

    const resp = await httpclient.getRecords(
      [this.sensorID],
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
      sensor: undefined as Sensor | undefined,
      allDataRecords: [] as DataRecord[]
    }
  },
  methods: {
    gotoStationsView() {
      this.$router.push({ query: { view: 'Stations' } })
    },
    gotoSensorsView() {
      this.$router.push({ query: { view: 'Sensors', station_id: this.stationID } })
    }
  },
  computed: {
    stationID(): number {
      const stationID = this.$route.query.station_id
      return stationID ? parseInt(stationID.toString()) : 0
    },
    sensorID(): number {
      const sensorID = this.$route.query.sensor_id
      return sensorID ? parseInt(sensorID.toString()) : 0
    }
  }
}
</script>

<!-- <template>
  <div>
    <div class="input-group">
      <button
        class="btn btn-outline-secondary"
        type="button"
        id="inputGroupFileAddon04"
        @click="handleTemplateDownload"
      >
        Download Template
      </button>
      <input
        type="file"
        class="form-control"
        id="inputGroupFile04"
        aria-describedby="inputGroupFileAddon04"
        aria-label="Upload"
        @change="selectCSVFile"
      />
      <button
        class="btn btn-outline-primary"
        type="button"
        id="inputGroupFileAddon04"
        :disabled="disableSubmitButton"
        @click="uploadCSVFile"
      >
        Submit
      </button>
    </div>

    <PagedTableComponent />
  </div>
</template>

<script setup lang="ts">
import httpclient, { type DataRecord } from '@/httpclient'
import PagedTableComponent from '@/components/PagedTableComponent.vue'
import { computed, reactive, ref, onMounted } from 'vue'

const uploading = ref(false)
const requesting = ref(false)
const file = ref<string | null>(null)
const disableSubmitButton = computed((): boolean => {
  return file.value == null
})
const records = reactive<{
  allRecords: DataRecord[]
}>({
  allRecords: []
})

onMounted(async () => {
  requesting.value = true
  const resp = await httpclient.getRecords(
    [1, 2],
    new Date('2023-04-02T14:00:00+08:00'),
    new Date('2023-04-02T14:30:00+08:00'),
    0,
    5
  )
  if (resp?.code == 200) {
    records.allRecords = resp.payload
    console.log(JSON.stringify(records.allRecords))
  } else {
  }
  requesting.value = false
})

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
</script> -->
