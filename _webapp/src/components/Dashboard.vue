<template>
  <div class="dashboardContainer">
    <div class="leftPanel d-flex flex-column justify-content-between">
      <!-- <tag-group :tags="dashboardStore.treeSensorSelectedTags" @on-tag-selected="onSelectTag" /> -->
      <TreeChart ref="treeChart" />
      <div class="space"></div>
      <LineChart
        :stations="dashboardStore.$state.treeStationsSelected"
        :sensors="dashboardStore.$state.treeSensorsSelected"
        :records="dashboardStore.$state.treeSensorRecordsLoaded"
        :loading="dashboardStore.$state.treeRecordsLoading"
        :maxLines="10"
        title="数据对比"
        default-text="请选择数据项"
        no-data-text="暂无数据"
      />
    </div>
    <div class="centerSpace"></div>
    <div class="rightPanel d-flex flex-column justify-content-between">
      <LineChart
        :sensors="airRelatedSensors"
        :records="airRelatedRecords"
        :show-default-text="!selectedStation"
        :loading="loadingDataForStation"
        :title="selectedStation ? `${selectedStation.name} - 空气参数` : '空气参数'"
        default-text="请选择站点"
        no-data-text="暂无数据"
      />
      <LineChart
        :sensors="soilRelatedSensors"
        :records="soilRelatedRecords"
        :show-default-text="!selectedStation"
        :loading="loadingDataForStation"
        :title="selectedStation ? `${selectedStation.name} - 土壤参数` : '土壤参数'"
        default-text="请选择站点"
        no-data-text="暂无数据"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import LineChart from '@/components/LineChart.vue'
import TreeChart from '@/components/tree/Tree.vue'
import httpclient, { type Station, type Sensor, type DataRecord } from '@/httpclient'
import { ref, type PropType, onMounted } from 'vue'
import { useDashboardStore } from '@/stores/DashboardStore'
import { airOptionNames, soilOptionNames } from '@/utils/constants'
import TagGroup from '@/components/tags/TagsGroup.vue'
import { type TagData } from './tags/TagData'

const dashboardStore = useDashboardStore()
const treeChart = ref<InstanceType<typeof TreeChart> | null>(null)

const loadingDataForStation = ref(false)
const stations = ref<Station[] | undefined>(undefined)
const selectedStation = ref<Station | undefined>(undefined)
const airRelatedSensors = ref<Sensor[] | undefined>(undefined)
const airRelatedRecords = ref<DataRecord[] | undefined>(undefined)
const soilRelatedSensors = ref<Sensor[] | undefined>(undefined)
const soilRelatedRecords = ref<DataRecord[] | undefined>(undefined)

const selectStation = async (station: Station | undefined) => {
  console.log('[dashboard] select station: ', JSON.stringify(station))
  if (selectedStation.value?.id && selectedStation.value?.id == station?.id) {
    return
  }
  loadingDataForStation.value = true
  selectedStation.value = station
  airRelatedSensors.value = []
  airRelatedRecords.value = []
  soilRelatedSensors.value = []
  soilRelatedRecords.value = []
  const stationId = station?.id
  if (stationId) {
    const sensorsRes = await httpclient.getSensors({ stationID: stationId })
    const sensors = sensorsRes?.payload ?? []
    if (sensors.length > 0) {
      const airSensorIds = sensors
        ? (sensors
            .filter((sensor) => sensor.id && sensor.name && airOptionNames.includes(sensor.name))
            .map((sensor) => sensor.id) as number[])
        : []
      const soilSensorIds = sensors
        ? (sensors
            .filter((sensor) => sensor.id && sensor.name && soilOptionNames.includes(sensor.name))
            .map((sensor) => sensor.id) as number[])
        : []
      const recordsRes = await httpclient.getRecords({
        sensorIDs: airSensorIds.concat(soilSensorIds)
      })
      const records = recordsRes?.payload ?? []
      airRelatedSensors.value = sensors.filter((sensor) => airSensorIds.includes(sensor.id))
      soilRelatedSensors.value = sensors.filter((sensor) => soilSensorIds.includes(sensor.id))
      airRelatedRecords.value = records.filter((record) => airSensorIds.includes(record.sensor_id))
      soilRelatedRecords.value = records.filter((record) =>
        soilSensorIds.includes(record.sensor_id)
      )
      console.log('[dashboard] get records count: ', records.length)
    }
  }
  loadingDataForStation.value = false
}

const setStations = (updatedStations: Station[]) => {
  stations.value = updatedStations
  // treeChart.value?.setStations(updatedStations)
}

const onSelectTag = (tag: TagData) => {
  const sensor = tag.data as Sensor
  if (sensor) {
    dashboardStore.removeTreeNodeSelected(sensor)
  }
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
  min-width: 200px;
  height: 100%;
  flex: 1;
  overflow: hidden;
}

.centerSpace {
  flex: 1.5;
  min-width: 200px;
}

.rightPanel {
  min-width: 200px;
  height: 100%;
  flex: 1;
}

.space {
  flex-grow: 999;
}
</style>
