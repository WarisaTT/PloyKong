<template>
  <div class="exp-preview">
    <div v-if="!data.items || !data.items.length" class="empty-msg">Not have experiences</div>
    <div v-else class="exp-list">
      <div v-for="(item, i) in data.items.slice(0,3)" :key="i" class="exp-item">
        <div class="exp-dot"></div>
        <div>
          <div class="exp-pos">{{ item.position }} <span class="exp-at">@ {{ item.company }}</span></div>
          <div class="exp-date">{{ item.start_date }} – {{ item.is_current ? 'Present' : item.end_date }}</div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
defineProps<{ data: any }>()
</script>
<style scoped>
.empty-msg { font-size: 12px; color: var(--muted); font-style: italic; }
.exp-list { display: flex; flex-direction: column; gap: 10px; }
.exp-item { display: flex; gap: 12px; align-items: flex-start; }
.exp-dot { width: 8px; height: 8px; border-radius: 50%; background: var(--indigo); flex-shrink: 0; margin-top: 5px; }
.exp-pos { font-size: 13px; font-weight: 700; }
.exp-at { color: var(--neon-purple); font-weight: 400; font-size: 12px; }
.exp-date { font-size: 11px; color: var(--muted); margin-top: 2px; }
</style>
