<template>
  <div>
    <div :style="indent" :class="labelClasses" @click="clickNode">
      {{ node?.label }}
      <div v-if="hasChildren" class="bi" :class="iconClasses"></div>
    </div>
    <div v-if="hasChildren && showChildren">
      <Node v-for="child in $props.node?.children" :node="child" :depth="$props.depth + 1"></Node>
    </div>
  </div>
</template>

<script setup lang="ts">
import { type Tree } from './Tree'
import { ref, type PropType, computed } from 'vue'
import Node from './Node.vue'
import { useDashboardStore } from '@/stores/dashboard'
import type { Sensor } from '@/httpclient'

const dashboardStore = useDashboardStore()

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
      dashboardStore.addTreeNodeSelected(sensor, station)
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
