<template>
  <section class="pub-skills pub-section" :class="[
    'layout-' + (data.layout || 'centered'),
    { 'hide-title': data.hide_title, 'hide-divider': data.hide_divider }
  ]" :style="data.section_bg_color ? `background: ${data.section_bg_color} !important; --item-bg: ${data.section_bg_color};` : ''">
    <div class="section-header-wrapper">
      <h2 class="layered-title" :data-text="data.title || 'Skills'">{{ data.title || 'Skills' }}</h2>
    </div>
    <div class="pub-section-content">
      <div v-if="data.image_url" class="universal-section-img-container">
        <img :src="data.image_url" class="universal-section-img" alt="Section Cover" />
      </div>
      <div v-if="groupedSkills.length > 0" class="skills-groups">
        <div v-for="(group, gIdx) in groupedSkills" :key="gIdx" class="skills-group">
          <h3 v-if="group.category && (group.category !== 'Skills' || groupedSkills.length > 1)" class="group-title">{{ group.category }}</h3>
          <div class="skills-grid grid-tags">
            <div
              v-for="(skill, i) in group.items"
              :key="i"
              class="skill-item skill-tag animate-slide-in-left is-visible"
              v-intersect
              :style="{ transitionDelay: `${(Number(i) % 10) * 50}ms` }"
            >
              <div class="skill-header">
                <span class="skill-name">{{ skill.name }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div v-if="groupedSkills.length === 0 && !data.image_url" style="text-align: center; padding: 20px;">
        <span class="generic-hint">คลิกเพื่อ edit ใน Properties panel &rarr;</span>
      </div>
    </div>
  </section>
</template>
<script setup lang="ts">
import { computed } from "vue";

const props = defineProps<{ data: any }>();

const groupedSkills = computed(() => {
  if (!props.data || !props.data.items) return [];

  const groups: Record<string, any[]> = {};
  const items = props.data.items as any[];

  items.forEach(skill => {
    const cat = skill.category || "";
    if (!groups[cat]) {
      groups[cat] = [];
    }
    groups[cat].push(skill);
  });

  return Object.keys(groups).map(key => ({
    category: key,
    items: groups[key]
  }));
});
</script>
<style scoped>
.pub-section.hide-divider { border-top: none !important; }
.pub-section.hide-title .layered-title { display: none !important; }
.pub-section {

  padding: 40px 0;
  border-top: 1px solid var(--section-border);
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
.layout-left .section-header-wrapper,
.layout-split .section-header-wrapper {
  text-align: left;
}
.layout-left .skills-grid {
  justify-content: flex-start;
}
.layout-split .skills-grid {
  justify-content: flex-start;
}
.layout-split .skill-header{
  display: flex !important;
  flex-wrap: wrap;
  gap: 12px;
  justify-content: flex-start;
}

.layout-centered .group-title {
  text-align: center;
}
.layout-centered .grid-tags {
  justify-content: center;
}

:global(.theme-dark) {
  --section-border: rgba(255, 255, 255, 0.06);
}

:global(.theme-light) {
  --section-border: rgba(0, 0, 0, 0.06);
}

.skills-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}
@media (max-width: 600px) {
  .skills-grid {
    grid-template-columns: 1fr;
  }
}

.skill-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 6px;
}
.skill-name {
  font-size: 14px;
  font-weight: 600;
}
.skills-group {
  margin-bottom: 32px;
}
.group-title {
  font-size: 16px;
  font-weight: 700;
  margin-bottom: 16px;
  color: var(--primary);
  opacity: 0.9;
}

/* Tag style */
.grid-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  justify-content: center;
}
.skill-tag {
  background: var(--primary-glow);
  color: var(--primary);
  border: 1px solid var(--primary);
  padding: 6px 16px;
  border-radius: 100px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
}
.skill-tag.is-visible {
  opacity: 1 !important;
  transform: none !important;
}
.skill-tag:hover {
  transform: translateY(-2px) !important;
}
.skill-tag .skill-header {
  margin-bottom: 0;
  justify-content: center;
  display: inline-flex;
  width: auto;
}
.skill-tag .skill-name {
  font-size: 13px;
  font-weight: 700;
}

:global(.theme-light) .skill-tag {
  background: var(--primary-glow);
}

.icon-inline {
  vertical-align: text-bottom;
  margin-right: 6px;
}

.generic-hint {
  font-size: 13px;
  color: var(--muted);
  font-style: italic;
  opacity: 0.7;
}
</style>
