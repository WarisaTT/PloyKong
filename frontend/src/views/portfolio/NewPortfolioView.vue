<template>
  <div class="new-portfolio-page">
    <div class="new-bg"></div>
    <div class="new-card">
      <!-- Step indicator -->
      <div class="steps-indicator">
        <div v-for="i in 3" :key="i" class="step-dot" :class="{ active: step >= i, done: step > i }"></div>
      </div>

      <!-- Step 1: Role -->
      <div v-if="step === 1" class="step-content animate-fade-in">
        <h2>🎭 คุณเป็นสายไหน?</h2>
        <p>เลือกสายงานเพื่อให้ระบบแนะนำ Template ที่เหมาะสม</p>
        <div class="role-grid">
          <div v-for="r in roles" :key="r.value" class="role-card" :class="{ selected: form.role === r.value }" @click="form.role = r.value">
            <span class="role-icon">{{ r.icon }}</span>
            <span class="role-name">{{ r.label }}</span>
          </div>
        </div>
        <button class="btn btn-primary" style="width:100%;margin-top:20px" @click="step++" :disabled="!form.role">ถัดไป →</button>
      </div>

      <!-- Step 2: Info -->
      <div v-if="step === 2" class="step-content animate-fade-in">
        <h2>✍️ ข้อมูลพื้นฐาน</h2>
        <div class="form-field">
          <label class="form-label">ชื่อ Portfolio</label>
          <input v-model="form.title" class="form-input" placeholder="My Awesome Portfolio" />
        </div>
        <div class="form-field">
          <label class="form-label">URL ของคุณ</label>
          <div class="slug-row">
            <span class="slug-prefix">pk.io/</span>
            <input v-model="form.slug" class="form-input slug-input" placeholder="yourname" @input="slugError = ''" />
          </div>
          <div v-if="slugError" class="field-error">{{ slugError }}</div>
        </div>
        <div style="display:flex;gap:10px;margin-top:20px">
          <button class="btn btn-secondary" @click="step--">← ย้อนกลับ</button>
          <button class="btn btn-primary" style="flex:1" @click="validateSlug">ถัดไป →</button>
        </div>
      </div>

      <!-- Step 3: Theme -->
      <div v-if="step === 3" class="step-content animate-fade-in">
        <h2>🎨 เลือก Theme</h2>
        <div class="theme-grid">
          <div v-for="t in themes" :key="t.id" class="theme-card" :class="{ selected: form.theme === t.id }" @click="form.theme = t.id" :style="t.style">
            <div class="theme-preview-bar" :style="{ background: t.accent }"></div>
            <span class="theme-name">{{ t.name }}</span>
          </div>
        </div>
        <div style="display:flex;gap:10px;margin-top:20px">
          <button class="btn btn-secondary" @click="step--">← ย้อนกลับ</button>
          <button class="btn btn-primary" style="flex:1" @click="createPortfolio" :disabled="creating">
            {{ creating ? '⏳ กำลังสร้าง...' : '🚀 สร้างพอร์ตเลย!' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { usePortfolioStore } from '@/stores/portfolio'

const router = useRouter()
const store = usePortfolioStore()

const step = ref(1)
const slugError = ref('')
const creating = ref(false)

const form = reactive({ role: '', title: '', slug: '', theme: 'neon-dark' })

const roles = [
  { value: 'programmer', icon: '👨‍💻', label: 'Programmer' },
  { value: 'designer', icon: '🎨', label: 'Designer' },
  { value: 'first-jobber', icon: '🌱', label: 'First Jobber' },
  { value: 'business', icon: '💼', label: 'Business' },
  { value: 'creative', icon: '🎭', label: 'Creative' },
  { value: 'custom', icon: '⚡', label: 'Other' },
]

const themes = [
  { id: 'neon-dark', name: '🌙 Neon Dark', accent: 'linear-gradient(135deg,#4F46E5,#a855f7)', style: 'background:rgba(5,8,20,0.8);border-color:rgba(79,70,229,0.5)' },
  { id: 'clean-light', name: '☀️ Clean Light', accent: 'linear-gradient(135deg,#6366f1,#06b6d4)', style: 'background:rgba(240,245,255,0.1);border-color:rgba(99,102,241,0.4)' },
  { id: 'tokyo-night', name: '🏙️ Tokyo Night', accent: 'linear-gradient(135deg,#7c3aed,#ec4899)', style: 'background:rgba(10,5,25,0.8);border-color:rgba(124,58,237,0.5)' },
  { id: 'ocean-cyber', name: '🌊 Ocean Cyber', accent: 'linear-gradient(135deg,#06b6d4,#10b981)', style: 'background:rgba(5,15,25,0.8);border-color:rgba(6,182,212,0.5)' },
]

function validateSlug() {
  if (!form.slug) { slugError.value = 'กรุณาใส่ URL'; return }
  if (!/^[a-z0-9-]{3,30}$/.test(form.slug)) {
    slugError.value = 'ใช้ได้เฉพาะ a-z, 0-9, - (3-30 ตัว)'
    return
  }
  step.value++
}

async function createPortfolio() {
  creating.value = true
  const themeConfig = {
    mode: form.theme === 'clean-light' ? 'light' : 'dark',
    primary_color: { 'neon-dark': '#4F46E5', 'clean-light': '#6366f1', 'tokyo-night': '#7c3aed', 'ocean-cyber': '#06b6d4' }[form.theme] || '#4F46E5',
    font: 'Syne',
    layout: 'centered',
  }
  const result = await store.createPortfolio({ slug: form.slug, title: form.title || 'My Portfolio', theme: themeConfig })
  creating.value = false
  if (result) router.push(`/portfolios/${result.id}/edit`)
}
</script>

<style scoped>
.new-portfolio-page { min-height:100vh;display:flex;align-items:center;justify-content:center;padding:24px;position:relative; }
.new-bg { position:fixed;inset:0;background:radial-gradient(ellipse 60% 50% at 50% 30%,rgba(79,70,229,0.2) 0%,transparent 60%); }
.new-card { position:relative;z-index:1;background:rgba(10,14,30,0.95);border:1px solid var(--border);border-radius:24px;padding:44px;width:100%;max-width:480px;backdrop-filter:blur(30px); }
.steps-indicator { display:flex;gap:8px;justify-content:center;margin-bottom:32px; }
.step-dot { width:8px;height:8px;border-radius:50%;background:var(--border);transition:all 0.3s; }
.step-dot.active { background:var(--purple);transform:scale(1.3); }
.step-dot.done { background:var(--success); }
.step-content h2 { font-size:22px;font-weight:800;margin-bottom:8px; }
.step-content > p { color:var(--muted);font-size:14px;margin-bottom:24px; }
.role-grid { display:grid;grid-template-columns:1fr 1fr 1fr;gap:10px; }
.role-card { display:flex;flex-direction:column;align-items:center;gap:8px;padding:16px 8px;background:rgba(255,255,255,0.03);border:1px solid var(--border);border-radius:12px;cursor:pointer;transition:all 0.15s;text-align:center; }
.role-card:hover { border-color:rgba(168,85,247,0.4);background:rgba(168,85,247,0.06); }
.role-card.selected { border-color:var(--purple);background:rgba(168,85,247,0.12); }
.role-icon { font-size:24px; }
.role-name { font-size:12px;font-weight:600; }
.form-field { display:flex;flex-direction:column;margin-bottom:14px; }
.slug-row { display:flex;align-items:center;background:rgba(255,255,255,0.04);border:1px solid var(--border);border-radius:10px;overflow:hidden; }
.slug-prefix { padding:0 12px;color:var(--muted);font-size:13px;font-weight:600;border-right:1px solid var(--border);height:100%;display:flex;align-items:center;flex-shrink:0; }
.slug-input { flex:1;border:none;background:transparent;padding:11px 12px; }
.field-error { font-size:12px;color:var(--danger);margin-top:4px; }
.theme-grid { display:grid;grid-template-columns:1fr 1fr;gap:12px; }
.theme-card { padding:16px;border:2px solid rgba(255,255,255,0.08);border-radius:12px;cursor:pointer;transition:all 0.15s; }
.theme-card.selected { border-color:var(--purple); }
.theme-preview-bar { height:4px;border-radius:2px;margin-bottom:10px; }
.theme-name { font-size:13px;font-weight:600; }
</style>
