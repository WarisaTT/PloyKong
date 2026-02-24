<template>
  <div class="settings-page">
    <div class="settings-header">
      <RouterLink to="/dashboard" class="back-btn">← Dashboard</RouterLink>
      <h1>⚙️ Settings</h1>
    </div>

    <div class="settings-card">
      <h3>👤 โปรไฟล์</h3>
      <div class="form-field">
        <label class="form-label">ชื่อ</label>
        <input v-model="form.name" class="form-input" />
      </div>
      <div class="form-field">
        <label class="form-label">Email</label>
        <input :value="authStore.user?.email" class="form-input" disabled style="opacity:0.5" />
      </div>
      <button class="btn btn-primary" @click="saveProfile" :disabled="saving">
        {{ saving ? '⏳ Saving...' : '💾 บันทึก' }}
      </button>
    </div>

    <div class="settings-card danger-zone">
      <h3>⚠️ Danger Zone</h3>
      <p>การลบบัญชีจะลบข้อมูลและพอร์ตโฟลิโอทั้งหมด ไม่สามารถกู้คืนได้</p>
      <button class="btn btn-danger" @click="deleteAccount">ลบบัญชี</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { userAPI } from '@/api'

const authStore = useAuthStore()
const router = useRouter()
const saving = ref(false)
const form = reactive({ name: authStore.user?.name || '' })

async function saveProfile() {
  saving.value = true
  await userAPI.updateMe({ name: form.name })
  if (authStore.user) authStore.user.name = form.name
  saving.value = false
}

async function deleteAccount() {
  if (!confirm('ยืนยันการลบบัญชีทั้งหมด? ไม่สามารถย้อนกลับได้')) return
  await userAPI.deleteMe()
  authStore.logout()
  router.push('/')
}
</script>

<style scoped>
.settings-page { max-width:600px;margin:0 auto;padding:36px 24px; }
.settings-header { display:flex;align-items:center;gap:16px;margin-bottom:32px; }
.back-btn { color:var(--muted);text-decoration:none;font-size:14px; }
h1 { font-size:24px;font-weight:800; }
.settings-card { background:var(--surface);border:1px solid var(--border);border-radius:20px;padding:28px;margin-bottom:20px; }
.settings-card h3 { font-size:16px;font-weight:700;margin-bottom:20px; }
.form-field { display:flex;flex-direction:column;margin-bottom:14px; }
.danger-zone { border-color:rgba(239,68,68,0.3);background:rgba(239,68,68,0.05); }
.danger-zone p { font-size:13px;color:var(--muted);margin-bottom:16px; }
</style>
