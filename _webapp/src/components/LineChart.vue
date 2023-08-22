<template>
  <div class="wrapper row border border-light">
    <div v-if="noData" class="h3 text-white text-center align-self-center col">No Data yet</div>
    <div v-else-if="noStationSelected" class="h3 text-white text-center align-self-center col">
      Please select a station to view data
    </div>
    <div v-else class="col">
      <div class="h3 text-white">{{ $props.station?.name }}</div>
      <apexchart width="100%" height="100%" type="line" :options="chartOptions" :series="series" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, type PropType } from 'vue'
import { type DataRecord, type Sensor, type Station } from '@/httpclient'

const props = defineProps({
  station: Object as PropType<Station>,
  records: Array as PropType<DataRecord[]>,
  sensors: Array as PropType<Sensor[]>
})

const noData = computed(() => props.station && (!props.records || props.records?.length == 0))
const noStationSelected = computed(() => !props.station)

const chartOptions = computed(() => {
  const dates: Date[] = getDatesFromRecords(props.records)
  return {
    xaxis: {
      type: 'datetime',
      min: getMinDate(dates),
      max: getMaxDate(dates)
    }
  }
})

const series = computed(() => {
  var result = []
  const sensors = props.sensors ?? []
  const records = props.records ?? []
  if (sensors.length > 0 && records.length > 0) {
    for (var sensor of sensors) {
      const relatedRecords = records.filter(
        (record) => record.sensor_id && sensor.id && record.sensor_id == sensor.id && record.time
      )
      var data: (Date | number)[][] = []
      for (var record of relatedRecords) {
        if (record.time && record.value) {
          data.push([record.time, record.value])
        }
      }
      result.push({
        name: `${sensor.position ?? ''}-${sensor.group ?? ''}-${sensor.tag ?? ''}-${
          sensor.name ?? ''
        }`,
        data: data
      })
      console.log('[line-chart] series: ', result)
    }
  }
  return result.slice(0, 2)
})

const getMinDate = (dates: Date[]): Date => new Date(Math.min(...dates.map(Number)))
const getMaxDate = (dates: Date[]): Date => new Date(Math.max(...dates.map(Number)))
const getDatesFromRecords = (records?: DataRecord[]): Date[] => {
  return records ? (records.map((record) => record.time).filter((date) => date) as Date[]) : []
}
</script>

<style scoped>
.wrapper {
  pointer-events: initial;
  min-height: 300px;
  padding: 8px;
}

.chart {
  position: relative;
}
</style>
