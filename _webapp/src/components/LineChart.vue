<template>
  <div class="wrapper row border border-light">
    <div v-if="noData" class="h3 text-white text-center align-self-center col">
      <div class="h3 text-white">{{ $props.title }}</div>
      <div class="h5 text-white">{{ $props.noDataText }}</div>
    </div>
    <div v-else-if="noStationSelected" class="h3 text-white text-center align-self-center col">
      <div class="h3 text-white">{{ $props.title }}</div>
      <div class="h5 text-white">{{ $props.defaultText }}</div>
    </div>
    <div v-else class="col">
      <div class="h3 text-white">{{ $props.title }}</div>
      <apexchart width="100%" height="100%" type="line" :options="chartOptions" :series="series" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, type PropType } from 'vue'
import {
  getPositionName,
  type DataRecord,
  type Sensor,
  type Station,
  getSensorDisplayText
} from '@/httpclient'

const props = defineProps({
  station: Object as PropType<Station>,
  records: Array as PropType<DataRecord[]>,
  sensors: Array as PropType<Sensor[]>,
  title: String,
  defaultText: String,
  noDataText: String
})

const noData = computed(() => props.station && (!props.records || props.records?.length == 0))
const noStationSelected = computed(() => !props.station)

const chartOptions = computed(() => {
  const dates: Date[] = getDatesFromRecords(props.records)
  return {
    theme: {
      palette: 'palette3'
    },
    chart: {
      foreColor: '#d3d3d3'
    },
    xaxis: {
      type: 'datetime',
      min: getMinDate(dates),
      max: getMaxDate(dates),
      title: {
        text: '时间',
        offsetX: '-50%',
        offsetY: -20
      }
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
        name: getSensorDisplayText(sensor),
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
