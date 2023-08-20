<template>
  <div class="container">
    <div class="mx-auto my-5 login py-3 px-5">
      <div>
        <img class="logo" src="/logo.png" />
      </div>

      <h2 class="my-3 text-center">Register</h2>

      <div class="my-3 mb-5 text-center">
        Already have an account? <RouterLink to="/login" class="register-link">Login</RouterLink>
      </div>

      <form>
        <div class="form-group my-3">
          <label class="mb-2">Username</label>
          <input type="text" class="form-control" v-model="username" :disabled="requesting" />
        </div>

        <div class="form-group my-3">
          <label class="mb-2">Password</label>
          <input type="password" class="form-control" v-model="password" :disabled="requesting" />
        </div>

        <div class="form-group my-3">
          <label class="mb-2">Confirm password</label>
          <input type="password" class="form-control" v-model="passwordConfirm" />
          <!-- check password match-->
          <div v-show="showPasswordMatch" class="small mt-2">
            <span v-if="passwordMatch" class="text-success">Password match</span>
            <span v-else class="text-danger">Password not match</span>
          </div>
        </div>

        <button
          type="submit"
          class="mt-3 w-100 btn btn-lg btn-primary"
          :disabled="requesting || !passwordMatch"
          @click="onSubmit"
        >
          <!-- Register -->
          <div v-if="requesting" class="spinner-border text-light mx-auto d-block" role="status" />
          <div v-else>Register</div>
        </button>

        <div
          class="mt-5 d-flex justify-content-between mb-2"
          style="color: rgba(0, 0, 0, 0.5); font-size: 0.99rem"
        >
          <div style="cursor: pointer">Privacy Policy</div>
          <div style="cursor: pointer">Term of Service</div>
          <div style="cursor: pointer">Contact</div>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import httpclient from '@/httpclient'

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
  } else if (password.value !== passwordConfirm.value) {
    alert('Please confirm password')
  } else {
    register(username.value, password.value)
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
  background-color: rgba(86, 161, 208, 0.15);
  border: 1px solid #ccc;
  border-radius: 4px;
}

.logo {
  width: 72px;
  margin: 20px auto;
  display: block;
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
