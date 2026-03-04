<template>
  <section class="pub-social pub-section" :class="[
    'layout-' + (data.layout || 'centered'),
    { 'hide-title': data.hide_title, 'hide-divider': data.hide_divider }
  ]" :style="data.section_bg_color ? `background: ${data.section_bg_color} !important; --item-bg: ${data.section_bg_color};` : ''">
    <div class="section-header-wrapper">
      <h2 class="layered-title" :data-text="data.title || 'Connect'">{{ data.title || 'Connect Here' }}</h2>
    </div>
    
    <div class="pub-section-content">
      <div class="social-grid">
        <a 
          v-for="(link, i) in data.items" 
          :key="i"
          :href="link.url"
          target="_blank"
          rel="noopener noreferrer"
          class="social-btn animate-scale-up"
          v-intersect
          :style="{ transitionDelay: `${Number(i) * 50}ms` }"
        >
          <component :is="getIcon(link.platform)" size="20" class="social-icon" />
          <span class="social-label">{{ link.label || capitalize(link.platform) }}</span>
        </a>
      </div>

      <div v-if="!data.items || data.items.length === 0" style="text-align: center; padding: 20px;">
        <span class="generic-hint">คลิกเพื่อ edit ใน Properties panel &rarr;</span>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { Facebook, Twitter, Instagram, Youtube, Linkedin, Github, Globe, MessageCircle } from 'lucide-vue-next';

defineProps<{ data: any }>();

function getIcon(platform: string) {
  const map: Record<string, any> = {
    facebook: Facebook,
    twitter: Twitter,
    instagram: Instagram,
    youtube: Youtube,
    linkedin: Linkedin,
    github: Github,
    line: MessageCircle,
    message: MessageCircle,
  };
  return map[platform?.toLowerCase()] || Globe;
}

function capitalize(str: string) {
  if (!str) return 'Website';
  return str.charAt(0).toUpperCase() + str.slice(1);
}
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
.layout-left .social-grid {
  justify-content: flex-start;
}

.social-grid {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 16px;
}

.social-btn {
  display: inline-flex;
  align-items: center;
  gap: 10px;
  padding: 12px 24px;
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 100px;
  color: var(--text);
  text-decoration: none;
  font-weight: 600;
  font-size: 14px;
  transition: all 0.2s ease;
}

.social-btn:hover {
  transform: translateY(-2px);
  background: rgba(79, 70, 229, 0.1);
  border-color: var(--primary);
  color: var(--primary);
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
}

.social-icon {
  opacity: 0.8;
}

.social-btn:hover .social-icon {
  opacity: 1;
}

.generic-hint {
  font-size: 13px;
  color: var(--muted);
  font-style: italic;
  opacity: 0.7;
}
</style>
