<template>
  <div class="vh-100 w-100 p-2 text-light" style="background-color: rgb(0.1, 0.1, 0.1)">
    <div class="d-flex flex-column h-100 w-100">
      <h3>环境监测系统 - 天津商业大学</h3>
      <div class="mt-2 h-100 w-100 flex-grow-1">
        <div class="d-flex h-100 w-100">
          <div class="left-panel me-2 p-2 rounded bg-dark">Style 1</div>
          <div class="flex-grow-1 p-2 rounded bg-dark">
            <MapContainer />
          </div>
          <div class="right-panel ms-2 p-2 rounded bg-dark">Style 3</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import MapContainer from '@/components/map/MapContainer.vue'

export default {
  components: {
    MapContainer
  }
}
</script>

<style scoped>
.left-panel {
  width: 300px;
}

.right-panel {
  width: 300px;
}
</style>
