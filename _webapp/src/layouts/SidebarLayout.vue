<template>
  <div class="view">
    <NavBarComponent class="navbar" />

    <div class="content d-flex flex-md-row flex-column-reverse">
      <div class="sidebar p-3">
        <div
          v-for="(section, sectionIdx) in sections.filter((section) => !section.hidden)"
          :key="sectionIdx"
        >
          <h5 class="mb-4">{{ section.name }}</h5>

          <SideBarButtonComponent
            v-for="(subView, idx) in section.subViews.filter((v) => !v.hidden)"
            :key="idx"
            :name="subView.name"
            :icon="subView.icon"
            :selected="subView.name == subViewName"
            @click="clickSideBar(subView.name)"
          />
        </div>

        <hr style="color: rgba(0, 0, 0, 0.2)" />

        <h5 class="mb-4">链接</h5>

        <SideBarButtonComponent
          name="数据看板"
          icon="bi-card-heading"
          right-icon="bi-box-arrow-up-right"
          @click="open('/')"
        />

        <SideBarButtonComponent
          name="天津商业大学"
          icon="bi-info-circle"
          right-icon="bi-box-arrow-up-right"
          @click="open('https://www.tjcu.edu.cn/')"
        />

        <SideBarButtonComponent
          name="中国科学院电工研究所"
          icon="bi-info-circle"
          right-icon="bi-box-arrow-up-right"
          @click="open('http://www.iee.cas.cn/')"
        />

        <SideBarButtonComponent
          name="天津大学"
          icon="bi-info-circle"
          right-icon="bi-box-arrow-up-right"
          @click="open('https://www.tju.edu.cn/')"
        />

        <button
          v-if="backButtonPath"
          type="button"
          class="mt-3 w-100 btn btn-outline-primary"
          @click="back()"
        >
          {{ backButtonName }}
        </button>
        <button type="button" class="mt-3 w-100 btn btn-danger" @click="logout()">登出</button>
      </div>

      <div class="main flex-grow-1 p-3">
        <h1>{{ selectedSubView?.name }}</h1>
        <component :is="selectedComponent" />
        <div style="height: 120px"></div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import NavBarComponent from '@/components/NavBarComponent.vue'
import SideBarButtonComponent from '@/components/SideBarButtonComponent.vue'
import httpclient from '@/http-client'
import { toRaw, type Component } from 'vue'

export interface SubView {
  name: string
  icon: string
  hidden?: boolean
  component: Component
}

export interface SectionView {
  name: string
  hidden?: boolean
  subViews: Array<SubView>
}

export default {
  components: {
    NavBarComponent,
    SideBarButtonComponent
  },
  props: {
    sections: { type: Array<SectionView>, required: true },
    defaultSubView: { type: String, required: true },
    backButtonName: { type: String },
    backButtonPath: { type: String }
  },
  created() {
    if (!httpclient.token) {
      this.logout()
    }
    if (!this.selectedSubView) {
      this.$router.push({ query: { view: this.defaultSubView } })
    }
  },
  computed: {
    subViewName(): string {
      const tab = this.$route.query.view || this.defaultSubView
      return tab.toString()
    },
    selectedSubView(): SubView | null {
      const views = this.sections
        .flatMap((section) => section.subViews)
        .filter((v) => v.name == this.subViewName)
      return views[0]
    },
    selectedComponent(): Component | null {
      const component = this.selectedSubView?.component
      return component ? toRaw(component) : null
    }
  },
  methods: {
    clickSideBar(view: string) {
      this.$router.push({ query: { view: view } })
    },
    open(url: string) {
      window.open(url)
    },
    back() {
      this.$router.push(this.backButtonPath!)
    },
    async logout() {
      httpclient.logout()
      this.$router.push('/login')
    }
  }
}
</script>

<style scoped lang="scss">
.navbar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
}

.content {
  margin-top: 59px;
}

.sidebar {
  width: 100%;
  border-right: 1px solid #e0e0e0;
  background-color: $tux-sidebar-bg;
  min-width: 300px;
}

.main {
  min-height: 50vh;
}

@media (min-width: 768px) {
  .view {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
  }

  .content {
    height: calc(100vh - 59px);
  }

  .sidebar {
    overflow: auto;
    width: 300px;
  }

  .main {
    overflow: auto;
  }
}
</style>
@/http-client
