<template>
  <div class="tree">
    <node v-for="node in nodes" :node="node" :depth="0"></node>
  </div>
</template>

<script setup lang="ts">
import type { Station } from '@/httpclient'
import Node from './Node.vue'
import { type Tree } from './Tree'
import { reactive, ref, onMounted } from 'vue'
import httpclient, { getPositionName } from '@/httpclient'

const nodes = ref<Tree[]>([])
const stations = ref<Station[]>([])

onMounted(() => {})

const setStations = async (updatedStations: Station[]) => {
  stations.value = updatedStations
  await loadSensors()
}

const loadSensors = async () => {
  const sensorsRes = await httpclient.getSensors()
  const sensors = sensorsRes?.payload ?? []
  var result: Tree[] = []
  for (const station of stations.value) {
    var stationNode: Tree = {
      label: station.name ?? '',
      children: []
    }

    const sensorsInStation = sensors.filter((sensor) => sensor.station_id == station.id)
    const positions = [...new Set(sensorsInStation.map((sensor) => sensor.position))]
    var positionNodes: Tree[] = []
    for (const position of positions) {
      positionNodes.push({
        label: getPositionName(position ?? ''),
        children: []
      } as Tree)
    }
    stationNode.children = positionNodes
    result.push(stationNode)
  }
  nodes.value = result
}

defineExpose({
  setStations
})
</script>

<style scoped>
.tree-list ul {
  padding-left: 16px;
  margin: 6px;
}

.tree {
  pointer-events: initial;
  min-height: 300px;
  padding: 8px;
}
</style>
