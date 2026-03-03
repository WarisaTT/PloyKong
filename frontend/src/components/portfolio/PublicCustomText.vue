<template>
  <section class="pub-custom pub-section" :class="[
    'layout-' + (data.layout || 'centered'),
    { 'hide-title': data.hide_title, 'hide-divider': data.hide_divider }
  ]" :style="data.section_bg_color ? { background: data.section_bg_color + ' !important' } : {}">
    <div v-if="data.title" class="section-header-wrapper">
      <h2 class="layered-title" :data-text="data.title">{{ data.title }}</h2>
    </div>
    <div class="pub-section-content">
      <div v-if="data.image_url" class="universal-section-img-container" style="margin-bottom: 24px;">
        <img :src="data.image_url" class="universal-section-img" alt="Section Cover" />
      </div>
      <div
        class="custom-content animate-slide-up"
        v-intersect
        :style="{ textAlign: data.alignment || 'left' }"
      >
        <template v-if="data.content">{{ data.content }}</template>
        <div v-else-if="!data.title && !data.image_url" class="empty-hint">
          <span class="generic-hint">คลิกเพื่อ edit ใน Properties panel &rarr;</span>
        </div>
      </div>
    </div>
  </section>
</template>
<script setup lang="ts">
import { computed } from 'vue';
defineProps<{ data: any }>();
</script>
<style scoped>
.pub-section.hide-divider { border-top: none !important; }
.pub-section.hide-title .layered-title { display: none !important; }
.pub-section {

  padding: 40px 0;
  border-top: 1px solid rgba(255, 255, 255, 0.06);
}
.custom-content {
  font-size: 15px;
  line-height: 1.8;
  color: var(--muted, #94a3b8);
  white-space: pre-wrap;
}

.empty-hint {
  text-align: center;
  padding: 20px;
}
.generic-hint {
  font-size: 13px;
  color: var(--muted);
  font-style: italic;
  opacity: 0.7;
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
.layout-left .custom-content {
  text-align: left !important;
}
</style>
