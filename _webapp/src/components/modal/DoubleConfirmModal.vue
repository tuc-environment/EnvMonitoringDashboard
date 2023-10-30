<template>
  <div
    class="modal fade"
    :class="{ show: visible, hidden: !visible }"
    tabindex="-1"
    style="display: block; background-color: rgba(0, 0, 0, 0.5)"
  >
    <div class="modal-dialog modal-dialog-centered modal-dialog-scrollable">
      <div class="modal-content">
        <div v-if="title" class="modal-header">
          <h5 class="modal-title" id="staticBackdropLabel">{{ title }}</h5>
          <button type="button" class="btn-close" @click="$emit('close')"></button>
        </div>
        <div class="modal-body">
          <div class="container p-2">
            <div v-if="content" class="row mb-2">
              <div class="col-md-12">
                <p>{{ content }}</p>
              </div>
            </div>
            <div class="row mb-2 justify-content-around">
              <button
                type="button"
                class="btn btn-secondary btn-block col-5"
                :disabled="loading"
                @click="$emit('close')"
              >
                取消
              </button>
              <button
                type="button"
                class="btn btn-danger btn-block col-5"
                :disabled="loading"
                @click="onConfirm"
              >
                <div v-if="loading">
                  <span
                    class="spinner-grow spinner-grow-sm"
                    role="status"
                    aria-hidden="true"
                  ></span>
                  Loading...
                </div>
                <div v-if="!loading">确定</div>
              </button>
            </div>
            <div style="text-align: center" class="mb-2"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

defineProps({
  title: { type: String },
  content: { type: String },
  visible: { type: Boolean, default: false }
})

const emit = defineEmits<{
  (e: 'onConfirm'): void
  (e: 'close'): void
}>()

const loading = ref(false)

const onConfirm = async () => {
  loading.value = true
  await emit('onConfirm')
  loading.value = false
}
</script>

<style scoped>
.hidden {
  visibility: hidden;
}
</style>
