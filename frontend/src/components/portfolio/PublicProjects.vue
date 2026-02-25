<template>
  <section class="pub-projects pub-section">
    <div class="section-header-wrapper">
      <h2 class="layered-title" data-text="Projects">Projects</h2>
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
  </section>
</template>
<script setup lang="ts">
import { Rocket, Link as LinkIcon, Github } from "lucide-vue-next";
defineProps<{ data: any }>();
</script>
<style scoped>
.pub-section {
  padding: 48px 0;
  border-top: 1px solid rgba(255, 255, 255, 0.06);
}
.proj-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(260px, 1fr));
  gap: 20px;
}
.proj-card {
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
}
.proj-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  margin-bottom: 12px;
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
.proj-link {
  font-size: 12px;
  color: var(--neon-cyan, #22d3ee);
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
</style>
