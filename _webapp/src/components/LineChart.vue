<template>
  <div class="wrapper row">
    <div v-if="$props.loading" class="d-flex align-items-center justify-content-center">
      <div class="spinner-border text-light" role="status"></div>
    </div>
    <div v-else-if="showDefaultText" class="h3 text-white text-center align-self-center col">
      <div class="h3 text-white">{{ $props.title }}</div>
      <div class="h5 text-white">{{ $props.defaultText }}</div>
    </div>
    <div v-else-if="noData" class="h3 text-white text-center align-self-center col">
      <div class="h3 text-white">{{ $props.title }}</div>
      <div class="h5 text-white">{{ $props.noDataText }}</div>
    </div>

    <div v-else class="col">
      <div class="h3 text-white">{{ $props.title }}</div>
      <apexchart width="100%" height="100%" type="line" :options="chartOptions" :series="series" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, type PropType } from 'vue'
import { type DataRecord, type Sensor, getSensorDisplayText } from '@/httpclient'

const props = defineProps({
  records: Array as PropType<DataRecord[]>,
  sensors: Array as PropType<Sensor[]>,
  title: String,
  showDefaultText: Boolean,
  defaultText: String,
  noDataText: String,
  loading: Boolean
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
        result.push({
          name: getSensorDisplayText(sensor),
          data: data
        })
      }
    }
  }
  console.log('[line-chart] series: ', result)
  return result.slice(0, 4)
})

const getMinDate = (dates: Date[]): Date => new Date(Math.min(...dates.map(Number)))
const getMaxDate = (dates: Date[]): Date => new Date(Math.max(...dates.map(Number)))
const getDatesFromRecords = (records?: DataRecord[]): Date[] => {
  return records ? (records.map((record) => record.time).filter((date) => date) as Date[]) : []
}
</script>

<style scoped>
.wrapper {
  position: relative;
  padding-top: 20px;
  padding-bottom: 8px;
  padding-left: 8px;
  padding-right: 8px;
  pointer-events: initial;
  border: none;
  width: 100%;
  margin: 3px;
  min-height: 300px;
}

.wrapper::before,
.wrapper::after {
  position: absolute;
  width: 80px;
  height: 60px;
  content: '';
}

.wrapper::before {
  left: 0;
  top: 0;
  border-left: 1px solid lightgray;
  border-top: 1px solid lightgray;
}

.wrapper::after {
  margin-right: 4px;
  right: 0;
  bottom: 0;
  border-right: 1px solid lightgray;
  border-bottom: 1px solid lightgray;
}

.chart {
  position: relative;
}
</style>
