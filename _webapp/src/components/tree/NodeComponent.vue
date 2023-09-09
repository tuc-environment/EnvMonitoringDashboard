<template>
  <div>
    <div :style="indent" :class="labelClasses" @click="clickNode">
      <i v-if="selected || hasChildrenSelected" class="bi bi-check2-square"></i>
      {{ node?.label }}
      <div v-if="hasChildren" class="bi" :class="iconClasses"></div>
    </div>
    <div v-if="hasChildren && showChildren">
      <NodeComponent
        v-for="(child, idx) in $props.node?.children"
        :key="idx"
        :node="child"
        :depth="$props.depth + 1"
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

const iconClasses = computed(() => {
  return {
    'bi-plus-circle': !showChildren.value,
    'bi-dash-circle': showChildren.value
  }
})

const labelClasses = computed(() => {
  return {
    'has-children': hasChildren.value,
    'label-wrapper': true,
    'leaf-node': !hasChildren.value
  }
})
const indent = computed(() => {
  return {
    transform: `translate(${props.depth * 50}px)`
  }
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

<style scoped>
.label-wrapper {
  display: flex;
  flex-direction: row;
  gap: 8px;
  padding-bottom: 10px;
  margin-bottom: 10px;
  border-bottom: 1px solid #ccc;
  font-size: 14px;
  overflow: hidden;
}

.has-children {
  cursor: pointer;
  color: lightblue;
}

.leaf-node {
  color: white;
}
</style>
