<!-- PublicHero.vue - renders the hero section on the public portfolio page -->
<template>
  <section class="pub-hero" :class="'layout-' + (data.layout || 'centered')">
    <div class="pub-hero-inner">
    <img v-if="data.image_url" :src="data.image_url" class="universal-section-img" alt="Cover" />
    <div class="hero-avatar-wrap">
      <img
        v-if="data.avatar_url"
        :src="data.avatar_url"
        :alt="data.name"
        class="hero-img"
      />
      <div v-else class="hero-avatar-fallback">{{ (data.name || "?")[0] }}</div>
      <div class="hero-ring"></div>
    </div>
    <div class="hero-text">
      <div class="hero-role-badge animate-slide-up" v-intersect>
        {{ data.role }}
      </div>

      <h1
        class="hero-name animate-slide-up"
        v-intersect
        style="transition-delay: 100ms"
      >
        {{ data.name }}
      </h1>

      <p
        class="hero-tagline animate-slide-up"
        v-intersect
        style="transition-delay: 200ms"
      >
        {{ data.tagline }}
      </p>

      <div
        v-if="data.show_hire_me || data.show_resume"
        class="hero-cta animate-scale-up"
        v-intersect
        style="transition-delay: 300ms"
      >
        <a
          v-if="data.show_hire_me"
          href="#contact"
          class="hero-btn primary"
          @click="trackHire"
        >
          <Send :size="18" />
          <span>Hire Me</span>
        </a>

        <button v-if="data.show_resume" @click="downloadResume" :disabled="downloadingResume" class="hero-btn secondary">
          <Loader2 v-if="downloadingResume" :size="18" class="spin" />
          <FileText v-else :size="18" />
          <span>{{ downloadingResume ? 'Generating...' : 'Resume/CV' }}</span>
        </button>
      </div>
      </div>
    </div>
    
  </section>
</template>
<script setup lang="ts">
import { ref } from "vue";
import { publicAPI } from "@/api";
import { useRoute } from "vue-router";
import { Send, Mail, FileText, Loader2 } from "lucide-vue-next";
import Swal from "sweetalert2";

const route = useRoute();
defineProps<{ data: any; theme?: any; portfolio?: any }>();
function trackHire() {
  publicAPI.track(route.params.slug as string, "hire_click");
}

const downloadingResume = ref(false);

async function downloadResume() {
  const slug = route.params.slug as string;
  if (!slug) return;
  
  downloadingResume.value = true;
  try {
    // Optionally track the download event specifically here or backend
    const res = await publicAPI.exportPdf(slug);
    
    // Create blob link to download
    const url = window.URL.createObjectURL(new Blob([res.data], { type: "application/pdf" }));
    const link = document.createElement("a");
    link.href = url;
    link.setAttribute("download", `Resume.pdf`);
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
    window.URL.revokeObjectURL(url);
  } catch (error) {
    console.error("Failed to download resume:", error);
    Swal.fire("Error", "ไม่สามารถดาวน์โหลด Resume ได้ในขณะนี้", "error");
  } finally {
    downloadingResume.value = false;
  }
}
</script>
<style scoped>
/* =========================
   HERO LAYOUT
========================= */
.universal-section-img {
  width: 100%;
  height: 250px;
  object-fit: cover;
  border-radius: 16px;
  margin-bottom: 40px;
  box-shadow: 0 10px 30px rgba(0,0,0,0.15);
}

.pub-hero {
  width: 100% !important;
  position: relative;

  display: flex;
  justify-content: center; /* จัดกลางแนวนอน */

  min-height: max-content;
  padding: 40px 0;

  overflow: hidden;

  background:
    radial-gradient(circle at 20% 30%, var(--primary-glow), transparent 55%),
    radial-gradient(circle at 80% 70%, var(--primary-glow), transparent 55%);
}

/* Background glow */
.pub-hero::before {
  content: "";
  position: absolute;
  inset: 0;
}

.pub-hero-inner {
  width: 100%;
  padding: 0 24px;
  max-width: 1200px;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
}

/* =========================
   AVATAR
========================= */
.hero-avatar-wrap {
  position: relative;
  width: 250px;
  height: 250px;
  margin-bottom: 28px;
}

.hero-img,
.hero-avatar-fallback {
  width: 250px;
  height: 250px;
  border-radius: 50%;
}

.hero-img {
  object-fit: cover;
  box-shadow: 0 25px 70px var(--primary-glow);
  border: 4px solid var(--avatar-border, transparent);
}

/* :global(.theme-light) .hero-img,
:global(.theme-light) .hero-avatar-fallback
{
  box-shadow: 0 25px 70px var(--primary-glow);
  border: 4px solid var(--avatar-border, transparent)
} */

.hero-avatar-fallback {
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 56px;
  font-weight: 800;
  color: #fff;
  box-shadow: 0 25px 70px var(--primary-glow);
  border: 4px solid var(--avatar-border, transparent);
}

/* Static soft ring */
.hero-ring {
  position: absolute;
  inset: -4px;
  border-radius: 50%;
  background: transparent;
  box-shadow: 0 0 30px 10px var(--primary-glow);
}

/* =========================
   TEXT
========================= */
.hero-text {
  max-width: 760px;
}

.hero-role-badge {
  display: inline-block;
  padding: 8px 20px;
  border-radius: 999px;
  font-size: 13px;
  font-weight: 600;
  margin-bottom: 16px;
  background: var(--primary-glow);
  border: 1px solid var(--primary);
}

.hero-name {
  font-size: clamp(42px, 6vw, 64px);
  font-weight: 900;
  margin-bottom: 22px;
  line-height: 1.05;

  background: linear-gradient(90deg, var(--primary), var(--secondary));
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
}

.hero-tagline {
  font-size: 18px;
  color: var(--muted);
  max-width: 640px;
  margin: 0 auto 40px; /* จาก 40px เหลือ 32 */
  line-height: 1.8;
}

/* =========================
   CTA BUTTONS
========================= */
.hero-cta {
  display: flex;
  justify-content: center;
  gap: 18px;
  flex-wrap: wrap;
}

.hero-btn {
  min-width: 190px;
  height: 56px;

  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 10px;

  padding: 0 32px;
  border-radius: 999px;

  font-size: 15px;
  font-weight: 600;
  text-decoration: none;

  transition: all 0.25s ease;
}

/* PRIMARY */
.hero-btn.primary {
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  color: #fff;
  box-shadow: 0 18px 45px var(--primary-glow);
}

.hero-btn.primary:hover {
  transform: translateY(-4px);
  box-shadow: 0 28px 70px var(--primary-glow);
}

/* SECONDARY */
.hero-btn.secondary {
  background: transparent;
  border: 1.5px solid var(--primary);
  color: var(--primary);
}

.hero-btn.secondary:hover {
  background: #d7d1dc14;
  transform: translateY(-4px);
}

/* Light mode adjust */
:global(.theme-light) .hero-btn.secondary {
  border: 2px solid var(--secondary);
  color: var(--secondary);
  font-weight: 700;
}

:global(.theme-light) .hero-btn.secondary:hover {
  background: color-mix(in srgb, var(--secondary) 10%, transparent);
  border-color: color-mix(in srgb, var(--secondary) 80%, black);
  color: color-mix(in srgb, var(--secondary) 80%, black);
}

:global(.theme-light) .hero-btn.primary {
  box-shadow: 0 10px 25px color-mix(in srgb, var(--primary) 30%, transparent);
}

:global(.theme-light) .hero-btn.primary:hover {
  box-shadow: 0 15px 35px color-mix(in srgb, var(--primary) 100%, transparent);
}

/* =========================
   LAYOUT MODIFIERS
========================= */

/* LEFT ALIGNED */
.pub-hero.layout-left .pub-hero-inner {
  align-items: flex-start;
  text-align: left;
}

.pub-hero.layout-left .hero-tagline {
  margin: 0 0 32px 0;
}

.pub-hero.layout-left .hero-cta {
  justify-content: flex-start;
}

/* SPLIT / TWO-COLUMN */
.pub-hero.layout-split .pub-hero-inner {
  flex-direction: row-reverse;
  justify-content: space-between;
  text-align: left;
  gap: 60px;
  align-items: center;
}

.pub-hero.layout-split .hero-text {
  flex: 1;
  text-align: left;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.pub-hero.layout-split .hero-tagline {
  margin: 0 0 32px 0;
  max-width: 100%;
}

.pub-hero.layout-split .hero-cta {
  justify-content: flex-start;
}

@media (max-width: 768px) {
  .pub-hero.layout-split .pub-hero-inner {
    flex-direction: column;
    text-align: center;
    gap: 30px;
  }
  .pub-hero.layout-split .hero-text {
    align-items: center;
    text-align: center;
  }
  .pub-hero.layout-split .hero-tagline {
    margin: 0 auto 32px;
  }
  .pub-hero.layout-split .hero-cta {
    justify-content: center;
  }
}
</style>
