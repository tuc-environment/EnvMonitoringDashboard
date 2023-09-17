<template>
  <div class="modal fade" :class="{ show: visible, hidden: !visible }" tabindex="-1"
    style="display: block; background-color: rgba(0, 0, 0, 0.5)">
    <div class="modal-dialog modal-dialog-centered modal-dialog-scrollable">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="staticBackdropLabel">{{ title }}</h5>
          <button type="button" class="btn-close" @click="$emit('close')"></button>
        </div>
        <div class="modal-body">
          <div v-if="!showingResult" class="p-2">
            <div class="row mb-2">
              <div class="col-md-6">
                <LabelInputComponent label="纬度" type="number" :required="false" :disabled="true" placeholder="纬度" unit="°"
                  :field="`${latVal}`" />
              </div>
              <div class="col-md-6">
                <LabelInputComponent label="经度" type="number" :required="false" :disabled="true" placeholder="经度" unit="°"
                  :field="`${lngVal}`" />
              </div>

              <div class="col-md-6">
                <LabelInputComponent label="空气温度" type="number" :required="false" placeholder="输入空气温度" unit="°C"
                  @input="onTempChanged" />
              </div>
              <div class="col-md-6">
                <LabelInputComponent label="空气湿度" type="number" :required="false" placeholder="输入空气湿度" unit="%"
                  @input="onHumidityChanged" />
              </div>
              <div class="col-md-12">
                <LabelInputComponent label="大气压强" type="number" :required="false" placeholder="输入大气压强" unit="MPar"
                  @input="onBarometricPressureChanged" />
              </div>
              <div class="col-md-6">
                <LabelInputComponent label="土壤浅层温度" type="number" :required="false" placeholder="输入土壤浅层温度" unit="°C"
                  @input="onSoilTempShallowChanged" />
              </div>
              <div class="col-md-6">
                <LabelInputComponent label="土壤浅层含水量" type="number" :required="false" placeholder="输入土壤浅层含水量" unit="%"
                  @input="onSoilWaterContentShallowChanged" />
              </div>
              <div class="col-md-6">
                <LabelInputComponent label="土壤深层温度" type="number" :required="false" placeholder="输入土壤深层温度" unit="°C"
                  @input="onSoilTempDeepChanged" />
              </div>
              <div class="col-md-6">
                <LabelInputComponent label="土壤深层含水量" type="number" :required="false" placeholder="输入土壤深层含水量" unit="%"
                  @input="onSoilWaterContentDeepChanged" />
              </div>
              <div class="col-md-12">
                <LabelInputComponent label="土壤导电率" type="number" :disabled="true" :required="false" placeholder="输入土壤导电率"
                  unit="mS/cm" @input="onSoilElectricalConductivityChanged" />
              </div>
            </div>

            <div style="text-align: center" class="mb-2">
              <button type="button" class="btn btn-success" :disabled="loading" @click="onConfirm">
                <div v-if="loading">
                  <span class="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>
                  Loading...
                </div>
                <div v-if="!loading">开始预测</div>
              </button>
            </div>
          </div>
          <div v-else class="p-2">
            <div class="row mb-2">
              <div class="col-md-12">
                空气温度(°C): {{ predictionResult?.temp ?? 'N.A.' }}
              </div>
              <div class="col-md-6">
                空气湿度(%): {{ predictionResult?.humidity ?? 'N.A.' }}
              </div>
              <div class="col-md-12">
                大气压强(mbar): {{ predictionResult?.barometric_pressure ?? 'N.A.' }}
              </div>
              <div class="col-md-6">
                土壤浅层温度(°C): {{ predictionResult?.soil_temp_shallow ?? 'N.A.' }}
              </div>
              <div class="col-md-6">
                土壤浅层含水量(%): {{ predictionResult?.soil_temp_deep ?? 'N.A.' }}
              </div>
              <div class="col-md-6">
                土壤深层温度(°C): {{ predictionResult?.soil_water_content_shallow ?? 'N.A.' }}
              </div>
              <div class="col-md-6">
                土壤深层含水量(%): {{ predictionResult?.soil_water_content_deep ?? 'N.A.' }}
              </div>
              <div class="col-md-6">
                土壤电导率(μS/m): {{ predictionResult?.soil_electrical_conductivity ?? 'N.A.' }}
              </div>
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
import httpclient, { type StationPrediction } from '@/httpclient';
import { load } from '@amap/amap-jsapi-loader';

const props = defineProps({
  title: { type: String, default: 'Modal Title' },
  visible: { type: Boolean, default: false },
})

const loading = ref(false)
const showingResult = ref(false)
const latVal = ref<number | undefined>(undefined)
const lngVal = ref<number | undefined>(undefined)
const tempVal = ref<number | undefined>(undefined)
const humidityVal = ref<number | undefined>(undefined)
const barometricPressureVal = ref<number | undefined>(undefined)
const soilTempShallowVal = ref<number | undefined>(undefined)
const soilTempDeepVal = ref<number | undefined>(undefined)
const soilWaterContentShallowVal = ref<number | undefined>(undefined)
const soilWaterContentDeepVal = ref<number | undefined>(undefined)
const soilElectricalConductivityVal = ref<number | undefined>(undefined)
const predictionResult = ref<StationPrediction | undefined>(undefined)

const onTempChanged = (e: any) => {
  tempVal.value = +e.target.value
}

const onHumidityChanged = (e: any) => {
  humidityVal.value = +e.target.value
}

const onBarometricPressureChanged = (e: any) => {
  barometricPressureVal.value = +e.target.value
}

const onSoilTempShallowChanged = (e: any) => {
  soilTempShallowVal.value = +e.target.value
}

const onSoilWaterContentShallowChanged = (e: any) => {
  soilWaterContentShallowVal.value = +e.target.value
}

const onSoilTempDeepChanged = (e: any) => {
  soilTempDeepVal.value = +e.target.value
}

const onSoilWaterContentDeepChanged = (e: any) => {
  soilWaterContentDeepVal.value = +e.target.value
}

const onSoilElectricalConductivityChanged = (e: any) => {
  soilElectricalConductivityVal.value = +e.target.value
}

const onConfirm = async () => {
  loading.value = true
  const result = await httpclient.predictStation({
    lat: latVal.value,
    lng: lngVal.value,
    temp: tempVal.value,
    humidity: humidityVal.value,
    barometric_pressure: barometricPressureVal.value,
    soil_temp_shallow: soilTempShallowVal.value,
    soil_temp_deep: soilTempDeepVal.value,
    soil_water_content_shallow: soilWaterContentShallowVal.value,
    soil_water_content_deep: soilWaterContentDeepVal.value,
    soil_electrical_conductivity: soilElectricalConductivityVal.value,
  })
  if (result?.code == 200) {
    predictionResult.value = result.payload
    showingResult.value = true
  }
  loading.value = false
}

const setCoordinate = (lat?: number, lng?: number) => {
  latVal.value = lat
  lngVal.value = lng
  tempVal.value = undefined
  humidityVal.value = undefined
  barometricPressureVal.value = undefined
  soilTempShallowVal.value = undefined
  soilTempDeepVal.value = undefined
  soilWaterContentShallowVal.value = undefined
  soilWaterContentDeepVal.value = undefined
  soilElectricalConductivityVal.value = undefined
  loading.value = false
  showingResult.value = false
  predictionResult.value = undefined
}

defineExpose({
  setCoordinate,
})
</script>

<style scoped>
.hidden {
  visibility: hidden;
}
</style>
