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
          <h5 class="modal-title" id="staticBackdropLabel">选择与该站点关联的传感器</h5>
          <button type="button" class="btn-close" @click="$emit('close')"></button>
        </div>
        <div class="modal-body">
          <div class="container p-2">
            <unlinked-sensor-tree
              ref="treeChart"
              @did-update-sensors="emit('onConfirm')"
              @on-cancel="emit('close')"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { type Station } from '@/http-client'
import { ref } from 'vue'
import UnlinkedSensorTree from './UnlinkedSensorTree.vue'

const treeChart = ref<InstanceType<typeof UnlinkedSensorTree> | null>(null)

defineProps({
  visible: { type: Boolean, default: false }
})

const emit = defineEmits<{
  (e: 'onConfirm'): void
  (e: 'close'): void
}>()

// const loading = ref(false)

const setStation = async (updatedStation: Station | undefined) => {
  treeChart.value?.setStation(updatedStation)
}

// const onConfirm = async () => {
//   loading.value = true
//   await emit('onConfirm')
//   loading.value = false
// }

defineExpose({
  setStation
})
</script>

<style scoped>
.hidden {
  visibility: hidden;
}
</style>
