<template>
  <div class="tagWrapper" @click="onTagSelected">
    <i class="bi bi-x-lg"></i>
    {{ tag.title }}
  </div>
</template>
<script setup lang="ts">
import { type PropType, computed } from 'vue'
import { type TagData } from '@/components/tags/TagData'

const props = defineProps({
  tag: {
    type: Object as PropType<TagData>,
    required: true
  }
})
const emit = defineEmits<{
  (e: 'onTagSelected', tag: TagData): void
}>()

const onTagSelected = () => {
  emit('onTagSelected', props.tag)
}
</script>
<style scoped>
.tagWrapper {
  height: 20;
  border-radius: 20px;
  border: 1px solid lightgrey;
  padding-left: 10px;
  padding-right: 10px;
  padding-top: 8px;
  padding-bottom: 8px;
  font-size: 10px;
  color: lightgrey;
}
</style>
