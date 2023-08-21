<template>
  <div class="homePageContainer">
    <MapContainer />
    <Dashboard />
  </div>
</template>

<script setup lang="ts">
import MapContainer from '@/components/MapContainer.vue'
import Dashboard from '@/components/Dashboard.vue'
</script>

<style scoped>
.homePageContainer {
  background-color: black;
}
</style>
