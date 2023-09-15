<template>
  <SidebarLayout :sections="sections" default-sub-view="Account" />
</template>

<script lang="ts">
import SidebarLayout from '@/layouts/SidebarLayout.vue'

import AccountView from '@/views/admin/AccountView.vue'
import StationView from '@/views/admin/management/StationView.vue'
import SensorView from '@/views/admin/management/SensorView.vue'

export default {
  components: {
    SidebarLayout
  },
  data() {
    return {
      sections: [
        {
          name: '管理',
          subViews: [
            {
              name: '账户管理',
              icon: 'bi-person',
              component: AccountView
            },
            {
              name: '站点管理',
              icon: 'bi-building',
              component: StationView
            },
            {
              name: '传感器管理',
              icon: 'bi-camera-video',
              component: SensorView,
              hidden: true
            }
          ]
        }
        // {
        //   name: '管理',
        //   subViews: [

        //   ]
        // }
      ]
    }
  }
}
</script>
