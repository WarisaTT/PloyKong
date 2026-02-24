<template>
  <section class="pub-custom pub-section">
    <div class="custom-content animate-slide-up" v-intersect :style="{ textAlign: data.alignment || 'left' }">
      {{ data.content }}
    </div>
  </section>
</template>
<script setup lang="ts">
defineProps<{ data: any }>()
</script>
<style scoped>
.pub-section { padding: 48px 0; border-top: 1px solid rgba(255,255,255,0.06); }
.custom-content { font-size: 15px; line-height: 1.8; color: var(--muted, #94a3b8); white-space: pre-wrap; }
</style>
