<template>
  <div>
    <div class="tree">
      <UnlinkedSensorNode
        v-for="(node, idx) in nodes"
        :key="idx"
        :node="node"
        :depth="0"
        @select="onSelectNode"
        @deselect="onDeselectNode"
      ></UnlinkedSensorNode>
    </div>
    <div class="container p-2">
      <div class="row mb-2 justify-content-around">
        <button
          type="button"
          class="btn btn-secondary btn-block col-5"
          :disabled="loading"
          @click="$emit('onCancel')"
        >
          取消
        </button>
        <button
          type="button"
          class="btn btn-success btn-block col-5"
          :disabled="loading"
          @click="onConfirm"
        >
          <div v-if="loading">
            <span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>
            更新中...
          </div>
          <div v-if="!loading">确定</div>
        </button>
      </div>
      <div style="text-align: center" class="mb-2"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { type CommonNode } from '@/components/modal/link-sensors-model/CommonNode'
import UnlinkedSensorNode from '@/components/modal/link-sensors-model/UnlinkedSensorNode.vue'
import httpclient, { type Sensor, type Station } from '@/http-client'
import { nextTick, ref } from 'vue'

const station = ref<Station | undefined>(undefined)
const nodes = ref<CommonNode[]>([])
const sensors = ref<Sensor[]>([])
const selectedSensorIds = ref<Set<number>>(new Set<number>())
const originalSelectedSensorIds = ref<Set<number>>(new Set<number>())

const loading = ref<boolean>(false)

const emit = defineEmits<{
  (e: 'didUpdateSensors'): void
  (e: 'onCancel'): void
}>()

const onConfirm = async () => {
  try {
    loading.value = true
    const sensorLinked = sensors.value
      .filter(
        (sensor) =>
          !originalSelectedSensorIds.value.has(sensor.id) && selectedSensorIds.value.has(sensor.id)
      )
      .map((sensor): Sensor => {
        return { ...sensor, station_id: station.value?.id ?? 0 }
      })
    const sensorUnlinked = sensors.value
      .filter(
        (sensor) =>
          originalSelectedSensorIds.value.has(sensor.id) && !selectedSensorIds.value.has(sensor.id)
      )
      .map((sensor): Sensor => {
        return { ...sensor, station_id: 0 }
      })
    console.log(originalSelectedSensorIds.value)
    console.log(selectedSensorIds.value)
    await httpclient.batchUpsertSensor(sensorLinked.concat(sensorUnlinked))
    loading.value = false
    emit('didUpdateSensors')
  } catch (_) {
    loading.value = false
  }
}

const setStation = async (updatedStation: Station | undefined) => {
  station.value = updatedStation
  await loadSensors()
}

const loadSensors = async () => {
  nodes.value = []
  originalSelectedSensorIds.value = new Set<number>()
  selectedSensorIds.value = new Set<number>()

  const assignedSensorsRes = await httpclient.getSensors({ stationID: station.value?.id })
  const unassignedSensorsRes = await httpclient.getUnassignedSensors()
  const assignedSensors = assignedSensorsRes?.payload ?? []
  const unassignedSensors = unassignedSensorsRes?.payload ?? []
  const assignedSensorIds = assignedSensors.map((s) => s.id)
  selectedSensorIds.value = new Set<number>(assignedSensorIds)
  originalSelectedSensorIds.value = new Set<number>(assignedSensorIds)
  sensors.value = assignedSensors.concat(unassignedSensors)
  await nextTick()
  var nodeData: CommonNode[] = []
  nodeData = mapSensorsToNodes(sensors.value)
  nodes.value = nodeData
}

const mapSensorsToNodes = (sensors: Sensor[]): CommonNode[] => {
  var data: CommonNode[] = []
  var selectedIMEIs: Set<string> = new Set<string>()
  for (const assignedSensor of sensors) {
    const imei = assignedSensor.imei
    if (imei) {
      selectedIMEIs.add(imei)
    }
  }
  for (const imei of selectedIMEIs.values()) {
    const sensorsWithSameIMEI: CommonNode[] = sensors
      .filter((s) => s.imei == imei)
      .map((s) => {
        return {
          label: s.sensor_report_code + ': ' + s.name,
          id: s.id,
          children: new Array<CommonNode>(),
          selected: selectedSensorIds.value.has(s.id)
        }
      })
    var selected = true
    sensorsWithSameIMEI.forEach((s) => {
      selected = selected && (s.selected ?? false)
    })
    data.push({
      label: 'IMEI: ' + imei,
      children: sensorsWithSameIMEI,
      selected: selected
    })
  }
  return data
}

const onSelectNode = (ids: number[]) => {
  console.log(`[link-sensor-tree] select ids: ${ids}`)
  const updated = selectedSensorIds.value
  ids.forEach((id) => updated.add(id))
  selectedSensorIds.value = updated
  updateNodes()
}

const onDeselectNode = (ids: number[]) => {
  console.log(`[link-sensor-tree] deselect ids: ${ids}`)
  const updated = selectedSensorIds.value
  ids.forEach((id) => updated.delete(id))
  selectedSensorIds.value = updated
  updateNodes()
}

const updateNodes = async () => {
  nodes.value = mapSensorsToNodes(sensors.value)
}

defineExpose({
  setStation
})
</script>

<style scoped>
.tree {
  pointer-events: initial;
  overflow-y: scroll;
  overflow-x: hidden;
  padding: 8px;
  width: 100%;
  height: 100%;
}

::-webkit-scrollbar {
  -webkit-appearance: none;
  width: 7px;
}

::-webkit-scrollbar-thumb {
  display: none;
  border-radius: 4px;
  background-color: rgba(0, 0, 0, 0.5);
  box-shadow: 0 0 1px rgba(255, 255, 255, 0.5);
}
</style>
