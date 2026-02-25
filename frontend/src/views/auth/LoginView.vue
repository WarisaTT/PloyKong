<template>
  <div class="auth-page">
    <div class="auth-bg"></div>
    <div class="auth-card">
      <div class="auth-logo">
        <span class="logo-pk">PK</span>
        <span class="logo-text grad-text">PloyKong</span>
      </div>
      <h2 class="auth-title">เข้าสู่ระบบ</h2>
      <p class="auth-sub">ยินดีต้อนรับกลับ!</p>

      <form @submit.prevent="handleLogin" class="auth-form">
        <div class="form-group">
          <label class="form-label">Email</label>
          <input
            v-model="form.email"
            type="email"
            class="form-input"
            placeholder="you@example.com"
            required
          />
        </div>
        <div class="form-group">
          <label class="form-label">Password</label>
          <input
            v-model="form.password"
            type="password"
            class="form-input"
            placeholder="••••••••"
            required
          />
        </div>

        <div v-if="authStore.error" class="auth-error">
          <AlertTriangle :size="14" class="icon-inline" /> {{ authStore.error }}
        </div>

        <button
          type="submit"
          class="btn btn-primary btn-lg btn-icon-text"
          style="width: 100%"
          :disabled="authStore.loading"
        >
          <span v-if="authStore.loading" class="btn-icon-text"
            ><Loader2 :size="18" class="spin" /> กำลังเข้าสู่ระบบ...</span
          >
          <span v-else class="btn-icon-text"
            ><Rocket :size="18" /> เข้าสู่ระบบ</span
          >
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
import { reactive } from "vue";
import { useRouter, RouterLink } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { Hand, AlertTriangle, Loader2, Rocket } from "lucide-vue-next";

const authStore = useAuthStore();
const router = useRouter();

const form = reactive({ email: "", password: "" });

async function handleLogin() {
  const success = await authStore.login(form.email, form.password);
  if (success) router.push("/dashboard");
}
</script>

<style scoped>
.auth-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
  position: relative;
  overflow: hidden;
}

/* ✨ Animated Background */
.auth-bg {
  position: fixed;
  inset: 0;
  background: linear-gradient(135deg, var(--indigo), var(--purple));
  animation: floatBg 12s ease-in-out infinite alternate;
}

@keyframes floatBg {
  0% {
    transform: scale(1) rotate(0deg);
  }
  100% {
    transform: scale(1.05) rotate(2deg);
  }
}

/* ✨ Card */
.auth-card {
  position: relative;
  z-index: 1;
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 28px;
  padding: 48px 44px;
  width: 100%;
  max-width: 420px;

  backdrop-filter: blur(40px);
  box-shadow:
    0 30px 80px rgba(0, 0, 0, 0.15),
    0 0 60px rgba(147, 51, 234, 0.08);

  transition: transform 0.3s ease;
}

.auth-card:hover {
  transform: translateY(-4px);
}

/* ✨ Logo */
.auth-logo {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 32px;
  justify-content: center;
}

.logo-pk {
  width: 48px;
  height: 48px;
  border-radius: 14px;
  background: linear-gradient(135deg, var(--indigo), var(--purple));
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: var(--font-display);
  font-size: 18px;
  font-weight: 800;
  color: #fff;
  box-shadow: 0 8px 24px rgba(147, 51, 234, 0.4);
}

.logo-text {
  font-family: var(--font-display);
  font-size: 24px;
  font-weight: 800;
}

/* ✨ Title */
.auth-title {
  font-size: 26px;
  font-weight: 800;
  text-align: center;
  margin-bottom: 6px;
}

.auth-sub {
  color: var(--muted);
  font-size: 14px;
  text-align: center;
  margin-bottom: 36px;
}

/* ✨ Form */
.auth-form {
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

/* ✨ Inputs */
/* ✨ Inputs */
.form-input {
  padding: 12px 16px;
  border-radius: 14px;
  border: 1px solid var(--border);

  background: rgba(15, 20, 40, 0.85); /* แทนสีเทาขาว */
  color: #f1f5f9; /* 👈 ตัวหนังสือชัดขึ้น */
  caret-color: var(--purple);

  font-size: 14px;
  transition: all 0.2s ease;
  outline: none;
}

/* placeholder */
.form-input::placeholder {
  color: #94a3b8; /* ไม่จางเกิน */
  opacity: 0.8;
}

/* focus */
.form-input:focus {
  background: rgba(20, 25, 50, 0.95);
  border-color: var(--purple);
  box-shadow: 0 0 0 4px rgba(147, 51, 234, 0.2);
}

/* ✨ Error */
.auth-error {
  padding: 12px 14px;
  background: rgba(239, 68, 68, 0.08);
  border: 1px solid rgba(239, 68, 68, 0.3);
  border-radius: 12px;
  font-size: 13px;
  color: var(--danger);
  display: flex;
  align-items: center;
  gap: 6px;
}

/* ✨ Button */
.btn-lg {
  padding: 12px 20px;
  font-size: 15px;
  border-radius: 14px;
}

.btn-primary {
  background: linear-gradient(135deg, var(--indigo), var(--purple));
  box-shadow: 0 10px 30px rgba(79, 70, 229, 0.35);
  transition: all 0.25s ease;
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 16px 40px rgba(147, 51, 234, 0.5);
}

/* ✨ Switch */
.auth-switch {
  text-align: center;
  margin-top: 24px;
  font-size: 14px;
  color: var(--muted);
}

.auth-link {
  color: var(--purple);
  text-decoration: none;
  font-weight: 600;
  transition: 0.2s;
}

.auth-link:hover {
  color: var(--indigo);
  text-decoration: underline;
}

/* ✨ Icons */
.icon-inline {
  vertical-align: middle;
}

.btn-icon-text {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

@keyframes spin {
  100% {
    transform: rotate(360deg);
  }
}

.spin {
  animation: spin 1s linear infinite;
}
</style>
