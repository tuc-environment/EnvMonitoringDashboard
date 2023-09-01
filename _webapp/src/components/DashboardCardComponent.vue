<template>
  <div class="dashboard-card w-100 h-100 p-2 d-flex flex-column">
    <div>
      <b>{{ title }}</b>
    </div>
    <div class="my-2 w-100 h-100 text-center flex-grow-1 d-flex flex-column justify-content-center">
      <slot />
    </div>
  </div>
</template>

<script lang="ts">
export default {
  props: {
    title: { type: String, default: 'Card Title' }
  }
}
</script>

<style scoped>
.dashboard-card {
  background-color: rgb(24, 27, 31);
  border: 1px solid rgba(204, 204, 220, 0.12);
  border-radius: 2px;
  min-height: 168px;
}
</style>
