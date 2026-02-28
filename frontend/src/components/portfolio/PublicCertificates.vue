<template>
  <section class="pub-certs pub-section" :class="'layout-' + (data.layout || 'centered')" :style="data.section_bg_color ? { background: data.section_bg_color } : {}">
    <div class="section-header-wrapper">
      <h2 class="layered-title" data-text="Certificates">Certificates & Awards</h2>
    </div>
    <div v-if="data.image_url" class="universal-section-img-container">
      <img :src="data.image_url" class="universal-section-img" alt="Section Cover" />
    </div>

    <div class="certs-grid">
      <div
        v-for="(cert, i) in data.items"
        :key="i"
        class="cert-card animate-scale-up"
        v-intersect
        :style="{ transitionDelay: `${Number(i) * 100}ms` }"
      >
        <div v-if="cert.image_url" class="cert-img-wrapper">
          <img :src="cert.image_url" :alt="cert.title" class="cert-img" loading="lazy" />
        </div>
        <div class="cert-content">
          <h3 class="cert-title">{{ cert.title }}</h3>
          <div class="cert-meta">
            <span class="cert-issuer">{{ cert.issuer }}</span>
            <span v-if="cert.date" class="cert-date">• {{ cert.date }}</span>
          </div>
          <p v-if="cert.description" class="cert-desc">{{ cert.description }}</p>
        </div>
      </div>
    </div>

    <div v-if="(!data.items || data.items.length === 0) && !data.image_url" style="text-align: center; padding: 20px;">
      <span class="generic-hint">คลิกเพื่อ edit ใน Properties panel &rarr;</span>
    </div>
  </section>
</template>

<script setup lang="ts">
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
  flex-grow: 1;
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 16px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  transition: transform 0.3s ease, box-shadow 0.3s ease, border-color 0.3s ease;
}
.cert-card:hover {
  transform: translateY(-4px);
  border-color: rgba(255, 255, 255, 0.2);
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
}
:global(.theme-light) .cert-card:hover {
  border-color: rgba(0, 0, 0, 0.15);
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.05);
}

.cert-img-wrapper {
  width: 100%;
  height: 180px;
  overflow: hidden;
  background: rgba(0, 0, 0, 0.2);
  border-bottom: 1px solid var(--border);
}
.cert-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.5s ease;
}
.cert-card:hover .cert-img {
  transform: scale(1.05);
}

.cert-content {
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 8px;
  flex: 1;
}
.layout-centered .cert-content {
  align-items: center;
  text-align: center;
}

.cert-title {
  font-size: 18px;
  font-weight: 700;
  margin: 0;
  line-height: 1.3;
  color: var(--text);
}

.cert-meta {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  font-weight: 500;
  flex-wrap: wrap;
}
.layout-centered .cert-meta {
  justify-content: center;
}
.cert-issuer {
  color: var(--primary);
}
.cert-date {
  color: var(--muted);
}

.cert-desc {
  font-size: 14px;
  color: var(--muted);
  line-height: 1.6;
  margin: 0;
  margin-top: 4px;
}

.generic-hint {
  font-size: 13px;
  color: var(--muted);
  font-style: italic;
  opacity: 0.7;
}
</style>
