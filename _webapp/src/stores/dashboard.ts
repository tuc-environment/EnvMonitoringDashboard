import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import httpclient, { type Station, type Sensor, type DataRecord } from '@/httpclient'
export const useDashboardStore = defineStore('dashboard', () => {
  const treeSensorSelected = ref<Sensor | undefined>(undefined)
  const treeStationSelected = ref<Station | undefined>(undefined)
  const treeSensorRecordsLoaded = ref<DataRecord[]>([])
  const setTreeNodeSelected = (sensor: Sensor | undefined, station: Station | undefined) => {
    treeSensorSelected.value = sensor
    treeStationSelected.value = station
    treeSensorRecordsLoaded.value = []
    loadRecordsForSelectedTreeSensor()
  }

  const loadRecordsForSelectedTreeSensor = async () => {
    const sensor = treeSensorSelected.value
    if (sensor) {
      try {
        const result = await httpclient.getRecords([sensor.id])
        treeSensorRecordsLoaded.value = result?.payload ?? []
      } catch (e) {
        console.log('[dashboard-store]', e)
      }
    }
  }

  return { treeSensorSelected, treeStationSelected, treeSensorRecordsLoaded, setTreeNodeSelected }
})
