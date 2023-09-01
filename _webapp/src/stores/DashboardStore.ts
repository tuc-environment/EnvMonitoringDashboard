import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import httpclient, {
  type Station,
  type Sensor,
  type DataRecord,
  getSensorDisplayText
} from '@/httpclient'
import { type TagData } from '@/components/tags/TagData'
import { airOptionNames, soilOptionNames } from '@/utils/constants'

export const useDashboardStore = defineStore('dashboard', () => {
  // tree
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

  // stations
  const stations = ref<Station[]>([])
  const loadingStations = ref(false)
  const selectedStation = ref<Station | undefined>(undefined)

  // selected station data
  const loadingDataForStation = ref(false)
  const airRelatedSensors = ref<Sensor[] | undefined>(undefined)
  const airRelatedRecords = ref<DataRecord[] | undefined>(undefined)
  const soilRelatedSensors = ref<Sensor[] | undefined>(undefined)
  const soilRelatedRecords = ref<DataRecord[] | undefined>(undefined)

  // counts

  const totalStations = ref<number | undefined>(undefined)
  const totalSensors = ref<number | undefined>(undefined)
  const totalRecords = ref<number | undefined>(undefined)
  const totalRecordsToday = ref<number | undefined>(undefined)

  // tree related actions
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
        const result = await httpclient.getRecords({ sensorIDs: sensors.map((sensor) => sensor.id), })
        treeSensorRecordsLoaded.value = result?.payload ?? []
      } catch (e) {
        console.log('[dashboard-store]', e)
      }
    }
    treeRecordsLoading.value = false
  }

  // station releated actions

  const loadStations = async () => {
    loadingStations.value = true
    try {
      const result = await httpclient.getStations()
      totalStations.value = result?.total
      stations.value = result?.payload ?? []
    } catch (_) { }
    loadingStations.value = false
  }

  const loadTotalCounts = async () => {
    try {
      const sensorRes = await httpclient.getSensors({ limit: 0, })
      totalSensors.value = sensorRes?.total
      const recordRes = await httpclient.getRecords({ limit: 0, })
      totalRecords.value = recordRes?.total
    } catch (_) { }
  }

  const setMapSelectedStation = async (station: Station | undefined) => {
    console.log('[dashboard] select station: ', JSON.stringify(station))
    if (
      selectedStation.value &&
      selectedStation.value?.id &&
      selectedStation.value?.id == station?.id
    ) {
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
        const recordsRes = await httpclient.getRecords({ sensorIDs: airSensorIds.concat(soilSensorIds) })
        const records = recordsRes?.payload ?? []
        airRelatedSensors.value = sensors.filter((sensor) => airSensorIds.includes(sensor.id))
        soilRelatedSensors.value = sensors.filter((sensor) => soilSensorIds.includes(sensor.id))
        airRelatedRecords.value = records.filter((record) =>
          airSensorIds.includes(record.sensor_id)
        )
        soilRelatedRecords.value = records.filter((record) =>
          soilSensorIds.includes(record.sensor_id)
        )
        console.log('[dashboard] get records count: ', records.length)
      }
    }
    loadingDataForStation.value = false
  }

  return {
    treeRecordsLoading,
    treeSensorsSelected,
    treeStationsSelected,
    treeSensorRecordsLoaded,
    treeSensorSelectedTags,
    loadingStations,
    stations,
    loadingDataForStation,
    selectedStation,
    airRelatedSensors,
    airRelatedRecords,
    soilRelatedSensors,
    soilRelatedRecords,
    totalRecords,
    totalSensors,
    totalStations,
    totalRecordsToday,
    loadTotalCounts,
    addTreeNodeSelected,
    removeTreeNodeSelected,
    loadStations,
    setMapSelectedStation
  }
})
