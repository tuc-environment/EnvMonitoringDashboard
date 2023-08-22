<template>
  <div class="homePageContainer">
    <MapContainer @did-select-station="selectStationHandler" />
    <Dashboard ref="dashboard" :selected-station="selectedStation" />
  </div>
</template>

<script setup lang="ts">
import MapContainer from '@/components/MapContainer.vue'
import Dashboard from '@/components/Dashboard.vue'
import { type Station } from '@/httpclient'
import { ref } from 'vue'

const selectedStation = ref<Station | undefined>(undefined)
const dashboard = ref<InstanceType<typeof Dashboard> | null>(null)

const selectStationHandler = (station: Station | undefined) => {
  selectedStation.value = station
  dashboard.value?.selectStation(station)
}
</script>

<style scoped>
.homePageContainer {
  background-color: black;
}
</style>
