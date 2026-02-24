<template>
  <section class="pub-exp pub-section">
    <h2 class="pub-section-title">💼 Experience</h2>
    <div class="timeline">
      <div v-for="(item, i) in data.items" :key="i" class="timeline-item">
        <div class="timeline-dot"></div>
        <div class="timeline-line" v-if="i < data.items.length - 1"></div>
        <div class="timeline-content">
          <div class="tl-header">
            <div class="tl-position">{{ item.position }}</div>
            <div class="tl-date">{{ item.start_date }} – {{ item.is_current ? 'Present' : item.end_date }}</div>
          </div>
          <div class="tl-company">{{ item.company }} <span v-if="item.location" class="tl-loc">· {{ item.location }}</span></div>
          <div class="tl-desc">{{ item.description }}</div>
        </div>
      </div>
    </div>
  </section>
</template>
<script setup lang="ts">
defineProps<{ data: any }>()
</script>
<style scoped>
.pub-section { padding: 48px 0; border-top: 1px solid rgba(255,255,255,0.06); }
.pub-section-title { font-size: 22px; font-weight: 800; margin-bottom: 28px; }
.timeline { display: flex; flex-direction: column; gap: 0; padding-left: 20px; }
.timeline-item { position: relative; padding-left: 28px; padding-bottom: 32px; }
.timeline-dot { position: absolute; left: -8px; top: 4px; width: 14px; height: 14px; border-radius: 50%; background: linear-gradient(135deg, var(--indigo, #4F46E5), #a855f7); flex-shrink: 0; }
.timeline-line { position: absolute; left: -2px; top: 18px; bottom: 0; width: 2px; background: rgba(255,255,255,0.08); }
.tl-header { display: flex; justify-content: space-between; align-items: flex-start; gap: 12px; margin-bottom: 4px; flex-wrap: wrap; }
.tl-position { font-size: 15px; font-weight: 700; }
.tl-date { font-size: 12px; color: #818cf8; font-weight: 600; }
.tl-company { font-size: 13px; color: var(--muted, #94a3b8); margin-bottom: 8px; }
.tl-loc { color: rgba(148,163,184,0.6); }
.tl-desc { font-size: 13px; color: var(--muted, #94a3b8); line-height: 1.7; }
</style>
