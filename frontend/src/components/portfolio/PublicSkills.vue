<template>
  <section class="pub-skills pub-section">
    <h2 class="pub-section-title"><Wrench :size="24" class="icon-inline" /> Skills</h2>
    <div class="skills-grid">
      <div v-for="(skill, i) in data.items" :key="i" class="skill-tag animate-slide-in-left" v-intersect :style="{ transitionDelay: `${(Number(i) % 10) * 50}ms` }">
        <div class="skill-header">
          <span class="skill-name">{{ skill.name }}</span>
          <span class="skill-pct">{{ skill.level }}%</span>
        </div>
        <div class="skill-bar-bg">
          <div class="skill-bar-fill" :style="{ width: skill.level + '%' }"></div>
        </div>
      </div>
    </div>
  </section>
</template>
<script setup lang="ts">
import { Wrench } from 'lucide-vue-next'
defineProps<{ data: any }>()
</script>
<style scoped>
.pub-section { padding: 48px 0; border-top: 1px solid rgba(255,255,255,0.06); }
:global(.theme-light) .pub-section { border-top-color: rgba(0,0,0,0.06); }
.pub-section-title { font-size: 22px; font-weight: 800; margin-bottom: 28px; }
.skills-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 16px; }
@media (max-width: 600px) { .skills-grid { grid-template-columns: 1fr; } }

.skill-header { display: flex; justify-content: space-between; margin-bottom: 6px; }
.skill-name { font-size: 14px; font-weight: 600; }
.skill-pct { font-size: 12px; color: var(--muted, #94a3b8); }
.skill-bar-bg { height: 6px; background: rgba(255,255,255,0.08); border-radius: 3px; overflow: hidden; }
:global(.theme-light) .skill-bar-bg { background: rgba(0,0,0,0.15); }
.skill-bar-fill { height: 100%; background: linear-gradient(90deg, var(--primary), var(--secondary)); border-radius: 3px; transition: width 1s ease; }
:global(.theme-light) .skill-bar-fill { background: linear-gradient(90deg, var(--primary), var(--secondary)); }

.icon-inline { vertical-align: text-bottom; margin-right: 6px; }
</style>
