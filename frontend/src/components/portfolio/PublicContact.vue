<template>
  <section id="contact" class="pub-contact pub-section" :class="'layout-' + (data.layout || 'centered')">
    <div class="section-header-wrapper">
      <h2 class="layered-title" data-text="Contact">Contact</h2>
    </div>
    <div v-if="data.image_url" class="universal-section-img-container">
      <img :src="data.image_url" class="universal-section-img" alt="Section Cover" />
    </div>
    <div class="contact-grid">
      <a
        v-if="data.email"
        :href="`mailto:${data.email}`"
        class="contact-card animate-scale-up"
        v-intersect
      >
        <Mail :size="24" class="contact-icon" />
        <span class="contact-label">Email</span>
        <span class="contact-val">{{ data.email }}</span>
      </a>
      <a
        v-if="data.linkedin"
        :href="data.linkedin"
        target="_blank"
        class="contact-card animate-scale-up"
        v-intersect
        style="transition-delay: 100ms"
      >
        <Linkedin :size="24" class="contact-icon" />
        <span class="contact-label">LinkedIn</span>
        <span class="contact-val">View Profile</span>
      </a>
      <a
        v-if="data.github"
        :href="data.github"
        target="_blank"
        class="contact-card animate-scale-up"
        v-intersect
        style="transition-delay: 200ms"
      >
        <Github :size="24" class="contact-icon" />
        <span class="contact-label">GitHub</span>
        <span class="contact-val">View Profile</span>
      </a>
      <div
        v-if="data.location"
        class="contact-card static animate-scale-up"
        v-intersect
        style="transition-delay: 300ms"
      >
        <MapPin :size="24" class="contact-icon" />
        <span class="contact-label">Location</span>
        <span class="contact-val">{{ data.location }}</span>
      </div>

      <!-- Custom Contacts -->
      <component
        v-for="(contact, i) in data.custom_items"
        :key="contact.id"
        :is="contact.link ? 'a' : 'div'"
        :href="contact.link || undefined"
        :target="
          contact.link &&
          !contact.link.startsWith('tel:') &&
          !contact.link.startsWith('mailto:')
            ? '_blank'
            : undefined
        "
        class="contact-card animate-scale-up"
        :class="{ static: !contact.link }"
        v-intersect
        :style="{ transitionDelay: `${400 + Number(i) * 100}ms` }"
      >
        <component
          :is="getIcon(contact.platform)"
          :size="24"
          class="contact-icon"
        />
        <span class="contact-label">{{ contact.label }}</span>
        <span class="contact-val">{{ contact.value }}</span>
      </component>
    </div>
  </section>
</template>
<script setup lang="ts">
import {
  Mail,
  Linkedin,
  Github,
  MapPin,
  Facebook,
  Twitter,
  Instagram,
  Youtube,
  MessageCircle,
  Phone,
  Link,
} from "lucide-vue-next";

defineProps<{ data: any }>();

function getIcon(platform: string) {
  switch (platform) {
    case "facebook":
      return Facebook;
    case "twitter":
      return Twitter;
    case "instagram":
      return Instagram;
    case "youtube":
      return Youtube;
    case "message-circle":
      return MessageCircle;
    case "phone":
      return Phone;
    default:
      return Link;
  }
}
</script>
<style scoped>
.pub-section.hide-divider { border-top: none !important; }
.pub-section.hide-title .layered-title { display: none !important; }
.pub-section {

  padding: 40px 0;
}

/* Dark theme */
:global(.theme-dark) .pub-section.hide-divider { border-top: none !important; }
.pub-section.hide-title .layered-title { display: none !important; }
.pub-section {

  border-top: 1px solid rgba(255, 255, 255, 0.06);
}

/* Light theme */
:global(.theme-light) .pub-section.hide-divider { border-top: none !important; }
.pub-section.hide-title .layered-title { display: none !important; }
.pub-section {

  border-top: 1px solid rgba(15, 23, 42, 0.08);
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
.layout-left .contact-grid {
  justify-content: flex-start;
}

.contact-grid {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 14px;
}

.contact-card {
  width: calc(33.333% - 10px);
  min-width: 180px;
  max-width: 280px;
  flex-grow: 1;
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
  border-color: rgba(0, 0, 0, 0.15);
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
