<template>
  <div class="dashboardContainer">
    <div class="leftPanel column">
      <TreeChart ref="treeChart" />
      <LineChart
        class="treeSelectedChart"
        :station="dashboardStore.$state.treeStationSelected"
        :sensors="
          dashboardStore.$state.treeSensorSelected ? [dashboardStore.$state.treeSensorSelected] : []
        "
        :records="dashboardStore.$state.treeSensorRecordsLoaded"
      />
    </div>
    <div class="centerSpace"></div>
    <div class="rightPanel column">
      <LineChart
        class="chart"
        :station="selectedStation"
        :sensors="relatedSensors"
        :records="relatedRecords"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import LineChart from '@/components/LineChart.vue'
import TreeChart from '@/components/tree/Tree.vue'
import httpclient, { type Station, type Sensor, type DataRecord } from '@/httpclient'
import { ref, type PropType, onMounted } from 'vue'
import { useDashboardStore } from '@/stores/dashboard'

const dashboardStore = useDashboardStore()
const treeChart = ref<InstanceType<typeof TreeChart> | null>(null)

const stations = ref<Station[] | undefined>(undefined)
const selectedStation = ref<Station | undefined>(undefined)
const relatedSensors = ref<Sensor[] | undefined>(undefined)
const relatedRecords = ref<DataRecord[] | undefined>(undefined)

const selectStation = async (station: Station | undefined) => {
  console.log('[dashboard] select station: ', JSON.stringify(station))
  if (selectedStation.value?.id && selectedStation.value?.id == station?.id) {
    return
  }
  selectedStation.value = station
  relatedSensors.value = []
  relatedRecords.value = []
  const stationId = station?.id
  if (stationId) {
    const sensorsRes = await httpclient.getSensors(stationId)
    const sensors = sensorsRes?.payload ?? []
    if (sensors.length > 0) {
      const sensorIds = sensors
        ? (sensors.map((sensors) => sensors.id).filter((id) => id) as number[])
        : []
      const recordsRes = await httpclient.getRecords(sensorIds)
      const records = recordsRes?.payload ?? []
      relatedSensors.value = sensors
      relatedRecords.value = records
      console.log('[dashboard] get records count: ', records.length)
    }
  }
}

const setStations = (updatedStations: Station[]) => {
  stations.value = updatedStations
  treeChart.value?.setStations(updatedStations)
}

defineExpose({
  selectStation,
  setStations
})
</script>

<style scoped>
.dashboardContainer {
  position: absolute;
  top: 0;
  left: 0;
  display: flex;
  flex-direction: row;
  padding: 0px;
  margin: 0px;
  width: 100%;
  height: calc(100vh);
  pointer-events: none;
}

.leftPanel {
  width: 200px;
  min-width: 200px;
  height: 100%;
  flex: 1;
}

.centerSpace {
  flex: 1.7;
  min-width: 200px;
}

.rightPanel {
  width: 200px;
  min-width: 200px;
  height: 100%;
  flex: 1;
}
</style>
