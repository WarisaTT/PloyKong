<!-- HeroPreview.vue -->
<template>
  <div class="hero-preview">
    <div class="hero-avatar">{{ (data.name || '?')[0].toUpperCase() }}</div>
    <div>
      <div class="hero-name">{{ data.name || 'Your Name' }}</div>
      <div class="hero-role">{{ data.role || 'Your Role' }}</div>
      <div class="hero-tagline">{{ data.tagline || 'Your tagline here...' }}</div>
    </div>
    <div v-if="data.show_hire_me" class="hire-badge">🎯 Hire Me</div>
  </div>
</template>
<script setup lang="ts">
defineProps<{ data: any }>()
</script>
<style scoped>
.hero-preview { display: flex; align-items: center; gap: 14px; flex-wrap: wrap; }
.hero-avatar {
  width: 44px; height: 44px; border-radius: 50%;
  background: linear-gradient(135deg, var(--indigo), var(--purple));
  display: flex; align-items: center; justify-content: center;
  font-family: var(--font-display); font-size: 18px; font-weight: 800; flex-shrink: 0;
}
.hero-name { font-size: 15px; font-weight: 700; }
.hero-role { font-size: 12px; color: var(--neon-purple); font-weight: 600; }
.hero-tagline { font-size: 11px; color: var(--muted); margin-top: 2px; }
.hire-badge {
  margin-left: auto; padding: 5px 14px; border-radius: 100px;
  background: linear-gradient(135deg, var(--indigo), var(--purple));
  font-size: 11px; font-weight: 700;
}
</style>
