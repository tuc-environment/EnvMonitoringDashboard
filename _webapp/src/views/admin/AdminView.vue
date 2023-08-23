<template>
  <SidebarLayout :sections="sections" default-sub-view="Account" />
</template>

<script lang="ts">
import SidebarLayout from '@/layouts/SidebarLayout.vue'

import AccountView from '@/views/admin/AccountView.vue'
import StationView from '@/views/admin/management/StationView.vue'
import SensorView from '@/views/admin/management/SensorView.vue'
import RecordView from '@/views/admin/management/RecordView.vue'

export default {
  components: {
    SidebarLayout
  },
  data() {
    return {
      sections: [
        {
          name: 'Account',
          subViews: [
            {
              name: 'Account',
              icon: 'bi-person',
              component: AccountView
            }
          ]
        },
        {
          name: 'Management',
          subViews: [
            {
              name: 'Stations',
              icon: 'bi-building',
              component: StationView
            },
            {
              name: 'Sensors',
              icon: 'bi-camera-video',
              component: SensorView,
              hidden: true
            },
            {
              name: 'Records',
              icon: 'bi-activity',
              component: RecordView,
              hidden: true
            }
          ]
        }
      ]
    }
  }
}
</script>
