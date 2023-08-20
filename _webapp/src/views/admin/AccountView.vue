<template>
  <div class="main-content">
    <SectionLayout title="Overview">
      <LabelInputComponent label="Username" type="text" v-model:field="username" disabled />
      <LabelInputComponent
        label="Access Token"
        type="text"
        v-model:field="token"
        footnote="Please keep access token secretly."
        disabled
      />
      <button class="btn btn-primary" @click="regenrateToken">Regenerate Token</button>
    </SectionLayout>

    <SectionLayout title="Change Password">
      <LabelInputComponent
        label="Password"
        type="password"
        footnote="Please make sure that your password is at least 8 characters long, includes a number, uppercase and lowercase letters."
        v-model:field="password"
      />
      <LabelInputComponent
        label="Confirm Password"
        type="password"
        v-model:field="confirmPassword"
      />
      <div v-if="password" class="small mt-2">
        <span v-if="password != confirmPassword" class="text-danger">Password not match</span>
        <span v-else class="text-success">Password match</span>
      </div>
      <button
        class="btn btn-primary mt-2"
        @click="changePassword"
        :disabled="password == '' || password != confirmPassword"
      >
        Change Password
      </button>
    </SectionLayout>
  </div>
</template>

<script lang="ts">
import SectionLayout from '@/layouts/SectionLayout.vue'
import LabelInputComponent from '@/components/LabelInputComponent.vue'
import httpclient from '@/httpclient'

export default {
  components: {
    SectionLayout,
    LabelInputComponent
  },
  async created() {
    const resp = await httpclient.getAccount()
    this.username = resp?.payload.username || ''
    this.token = resp?.payload.token || ''
  },
  data() {
    return {
      username: '',
      token: '',
      password: '',
      confirmPassword: ''
    }
  },
  methods: {
    async regenrateToken() {
      const resp = await httpclient.regenrateToken()
      this.username = resp?.payload.username || ''
      this.token = resp?.payload.token || ''
    },
    async changePassword() {
      const resp = await httpclient.changePassword(this.password)
      alert('Password updated')
    }
  }
}
</script>
