<template>
  <div class="auth-page">

    <!-- ══ LEFT: Brand Panel ════════════════════════════════ -->
    <aside class="brand-panel">
      <div class="bp-mesh">
        <div class="bp-orb bp-orb--1"></div>
        <div class="bp-orb bp-orb--2"></div>
        <div class="bp-grid"></div>
      </div>

      <div class="bp-inner">
        <!-- Logo -->
        <div class="bp-logo">
          <div class="bp-logo-mark">
            <img src="/favicon.png" alt="PloyKong Logo" class="logo-img" />
          </div>
          <span class="bp-logo-name">Ploy<em>Kong</em></span>
        </div>

        <!-- Headline -->
        <div class="bp-headline">
          <p class="bp-eyebrow">No-Code · AI-Powered · Gen-Z</p>
          <h1 class="bp-h1">
            มีดีต้อง<br/>
            <span class="bp-accent">ปล่อยของ</span>
          </h1>
          <p class="bp-body">สร้างเว็บ Portfolio ระดับมืออาชีพ<br/>ด้วย AI ในไม่กี่คลิก — ไม่ต้องรู้โค้ด</p>
        </div>

        <!-- Features -->
        <ul class="bp-feats">
          <li v-for="f in features" :key="f.label" class="bp-feat">
            <span class="bp-feat-ico">{{ f.icon }}</span>
            <span>{{ f.label }}</span>
          </li>
        </ul>

        <!-- Live badge -->
        <div class="bp-badge">
          <span class="bp-badge-dot"></span>
          HR จาก Google Thailand กดดูพอร์ตของคุณ เมื่อกี้
        </div>
      </div>

      <!-- Stats strip -->
      <div class="bp-stats">
        <div v-for="s in stats" :key="s.label" class="bp-stat">
          <span class="bp-stat-val">{{ s.val }}</span>
          <span class="bp-stat-label">{{ s.label }}</span>
        </div>
      </div>
    </aside>

    <!-- ══ RIGHT: Form Panel ════════════════════════════════ -->
    <main class="form-panel">
      <div class="fp-glow"></div>

      <!-- Theme toggle top right -->
      <div class="fp-theme">
        <ThemeToggle />
      </div>

      <div class="auth-card">
        <!-- Mobile logo -->
        <div class="mobile-logo">
          <div class="m-logo-mark">
            <img src="/favicon.png" alt="PloyKong Logo" class="logo-img" />
          </div>
          <span class="m-logo-text">PloyKong</span>
        </div>

        <p class="ac-eyebrow">Welcome back</p>
        <h2 class="ac-title">เข้าสู่ระบบ</h2>
        <p class="ac-sub">ยินดีต้อนรับกลับ — พอร์ตคุณรอคุณอยู่ 👋</p>

        <form @submit.prevent="handleLogin" class="ac-form">

          <div class="field">
            <label class="field-lbl">Email Address</label>
            <div class="field-wrap">
              <Mail class="field-ico" :size="15" />
              <input v-model="form.email" type="email" class="field-inp" placeholder="you@example.com" autocomplete="email" required/>
            </div>
          </div>

          <div class="field">
            <div class="field-hdr">
              <label class="field-lbl">Password</label>
              <a href="#" class="field-forgot">ลืมรหัสผ่าน?</a>
            </div>
            <div class="field-wrap">
              <Lock :size="15" class="field-ico" />
              <input v-model="form.password" :type="showPw ? 'text' : 'password'" class="field-inp" placeholder="••••••••" autocomplete="current-password" required/>
              <button type="button" class="field-eye" @click="showPw = !showPw">
                <Eye v-if="!showPw" :size="15" />
                <EyeOff v-else :size="15" />
              </button>
            </div>
          </div>

          <Transition name="err-fade">
            <div v-if="authStore.error" class="ac-error">
              <AlertTriangle :size="14" />{{ authStore.error }}
            </div>
          </Transition>

          <button type="submit" class="btn-submit" :disabled="authStore.loading">
            <span class="btn-inner">
              <Loader2 v-if="authStore.loading" :size="16" class="spin" />
              <Rocket v-else :size="16" />
              {{ authStore.loading ? 'กำลังเข้าสู่ระบบ...' : 'เข้าสู่ระบบ' }}
            </span>
          </button>
        </form>

        <div class="ac-divider"><span>หรือ</span></div>

        <button class="btn-google" type="button">
          <svg width="17" height="17" viewBox="0 0 24 24"><path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"/><path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/><path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05"/><path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/></svg>
          ดำเนินการต่อด้วย Google
        </button>

        <p class="ac-switch">
          ยังไม่มีบัญชี?
          <RouterLink to="/register" class="ac-link">สมัครสมาชิกฟรี →</RouterLink>
        </p>
      </div>
    </main>

  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useThemeStore } from '@/stores/theme'
import ThemeToggle from '@/components/ThemeToggle.vue'
import { AlertTriangle, Loader2, Rocket, Mail, Lock, Eye, EyeOff } from 'lucide-vue-next'

const authStore = useAuthStore()
const router    = useRouter()
const form      = reactive({ email: '', password: '' })
const showPw    = ref(false)

const features = [
  { icon: '🎨', label: 'Drag & Drop Builder — ลากวางได้เลย' },
  { icon: '🤖', label: 'AI เกลาเนื้อหาให้เป็นมืออาชีพ' },
  { icon: '📊', label: 'Analytics + Real-time Notifications' },
  { icon: '🔒', label: 'Password-Protected Sections' },
]
const stats = [
  { val: '48',    label: 'Templates' },
  { val: '<5min', label: 'Build Time' },
  { val: '∞',     label: 'Profiles' },
  { val: 'Free',  label: 'Forever' },
]

async function handleLogin() {
  const ok = await authStore.login(form.email, form.password)
  if (ok) router.push('/dashboard')
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Noto+Sans+Thai:wght@400;500;600;700;800&display=swap');

:root {
  --font-display: "Noto Sans Thai", sans-serif;
  --font-body: "Noto Sans Thai", sans-serif;
}

/* ── Page layout ─────────────────────────────────────────── */
.auth-page {
  display: grid;
  grid-template-columns: 1fr 1fr;
  min-height: 100vh;
  font-family: 'Inter', 'Anuphan', system-ui, sans-serif;
  background: var(--bg);
  color: var(--text);
}
@media (max-width: 860px) {
  .auth-page { grid-template-columns: 1fr; }
  .brand-panel { display: none; }
}

/* ════ BRAND PANEL ═════════════════════════════════════════ */
.brand-panel {
  position: relative;
  display: flex; flex-direction: column; justify-content: space-between;
  padding: 1px 95px 78px; overflow: hidden;
  overflow: hidden;
  background: var(--bg-2);
  border-right: 1px solid var(--border);
}

.bp-mesh { position: absolute; inset: 0; pointer-events: none; }
.bp-orb {
  position: absolute; border-radius: 50%; filter: blur(80px);
}
.bp-orb--1 {
  width: 500px; height: 500px;
  background: radial-gradient(circle, var(--accent-glow), transparent 70%);
  top: -120px; left: -100px;
  animation: orbDrift 20s ease-in-out infinite alternate;
}
.bp-orb--2 {
  width: 400px; height: 400px;
  background: radial-gradient(circle, rgba(232,121,249,0.1), transparent 70%);
  bottom: -80px; right: -60px;
  animation: orbDrift 16s ease-in-out infinite alternate-reverse;
}
.bp-grid {
  position: absolute; inset: 0;
  background-image:
    linear-gradient(var(--border) 1px, transparent 1px),
    linear-gradient(90deg, var(--border) 1px, transparent 1px);
  background-size: 40px 40px;
  mask-image: radial-gradient(ellipse 80% 70% at 50% 50%, black, transparent 80%);
}

@keyframes orbDrift {
  from { transform: translate(0,0) scale(1); }
  to   { transform: translate(30px,20px) scale(1.1); }
}

.bp-inner {
  position: relative; z-index: 2;
  flex: 1; display: flex; flex-direction: column; justify-content: center; gap: 36px;
}

/* Logo */
.bp-logo { display: flex; align-items: center; gap: 12px; }
.bp-logo-mark {
  width: 48px; height: 48px; border-radius: 14px;
  background: var(--surface-2); border: 1px solid var(--border-2);
  display: flex; align-items: center; justify-content: center;
  box-shadow: 0 0 24px var(--accent-glow);
}
.bp-logo-name {
  font-family: 'Syne', sans-serif; font-size: 26px; font-weight: 800; letter-spacing: -0.5px;
  color: var(--text);
}
.bp-logo-name em {
  font-style: normal;
  background: var(--grad-pk);
  -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text;
}

/* Headline */
.bp-eyebrow {
  font-size: 10px; font-weight: 700; letter-spacing: 3px; text-transform: uppercase;
  color: var(--accent-2); margin-bottom: 12px;
}
.bp-h1 {
  font-family: 'Syne', sans-serif;
  font-size: clamp(36px, 3.5vw, 50px); font-weight: 900; letter-spacing: -1.5px;
  line-height: 1.05; color: var(--text); margin-bottom: 14px;
}
.bp-accent {
  background: var(--grad-pk);
  -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text;
}
.bp-body { font-size: 15px; line-height: 1.7; color: var(--text-2); font-weight: 300; }

/* Features */
.bp-feats { list-style: none; display: flex; flex-direction: column; gap: 12px; }
.bp-feat {
  display: flex; align-items: center; gap: 12px;
  font-size: 13.5px; color: var(--text-2);
}
.bp-feat-ico {
  width: 34px; height: 34px; border-radius: 10px;
  background: var(--surface-2); border: 1px solid var(--border);
  display: flex; align-items: center; justify-content: center; font-size: 16px;
  flex-shrink: 0;
}

/* Badge */
.bp-badge {
  display: inline-flex; align-items: center; gap: 10px; width: fit-content;
  padding: 9px 16px; border-radius: 100px;
  background: rgba(52,211,153,0.08); border: 1px solid rgba(52,211,153,0.2);
  font-size: 12px; color: var(--success);
  animation: badgeFloat 4s ease-in-out infinite;
}
.bp-badge-dot {
  width: 7px; height: 7px; border-radius: 50%; flex-shrink: 0;
  background: var(--success); box-shadow: 0 0 8px var(--success);
  animation: pulseDot 2s ease-in-out infinite;
}

/* Stats */
.bp-stats {
  position: relative; z-index: 2;
  display: flex; gap: 32px; padding-top: 28px; border-top: 1px solid var(--border);
}
.bp-stat { display: flex; flex-direction: column; gap: 2px; }
.bp-stat-val {
  font-family: 'Syne', sans-serif; font-size: 22px; font-weight: 800;
  background: var(--grad-pk);
  -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text;
}
.bp-stat-label { font-size: 10px; color: var(--text-3); letter-spacing: 0.5px; text-transform: uppercase; }

@keyframes badgeFloat { 0%,100%{ transform:translateY(0) } 50%{ transform:translateY(-5px) } }
@keyframes pulseDot { 0%,100%{ box-shadow:0 0 6px var(--success) } 50%{ box-shadow:0 0 14px var(--success) } }

/* ════ FORM PANEL ══════════════════════════════════════════ */
.form-panel {
  position: relative;
  display: flex; align-items: center; justify-content: center;
  padding: 52px 40px; background: var(--surface); overflow: hidden;
}

.fp-glow {
  position: absolute; width: 500px; height: 500px; border-radius: 50%;
  background: radial-gradient(circle, var(--accent-glow), transparent 65%);
  top: 50%; left: 50%; transform: translate(-50%,-50%);
  pointer-events: none; opacity: 0.5;
}

.fp-theme {
  position: absolute; top: 20px; right: 20px; z-index: 10;
}

.auth-card {
  position: relative; z-index: 2; width: 100%; max-width: 400px;
  animation: cardIn 0.5s cubic-bezier(0.16,1,0.3,1) both;
}
@keyframes cardIn { from { opacity:0; transform:translateY(18px) } to { opacity:1; transform:translateY(0) } }

/* Mobile logo */
.mobile-logo {
  display: none; align-items: center; gap: 10px; margin-bottom: 28px;
}
@media (max-width: 860px) { .mobile-logo { display: flex; } }
.m-logo-mark {
  width: 36px; height: 36px; border-radius: 10px;
  background: var(--grad-primary);
  display: flex; align-items: center; justify-content: center;
  font-family: 'Syne', sans-serif; font-size: 14px; font-weight: 800; color: #fff;
  box-shadow: 0 4px 12px var(--accent-glow);
}
.m-logo-text { font-family: 'Syne', sans-serif; font-size: 18px; font-weight: 700; }

.ac-eyebrow {
  font-size: 10px; font-weight: 700; letter-spacing: 3px; text-transform: uppercase;
  color: var(--accent-2); margin-bottom: 8px;
}
.ac-title { font-family: 'Syne', sans-serif; font-size: 32px; font-weight: 800; letter-spacing: -1px; margin-bottom: 6px; }
.ac-sub { font-size: 14px; color: var(--text-2); margin-bottom: 36px; font-weight: 300; }

/* Form */
.ac-form { display: flex; flex-direction: column; gap: 20px; }
.field { display: flex; flex-direction: column; gap: 7px; }
.field-hdr { display: flex; justify-content: space-between; align-items: center; }
.field-lbl {
  font-size: 11px; font-weight: 700; letter-spacing: 0.7px; text-transform: uppercase;
  color: var(--text-2);
}
.field-forgot {
  font-size: 12px; color: var(--text-2); text-decoration: none; transition: color 0.2s;
}
.field-forgot:hover { color: var(--accent-2); }
.field-wrap { position: relative; }
.field-ico {
  position: absolute; left: 14px; top: 50%; transform: translateY(-50%);
  color: var(--text-3); pointer-events: none; transition: color 0.2s;
}
.field-inp {
  width: 100%; padding: 13px 42px;
  background: var(--surface-2); border: 1px solid var(--border);
  border-radius: 13px; color: var(--text); font-family: inherit; font-size: 14px;
  outline: none; caret-color: var(--accent-2);
  transition: all 0.2s;
}
.field-inp::placeholder { color: var(--text-3); font-weight: 300; }
.field-inp:focus {
  border-color: var(--accent); background: var(--accent-subtle);
  box-shadow: 0 0 0 3px var(--accent-glow);
}
.field-wrap:focus-within .field-ico { color: var(--accent-2); }
.field-eye {
  position: absolute; right: 13px; top: 50%; transform: translateY(-50%);
  background: none; border: none; cursor: pointer; color: var(--text-3);
  padding: 3px; transition: color 0.2s;
}
.field-eye:hover { color: var(--text); }

/* Error */
.ac-error {
  padding: 11px 14px; gap: 8px;
  background: rgba(248,113,113,0.08); border: 1px solid rgba(248,113,113,0.25);
  border-radius: 11px; font-size: 12.5px; color: var(--danger);
  display: flex; align-items: center;
}
.err-fade-enter-active, .err-fade-leave-active { transition: all 0.2s ease; }
.err-fade-enter-from, .err-fade-leave-to { opacity: 0; transform: translateY(-4px); }

/* Submit */
.btn-submit {
  width: 100%; padding: 14px; border: none; border-radius: 13px;
  background: var(--grad-primary); color: #fff;
  font-family: inherit; font-size: 14.5px; font-weight: 700; cursor: pointer;
  box-shadow: 0 6px 24px var(--accent-glow);
  transition: all 0.22s;
}
.btn-submit:hover:not(:disabled) { transform: translateY(-2px); box-shadow: 0 12px 36px var(--accent-glow); }
.btn-submit:active:not(:disabled) { transform: translateY(0); }
.btn-submit:disabled { opacity: 0.5; cursor: not-allowed; }
.btn-inner { display: inline-flex; align-items: center; justify-content: center; gap: 8px; }

/* Divider */
.ac-divider {
  display: flex; align-items: center; gap: 12px; margin: 22px 0;
}
.ac-divider::before, .ac-divider::after { content: ''; flex: 1; height: 1px; background: var(--border); }
.ac-divider span { font-size: 11px; color: var(--text-3); }

/* Google */
.btn-google {
  width: 100%; padding: 12px 20px;
  background: var(--surface-2); border: 1px solid var(--border);
  border-radius: 13px; color: var(--text-2); font-family: inherit;
  font-size: 13.5px; font-weight: 500; cursor: pointer;
  display: flex; align-items: center; justify-content: center; gap: 10px;
  transition: all 0.2s;
}
.btn-google:hover { background: var(--surface-3); border-color: var(--border-2); color: var(--text); }

/* Switch */
.ac-switch {
  text-align: center; margin-top: 28px; font-size: 13.5px; color: var(--text-2); font-weight: 300;
}
.ac-link {
  color: var(--accent-2); text-decoration: none; font-weight: 700; transition: color 0.2s;
}
.ac-link:hover { color: var(--accent); }

@keyframes spin { to { transform: rotate(360deg); } }
.spin { animation: spin 0.9s linear infinite; }
</style>
