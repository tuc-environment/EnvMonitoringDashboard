<template>
  <BackgroundComponent />

  <DashboardLayout>
    <div class="container">
      <div class="mx-auto my-5 login py-3 px-5">
        <div class="d-flex justify-content-center w-100">
          <img class="logo" src="/tianjin-logo.png" />
          <img class="logo" src="/diangong-logo.png" />
        </div>

        <h2 class="mt-3 text-center">环境监测系统 API</h2>
        <div class="text-center">天津商业大学 | 中科院电工研究所</div>

        <div class="my-3 mb-5 text-center">
          新用户？<RouterLink to="/register" class="register-link">注册账号</RouterLink>
        </div>

        <div v-if="errorMsg" class="alert alert-danger" role="alert">{{ errorMsg }}</div>

        <form @submit.prevent="preventDefault">
          <div class="form-group my-3">
            <label class="mb-2">用户名</label>
            <input type="text" class="form-control" v-model="username" :disabled="requesting" />
          </div>

          <div class="form-group my-3">
            <label class="mb-2">密码</label>
            <input type="password" class="form-control" v-model="password" :disabled="requesting" />
          </div>

          <button
            type="submit"
            class="mt-3 w-100 btn btn-lg btn-primary"
            :disabled="requesting"
            @click="login(username, password)"
          >
            <div
              v-if="requesting"
              class="spinner-border text-light mx-auto d-block"
              role="status"
            />
            <div v-else>登入</div>
          </button>

          <div
            class="mt-5 d-flex justify-content-center mb-2"
            style="color: rgba(0, 0, 0, 0.5); font-size: 0.99rem"
          >
            <a
            class="mx-2"
              style="cursor: pointer; text-decoration: none"
              href="https://www.tjcu.edu.cn/"
              target="_blank"
              >天津商业大学 <i class="bi bi-box-arrow-up-right"></i
            ></a>
            <a
            class="mx-2"
              style="cursor: pointer; text-decoration: none"
              href="http://www.iee.cas.cn/"
              target="_blank"
              >中科院电工研究所 <i class="bi bi-box-arrow-up-right"></i
            ></a>
          </div>
        </form>
      </div>
    </div>

    <!-- magic padding -->
    <div style="height: 1px"></div>
  </DashboardLayout>
</template>

<script setup lang="ts">
import BackgroundComponent from '@/components/BackgroundComponent.vue'
import httpclient from '@/http-client'
import DashboardLayout from '@/layouts/DashboardLayout.vue'
import { ref } from 'vue'
import { RouterLink, useRouter } from 'vue-router'

const errorMsg = ref('')
const username = ref('')
const password = ref('')
const requesting = ref(false)
const router = useRouter()

if (httpclient.token) {
  router.push('/admin')
}

const login = async (username: string, password: string) => {
  requesting.value = true
  const resp = await httpclient.login(username, password)
  if (resp?.code == 200) {
    router.push('/admin')
  } else if (resp?.code == 400) {
    clearForm()
    errorMsg.value = '账号或密码错误!'
  }
  requesting.value = false
}

const clearForm = () => {
  username.value = ''
  password.value = ''
}

function preventDefault(e: Event) {
  e.preventDefault()
}
</script>

<style scoped>
.login {
  width: 100%;
  max-width: 480px;
  background-color: rgb(255, 255, 255);
  border-radius: 4px;
}

.logo {
  width: 96px;
  display: block;
  border-radius: 5px;
}

.register-link {
  color: #007bff;
  text-decoration: none;
  font-weight: 600;
}
</style>
@/http-client
