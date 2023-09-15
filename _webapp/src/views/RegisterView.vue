<template>
  <BackgroundComponent />

  <DashboardLayout>
    <div class="container">
      <div class="mx-auto my-5 login py-3 px-5">
        <div>
          <img class="logo" src="/logo.png" />
        </div>

        <h2 class="my-3 text-center">注册账号</h2>

        <div class="my-3 mb-5 text-center">
          已经拥有账号？ <RouterLink to="/login" class="register-link">登入</RouterLink>
        </div>

        <div v-if="errorMsg" class="alert alert-danger" role="alert">{{ errorMsg }}</div>

        <form>
          <div class="form-group my-3">
            <label class="mb-2">用户名</label>
            <input type="text" class="form-control" v-model="username" :disabled="requesting" />
          </div>

          <div class="form-group my-3">
            <label class="mb-2">密码</label>
            <input type="password" class="form-control" v-model="password" :disabled="requesting" />
          </div>

          <div class="form-group my-3">
            <label class="mb-2">确认密码</label>
            <input type="password" class="form-control" v-model="passwordConfirm" />
            <!-- check password match-->
            <div v-show="showPasswordMatch" class="small mt-2">
              <span v-if="passwordMatch" class="text-success">密码匹配</span>
              <span v-else class="text-danger">密码不匹配</span>
            </div>
          </div>

          <button
            type="submit"
            class="mt-3 w-100 btn btn-lg btn-primary"
            :disabled="requesting || !passwordMatch"
            @click="onSubmit"
          >
            <!-- Register -->
            <div
              v-if="requesting"
              class="spinner-border text-light mx-auto d-block"
              role="status"
            />
            <div v-else>注册</div>
          </button>

          <div
            class="mt-5 d-flex justify-content-center mb-2"
            style="color: rgba(0, 0, 0, 0.5); font-size: 0.99rem"
          >
            <a
              style="cursor: pointer; text-decoration: none"
              href="https://www.tjcu.edu.cn/"
              target="_blank"
              >天津商业大学 <i class="bi bi-box-arrow-up-right"></i
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
import { computed, ref } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import httpclient from '@/httpclient'
import DashboardLayout from '@/layouts/DashboardLayout.vue'
import BackgroundComponent from '@/components/BackgroundComponent.vue'

let errorMsg = ref('')
let username = ref('')
let password = ref('')
let passwordConfirm = ref('')
let requesting = ref(false)

const router = useRouter()

const showPasswordMatch = computed((): boolean => {
  return password.value.length > 0 && passwordConfirm.value.length > 0
})

const passwordMatch = computed((): boolean => {
  return password.value == passwordConfirm.value
})

const register = async (username: string, password: string) => {
  requesting.value = true
  const resp = await httpclient.register(username, password)
  if (resp?.code == 200) {
    router.push('/login')
  } else if (resp?.code == 400) {
    errorMsg.value = '用户名已经被注册'
  }
  requesting.value = false
}

const onSubmit = async (evt: Event) => {
  evt.preventDefault()
  if (username.value.length == 0 || password.value.length == 0) {
    errorMsg.value = '用户名和密码不能为空'
    return
  } else if (password.value.length < 6) {
    errorMsg.value = '密码至少 6 位'
    return
  } else if (password.value !== passwordConfirm.value) {
    errorMsg.value = '密码不匹配'
    return
  } else {
    await register(username.value, password.value)
  }
}

const clearForm = () => {
  username.value = ''
  password.value = ''
  passwordConfirm.value = ''
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
  width: 72px;
  margin: 20px auto;
  display: block;
  border-radius: 5px;
}

.register-link {
  color: #007bff;
  text-decoration: none;
  font-weight: 600;
}

input {
  background-color: rgb(243, 247, 251);
}
</style>
