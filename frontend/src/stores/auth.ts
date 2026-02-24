import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authAPI, userAPI } from '@/api'
import type { User } from '@/types'

export const useAuthStore = defineStore('auth', () => {
  // ─── State ────────────────────────────────────────────────────────────────
  const user = ref<User | null>(null)
  const accessToken = ref<string | null>(localStorage.getItem('access_token'))
  const refreshToken = ref<string | null>(localStorage.getItem('refresh_token'))
  const loading = ref(false)
  const error = ref<string | null>(null)

  // ─── Getters ──────────────────────────────────────────────────────────────
  const isAuthenticated = computed(() => !!accessToken.value)
  const isPro = computed(() => user.value?.plan === 'pro' || user.value?.plan === 'team')

  // ─── Actions ──────────────────────────────────────────────────────────────
  async function login(email: string, password: string) {
    loading.value = true
    error.value = null
    try {
      const { data } = await authAPI.login({ email, password })
      setTokens(data.data.tokens.access_token, data.data.tokens.refresh_token)
      user.value = data.data.user
      return true
    } catch (e: any) {
      error.value = e.response?.data?.error || 'Login failed'
      return false
    } finally {
      loading.value = false
    }
  }

  async function register(email: string, password: string, name: string) {
    loading.value = true
    error.value = null
    try {
      const { data } = await authAPI.register({ email, password, name })
      setTokens(data.data.tokens.access_token, data.data.tokens.refresh_token)
      user.value = data.data.user
      return true
    } catch (e: any) {
      error.value = e.response?.data?.error || 'Registration failed'
      return false
    } finally {
      loading.value = false
    }
  }

  async function logout() {
    if (refreshToken.value) {
      await authAPI.logout(refreshToken.value).catch(() => {})
    }
    clearTokens()
    user.value = null
  }

  async function fetchMe() {
    if (!accessToken.value) return
    try {
      const { data } = await userAPI.getMe()
      user.value = data.data
    } catch {
      clearTokens()
      user.value = null
    }
  }

  function setTokens(access: string, refresh: string) {
    accessToken.value = access
    refreshToken.value = refresh
    localStorage.setItem('access_token', access)
    localStorage.setItem('refresh_token', refresh)
  }

  function clearTokens() {
    accessToken.value = null
    refreshToken.value = null
    localStorage.removeItem('access_token')
    localStorage.removeItem('refresh_token')
  }

  return {
    user, accessToken, refreshToken, loading, error,
    isAuthenticated, isPro,
    login, register, logout, fetchMe
  }
})
