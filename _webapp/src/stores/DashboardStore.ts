import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import httpclient, {
  type Station,
  type Sensor,
  type DataRecord,
  getSensorDisplayText
} from '@/httpclient'
import { type TagData } from '@/components/tags/TagData'
export const useDashboardStore = defineStore('dashboard', () => {
  const treeSensorsSelected = ref<Sensor[]>([])
  const treeSensorSelectedTags = computed<TagData[]>((): TagData[] =>
    treeSensorsSelected.value.map((sensor: Sensor): TagData => {
      const stationName = treeStationsSelected.value.find(
        (station) => station.id == sensor.station_id
      )?.name
      return { title: getSensorDisplayText(sensor, stationName), data: sensor }
    })
  )
  const treeStationsSelected = ref<Station[]>([])
  const treeSensorRecordsLoaded = ref<DataRecord[]>([])
  const treeRecordsLoading = ref(false)
  const addTreeNodeSelected = (sensor: Sensor, station: Station) => {
    var sensors = treeSensorsSelected.value
    const selectedSensorIds = sensors.map((sensor) => sensor.id)
    if (selectedSensorIds.includes(sensor.id)) {
      return
    }
    sensors.push(sensor)
    treeSensorsSelected.value = sensors
    var stations = treeStationsSelected.value
    stations.push(station)
    treeStationsSelected.value = stations
    loadRecordsForSelectedTreeSensor()
  }

  const removeTreeNodeSelected = (sensor: Sensor) => {
    const sensors = treeSensorsSelected.value
    const stations = treeStationsSelected.value
    const records = treeSensorRecordsLoaded.value
    const updatedSensors = sensors.filter((s) => s.id != sensor.id)
    const stationIds = [...new Set(updatedSensors.map((sensor) => sensor.station_id))]
    const updatedStations = stations.filter((station) => stationIds.includes(station.id))
    const updatedRecords = records.filter((record) => record.sensor_id != sensor.id)
    treeSensorsSelected.value = updatedSensors
    treeStationsSelected.value = updatedStations
    treeSensorRecordsLoaded.value = updatedRecords
  }

  const loadRecordsForSelectedTreeSensor = async () => {
    treeRecordsLoading.value = true
    const sensors = treeSensorsSelected.value
    if (sensors.length > 0) {
      try {
        const result = await httpclient.getRecords(sensors.map((sensor) => sensor.id))
        treeSensorRecordsLoaded.value = result?.payload ?? []
      } catch (e) {
        console.log('[dashboard-store]', e)
      }
    }
    treeRecordsLoading.value = false
  }

  return {
    treeRecordsLoading,
    treeSensorsSelected,
    treeStationsSelected,
    treeSensorRecordsLoaded,
    treeSensorSelectedTags,
    addTreeNodeSelected,
    removeTreeNodeSelected
  }
})
