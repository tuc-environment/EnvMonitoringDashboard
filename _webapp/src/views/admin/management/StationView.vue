<template>
  <table class="table table-bordered align-middle">
    <thead class="table-dark">
      <tr>
        <th scope="col" width="10%">#</th>
        <th scope="col" width="30%">名称</th>
        <th scope="col" width="10%">纬度</th>
        <th scope="col" width="10%">经度</th>
        <th scope="col" width="10%">海拔</th>
        <th scope="col" width="30%">操作</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="station in allStations" :key="station.id">
        <td>{{ station.id }}</td>
        <td>{{ station.name }}</td>
        <td>{{ station.lat }}</td>
        <td>{{ station.lng }}</td>
        <td>{{ station.altitude }}</td>
        <td>
          <button
            type="button"
            class="btn btn-outline-success btn-sm mx-2"
            @click="viewSensorData(station.id)"
          >
            查看所有传感器
          </button>
          <button type="button" class="btn btn-outline-primary btn-sm mx-2" disabled>编辑</button>
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script lang="ts">
import httpclient, { type Station } from '@/httpclient'

export default {
  components: {},
  async created() {
    const resp = await httpclient.getStations()
    this.allStations = resp?.payload || []
  },
  data() {
    return {
      requesting: true,
      allStations: [] as Station[]
    }
  },
  methods: {
    viewSensorData(stationID: number) {
      this.$router.push({ query: { view: 'Sensors', station_id: stationID } })
    }
  }
}
</script>
