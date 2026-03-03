<template>
  <textarea
    ref="taRef"
    :value="modelValue"
    :placeholder="placeholder"
    :class="['form-input', 'auto-textarea', className]"
    :style="style"
    rows="1"
    @input="onInput"
    @keydown.enter.exact.prevent="insertNewline"
  />
</template>

<script setup lang="ts">
import { ref, watch, nextTick, onMounted } from 'vue'

const props = defineProps<{
  modelValue: string
  placeholder?: string
  className?: string
  style?: any
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', val: string): void
  (e: 'input'): void
}>()

const taRef = ref<HTMLTextAreaElement | null>(null)

function resize() {
  const el = taRef.value
  if (!el) return
  el.style.height = 'auto'
  el.style.height = el.scrollHeight + 'px'
}

function onInput(e: Event) {
  const val = (e.target as HTMLTextAreaElement).value
  emit('update:modelValue', val)
  emit('input')
  nextTick(resize)
}

function insertNewline(e: Event) {
  const el = taRef.value
  if (!el) return
  const start = el.selectionStart
  const end = el.selectionEnd
  const val = el.value
  const newVal = val.substring(0, start) + '\n' + val.substring(end)
  emit('update:modelValue', newVal)
  nextTick(() => {
    el.selectionStart = el.selectionEnd = start + 1
    resize()
  })
}

onMounted(() => nextTick(resize))

watch(() => props.modelValue, () => nextTick(resize))
</script>

<style scoped>
.auto-textarea {
  resize: none;
  overflow: hidden;
  min-height: 38px;
  line-height: 1.6;
  white-space: pre-wrap;
  word-break: break-word;
  font-family: inherit;
  transition: height 0.05s ease;
}
</style>
