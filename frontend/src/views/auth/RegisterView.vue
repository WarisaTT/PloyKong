<template>
  <div class="auth-page">

    <!-- ══ LEFT: Brand Panel ════════════════════════════════ -->
    <aside class="brand-panel">
      <div class="bp-mesh">
        <div class="bp-orb bp-orb--1"></div>
        <div class="bp-orb bp-orb--2"></div>
        <div class="bp-orb bp-orb--3"></div>
        <div class="bp-grid"></div>
      </div>

      <!-- Big BG text -->
      <div class="bp-watermark" aria-hidden="true">PK</div>

      <div class="bp-inner">
        <!-- Badge top -->
        <div class="bp-badge-top">✦ เริ่มต้นได้ฟรี · ไม่ต้องใส่บัตรเครดิต</div>

        <!-- Headline -->
        <div class="bp-headline">
          <h1 class="bp-h1">
            พอร์ตสวย<br/>
            <span class="bp-accent">ใน 5 นาที</span><br/>
            จริงๆ
          </h1>
          <p class="bp-body">
            ผู้สมัครที่มีเว็บพอร์ตโฟลิโอได้รับการตอบรับจาก HR<br/>
            มากกว่าถึง <strong style="color: var(--text)">3.4 เท่า</strong>
          </p>
        </div>

        <!-- Steps -->
        <div class="bp-steps">
          <div v-for="(step, i) in steps" :key="i" class="bp-step">
            <div class="bp-step-num" :style="{ background: step.color }">{{ i + 1 }}</div>
            <div>
              <div class="bp-step-title">{{ step.title }}</div>
              <div class="bp-step-desc">{{ step.desc }}</div>
            </div>
          </div>
        </div>
      </div>

      <!-- Stats -->
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

      <!-- Theme toggle -->
      <div class="fp-theme">
        <ThemeToggle />
      </div>

      <div class="auth-card">
        <!-- Card logo -->
        <div class="card-logo">
          <div class="logo-mark">
            <img src="/favicon.png" alt="PloyKong Logo" class="logo-img" />
          </div>
          <div class="logo-name">PloyKong</div>
        </div>

        <h2 class="ac-title">สร้างบัญชีฟรี</h2>
        <p class="ac-sub">เริ่มสร้างพอร์ตโฟลิโอแรกของคุณในวันนี้ ✦</p>

        <div id="google-btn-wrapper" class="google-btn-container"></div>

        <div class="ac-divider"><span>หรือกรอกข้อมูล</span></div>

        <form @submit.prevent="handleRegister" class="ac-form">

          <div class="field">
            <label class="field-lbl">ชื่อ-นามสกุล</label>
            <div class="field-wrap">
              <User class="field-ico" :size="15" />
              <input v-model="form.name" type="text" class="field-inp" placeholder="Ploy Kong" required />
            </div>
          </div>

          <div class="field">
            <label class="field-lbl">Email Address</label>
            <div class="field-wrap">
              <Mail class="field-ico" :size="15" />
              <input v-model="form.email" type="email" class="field-inp" placeholder="ploykong@example.com" autocomplete="email" required />
            </div>
          </div>

          <div class="field">
            <label class="field-lbl">Password</label>
            <div class="field-wrap">
              <Lock class="field-ico" :size="15" />
              <input v-model="form.password" :type="showPw ? 'text' : 'password'" class="field-inp" placeholder="อย่างน้อย 8 ตัวอักษร" minlength="8" required />
              <button type="button" class="field-eye" @click="showPw = !showPw">
                <Eye v-if="!showPw" :size="15" />
                <EyeOff v-else :size="15" />
              </button>
            </div>
            <!-- Password strength -->
            <div class="pw-strength">
              <div class="pw-bars">
                <div class="pw-bar" v-for="i in 4" :key="i" :class="{ active: i <= pwStrength }"></div>
              </div>
              <div class="pw-lbl">{{ pwLabel }}</div>
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
              <Check v-else :size="16" />
              {{ authStore.loading ? 'กำลังสมัคร...' : 'สมัครสมาชิกฟรี' }}
            </span>
          </button>
        </form>

        <p class="ac-switch">
          มีบัญชีแล้ว?
          <RouterLink to="/login" class="ac-link">← เข้าสู่ระบบ</RouterLink>
        </p>
      </div>
    </main>
  </div>
</template>

<script setup>
import { reactive, ref, computed, onMounted } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import ThemeToggle from '@/components/ThemeToggle.vue'
import { AlertTriangle, Loader2, User, Mail, Lock, Check, Eye, EyeOff } from 'lucide-vue-next'

const authStore = useAuthStore()
const router    = useRouter()
const form      = reactive({ name: '', email: '', password: '' })
const showPw    = ref(false)

onMounted(() => {
  if (window.google && import.meta.env.VITE_GOOGLE_CLIENT_ID) {
    window.google.accounts.id.initialize({
      client_id: import.meta.env.VITE_GOOGLE_CLIENT_ID,
      callback: handleGoogleCallback,
    })
    
    window.google.accounts.id.renderButton(
      document.getElementById('google-btn-wrapper'),
      { 
        theme: 'outline', 
        size: 'large', 
        width: '400', 
        shape: 'circle',
        logo_alignment: 'left',
        text: 'signup_with'
      }
    )
  }
})

async function handleGoogleCallback(response) {
  const ok = await authStore.loginWithGoogle(response.credential)
  if (ok) router.push('/dashboard')
}

// Password strength
const pwStrength = computed(() => {
  const pw = form.password
  if (!pw) return 0
  let s = 0
  if (pw.length >= 8) s++
  if (/[A-Z]/.test(pw)) s++
  if (/[0-9]/.test(pw)) s++
  if (/[^A-Za-z0-9]/.test(pw)) s++
  return s
})
const pwLabel = computed(() => {
  return ['', 'อ่อนมาก', 'อ่อน', 'ปานกลาง', 'แข็งแกร่ง'][pwStrength.value] || ''
})

const steps = [
  { color: 'linear-gradient(135deg,#7C6FE0,#A78BFA)', title: 'สมัครฟรี · ใช้ได้ทันที', desc: 'ไม่มีค่าใช้จ่าย ไม่ต้องใส่ข้อมูลบัตร' },
  { color: 'linear-gradient(135deg,#A78BFA,#22D3EE)', title: 'เลือก Template · ลากวาง', desc: '48 template พร้อมใช้ ตามสายงานคุณ' },
  { color: 'linear-gradient(135deg,#22D3EE,#34D399)', title: 'Publish · แชร์ทันที', desc: 'URL สวยงาม ส่งให้ HR ได้เลย' },
]

const stats = [
  { val: '2.4k+', label: 'New This Week' },
  { val: '3.4x',  label: 'Higher Response' },
  { val: '0฿',    label: 'Cost to Start' },
]

async function handleRegister() {
  const success = await authStore.register(form.email, form.password, form.name)
  if (success) router.push('/dashboard')
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Noto+Sans+Thai:wght@400;500;600;700;800&display=swap');

:root {
  --font-display: "Noto Sans Thai", sans-serif;
  --font-body: "Noto Sans Thai", sans-serif;
}

.auth-page {
  display: grid; grid-template-columns: 1.1fr 0.9fr;
  min-height: 100vh;
  font-family: 'Inter', 'Anuphan', system-ui, sans-serif;
  background: var(--bg); color: var(--text);
}
@media (max-width: 900px) {
  .auth-page { grid-template-columns: 1fr; }
  .brand-panel { display: none; }
  .form-panel { padding: 40px 20px; }
}

/* ════ BRAND PANEL ═════════════════════════════════════════ */
.brand-panel {
  position: relative;
  display: flex; flex-direction: column; justify-content: space-between;
  padding: 20px 95px 44px; overflow: hidden;
  background: var(--bg-2); border-right: 1px solid var(--border);
}

.bp-mesh { position: absolute; inset: 0; pointer-events: none; }
.bp-orb { position: absolute; border-radius: 50%; filter: blur(80px); }
.bp-orb--1 {
  width: 500px; height: 500px;
  background: radial-gradient(circle, var(--accent-glow), transparent 70%);
  top: -100px; left: -100px;
}
.bp-orb--2 {
  width: 400px; height: 400px;
  background: radial-gradient(circle, rgba(232,121,249,0.1), transparent 70%);
  bottom: -80px; right: -60px;
}
.bp-orb--3 {
  width: 300px; height: 300px;
  background: radial-gradient(circle, rgba(34,211,238,0.07), transparent 70%);
  top: 40%; right: 10%;
}

.bp-grid {
  position: absolute; inset: 0;
  background-image:
    linear-gradient(var(--border) 1px, transparent 1px),
    linear-gradient(90deg, var(--border) 1px, transparent 1px);
  background-size: 40px 40px;
}

.bp-watermark {
  position: absolute; pointer-events: none; user-select: none;
  font-family: 'Syne', sans-serif; font-weight: 900;
  font-size: clamp(200px, 22vw, 320px); line-height: 1;
  top: 50%; left: -20px; transform: translateY(-50%);
  color: var(--accent); opacity: 0.04; letter-spacing: -10px;
}

.bp-inner {
  position: relative; z-index: 2;
  flex: 1; display: flex; flex-direction: column; justify-content: center; gap: 32px;
}

.bp-badge-top {
  display: inline-flex; align-items: center; gap: 8px; width: fit-content;
  padding: 6px 16px; border-radius: 100px;
  background: var(--accent-subtle); border: 1px solid var(--accent-glow);
  font-size: 11px; font-weight: 700; color: var(--accent-2);
  letter-spacing: 1px; text-transform: uppercase;
}

.bp-h1 {
  font-family: 'Syne', sans-serif;
  font-size: clamp(38px, 4vw, 56px); font-weight: 900; letter-spacing: -1.5px;
  line-height: 1.05; color: var(--text);
}
.bp-accent {
  background: var(--grad-pk);
  -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text;
}
.bp-body { font-size: 15px; line-height: 1.8; color: var(--text-2); font-weight: 300; }

.bp-steps { display: flex; flex-direction: column; gap: 20px; }
.bp-step { display: flex; align-items: flex-start; gap: 16px; }
.bp-step-num {
  width: 32px; height: 32px; border-radius: 50%;
  display: flex; align-items: center; justify-content: center;
  font-family: 'Syne', sans-serif; font-size: 13px; font-weight: 800;
  color: #fff; flex-shrink: 0; box-shadow: 0 4px 14px var(--accent-glow);
}
.bp-step-title { font-size: 14px; font-weight: 600; color: var(--text); margin-bottom: 2px; }
.bp-step-desc { font-size: 12.5px; color: var(--text-2); }

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
  pointer-events: none; opacity: 0.4;
}

.fp-theme { position: absolute; top: 20px; right: 20px; z-index: 10; }

.auth-card {
  position: relative; z-index: 2; width: 100%; max-width: 400px;
  animation: cardIn 0.5s cubic-bezier(0.16,1,0.3,1) both;
}
@keyframes cardIn { from { opacity:0; transform:translateY(18px) } to { opacity:1; transform:translateY(0) } }

.card-logo { display: flex; align-items: center; gap: 10px; margin-bottom: 24px; }
.logo-mark {
  width: 36px; height: 36px; border-radius: 10px;
  background: var(--surface-2); border: 1px solid var(--border-2);
  display: flex; align-items: center; justify-content: center;
  box-shadow: 0 0 16px var(--accent-glow);
}
.logo-name { font-family: 'Syne', sans-serif; font-size: 18px; font-weight: 700; color: var(--text); }

.ac-sub { font-size: 14px; color: var(--text-2); margin-bottom: 24px; font-weight: 300; }

@media (max-width: 480px) {
  .ac-title { font-size: 26px; }
  .auth-card { padding: 0 10px; }
}

.ac-form { display: flex; flex-direction: column; gap: 16px; }
.field { display: flex; flex-direction: column; gap: 7px; }
.field-lbl {
  font-size: 11px; font-weight: 700; letter-spacing: 0.7px; text-transform: uppercase;
  color: var(--text-2);
}
.field-wrap { position: relative; }
.field-ico {
  position: absolute; left: 14px; top: 50%; transform: translateY(-50%);
  color: var(--text-3); pointer-events: none; transition: color 0.2s;
}
.field-inp {
  width: 100%; padding: 13px 42px;
  background: var(--surface-2); border: 1px solid var(--border);
  border-radius: 13px; color: var(--text); font-family: inherit; font-size: 14px;
  outline: none; caret-color: var(--accent-2); transition: all 0.2s;
}
.field-inp::placeholder { color: var(--text-3); font-weight: 300; }
.field-inp:focus {
  border-color: var(--accent); background: var(--accent-subtle);
  box-shadow: 0 0 0 3px var(--accent-glow);
}
.field-wrap:focus-within .field-ico { color: var(--accent-2); }
.field-eye {
  position: absolute; right: 13px; top: 50%; transform: translateY(-50%);
  background: none; border: none; cursor: pointer; color: var(--text-3); padding: 3px;
}
.field-eye:hover { color: var(--text); }

/* Password strength */
.pw-strength { margin-top: 4px; }
.pw-bars { display: flex; gap: 4px; margin-bottom: 5px; }
.pw-bar {
  flex: 1; height: 3px; border-radius: 100px; background: var(--surface-3);
  transition: background 0.3s;
}
.pw-bar.active { background: var(--grad-primary); }
.pw-lbl { font-size: 10.5px; color: var(--text-3); }

/* Error */
.ac-error {
  padding: 11px 14px; gap: 8px;
  background: rgba(248,113,113,0.08); border: 1px solid rgba(248,113,113,0.25);
  border-radius: 11px; font-size: 12.5px; color: var(--danger);
  display: flex; align-items: center;
}
.err-fade-enter-active, .err-fade-leave-active { transition: all 0.2s ease; }
.err-fade-enter-from, .err-fade-leave-to { opacity: 0; transform: translateY(-4px); }

/* Divider */
.ac-divider {
  display: flex; align-items: center; gap: 12px; margin: 16px 0;
}
.ac-divider::before, .ac-divider::after { content: ''; flex: 1; height: 1px; background: var(--border); }
.ac-divider span { font-size: 11px; color: var(--text-3); text-transform: uppercase; letter-spacing: 0.5px; }

.google-btn-container {
  display: flex;
  justify-content: center;
  width: 100%;
  margin-bottom: 20px;
  min-height: 40px;
}

/* Submit */
.btn-submit {
  width: 100%; padding: 14px; border: none; border-radius: 13px;
  background: var(--grad-primary); color: #fff;
  font-family: inherit; font-size: 14.5px; font-weight: 700; cursor: pointer;
  box-shadow: 0 6px 24px var(--accent-glow); transition: all 0.22s;
}
.btn-submit:hover:not(:disabled) { transform: translateY(-2px); box-shadow: 0 12px 36px var(--accent-glow); }
.btn-submit:disabled { opacity: 0.5; cursor: not-allowed; }
.btn-inner { display: inline-flex; align-items: center; justify-content: center; gap: 8px; }

/* Switch */
.ac-switch {
  text-align: center; margin-top: 24px; font-size: 13.5px; color: var(--text-2); font-weight: 300;
}
.ac-link { color: var(--accent-2); text-decoration: none; font-weight: 700; }
.ac-link:hover { color: var(--accent); }

@keyframes spin { to { transform: rotate(360deg); } }
.spin { animation: spin 1s linear infinite; }
</style>
