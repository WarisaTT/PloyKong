<template>
  <div class="auth-page">
    <div class="auth-bg"></div>
    <div class="auth-card">
      <div class="auth-logo">
        <span class="logo-pk">PK</span>
        <span class="logo-text grad-text">PloyKong</span>
      </div>
      <h2 class="auth-title">สมัครสมาชิกฟรี</h2>
      <p class="auth-sub">
        เริ่มสร้างพอร์ตโฟลิโอในไม่กี่คลิก
        <Rocket :size="14" class="icon-inline" />
      </p>

      <form @submit.prevent="handleRegister" class="auth-form">
        <div class="form-group">
          <label class="form-label">ชื่อ-นามสกุล</label>
          <input
            v-model="form.name"
            type="text"
            class="form-input"
            placeholder="Napat Kankham"
            required
          />
        </div>
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
            placeholder="อย่างน้อย 8 ตัวอักษร"
            required
            minlength="8"
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
            ><Loader2 :size="18" class="spin" /> กำลังสมัคร...</span
          >
          <span v-else class="btn-icon-text"
            ><PartyPopper :size="18" /> สมัครสมาชิกฟรี</span
          >
        </button>
      </form>

      <p class="auth-switch">
        มีบัญชีแล้ว?
        <RouterLink to="/login" class="auth-link">เข้าสู่ระบบ</RouterLink>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from "vue";
import { useRouter, RouterLink } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { Rocket, AlertTriangle, Loader2, PartyPopper } from "lucide-vue-next";

const authStore = useAuthStore();
const router = useRouter();
const form = reactive({ name: "", email: "", password: "" });

async function handleRegister() {
  const success = await authStore.register(
    form.email,
    form.password,
    form.name,
  );
  if (success) router.push("/dashboard");
}
</script>

<style scoped>
.auth-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
  position: relative;
}
.auth-bg {
  position: fixed;
  inset: 0;
  background:
    radial-gradient(
      ellipse 70% 60% at 70% 30%,
      rgba(79, 70, 229, 0.25) 0%,
      transparent 60%
    ),
    radial-gradient(
      ellipse 50% 40% at 20% 80%,
      rgba(6, 182, 212, 0.15) 0%,
      transparent 50%
    );
}
.auth-card {
  position: relative;
  z-index: 1;
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 24px;
  padding: 44px;
  width: 100%;
  max-width: 420px;
  backdrop-filter: blur(30px);
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.15);
}
.auth-logo {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 28px;
  justify-content: center;
}
.logo-pk {
  width: 44px;
  height: 44px;
  border-radius: 12px;
  background: linear-gradient(135deg, var(--indigo), var(--purple));
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: var(--font-display);
  font-size: 16px;
  font-weight: 800;
  color: #fff;
}
.logo-text {
  font-family: var(--font-display);
  font-size: 22px;
  font-weight: 800;
}
.auth-title {
  font-size: 24px;
  font-weight: 800;
  text-align: center;
  margin-bottom: 6px;
}
.auth-sub {
  color: var(--muted);
  font-size: 14px;
  text-align: center;
  margin-bottom: 32px;
}
.auth-form {
  display: flex;
  flex-direction: column;
  gap: 16px;
}
.form-group {
  display: flex;
  flex-direction: column;
}
.auth-error {
  padding: 10px 14px;
  background: rgba(239, 68, 68, 0.1);
  border: 1px solid rgba(239, 68, 68, 0.3);
  border-radius: 10px;
  font-size: 13px;
  color: var(--danger);
  display: flex;
  align-items: center;
  gap: 6px;
}
.auth-switch {
  text-align: center;
  margin-top: 20px;
  font-size: 14px;
  color: var(--muted);
}
.auth-link {
  color: var(--neon-purple);
  text-decoration: none;
  font-weight: 600;
}

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
