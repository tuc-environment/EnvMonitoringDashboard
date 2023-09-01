<template>
  <div
    style="
      position: absolute;
      left: 0;
      right: 0;
      top: 0;
      bottom: 0;
      background-color: rgb(17, 18, 23);
      z-index: -100;
    "
  ></div>

  <div class="view" style="background-color: rgb(17, 18, 23)">
    <NavBarComponent class="navbar" />

    <div class="content d-flex text-light">
      <div class="main flex-grow-1 p-3 d-flex w-100">
        <slot />
      </div>
    </div>
  </div>

  <!-- <div
    style="
      position: absolute;
      left: 0;
      right: 0;
      top: 0;
      bottom: 0;
      background-color: rgb(17, 18, 23);
      z-index: -100;
    "
  ></div>

  <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <div class="container-fluid">
      <a class="navbar-brand" href="/">
        <div
          style="
            border: 0px solid white;
            border-radius: 4px;
            padding: 5px;
            background-color: white;
            background-image: url('/logo.png');
            background-size: cover;
            width: 30px;
            height: 30px;
          "
          class="d-inline-block align-top"
        ></div>
        <span class="small ms-2">环境监测系统 - 天津商业大学</span>
      </a>

      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarNav"
        aria-controls="navbarNav"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarNav">
        <div class="me-auto"></div>
        <ul class="navbar-nav me-2">
          <li class="nav-item">
            <a class="nav-link" href="/">主页</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="/cross-station">站点对比</a>
          </li>
        </ul>
        <form class="form-inline my-2">
          <button class="btn btn-sm btn-outline-primary my-2 my-sm-0" type="submit">
            登入管理后台
          </button>
        </form>
      </div>
    </div>
  </nav> -->

  <!-- <div
    class="p-3 text-light"
    style="
      position: absolute;
      left: 0;
      right: 0;
      top: 68px;
      bottom: 0;
      background-color: rgb(17, 18, 23);
      z-index: 1000;
      overflow-y: auto;
    "
  >
    <slot />
  </div> -->
</template>

<script setup lang="ts">
import NavBarComponent from '@/components/NavBarComponent.vue'
</script>

<style scoped lang="scss">
.navbar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
}

.content {
  margin-top: 56px;
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
  .view {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
  }

  .content {
    height: calc(100vh - 56px);
  }

  .sidebar {
    overflow: auto;
    width: 300px;
  }

  .main {
    overflow: auto;
  }
}
</style>
