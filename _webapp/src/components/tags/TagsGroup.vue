<template>
  <div v-if="$props.tags && $props.tags.length > 0" class="tagGroup">
    <tag v-for="tag in $props.tags" :tag="tag" @on-tag-selected="onTagSelected" />
  </div>
</template>
<script setup lang="ts">
import { ref, type PropType, computed } from 'vue'
import Tag from '@/components/tags/Tag.vue'
import type { TagData } from './TagData'
const props = defineProps({
  tags: Array as PropType<TagData[]>
})

const emit = defineEmits<{
  (e: 'onTagSelected', tag: TagData): void
}>()

const onTagSelected = (tag: TagData) => {
  emit('onTagSelected', tag)
}
</script>
<style scoped>
.tagGroup {
  padding-top: 8px;
  padding-left: 8px;
  padding-right: 8px;
  pointer-events: initial;
  position: relative;
  width: 100%;
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
  flex-direction: row;
  flex: 1;
  max-height: 10%;
}
</style>
