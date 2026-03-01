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
        <!-- Badge -->
        <div class="bp-badge-top">✦ เริ่มต้นได้ฟรี · ไม่ต้องใส่บัตรเครดิต</div>

        <!-- Headline -->
        <div class="bp-headline">
          <h1 class="bp-h1">พอร์ตสวย<br/><span class="bp-accent">ใน 5 นาที</span><br/>จริงๆ</h1>
          <p class="bp-body">
            ผู้สมัครที่มีเว็บพอร์ตโฟลิโอได้รับการตอบรับจาก HR<br/>
            มากกว่าถึง <strong style="color: #fff">3.4 เท่า</strong> เมื่อเทียบกับ Resume ธรรมดา
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

      <!-- Stats strip (matching login style) -->
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
        <div class="card-logo">
          <div class="logo-mark">PK</div>
          <div class="logo-name">PloyKong</div>
        </div>

        <h2 class="ac-title">สร้างบัญชีฟรี</h2>
        <p class="ac-sub">เริ่มสร้างพอร์ตโฟลิโอแรกของคุณในวันนี้ ✦</p>

        <button class="btn-google" type="button">
          <svg width="17" height="17" viewBox="0 0 24 24"><path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"/><path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/><path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05"/><path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/></svg>
          สมัครด้วย Google — เร็วสุด
        </button>

        <div class="ac-divider"><span>หรือกรอกข้อมูล</span></div>

        <form @submit.prevent="handleRegister" class="ac-form">
          <div class="field">
            <label class="field-lbl">ชื่อ-นามสกุล</label>
            <div class="field-wrap">
              <User class="field-ico" :size="15" />
              <input v-model="form.name" type="text" class="field-inp" placeholder="Napat Kankham" required/>
            </div>
          </div>

          <div class="field">
            <label class="field-lbl">Email Address</label>
            <div class="field-wrap">
              <Mail class="field-ico" :size="15" />
              <input v-model="form.email" type="email" class="field-inp" placeholder="you@example.com" autocomplete="email" required/>
            </div>
          </div>

          <div class="field">
            <label class="field-lbl">Password</label>
            <div class="field-wrap">
              <Lock class="field-ico" :size="15" />
              <input v-model="form.password" :type="showPw ? 'text' : 'password'" class="field-inp" placeholder="อย่างน้อย 8 ตัวอักษร" minlength="8" required/>
              <button type="button" class="field-eye" @click="showPw = !showPw">
                <Eye v-if="!showPw" :size="15" />
                <EyeOff v-else :size="15" />
              </button>
            </div>
          </div>

          <!-- Password strength indicator mockup -->
          <div class="pw-strength">
            <div class="pw-bars">
              <div class="pw-bar active"></div>
              <div class="pw-bar active"></div>
              <div class="pw-bar"></div>
              <div class="pw-bar"></div>
            </div>
            <div class="pw-lbl">ความปลอดภัยของรหัสผ่าน: ปานกลาง</div>
          </div>

          <transition name="err-fade">
            <div v-if="authStore.error" class="ac-error">
              <AlertTriangle :size="14"/>{{ authStore.error }}
            </div>
          </transition>

          <button type="submit" class="btn-submit" :disabled="authStore.loading">
            <span class="btn-inner">
              <Loader2 v-if="authStore.loading" :size="16" class="spin"/>
              <Check v-else :size="16"/>
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

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { AlertTriangle, Loader2, User, Mail, Lock, Check, Eye, EyeOff } from 'lucide-vue-next'

const authStore = useAuthStore()
const router    = useRouter()
const form      = reactive({ name: '', email: '', password: '' })
const showPw    = ref(false)

const steps = [
  { color: 'linear-gradient(135deg,#4F46E5,#7C3AED)', title: 'สมัครฟรี · ใช้ได้ทันที', desc: 'ไม่มีค่าใช้จ่าย ไม่ต้องใส่ข้อมูลบัตร' },
  { color: 'linear-gradient(135deg,#7C3AED,#06B6D4)', title: 'เลือก Template · ลากวาง', desc: '48 template พร้อมใช้ ตามสายงานคุณ' },
  { color: 'linear-gradient(135deg,#06B6D4,#10B981)', title: 'Publish · แชร์ทันที', desc: 'URL สวยงาม ส่งให้ HR ได้เลย' },
]

const stats = [
  { val: '2.4k+', label: 'New This Week' },
  { val: '3.4x', label: 'Higher Response' },
  { val: '0฿',   label: 'Cost to Start' },
]

async function handleRegister() {
  const success = await authStore.register(form.email, form.password, form.name)
  if (success) router.push('/dashboard')
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Sora:wght@300;400;500;600;700;800&family=Syne:wght@700;800&display=swap');

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
  grid-template-columns: 1.1fr 0.9fr;
  min-height: 100vh;
  font-family: 'Sora', 'Noto Sans Thai', system-ui, sans-serif;
  background: #080A12;
}

@media (max-width: 900px) {
  .auth-page { grid-template-columns: 1fr; }
  .brand-panel { display: none; }
}

/* ════ BRAND PANEL ═════════════════════════════════════════ */
.brand-panel {
  position: relative;
  display: flex; flex-direction: column; justify-content: space-between;
  padding: 52px 72px 44px;
  overflow: hidden;
  background: #07090E;
}

.brand-panel::before {
  content: 'PK';
  position: absolute;
  font-family: 'Syne', sans-serif;
  font-weight: 800;
  font-size: clamp(200px, 25vw, 350px);
  line-height: 1;
  top: 50%;
  left: -20px;
  transform: translateY(-50%);
  background: linear-gradient(160deg, rgba(99,102,241,0.12) 0%, rgba(56,189,248,0.06) 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  pointer-events: none;
  user-select: none;
  letter-spacing: -15px;
}

.bp-orb {
  position: absolute; border-radius: 50%;
  filter: blur(80px); pointer-events: none;
}
.bp-orb--1 { width:500px; height:500px; background: radial-gradient(circle, rgba(99,102,241,0.15), transparent 70%); top:-100px; left:-100px; }
.bp-orb--2 { width:400px; height:400px; background: radial-gradient(circle, rgba(232,197,71,0.08), transparent 70%); bottom:-80px; right:80px; }
.bp-orb--3 { width:300px; height:300px; background: radial-gradient(circle, rgba(224,64,251,0.06), transparent 70%); top:35%; right:200px; }

.bp-grain {
  position:absolute; inset:0; pointer-events:none; opacity:0.25;
  background-image: url("data:image/svg+xml,%3Csvg viewBox='0 0 200 200' xmlns='http://www.w3.org/2000/svg'%3E%3Cfilter id='n'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.75' numOctaves='4' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23n)'/%3E%3C/svg%3E");
  background-size:160px; mix-blend-mode:overlay;
}
.bp-dots {
  position:absolute; inset:0; pointer-events:none;
  background-image: radial-gradient(circle, rgba(255,255,255,0.06) 1px, transparent 1px);
  background-size: 32px 32px;
}

.bp-inner { position:relative; z-index:2; flex:1; display:flex; flex-direction:column; justify-content:center; gap:32px; }

.bp-badge-top {
  display: inline-flex; align-items: center; gap: 8px;
  padding: 6px 16px; border-radius: 100px;
  border: 1px solid rgba(123,97,255,0.3);
  background: rgba(123,97,255,0.08);
  font-size: 11px; font-weight: 700; color: #A89EFF;
  letter-spacing: 1px; text-transform: uppercase; width: fit-content;
}

.bp-h1 { font-family: 'Syne', sans-serif; font-size:clamp(36px, 4vw, 56px); font-weight:800; line-height:1.05; letter-spacing:-1.5px; color:#fff; }
.bp-accent { background: linear-gradient(135deg, #7B61FF 0%, #E040FB 50%, #06B6D4 100%); -webkit-background-clip:text; -webkit-text-fill-color:transparent; background-clip:text; }
.bp-body { font-size:15px; line-height:1.8; color:var(--muted); font-weight:300; }

.bp-steps { display:flex; flex-direction:column; gap:20px; }
.bp-step { display:flex; align-items:flex-start; gap:16px; }
.bp-step-num { width:32px; height:32px; border-radius:50%; display:flex; align-items:center; justify-content:center; font-family: 'Syne', sans-serif; font-size:13px; font-weight:800; color:#fff; flex-shrink:0; box-shadow:0 4px 14px rgba(79,70,229,0.3); }
.bp-step-title { font-size:14px; font-weight:600; color:#fff; margin-bottom:2px; }
.bp-step-desc { font-size:12.5px; color:var(--muted); }

.bp-stats { display:flex; gap:32px; padding-top:28px; border-top:1px solid var(--border); }
.bp-stat { display:flex; flex-direction:column; gap:2px; }
.bp-stat-val { font-size:22px; font-weight:800; letter-spacing:-0.5px; color: #fff; }
.bp-stat-label { font-size:10px; color:var(--muted); letter-spacing:0.5px; text-transform: uppercase; }

/* ════ FORM PANEL ══════════════════════════════════════════ */
.form-panel { position:relative; display:flex; align-items:center; justify-content:center; padding:52px 40px; background:var(--surface); overflow:hidden; }
.fp-glow { position:absolute; width:600px; height:600px; border-radius:50%; background:radial-gradient(circle, rgba(99,102,241,0.07), transparent 65%); top:50%; left:50%; transform:translate(-50%,-50%); pointer-events:none; }

.auth-card { position:relative; z-index:2; width:100%; max-width:400px; animation:cardIn 0.5s cubic-bezier(0.16,1,0.3,1) both; }
@keyframes cardIn { from { opacity:0; transform:translateY(18px); } to { opacity:1; transform:translateY(0); } }

.card-logo { display:flex; align-items:center; gap:10px; margin-bottom:28px; }
.logo-mark { width:36px; height:36px; border-radius:10px; background:linear-gradient(135deg, #4F46E5, #7C3AED, #06B6D4); display:flex; align-items:center; justify-content:center; font-family: 'Syne', sans-serif; font-size:14px; font-weight:800; color:#fff; box-shadow:0 4px 12px rgba(123,97,255,0.4); }
.logo-name { font-family: 'Syne', sans-serif; font-size:18px; font-weight:700; color:#fff; }

.ac-title { font-size:32px; font-weight:800; letter-spacing:-1px; color:var(--text); margin-bottom:6px; line-height:1.1; }
.ac-sub { font-size:14px; color:var(--muted); margin-bottom:24px; font-weight:300; }

.ac-form { display:flex; flex-direction:column; gap:16px; }
.field { display:flex; flex-direction:column; gap:7px; }
.field-lbl { font-size:11px; font-weight:700; letter-spacing:0.7px; text-transform:uppercase; color:var(--dim); }
.field-wrap { position:relative; }
.field-ico { position:absolute; left:14px; top:50%; transform:translateY(-50%); color:var(--muted); pointer-events:none; transition:color 0.2s; }
.field-inp { width:100%; padding:13px 42px; background:var(--s2); border:1px solid var(--border); border-radius:13px; color:var(--text); font-family:inherit; font-size:14px; caret-color:var(--indigo); outline:none; transition:all 0.2s; }
.field-inp:focus { border-color:var(--indigo); background:rgba(99,102,241,0.03); box-shadow:0 0 0 3px rgba(99,102,241,0.1); }
.field-wrap:focus-within .field-ico { color:var(--indigo); }
.field-eye { position:absolute; right:13px; top:50%; transform:translateY(-50%); background:none; border:none; cursor:pointer; color:var(--muted); padding:3px; }

.pw-strength { margin-top: -8px; }
.pw-bars { display:flex; gap:4px; margin-bottom:6px; }
.pw-bar { flex:1; height:3px; border-radius:100px; background:var(--border); }
.pw-bar.active { background: linear-gradient(90deg, #4F46E5, #7C3AED); }
.pw-lbl { font-size:10.5px; color:var(--muted); }

.ac-divider { display:flex; align-items:center; gap:12px; margin:16px 0; }
.ac-divider::before, .ac-divider::after { content:''; flex:1; height:1px; background:var(--border); }
.ac-divider span { font-size:11px; color:var(--muted); text-transform: uppercase; letter-spacing: 0.5px; }

.btn-google { width:100%; padding:12px; background:var(--s2); border:1px solid var(--border); border-radius:13px; color:var(--dim); font-size:13.5px; font-weight:500; cursor:pointer; display:flex; align-items:center; justify-content:center; gap:10px; transition:all 0.2s; margin-bottom: 8px; }
.btn-google:hover { background:var(--s3); border-color:rgba(255,255,255,0.15); color:var(--text); }

.btn-submit { width:100%; padding:14px; background:linear-gradient(135deg, #4F46E5, #7C3AED 55%, #06B6D4); border:none; border-radius:13px; color:#fff; font-size:14.5px; font-weight:700; cursor:pointer; position:relative; overflow:hidden; box-shadow:0 6px 24px rgba(99,102,241,0.3); transition:all 0.22s; }
.btn-submit:hover:not(:disabled) { transform:translateY(-2px); box-shadow:0 12px 36px rgba(99,102,241,0.45); }
.btn-submit:disabled { opacity:0.5; cursor:not-allowed; }
.btn-inner { display:inline-flex; align-items:center; justify-content:center; gap:8px; }

.ac-switch { text-align:center; margin-top:24px; font-size:13.5px; color:var(--muted); font-weight:300; }
.ac-link { color:var(--indigo); text-decoration:none; font-weight:700; }
.ac-link:hover { color:#818CF8; }

.ac-error { padding:11px 14px; gap:8px; background:rgba(248,113,113,0.07); border:1px solid rgba(248,113,113,0.25); border-radius:11px; font-size:12.5px; color:var(--danger); display:flex; align-items:center; }
.err-fade-enter-active, .err-fade-leave-active { transition:all 0.2s ease; }
.err-fade-enter-from, .err-fade-leave-to { opacity:0; transform:translateY(-4px); }

.spin { animation: spin 1s linear infinite; }
@keyframes spin { to { transform:rotate(360deg); } }
</style>
