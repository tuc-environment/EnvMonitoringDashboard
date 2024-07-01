<template>
  <div class="custom-btn d-flex align-items-center justify-content-center" @click="onClick">
    {{ title }}
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const props = defineProps({
  title: String,
  link: String
})

function onClick() {
  window.location.href = props.link!
}
</script>

<style scoped lang="scss">
.custom-btn {
  color: white;
  text-align: center;
  width: 108px;
  height: 32px;
  background-image: url('./home/button.png');
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-position: center;
}

.custom-btn:hover {
  cursor: pointer;
  background-image: url('./home/button-hover.png');
}
</style>
