<template>
  <div class="dashboardContainer">
    <div class="leftPanel">
      <button style="background-color: green; width: 100px; height: 100px" @click="test">
        test
      </button>
    </div>
    <div class="centerSpace"></div>
    <div class="rightPanel"></div>
  </div>
</template>

<script setup lang="ts">
const test = () => {
  console.log('testing')
}
</script>

<style scoped>
.dashboardContainer {
  position: absolute;
  top: 0;
  left: 0;
  display: flex;
  flex-direction: row;
  padding: 0px;
  margin: 0px;
  width: 100%;
  height: calc(100vh);
  pointer-events: none;
}

.leftPanel {
  width: 200px;
  min-width: 200px;
  height: 100%;
  background-color: red;
  flex: 1;
  pointer-events: initial;
}

.centerSpace {
  flex: 1.7;
  min-width: 200px;
}

.rightPanel {
  width: 200px;
  min-width: 200px;
  height: 100%;
  background-color: blue;
  flex: 1;
  pointer-events: initial;
}
</style>
