<template>
  <div class="container">
    <div class="d-flex flex-row my-1 align-items-center">
      <div class="h5 mr-3">Create new station:</div>
      <button type="button" class="btn btn-success" @click="onAddButtonClicked">Add</button>
    </div>
    <div v-if="!requesting" class="row">
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
              <div v-if="operatingStationId == station.id">
                <span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>
                修改中...
              </div>
              <div v-else>
                <button
                  type="button"
                  class="btn btn-outline-success btn-sm mx-2"
                  @click="viewSensorData(station.id)"
                >
                  查看所有传感器
                </button>
                <button
                  type="button"
                  class="btn btn-outline-primary btn-sm mx-2"
                  @click="showStationUpsertModal(station)"
                >
                  编辑
                </button>
                <button
                  type="button"
                  class="btn btn-outline-danger btn-sm mx-2"
                  @click="onDeleteStation(station.id)"
                >
                  删除
                </button>
              </div>
            </td>
          </tr>
        </tbody>
        <tfoot>
          <table-paginator
            :offset="offset"
            :limit="limit"
            :total="total"
            @to-previous="toPreviousPage"
            @to-index="toOffset"
            @to-next="toNextPage"
          />
        </tfoot>
      </table>
    </div>
    <upsert-station-modal
      ref="upsertStationModal"
      :visible="isShowingStationUpsertModal"
      title="Create Station"
      @did-upsert-station="onStationUpserted"
      @close="onModalClosed"
    />
    <double-confirm-modal
      :visible="isShowingDeleteCreationModal"
      :title="'删除站点 ' + operatingStationId"
      content="确定删除站点？"
      @close="onDeleteStationModelClosed"
      @on-confirm="confirmDeletingStation"
    />
  </div>
</template>

<script setup lang="ts">
import httpclient, { type Station } from '@/httpclient'
import { ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import UpsertStationModal from '@/components/modal/UpsertStationModal.vue'
import DoubleConfirmModal from '@/components/modal/DoubleConfirmModal.vue'
import { useLoading } from 'vue3-loading-overlay'
import TablePaginator from '@/components/TablePaginator.vue'

const offset = ref(0)
// temp: for testing
const limit = 2
const total = ref(0)
const loader = useLoading()
const requesting = ref(false)
const allStations = ref<Station[]>([])
const router = useRouter()
const isShowingStationUpsertModal = ref(false)
const isShowingDeleteCreationModal = ref(false)
const operatingStationId = ref<number | undefined>(undefined)
const upsertStationModal = ref<InstanceType<typeof UpsertStationModal> | null>(null)

watch(requesting, (newVal, _) => {
  if (newVal) {
    loader.show()
  } else {
    loader.hide()
  }
})

const loadStations = async (currentOffset: number, showLoadingIndicator: boolean) => {
  if (showLoadingIndicator) {
    requesting.value = true
  }
  const resp = await httpclient.getStations({ offset: currentOffset, limit: limit })
  allStations.value = resp?.payload || []
  total.value = resp?.total ?? 0
  requesting.value = false
}

const viewSensorData = (stationID: number) => {
  router.push({ query: { view: 'Sensors', station_id: stationID } })
}

// station update

const showStationUpsertModal = (station: Station) => {
  isShowingStationUpsertModal.value = true
  operatingStationId.value = station.id
  upsertStationModal.value?.setStation(station)
}

const onAddButtonClicked = () => {
  upsertStationModal.value?.setStation(undefined)
  isShowingStationUpsertModal.value = true
}

const onModalClosed = () => {
  console.log('[station-view] on station creation closed')
  isShowingStationUpsertModal.value = false
  operatingStationId.value = undefined
  upsertStationModal.value?.setStation(undefined)
}

const onStationUpserted = async () => {
  await loadStations(offset.value, false)
  isShowingStationUpsertModal.value = false
  operatingStationId.value = undefined
  upsertStationModal.value?.setStation(undefined)
}

// station deletion

const onDeleteStationModelClosed = () => {
  isShowingDeleteCreationModal.value = false
  operatingStationId.value = undefined
}

const onDeleteStation = (stationId: number | undefined) => {
  if (stationId) {
    operatingStationId.value = stationId
    isShowingDeleteCreationModal.value = true
  }
}

const confirmDeletingStation = async () => {
  try {
    const stationId = operatingStationId.value
    if (stationId) {
      await httpclient.deleteStation(stationId)
      await loadStations(offset.value, false)
      isShowingDeleteCreationModal.value = false
    }
  } catch (_) {}
  operatingStationId.value = undefined
}

const toPreviousPage = (offsetVal: number) => {
  offset.value = offsetVal
  loadStations(offset.value, false)
}

const toNextPage = (offsetVal: number) => {
  offset.value = offsetVal
  loadStations(offset.value, false)
}

const toOffset = (offsetVal: number) => {
  offset.value = offsetVal
  loadStations(offset.value, false)
}

loadStations(offset.value, true)
</script>
