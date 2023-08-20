<template>
  <div class="main-content">
    <SectionLayout title="Overview">
      <LabelInputComponent label="Username" type="text" :field="username" disabled />
      <LabelInputComponent
        label="Access Token"
        type="text"
        :field="token"
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
      />
      <LabelInputComponent label="Confirm Password" type="password" />
      <button class="btn btn-primary">Change Password</button>
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
      token: ''
    }
  },
  methods: {
    async regenrateToken() {
      const resp = await httpclient.regenrateToken()
      this.username = resp?.payload.username || ''
      this.token = resp?.payload.token || ''
    },
    async changePassword() {}
  }
}
</script>
