<template>
  <div style="text-align: left">
    <div class="d-flex my-2 small node p-2 rounded" @click="clickNode">
      <div :style="{ width: `${depth * 20}px` }"></div>
      <div v-if="hasChildren" class="flex-grow-1">
        <div>
          <i v-if="!showChildren" class="bi bi-caret-right-fill me-2"></i>
          <i v-else class="bi bi-caret-down-fill me-2"></i>

          {{ node?.label }}

          <i v-if="hasChildrenSelected" class="ms-2 bi bi-check-circle"></i>
        </div>
      </div>

      <div v-else class="d-flex align-items-center">
        <input
          class="form-check-input me-2"
          type="checkbox"
          :checked="isSensorSelected(node?.sensor?.id)"
        />
        <label class="form-check-label"> {{ node?.label }} </label>
      </div>
    </div>
    <div v-if="hasChildren && showChildren">
      <NodeComponent
        v-for="(child, idx) in node?.children"
        :key="idx"
        :node="child"
        :depth="depth + 1"
      ></NodeComponent>
    </div>
  </div>
</template>

<script setup lang="ts">
import { type Tree, traverseChildren } from './Tree'
import { ref, type PropType, computed } from 'vue'
import NodeComponent from './NodeComponent.vue'
import { useDashboardStore } from '@/stores/DashboardStore'

const dashboardStore = useDashboardStore()

const selected = computed(
  () =>
    props.node?.sensor &&
    dashboardStore.$state.treeSensorsSelected
      .map((sensor) => sensor.id)
      .includes(props.node.sensor.id)
)

function isSensorSelected(sensorId: number | undefined) {
  sensorId = sensorId || -1
  return dashboardStore.$state.treeSensorsSelected.map((sensor) => sensor.id).includes(sensorId)
}

const hasChildrenSelected = computed(() => {
  const childrenSensorIds = traverseChildren(props.node).map((sensor) => sensor.id)
  const selectedSensorIds = dashboardStore.$state.treeSensorsSelected.map((sensor) => sensor.id)
  for (const childSensorId of childrenSensorIds) {
    if (selectedSensorIds.includes(childSensorId)) {
      return true
    }
  }
  return false
})

const props = defineProps({
  node: Object as PropType<Tree>,
  depth: {
    type: Number,
    required: true
  }
})

const showChildren = ref(false)
const hasChildren = computed((): boolean => {
  return props.node?.children != undefined && props.node.children.length > 0
})

const clickNode = () => {
  if (hasChildren.value) {
    showChildren.value = !showChildren.value
  } else {
    const sensor = props.node?.sensor
    const station = props.node?.station
    if (sensor && station) {
      if (selected.value) {
        dashboardStore.removeTreeNodeSelected(sensor)
      } else {
        dashboardStore.addTreeNodeSelected(sensor, station)
      }
    }
  }
}
</script>

<style>
.node {
  cursor: pointer;
}

.node:hover {
  background-color: #c3c1c126;
}
</style>
