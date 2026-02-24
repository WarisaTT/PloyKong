<template>
  <div class="contact-preview">
    <div v-if="data.email" class="contact-row">📧 {{ data.email }}</div>
    <div v-if="data.linkedin" class="contact-row">💼 {{ data.linkedin }}</div>
    <div v-if="data.github" class="contact-row">🐙 {{ data.github }}</div>
    <div v-if="data.location" class="contact-row">📍 {{ data.location }}</div>
    <div v-if="!data.email && !data.linkedin" class="empty-msg">ยังไม่มีข้อมูลติดต่อ</div>
  </div>
</template>
<script setup lang="ts">
defineProps<{ data: any }>()
</script>
<style scoped>
.contact-preview { display: flex; flex-direction: column; gap: 6px; }
.contact-row { font-size: 12px; color: var(--muted); }
.empty-msg { font-size: 12px; color: var(--muted); font-style: italic; }
</style>
