<template>
  <section class="pub-exp pub-section" :class="[
    'layout-' + (data.layout || 'centered'),
    { 'hide-title': data.hide_title, 'hide-divider': data.hide_divider }
  ]" :style="data.section_bg_color ? { background: data.section_bg_color + ' !important' } : {}">
    <div class="section-header-wrapper">
      <h2 class="layered-title" :data-text="data.title || 'Experience'">{{ data.title || 'Experience' }}</h2>
    </div>
    <div v-if="data.image_url" class="universal-section-img-container">
      <img :src="data.image_url" class="universal-section-img" alt="Section Cover" />
    </div>
    <div class="experience-list">
      <div
        v-for="(item, i) in data.items"
        :key="i"
        class="experience-item animate-slide-in-right"
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
            <div class="tl-position">{{ item.position }}</div>
            <div class="tl-date">
              {{ item.start_date }} –
              {{ item.is_current ? "Present" : item.end_date }}
            </div>
          </div>
          <div class="tl-company">
            {{ item.company }}
            <span v-if="item.location" class="tl-loc"
              >· {{ item.location }}</span
            >
          </div>
          <div class="tl-desc">{{ item.description }}</div>
          
          <!-- Experience Image Gallery (New) -->
          <div v-if="item.image_urls && item.image_urls.length > 0" class="tl-gallery">
            <img 
              v-for="(img, imgIdx) in item.image_urls" 
              :key="imgIdx" 
              :src="img" 
              class="tl-gallery-img" 
              alt="Experience photo"
              loading="lazy"
            />
          </div>
          
        </div>
      </div>
    </div>
    <div v-if="(!data.items || data.items.length === 0) && !data.image_url" style="text-align: center; padding: 20px;">
      <span class="generic-hint">คลิกเพื่อ edit ใน Properties panel &rarr;</span>
    </div>
  </section>
</template>
<script setup lang="ts">
import { Link2 } from "lucide-vue-next";
defineProps<{ data: any }>();
</script>
<style scoped>
.pub-section.hide-divider { border-top: none !important; }
.pub-section.hide-title .layered-title { display: none !important; }
.pub-section {

  padding: 40px 0;
  border-top: 1px solid rgba(255, 255, 255, 0.06);
}
:global(.theme-light) .pub-section.hide-divider { border-top: none !important; }
.pub-section.hide-title .layered-title { display: none !important; }
.pub-section {

  border-top-color: rgba(0, 0, 0, 0.06);
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

.layout-left .experience-list {
  max-width: 100%;
  margin: 0;
}

.experience-list {
  display: flex;
  flex-direction: column;
  gap: 0;
  max-width: 600px;
  margin: 0 auto;
}
.experience-item {
  position: relative;
  padding-bottom: 32px;
}
.experience-item.has-timeline {
  padding-left: 32px; /* Increased padding from line */
}
.timeline-dot {
  position: absolute;
  left: 0px;
  top: 8px; /* Adjusted to align better with first line of text */
  width: 14px;
  height: 14px;
  border-radius: 50%;
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  flex-shrink: 0;
}
.timeline-line {
  position: absolute;
  left: 6px; /* Centered relative to the 14px dot */
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
  align-items: flex-start; /* align items to the start */
  gap: 4px;
  margin-bottom: 4px;
}
.tl-position {
  font-size: 16px; /* slightly larger */
  font-weight: 700;
  color: var(--title-color, var(--text));
}
.tl-date {
  font-size: 13px;
  color: var(--primary);
  font-weight: 600;
  white-space: nowrap; /* keep date on one line */
}
.tl-company {
  font-size: 13px;
  color: var(--muted, #94a3b8);
  margin-bottom: 8px;
}
.tl-loc {
  color: rgba(148, 163, 184, 0.6);
}
.tl-desc {
  font-size: 14px;
  line-height: 1.6;
  color: var(--muted);
  white-space: pre-wrap;
}

/* New Gallery Styles */
.tl-gallery {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  margin-top: 16px;
}
.layout-centered .tl-gallery {
  justify-content: center;
}
.tl-gallery-img {
  width: 120px;
  height: 80px;
  object-fit: cover;
  border-radius: 8px;
  border: 1px solid var(--border);
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}
.tl-gallery-img:hover {
  transform: translateY(-2px) scale(1.05);
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
  z-index: 2;
  position: relative;
}

.icon-inline {
  vertical-align: text-bottom;
  margin-right: 6px;
}
</style>
