<template>
  <section class="pub-stats pub-section" :class="[
    'layout-' + (data.layout || 'centered'),
    { 'hide-title': data.hide_title, 'hide-divider': data.hide_divider }
  ]" :style="data.section_bg_color ? { background: data.section_bg_color } : {}">
    <div class="section-header-wrapper">
      <h2 class="layered-title" :data-text="data.title || 'Stats'">{{ data.title || 'Stats & Numbers' }}</h2>
    </div>
    
    <div class="pub-section-content">
      <div class="stats-grid">
        <div 
          v-for="(stat, i) in data.items" 
          :key="i"
          class="stat-box animate-scale-up"
          v-intersect
          :style="{ transitionDelay: `${Number(i) * 100}ms` }"
        >
          <div class="stat-value">{{ stat.value }}</div>
          <div class="stat-label">{{ stat.label }}</div>
        </div>
      </div>

      <div v-if="!data.items || data.items.length === 0" style="text-align: center; padding: 20px;">
        <span class="generic-hint">คลิกเพื่อ edit ใน Properties panel &rarr;</span>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
defineProps<{ data: any }>();
</script>

<style scoped>
.pub-section.hide-divider { border-top: none !important; }
.pub-section.hide-title .layered-title { display: none !important; }
.pub-section {
  padding: 40px 0;
  border-top: 1px solid var(--section-border);
}

:global(.theme-dark) {
  --section-border: rgba(255, 255, 255, 0.06);
}
:global(.theme-light) {
  --section-border: rgba(0, 0, 0, 0.06);
}

.layout-left .section-header-wrapper {
  text-align: left;
}
.layout-left .stats-grid {
  justify-content: flex-start;
}

.stats-grid {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 24px;
}

.stat-box {
  background: var(--surface);
  border: 1px solid var(--border);
  padding: 24px 32px;
  border-radius: 16px;
  min-width: 140px;
  text-align: center;
  transition: transform 0.3s;
}

.stat-box:hover {
  transform: translateY(-4px);
  border-color: var(--primary);
  box-shadow: 0 10px 30px rgba(0,0,0,0.1);
}

.stat-value {
  font-size: 36px;
  font-weight: 800;
  color: var(--primary);
  line-height: 1;
  margin-bottom: 8px;
}

.stat-label {
  font-size: 14px;
  font-weight: 600;
  color: var(--muted);
}

.generic-hint {
  font-size: 13px;
  color: var(--muted);
  font-style: italic;
  opacity: 0.7;
}
</style>
