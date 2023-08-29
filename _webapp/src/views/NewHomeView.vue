<template>
  <div class="my-body">
    <div
      class="my-navbar d-flex text-light p-2 align-items-center"
      style="background-color: rgb(24, 27, 31); height: 48px"
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

    <div class="p-3 text-light">
      <div class="row">
        <div class="col-md-3 h-100 my-2">
          <DashboardCardComponent title="站点数量">
            <div class="text-danger" style="font-size: 5rem">5</div>
          </DashboardCardComponent>
        </div>
        <div class="col-md-3 h-100 my-2">
          <DashboardCardComponent title="传感器数量">
            <div class="text-danger" style="font-size: 5rem">189</div>
          </DashboardCardComponent>
        </div>
        <div class="col-md-3 h-100 my-2">
          <DashboardCardComponent title="今日报警数量">
            <div class="text-danger" style="font-size: 5rem">822</div>
          </DashboardCardComponent>
        </div>
        <div class="col-md-3 h-100 my-2">
          <DashboardCardComponent title="今日更新数量">
            <div class="text-danger" style="font-size: 5rem">1822</div>
          </DashboardCardComponent>
        </div>

        <div class="col-md-6 h-100 my-2">
          <DashboardCardComponent title="选择传感器">选择传感器</DashboardCardComponent>
        </div>
        <div class="col-md-6 h-100 my-2">
          <DashboardCardComponent title="传感器数据">
            <LineChart
              :sensors="airRelatedSensors"
              :records="airRelatedRecords"
              :show-default-text="!selectedStation"
              :loading="false"
              :title="selectedStation ? `${selectedStation.name} - 空气参数` : '空气参数'"
              default-text="请选择站点"
              no-data-text="暂无数据"
            />
          </DashboardCardComponent>
        </div>
        <div class="col-md-12 h-100 my-2">
          <DashboardCardComponent title="传感器地理位置">
            <MapContainer />
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

<script lang="ts">
import DashboardCardComponent from '@/components/DashboardCardComponent.vue'
import LineChart from '@/components/LineChart.vue'
import MapContainer from '@/components/map/MapContainer.vue'
import { SensorPosition, type DataRecord, type Sensor, type Station } from '@/httpclient'

export default {
  components: {
    DashboardCardComponent,
    LineChart,
    MapContainer
  },
  created() {
    this.airRelatedSensors = [
      {
        id: 1,
        station_id: 1,
        position: SensorPosition.up,
        tag: 'test',
        name: '空气',
        group: '1',
        unit: 'C'
      }
    ]
  },
  data() {
    return {
      airRelatedSensors: [] as Sensor[],
      airRelatedRecords: [] as DataRecord[],
      selectedStation: null as Station | null
    }
  }
}
</script>

<style scoped>
.my-body {
  background-color: rgb(17, 18, 23);
}

.my-navbar {
  border-bottom: 1px solid rgba(204, 204, 220, 0.12);
}
</style>
