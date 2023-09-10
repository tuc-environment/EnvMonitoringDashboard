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
                  创建中...
                </div>
                <div v-if="!loading">创建</div>
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
import LabelInputComponent from './LabelInputComponent.vue'
import httpclient from '@/httpclient'

const props = defineProps({
  title: { type: String, default: 'Modal Title' },
  visible: { type: Boolean, default: false }
})

const emit = defineEmits<{
  (e: 'didCreateStation'): void
}>()

const loading = ref(false)
const nameVal = ref<string | undefined>('')
const latVal = ref<number | undefined>(undefined)
const lngVal = ref<number | undefined>(undefined)
const altitudeVal = ref<number | undefined>(undefined)

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
    const validName = name && name.trim().length > 0
    const validLat = lat != undefined && lat > 3 && lat < 54
    const validLng = lng != undefined && lng > 73 && lng < 136
    const validAltitude = altitude != undefined
    if (validName && validLat && validLng && validAltitude) {
      const result = await httpclient.upsertStation({ name, lat, lng, altitude })
      if (result?.code == 200) {
        nameVal.value = undefined
        latVal.value = undefined
        lngVal.value = undefined
        altitudeVal.value = undefined
        await emit('didCreateStation')
      }
    } else {
      alert('输入有误')
    }
  } catch (_) {}
  loading.value = false
}
</script>

<style scoped>
.hidden {
  visibility: hidden;
}
</style>
