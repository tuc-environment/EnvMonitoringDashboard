<template>
  <div class="dashboard-card w-100 h-100 p-2 d-flex flex-column">
    <div>
      <b>{{ title }}</b>
    </div>
    <div
      class="my-2 w-100 text-center flex-grow-1 d-flex flex-column justify-content-center"
      style="overflow: auto"
    >
      <div class="w-100 h-100" :class="{ 'd-none': loading }">
        <slot />
      </div>
      <div :class="{ 'd-none': !loading }">
        <div class="spinner-border" role="status"></div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  props: {
    title: { type: String, default: 'Card Title' },
    loading: { type: Boolean, default: false }
  }
}
</script>

<style scoped>
.dashboard-card {
  background: linear-gradient(144deg, rgba(225, 255, 177, 0.2) 0%, rgba(182, 227, 136, 0.15) 100%);
  border: 1px solid rgba(0, 0, 0, 0.2);
  border-radius: 5px;
  min-height: 168px;
}
</style>
