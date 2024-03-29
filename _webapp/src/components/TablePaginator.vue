<template>
  <div
    v-if="
      $props.offset != undefined &&
      $props.limit != undefined &&
      $props.total != undefined &&
      $props.total > $props.limit
    "
  >
    <div class="d-flex justify-content-end align-items-center">
      <div class="my-2 mx-2">
        第{{ currentPageIdx + 1 }}页, 总数: {{ $props.total }}, 当前显示 {{ $props.offset + 1 }} -
        {{ +$props.offset + +$props.limit }}
      </div>
      <nav>
        <ul class="pagination pagination-sm justify-content-end my-2">
          <li
            class="page-item"
            :class="{ disabled: currentPageIdx == 0 }"
            style="cursor: pointer"
            @click="clickPrevious"
          >
            <a class="page-link" tabindex="-1">向前</a>
          </li>
          <li
            v-if="currentPageIdx >= Math.ceil(displayedPageCount / 2)"
            style="cursor: pointer"
            class="page-item"
            :key="0"
            @click="clickPageIndex(0)"
          >
            <a class="page-link">{{ 1 }}</a>
          </li>
          <li v-if="currentPageIdx > Math.ceil(displayedPageCount / 2)" class="page-item">...</li>
          <li
            class="page-item"
            v-for="label in pageLabels"
            :key="label"
            :class="{ disabled: label - 1 == currentPageIdx }"
            style="cursor: pointer"
            @click="clickPageIndex(label - 1)"
          >
            <a class="page-link">{{ label }}</a>
          </li>
          <li
            v-if="currentPageIdx < pagesCount - Math.ceil(displayedPageCount / 2) - 1"
            class="page-item"
          >
            ...
          </li>
          <li
            v-if="currentPageIdx < pagesCount - Math.ceil(displayedPageCount / 2)"
            class="page-item"
            style="cursor: pointer"
            :key="0"
            @click="clickPageIndex(pagesCount - 1)"
          >
            <a class="page-link">{{ pagesCount }}</a>
          </li>
          <li
            class="page-item"
            :class="{ disabled: currentPageIdx == pagesCount }"
            style="cursor: pointer"
            @click="clickNext"
          >
            <a class="page-link" tabindex="-1">向后</a>
          </li>
        </ul>
      </nav>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps({
  offset: Number,
  limit: Number,
  total: Number
})

const emit = defineEmits<{
  (e: 'toNext', offset: number): void
  (e: 'toPrevious', offset: number): void
  (e: 'toIndex', offset: number): void
}>()

const currentPageIdx = computed(() =>
  props.offset != undefined && props.limit != undefined ? Math.ceil(props.offset / props.limit) : 0
)
const pagesCount = computed(() =>
  props.total != undefined && props.limit != undefined ? Math.ceil(props.total / props.limit) : 0
)
const displayedPageCount = 5
const pageLabels = computed((): Array<number> => {
  var labels = []
  const base = currentPageIdx.value + 1 - Math.ceil(displayedPageCount / 2)
  if (base <= 0) {
    for (let i = 0; i < displayedPageCount; i++) {
      const val = i + 1
      if (val > 0 && val <= pagesCount.value) {
        labels.push(val)
      }
    }
  } else if (base >= pagesCount.value - displayedPageCount) {
    for (let i = 0; i < displayedPageCount; i++) {
      const val = pagesCount.value - displayedPageCount + i + 1
      if (val > 0 && val <= pagesCount.value) {
        labels.push(val)
      }
    }
  } else {
    for (let i = 1; i <= displayedPageCount; i++) {
      const val = base + i
      if (val > 0 && val <= pagesCount.value) {
        labels.push(val)
      }
    }
  }
  return labels
})

const clickPrevious = async () => {
  if (props.offset != undefined && props.limit != undefined && props.offset > 0) {
    const offset = props.offset - props.limit
    emit('toPrevious', offset)
  }
}

const clickNext = () => {
  if (
    props.offset != undefined &&
    props.limit != undefined &&
    props.total != undefined &&
    props.offset + props.limit < props.total
  ) {
    const offset = props.offset + props.limit
    emit('toNext', offset)
  }
}

const clickPageIndex = (pageIndex: number) => {
  if (props.limit != undefined) {
    const offset = props.limit * pageIndex
    emit('toIndex', offset)
  }
}
</script>

<style scope>
tr,
th {
  vertical-align: middle;
}
</style>
