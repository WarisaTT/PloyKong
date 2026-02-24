<!-- HeroPreview.vue -->
<template>
  <div class="hero-preview">
    <img v-if="data.avatar_url" :src="data.avatar_url" class="hero-avatar hero-img" />
    <div v-else class="hero-avatar">{{ (data.name || '?')[0].toUpperCase() }}</div>
    <div>
      <div class="hero-name">{{ data.name || 'Your Name' }}</div>
      <div class="hero-role">{{ data.role || 'Your Role' }}</div>
      <div class="hero-tagline">{{ data.tagline || 'Your tagline here...' }}</div>
    </div>
    <div v-if="data.show_hire_me" class="hire-badge">
      <Send :size="14" />
      Hire Me
    </div>
  </div>
</template>
<script setup lang="ts">
import { Send } from 'lucide-vue-next'
defineProps<{ data: any }>()
</script>

<style scoped>
.hero-preview { display: flex; align-items: center; gap: 14px; flex-wrap: wrap; }
.hero-avatar {
  width: 44px; height: 44px; border-radius: 50%;
  background: linear-gradient(135deg, var(--indigo), var(--purple));
  color: #ffffff;
  display: flex; align-items: center; justify-content: center;
  font-family: var(--font-display); font-size: 18px; font-weight: 800; flex-shrink: 0;
}
.hero-img { object-fit: cover; }
.hero-name { font-size: 15px; font-weight: 700; }
.hero-role { font-size: 12px; color: var(--neon-purple); font-weight: 600; }
.hero-tagline { font-size: 11px; color: var(--muted); margin-top: 2px; }
.hire-badge {
  margin-left: auto;
  padding: 6px 14px;
  border-radius: 100px;
  background: linear-gradient(135deg, var(--indigo), var(--purple));
  color: #ffffff;
  font-size: 11px;
  font-weight: 700;
  display: inline-flex;
  align-items: center;
  gap: 6px;
}
:global(.theme-light) .hero-name { color: #111827; }
:global(.theme-light) .hero-role { color: var(--indigo); font-weight: 700; }
:global(.theme-light) .hero-tagline { color: #4b5563; }
:global(.theme-light) .hire-badge {
  background: linear-gradient(135deg, var(--indigo), color-mix(in srgb, var(--indigo) 70%, black));
  box-shadow: 0 4px 12px color-mix(in srgb, var(--indigo) 30%, transparent);
}
</style>

