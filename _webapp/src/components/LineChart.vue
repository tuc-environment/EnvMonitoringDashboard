<template>
  <div class="w-100 h-100" style="min-height: 360px">
    <div v-if="loading" class="w-100 h-100 text-center d-flex flex-column justify-content-center">
      <div class="mx-auto spinner-border"></div>
    </div>
    <div v-else-if="showDefaultText"
      class="small text-center text-secondary w-100 h-100 d-flex flex-column justify-content-center">
      {{ defaultText }}
    </div>
    <div v-else-if="noData"
      class="small text-center text-secondary w-100 h-100 d-flex flex-column justify-content-center">
      {{ noDataText }}
    </div>

    <div v-else class="d-flex flex-column w-100 h-100 text-dark">
      <div v-if="showSelections" class="d-flex flex-row mb-3">
        <select v-if="availablePositions.length > 0" class="form-select form-select-sm me-2" @input="onPositionChanged">
          <option :value="false">可选位置</option>
          <option v-for="(position, idx) in availablePositions" :key="idx" :value="position"
            :selected="position == selectedPosition">
            {{ getPositionName(position) }}
          </option>
        </select>
        <select v-if="avaialebleTags.length > 0" class="form-select form-select-sm me-2" @input="onTagChanged">
          <option :value="false">可选深度</option>
          <option v-for="tag in avaialebleTags" :key="tag" :value="tag" :selected="selectedTag == tag">
            {{ tag }}
          </option>
        </select>
        <select v-if="avaialebleGroups.length > 0" class="form-select form-select-sm me-2" @input="onGroupChanged">
          <option :value="false">可选组</option>
          <option v-for="(group, idx) in avaialebleGroups" :key="idx" :value="group" :selected="selectedGroup == group">
            {{ group }}
          </option>
        </select>
        <select v-if="availableOptionNames.length > 0" class="form-select form-select-sm" @input="onOptionNameChanged">
          <option :value="false">可选数据项</option>
          <option v-for="optionName in availableOptionNames" :key="optionName" :value="optionName"
            :selected="selectedOptionName == optionName">
            {{ optionName }}
          </option>
        </select>
      </div>
      <div class="flex-grow-1 w-100">
        <apexchart width="100%" height="100%" type="line" :options="chartOptions" :series="series" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  SensorPosition,
  getPositionName,
  getSensorDisplayText,
  type DataRecord,
  type Sensor,
  type Station
} from '@/http-client';
import { computed, ref, watch, type PropType } from 'vue';

const props = defineProps({
  records: Array as PropType<DataRecord[]>,
  sensors: Array as PropType<Sensor[]>,
  stations: Array as PropType<Station[]>,
  showDefaultText: Boolean,
  defaultText: String,
  noDataText: String,
  loading: Boolean,
  showSelections: {
    type: Boolean,
    default: false
  },
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
      palette: 'palette7'
    },
    chart: {
      animations: { enabled: false }
    },
    dataLabels: { enabled: false },
    stroke: {
      width: 1,
      lineCap: 'round'
      // curve: 'smooth'
    },
    markers: {
      size: 3
    },
    xaxis: {
      type: 'datetime',
      min: getMinDate(dates),
      max: getMaxDate(dates),
      labels: { show: true },
      tooltip: { enabled: false },
      crossharis: { show: false },
      axisBorder: { show: false },
      axisTicks: { show: true }
    },
    tooltip: {
      x: {
        format: 'dd MMM HH:mm'
      }
    },
    legend: {
      show: false
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
      if (selectedPosition.value && sensor.position != selectedPosition.value) {
        continue
      }
      if (selectedGroup.value && sensor.group != selectedGroup.value) {
        continue
      }
      if (selectedTag.value && sensor.tag != selectedTag.value) {
        continue
      }
      if (selectedOptionName.value && sensor.name != selectedOptionName.value) {
        continue
      }
      const relatedRecords = records.filter(
        (record) =>
          record.sensor_id &&
          sensor.id &&
          record.sensor_id == sensor.id &&
          record.time &&
          record.value
      )
      if (relatedRecords.length > 1) {
        var data: { x: Date; y: number }[] = []
        for (var record of relatedRecords) {
          data.push({
            x: new Date(record.time!),
            y: record.value!
          })
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
  return result
})

const getMinDate = (dates: Date[]): Date => new Date(Math.min(...dates.map(Number)))
const getMaxDate = (dates: Date[]): Date => new Date(Math.max(...dates.map(Number)))
const getDatesFromRecords = (records?: DataRecord[]): Date[] => {
  var result = new Set<Date>()
  for (const record of records ?? []) {
    if (record.time) {
      result.add(record.time)
    }
  }
  const arr = Array.from(result.values())
  console.log(`[line-chart] dates in records: ${arr}`)
  return arr
}

// selector logic
const availablePositions = computed<SensorPosition[]>(() => {
  var result: Set<SensorPosition> = new Set<SensorPosition>()
  const sensors = props.sensors ?? []
  if (sensors.length > 0) {
    for (var sensor of sensors) {
      if (sensor.position) {
        result.add(sensor.position)
      }
    }
  }
  console.log('[line-chart] availablePositions: ', result)
  return Array.from(result.values())
})
const selectedPosition = ref<SensorPosition | undefined>(undefined)
const onPositionChanged = (e: any) => {
  if (e.target.value != 'false') {
    selectedPosition.value = e.target.value
  } else {
    selectedPosition.value = undefined
  }
}

const avaialebleGroups = computed<string[]>(() => {
  var result: Set<string> = new Set<string>()
  const sensors = props.sensors ?? []
  if (sensors.length > 0) {
    for (var sensor of sensors) {
      if (sensor.group) {
        result.add(sensor.group)
      }
    }
  }
  console.log('[line-chart] avaialebleGroups: ', result)
  return Array.from(result.values())
})
const selectedGroup = ref<string | undefined>(undefined)
const onGroupChanged = (e: any) => {
  if (e.target.value != 'false') {
    selectedGroup.value = e.target.value
  } else {
    selectedGroup.value = undefined
  }
  console.log(selectedGroup.value)
}

const avaialebleTags = computed<string[]>(() => {
  var result: Set<string> = new Set<string>()
  const sensors = props.sensors ?? []
  if (sensors.length > 0) {
    for (var sensor of sensors) {
      if (sensor.tag) {
        result.add(sensor.tag)
      }
    }
  }
  console.log('[line-chart] avaialebleTags: ', result)
  return Array.from(result.values())
})
const selectedTag = ref<string | undefined>(undefined)
const onTagChanged = (e: any) => {
  if (e.target.value != 'false') {
    selectedTag.value = e.target.value
  } else {
    selectedTag.value = undefined
  }
}

const availableOptionNames = computed<string[]>(() => {
  var result: Set<string> = new Set<string>()
  const sensors = props.sensors ?? []
  if (sensors.length > 0) {
    for (var sensor of sensors) {
      if (sensor.name) {
        result.add(sensor.name)
      }
    }
  }
  console.log('[line-chart] availableOptions: ', result)
  return Array.from(result.values())
})
const selectedOptionName = ref<string | undefined>(undefined)
const onOptionNameChanged = (e: any) => {
  if (e.target.value != 'false') {
    selectedOptionName.value = e.target.value
  } else {
    selectedOptionName.value = undefined
  }
}

watch(
  () => availablePositions.value,
  (position) => {
    if (position.length > 0) {
      selectedPosition.value = position[0]
    }
  }
)

watch(
  () => avaialebleTags.value,
  (tags) => {
    if (tags.length > 0) {
      selectedTag.value = tags[0]
    }
  }
)

watch(
  () => avaialebleGroups.value,
  (groups) => {
    if (groups.length > 0) {
      selectedGroup.value = groups[0]
    }
  }
)

watch(
  () => availableOptionNames.value,
  (options) => {
    if (options.length > 0) {
      selectedOptionName.value = options[0]
    }
  }
)

</script>
@/http-client
