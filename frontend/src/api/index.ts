import axios from "axios";
import { showError } from "@/utils/alert";

const getErrorMessage = (error: any) => {
  try {
    return (
      error?.response?.data?.error ||
      error?.response?.data?.message ||
      error?.message ||
      "An unexpected error occurred"
    );
  } catch (e) {
    return "An unexpected error occurred";
  }
};

const isNgrok = (import.meta.env.VITE_API_URL || "").includes("ngrok");

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || "/api/v1",
  timeout: 30000,
  headers: {
    "Content-Type": "application/json",
    ...(isNgrok && { "ngrok-skip-browser-warning": "true" }),
  },
});

// ─── Request Interceptor: attach access token ─────────────────────────────────
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem("access_token");
    if (token) config.headers.Authorization = `Bearer ${token}`;
    return config;
  },
  (error) => Promise.reject(error),
);

// ─── Response Interceptor: handle 401 → refresh token ────────────────────────
let isRefreshing = false;
let failedQueue: Array<{ resolve: Function; reject: Function }> = [];

const processQueue = (error: any, token: string | null = null) => {
  failedQueue.forEach((p) => (error ? p.reject(error) : p.resolve(token)));
  failedQueue = [];
};

api.interceptors.response.use(
  (response) => response,
  async (error) => {
    const original = error.config;
    if (error.response?.status === 401 && !original._retry) {
      if (isRefreshing) {
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject });
        }).then((token) => {
          original.headers.Authorization = `Bearer ${token}`;
          return api(original);
        });
      }
      original._retry = true;
      isRefreshing = true;
      const refreshToken = localStorage.getItem("refresh_token");
      if (!refreshToken) {
        isRefreshing = false;
        localStorage.clear();
        window.location.href = "/login";
        return Promise.reject(error);
      }
      try {
        const res = await axios.post(
          `${import.meta.env.VITE_API_URL || "/api/v1"}/auth/refresh`,
          { refresh_token: refreshToken },
          { headers: isNgrok ? { "ngrok-skip-browser-warning": "true" } : {} },
        );
        const { access_token, refresh_token } = res.data.data;
        localStorage.setItem("access_token", access_token);
        localStorage.setItem("refresh_token", refresh_token);
        api.defaults.headers.common.Authorization = `Bearer ${access_token}`;
        processQueue(null, access_token);
        original.headers.Authorization = `Bearer ${access_token}`;
        return api(original);
      } catch (e) {
        processQueue(e, null);
        localStorage.clear();
        window.location.href = "/login";
        // Silent failure - handled by redirect
        return Promise.reject(e);
      } finally {
        isRefreshing = false;
      }
    }
    // Non-401 errors are now handled by the calling code/stores.
    // We remove the global showError which was causing disruptive modals for ignorable background tasks.
    if (error?.response?.status >= 500) {
      // eslint-disable-next-line no-console
      console.error("Server Error:", getErrorMessage(error));
    }

    return Promise.reject(error);
  },
);

// ─── Auth ─────────────────────────────────────────────────────────────────────
export const authAPI = {
  register: (data: { email: string; password: string; name: string }) =>
    api.post("/auth/register", data),
  login: (data: { email: string; password: string }) =>
    api.post("/auth/login", data),
  refresh: (refreshToken: string) =>
    api.post("/auth/refresh", { refresh_token: refreshToken }),
  logout: (refreshToken: string) =>
    api.post("/auth/logout", { refresh_token: refreshToken }),
  loginWithGoogle: (credential: string) =>
    api.post("/auth/google", { credential }),
};

// ─── User ─────────────────────────────────────────────────────────────────────
export const userAPI = {
  getMe: () => api.get("/users/me"),
  updateMe: (data: any) => api.patch("/users/me", data),
  deleteMe: () => api.delete("/users/me"),
};

// ─── Portfolio ────────────────────────────────────────────────────────────────
export const portfolioAPI = {
  list: () => api.get("/portfolios"),
  create: (data: any) => api.post("/portfolios", data),
  getById: (id: string) => api.get(`/portfolios/${id}`),
  update: (id: string, data: any) => api.patch(`/portfolios/${id}`, data),
  delete: (id: string) => api.delete(`/portfolios/${id}`),
  exportPdf: (id: string) => api.get(`/public/id/${id}/pdf`, { responseType: "blob" }),
  publish: (id: string) => api.post(`/portfolios/${id}/publish`),
  unpublish: (id: string) => api.post(`/portfolios/${id}/unpublish`),
  duplicate: (id: string) => api.post(`/portfolios/${id}/duplicate`),
};

// ─── Section ──────────────────────────────────────────────────────────────────
export const sectionAPI = {
  list: (portfolioId: string) => api.get(`/portfolios/${portfolioId}/sections`),
  create: (portfolioId: string, data: any) =>
    api.post(`/portfolios/${portfolioId}/sections`, data),
  update: (id: string, data: any) => api.patch(`/sections/${id}`, data),
  delete: (id: string) => api.delete(`/sections/${id}`),
  duplicate: (id: string) => api.post(`/sections/${id}/duplicate`),
  reorder: (order: { id: string; position: number }[]) =>
    api.post("/sections/reorder", { order }),
};

// ─── Analytics ────────────────────────────────────────────────────────────────
export const analyticsAPI = {
  summary: (portfolioId: string) =>
    api.get(`/analytics/${portfolioId}/summary`),
  visitors: (portfolioId: string) =>
    api.get(`/analytics/${portfolioId}/visitors`),
  sources: (portfolioId: string) =>
    api.get(`/analytics/${portfolioId}/sources`),
};

// ─── AI ───────────────────────────────────────────────────────────────────────
export const aiAPI = {
  improveText: (text: string, context: string) =>
    api.post("/ai/improve-text", { text, context }),
  scoreResume: (content: string, targetRole: string) =>
    api.post("/ai/score-resume", {
      portfolio_content: content,
      target_role: targetRole,
    }),
  suggestSkills: (role: string, currentSkills: string[]) =>
    api.post("/ai/suggest-skills", { role, current_skills: currentSkills }),
};

// ─── Upload ───────────────────────────────────────────────────────────────────
export const uploadAPI = {
  image: (file: File) => {
    const formData = new FormData();
    formData.append("image", file);
    return api.post("/upload", formData, {
      headers: { "Content-Type": "multipart/form-data" },
    });
  },
  file: (file: File) => {
    const formData = new FormData();
    formData.append("file", file);
    return api.post("/upload", formData, {
      headers: { "Content-Type": "multipart/form-data" },
    });
  },
};

// ─── Public ───────────────────────────────────────────────────────────────────
// ⚠️ ใช้ `api` instance เสมอ (ไม่ใช่ axios ตรงๆ)
// เพื่อให้ได้ baseURL + ngrok header อัตโนมัติ
export const publicAPI = {
  view: (slug: string, password?: string) =>
    api.get(`/public/p/${slug}`, {
      headers: password ? { "X-Portfolio-Password": password } : {},
    }),
  track: (slug: string, eventType: string, source?: string) =>
    api.post(`/public/p/${slug}/track`, { event_type: eventType, source }),
  chat: (slug: string, message: string, sessionId: string) =>
    api.post(`/public/p/${slug}/chat`, { message, session_id: sessionId }),
  exportPdf: (slug: string) =>
    api.get(`/public/p/${slug}/pdf`, { responseType: "blob" })
};

export default api;
