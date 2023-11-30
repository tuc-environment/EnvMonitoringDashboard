<template>
  <DashboardLayout>
    <div>
      <div
        class="w-100 background-image d-flex align-items-center"
        style="
          aspect-ratio: 21/9;
          min-height: 200px;
          max-height: 460px;
          background-image: url('/all_banner.jpg');
        "
      >
        <div class="container text-white" style="z-index: 100">
          <div class="w-100" style="max-width: 500px">
            <div style="font-size: 4rem">
              <strong>联系我们</strong>
            </div>
            <div style="font-size: 1.2rem">天津商业大学 | 中科院电工研究所 | 天津大学</div>
          </div>
        </div>
      </div>

      <div class="container my-5">
        <div class="row align-items-center my-3">
          <div class="col-md-6 my-2 text-center">
            <img class="rounded w-100" src="/tianjin-map.png" style="aspect-ratio: 16/9" />
          </div>
          <div class="col-md-6 my-2">
            <div class="h1">天津商业大学</div>
            <div class="my-3">
              <div class="h5">地址</div>
              <div>天津市北辰区光荣道409号</div>
            </div>
            <div class="my-3">
              <div class="h5">电话</div>
              <div>(022) 2667 5774</div>
            </div>
            <div class="my-3">
              <div class="h5">邮箱</div>
              <div>president@tjcu.edu.cn</div>
            </div>
          </div>
        </div>
      </div>

      <div class="container my-5">
        <div class="row align-items-center my-3">
          <div class="col-md-6 my-2 text-center">
            <img class="rounded w-100" src="/diangong-map.png" style="aspect-ratio: 16/9" />
          </div>
          <div class="col-md-6 my-2">
            <div class="h1">中科院电工研究所</div>
            <div class="my-3">
              <div class="h5">地址</div>
              <div>北京市海淀区中关村北二条六号</div>
            </div>
            <div class="my-3">
              <div class="h5">电话</div>
              <div>(010) 8254 7001</div>
            </div>
            <div class="my-3">
              <div class="h5">邮箱</div>
              <div>office@mail.iee.ac.cn</div>
            </div>
          </div>
        </div>
      </div>

      <div class="container my-5">
        <div class="row align-items-center my-3">
          <div class="col-md-6 my-2 text-center">
            <img class="rounded w-100" src="/tianjin-map.png" style="aspect-ratio: 16/9" />
          </div>
          <div class="col-md-6 my-2">
            <div class="h1">天津大学</div>
            <div class="my-3">
              <div class="h5">地址</div>
              <div>天津市南开区卫津路92号</div>
            </div>
            <div class="my-3">
              <div class="h5">网站</div>
              <div><a href="https://www.tju.edu.cn/">https://www.tju.edu.cn/</a></div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div style="height: 64px"></div>
  </DashboardLayout>
</template>

<script setup lang="ts">
import DashboardLayout from '@/layouts/DashboardLayout.vue'
</script>

<style scoped>
.background-image {
  z-index: 0;
  position: relative;
  background-size: cover;
  background-position: center;
}

.background-image::before {
  z-index: 0;
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(
    180deg,
    rgba(29, 31, 42, 0.9) 1%,
    rgba(29, 31, 42, 0.45) 34%,
    rgba(29, 31, 42, 0.45) 66%,
    rgba(29, 31, 42, 0.9) 99%
  );
}
</style>
