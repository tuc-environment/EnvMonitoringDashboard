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
        Count: {{ $props.total }}, Showing {{ $props.offset + 1 }} -
        {{ $props.offset + $props.limit }}
      </div>
      <nav>
        <ul class="pagination pagination-sm justify-content-end my-2">
          <li class="page-item" :class="{ disabled: currentPageIdx == 0 }" @click="clickPrevious">
            <a class="page-link" tabindex="-1">Previous</a>
          </li>
          <li
            class="page-item"
            v-for="(label, idx) in pageLabels"
            :key="idx"
            :class="{ disabled: idx == currentPageIdx }"
            @click="clickPageIndex(idx)"
          >
            <a class="page-link">{{ label }}</a>
          </li>
          <li
            class="page-item"
            :class="{ disabled: currentPageIdx == pagesCount }"
            @click="clickNext"
          >
            <a class="page-link" tabindex="-1">Next</a>
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
const pageLabels = computed((): Array<number> => {
  var labels = []
  for (let i = 1; i <= pagesCount.value; i++) {
    labels.push(i)
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
