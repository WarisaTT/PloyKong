<template>
  <div class="auth-page">

    <!-- ══ LEFT: Brand Panel ════════════════════════════════ -->
    <aside class="brand-panel">
      <div class="bp-orb bp-orb--1"></div>
      <div class="bp-orb bp-orb--2"></div>
      <div class="bp-orb bp-orb--3"></div>
      <div class="bp-grain"></div>
      <div class="bp-dots"></div>

      <div class="bp-inner">
        <!-- Logo -->
        <div class="bp-logo">
          <div class="bp-logo-mark">
            <svg width="26" height="26" viewBox="0 0 26 26" fill="none">
              <path d="M4 3h9a7 7 0 010 14H4V3z" fill="url(#lga)"/>
              <path d="M4 17h13l6 7H4v-7z" fill="url(#lgb)" opacity=".88"/>
              <defs>
                <linearGradient id="lga" x1="4" y1="3" x2="18" y2="17"><stop stop-color="#E8C547"/><stop offset="1" stop-color="#A78BFA"/></linearGradient>
                <linearGradient id="lgb" x1="4" y1="17" x2="23" y2="24"><stop stop-color="#818CF8"/><stop offset="1" stop-color="#38BDF8"/></linearGradient>
              </defs>
            </svg>
          </div>
          <span class="bp-logo-name">Ploy<em>Kong</em></span>
        </div>

        <!-- Headline -->
        <div class="bp-headline">
          <p class="bp-eyebrow">No-Code · AI-Powered · Gen-Z</p>
          <h1 class="bp-h1">มีดีต้อง<br/><span class="bp-accent">ปล่อยของ</span></h1>
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

      <div class="auth-card">
        <p class="ac-eyebrow">Welcome back</p>
        <h2 class="ac-title">เข้าสู่ระบบ</h2>
        <p class="ac-sub">ยินดีต้อนรับกลับ — พอร์ตคุณรอคุณอยู่ 👋</p>

        <form @submit.prevent="handleLogin" class="ac-form">

          <div class="field">
            <label class="field-lbl">Email Address</label>
            <div class="field-wrap">
              <svg class="field-ico" width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><rect x="2" y="4" width="20" height="16" rx="3"/><path d="m2 7 10 7 10-7"/></svg>
              <input v-model="form.email" type="email" class="field-inp" placeholder="you@example.com" autocomplete="email" required/>
            </div>
          </div>

          <div class="field">
            <div class="field-hdr">
              <label class="field-lbl">Password</label>
              <a href="#" class="field-forgot">ลืมรหัสผ่าน?</a>
            </div>
            <div class="field-wrap">
              <svg class="field-ico" width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><rect x="3" y="11" width="18" height="11" rx="2"/><path d="M7 11V7a5 5 0 0110 0v4"/></svg>
              <input v-model="form.password" :type="showPw ? 'text' : 'password'" class="field-inp" placeholder="••••••••" autocomplete="current-password" required/>
              <button type="button" class="field-eye" @click="showPw = !showPw">
                <svg v-if="!showPw" width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M1 12s4-8 11-8 11-8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/></svg>
                <svg v-else width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M17.94 17.94A10.07 10.07 0 0112 20c-7 0-11-8-11-8a18.45 18.45 0 015.06-5.94M9.9 4.24A9.12 9.12 0 0112 4c7 0 11 8 11 8a18.5 18.5 0 01-2.16 3.19m-6.72-1.07a3 3 0 11-4.24-4.24"/><line x1="1" y1="1" x2="23" y2="23"/></svg>
              </button>
            </div>
          </div>

          <transition name="err-fade">
            <div v-if="authStore.error" class="ac-error">
              <AlertTriangle :size="14"/>{{ authStore.error }}
            </div>
          </transition>

          <button type="submit" class="btn-submit" :disabled="authStore.loading">
            <span class="btn-inner">
              <Loader2 v-if="authStore.loading" :size="16" class="spin"/>
              <Rocket v-else :size="16"/>
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

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { AlertTriangle, Loader2, Rocket } from 'lucide-vue-next'

const authStore = useAuthStore()
const router    = useRouter()
const form      = reactive({ email: '', password: '' })
const showPw    = ref(false)

const features = [
  { icon: '🎨', label: 'Drag & Drop Builder — ลากวางได้เลย'   },
  { icon: '🤖', label: 'AI เกลาเนื้อหาให้เป็นมืออาชีพ'       },
  { icon: '📊', label: 'Analytics + Real-time Notifications'   },
  { icon: '🔒', label: 'Password-Protected Sections'           },
]
const stats = [
  { val: '48',    label: 'Templates' },
  { val: '<5min', label: 'Build Time'},
  { val: '∞',     label: 'Profiles'  },
  { val: 'Free',  label: 'Forever'   },
]

async function handleLogin() {
  const ok = await authStore.login(form.email, form.password)
  if (ok) router.push('/dashboard')
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Sora:wght@300;400;500;600;700;800&display=swap');

/* ── Tokens ─────────────────────────────────────────────── */
.auth-page {
  --gold:      #E8C547;
  --gold-s:    #F5D87E;
  --indigo:    #6366F1;
  --surface:   #0F1117;
  --s2:        rgba(255,255,255,0.04);
  --s3:        rgba(255,255,255,0.07);
  --border:    rgba(255,255,255,0.09);
  --bdr-gold:  rgba(232,197,71,0.38);
  --text:      #F1F5F9;
  --muted:     #64748B;
  --dim:       #94A3B8;
  --danger:    #FCA5A5;

  display: grid;
  grid-template-columns: 1fr 1fr;
  min-height: 100vh;
  font-family: 'Sora', 'Noto Sans Thai', system-ui, sans-serif;
  background: #080A12;
}
@media (max-width: 860px) {
  .auth-page { grid-template-columns: 1fr; }
  .brand-panel { display: none; }
}

/* ════ BRAND PANEL ═════════════════════════════════════════ */
.brand-panel {
  position: relative;
  display: flex; flex-direction: column; justify-content: space-between;
  padding: 52px 52px 44px;
  overflow: hidden;
  background: #07090E;
}

.bp-orb {
  position: absolute; border-radius: 50%;
  filter: blur(80px); pointer-events: none;
}
.bp-orb--1 {
  width:520px; height:520px;
  background: radial-gradient(circle, rgba(99,102,241,0.18), transparent 70%);
  top:-130px; left:-130px;
  animation: orbDrift 18s ease-in-out infinite alternate;
}
.bp-orb--2 {
  width:420px; height:420px;
  background: radial-gradient(circle, rgba(232,197,71,0.1), transparent 70%);
  bottom:-70px; right:-70px;
  animation: orbDrift 14s ease-in-out infinite alternate-reverse;
}
.bp-orb--3 {
  width:300px; height:300px;
  background: radial-gradient(circle, rgba(56,189,248,0.07), transparent 70%);
  top:50%; left:35%;
  animation: orbDrift 22s ease-in-out infinite alternate;
}
@keyframes orbDrift {
  from { transform: translate(0,0)      scale(1);   }
  to   { transform: translate(35px,25px) scale(1.1); }
}

.bp-grain {
  position:absolute; inset:0; pointer-events:none; opacity:0.25;
  background-image: url("data:image/svg+xml,%3Csvg viewBox='0 0 200 200' xmlns='http://www.w3.org/2000/svg'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.75' numOctaves='4' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23n)'/%3E%3C/svg%3E");
  background-size:160px; mix-blend-mode:overlay;
}
.bp-dots {
  position:absolute; inset:0; pointer-events:none;
  background-image: radial-gradient(circle, rgba(255,255,255,0.1) 1px, transparent 1px);
  background-size: 32px 32px;
  mask-image: radial-gradient(ellipse 80% 80% at 50% 50%, black 20%, transparent 76%);
}

.bp-inner {
  position:relative; z-index:2;
  flex:1; display:flex; flex-direction:column; justify-content:center; gap:38px;
}

/* Logo */
.bp-logo { display:flex; align-items:center; gap:12px; }
.bp-logo-mark {
  width:48px; height:48px; border-radius:14px; flex-shrink:0;
  background: linear-gradient(135deg, #1a1a4e, #2d1b69);
  border: 1px solid rgba(232,197,71,0.3);
  display:flex; align-items:center; justify-content:center;
  box-shadow: 0 0 24px rgba(99,102,241,0.3), inset 0 1px 0 rgba(255,255,255,0.1);
}
.bp-logo-name { font-size:26px; font-weight:800; letter-spacing:-0.5px; color:#fff; }
.bp-logo-name em {
  font-style:normal;
  background: linear-gradient(90deg, var(--gold), var(--gold-s));
  -webkit-background-clip:text; -webkit-text-fill-color:transparent; background-clip:text;
}

/* Headline */
.bp-eyebrow {
  font-size:10px; font-weight:700; letter-spacing:3px; text-transform:uppercase;
  color:var(--gold); opacity:0.75; margin-bottom:12px;
}
.bp-h1 {
  font-size:clamp(38px,3.8vw,50px); font-weight:800; letter-spacing:-1.5px;
  line-height:1.05; color:#fff; margin-bottom:14px;
}
.bp-accent {
  background: linear-gradient(135deg, var(--gold) 0%, var(--gold-s) 40%, var(--indigo) 100%);
  -webkit-background-clip:text; -webkit-text-fill-color:transparent; background-clip:text;
}
.bp-body { font-size:15px; line-height:1.7; color:var(--muted); font-weight:300; }

/* Features */
.bp-feats { list-style:none; display:flex; flex-direction:column; gap:11px; }
.bp-feat  { display:flex; align-items:center; gap:11px; font-size:13.5px; color:var(--dim); }
.bp-feat-ico {
  width:32px; height:32px; border-radius:9px; flex-shrink:0;
  background:var(--s3); border:1px solid var(--border);
  display:flex; align-items:center; justify-content:center; font-size:15px;
}

/* Badge */
.bp-badge {
  display:inline-flex; align-items:center; gap:9px; width:fit-content;
  padding:9px 16px; border-radius:100px;
  background:rgba(52,211,153,0.06); border:1px solid rgba(52,211,153,0.22);
  font-size:12px; color:#6EE7B7;
  animation:badgeFloat 4s ease-in-out infinite;
}
.bp-badge-dot {
  width:7px; height:7px; border-radius:50%; flex-shrink:0;
  background:#34D399; box-shadow:0 0 8px #34D399;
  animation:pulseDot 2s ease-in-out infinite;
}
@keyframes badgeFloat { 0%,100%{transform:translateY(0)} 50%{transform:translateY(-5px)} }
@keyframes pulseDot   { 0%,100%{box-shadow:0 0 6px #34D399} 50%{box-shadow:0 0 16px #34D399} }

/* Stats */
.bp-stats {
  position:relative; z-index:2;
  display:flex; gap:32px; padding-top:28px; border-top:1px solid var(--border);
}
.bp-stat  { display:flex; flex-direction:column; gap:2px; }
.bp-stat-val {
  font-size:22px; font-weight:800; letter-spacing:-0.5px;
  background: linear-gradient(90deg, var(--gold), var(--gold-s));
  -webkit-background-clip:text; -webkit-text-fill-color:transparent; background-clip:text;
}
.bp-stat-label { font-size:10px; color:var(--muted); letter-spacing:0.5px; }

/* ════ FORM PANEL ══════════════════════════════════════════ */
.form-panel {
  position:relative;
  display:flex; align-items:center; justify-content:center;
  padding:52px 40px; background:var(--surface); overflow:hidden;
}
.fp-glow {
  position:absolute; width:600px; height:600px; border-radius:50%;
  background:radial-gradient(circle, rgba(99,102,241,0.07), transparent 65%);
  top:50%; left:50%; transform:translate(-50%,-50%);
  animation:orbDrift 20s ease-in-out infinite alternate; pointer-events:none;
}

.auth-card {
  position:relative; z-index:2; width:100%; max-width:400px;
  animation:cardIn 0.5s cubic-bezier(0.16,1,0.3,1) both;
}
@keyframes cardIn {
  from { opacity:0; transform:translateY(18px); }
  to   { opacity:1; transform:translateY(0);    }
}

.ac-eyebrow {
  font-size:10px; font-weight:700; letter-spacing:3px; text-transform:uppercase;
  color:var(--gold); opacity:0.8; margin-bottom:8px;
}
.ac-title  { font-size:32px; font-weight:800; letter-spacing:-1px; color:var(--text); margin-bottom:6px; line-height:1.1; }
.ac-sub    { font-size:14px; color:var(--muted); margin-bottom:36px; font-weight:300; }

/* Form fields */
.ac-form   { display:flex; flex-direction:column; gap:20px; }
.field     { display:flex; flex-direction:column; gap:7px; }
.field-hdr { display:flex; justify-content:space-between; align-items:center; }
.field-lbl {
  font-size:11px; font-weight:700; letter-spacing:0.7px; text-transform:uppercase;
  color:var(--dim);
}
.field-forgot {
  font-size:12px; color:var(--muted); text-decoration:none; transition:color 0.2s;
}
.field-forgot:hover { color:var(--gold); }
.field-wrap { position:relative; }
.field-ico {
  position:absolute; left:14px; top:50%; transform:translateY(-50%);
  color:var(--muted); pointer-events:none; transition:color 0.2s;
}
.field-inp {
  width:100%; padding:13px 42px 13px 42px;
  background:var(--s2); border:1px solid var(--border); border-radius:13px;
  color:var(--text); font-family:inherit; font-size:14px;
  caret-color:var(--gold); outline:none;
  transition:border-color 0.2s, background 0.2s, box-shadow 0.2s;
}
.field-inp::placeholder { color:var(--muted); font-weight:300; }
.field-inp:focus {
  border-color:var(--bdr-gold); background:rgba(232,197,71,0.03);
  box-shadow:0 0 0 3px rgba(232,197,71,0.09);
}
.field-wrap:focus-within .field-ico { color:var(--gold); }
.field-eye {
  position:absolute; right:13px; top:50%; transform:translateY(-50%);
  background:none; border:none; cursor:pointer; color:var(--muted);
  padding:3px; transition:color 0.2s;
}
.field-eye:hover { color:var(--gold); }

/* Error */
.ac-error {
  padding:11px 14px; gap:8px;
  background:rgba(248,113,113,0.07); border:1px solid rgba(248,113,113,0.25);
  border-radius:11px; font-size:12.5px; color:var(--danger);
  display:flex; align-items:center;
}
.err-fade-enter-active, .err-fade-leave-active { transition:all 0.2s ease; }
.err-fade-enter-from, .err-fade-leave-to { opacity:0; transform:translateY(-4px); }

/* Submit */
.btn-submit {
  width:100%; padding:14px;
  background:linear-gradient(135deg, #4F46E5, #7C3AED 55%, #6366F1);
  border:none; border-radius:13px; color:#fff;
  font-family:inherit; font-size:14.5px; font-weight:700;
  cursor:pointer; position:relative; overflow:hidden;
  box-shadow:0 6px 24px rgba(99,102,241,0.38);
  transition:transform 0.22s, box-shadow 0.22s;
}
.btn-submit::after {
  content:''; position:absolute; inset:0;
  background:linear-gradient(135deg, transparent 30%, rgba(255,255,255,0.13), transparent 70%);
  transform:translateX(-100%); transition:transform 0.5s;
}
.btn-submit:hover:not(:disabled) { transform:translateY(-2px); box-shadow:0 12px 36px rgba(99,102,241,0.5); }
.btn-submit:hover:not(:disabled)::after { transform:translateX(100%); }
.btn-submit:active:not(:disabled) { transform:translateY(0); }
.btn-submit:disabled { opacity:0.5; cursor:not-allowed; }
.btn-inner { display:inline-flex; align-items:center; justify-content:center; gap:8px; }

/* Divider */
.ac-divider {
  display:flex; align-items:center; gap:12px; margin:22px 0;
}
.ac-divider::before,
.ac-divider::after { content:''; flex:1; height:1px; background:var(--border); }
.ac-divider span { font-size:11px; color:var(--muted); }

/* Google */
.btn-google {
  width:100%; padding:12px 20px;
  background:var(--s2); border:1px solid var(--border); border-radius:13px;
  color:var(--dim); font-family:inherit; font-size:13.5px; font-weight:500;
  cursor:pointer; display:flex; align-items:center; justify-content:center; gap:10px;
  transition:background 0.2s, border-color 0.2s, color 0.2s;
}
.btn-google:hover { background:var(--s3); border-color:rgba(255,255,255,0.15); color:var(--text); }

/* Switch */
.ac-switch {
  text-align:center; margin-top:28px; font-size:13.5px; color:var(--muted); font-weight:300;
}
.ac-link {
  color:var(--gold); text-decoration:none; font-weight:700; transition:color 0.2s, text-shadow 0.2s;
}
.ac-link:hover { color:var(--gold-s); text-shadow:0 0 12px rgba(232,197,71,0.35); }

@keyframes spin { to { transform:rotate(360deg); } }
.spin { animation:spin 0.9s linear infinite; }
</style>
