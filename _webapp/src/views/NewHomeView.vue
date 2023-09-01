<template>
  <DashboardLayout>
    <div class="d-flex flex-column w-100">
      <div class="row align-items-stretch">
        <div class="col-md-6 col-lg-4 col-xl-2 mb-2">
          <DashboardCardComponent title="站点数量">
            <div
              v-if="dashboardStore.$state.totalStations == undefined"
              class="spinner-border text-light align-self-center"
            ></div>
            <div v-else class="text-danger" style="font-size: 4rem">
              {{ formatDisplayNumber(dashboardStore.$state.totalStations ?? 0) }}
            </div>
          </DashboardCardComponent>
        </div>
        <div class="col-md-6 col-lg-4 col-xl-2 mb-2">
          <DashboardCardComponent title="传感器数量">
            <div
              v-if="dashboardStore.$state.totalSensors == undefined"
              class="spinner-border text-light align-self-center"
            ></div>
            <div v-else class="text-danger" style="font-size: 4rem">
              {{ formatDisplayNumber(dashboardStore.$state.totalSensors ?? 0) }}
            </div>
          </DashboardCardComponent>
        </div>
        <div class="col-md-6 col-lg-4 col-xl-2 mb-2">
          <DashboardCardComponent title="数据总量">
            <div
              v-if="dashboardStore.$state.totalRecords == undefined"
              class="spinner-border text-light align-self-center"
            ></div>
            <div v-else class="text-danger" style="font-size: 4rem">
              {{ formatDisplayNumber(dashboardStore.$state.totalRecords ?? 0) }}
            </div>
          </DashboardCardComponent>
        </div>
        <div class="col-md-6 col-lg-4 col-xl-2 mb-2">
          <DashboardCardComponent title="今日更新数量">
            <div class="text-danger" style="font-size: 4vmax">1822</div>
          </DashboardCardComponent>
        </div>
        <div class="col-md-6 col-lg-4 col-xl-2 mb-2">
          <DashboardCardComponent title="今日更新数量">
            <div class="text-danger" style="font-size: 4vmax">1822</div>
          </DashboardCardComponent>
        </div>
        <div class="col-md-6 col-lg-4 col-xl-2 mb-2">
          <DashboardCardComponent title="今日更新数量">
            <div class="text-danger" style="font-size: 4vmax">1822</div>
          </DashboardCardComponent>
        </div>
      </div>

      <div class="row align-items-stretch flex-grow-1">
        <div class="col-lg-12 col-xl-8 my-2">
          <DashboardCardComponent title="站点位置" class="info-card">
            <map-container ref="mapContainer" @did-select-station="selectStationHandler" />
          </DashboardCardComponent>
        </div>
        <div class="col-lg-12 col-xl-4 my-2">
          <div class="h-100 d-flex flex-column justify-content-between">
            <div class="mb-2" style="flex: 1">
              <DashboardCardComponent
                :title="
                  dashboardStore.$state.selectedStation
                    ? `${dashboardStore.$state.selectedStation.name} - 空气参数`
                    : '空气参数'
                "
              >
                <line-chart
                  :sensors="dashboardStore.$state.airRelatedSensors"
                  :records="dashboardStore.$state.airRelatedRecords"
                  :show-default-text="!dashboardStore.$state.selectedStation"
                  :loading="dashboardStore.$state.loadingDataForStation"
                  default-text="点击地图选择站点查看数据"
                  no-data-text="暂无数据"
                />
              </DashboardCardComponent>
            </div>
            <div style="flex: 1">
              <DashboardCardComponent
                :title="
                  dashboardStore.$state.selectedStation
                    ? `${dashboardStore.$state.selectedStation.name} - 土壤参数`
                    : '土壤参数'
                "
              >
                <line-chart
                  :sensors="dashboardStore.$state.soilRelatedSensors"
                  :records="dashboardStore.$state.soilRelatedRecords"
                  :show-default-text="!dashboardStore.$state.selectedStation"
                  :loading="dashboardStore.$state.loadingDataForStation"
                  default-text="点击地图选择站点查看数据"
                  no-data-text="暂无数据"
                />
              </DashboardCardComponent>
            </div>
          </div>
        </div>
      </div>
    </div>
  </DashboardLayout>

  <ModalComponent
    title="预测站点信息"
    :visible="showPredictionModal"
    @close="showPredictionModal = false"
  />
</template>

<script setup lang="ts">
import DashboardLayout from '@/layouts/DashboardLayout.vue'
import DashboardCardComponent from '@/components/DashboardCardComponent.vue'
import ModalComponent from '@/components/ModalComponent.vue'
import LineChart from '@/components/LineChart.vue'
import TreeChart from '@/components/tree/Tree.vue'
import MapContainer from '@/components/map/MapContainer.vue'
import { type Station } from '@/httpclient'
import { useDashboardStore } from '@/stores/DashboardStore'
import { ref } from 'vue'
import { formatDisplayNumber } from '@/utils/utils'

const showPredictionModal = ref(false)
const treeChart = ref<InstanceType<typeof TreeChart> | null>(null)
const dashboardStore = useDashboardStore()
const stations = ref<Station[]>([])
const mapContainer = ref<InstanceType<typeof MapContainer> | null>(null)

dashboardStore.$subscribe((_, state) => {
  if (stations.value.length != state.stations.length) {
    stations.value = state.stations
    treeChart.value?.setStations(stations.value)
    mapContainer.value?.setStations(stations.value)
  }
})

const selectStationHandler = (station: Station | undefined) => {
  dashboardStore.setMapSelectedStation(station)
}

dashboardStore.loadStations()
dashboardStore.loadTotalCounts()
</script>
