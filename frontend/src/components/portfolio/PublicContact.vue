<template>
  <section id="contact" class="pub-contact pub-section">
    <h2 class="pub-section-title">
      <Mail :size="24" class="icon-inline" /> Contact
    </h2>
    <div class="contact-grid">
      <a v-if="data.email" :href="`mailto:${data.email}`" class="contact-card animate-scale-up" v-intersect>
        <Mail :size="24" class="contact-icon" />
        <span class="contact-label">Email</span>
        <span class="contact-val">{{ data.email }}</span>
      </a>
      <a v-if="data.linkedin" :href="data.linkedin" target="_blank" class="contact-card animate-scale-up" v-intersect style="transition-delay: 100ms;">
        <Linkedin :size="24" class="contact-icon" />
        <span class="contact-label">LinkedIn</span>
        <span class="contact-val">View Profile</span>
      </a>
      <a v-if="data.github" :href="data.github" target="_blank" class="contact-card animate-scale-up" v-intersect style="transition-delay: 200ms;">
        <Github :size="24" class="contact-icon" />
        <span class="contact-label">GitHub</span>
        <span class="contact-val">View Profile</span>
      </a>
      <div v-if="data.location" class="contact-card static animate-scale-up" v-intersect style="transition-delay: 300ms;">
        <MapPin :size="24" class="contact-icon" />
        <span class="contact-label">Location</span>
        <span class="contact-val">{{ data.location }}</span>
      </div>
    </div>
  </section>
</template>
<script setup lang="ts">
import { Mail, Linkedin, Github, MapPin } from 'lucide-vue-next'
defineProps<{ data: any }>()
</script>
<style scoped>
.pub-section {
  padding: 48px 0;
}

/* Dark theme */
:global(.theme-dark) .pub-section {
  border-top: 1px solid rgba(255, 255, 255, 0.06);
}

/* Light theme */
:global(.theme-light) .pub-section {
  border-top: 1px solid rgba(15, 23, 42, 0.08);
}

.pub-section-title {
  font-size: 22px;
  font-weight: 800;
  margin-bottom: 28px;
}

.contact-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: 14px;
}

.contact-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
  padding: 22px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 16px;
  text-decoration: none;
  color: var(--text, #e2e8f0);
  transition: all 0.2s;
  text-align: center;
}

:global(.theme-light) .contact-card {
  background: var(--surface);
  border-color: rgba(0, 0, 0, 0.08);
}

.contact-card:not(.static):hover {
  border-color: var(--primary);
  transform: translateY(-3px);
  background: var(--primary-glow);
}

:global(.theme-light) .contact-card:not(.static):hover {
  box-shadow: 0 16px 40px color-mix(in srgb, var(--primary) 15%, transparent);
}

.contact-icon {
  margin-bottom: 4px;
}

.contact-label {
  font-size: 11px;
  color: var(--muted, #94a3b8);
  font-weight: 600;
  letter-spacing: 1px;
  text-transform: uppercase;
}

.contact-val {
  font-size: 13px;
  font-weight: 600;
  color: var(--primary);
}

.icon-inline {
  vertical-align: text-bottom;
  margin-right: 6px;
}
</style>
