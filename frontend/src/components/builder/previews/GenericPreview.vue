<template>
  <div class="generic-preview">
    <span class="generic-type">{{ section.type }}</span>
    <span class="generic-hint">คลิกเพื่อ edit ใน Properties panel →</span>
  </div>
</template>
<script setup lang="ts">
import type { Section } from '@/types'
defineProps<{ section: Section }>()
</script>
<style scoped>
.generic-preview { display: flex; align-items: center; gap: 10px; }
.generic-type { font-size: 12px; font-weight: 700; text-transform: capitalize; background: rgba(79,70,229,0.15); padding: 3px 10px; border-radius: 100px; color: #818cf8; }
.generic-hint { font-size: 11px; color: var(--muted); font-style: italic; }
</style>
