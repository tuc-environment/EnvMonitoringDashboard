<template>
  <div class="container">
    <div class="mx-auto my-5 login py-3 px-5">
      <div>
        <img class="logo" src="/logo.png" />
      </div>

      <h2 class="my-3 text-center">Sign In to TUC Environment</h2>

      <div class="my-3 mb-5 text-center">
        New Here? <RouterLink to="/register" class="register-link">Create an Account</RouterLink>
      </div>

      <form @submit.prevent="preventDefault">
        <div class="form-group my-3">
          <label class="mb-2">Username</label>
          <input type="text" class="form-control" v-model="username" :disabled="requesting" />
        </div>

        <div class="form-group my-3">
          <label class="mb-2">Password</label>
          <input type="password" class="form-control" v-model="password" :disabled="requesting" />
        </div>

        <button
          type="submit"
          class="mt-3 w-100 btn btn-lg btn-primary"
          :disabled="requesting"
          @click="login(username, password)"
        >
          <div v-if="requesting" class="spinner-border text-light" role="status" />
          <div v-else>Login</div>
        </button>

        <div
          class="mt-5 d-flex justify-content-between mb-2"
          style="color: rgba(0, 0, 0, 0.5); font-size: 0.99rem"
        >
          <div>Privacy Policy</div>
          <div>Term of Service</div>
          <div>Contact</div>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink, useRouter } from 'vue-router'
import httpclient from '@/httpclient'
import { ref } from 'vue'

const username = ref('')
const password = ref('')
const requesting = ref(false)
const router = useRouter()

const login = async (username: string, password: string) => {
  requesting.value = true
  const resp = await httpclient.login(username, password)
  if (resp?.code == 200) {
    router.push('/admin')
  } else if (resp?.code == 400) {
    clearForm()
    alert(resp?.error)
  }
  requesting.value = false
}

const onSubmit = () => {
  if (username.value.length == 0 || password.value.length == 0) {
    alert('Username & password cannot be empty')
  } else if (password.value.length < 6) {
    alert('Password is least 6 digits')
  } else {
    login(username.value, password.value)
  }
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
  background-color: rgba(86, 161, 208, 0.15);
  border: 1px solid #ccc;
  border-radius: 4px;
}

.logo {
  width: 72px;
  margin: 20px auto;
  display: block;
  border-radius: 5px;
}
</style>
