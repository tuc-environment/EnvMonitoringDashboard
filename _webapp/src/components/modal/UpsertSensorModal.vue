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
                <div class="mb-2">
                  <label class="form-label small">数据项</label
                  ><span class="ms-1 text-danger">*</span>
                  <div class="input-group">
                    <select class="form-select" @input="onNameChanged">
                      <option :value="undefined">选择数据项</option>
                      <option
                        v-for="optionName in allOptionNames"
                        :value="optionName"
                        :selected="optionName == nameVal"
                      >
                        {{ optionName }}
                      </option>
                    </select>
                  </div>
                </div>
              </div>
              <div class="col-md-12">
                <div class="mb-2">
                  <label class="form-label small">传感器位置</label
                  ><span class="ms-1 text-danger">*</span>
                  <div class="input-group">
                    <select class="form-select" @input="onPostionChanged">
                      <option :value="undefined">选择传感器位置</option>
                      <option
                        v-for="position in positions"
                        :value="position"
                        :selected="position == positionVal"
                      >
                        {{ getPositionName(position) }}
                      </option>
                    </select>
                  </div>
                </div>
              </div>
              <div class="col-md-12">
                <LabelInputComponent
                  label="单位"
                  type="string"
                  placeholder="单位"
                  :field="unitVal"
                  @input="onUnitChanged"
                />
              </div>

              <div class="col-md-12">
                <LabelInputComponent
                  label="标签"
                  type="string"
                  placeholder="标签"
                  :field="tagVal"
                  @input="onTagChanged"
                />
              </div>

              <div class="col-md-12">
                <LabelInputComponent
                  label="组"
                  type="string"
                  placeholder="组"
                  :field="groupVal"
                  @input="onGroupChanged"
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
                  {{ sensorVal ? '更新中...' : '创建中...' }}
                </div>
                <div v-if="!loading">
                  {{ sensorVal ? '更新' : '创建' }}
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
import httpclient, { type Sensor } from '@/httpclient'
import { allOptionNames } from '@/utils/constants'
import { getPositionName, SensorPosition } from '@/httpclient'

const props = defineProps({
  title: { type: String, default: 'Modal Title' },
  visible: { type: Boolean, default: false }
})

const emit = defineEmits<{
  (e: 'didUpsertSensor'): void
}>()

const loading = ref(false)
const nameVal = ref<string | undefined>(undefined)
const positionVal = ref<SensorPosition | undefined>(undefined)
const positions: SensorPosition[] = [SensorPosition.up, SensorPosition.middle, SensorPosition.down]
const unitVal = ref<string | undefined>(undefined)
const tagVal = ref<string | undefined>(undefined)
const groupVal = ref<string | undefined>(undefined)
const stationIdVal = ref<number | undefined>(undefined)
const sensorVal = ref<Sensor | undefined>(undefined)

const onConfirm = async () => {
  loading.value = true
  try {
    const name = nameVal.value
    const position = positionVal.value
    const unit = unitVal.value
    const tag = tagVal.value
    const group = groupVal.value
    const stationId = stationIdVal.value
    const validName = name && name.trim().length > 0
    const validPosition = position != undefined
    const validStationId = stationIdVal.value != undefined

    if (validName && validPosition && validStationId) {
      const result = await httpclient.upsertSensor({
        station_id: stationId!,
        id: sensorVal.value?.id,
        name,
        position,
        tag,
        unit,
        group
      })
      if (result?.code == 200) {
        nameVal.value = undefined
        positionVal.value = undefined
        unitVal.value = undefined
        tagVal.value = undefined
        groupVal.value = undefined
        await emit('didUpsertSensor')
      }
    } else {
      alert('输入有误')
    }
  } catch (_) {}
  loading.value = false
}

const onNameChanged = (e: any) => {
  nameVal.value = e.target.value
}

const onPostionChanged = (e: any) => {
  positionVal.value = e.target.value
}

const onUnitChanged = (e: any) => {
  unitVal.value = e.target.value
}

const onTagChanged = (e: any) => {
  tagVal.value = e.target.value
}

const onGroupChanged = (e: any) => {
  groupVal.value = e.target.value
}

const setSensor = (stationId: number, sensor?: Sensor) => {
  console.log(
    `[upsert-sensor-modal] set sensor with stationId: ${stationId}, sensor: ${
      sensor ? JSON.stringify(sensor) : undefined
    }`
  )
  stationIdVal.value = stationId
  sensorVal.value = sensor
  nameVal.value = sensor?.name
  positionVal.value = sensor?.position
  tagVal.value = sensor?.tag
  groupVal.value = sensor?.group
}

defineExpose({
  setSensor
})
</script>

<style scoped>
.hidden {
  visibility: hidden;
}
</style>
