<template>
  <section class="pub-certs pub-section" :class="[
    'layout-' + (data.layout || 'centered'),
    { 'hide-title': data.hide_title, 'hide-divider': data.hide_divider }
  ]" :style="data.section_bg_color ? `background: ${data.section_bg_color} !important; --item-bg: ${data.section_bg_color};` : ''">
    <div class="section-header-wrapper animate-slide-up" v-intersect>
      <h2 class="layered-title" :data-text="data.title || 'Certificates'">{{ data.title || 'Certificates & Awards' }}</h2>
    </div>
    
    <div class="pub-section-content">
      <div v-if="data.image_url" class="universal-section-img-container">
        <img :src="data.image_url" class="universal-section-img" alt="Section Cover" />
      </div>

      <div class="certs-grid">
        <component
          :is="cert.certificate_url ? 'a' : 'div'"
          v-for="(cert, i) in data.items"
          :key="i"
          :href="cert.certificate_url || undefined"
          :target="cert.certificate_url ? '_blank' : undefined"
          class="cert-card animate-scale-up premium-card-shine"
          :class="{ 'is-link': !!cert.certificate_url }"
          v-intersect
          :style="{ transitionDelay: `${Number(i) * 100}ms` }"
        >
          <div v-if="cert.image_url" class="cert-img-wrapper">
            <img :src="cert.image_url" :alt="cert.title" class="cert-img" loading="lazy" />
            <div v-if="cert.certificate_url" class="cert-link-overlay">
              <ExternalLink :size="24" />
            </div>
          </div>
          <div class="cert-content">
            <h3 class="cert-title">{{ cert.title }}</h3>
            <div class="cert-meta">
              <span class="cert-issuer">{{ cert.issuer }}</span>
              <span v-if="cert.date" class="cert-date">• {{ cert.date }}</span>
            </div>
            <p v-if="cert.description" class="cert-desc">{{ cert.description }}</p>
          </div>
        </component>
      </div>

      <div v-if="(!data.items || data.items.length === 0) && !data.image_url" style="text-align: center; padding: 20px;">
        <span class="generic-hint">คลิกเพื่อ edit ใน Properties panel &rarr;</span>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { ExternalLink, FolderOpen } from "lucide-vue-next";
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
:global(.theme-light) .pub-section {
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
.layout-left .certs-grid {
  justify-content: flex-start;
}

.certs-grid {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 24px;
}

.cert-card {
  width: calc(33.333% - 16px);
  min-width: 280px;
  max-width: 400px;
  flex: 0 1 auto;
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 16px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  text-decoration: none !important;
  color: inherit !important;
}

.cert-card.is-link {
  cursor: pointer;
}

.layout-centered .cert-card {
  flex-grow: 0;
}

.layout-left .cert-card,
.layout-split .cert-card {
  flex-grow: 1;
}

@media (max-width: 900px) {
  .cert-card {
    width: calc(50% - 12px);
  }
}

@media (max-width: 600px) {
  .cert-card {
    width: 100%;
  }
}

.cert-card:hover {
  transform: translateY(-5px);
  border-color: var(--primary);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.3);
}

:global(.theme-light) .cert-card:hover {
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.08);
}

.cert-img-wrapper {
  width: 100%;
  height: 180px;
  overflow: hidden;
  position: relative;
  background: rgba(0, 0, 0, 0.2);
  border-bottom: 1px solid var(--border);
}

.cert-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.6s ease;
}

.cert-link-overlay {
  position: absolute;
  inset: 0;
  background: rgba(var(--primary-rgb, 79, 70, 229), 0.4);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.3s ease;
  color: white;
  backdrop-filter: blur(4px);
}

.cert-card:hover .cert-link-overlay {
  opacity: 1;
}

.cert-card:hover .cert-img {
  transform: scale(1.1);
}

.cert-content {
  padding: 24px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  flex: 1;
}

.layout-centered .cert-content {
  align-items: center;
  text-align: center;
}

.cert-title {
  font-size: 19px;
  font-weight: 800;
  margin: 0;
  line-height: 1.25;
  color: var(--text);
  letter-spacing: -0.01em;
}

.cert-meta {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  font-weight: 300;
  flex-wrap: wrap;
}

.layout-centered .cert-meta {
  justify-content: center;
}

.cert-issuer {
  color: var(--surface);
  text-transform: uppercase;
  letter-spacing: 0.05em;
  font-size: 12px;
}

.cert-date {
  color: var(--muted);
  font-weight: 400;
}

.cert-desc {
  font-size: 10px;
  color: var(--muted);
  line-height: 1.6;
  margin: 0;
  margin-top: 6px;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.cert-acquired {
  margin-top: 12px;
  padding: 12px;
  background: rgba(var(--primary-rgb, 79, 70, 229), 0.05);
  border-radius: 8px;
  border-left: 3px solid var(--primary);
}

.acquired-label {
  display: block;
  font-size: 11px;
  font-weight: 800;
  text-transform: uppercase;
  color: var(--primary);
  margin-bottom: 4px;
  letter-spacing: 0.05em;
}

.acquired-text {
  font-size: 13px;
  color: var(--text);
  line-height: 1.5;
  margin: 0;
  opacity: 0.9;
}

.cert-action-link {
  margin-top: auto;
  padding-top: 12px;
  font-size: 13px;
  font-weight: 700;
  color: var(--text);
  display: flex;
  align-items: center;
  gap: 6px;
  opacity: 0.8;
  transition: opacity 0.2s;
}

.cert-card:hover .cert-action-link {
  opacity: 1;
}

.section-footer-link {
  margin-top: 5px;
  display: flex;
  justify-content: center;
  color: var(--muted);
}

.generic-hint {
  font-size: 13px;
  color: var(--muted);
  font-style: italic;
  opacity: 0.7;
}
</style>
