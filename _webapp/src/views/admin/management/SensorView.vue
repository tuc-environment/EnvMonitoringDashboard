<template>
  <div class="p-3 border rounded my-2" style="max-width: 420px">
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

  <table class="table table-bordered align-middle">
    <thead class="table-dark">
      <tr>
        <th scope="col" width="10%">#</th>
        <th scope="col" width="30%">名称</th>
        <th scope="col" width="10%">位置</th>
        <th scope="col" width="10%">单位</th>
        <th scope="col" width="10%">组</th>
        <th scope="col" width="30%">操作</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="sensor in allSensors" :key="sensor.id">
        <td>{{ sensor.id }}</td>
        <td>{{ sensor.name }}</td>
        <td>{{ sensor.position }}</td>
        <td>{{ sensor.unit }}</td>
        <td>{{ sensor.group }}</td>
        <td>
          <button
            type="button"
            class="btn btn-outline-success btn-sm mx-2"
            @click="viewSensorRecord(sensor.id)"
          >
            查看传感器数据
          </button>
          <button type="button" class="btn btn-outline-primary btn-sm mx-2" disabled>编辑</button>
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script lang="ts">
import httpclient, { type Station, type Sensor } from '@/httpclient'

export default {
  async created() {
    this.station = (await httpclient.getStations())?.payload
      ?.filter((station) => station.id == this.stationID)
      .at(0)
    const resp = await httpclient.getSensors(this.stationID)
    this.allSensors = resp?.payload || []
  },
  data() {
    return {
      station: undefined as Station | undefined,
      allSensors: [] as Sensor[]
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
    }
  },
  computed: {
    stationID(): number {
      const stationID = this.$route.query.station_id
      return stationID ? parseInt(stationID.toString()) : 0
    }
  }
}
</script>
