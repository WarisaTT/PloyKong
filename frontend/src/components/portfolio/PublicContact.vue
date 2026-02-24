<template>
  <section id="contact" class="pub-contact pub-section">
    <h2 class="pub-section-title">📬 Contact</h2>
    <div class="contact-grid">
      <a v-if="data.email" :href="`mailto:${data.email}`" class="contact-card">
        <span class="contact-icon">📧</span>
        <span class="contact-label">Email</span>
        <span class="contact-val">{{ data.email }}</span>
      </a>
      <a v-if="data.linkedin" :href="data.linkedin" target="_blank" class="contact-card">
        <span class="contact-icon">💼</span>
        <span class="contact-label">LinkedIn</span>
        <span class="contact-val">View Profile</span>
      </a>
      <a v-if="data.github" :href="data.github" target="_blank" class="contact-card">
        <span class="contact-icon">🐙</span>
        <span class="contact-label">GitHub</span>
        <span class="contact-val">View Profile</span>
      </a>
      <div v-if="data.location" class="contact-card static">
        <span class="contact-icon">📍</span>
        <span class="contact-label">Location</span>
        <span class="contact-val">{{ data.location }}</span>
      </div>
    </div>
  </section>
</template>
<script setup lang="ts">
defineProps<{ data: any }>()
</script>
<style scoped>
.pub-section { padding: 48px 0; border-top: 1px solid rgba(255,255,255,0.06); }
.pub-section-title { font-size: 22px; font-weight: 800; margin-bottom: 28px; }
.contact-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(180px, 1fr)); gap: 14px; }
.contact-card {
  display: flex; flex-direction: column; align-items: center; gap: 6px;
  padding: 22px; background: rgba(255,255,255,0.03); border: 1px solid rgba(255,255,255,0.08);
  border-radius: 16px; text-decoration: none; color: var(--text, #e2e8f0);
  transition: all 0.2s; text-align: center;
}
.contact-card:not(.static):hover { border-color: rgba(79,70,229,0.4); transform: translateY(-3px); background: rgba(79,70,229,0.06); }
.contact-icon { font-size: 24px; }
.contact-label { font-size: 11px; color: var(--muted, #94a3b8); font-weight: 600; letter-spacing: 1px; text-transform: uppercase; }
.contact-val { font-size: 13px; font-weight: 600; color: var(--neon-cyan, #22d3ee); }
</style>
