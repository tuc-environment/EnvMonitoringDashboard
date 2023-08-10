<template>
  <SidebarLayout name="Account" :sub-views="subViews" default-sub-view="Account" />
</template>

<script lang="ts">
import SidebarLayout from '@/layouts/SidebarLayout.vue'

import AccountView from '@/views/admin/AccountView.vue'

export default {
  components: {
    SidebarLayout
  },
  data() {
    return {
      subViews: [
        {
          name: 'Account',
          icon: 'bi-person',
          component: AccountView
        }
      ]
    }
  }
}
</script>
