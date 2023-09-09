<template>
  <div>
    <div
      class="d-flex my-2 small"
      :style="{ transform: `translate(${depth * 20}px)` }"
      style="cursor: pointer"
      @click="clickNode"
    >
      <!-- <i v-if="selected || hasChildrenSelected" class="bi bi-check2-square"></i> -->

      <div v-if="hasChildren">
        <i v-if="!showChildren" class="bi bi-caret-right-fill me-2"></i>
        <i v-else class="bi bi-caret-down-fill me-2"></i>

        {{ node?.label }}
      </div>
      <div v-else class="form-check">
        <input class="form-check-input" type="checkbox" />
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
    console.log('select node')
    showChildren.value = !showChildren.value
  } else {
    console.log('select leaf')
    const sensor = props.node?.sensor
    const station = props.node?.station
    if (sensor && station) {
      if (selected.value) {
        dashboardStore.removeTreeNodeSelected(sensor)
        console.log('[tree-node] deselect leaf: ', JSON.stringify(sensor))
      } else {
        dashboardStore.addTreeNodeSelected(sensor, station)
        console.log('[tree-node] select leaf: ', JSON.stringify(sensor))
      }
    }
  }
}
</script>
