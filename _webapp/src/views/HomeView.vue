<template>
  <DashboardLayout>
    <div class="mt-5 mb-2 text-center">
      <h1>光伏电站环境影响监测和评估系统</h1>
      <div class="mx-auto" style="max-width: 480px">
        本环境监测系统依托光伏发电系统生态环境监测分析硬件平台，通过高质量、高频率的环境参数监测，实现对光伏电站建设对周边生态环境的全面评估。
      </div>
    </div>

    <div class="my-2">
      <div class="row mx-auto" style="max-width: 640px">
        <div class="col-md-3 col-6 p-2 text-center">
          <div style="font-size: 3rem" class="text-success">
            {{ formatDisplayNumber(dashboardStore.$state.totalStations ?? 0) }}
          </div>
          <div>站点数量</div>
        </div>
        <div class="col-md-3 col-6 p-2 text-center">
          <div style="font-size: 3rem" class="text-success">
            {{ formatDisplayNumber(dashboardStore.$state.totalSensors ?? 0) }}
          </div>
          <div>传感器数量</div>
        </div>
        <div class="col-md-3 col-6 p-2 text-center">
          <div style="font-size: 3rem" class="text-success">
            {{ formatDisplayNumber(dashboardStore.$state.totalRecords ?? 0) }}
          </div>
          <div>数据总量</div>
        </div>
        <div class="col-md-3 col-6 p-2 text-center">
          <div style="font-size: 3rem" class="text-success">
            {{ formatDisplayNumber(dashboardStore.$state.totalCreatedToday ?? 0) }}
          </div>
          <div>今日更新数据量</div>
        </div>
      </div>
    </div>

    <hr />

    <div class="my-2">
      <div class="text-center mx-auto mb-2">
        <h2>站点位置</h2>
        <div class="mx-auto" style="max-width: 480px">
          点击地图站点查看站点数据，或者新建预测站点。
        </div>
      </div>

      <div class="row mx-auto" style="max-width: 720px; aspect-ratio: 4/3">
        <div class="col-md-12 p-2">
          <map-container
            ref="mapContainer"
            @did-select-station="selectStationHandler"
            @on-confirm-prediction-marker="onPredictionMarkerConfirmed"
          />
        </div>
      </div>

      <div class="row mx-auto">
        <div class="col-md-6 p-2">
          <div class="text-center my-2">
            {{
              dashboardStore.$state.selectedStation
                ? `${dashboardStore.$state.selectedStation.name} - 空气参数`
                : '空气参数'
            }}
          </div>
          <line-chart
            :sensors="dashboardStore.$state.airRelatedSensors"
            :records="dashboardStore.$state.airRelatedRecords"
            :show-default-text="!dashboardStore.$state.selectedStation"
            :loading="dashboardStore.$state.loadingDataForStation"
            :show-selections="true"
            default-text="点击地图选择站点查看数据"
            no-data-text="暂无数据"
          />
        </div>
        <div class="col-md-6 p-2">
          <div class="text-center my-2">
            {{
              dashboardStore.$state.selectedStation
                ? `${dashboardStore.$state.selectedStation.name} - 土壤参数`
                : '土壤参数'
            }}
          </div>
          <line-chart
            :sensors="dashboardStore.$state.soilRelatedSensors"
            :records="dashboardStore.$state.soilRelatedRecords"
            :show-default-text="!dashboardStore.$state.selectedStation"
            :loading="dashboardStore.$state.loadingDataForStation"
            :show-selections="true"
            default-text="点击地图选择站点查看数据"
            no-data-text="暂无数据"
          />
        </div>
      </div>
    </div>

    <hr />

    <div class="my-2 text-center">
      <div class="mx-auto text-secondary small" style="max-width: 320px">
        Environmental Impact Monitoring and Assessment System for Photovoltaic Power Plants
      </div>
      <div class="mx-auto text-secondary small" style="max-width: 320px">
        All Rights Reserved © 2024
      </div>
    </div>

    <div style="height: 64px"></div>
  </DashboardLayout>

  <PredictionModalComponent
    ref="predictionModal"
    title="预测站点信息"
    :visible="showPredictionModal"
    @close="onModalClosed"
  />
</template>

<script setup lang="ts">
import LineChart from '@/components/LineChart.vue'
import MapContainer from '@/components/map/MapContainer.vue'
import PredictionModalComponent from '@/components/modal/PredictionModalComponent.vue'
import TreeChart from '@/components/tree/TreeComponent.vue'
import { type Station } from '@/http-client'
import DashboardLayout from '@/layouts/DashboardLayout.vue'
import { useDashboardStore } from '@/stores/DashboardStore'
import { formatDisplayNumber } from '@/utils/utils'
import { ref } from 'vue'

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
