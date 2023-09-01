<template>
  <div
    style="
      position: absolute;
      left: 0;
      right: 0;
      top: 0;
      bottom: 0;
      background-color: rgb(17, 18, 23);
      z-index: -100;
    "
  ></div>

  <div style="background-color: rgb(17, 18, 23)">
    <div
      class="d-flex text-light p-2 align-items-center fixed-top"
      style="
        background-color: rgb(24, 27, 31);
        height: 48px;
        border-bottom: 1px solid rgba(204, 204, 220, 0.12);
      "
    >
      <div>
        <img class="logo rounded me-2" src="/logo.png" style="height: 28px" />
      </div>
      <div>环境监测系统 - 天津商业大学</div>
      <div class="flex-grow-1"></div>
      <div>
        <a href="/login" class="text-light">Sign in</a>
      </div>
    </div>

    <div
      class="p-3 text-light"
      style="
        position: absolute;
        left: 0;
        right: 0;
        top: 50px;
        bottom: 0;
        background-color: rgb(17, 18, 23);
        z-index: 1000;
        overflow-y: auto;
      "
    >
      <div class="row">
        <div class="col-lg-6 col-xl-6 h-100 my-1">
          <DashboardCardComponent title="选择传感器" class="info-card">
            <tag-group
              :tags="dashboardStore.treeSensorSelectedTags"
              @on-tag-selected="onSelectTag"
            />
            <tree-chart ref="treeChart" />
          </DashboardCardComponent>
        </div>
        <div class="col-lg-6 col-xl-6 h-100 my-1">
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
  </div>
</template>

<script setup lang="ts">
import DashboardCardComponent from '@/components/DashboardCardComponent.vue'
import LineChart from '@/components/LineChart.vue'
import TreeChart from '@/components/tree/Tree.vue'
import { type Sensor, type Station } from '@/httpclient'
import { useDashboardStore } from '@/stores/DashboardStore'
import TagGroup from '@/components/tags/TagsGroup.vue'
import { ref } from 'vue'
import type { TagData } from '@/components/tags/TagData'

const treeChart = ref<InstanceType<typeof TreeChart> | null>(null)
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
