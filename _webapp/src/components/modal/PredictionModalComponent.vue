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
          <div v-if="!showingResult" class="p-2">
            <div class="row mb-2">
              <div class="col-md-6">
                <LabelInputComponent
                  label="纬度"
                  type="number"
                  :required="false"
                  :disabled="true"
                  placeholder="纬度"
                  unit="°"
                  :field="`${latVal}`"
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
                  :field="`${lngVal}`"
                />
              </div>

              <div class="col-md-6">
                <LabelInputComponent
                  label="空气温度"
                  type="number"
                  :required="true"
                  placeholder="输入空气温度"
                  unit="°C"
                  @input="onTempChanged"
                />
              </div>
              <div class="col-md-6">
                <LabelInputComponent
                  label="空气湿度"
                  type="number"
                  :required="true"
                  placeholder="输入空气湿度"
                  unit="%"
                  @input="onHumidityChanged"
                />
              </div>
              <div class="col-md-12">
                <LabelInputComponent
                  label="大气压强"
                  type="number"
                  :required="true"
                  placeholder="输入大气压强"
                  unit="MPar"
                  @input="onBarometricPressureChanged"
                />
              </div>
              <div class="col-md-6">
                <LabelInputComponent
                  label="土壤浅层温度"
                  type="number"
                  :required="true"
                  placeholder="输入土壤浅层温度"
                  unit="°C"
                  @input="onSoilTempShallowChanged"
                />
              </div>
              <div class="col-md-6">
                <LabelInputComponent
                  label="土壤浅层含水量"
                  type="number"
                  :required="true"
                  placeholder="输入土壤浅层含水量"
                  unit="%"
                  @input="onSoilWaterContentShallowChanged"
                />
              </div>
              <div class="col-md-6">
                <LabelInputComponent
                  label="土壤深层温度"
                  type="number"
                  :required="true"
                  placeholder="输入土壤深层温度"
                  unit="°C"
                  @input="onSoilTempDeepChanged"
                />
              </div>
              <div class="col-md-6">
                <LabelInputComponent
                  label="土壤深层含水量"
                  type="number"
                  :required="true"
                  placeholder="输入土壤深层含水量"
                  unit="%"
                  @input="onSoilWaterContentDeepChanged"
                />
              </div>
              <div class="col-md-12">
                <LabelInputComponent
                  label="土壤导电率"
                  type="number"
                  :required="true"
                  placeholder="输入土壤导电率"
                  unit="mS/cm"
                  @input="onSoilElectricalConductivityChanged"
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
          <div v-else class="p-2">
            <h5>预测结果</h5>
            <table class="table table-striped my-3">
              <tbody>
                <tr>
                  <td width="20%">板下浅层土壤温度:</td>
                  <td>
                    {{
                      predictionResult?.down_soil_temp_shallow == null
                        ? 'N.A.'
                        : predictionResult?.down_soil_temp_shallow.toFixed(2)
                    }}℃
                  </td>
                </tr>
                <tr>
                  <td>板下深层土壤温度:</td>
                  <td>
                    {{
                      predictionResult?.down_soil_temp_deep == null
                        ? 'N.A.'
                        : predictionResult?.down_soil_temp_deep.toFixed(2)
                    }}℃
                  </td>
                </tr>
                <tr>
                  <td width="70%">板下浅层土壤湿度:</td>
                  <td>
                    {{
                      predictionResult?.down_soil_water_content_shallow == null
                        ? 'N.A.'
                        : predictionResult?.down_soil_water_content_shallow.toFixed(2)
                    }}%
                  </td>
                </tr>
                <tr>
                  <td>板下深层土壤湿度:</td>
                  <td>
                    {{
                      predictionResult?.down_soil_water_content_deep == null
                        ? 'N.A.'
                        : predictionResult?.down_soil_water_content_deep.toFixed(2)
                    }}%
                  </td>
                </tr>
                <tr>
                  <td width="20%">板下大气温度:</td>
                  <td>
                    {{
                      predictionResult?.down_temp == null
                        ? 'N.A.'
                        : predictionResult?.down_temp.toFixed(2)
                    }}℃
                  </td>
                </tr>
                <tr>
                  <td>板下大气湿度:</td>
                  <td>
                    {{
                      predictionResult?.down_humidity == null
                        ? 'N.A.'
                        : predictionResult?.down_humidity.toFixed(2)
                    }}%
                  </td>
                </tr>

                <tr>
                  <td width="20%">板间浅层土壤温度:</td>
                  <td>
                    {{
                      predictionResult?.middle_soil_temp_shallow == null
                        ? 'N.A.'
                        : predictionResult?.middle_soil_temp_shallow.toFixed(2)
                    }}℃
                  </td>
                </tr>
                <tr>
                  <td>板间深层土壤温度:</td>
                  <td>
                    {{
                      predictionResult?.middle_soil_temp_deep == null
                        ? 'N.A.'
                        : predictionResult?.middle_soil_temp_deep.toFixed(2)
                    }}℃
                  </td>
                </tr>
                <tr>
                  <td width="20%">板间浅层土壤湿度:</td>
                  <td>
                    {{
                      predictionResult?.middle_soil_water_content_shallow == null
                        ? 'N.A.'
                        : predictionResult?.middle_soil_water_content_shallow.toFixed(2)
                    }}%
                  </td>
                </tr>
                <tr>
                  <td>板间深层土壤湿度:</td>
                  <td>
                    {{
                      predictionResult?.middle_soil_water_content_deep == null
                        ? 'N.A.'
                        : predictionResult?.middle_soil_water_content_deep.toFixed(2)
                    }}%
                  </td>
                </tr>
                <tr>
                  <td width="20%">板间大气温度:</td>
                  <td>
                    {{
                      predictionResult?.middle_temp == null
                        ? 'N.A.'
                        : predictionResult?.middle_temp.toFixed(2)
                    }}℃
                  </td>
                </tr>
                <tr>
                  <td>板间大气湿度:</td>
                  <td>
                    {{
                      predictionResult?.middle_humidity == null
                        ? 'N.A.'
                        : predictionResult?.middle_humidity.toFixed(2)
                    }}%
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import LabelInputComponent from '@/components/LabelInputComponent.vue'
import httpclient, { type StationPrediction } from '@/httpclient'
import { load } from '@amap/amap-jsapi-loader'

const props = defineProps({
  title: { type: String, default: 'Modal Title' },
  visible: { type: Boolean, default: false }
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
  const params = {
    lat: latVal.value,
    lng: lngVal.value,
    temp: tempVal.value,
    humidity: humidityVal.value,
    barometric_pressure: barometricPressureVal.value,
    soil_temp_shallow: soilTempShallowVal.value,
    soil_temp_deep: soilTempDeepVal.value,
    soil_water_content_shallow: soilWaterContentShallowVal.value,
    soil_water_content_deep: soilWaterContentDeepVal.value,
    soil_electrical_conductivity: soilElectricalConductivityVal.value
  }
  console.log(params)
  const result = await httpclient.predictStation(params)
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
  setCoordinate
})
</script>

<style scoped>
.hidden {
  visibility: hidden;
}
</style>
