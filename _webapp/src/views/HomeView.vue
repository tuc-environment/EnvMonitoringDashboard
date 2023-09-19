<template>
  <DashboardLayout>
    <div class="p-2 d-flex flex-column w-100 h-100">
      <div class="row align-items-stretch mx-1">
        <div class="col-md-6 col-lg-3 col-xl-3 p-2">
          <DashboardCardComponent title="站点数量" :loading="!dashboardStore.$state.totalStations">
            <div style="font-size: 4rem">
              {{ formatDisplayNumber(dashboardStore.$state.totalStations ?? 0) }}
            </div>
          </DashboardCardComponent>
        </div>
        <div class="col-md-6 col-lg-3 col-xl-3 p-2">
          <DashboardCardComponent title="传感器数量" :loading="!dashboardStore.$state.totalSensors">
            <div style="font-size: 4rem">
              {{ formatDisplayNumber(dashboardStore.$state.totalSensors ?? 0) }}
            </div>
          </DashboardCardComponent>
        </div>
        <div class="col-md-6 col-lg-3 col-xl-3 p-2">
          <DashboardCardComponent title="数据总量" :loading="!dashboardStore.$state.totalRecords">
            <div style="font-size: 4rem">
              {{ formatDisplayNumber(dashboardStore.$state.totalRecords ?? 0) }}
            </div>
          </DashboardCardComponent>
        </div>
        <div class="col-md-6 col-lg-3 col-xl-3 p-2">
          <DashboardCardComponent title="今日更新数据量">
            <div style="font-size: 4rem">
              {{ formatDisplayNumber(dashboardStore.$state.totalCreatedToday ?? 0) }}
            </div>
          </DashboardCardComponent>
        </div>
      </div>

      <div class="row align-items-stretch flex-grow-1 mx-1">
        <div class="col-lg-12 col-xl-8 p-2">
          <DashboardCardComponent title="站点位置" class="info-card">
            <map-container
              ref="mapContainer"
              @did-select-station="selectStationHandler"
              @on-confirm-prediction-marker="onPredictionMarkerConfirmed"
            />
          </DashboardCardComponent>
        </div>
        <div class="col-lg-12 col-xl-4 p-2">
          <div class="h-100 d-flex flex-column justify-content-between">
            <div class="mb-3" style="flex: 1">
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
                  :show-selections="true"
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
                  :show-selections="true"
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

  <PredictionModalComponent
    ref="predictionModal"
    title="预测站点信息"
    :visible="showPredictionModal"
    @close="onModalClosed"
  />
</template>

<script setup lang="ts">
import DashboardLayout from '@/layouts/DashboardLayout.vue'
import DashboardCardComponent from '@/components/DashboardCardComponent.vue'
import PredictionModalComponent from '@/components/modal/PredictionModalComponent.vue'
import LineChart from '@/components/LineChart.vue'
import TreeChart from '@/components/tree/TreeComponent.vue'
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
const predictionModal = ref<InstanceType<typeof PredictionModalComponent> | null>(null)

// predict marker
const lngVal = ref<number | undefined>(undefined)
const latVal = ref<number | undefined>(undefined)

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

const onPredictionMarkerConfirmed = (lng: number, lat: number) => {
  lngVal.value = lng
  latVal.value = lat
  predictionModal.value?.setCoordinate(lat, lng)
  showPredictionModal.value = true
}

const onModalClosed = () => {
  showPredictionModal.value = false
  mapContainer.value?.closePredictionMarker()
}

dashboardStore.loadStations()
dashboardStore.loadTotalCounts()
</script>

<style scoped>
.row > * {
  padding-left: 0;
  padding-right: 0;
}
</style>
