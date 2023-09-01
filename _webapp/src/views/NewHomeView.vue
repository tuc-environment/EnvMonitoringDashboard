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
      <div class="row align-items-stretch">
        <div class="col-md-6 col-lg-4 col-xl-2 mb-2">
          <DashboardCardComponent title="站点数量">
            <div class="text-danger" style="font-size: 4rem">5</div>
          </DashboardCardComponent>
        </div>
        <div class="col-md-6 col-lg-4 col-xl-2 mb-2">
          <DashboardCardComponent title="传感器数量">
            <div class="text-danger" style="font-size: 4rem">189</div>
          </DashboardCardComponent>
        </div>
        <div class="col-md-6 col-lg-4 col-xl-2 mb-2">
          <DashboardCardComponent title="今日报警数量">
            <div class="text-danger" style="font-size: 4rem">822</div>
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

      <div class="row align-items-stretch">
        <div class="col-lg-12 col-xl-8 my-2">
          <DashboardCardComponent title="传感器地理位置" class="info-card">
            <map-container ref="mapContainer" @did-select-station="selectStationHandler" />
          </DashboardCardComponent>
        </div>
        <div class="col-lg-12 col-xl-4 my-2">
          <div class="h-100 d-flex flex-column justify-content-between">
            <div class="mb-2">
              <DashboardCardComponent
                :title="
                  dashboardStore.$state.selectedStation
                    ? `${dashboardStore.$state.selectedStation.name} - 空气参数`
                    : '空气参数'
                "
                class="info-card"
              >
                <line-chart
                  :sensors="dashboardStore.$state.airRelatedSensors"
                  :records="dashboardStore.$state.airRelatedRecords"
                  :show-default-text="!dashboardStore.$state.selectedStation"
                  :loading="dashboardStore.$state.loadingDataForStation"
                  default-text="请选择站点"
                  no-data-text="暂无数据"
                />
              </DashboardCardComponent>
            </div>
            <div>
              <DashboardCardComponent
                :title="
                  dashboardStore.$state.selectedStation
                    ? `${dashboardStore.$state.selectedStation.name} - 土壤参数`
                    : '土壤参数'
                "
                class="info-card"
              >
                <line-chart
                  :sensors="dashboardStore.$state.soilRelatedSensors"
                  :records="dashboardStore.$state.soilRelatedRecords"
                  :show-default-text="!dashboardStore.$state.selectedStation"
                  :loading="dashboardStore.$state.loadingDataForStation"
                  default-text="请选择站点"
                  no-data-text="暂无数据"
                />
              </DashboardCardComponent>
            </div>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col-lg-6 col-xl-3 h-100 my-1">
          <DashboardCardComponent title="选择传感器" class="info-card">
            <tag-group
              :tags="dashboardStore.treeSensorSelectedTags"
              @on-tag-selected="onSelectTag"
            />
            <tree-chart ref="treeChart" />
          </DashboardCardComponent>
        </div>
        <div class="col-lg-6 col-xl-3 h-100 my-1">
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
  <!-- <div class="vh-100 w-100 p-2 text-light" style="background-color: rgb(0.1, 0.1, 0.1)">
    <div class="d-flex flex-column h-100 w-100">
      <h3>环境监测系统 - 天津商业大学</h3>
      <div class="mt-2 h-100 w-100 flex-grow-1">
        <div class="d-flex h-100 w-100">
          <div class="left-panel me-2 p-2 rounded bg-dark">Style 1</div>
          <div class="flex-grow-1 p-2 rounded bg-dark">
            <MapContainer />
          </div>
          <div class="right-panel ms-2 p-2 rounded bg-dark">Style 3</div>
        </div>
      </div>
    </div>
  </div> -->
</template>

<script setup lang="ts">
import DashboardCardComponent from '@/components/DashboardCardComponent.vue'
import LineChart from '@/components/LineChart.vue'
import TreeChart from '@/components/tree/Tree.vue'
import MapContainer from '@/components/map/MapContainer.vue'
import { type Sensor, type Station } from '@/httpclient'
import { useDashboardStore } from '@/stores/DashboardStore'
import TagGroup from '@/components/tags/TagsGroup.vue'
import { ref } from 'vue'
import type { TagData } from '@/components/tags/TagData'

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

const onSelectTag = (tag: TagData) => {
  const sensor = tag.data as Sensor
  if (sensor) {
    dashboardStore.removeTreeNodeSelected(sensor)
  }
}

const selectStationHandler = (station: Station | undefined) => {
  dashboardStore.setMapSelectedStation(station)
}

dashboardStore.loadStations()
</script>

<!-- <style scoped>
.info-card {
  height: 960px;
}

@media (max-width: 1199.98px) {
  .info-card {
    height: 400px;
    min-height: 400px;
  }
}
</style> -->
