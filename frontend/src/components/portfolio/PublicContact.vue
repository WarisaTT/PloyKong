<template>
  <section id="contact" class="pub-contact pub-section" :class="'layout-' + (data.layout || 'centered')" :style="data.section_bg_color ? `background: ${data.section_bg_color} !important; --item-bg: ${data.section_bg_color};` : ''">
    <div class="section-header-wrapper">
      <h2 class="layered-title" :data-text="data.title || 'Contact'">{{ data.title || 'Contact' }}</h2>
    </div>
    <div v-if="data.image_url" class="universal-section-img-container">
      <img :src="data.image_url" class="universal-section-img" alt="Section Cover" />
    </div>
    <div class="contact-grid">
      <a
        v-if="data.email"
        id="contact-email"
        href="javascript:void(0)"
        class="contact-card animate-scale-up"
        v-intersect
        @click="showContactModal = true"
      >
        <Mail :size="24" class="contact-icon" />
        <span class="contact-label">Email</span>
        <span class="contact-val">{{ data.email }}</span>
        <div class="contact-hint">Click to view contact</div>
      </a>

      <!-- Premium Contact Modal -->
      <Teleport to="body">
        <Transition name="modal-fade">
          <div v-if="showContactModal" class="contact-modal-overlay" @click.self="showContactModal = false">
            <div class="contact-modal-content">
              <button class="modal-close" @click="showContactModal = false">&times;</button>
              
              <div class="modal-header">
                <div class="modal-logo-container">
                  <div class="logo-inner-glow"></div>
                  <Mail :size="32" class="modal-icon-premium" />
                </div>
                <div class="modal-brand-name">PLOYKONG EXECUTIVE</div>
                <h3 class="gradient-text">{{ data.title || 'Contact Details' }}</h3>
                <p>ช่องทางติดต่อสื่อสารส่วนตัวและข้อมูลสำคัญ</p>
              </div>

              <div class="modal-info-grid">
                <div v-if="data.email" class="info-item">
                  <Mail :size="18" class="info-icon" />
                  <div class="info-content">
                    <label>Email Address</label>
                    <a :href="`mailto:${data.email}`" class="info-link">{{ data.email }}</a>
                  </div>
                </div>
                
                <div v-if="data.linkedin" class="info-item">
                  <Linkedin :size="18" class="info-icon" />
                  <div class="info-content">
                    <label>LinkedIn Professional</label>
                    <a :href="data.linkedin" target="_blank" class="info-link">View Profile</a>
                  </div>
                </div>

                <div v-if="data.github" class="info-item">
                  <Github :size="18" class="info-icon" />
                  <div class="info-content">
                    <label>GitHub Repository</label>
                    <a :href="data.github" target="_blank" class="info-link">View Projects</a>
                  </div>
                </div>

                <div v-if="data.location" class="info-item">
                  <MapPin :size="18" class="info-icon" />
                  <div class="info-content">
                    <label>Current Location</label>
                    <span class="info-text">{{ data.location }}</span>
                  </div>
                </div>

                <!-- Custom Items -->
                <div v-for="contact in data.custom_items" :key="contact.id" class="info-item">
                  <component :is="getIcon(contact.platform)" :size="18" class="info-icon" />
                  <div class="info-content">
                    <label>{{ contact.label }}</label>
                    <a v-if="contact.link" :href="contact.link" target="_blank" class="info-link">{{ contact.value }}</a>
                    <span v-else class="info-text">{{ contact.value }}</span>
                  </div>
                </div>
              </div>
              
              <div class="modal-footer">
                <p>Feel free to reach out via any preferred channel!</p>
              </div>
            </div>
          </div>
        </Transition>
      </Teleport>
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



import { ref, onMounted, onUnmounted } from "vue";

const props = defineProps<{ data: any }>();
const showContactModal = ref(false);

function getIcon(platform: string) {
  switch (platform) {
    case "facebook": return Facebook;
    case "twitter": return Twitter;
    case "instagram": return Instagram;
    case "youtube": return Youtube;
    case "message-circle": return MessageCircle;
    case "phone": return Phone;
    default: return Link;
  }
}

// Global hook for #hire-me hash from Hero
const checkHash = () => {
  if (window.location.hash === '#hire-me') {
    showContactModal.value = true;
    window.location.hash = '#contact'; // Jump to section but show modal
  }
};

onMounted(() => {
  window.addEventListener('hashchange', checkHash);
  checkHash();
});

onUnmounted(() => {
  window.removeEventListener('hashchange', checkHash);
});
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
.contact-grid {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 20px;
}

.layout-left .contact-grid {
  justify-content: flex-start;
}

.contact-card {
  flex: 0 1 280px;
  min-width: 240px;
}

@media (max-width: 640px) {
  .contact-card {
    flex: 1 1 100%;
  }
}

.contact-card {
  display: flex;
  flex-wrap: wrap;
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

.icon-sm {
  vertical-align: -2px;
}

/* Contact Highlight Animation */
@keyframes contact-highlight {
  0% { transform: scale(1); box-shadow: 0 0 0 0 var(--primary); }
  30% { transform: scale(1.15); box-shadow: 0 0 30px 15px var(--primary-glow); border-color: var(--primary); }
  60% { transform: scale(1.05); box-shadow: 0 0 20px 10px var(--primary-glow); }
  100% { transform: scale(1); box-shadow: 0 0 0 0 transparent; }
}

.contact-card:target {
  animation: contact-highlight 1.2s ease-out 2;
  border-color: var(--primary) !important;
  z-index: 100;
  background: var(--primary-glow) !important;
}

.contact-hint {
  font-size: 10px;
  opacity: 0.5;
  margin-top: 4px;
}

/* --- Premium Modal Global --- */
:global(.contact-modal-overlay) {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.75);
  backdrop-filter: blur(20px) saturate(180%);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 99999;
  padding: 24px;
}

:global(.contact-modal-content) {
  background: linear-gradient(145deg, rgba(20, 20, 35, 0.9), rgba(10, 10, 18, 0.95));
  backdrop-filter: blur(30px);
  width: 100%;
  max-width: 480px;
  border-radius: 40px;
  border: 1px solid rgba(255, 255, 255, 0.08);
  padding: 48px;
  position: relative;
  box-shadow: 
    0 50px 100px -20px rgba(0, 0, 0, 0.8),
    inset 0 0 0 1px rgba(255, 255, 255, 0.03);
  color: white;
}

:global(.theme-light .contact-modal-content) {
  background: linear-gradient(165deg, #ffffff 0%, #fcfdfe 100%);
  color: #0f172a;
  border-color: rgba(15, 23, 42, 0.08);
  box-shadow: 
    0 10px 30px -10px rgba(0, 0, 0, 0.05),
    0 40px 120px -20px rgba(15, 23, 42, 0.12),
    inset 0 0 0 1px rgba(255, 255, 255, 0.82);
}

:global(.modal-close) {
  position: absolute;
  top: 32px;
  right: 32px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.08);
  width: 38px;
  height: 38px;
  border-radius: 50%;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  font-size: 18px;
}

:global(.theme-light .modal-close) {
  background: rgba(15, 23, 42, 0.05);
  border-color: rgba(15, 23, 42, 0.1);
  color: #1e293b;
}

:global(.modal-close:hover) {
  background: var(--primary);
  border-color: var(--primary);
  color: white;
  transform: rotate(90deg) scale(1.1);
  box-shadow: 0 0 20px var(--primary-glow);
}

:global(.modal-header) {
  text-align: center;
  margin-bottom: 40px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

:global(.modal-logo-container) {
  width: 88px;
  height: 88px;
  background: linear-gradient(135deg, var(--primary) 0%, color-mix(in srgb, var(--primary), #000 25%) 100%);
  border-radius: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 28px;
  position: relative;
  box-shadow: 
    0 25px 50px -12px color-mix(in srgb, var(--primary) 60%, transparent),
    inset 0 0 20px rgba(255, 255, 255, 0.15);
  border: 4px solid rgba(255, 255, 255, 0.08); /* Outer subtle ring */
}

:global(.modal-logo-container::after) {
  content: '';
  position: absolute;
  inset: -6px;
  border-radius: 34px;
  border: 1px solid color-mix(in srgb, var(--primary) 30%, transparent);
  pointer-events: none;
}

:global(.logo-inner-glow) {
  position: absolute;
  inset: 2px;
  border-radius: inherit;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.2), transparent);
  opacity: 0.5;
}

:global(.modal-icon-premium) {
  color: white;
  z-index: 1;
  filter: drop-shadow(0 4px 12px rgba(0,0,0,0.3));
  transform: scale(1.1);
}

:global(.modal-brand-name) {
  font-size: 11px;
  font-weight: 900;
  color: var(--primary);
  letter-spacing: 3px;
  margin-bottom: 8px;
  text-transform: uppercase;
  text-shadow: 0 0 15px color-mix(in srgb, var(--primary) 30%, transparent);
}

:global(.gradient-text) {
  font-size: 34px;
  font-weight: 900;
  background: linear-gradient(to bottom, #ffffff 10%, rgba(255,255,255,0.5) 100%);
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
  margin-bottom: 10px;
  letter-spacing: -1px;
  filter: drop-shadow(0 2px 4px rgba(0,0,0,0.1));
}

:global(.theme-light .gradient-text) {
  background: linear-gradient(to bottom, #0f172a 30%, #475569);
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
}

:global(.modal-header p) {
  font-size: 15px;
  color: rgba(255, 255, 255, 0.6);
  max-width: 340px;
  line-height: 1.5;
}

:global(.theme-light .modal-header p) {
  color: #64748b;
  font-weight: 500;
  opacity: 0.9;
}

:global(.modal-info-grid) {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

:global(.info-item) {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 18px 24px;
  background: rgba(255, 255, 255, 0.02);
  border: 1px solid rgba(255, 255, 255, 0.04);
  border-radius: 24px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

:global(.theme-light .info-item) {
  background: rgba(241, 245, 249, 0.35);
  border-color: rgba(15, 23, 42, 0.06);
}

:global(.theme-light .info-item:hover) {
  background: #ffffff;
  border-color: var(--primary);
  transform: translateY(-2px) translateX(6px);
  box-shadow: 0 15px 35px -5px rgba(15, 23, 42, 0.08);
}

:global(.info-icon) {
  color: var(--primary);
  opacity: 0.9;
}

:global(.info-content label) {
  display: block;
  font-size: 10px;
  font-weight: 800;
  text-transform: uppercase;
  letter-spacing: 1px;
  color: var(--primary);
  margin-bottom: 3px;
  opacity: 0.7;
}

:global(.info-link), :global(.info-text) {
  font-size: 15px;
  font-weight: 700;
  text-decoration: none;
  color: inherit;
  letter-spacing: -0.2px;
}

:global(.info-link) {
  color: var(--primary);
  transition: opacity 0.2s;
}

:global(.info-link:hover) {
  opacity: 0.8;
}

:global(.modal-footer) {
  margin-top: 32px;
  text-align: center;
  font-size: 12px;
  color: rgba(255, 255, 255, 0.4);
  font-style: italic;
  font-weight: 500;
}

:global(.theme-light .modal-footer) {
  color: #94a3b8;
}

/* Modal Transitions */
:global(.modal-fade-enter-active) { transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1); }
:global(.modal-fade-leave-active) { transition: all 0.2s ease; }
:global(.modal-fade-enter-from) { opacity: 0; transform: scale(0.95); }
:global(.modal-fade-leave-to) { opacity: 0; transform: scale(0.98); }
</style>
