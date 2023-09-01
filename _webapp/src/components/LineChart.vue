<template>
  <div class="d-flex flex-column justify-content-center" style="min-height: 240px">
    <div v-if="loading" class="small text-center text-secondary">
      <div class="spinner-border text-secondary"></div>
    </div>
    <div v-else-if="showDefaultText" class="small text-center text-secondary">
      {{ defaultText }}
    </div>
    <div v-else-if="noData" class="small text-center text-secondary">
      {{ noDataText }}
    </div>

    <div v-else class="w-100 h-100">
      <apexchart height="100%" type="line" :options="chartOptions" :series="series" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, type PropType } from 'vue'
import { type DataRecord, type Sensor, getSensorDisplayText, type Station } from '@/httpclient'

const props = defineProps({
  records: Array as PropType<DataRecord[]>,
  sensors: Array as PropType<Sensor[]>,
  stations: Array as PropType<Station[]>,
  showDefaultText: Boolean,
  defaultText: String,
  noDataText: String,
  loading: Boolean,
  maxLines: {
    type: Number,
    default: 4
  }
})

const noData = computed(() => !props.records || props.records?.length == 0)
const showDefaultText = computed(() => props.showDefaultText)

const chartOptions = computed(() => {
  const dates: Date[] = getDatesFromRecords(props.records)
  return {
    theme: {
      palette: 'palette2'
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
  const stations = props.stations ?? []
  const sensors = props.sensors ?? []
  const records = props.records ?? []
  if (sensors.length > 0 && records.length > 0) {
    for (var sensor of sensors) {
      const relatedRecords = records.filter(
        (record) =>
          record.sensor_id &&
          sensor.id &&
          record.sensor_id == sensor.id &&
          record.time &&
          record.value
      )
      if (relatedRecords.length > 1) {
        var data: (Date | number)[][] = []
        for (var record of relatedRecords) {
          data.push([record.time!, record.value!])
        }
        const station = stations.find((station) => station.id == sensor.station_id)
        const displayText = station
          ? getSensorDisplayText(sensor, station.name)
          : getSensorDisplayText(sensor)
        result.push({
          name: displayText,
          data: data
        })
      }
    }
  }
  console.log('[line-chart] series: ', result)
  if (result.length > props.maxLines) {
    return result.slice(0, props.maxLines)
  } else {
    return result
  }
})

const getMinDate = (dates: Date[]): Date => new Date(Math.min(...dates.map(Number)))
const getMaxDate = (dates: Date[]): Date => new Date(Math.max(...dates.map(Number)))
const getDatesFromRecords = (records?: DataRecord[]): Date[] => {
  return records ? (records.map((record) => record.time).filter((date) => date) as Date[]) : []
}
</script>
