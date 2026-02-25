<template>
  <section class="pub-skills pub-section">
    <div class="section-header-wrapper">
      <h2 class="layered-title" data-text="Skills">Skills</h2>
    </div>
    <div class="skills-grid">
      <div
        v-for="(skill, i) in data.items"
        :key="i"
        class="skill-tag animate-slide-in-left"
        v-intersect
        :style="{ transitionDelay: `${(Number(i) % 10) * 50}ms` }"
      >
        <div class="skill-header">
          <span class="skill-name">{{ skill.name }}</span>
          <span class="skill-pct">{{ skill.level }}%</span>
        </div>
        <div class="skill-bar-bg">
          <div
            class="skill-bar-fill"
            :style="{ width: skill.level + '%' }"
          ></div>
        </div>
      </div>
    </div>
  </section>
</template>
<script setup lang="ts">
defineProps<{ data: any }>();
</script>
<style scoped>
.pub-section {
  padding: 80px 0;
  border-top: 1px solid var(--section-border);
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
.skill-pct {
  font-size: 12px;
  color: var(--muted, #94a3b8);
}
.skill-bar-bg {
  height: 6px;
  border-radius: 3px;
  overflow: hidden;
  background: var(--skill-bg);
}

:global(.theme-dark) {
  --skill-bg: rgba(255, 255, 255, 0.08);
}

:global(.theme-light) {
  --skill-bg: rgba(0, 0, 0, 0.08);
}

.skill-bar-fill {
  height: 100%;
  background: linear-gradient(90deg, var(--primary), var(--secondary));
  border-radius: 3px;
  transition: width 1s ease;
}
:global(.theme-light) .skill-bar-fill {
  background: linear-gradient(20deg, #4f46e5, #8b5cf6);
}

.icon-inline {
  vertical-align: text-bottom;
  margin-right: 6px;
}
</style>
