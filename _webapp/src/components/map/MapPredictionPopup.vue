<template>
  <div class="popupWrapper">
    <form class="d-flex flex-column align-items-stretch">
      <div class="form-group row">
        <div class="form-group col-md-6">
          <label for="latitude">纬度</label>
          <input
            type="number"
            class="form-control"
            id="latitude"
            placeholder="°"
            v-model="latVal"
          />
        </div>
        <div class="form-group col-md-6">
          <label for="longtitude">经度</label>
          <input
            type="number"
            class="form-control"
            id="longtitude"
            placeholder="° "
            v-model="lngVal"
          />
        </div>
      </div>
      <!-- <div class="form-group">
        <label for="temp">空气温度</label>
        <input type="number" class="form-control" id="temp" placeholder="℃" v-model="tempVal" />
      </div>
      <div class="form-group">
        <label for="humidity">空气湿度</label>
        <input
          type="number"
          class="form-control"
          id="humidity"
          placeholder="%"
          v-model="humidityVal"
        />
      </div>
      <div class="form-group">
        <label for="barometric-pressure">大气压强</label>
        <input
          type="number"
          class="form-control"
          id="barometric-pressure"
          placeholder="mbar"
          v-model="barometricPressureVal"
        />
      </div>
      <div class="form-group">
        <label for="soil-temp-shallow">土壤浅层温度</label>
        <input
          type="number"
          class="form-control"
          id="soil-temp-shallow"
          placeholder="℃"
          v-model="soilTempShallowVal"
        />
      </div>
      <div class="form-group">
        <label for="soil-temp-deep">土壤深层温度</label>
        <input
          type="number"
          class="form-control"
          id="soil-temp-deep"
          placeholder="℃"
          v-model="soilTempDeepVal"
        />
      </div>
      <div class="form-group">
        <label for="soil-water-content-shallow">土壤浅层含水量</label>
        <input
          type="number"
          class="form-control"
          id="soil-water-content-shallow"
          placeholder="%"
          v-model="soilWaterContentShallowVal"
        />
      </div>
      <div class="form-group">
        <label for="soil-water-content-deep">土壤深层含水量</label>
        <input
          type="number"
          class="form-control"
          id="soil-water-content-deep"
          placeholder="%"
          v-model="soilWaterContentDeepVal"
        />
      </div>
      <div class="form-group">
        <label for="soil-electrical-conductivity">土壤电导率</label>
        <input
          type="number"
          class="form-control"
          id="soil-electrical-conductivity"
          placeholder="µS/m"
          v-model="soilElectricalConductivityVal"
        />
      </div> -->
      <button
        type="submit"
        class="btn btn-primary submit-btn"
        :disabled="loading"
        @click="onSubmit"
      >
        <div v-if="loading">
          <span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>
          Loading...
        </div>
        <div v-if="!loading">开始预测</div>
      </button>
    </form>
  </div>
</template>
<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'

const props = defineProps({
  lat: Number,
  lng: Number
})

const latVal = ref(0)
const lngVal = ref(0)
// const tempVal = ref<number | undefined>(undefined)
// const humidityVal = ref<number | undefined>(undefined)
// const barometricPressureVal = ref<number | undefined>(undefined)
// const soilTempShallowVal = ref<number | undefined>(undefined)
// const soilTempDeepVal = ref<number | undefined>(undefined)
// const soilWaterContentShallowVal = ref<number | undefined>(undefined)
// const soilWaterContentDeepVal = ref<number | undefined>(undefined)
// const soilElectricalConductivityVal = ref<number | undefined>(undefined)
const loading = ref(false)

// eslint-disable-next-line @typescript-eslint/no-unused-vars
watch(latVal, (newVal, oldVal) => {
  onPositionChanged()
})

// eslint-disable-next-line @typescript-eslint/no-unused-vars
watch(lngVal, (newVal, oldVal) => {
  onPositionChanged()
})

onMounted(() => {
  console.log(`[map-predict-popup] load popup with lat: ${props.lat}, lng: ${props.lng}`)
  latVal.value = props.lat ?? 0
  lngVal.value = props.lng ?? 0
})

const emit = defineEmits<{
  (e: 'onPositionChanged', lng: number, lat: number): void
  (e: 'onConfirm', lng: number, lat: number): void
}>()

const onPositionChanged = () => {
  emit('onPositionChanged', lngVal.value, latVal.value)
}

const onSubmit = () => {
  loading.value = true
  // const temp = tempVal.value
  // const humidity = humidityVal.value
  emit('onConfirm', lngVal.value, latVal.value)
}
</script>

<style scoped>
.popupWrapper {
  position: relative;
  z-index: 1;
  width: 300px;
  padding: 8px;
  overflow: hidden;
}

.form-group {
  margin-top: 8px;
}

.form-label {
  flex: 1;
}

.form-control {
  flex: 2;
}

.submit-btn {
  margin-top: 16px;
}
</style>
