import { type TagData } from '@/components/tags/TagData'
import httpclient, {
  getSensorDisplayText,
  type DataRecord,
  type Sensor,
  type Station
} from '@/http-client'
import { airOptionNames, soilOptionNames } from '@/utils/constants'
import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

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
  const totalCreatedToday = ref<number | undefined>(undefined)

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
        const startDate = new Date()
        startDate.setDate(startDate.getDay() - 180)
        const result = await httpclient.getRecords({
          sensorIDs: sensors.map((sensor) => sensor.id),
          offset: 0,
          limit: 900000,
          startTime: startDate
        })
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
    } catch (_) {}
    loadingStations.value = false
  }

  const loadTotalCounts = async () => {
    try {
      const sensorRes = await httpclient.getSensors({ limit: 0, visibleInDashboard: true })
      totalSensors.value = sensorRes?.total
      const recordRes = await httpclient.getRecords({ limit: 0 })
      totalRecords.value = recordRes?.total
      const now = new Date()
      const todayStart = new Date(now.getFullYear(), now.getMonth(), now.getDay(), 0, 0)
      const todayEnd = new Date(now.getFullYear(), now.getMonth(), now.getDay(), 23, 59)
      const todayRes = await httpclient.getRecords({
        limit: 0,
        afterCreatedAt: todayStart,
        beforeCreatedAt: todayEnd
      })
      totalCreatedToday.value = todayRes?.total
    } catch (_) {}
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
      const sensorsRes = await httpclient.getSensors({
        stationID: stationId,
        visibleInDashboard: true
      })
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
        const startDate = new Date()
        startDate.setDate(startDate.getDay() - 180)
        console.log(`[dashboard-store] half year before: ${startDate}`)

        const airRecordRes = await httpclient.getRecords({
          sensorIDs: airSensorIds,
          offset: 0,
          limit: 900000,
          startTime: startDate
        })
        const airRecords = airRecordRes?.payload ?? []

        const soilRecordRes = await httpclient.getRecords({
          sensorIDs: soilSensorIds,
          offset: 0,
          limit: 900000,
          startTime: startDate
        })
        const soilRecords = soilRecordRes?.payload ?? []

        airRelatedSensors.value = sensors.filter((sensor) => airSensorIds.includes(sensor.id))
        soilRelatedSensors.value = sensors.filter((sensor) => soilSensorIds.includes(sensor.id))
        airRelatedRecords.value = airRecords.filter((record) =>
          airSensorIds.includes(record.sensor_id)
        )
        soilRelatedRecords.value = soilRecords.filter((record) =>
          soilSensorIds.includes(record.sensor_id)
        )
        console.log('[dashboard] get records count: ', airRecords.length + soilRecords.length)
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
    totalCreatedToday,
    loadTotalCounts,
    addTreeNodeSelected,
    removeTreeNodeSelected,
    loadStations,
    setMapSelectedStation
  }
})
