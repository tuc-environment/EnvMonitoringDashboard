<template>
  <DashboardLayout>
    <div class="p-2 w-100 h-100">
      <div class="d-flex h-100 mx-1 card-container">
        <div class="p-2 card-col">
          <DashboardCardComponent title="选择传感器" :loading="loading">
            <TreeComponent ref="treeChart" class="align-self-start" />
          </DashboardCardComponent>
        </div>
        <div class="p-2 flex-grow-1">
          <DashboardCardComponent title="传感器数据">
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
import { type Station } from '@/httpclient'
import { useDashboardStore } from '@/stores/DashboardStore'
import { ref } from 'vue'

const treeChart = ref<InstanceType<typeof TreeComponent> | null>(null)
const dashboardStore = useDashboardStore()
const stations = ref<Station[]>([])
let loading = ref(true)

dashboardStore.$subscribe((_, state) => {
  if (stations.value.length != state.stations.length) {
    stations.value = state.stations
    treeChart.value?.setStations(stations.value)
  }
})

dashboardStore.loadStations().then(() => {
  loading.value = false
})
</script>

<style scoped>
.row > * {
  padding-left: 0;
  padding-right: 0;
}

.card-container {
  flex-direction: column;
}

@media (min-width: 992px) {
  .card-container {
    flex-direction: row;
  }
}

.card-col {
  height: 420px;
}

@media (min-width: 992px) {
  .card-col {
    height: 100%;
    width: 400px;
  }
}
</style>
