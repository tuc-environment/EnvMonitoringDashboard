<template>
  <table class="table table-bordered">
    <thead class="table-dark">
      <tr>
        <th scope="col" width="10%">#</th>
        <th scope="col" width="30%">Name</th>
        <th scope="col" width="20%">Latitude</th>
        <th scope="col" width="20%">Longitude</th>
        <th scope="col" width="20%">Altitude</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="station in allStations" :key="station.id">
        <td>{{ station.id }}</td>
        <td>{{ station.name }}</td>
        <td>{{ station.lat }}</td>
        <td>{{ station.lng }}</td>
        <td>{{ station.altitude }}</td>
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
  }
}
</script>
