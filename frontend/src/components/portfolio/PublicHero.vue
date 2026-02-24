<!-- PublicHero.vue - renders the hero section on the public portfolio page -->
<template>
  <section class="pub-hero">
    <div class="hero-avatar-wrap">
      <img v-if="data.avatar_url" :src="data.avatar_url" :alt="data.name" class="hero-img" />
      <div v-else class="hero-avatar-fallback">{{ (data.name || '?')[0] }}</div>
      <div class="hero-ring"></div>
    </div>
    <div class="hero-text">
      <div class="hero-role-badge">{{ data.role }}</div>
      <h1 class="hero-name">{{ data.name }}</h1>
      <p class="hero-tagline">{{ data.tagline }}</p>
      <div v-if="data.show_hire_me" class="hero-cta">
        <a :href="data.hire_me_link || 'mailto:'" class="btn btn-primary" @click="trackHire">
          🎯 Hire Me
        </a>
        <a href="#contact" class="btn btn-secondary">📬 Contact</a>
      </div>
    </div>
  </section>
</template>
<script setup lang="ts">
import { publicAPI } from '@/api'
import { useRoute } from 'vue-router'
const route = useRoute()
defineProps<{ data: any; theme?: any; portfolio?: any }>()
function trackHire() {
  publicAPI.track(route.params.slug as string, 'hire_click')
}
</script>
<style scoped>
.pub-hero { display: flex; flex-direction: column; align-items: center; text-align: center; padding: 80px 0 60px; gap: 24px; }
.hero-avatar-wrap { position: relative; width: 120px; height: 120px; }
.hero-img { width: 120px; height: 120px; border-radius: 50%; object-fit: cover; }
.hero-avatar-fallback {
  width: 120px; height: 120px; border-radius: 50%;
  background: linear-gradient(135deg, var(--indigo, #4F46E5), var(--purple, #a855f7));
  display: flex; align-items: center; justify-content: center;
  font-family: var(--font-display); font-size: 48px; font-weight: 800; color: #fff;
}
.hero-ring {
  position: absolute; inset: -6px; border-radius: 50%;
  border: 2px solid transparent;
  background: linear-gradient(135deg, var(--primary, #4F46E5), #22d3ee) border-box;
  -webkit-mask: linear-gradient(#fff 0 0) padding-box, linear-gradient(#fff 0 0);
  -webkit-mask-composite: destination-out;
  mask-composite: exclude;
  animation: spin 4s linear infinite;
}
@keyframes spin { to { transform: rotate(360deg); } }
.hero-role-badge { display: inline-block; padding: 5px 16px; border-radius: 100px; background: rgba(79,70,229,0.15); color: #818cf8; font-size: 13px; font-weight: 600; border: 1px solid rgba(79,70,229,0.3); margin-bottom: 12px; }
.hero-name { font-size: clamp(32px, 5vw, 52px); font-weight: 800; margin-bottom: 12px; }
.hero-tagline { font-size: 16px; color: var(--muted, #94a3b8); max-width: 500px; line-height: 1.7; margin-bottom: 20px; }
.hero-cta { display: flex; gap: 12px; justify-content: center; flex-wrap: wrap; }
</style>
