<template>
  <div class="auth-page">
    <div class="auth-bg"></div>
    <div class="auth-card">
      <div class="auth-logo">
        <span class="logo-pk">PK</span>
        <span class="logo-text grad-text">PloyKong</span>
      </div>
      <h2 class="auth-title">เข้าสู่ระบบ</h2>
      <p class="auth-sub">ยินดีต้อนรับกลับ! 👋</p>

      <form @submit.prevent="handleLogin" class="auth-form">
        <div class="form-group">
          <label class="form-label">Email</label>
          <input v-model="form.email" type="email" class="form-input" placeholder="you@example.com" required />
        </div>
        <div class="form-group">
          <label class="form-label">Password</label>
          <input v-model="form.password" type="password" class="form-input" placeholder="••••••••" required />
        </div>

        <div v-if="authStore.error" class="auth-error">
          ⚠️ {{ authStore.error }}
        </div>

        <button type="submit" class="btn btn-primary btn-lg" style="width:100%" :disabled="authStore.loading">
          <span v-if="authStore.loading">⏳ กำลังเข้าสู่ระบบ...</span>
          <span v-else">🚀 เข้าสู่ระบบ</span>
        </button>
      </form>

      <p class="auth-switch">
        ยังไม่มีบัญชี?
        <RouterLink to="/register" class="auth-link">สมัครสมาชิกฟรี</RouterLink>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
const router = useRouter()

const form = reactive({ email: '', password: '' })

async function handleLogin() {
  const success = await authStore.login(form.email, form.password)
  if (success) router.push('/dashboard')
}
</script>

<style scoped>
.auth-page {
  min-height: 100vh; display: flex; align-items: center; justify-content: center;
  padding: 20px; position: relative;
}
.auth-bg {
  position: fixed; inset: 0;
  background:
    radial-gradient(ellipse 70% 60% at 30% 20%, rgba(79,70,229,0.25) 0%, transparent 60%),
    radial-gradient(ellipse 50% 40% at 80% 80%, rgba(6,182,212,0.15) 0%, transparent 50%);
}
.auth-card {
  position: relative; z-index: 1;
  background: rgba(10, 14, 30, 0.9);
  border: 1px solid var(--border);
  border-radius: 24px; padding: 44px;
  width: 100%; max-width: 420px;
  backdrop-filter: blur(30px);
  box-shadow: 0 40px 80px rgba(0,0,0,0.5);
}
.auth-logo { display: flex; align-items: center; gap: 10px; margin-bottom: 28px; justify-content: center; }
.logo-pk {
  width: 44px; height: 44px; border-radius: 12px;
  background: linear-gradient(135deg, var(--indigo), var(--purple));
  display: flex; align-items: center; justify-content: center;
  font-family: var(--font-display); font-size: 16px; font-weight: 800; color: #fff;
}
.logo-text { font-family: var(--font-display); font-size: 22px; font-weight: 800; }
.auth-title { font-size: 24px; font-weight: 800; text-align: center; margin-bottom: 6px; }
.auth-sub { color: var(--muted); font-size: 14px; text-align: center; margin-bottom: 32px; }
.auth-form { display: flex; flex-direction: column; gap: 16px; }
.form-group { display: flex; flex-direction: column; }
.auth-error {
  padding: 10px 14px; background: rgba(239,68,68,0.1);
  border: 1px solid rgba(239,68,68,0.3); border-radius: 10px;
  font-size: 13px; color: #fca5a5;
}
.auth-switch { text-align: center; margin-top: 20px; font-size: 14px; color: var(--muted); }
.auth-link { color: var(--neon-purple); text-decoration: none; font-weight: 600; }
.auth-link:hover { text-decoration: underline; }
</style>
