<template>
  <div class="tree">
    <NodeComponent v-for="(node, idx) in nodes" :key="idx" :node="node" :depth="0"></NodeComponent>
  </div>
</template>

<script setup lang="ts">
import type { SensorPosition, Station } from '@/http-client'
import httpclient, { getPositionName } from '@/http-client'
import { onMounted, ref } from 'vue'
import NodeComponent from './NodeComponent.vue'
import { type Tree } from './Tree'

const nodes = ref<Tree[]>([])
const stations = ref<Station[]>([])

onMounted(() => {})

const setStations = async (updatedStations: Station[]) => {
  stations.value = updatedStations
  await loadSensors()
}

const loadSensors = async () => {
  const sensorsRes = await httpclient.getSensors({ visibleInDashboard: true })
  const sensors = sensorsRes?.payload ?? []
  var result: Tree[] = []
  for (const station of stations.value) {
    const sensorsInStation = sensors.filter((sensor) => sensor.station_id == station.id)
    const positions = [
      ...new Set(
        sensorsInStation
          .filter((sensor) => sensor.position != undefined && sensor.position != null)
          .map((sensor) => sensor.position)
      )
    ] as Array<SensorPosition>
    var positionNodes: Tree[] = []
    for (const position of positions) {
      const sensorsInPosition = sensorsInStation.filter((sensor) => sensor.position == position)
      const groups = [
        ...new Set(
          sensorsInPosition
            .filter((sensor) => sensor.group != undefined && sensor.group.trim() != '')
            .map((sensor) => sensor.group)
        )
      ]

      var groupNodes: Tree[] = []

      for (const group of groups) {
        const sensorsInGroup = sensorsInPosition.filter((sensor) => sensor.group == group)
        const tags = [
          ...new Set(
            sensorsInGroup
              .filter((sensor) => sensor.tag != undefined && sensor.tag.trim() != '')
              .map((sensor) => sensor.tag)
          )
        ]
        var tagNodes: Tree[] = []

        for (const tag of tags) {
          const sensorsWithTag = sensorsInGroup.filter((sensor) => sensor.tag == tag)
          const sensorNodes = sensorsWithTag.map((sensor): Tree => {
            return {
              label: sensor.name ?? '',
              children: [],
              sensor: sensor,
              station: station
            }
          })
          tagNodes.push({
            label: tag,
            children: sensorNodes
          } as Tree)
        }

        const noTagSensors = sensorsInGroup.filter(
          (sensor) => sensor.tag == undefined || sensor.tag.trim() == ''
        )
        for (const noTagSensor of noTagSensors) {
          tagNodes.push({
            label: noTagSensor.name,
            children: [],
            sensor: noTagSensor,
            station: station
          } as Tree)
        }

        const groupNode: Tree = {
          label: group ?? '',
          children: tagNodes,
          sensor: null,
          station: station
        }
        groupNodes.push(groupNode)
      }

      const noGroupSensors = sensorsInPosition.filter(
        (sensor) => sensor.group == undefined || sensor.group.trim() == ''
      )
      for (const noGroupSensor of noGroupSensors) {
        groupNodes.push({
          label: noGroupSensor.name,
          children: [],
          sensor: noGroupSensor,
          station: station
        } as Tree)
      }

      const positionNode: Tree = {
        label: getPositionName(position),
        children: groupNodes,
        sensor: null,
        station: station
      }
      positionNodes.push(positionNode)
    }
    const noPositionSensors = sensorsInStation.filter((sensor) => sensor.position == undefined)
    for (const noPositionSensor of noPositionSensors) {
      positionNodes.push({
        label: noPositionSensor.name,
        children: [],
        sensor: noPositionSensor,
        station: station
      } as Tree)
    }
    const stationNode: Tree = {
      label: station.name ?? '',
      children: positionNodes,
      sensor: null,
      station: station
    }
    result.push(stationNode)
  }
  nodes.value = result
}

defineExpose({
  setStations
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
@/http-client@/http-client
