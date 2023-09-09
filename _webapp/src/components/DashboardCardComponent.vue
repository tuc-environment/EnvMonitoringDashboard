<template>
  <div class="dashboard-card w-100 h-100 p-2 d-flex flex-column">
    <div>
      <b>{{ title }}</b>
    </div>
    <div class="my-2 w-100 h-100 text-center flex-grow-1 d-flex flex-column justify-content-center">
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
  background-color: rgb(24, 27, 31);
  border: 1px solid rgba(204, 204, 220, 0.12);
  border-radius: 2px;
  min-height: 168px;
}
</style>
