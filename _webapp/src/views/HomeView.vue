<template>
  <div class="homePageContainer">
    <MapContainer ref="mapContainer" @did-select-station="selectStationHandler" />
    <Dashboard ref="dashboard" :selected-station="selectedStation" />
  </div>
</template>

<script setup lang="ts">
import MapContainer from '@/components/map/MapContainer.vue'
import Dashboard from '@/components/Dashboard.vue'
import httpclient, { type Station } from '@/httpclient'
import { onMounted, ref } from 'vue'
import { toast, updateGlobalOptions } from 'vue3-toastify'
import 'vue3-toastify/dist/index.css'

const requesting = ref(false)
const stations = ref<Station[] | undefined>(undefined)
const selectedStation = ref<Station | undefined>(undefined)
const dashboard = ref<InstanceType<typeof Dashboard> | null>(null)
const mapContainer = ref<InstanceType<typeof MapContainer> | null>(null)

onMounted(async () => {
  await loadStations()
  toast('点击地图空白区域建立预测点', {
    position: toast.POSITION.TOP_CENTER,
    autoClose: 2500,
    theme: toast.THEME.DARK
  })
})

const selectStationHandler = (station: Station | undefined) => {
  selectedStation.value = station
  dashboard.value?.selectStation(station)
}

const loadStations = async () => {
  requesting.value = true
  const resp = await httpclient.getStations()
  if (resp?.code == 200) {
    stations.value = resp.payload
  } else {
  }
  requesting.value = false
  mapContainer.value?.setStations(stations.value ?? [])
  dashboard.value?.setStations(stations.value ?? [])
}
</script>

<style scoped>
.homePageContainer {
  background-color: black;
}
</style>
