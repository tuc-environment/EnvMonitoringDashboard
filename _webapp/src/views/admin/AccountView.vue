<template>
  <div class="main-content">
    <SectionLayout title="总览">
      <LabelInputComponent label="用户名" type="text" v-model:field="username" disabled />
      <LabelInputComponent
        label="访问令牌"
        type="text"
        v-model:field="token"
        footnote="请妥善保管访问令牌避免泄漏。"
        disabled
      />
      <button class="btn btn-primary" @click="regenrateToken">重新生成令牌</button>
    </SectionLayout>

    <SectionLayout title="修改密码">
      <LabelInputComponent
        label="密码"
        type="password"
        footnote="请确保您的密码至少包含 8 个字符，且包括数字、大写字母和小写字母。"
        v-model:field="password"
      />
      <LabelInputComponent label="确认密码" type="password" v-model:field="confirmPassword" />
      <div v-if="password" class="small mt-2">
        <span v-if="password != confirmPassword" class="text-danger">Password not match</span>
        <span v-else class="text-success">Password match</span>
      </div>
      <button
        class="btn btn-primary mt-2"
        @click="changePassword"
        :disabled="password == '' || password != confirmPassword"
      >
        修改密码
      </button>
    </SectionLayout>
  </div>
</template>

<script lang="ts">
import LabelInputComponent from '@/components/LabelInputComponent.vue'
import httpclient from '@/http-client'
import SectionLayout from '@/layouts/SectionLayout.vue'

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
      await httpclient.changePassword(this.password)
      alert('密码已更新')
    }
  }
}
</script>
