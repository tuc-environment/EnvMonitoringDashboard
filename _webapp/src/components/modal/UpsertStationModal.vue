<template>
  <div
    class="modal fade"
    :class="{ show: visible, hidden: !visible }"
    tabindex="-1"
    style="display: block; background-color: rgba(0, 0, 0, 0.5)"
  >
    <div class="modal-dialog modal-dialog-centered modal-dialog-scrollable">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="staticBackdropLabel">{{ title }}</h5>
          <button type="button" class="btn-close" @click="$emit('close')"></button>
        </div>
        <div class="modal-body">
          <div class="p-2">
            <div class="row mb-2">
              <div class="col-md-12">
                <LabelInputComponent
                  label="名字"
                  type="string"
                  :required="true"
                  placeholder="名字"
                  unit=""
                  :field="nameVal"
                  @input="onChangeName"
                />
              </div>
              <div class="col-md-12">
                <LabelInputComponent
                  label="纬度"
                  type="number"
                  :required="true"
                  placeholder="纬度"
                  unit="°"
                  :field="`${latVal}`"
                  @input="onChangeLat"
                />
              </div>
              <div class="col-md-12">
                <LabelInputComponent
                  label="经度"
                  type="number"
                  :required="true"
                  placeholder="经度"
                  unit="°"
                  :field="`${lngVal}`"
                  @input="onChangeLng"
                />
              </div>
              <div class="col-md-12">
                <LabelInputComponent
                  label="海拔"
                  type="number"
                  :required="true"
                  placeholder="海拔"
                  unit="m"
                  :field="`${altitudeVal}`"
                  @input="onChangeAltitude"
                />
              </div>
            </div>

            <div style="text-align: center" class="mb-2">
              <button type="button" class="btn btn-success" :disabled="loading" @click="onConfirm">
                <div v-if="loading">
                  <span
                    class="spinner-grow spinner-grow-sm"
                    role="status"
                    aria-hidden="true"
                  ></span>
                  {{ stationVal ? '更新中...' : '创建中...' }}
                </div>
                <div v-if="!loading">
                  {{ stationVal ? '更新' : '创建' }}
                </div>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import LabelInputComponent from '@/components/LabelInputComponent.vue'
import httpclient, { type Station } from '@/httpclient'

const props = defineProps({
  title: { type: String, default: 'Modal Title' },
  visible: { type: Boolean, default: false }
})

const emit = defineEmits<{
  (e: 'didUpsertStation'): void
}>()

const loading = ref(false)
const nameVal = ref<string | undefined>('')
const latVal = ref<number | undefined>(undefined)
const lngVal = ref<number | undefined>(undefined)
const altitudeVal = ref<number | undefined>(undefined)
const stationVal = ref<Station | undefined>(undefined)

const onChangeName = (e: any) => {
  nameVal.value = e.target.value
}

const onChangeLat = (e: any) => {
  latVal.value = +e.target.value
}

const onChangeLng = (e: any) => {
  lngVal.value = +e.target.value
}

const onChangeAltitude = (e: any) => {
  altitudeVal.value = +e.target.value
}

const onConfirm = async () => {
  loading.value = true
  try {
    const name = nameVal.value
    const lat = latVal.value
    const lng = lngVal.value
    const altitude = altitudeVal.value
    const id = stationVal.value?.id
    const validName = name && name.trim().length > 0
    const validLat = lat != undefined && lat > 3 && lat < 54
    const validLng = lng != undefined && lng > 73 && lng < 136
    const validAltitude = altitude != undefined
    if (validName && validLat && validLng && validAltitude) {
      const result = await httpclient.upsertStation({ id, name, lat, lng, altitude })
      if (result?.code == 200) {
        nameVal.value = undefined
        latVal.value = undefined
        lngVal.value = undefined
        altitudeVal.value = undefined
        stationVal.value = undefined
        await emit('didUpsertStation')
      }
    } else {
      alert('输入有误')
    }
  } catch (_) {}
  loading.value = false
}

const setStation = (station?: Station) => {
  stationVal.value = station
  nameVal.value = station?.name
  latVal.value = station?.lat
  lngVal.value = station?.lng
  altitudeVal.value = station?.altitude
}

defineExpose({
  setStation
})
</script>

<style scoped>
.hidden {
  visibility: hidden;
}
</style>
