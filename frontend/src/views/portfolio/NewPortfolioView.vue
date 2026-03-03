<template>
  <div class="new-portfolio-page" :style="pageBackground">
    <div class="new-card">
      <div class="back-btn-wrap">
        <RouterLink to="/dashboard" class="back-btn" title="Back to Dashboard">
          <ArrowLeft class="icon" :size="20" />
        </RouterLink>
      </div>

      <!-- Step indicator -->
      <div class="steps-indicator">
        <div v-for="i in 3" :key="i" class="step-dot" :class="{ active: step >= i, done: step > i }"></div>
      </div>

      <!-- Mode Toggle (Top right) -->
      <!-- Mode Toggle removed from here, moving down to Step 3 -->

      <!-- STEP 1: Layout Selection -->
      <div v-if="step === 1" class="step-content animate-fade-in">
        <div class="step-header">
          <div class="step-sub">Step 1 of 3</div>
          <h2>คุณเป็นสายไหน?</h2>
          <p>เลือก Layout ที่เหมาะกับสายงานของคุณ</p>
        </div>

        <div class="layout-grid">
          <button v-for="t in TEMPLATES" :key="t.id" @click="form.layout = t.id" class="layout-card"
            :class="{ selected: form.layout === t.id, 'layout-classic': t.id === 'classic' }"
            :style="form.layout === t.id ? { borderColor: '#C084FC', background: 'rgba(123,97,255,0.1)', boxShadow: '0 0 20px rgba(123,97,255,0.25)', color: 'var(--text)' } : { background: 'var(--bg2)', color: 'var(--text)' }">
            <div class="layout-icon"
              :style="form.layout === t.id ? { background: 'linear-gradient(135deg, #C084FC, #F9A8D4)', color: '#fff', boxShadow: '0 0 18px rgba(192,132,252,0.45)' } : { background: 'var(--border)', color: 'var(--muted)' }">
              <component :is="t.icon" :size="t.id === 'classic' ? 24 : 20" />
            </div>
            <div class="layout-info">
              <div class="layout-name">{{ t.label }}</div>
              <div class="layout-desc">{{ t.desc }}</div>
            </div>
            <div v-if="form.layout === t.id" class="selected-badge">✓ Selected</div>
          </button>
        </div>

        <button class="btn-next" @click="step = 2" :disabled="!form.layout" :style="form.layout ? {
          background: 'linear-gradient(135deg, #9333EA, #DB2777)',
          color: '#fff'
        } : {
          background: 'linear-gradient(135deg, #9333EA, #DB2777)',
          color: '#fff'
        }">
          ถัดไป →
        </button>
      </div>

      <!-- STEP 2: Palette Selection -->
      <div v-if="step === 2" class="step-content animate-fade-in">
        <div class="step-header">
          <div class="step-sub">Step 2 of 3</div>
          <h2>เลือก Palette & สี</h2>
          <p>4 palettes · แต่ละ palette มี 6 โทนสีหลัก</p>
        </div>

        <div class="palette-grid">
          <button v-for="p in PALETTES" :key="p.id" @click="selectPalette(p)" class="palette-card"
            :class="{ 'palette-classic': p.id === 'classic', selected: form.palette === p.id }"
            :style="form.palette === p.id ? { borderColor: p.colors[0], background: 'var(--bg2)', boxShadow: `0 0 20px ${p.colors[0]}35`, color: 'var(--text)' } : { background: 'var(--bg2)', color: 'var(--text)' }">
            <div class="palette-header">
              <span class="palette-name">{{ p.label }}</span>
              <span v-if="form.palette === p.id" class="selected-text" :style="{ color: p.colors[0] }">✓ Selected</span>
            </div>
            <div class="color-swatches">
              <div v-for="(c, i) in p.colors" :key="i" class="swatch"
                :style="{ background: c, boxShadow: `0 2px 4px ${c}40` }"></div>
            </div>
          </button>
        </div>

        <div class="step-footer">
          <button class="btn-back" @click="step = 1" title="Back">
            <ArrowLeft :size="20" />
          </button>
          <button class="btn-next" @click="step = 3" :disabled="!form.palette"
            :style="form.layout ? {
              background: 'linear-gradient(135deg, #9333EA, #DB2777)',
              color: '#fff'
            } : {
              background: 'linear-gradient(135deg, #9333EA, #DB2777)',
              color: '#fff'
            }">
            ถัดไป →
          </button>
        </div>
      </div>

      <!-- STEP 3: Info & Preview -->
      <div v-if="step === 3" class="step-content animate-fade-in">
        <div class="step-header">
          <div class="step-sub">Step 3 of 3</div>
          <h2>ตั้งชื่อและ URL ของคุณ</h2>
          <p>รับชม Preview ก่อนสร้างพอร์ต</p>
        </div>

        <div class="info-form">
          <div class="form-field">
            <label class="form-label">ชื่อ Portfolio</label>
            <input v-model="form.title" class="form-input" placeholder="My Awesome Portfolio" />
          </div>
          <div class="form-field">
            <label class="form-label">URL ของคุณ</label>
            <div class="slug-row">
              <span class="slug-prefix">pk.io/p/</span>
              <input v-model="form.slug" class="form-input slug-input" placeholder="yourname" @input="slugError = ''" />
            </div>
            <div v-if="slugError" class="field-error">{{ slugError }}</div>
          </div>
        </div>

        <!-- Live Preview -->
        <div class="preview-box">
          <div style="display: flex; justify-content: space-between; align-items: flex-end; margin-bottom: 10px;">
            <div class="preview-header" style="margin-bottom: 0;">Live Preview: {{ selectedTemplateObj?.label }} ({{
              selectedPaletteObj?.label }})</div>

            <div class="theme-toggle-group" style="scale: 0.85; transform-origin: right bottom;">
              <button :class="['theme-btn', { active: form.mode === 'dark' }]" @click="form.mode = 'dark'"
                title="Dark Mode Preview">
                <Moon :size="14" />
              </button>
              <button :class="['theme-btn', { active: form.mode === 'light' }]" @click="form.mode = 'light'"
                title="Light Mode Preview">
                <Sun :size="14" />
              </button>
            </div>
          </div>
          <div class="preview-wrapper"
            :style="{ borderColor: resolvedTheme.primary + '40', boxShadow: isDark ? `0 0 40px ${resolvedTheme.primary}20` : '0 20px 60px rgba(0,0,0,0.15)' }">
            <TemplatePreviews :layout="form.layout" :theme="resolvedTheme" />
          </div>
        </div>

        <!-- Create Button -->
        <div class="step-footer">
          <button class="btn-back" @click="step = 2" title="Back">
            <ArrowLeft :size="20" />
          </button>
          <button class="btn-next create-btn" @click="validateAndCreate" :disabled="creating"
            :style="{
              background: 'linear-gradient(135deg, #9333EA, #DB2777)',
              color: '#fff'
            }">
            <template v-if="creating">
              <Loader2 :size="16" class="spin icon-inline" /> กำลังสร้าง...
            </template>
            <template v-else>✦ สร้าง Portfolio ของฉัน</template>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, markRaw } from "vue";
import { useRouter } from "vue-router";
import { usePortfolioStore } from "@/stores/portfolio";
import { Code, PenTool, Sprout, Send, Sparkles, Zap, Moon, Sun, Loader2, Layout, ArrowLeft } from "lucide-vue-next";
import TemplatePreviews from "@/components/builder/TemplatePreviews.vue";

const router = useRouter();
const store = usePortfolioStore();

const step = ref(1);
const slugError = ref("");
const creating = ref(false);

const form = reactive({
  layout: "classic",
  palette: "classic",
  primary_color: "#4F46E5",
  mode: "dark",
  title: "",
  slug: "",
});

const isDark = computed(() => form.mode === "dark");

// ─── Constants (Migrated from JSX) ───────────────────────────
const TEMPLATES = [
  { id: "classic", icon: markRaw(Layout), label: "Classic", desc: "รูปแบบดั้งเดิม เรียบง่าย ชัดเจน ใช้งานง่าย" },
  { id: "programmer", icon: markRaw(Code), label: "Programmer", desc: "เน้น Tech Stack / Terminal / Hacker Vibe" },
  { id: "designer", icon: markRaw(PenTool), label: "Designer", desc: "เน้นงาน Visual ภาพใหญ่ๆ ฟอนต์เก๋ๆ" },
  { id: "firstjobber", icon: markRaw(Sprout), label: "First Jobber", desc: "เรียบร้อย คลีนๆ เหมาะยื่นสมัครงาน" },
  { id: "business", icon: markRaw(Send), label: "Business", desc: "เคร่งขรึม เป็นทางการ โชว์ผลประกอบการ" },
  { id: "creative", icon: markRaw(Sparkles), label: "Creative", desc: "ระเบิดไอเดีย สนุกสนาน สีสันฉูดฉาด" },
  { id: "other", icon: markRaw(Zap), label: "Other", desc: "รูปแบบกลางๆ ปรับใช้ได้กับทุกคน" }
];

const PALETTES = [
  { id: "classic", label: "Classic Colors", colors: ["#4F46E5", "#9333EA", "#DB2777", "#06B6D4", "#10B981", "#F97316"] },
  { id: "neon", label: "Neon Cyber", colors: ["#00F5FF", "#7B61FF", "#FF00C8", "#39FF14", "#FFEA00", "#FF3366"] },
  { id: "sunset", label: "Sunset Warm", colors: ["#FF5E3A", "#FF9500", "#FF2A55", "#D35DD8", "#FFB800", "#FF7043"] },
  { id: "forest", label: "Forest Nature", colors: ["#10B981", "#059669", "#34D399", "#A3E635", "#06B6D4", "#3B82F6"] },
  { id: "gold", label: "Luxury Gold", colors: ["#D4AF37", "#FBBF24", "#FCD34D", "#B45309", "#92400E", "#F59E0B"] }
];

// Computed
const selectedTemplateObj = computed(() => TEMPLATES.find(t => t.id === form.layout));
const selectedPaletteObj = computed(() => PALETTES.find(p => p.id === form.palette));

// Using global theme CSS variables (themeStore.mode influences data-theme)
const pageBackground = computed(() => {
  return {}; // Use default CSS var(--bg) defined in style
});

const resolvedTheme = computed(() => {
  const p = form.primary_color || "#7B61FF";
  return {
    isDark: isDark.value,
    bg: isDark.value ? '#0A0E1A' : '#FFFFFF',
    surface: isDark.value ? '#111827' : '#F8FAFC',
    card: isDark.value ? '#1F2937' : '#F1F5F9',
    border: isDark.value ? '#334155' : '#E2E8F0',
    text: isDark.value ? '#FFFFFF' : '#0F172A',
    muted: isDark.value ? '#94A3B8' : '#64748B',
    primary: p,
    secondary: PALETTES[0].colors[1], // mock secondary
  };
});

// Methods
function selectPalette(p: any) {
  form.palette = p.id;
  form.primary_color = p.colors[0];
}

function validateAndCreate() {
  if (!form.slug) {
    slugError.value = "กรุณาใส่ URL";
    return;
  }
  if (!/^[a-z0-9-]{3,30}$/.test(form.slug)) {
    slugError.value = "ใช้ได้เฉพาะ a-z, 0-9, - (3-30 ตัว)";
    return;
  }
  createPortfolio();
}

async function createPortfolio() {
  creating.value = true;
  const themeConfig = {
    mode: form.mode,
    primary_color: form.primary_color,
    palette: form.palette,
    template: form.layout,
    font: form.layout === 'programmer' ? 'Fira Code' : form.layout === 'designer' ? 'Playfair Display' : 'Syne',
    layout: form.layout === 'firstjobber' ? 'left' : form.layout === 'designer' ? 'split' : 'centered',
  };
  const result = await store.createPortfolio({
    slug: form.slug,
    title: form.title || `${selectedTemplateObj.value?.label} Portfolio`,
    theme: themeConfig,
  });
  creating.value = false;
  if (result) router.push(`/portfolios/${result.id}/edit`);
}
</script>

<style scoped>
.new-portfolio-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
  transition: all 0.4s ease;
  font-family: system-ui, sans-serif;
  background: var(--bg);
  color: var(--text);
}

.new-card {
  position: relative;
  z-index: 1;
  width: 100%;
  max-width: 720px;
  margin: 0 auto;
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 24px;
  padding: 44px;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.05);
}

/* Back Button Wrapper (absolute on desktop, relative on mobile) */
.back-btn-wrap {
  position: absolute;
  top: 24px;
  left: 24px;
}

.steps-indicator {
  display: flex;
  gap: 8px;
  justify-content: center;
  margin-bottom: 24px;
}

.back-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  color: var(--muted);
  text-decoration: none;
  border-radius: 8px;
  background: var(--surface);
  border: 1px solid var(--border);
  transition: all 0.2s ease;
}

.back-btn:hover {
  color: var(--text);
  border-color: var(--indigo);
  background: rgba(79, 70, 229, 0.1);
  transform: translateY(-1px);
}

.back-btn:hover .icon {
  transform: translateX(-2px);
  transition: transform 0.2s ease;
}

.step-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: var(--border);
  transition: all 0.3s;
}

.step-dot.active {
  background: #7B61FF;
  transform: scale(1.3);
}

.step-dot.done {
  background: #00F5FF;
}

/* Mode toggle (now near preview) */
.theme-toggle-group {
  display: flex;
  background: var(--border);
  border-radius: 100px;
  padding: 4px;
  gap: 4px;
}

.theme-btn {
  background: transparent;
  border: none;
  border-radius: 100px;
  padding: 6px 12px;
  cursor: pointer;
  color: var(--muted);
  transition: all 0.2s;
}

.theme-btn.active {
  background: var(--surface);
  color: var(--text);
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
}

.step-header {
  text-align: center;
  margin-bottom: 24px;
}

.step-sub {
  color: var(--muted);
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 2px;
  text-transform: uppercase;
  margin-bottom: 6px;
}

.step-header h2 {
  margin: 0;
  font-size: 24px;
  font-weight: 900;
}

.step-header p {
  color: var(--muted);
  font-size: 13px;
  margin-top: 6px;
}

/* Layout Grid */
.layout-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 12px;
  margin-bottom: 24px;
}

.layout-card {
  border: 1.5px solid var(--border);
  border-radius: 16px;
  padding: 20px 12px;
  cursor: pointer;
  transition: all 0.2s;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  text-align: center;
}

.layout-icon {
  width: 44px;
  height: 44px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
  flex-shrink: 0;
}

.layout-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.layout-name {
  font-size: 13px;
  font-weight: 800;
}

.layout-desc {
  font-size: 10px;
  color: var(--muted);
  line-height: 1.4;
}

.layout-card.layout-classic {
  grid-column: 1 / -1;
  flex-direction: row;
  text-align: left;
  padding: 24px;
  align-items: center;
  gap: 16px;
}

.layout-card.layout-classic .layout-icon {
  width: 56px;
  height: 56px;
  border-radius: 16px;
}

.layout-card.layout-classic .layout-name {
  font-size: 16px;
}

.layout-card.layout-classic .layout-desc {
  font-size: 12px;
}

.layout-card.layout-classic .selected-badge {
  margin-top: 0;
  margin-left: auto;
}

.selected-badge {
  background: linear-gradient(135deg, #C084FC, #F9A8D4);
  color: #fff;
  font-size: 9px;
  font-weight: 800;
  letter-spacing: 1px;
  padding: 3px 10px;
  border-radius: 100px;
  margin-top: 4px;
}

/* Palette Grid */
.palette-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
  margin-bottom: 24px;
}

.palette-card {
  border: 1.5px solid var(--border);
  border-radius: 14px;
  padding: 14px;
  cursor: pointer;
  transition: all 0.2s;
  text-align: left;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.palette-card.palette-classic {
  grid-column: 1 / -1;
  flex-direction: row;
  align-items: center;
  padding: 20px 24px;
}

.palette-card.palette-classic .palette-header {
  margin-bottom: 0;
  flex-direction: column;
  align-items: flex-start;
  gap: 4px;
}

.palette-card.palette-classic .color-swatches {
  gap: 8px;
}

.palette-card.palette-classic .swatch {
  width: 32px;
  height: 32px;
}

.palette-card.palette-classic .palette-name {
  font-size: 16px;
}

.palette-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.palette-name {
  font-size: 13px;
  font-weight: 800;
}

.selected-text {
  font-size: 10px;
  font-weight: 800;
}

.color-swatches {
  display: flex;
  gap: 4px;
}

.swatch {
  flex: 1;
  height: 20px;
  border-radius: 4px;
  outline-offset: 1px;
}

/* Color Picker Box */
.color-picker-box {
  border: 1px solid var(--border);
  border-radius: 16px;
  padding: 16px 18px;
  margin-bottom: 24px;
}

.picker-label {
  color: var(--muted);
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 2px;
  text-transform: uppercase;
  margin-bottom: 14px;
}

.picker-swatches {
  display: flex;
  gap: 8px;
}

.picker-swatch {
  flex: 1;
  padding: 12px 0;
  border-radius: 10px;
  cursor: pointer;
  transition: all 0.15s;
  display: flex;
  justify-content: center;
}

.check-mark {
  color: #000;
  font-size: 12px;
  font-weight: 900;
  text-shadow: 0 1px 2px rgba(255, 255, 255, 0.5);
}

/* Buttons */
.step-footer {
  display: flex;
  gap: 12px;
  margin-top: 12px;
}

.step-footer .btn-next {
  width: auto;
  flex: 1;
}

.btn-next {
  width: 100%;
  padding: 14px;
  border-radius: 100px;
  border: none;
  font-size: 15px;
  font-weight: 800;
  cursor: pointer;
  transition: all 0.3s;
}

.btn-next:disabled {
  cursor: not-allowed;
  box-shadow: none !important;
}

.btn-back {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  background: var(--surface);
  border: 1.5px solid var(--border);
  color: var(--muted);
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s;
  flex-shrink: 0;
}

.btn-back:hover {
  background: var(--border);
  color: var(--text);
}

.create-btn {
  color: #000 !important;
}

/* Info Form */
.info-form {
  background: var(--surface);
  border: 1px solid var(--border);
  padding: 20px;
  border-radius: 16px;
  margin-bottom: 24px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
}

.form-field {
  display: flex;
  flex-direction: column;
  margin-bottom: 16px;
}

.form-label {
  font-size: 12px;
  font-weight: 700;
  margin-bottom: 8px;
}

.form-input {
  background: var(--bg);
  border: 1px solid var(--border);
  padding: 12px;
  border-radius: 8px;
  color: var(--text);
  font-size: 14px;
}

.slug-row {
  display: flex;
  align-items: center;
  background: var(--bg);
  border: 1px solid var(--border);
  border-radius: 8px;
  overflow: hidden;
}

.slug-prefix {
  padding: 0 12px;
  color: var(--muted);
  font-size: 13px;
  font-weight: 600;
  border-right: 1px solid var(--border);
  height: 100%;
  display: flex;
  align-items: center;
}

.slug-input {
  flex: 1;
  border: none;
  background: transparent;
  padding: 12px;
}

.field-error {
  font-size: 12px;
  color: #ef4444;
  margin-top: 6px;
}

/* Preview */
.preview-box {
  margin-bottom: 24px;
}

.preview-header {
  color: var(--muted);
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 2px;
  text-transform: uppercase;
  margin-bottom: 10px;
}

.preview-wrapper {
  border-radius: 20px;
  overflow: hidden;
  border: 1px solid;
  transform: scale(1);
  background: var(--bg);
}

.icon-inline {
  vertical-align: middle;
}

@keyframes spin {
  100% {
    transform: rotate(360deg);
  }
}

.spin {
  animation: spin 1s linear infinite;
}

/* ── Mobile Responsive ─────────────────────────────────── */
@media (max-width: 640px) {
  .new-portfolio-page {
    padding: 12px;
    align-items: flex-start;
    padding-top: 20px;
    overflow-y: auto;
  }

  .new-card {
    padding: 24px 20px;
    border-radius: 20px;
    max-width: 100%;
  }

  /* Back button: collapse into flow */
  .back-btn-wrap {
    position: relative;
    top: auto;
    left: auto;
    margin-bottom: 12px;
  }

  .step-header h2 {
    font-size: 20px;
  }

  .step-header p {
    font-size: 12px;
  }

  /* Layout grid: 2 columns on mobile */
  .layout-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 8px;
  }

  /* Classic card stays full-width */
  .layout-card.layout-classic {
    grid-column: 1 / -1;
    padding: 16px;
    gap: 12px;
  }

  .layout-card.layout-classic .layout-icon {
    width: 44px;
    height: 44px;
  }

  .layout-card.layout-classic .layout-name {
    font-size: 14px;
  }

  .layout-card {
    padding: 14px 8px;
    border-radius: 12px;
  }

  .layout-icon {
    width: 36px;
    height: 36px;
    border-radius: 10px;
  }

  .layout-name {
    font-size: 12px;
  }

  .layout-desc {
    font-size: 9px;
  }

  .selected-badge {
    font-size: 8px;
    padding: 2px 7px;
  }

  /* Palette: 1 column on very small screens */
  .palette-grid {
    grid-template-columns: 1fr;
    gap: 10px;
  }

  .palette-card.palette-classic {
    flex-direction: column;
    gap: 12px;
    padding: 16px;
    align-items: flex-start;
  }

  .palette-card.palette-classic .color-swatches {
    width: 100%;
  }

  .palette-card.palette-classic .swatch {
    width: auto;
    height: 24px;
    flex: 1;
  }

  /* Preview: scale down the wrapper */
  .preview-wrapper {
    transform: scale(0.9);
    transform-origin: top center;
    border-radius: 14px;
    margin-bottom: -10%;
  }

  /* Theme toggle in step 3 */
  .theme-toggle-group {
    scale: 0.85;
    transform-origin: right center;
  }

  /* Btn next */
  .btn-next {
    font-size: 14px;
    padding: 12px;
  }

  .btn-back {
    width: 42px;
    height: 42px;
  }

  /* Info form (step 3) */
  .info-form {
    padding: 16px;
  }

  .slug-prefix {
    font-size: 11px;
    padding: 0 8px;
    white-space: nowrap;
  }
}

@media (max-width: 400px) {
  .layout-grid {
    grid-template-columns: 1fr 1fr;
    gap: 6px;
  }

  .new-card {
    padding: 20px 14px;
  }
}
</style>
