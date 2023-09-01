<template>
  <div class="mb-2">
    <label class="form-label small">{{ label }}</label
    ><span v-show="required" class="ms-1 text-danger">*</span>
    <div class="input-group">
      <input
        :type="type"
        class="form-control"
        :value="field"
        :placeholder="placeholder"
        @input="onInput"
        :disabled="disabled"
      />
      <span v-if="unit" class="input-group-text">{{ unit }}</span>
    </div>

    <label class="text-secondary small" v-if="footnote">{{ footnote }}</label>
  </div>
</template>

<script lang="ts">
export default {
  props: {
    label: { type: String, required: true },
    type: { type: String, required: true },
    field: { type: String },
    placeholder: { type: String },
    unit: { type: String },
    footnote: { type: String },
    required: { type: Boolean, default: false },
    disabled: { type: Boolean, default: false }
  },
  emits: ['update:field'],
  methods: {
    onInput(e: any) {
      this.$emit('update:field', e.target.value)
    }
  }
}
</script>
