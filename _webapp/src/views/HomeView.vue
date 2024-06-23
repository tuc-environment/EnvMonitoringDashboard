<template>
  <div class="home-container p-3">
    <div class="header">光伏发电生态设计评估系统</div>

    <!-- nav buttons -->

    <div class="mx-auto d-flex" style="width: 90%; margin-top: -10px">
      <ButtonComponent class="me-5" title="数据看板" />
      <ButtonComponent class="me-auto" title="数据对比" />
      <ButtonComponent class="ms-auto" title="联系我们" />
      <ButtonComponent class="ms-5" title="进入管理后台" />
    </div>

    <!-- status -->
    <div class="mx-auto my-3 d-flex justify-content-between" style="width: 60%">
      <ButtonComponent2
        title="站点数量"
        :value="formatDisplayNumber(dashboardStore.$state.totalStations ?? 0)"
      />
      <ButtonComponent2
        title="传感器数量"
        :value="formatDisplayNumber(dashboardStore.$state.totalSensors ?? 0)"
      />
      <ButtonComponent2
        title="数据总量"
        :value="formatDisplayNumber(dashboardStore.$state.totalRecords ?? 0)"
      />
      <ButtonComponent2
        title="今日更新数据量"
        :value="formatDisplayNumber(dashboardStore.$state.totalCreatedToday ?? 0)"
      />
    </div>

    <!-- content -->
    <div class="d-flex">
      <!-- left -->
      <div>
        <div class="left-container">
          <map-container
            ref="mapContainer"
            @did-select-station="selectStationHandler"
            @on-confirm-prediction-marker="onPredictionMarkerConfirmed"
          />
        </div>
      </div>

      <!-- right -->
      <div>
        <div class="right-top-container">
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
        <div class="right-bottom-container">
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
  </div>

  <PredictionModalComponent
    ref="predictionModal"
    title="预测站点信息"
    :visible="showPredictionModal"
    @close="onModalClosed"
  />
</template>

<script setup lang="ts">
import ButtonComponent from '@/components/home/ButtonComponent.vue'
import ButtonComponent2 from '@/components/home/ButtonComponent2.vue'

import LineChart from '@/components/LineChart.vue'
import MapContainer from '@/components/map/MapContainer.vue'
import PredictionModalComponent from '@/components/modal/PredictionModalComponent.vue'
import TreeChart from '@/components/tree/TreeComponent.vue'
import { type Station } from '@/http-client'
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

<style scoped lang="scss">
.home-container {
  background-image: url('./home/bg.png');
  background-repeat: no-repeat;
  background-position: center;
  background-size: 100% 100%;
  width: 100%;
  height: 100%;
}

.header {
  color: green;
  font-size: 24px;
  background-image: url('./home/top.png');
  background-repeat: no-repeat;
  background-position: center;
  background-size: 100% 100%;
  text-align: center;
}

.left-container {
  background-image: url('./home/left-container.png');
  background-repeat: no-repeat;
  background-position: center;
  background-size: 100% 100%;
  width: 480px;
  height: 320px;
}

.right-top-container {
  background-image: url('./home/right-top-container.png');
  background-repeat: no-repeat;
  background-position: center;
  background-size: 100% 100%;
  width: 480px;
  height: 160px;
}

.right-bottom-container {
  background-image: url('./home/right-bottom-container.png');
  background-repeat: no-repeat;
  background-position: center;
  background-size: 100% 100%;
  width: 480px;
  height: 160px;
}
</style>
