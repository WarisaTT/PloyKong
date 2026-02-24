import axios from 'axios'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || '/api/v1',
  timeout: 30000,
  headers: { 'Content-Type': 'application/json' }
})

// ─── Request Interceptor: attach access token ────────────────────────────────
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('access_token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => Promise.reject(error)
)

// ─── Response Interceptor: handle 401 → refresh token ───────────────────────
let isRefreshing = false
let failedQueue: Array<{ resolve: Function; reject: Function }> = []

const processQueue = (error: any, token: string | null = null) => {
  failedQueue.forEach((prom) => {
    if (error) prom.reject(error)
    else prom.resolve(token)
  })
  failedQueue = []
}

api.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config

    if (error.response?.status === 401 && !originalRequest._retry) {
      if (isRefreshing) {
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject })
        })
          .then((token) => {
            originalRequest.headers.Authorization = `Bearer ${token}`
            return api(originalRequest)
          })
          .catch((err) => Promise.reject(err))
      }

      originalRequest._retry = true
      isRefreshing = true

      const refreshToken = localStorage.getItem('refresh_token')
      if (!refreshToken) {
        isRefreshing = false
        localStorage.clear()
        window.location.href = '/login'
        return Promise.reject(error)
      }

      try {
        const { data } = await axios.post('/api/v1/auth/refresh', {
          refresh_token: refreshToken
        })

        const { access_token, refresh_token } = data.data
        localStorage.setItem('access_token', access_token)
        localStorage.setItem('refresh_token', refresh_token)

        api.defaults.headers.common.Authorization = `Bearer ${access_token}`
        processQueue(null, access_token)

        return api(originalRequest)
      } catch (refreshError) {
        processQueue(refreshError, null)
        localStorage.clear()
        window.location.href = '/login'
        return Promise.reject(refreshError)
      } finally {
        isRefreshing = false
      }
    }

    return Promise.reject(error)
  }
)

export default api

// ─── Typed API Services ──────────────────────────────────────────────────────

export const authAPI = {
  register: (data: { email: string; password: string; name: string }) =>
    api.post('/auth/register', data),
  login: (data: { email: string; password: string }) =>
    api.post('/auth/login', data),
  logout: (refreshToken: string) =>
    api.post('/auth/logout', { refresh_token: refreshToken }),
  refresh: (refreshToken: string) =>
    api.post('/auth/refresh', { refresh_token: refreshToken })
}

export const userAPI = {
  getMe: () => api.get('/users/me'),
  updateMe: (data: Partial<{ name: string; avatar_url: string }>) =>
    api.put('/users/me', data),
  deleteMe: () => api.delete('/users/me')
}

export const portfolioAPI = {
  list: () => api.get('/portfolios'),
  create: (data: { slug: string; title: string; theme?: any }) =>
    api.post('/portfolios', data),
  getById: (id: string) => api.get(`/portfolios/${id}`),
  update: (id: string, data: any) => api.put(`/portfolios/${id}`, data),
  delete: (id: string) => api.delete(`/portfolios/${id}`),
  publish: (id: string) => api.post(`/portfolios/${id}/publish`),
  unpublish: (id: string) => api.post(`/portfolios/${id}/unpublish`),
  exportPDF: (id: string) => api.get(`/portfolios/${id}/export-pdf`)
}

export const sectionAPI = {
  list: (portfolioId: string) =>
    api.get(`/portfolios/${portfolioId}/sections`),
  create: (portfolioId: string, data: any) =>
    api.post(`/portfolios/${portfolioId}/sections`, data),
  update: (portfolioId: string, id: string, data: any) =>
    api.put(`/portfolios/${portfolioId}/sections/${id}`, data),
  delete: (portfolioId: string, id: string) =>
    api.delete(`/portfolios/${portfolioId}/sections/${id}`),
  reorder: (portfolioId: string, order: { id: string; position: number }[]) =>
    api.post(`/portfolios/${portfolioId}/sections/reorder`, { order })
}

export const analyticsAPI = {
  summary: (portfolioId: string) =>
    api.get(`/analytics/${portfolioId}/summary`),
  visitors: (portfolioId: string) =>
    api.get(`/analytics/${portfolioId}/visitors`),
  sources: (portfolioId: string) =>
    api.get(`/analytics/${portfolioId}/sources`)
}

export const aiAPI = {
  improveText: (text: string, context: string) =>
    api.post('/ai/improve-text', { text, context }),
  scoreResume: (content: string, targetRole: string) =>
    api.post('/ai/score-resume', { portfolio_content: content, target_role: targetRole }),
  suggestSkills: (role: string, currentSkills: string[]) =>
    api.post('/ai/suggest-skills', { role, current_skills: currentSkills })
}

export const publicAPI = {
  view: (slug: string, password?: string) =>
    axios.get(`/api/v1/public/p/${slug}`, {
      headers: password ? { 'X-Portfolio-Password': password } : {}
    }),
  track: (slug: string, eventType: string, source?: string) =>
    axios.post(`/api/v1/public/p/${slug}/track`, { event_type: eventType, source }),
  chat: (slug: string, message: string, sessionId: string) =>
    axios.post(`/api/v1/public/p/${slug}/chat`, { message, session_id: sessionId })
}
