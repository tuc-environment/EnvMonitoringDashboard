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
              <div class="col-md-6">
                <LabelInputComponent
                  label="纬度"
                  type="number"
                  :required="false"
                  :disabled="true"
                  placeholder="纬度"
                  unit="°"
                  :field="`${props.lat}`"
                />
              </div>
              <div class="col-md-6">
                <LabelInputComponent
                  label="经度"
                  type="number"
                  :required="false"
                  :disabled="true"
                  placeholder="经度"
                  unit="°"
                  :field="`${props.lng}`"
                />
              </div>

              <div class="col-md-6">
                <LabelInputComponent
                  label="空气温度"
                  type="number"
                  :required="false"
                  placeholder="输入空气温度"
                  unit="°C"
                />
              </div>
              <div class="col-md-6">
                <LabelInputComponent
                  label="空气湿度"
                  type="number"
                  :required="false"
                  placeholder="输入空气湿度"
                  unit="%"
                />
              </div>
              <div class="col-md-12">
                <LabelInputComponent
                  label="大气压强"
                  type="number"
                  :required="false"
                  placeholder="输入大气压强"
                  unit="MPar"
                />
              </div>
              <div class="col-md-6">
                <LabelInputComponent
                  label="土壤浅层温度"
                  type="number"
                  :required="false"
                  placeholder="输入土壤浅层温度"
                  unit="°C"
                />
              </div>
              <div class="col-md-6">
                <LabelInputComponent
                  label="土壤浅层含水量"
                  type="number"
                  :required="false"
                  placeholder="输入土壤浅层含水量"
                  unit="%"
                />
              </div>
              <div class="col-md-6">
                <LabelInputComponent
                  label="土壤深层温度"
                  type="number"
                  :required="false"
                  placeholder="输入土壤深层温度"
                  unit="°C"
                />
              </div>
              <div class="col-md-6">
                <LabelInputComponent
                  label="土壤深层含水量"
                  type="number"
                  :required="false"
                  placeholder="输入土壤深层含水量"
                  unit="%"
                />
              </div>
              <div class="col-md-12">
                <LabelInputComponent
                  label="土壤导电率"
                  type="number"
                  :disabled="true"
                  :required="false"
                  placeholder="输入土壤导电率"
                  unit="mS/cm"
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
                  Loading...
                </div>
                <div v-if="!loading">开始预测</div>
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

const props = defineProps({
  title: { type: String, default: 'Modal Title' },
  visible: { type: Boolean, default: false },
  lat: Number,
  lng: Number
})

const loading = ref(false)

const onConfirm = () => {
  loading.value = true
}
</script>

<style scoped>
.hidden {
  visibility: hidden;
}
</style>
