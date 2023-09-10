<template>
  <div class="container">
    <div class="d-flex flex-row my-1 align-items-center">
      <div class="h5 mr-3">Create new station:</div>
      <button type="button" class="btn btn-success" @click="onAddButtonClicked">Add</button>
    </div>
    <div class="row">
      <table class="table table-bordered align-middle">
        <thead class="table-dark">
          <tr>
            <th scope="col" width="10%">#</th>
            <th scope="col" width="30%">名称</th>
            <th scope="col" width="10%">纬度</th>
            <th scope="col" width="10%">经度</th>
            <th scope="col" width="10%">海拔</th>
            <th scope="col" width="30%">操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="station in allStations" :key="station.id">
            <td>{{ station.id }}</td>
            <td>{{ station.name }}</td>
            <td>{{ station.lat }}</td>
            <td>{{ station.lng }}</td>
            <td>{{ station.altitude }}</td>
            <td>
              <button
                type="button"
                class="btn btn-outline-success btn-sm mx-2"
                @click="viewSensorData(station.id)"
              >
                查看所有传感器
              </button>
              <button type="button" class="btn btn-outline-primary btn-sm mx-2" disabled>
                编辑
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <create-station-modal
      :visible="showStationCreationModal"
      title="Create Station"
      @did-create-station="onStationCreated"
      @close="onModalClosed"
    />
  </div>
</template>

<script setup lang="ts">
import httpclient, { type Station } from '@/httpclient'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import CreateStationModal from '@/components/CreateStationModal.vue'

const requesting = ref(false)
const allStations = ref<Station[]>([])
const router = useRouter()
const showStationCreationModal = ref(false)

const loadStations = async () => {
  requesting.value = true
  const resp = await httpclient.getStations()
  allStations.value = resp?.payload || []
  requesting.value = false
}

const viewSensorData = (stationID: number) => {
  router.push({ query: { view: 'Sensors', station_id: stationID } })
}

const onAddButtonClicked = () => {
  showStationCreationModal.value = true
}

const onModalClosed = () => {
  console.log('[station-view] on station creation closed')
  showStationCreationModal.value = false
}

const onStationCreated = async () => {
  await loadStations()
  showStationCreationModal.value = false
}

loadStations()
</script>
