<template>
  <DashboardLayout>
    <div class="p-2 w-100 h-100">
      <div class="row h-100">
        <div class="col-lg-5 col-xl-5 my-1" style="min-height: 380px">
          <DashboardCardComponent title="选择传感器" class="info-card">
            <tag-group
              :tags="dashboardStore.treeSensorSelectedTags"
              @on-tag-selected="onSelectTag"
            />
            <TreeComponent ref="treeChart" />
          </DashboardCardComponent>
        </div>
        <div class="col-lg-7 col-xl-7 my-1" style="min-height: 420px">
          <DashboardCardComponent title="传感器数据" class="info-card">
            <line-chart
              :stations="dashboardStore.$state.treeStationsSelected"
              :sensors="dashboardStore.$state.treeSensorsSelected"
              :records="dashboardStore.$state.treeSensorRecordsLoaded"
              :loading="dashboardStore.$state.treeRecordsLoading"
              :maxLines="10"
              title="数据对比"
              default-text="请选择数据项"
              no-data-text="暂无数据"
            />
          </DashboardCardComponent>
        </div>
      </div>
    </div>
  </DashboardLayout>
</template>

<script setup lang="ts">
import DashboardLayout from '@/layouts/DashboardLayout.vue'
import DashboardCardComponent from '@/components/DashboardCardComponent.vue'
import LineChart from '@/components/LineChart.vue'
import TreeComponent from '@/components/tree/TreeComponent.vue'
import { type Sensor, type Station } from '@/httpclient'
import { useDashboardStore } from '@/stores/DashboardStore'
import TagGroup from '@/components/tags/TagsGroup.vue'
import { ref } from 'vue'
import type { TagData } from '@/components/tags/TagData'

const treeChart = ref<InstanceType<typeof TreeComponent> | null>(null)
const dashboardStore = useDashboardStore()
const stations = ref<Station[]>([])

dashboardStore.$subscribe((_, state) => {
  if (stations.value.length != state.stations.length) {
    stations.value = state.stations
    treeChart.value?.setStations(stations.value)
  }
})

const onSelectTag = (tag: TagData) => {
  const sensor = tag.data as Sensor
  if (sensor) {
    dashboardStore.removeTreeNodeSelected(sensor)
  }
}

dashboardStore.loadStations()
// dashboardStore.loadSensors()
</script>
