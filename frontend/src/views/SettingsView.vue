<template>
  <div class="settings-page">
    <div class="settings-header">
      <RouterLink to="/dashboard" class="back-btn" title="Dashboard">
        <ArrowLeft class="icon" :size="20" />
      </RouterLink>
      <h1><SettingsIcon :size="28" class="icon-inline" /> Settings</h1>
    </div>

    <div class="settings-card">
      <h3><User :size="18" class="icon-inline" /> Profile</h3>
      <div class="form-field">
        <label class="form-label">Name</label>
        <input v-model="form.name" class="form-input" />
      </div>
      <div class="form-field">
        <label class="form-label">Email</label>
        <input :value="authStore.user?.email" class="form-input" disabled style="opacity:0.5" />
      </div>
      <button class="btn btn-primary btn-icon-text" @click="saveProfile" :disabled="saving">
        <template v-if="saving"><Loader2 :size="16" class="spin" /> Saving...</template>
        <template v-else><Save :size="16" /> Save</template>
      </button>
    </div>

    <div class="settings-card">
      <h3><Palette :size="18" class="icon-inline" />App Appearance</h3>
      <div class="theme-options">
        <label
          class="theme-option"
          :class="{ active: themeStore.mode === 'light' }"
          @click="themeStore.setMode('light')"
        >
          <Sun :size="24" class="theme-icon" /> Light
        </label>
        <label
          class="theme-option"
          :class="{ active: themeStore.mode === 'dark' }"
          @click="themeStore.setMode('dark')"
        >
          <Moon :size="24" class="theme-icon" /> Dark
        </label>
        <label
          class="theme-option"
          :class="{ active: themeStore.mode === 'auto' }"
          @click="themeStore.setMode('auto')"
        >
          <Monitor :size="24" class="theme-icon" /> System
        </label>
      </div>
    </div>

    <div class="settings-card danger-zone">
      <h3><AlertTriangle :size="18" class="icon-inline" />Danger Zone</h3>
      <p>Deleting your account will delete all your data and portfolios, and cannot be undone.</p>
      <button class="btn btn-danger" @click="deleteAccount">Delete Account</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import Swal from 'sweetalert2'
import { useThemeStore } from '@/stores/theme'
import { userAPI } from '@/api'
import { Settings as SettingsIcon, User, Loader2, Save, Palette, Sun, Moon, Monitor, AlertTriangle, ArrowLeft } from 'lucide-vue-next'


const authStore = useAuthStore()
const themeStore = useThemeStore()
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
  const { isConfirmed } = await Swal.fire({
    title: 'Confirm deletion?',
    text: 'This action cannot be undone.',
    icon: 'warning',
    showCancelButton: true,
    confirmButtonColor: '#ef4444',
    cancelButtonColor: '#94a3b8',
    confirmButtonText: 'Confirm deletion'
  })
  if (!isConfirmed) return
  await userAPI.deleteMe()
  authStore.logout()
  router.push('/')
}
</script>

<style scoped>
.settings-page { max-width:600px;margin:0 auto;padding:36px 24px; }
.settings-header { display:flex;align-items:center;gap:16px;margin-bottom:32px; }
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
h1 { font-size:24px;font-weight:800; }
.settings-card { background:var(--surface);border:1px solid var(--border);border-radius:20px;padding:28px;margin-bottom:20px; }
.settings-card h3 { font-size:16px;font-weight:700;margin-bottom:20px; }
.form-field { display:flex;flex-direction:column;margin-bottom:14px; }
.danger-zone { border-color: rgba(239, 68, 68, 0.3); background: rgba(239, 68, 68, 0.05); }
.danger-zone p { font-size:13px;color:var(--danger);margin-bottom:16px; }

.theme-options {
  display: flex;
  gap: 12px;
}
.theme-option {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  padding: 16px;
  border: 1px solid var(--border);
  border-radius: 12px;
  cursor: pointer;
  background: var(--bg);
  transition: all 0.2s;
  color: var(--muted);
  font-weight: 600;
  font-size: 14px;
}
.theme-option:hover {
  background: var(--bg2);
}
.theme-option.active {
  background: rgba(79, 70, 229, 0.15);
  border-color: var(--indigo);
  color: var(--indigo);
}
.theme-icon {
  margin-bottom: 4px;
}
.icon-inline { vertical-align: text-bottom; margin-right: 4px; }
.btn-icon-text { display: inline-flex; align-items: center; justify-content: center; gap: 8px; }
@keyframes spin { 100% { transform: rotate(360deg); } }
.spin { animation: spin 1s linear infinite; }
</style>
