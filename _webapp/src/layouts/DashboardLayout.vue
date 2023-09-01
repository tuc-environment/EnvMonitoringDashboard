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

  <div style="background-color: rgb(17, 18, 23)">
    <div
      class="d-flex text-light p-2 align-items-center fixed-top"
      style="
        background-color: rgb(24, 27, 31);
        height: 48px;
        border-bottom: 1px solid rgba(204, 204, 220, 0.12);
      "
    >
      <div>
        <img class="logo rounded me-2" src="/logo.png" style="height: 28px" />
      </div>
      <div>环境监测系统 - 天津商业大学</div>
      <div class="flex-grow-1"></div>
      <div>
        <a href="/login" class="text-light">Sign in</a>
      </div>
    </div>

    <slot />
  </div>
</template>
