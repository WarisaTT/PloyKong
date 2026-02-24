<template>
  <div class="stat-card" :class="`color-${color}`">
    <div class="stat-label">
      <component v-if="icon" :is="icon" :size="14" class="icon-inline" />
      {{ label }}
    </div>
    <div class="stat-value">{{ value.toLocaleString() }}</div>
    <div class="stat-change">{{ change }}</div>
  </div>
</template>

<script setup lang="ts">
defineProps<{ label: string; value: number; change: string; color: string; icon?: any }>()
</script>

<style scoped>
.stat-card {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 16px;
  padding: 22px;
  transition: all 0.2s;
}
.stat-card:hover { transform: translateY(-2px); }
.stat-label { font-size: 13px; color: var(--muted); font-weight: 600; margin-bottom: 10px; }
.stat-value { font-family: var(--font-display); font-size: 32px; font-weight: 800; margin-bottom: 6px; }
.stat-change { font-size: 12px; font-weight: 600; }

.color-indigo .stat-value { color: var(--indigo); }
.color-indigo .stat-change { color: var(--indigo); }
.color-cyan .stat-value { color: var(--cyan); }
.color-cyan .stat-change { color: var(--cyan); }
.color-pink .stat-value { color: var(--pink); }
.color-pink .stat-change { color: var(--pink); }
.color-green .stat-value { color: var(--success); }
.color-green .stat-change { color: var(--success); }

.icon-inline { vertical-align: text-bottom; margin-right: 6px; }
</style>
