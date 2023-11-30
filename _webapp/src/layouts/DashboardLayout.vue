<template>
  <div style="position: fixed; width: 100%; height: 100%; left: 0; top: 0; z-index: -2000"></div>

  <div class="view">
    <NavBarComponent class="navbar" />

    <div class="content d-flex">
      <div class="main flex-grow-1 w-100 h-100">
        <div class="container h-100" style="max-width: 1660px">
          <slot />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import NavBarComponent from '@/components/NavBarComponent.vue'
</script>

<style scoped lang="scss">
.view {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}

.content {
  height: calc(100vh - 59px);
  overflow: auto;
}

.navbar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
}

.content {
  margin-top: 59px;
}

.sidebar {
  width: 100%;
  border-right: 1px solid #e0e0e0;
  background-color: $tux-sidebar-bg;
  min-width: 300px;
}

.main {
  min-height: 50vh;
}

@media (min-width: 768px) {
  .sidebar {
    overflow: auto;
    width: 300px;
  }

  .main {
    overflow: auto;
  }
}
</style>
