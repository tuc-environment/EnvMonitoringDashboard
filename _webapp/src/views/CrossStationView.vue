<template>
  <DashboardLayout>
    <div class="p-2 w-100 h-100">
      <div class="row h-100">
        <div class="col-lg-5 col-xl-5 mb-2 card-col">
          <DashboardCardComponent title="选择传感器" :loading="loading">
            <TreeComponent ref="treeChart" class="align-self-start" />
          </DashboardCardComponent>
        </div>
        <div class="col-lg-7 col-xl-7 mb-2 card-col" style="min-height: 420px">
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
.card-col {
  min-height: 420px;
  height: 420px;
}

@media (min-width: 992px) {
  .card-col {
    height: 100%;
  }
}
</style>
