<template>
  <div class="custom-btn d-flex align-items-center justify-content-center">
    <div>
      <div>{{ title }}</div>
      <div style="font-size: 1.5rem; color: green">{{ value }}</div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps({
  title: String,
  value: String
})
</script>

<style scoped lang="scss">
.custom-btn {
  color: white;
  text-align: center;
  width: 168px;
  height: 72px;
  background-image: url('./home/button2.png');
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-position: center;
}
</style>
