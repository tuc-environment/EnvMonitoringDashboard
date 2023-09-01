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
      class="p-3 text-light d-flex flex-column"
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
      <div class="row align-items-stretch">
        <div class="col-md-6 col-lg-4 col-xl-2 mb-2">
          <DashboardCardComponent title="站点数量">
            <div
              v-if="dashboardStore.$state.loadingStations"
              class="spinner-border text-light align-self-center"
            ></div>
            <div v-else class="text-danger" style="font-size: 4rem">
              {{ formatDisplayNumber(dashboardStore.$state.stations.length ?? 0) }}
            </div>
          </DashboardCardComponent>
        </div>
        <div class="col-md-6 col-lg-4 col-xl-2 mb-2">
          <DashboardCardComponent title="传感器数量">
            <div
              v-if="dashboardStore.$state.loadingSensors"
              class="spinner-border text-light align-self-center"
            ></div>
            <div v-else class="text-danger" style="font-size: 4rem">
              {{ formatDisplayNumber(dashboardStore.$state.sensors.length ?? 0) }}
            </div>
          </DashboardCardComponent>
        </div>
        <div class="col-md-6 col-lg-4 col-xl-2 mb-2">
          <DashboardCardComponent title="今日新增数据">
            <div
              v-if="dashboardStore.$state.loadingSensors"
              class="spinner-border text-light align-self-center"
            ></div>
            <div v-else class="text-danger" style="font-size: 4rem">
              {{ formatDisplayNumber(1100000000) }}
            </div>
          </DashboardCardComponent>
        </div>
        <div class="col-md-6 col-lg-4 col-xl-2 mb-2">
          <DashboardCardComponent title="今日更新数量">
            <div class="text-danger" style="font-size: 4rem">1822</div>
          </DashboardCardComponent>
        </div>
        <div class="col-md-6 col-lg-4 col-xl-2 mb-2">
          <DashboardCardComponent title="今日更新数量">
            <div class="text-danger" style="font-size: 4rem">1822</div>
          </DashboardCardComponent>
        </div>
        <div class="col-md-6 col-lg-4 col-xl-2 mb-2">
          <DashboardCardComponent title="今日更新数量">
            <div class="text-danger" style="font-size: 4rem">1822</div>
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
  </div>
</template>

<script setup lang="ts">
import DashboardCardComponent from '@/components/DashboardCardComponent.vue'
import LineChart from '@/components/LineChart.vue'
import TreeChart from '@/components/tree/Tree.vue'
import MapContainer from '@/components/map/MapContainer.vue'
import { type Station } from '@/httpclient'
import { useDashboardStore } from '@/stores/DashboardStore'
import { ref } from 'vue'
import { formatDisplayNumber } from '@/utils/utils'

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
dashboardStore.loadSensors()
</script>
