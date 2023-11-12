<template>
  <div style="text-align: left">
    <div class="d-flex my-2 small node p-2 rounded">
      <div :style="{ width: `${depth * 20}px` }"></div>
      <div v-if="hasChildren" class="flex-grow-1" style="display: flex; flex-direction: row">
        <input
          class="form-check-input me-2"
          type="checkbox"
          :checked="selected"
          @input="onChangeSelected"
        />
        <div @click="clickNode">
          <i v-if="!showChildren" class="bi bi-caret-right-fill me-2"></i>
          <i v-else class="bi bi-caret-down-fill me-2"></i>
          {{ $props.node?.label }}
        </div>
      </div>

      <div v-else>
        <input
          class="form-check-input me-2"
          type="checkbox"
          :checked="selected"
          @input="onChangeSelected"
        />
        <label class="form-check-label"> {{ $props.node?.label }} </label>
      </div>
    </div>
    <div v-if="hasChildren && showChildren">
      <UnlinkedSensorNode
        v-for="(child, idx) in $props.node?.children"
        :key="idx"
        :node="child"
        :depth="depth + 1"
        @select="onSelected"
        @deselect="onDeselected"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
// import { useLinkSensorStore } from '@/stores/LinkSensorsStore'
import { computed, onMounted, ref, watch, type PropType } from 'vue'
import { type CommonNode } from './CommonNode'
import UnlinkedSensorNode from './UnlinkedSensorNode.vue'
// const linkSensorStore = useLinkSensorStore()
const selected = ref<boolean>(false)

// const hasChildrenSelected = computed<boolean>(() => hasSelectedChildren(props.node))

const props = defineProps({
  node: Object as PropType<CommonNode>,
  depth: {
    type: Number,
    required: true
  }
})
onMounted(() => {
  selected.value = props.node?.selected ?? false
})

const emit = defineEmits<{
  (e: 'select', ids: number[]): void
  (e: 'deselect', ids: number[]): void
}>()

watch(props, async (newVal) => {
  selected.value = newVal.node?.selected ?? false
})

const showChildren = ref(false)
const hasChildren = computed((): boolean => {
  return props.node?.children != undefined && props.node.children.length > 0
})

const clickNode = () => {
  if (hasChildren.value) {
    showChildren.value = !showChildren.value
  }
}

const onSelected = (ids: number[]) => {
  emit('select', ids)
}

const onDeselected = (ids: number[]) => {
  emit('deselect', ids)
}

const onChangeSelected = (e: any) => {
  const node = props.node
  if (node) {
    const checked: boolean = e.target.checked
    if (hasChildren.value) {
      const ids = node.children.map((n) => n.id).filter((id) => id) as number[]
      if (checked) {
        emit('select', ids)
      } else {
        emit('deselect', ids)
      }
    } else if (node.id) {
      if (checked) {
        emit('select', [node.id])
      } else {
        emit('deselect', [node.id])
      }
    }
  }
}
</script>

<style>
.node {
  cursor: pointer;
}

.node:hover {
  background-color: #c3c1c126;
}
</style>
