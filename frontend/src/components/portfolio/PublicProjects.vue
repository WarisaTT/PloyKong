<template>
  <section class="pub-projects pub-section" :class="[
    'layout-' + (data.layout || 'centered'),
    { 'hide-title': data.hide_title, 'hide-divider': data.hide_divider }
  ]" :style="data.section_bg_color ? { background: data.section_bg_color + ' !important' } : {}">
    <div class="section-header-wrapper">
      <h2 class="layered-title" :data-text="data.title || 'Projects'">{{ data.title || 'Projects' }}</h2>
    </div>
    <div v-if="data.image_url" class="universal-section-img-container">
      <img :src="data.image_url" class="universal-section-img" alt="Section Cover" />
    </div>
    <div class="proj-grid">
      <div
        v-for="(proj, i) in data.items"
        :key="i"
        class="proj-card animate-scale-up"
        v-intersect
        :style="{ transitionDelay: `${Number(i) * 100}ms` }"
      >
        <img
          v-if="proj.image_url"
          :src="proj.image_url"
          :alt="proj.title"
          class="proj-img"
        />
        <div class="proj-body">
          <div class="proj-title">{{ proj.title }}</div>
          <div class="proj-desc">{{ proj.description }}</div>
          <div class="proj-tags">
            <span v-for="tag in proj.tags" :key="tag" class="proj-tag">{{
              tag
            }}</span>
          </div>
          <div class="proj-links">
            <a
              v-if="proj.live_url"
              :href="proj.live_url"
              target="_blank"
              class="proj-link"
              ><LinkIcon :size="12" class="icon-sm" /> Live</a
            >
            <a
              v-if="proj.github_url"
              :href="proj.github_url"
              target="_blank"
              class="proj-link"
              ><Github :size="12" class="icon-sm" /> GitHub</a
            >
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
import { Rocket, Link as LinkIcon, Github } from "lucide-vue-next";
defineProps<{ data: any }>();
</script>
<style scoped>
.pub-section.hide-divider { border-top: none !important; }
.pub-section.hide-title .layered-title { display: none !important; }
.pub-section {
  padding: 20px 0;
  border-top: 1px solid rgba(255, 255, 255, 0.06);
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
.layout-left .proj-grid {
  justify-content: flex-start;
}

@media (max-width: 600px) {
  .layout-split .proj-card,
  .layout-left .proj-card {
    width: 100%;
  }
}

.proj-grid {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 20px;
}
.proj-card {
  width: calc(33.333% - 14px);
  min-width: 260px;
  max-width: 380px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 16px;
  overflow: hidden;
  transition: all 0.25s;
}
.proj-card:hover {
  transform: translateY(-4px);
  border-color: var(--primary);
  box-shadow: 0 16px 40px rgba(0, 0, 0, 0.4);
}

:global(.theme-light) .proj-card {
  background: var(--surface);
  border-color: rgba(0, 0, 0, 0.08);
}
:global(.theme-light) .proj-card:hover {
  border-color: var(--primary);
  box-shadow: 0 16px 40px color-mix(in srgb, var(--primary) 20%, transparent);
}
.proj-img {
  width: 100%;
  height: 160px;
  object-fit: cover;
}
.proj-body {
  padding: 18px;
}

.layout-centered .proj-body {
  text-align: center;
}
.layout-split .proj-body {
  text-align: left;
}
.proj-title {
  font-size: 15px;
  font-weight: 700;
  margin-bottom: 8px;
}
.proj-desc {
  font-size: 13px;
  color: var(--muted, #94a3b8);
  line-height: 1.6;
  margin-bottom: 12px;
  word-break: break-word; /* Fix: handle long unbreakable strings */
  overflow-wrap: break-word;
}
.proj-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  margin-bottom: 12px;
}
.layout-centered .proj-tags {
  justify-content: center;
}
.layout-split .proj-tags {
  justify-content: flex-start;
}

.proj-tag {
  font-size: 11px;
  padding: 3px 10px;
  border-radius: 100px;
  background: var(--primary-glow);
  color: var(--primary);
  border: 1px solid var(--primary);
}
.proj-links {
  display: flex;
  gap: 10px;
}
.layout-centered .proj-links {
  justify-content: center;
}
.layout-split .proj-links {
  justify-content: flex-start;
}
.proj-link {
  font-size: 12px;
  color: var(--neon-cyan, --primary-glow);
  text-decoration: none;
  font-weight: 600;
  display: inline-flex;
  align-items: center;
  gap: 4px;
}
.proj-link:hover {
  text-decoration: underline;
}

.icon-inline {
  vertical-align: text-bottom;
  margin-right: 6px;
}
.icon-sm {
  vertical-align: -2px;
}

.generic-hint {
  font-size: 13px;
  color: var(--muted);
  font-style: italic;
  opacity: 0.7;
}
</style>
