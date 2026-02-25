import { createRouter, createWebHistory } from "vue-router";
import { useAuthStore } from "@/stores/auth";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    // ─── Public Routes ────────────────────────────────────────────────────
    {
      path: "/",
      name: "landing",
      component: () => import("@/views/LandingView.vue"),
      meta: { guestOnly: true },
    },
    {
      path: "/login",
      name: "login",
      component: () => import("@/views/auth/LoginView.vue"),
      meta: { guestOnly: true },
    },
    {
      path: "/register",
      name: "register",
      component: () => import("@/views/auth/RegisterView.vue"),
      meta: { guestOnly: true },
    },

    // ─── Authenticated Routes ─────────────────────────────────────────────
    {
      path: "/dashboard",
      name: "dashboard",
      component: () => import("@/views/DashboardView.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/portfolios/new",
      name: "portfolio-new",
      component: () => import("@/views/portfolio/NewPortfolioView.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/portfolios/:id/edit",
      name: "portfolio-edit",
      component: () => import("@/views/portfolio/BuilderView.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/portfolios/:id/analytics",
      name: "portfolio-analytics",
      component: () => import("@/views/portfolio/AnalyticsView.vue"),
      meta: { requiresAuth: true },
    },
    {
      path: "/settings",
      name: "settings",
      component: () => import("@/views/SettingsView.vue"),
      meta: { requiresAuth: true },
    },

    // ─── Public Portfolio View (by slug) ──────────────────────────────────
    {
      path: "/p/:slug",
      name: "portfolio-public",
      component: () => import("@/views/portfolio/PublicPortfolioView.vue"),
    },

    // ─── 404 ──────────────────────────────────────────────────────────────
    {
      path: "/:pathMatch(.*)*",
      name: "not-found",
      component: () => import("@/views/NotFoundView.vue"),
    },
  ],
  scrollBehavior() {
    return { top: 0 };
  },
});

// ─── Navigation Guards ────────────────────────────────────────────────────────
router.beforeEach(async (to, _from, next) => {
  const authStore = useAuthStore();

  // Fetch user if authenticated but no user data yet
  if (authStore.isAuthenticated && !authStore.user) {
    await authStore.fetchMe();
  }

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    return next({ name: "login", query: { redirect: to.fullPath } });
  }

  if (to.meta.guestOnly && authStore.isAuthenticated) {
    return next({ name: "dashboard" });
  }

  next();
});

export default router;
