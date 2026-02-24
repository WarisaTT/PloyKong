<template>
  <div class="stat-card" :class="`color-${color}`">
    <div class="stat-label">{{ label }}</div>
    <div class="stat-value">{{ value.toLocaleString() }}</div>
    <div class="stat-change">{{ change }}</div>
  </div>
</template>

<script setup lang="ts">
defineProps<{ label: string; value: number; change: string; color: string }>()
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

.color-indigo .stat-value { color: #818cf8; }
.color-indigo .stat-change { color: #818cf8; }
.color-cyan .stat-value { color: var(--neon-cyan); }
.color-cyan .stat-change { color: var(--neon-cyan); }
.color-pink .stat-value { color: #f472b6; }
.color-pink .stat-change { color: #f472b6; }
.color-green .stat-value { color: var(--success); }
.color-green .stat-change { color: var(--success); }
</style>
