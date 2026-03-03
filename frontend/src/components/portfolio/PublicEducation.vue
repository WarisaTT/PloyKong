<template>
  <section class="pub-edu pub-section" :class="[
    'layout-' + (data.layout || 'centered'),
    { 'hide-title': data.hide_title, 'hide-divider': data.hide_divider }
  ]" :style="data.section_bg_color ? { background: data.section_bg_color } : {}">
    <div class="section-header-wrapper">
      <h2 class="layered-title" :data-text="data.title || 'Education'">{{ data.title || 'Education' }}</h2>
    </div>
    <div class="pub-section-content">
      <div v-if="data.image_url" class="universal-section-img-container">
        <img :src="data.image_url" class="universal-section-img" alt="Section Cover" />
      </div>

      <div class="education-list">
        <div
          v-for="(item, i) in data.items"
          :key="i"
          class="education-item animate-slide-in-right"
          :class="{ 'has-timeline': data.items.length > 1 }"
          v-intersect
          :style="{ transitionDelay: `${Number(i) * 100}ms` }"
        >
          <div v-if="data.items.length > 1" class="timeline-dot"></div>
          <div
            class="timeline-line"
            v-if="data.items.length > 1 && Number(i) < data.items.length - 1"
          ></div>
          <div class="timeline-content">
            <div class="tl-header">
              <div class="tl-position">{{ item.degree }} <span v-if="item.field">in {{ item.field }}</span></div>
              <div class="tl-date">
                {{ item.start_year }} – {{ item.end_year || "Present" }}
              </div>
            </div>
            <div class="tl-company">
              {{ item.school }}
              <span v-if="item.gpa" class="tl-loc"
                >· GPA: {{ item.gpa }}</span
              >
            </div>
          </div>
        </div>
      </div>
      <div v-if="(!data.items || data.items.length === 0) && !data.image_url" style="text-align: center; padding: 20px;">
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

.universal-section-img-container {
  margin-bottom: 32px;
  text-align: center;
}
.universal-section-img {
  width: 100%;
  max-width: 800px;
  height: 250px;
  object-fit: cover;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0,0,0,0.1);
}

/* Layout alignment */
.layout-left .section-header-wrapper {
  text-align: left;
}

.layout-left .education-list {
  max-width: 100%;
  margin: 0;
}

.education-list {
  display: flex;
  flex-direction: column;
  gap: 0;
  max-width: 600px;
  margin: 0 auto;
}
.education-item {
  position: relative;
  padding-bottom: 32px;
}
.education-item.has-timeline {
  padding-left: 32px;
}
.timeline-dot {
  position: absolute;
  left: 0px;
  top: 8px;
  width: 14px;
  height: 14px;
  border-radius: 50%;
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  flex-shrink: 0;
}
.timeline-line {
  position: absolute;
  left: 6px;
  top: 24px;
  bottom: 0;
  width: 2px;
  background: var(--line-color);
}
:global(.theme-light) .timeline-line {
  --line-color: rgba(0, 0, 0, 0.5);
}

:global(.theme-dark) .timeline-line {
  --line-color: rgba(255, 255, 255, 0.5);
}

.tl-header {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 4px;
  margin-bottom: 4px;
}
.tl-position {
  font-size: 16px;
  font-weight: 700;
  color: var(--title-color, var(--text));
}
.tl-date {
  font-size: 13px;
  color: var(--primary);
  font-weight: 600;
  white-space: nowrap;
}
.tl-company {
  font-size: 14px;
  color: var(--text);
  margin-bottom: 8px;
}
.tl-loc {
  color: var(--muted);
}

.generic-hint {
  font-size: 13px;
  color: var(--muted);
  font-style: italic;
  opacity: 0.7;
}
</style>
