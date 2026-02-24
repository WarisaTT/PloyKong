<template>
  <div class="dashboard-layout">
    <!-- Sidebar -->
    <aside class="sidebar">
      <div class="sidebar-logo">
        <span class="logo-pk">PK</span>
        <span class="logo-text">PloyKong</span>
      </div>
      <nav class="sidebar-nav">
        <RouterLink to="/dashboard" class="nav-item" exact-active-class="active">
          <BarChart2 :size="16" /> Dashboard
        </RouterLink>
        <RouterLink to="/portfolios/new" class="nav-item" exact-active-class="active">
          <Plus :size="16" /> สร้างพอร์ตใหม่
        </RouterLink>
        <RouterLink to="/settings" class="nav-item" exact-active-class="active">
          <Settings :size="16" /> Settings
        </RouterLink>

        <div class="sidebar-theme-toggle">
          <div class="theme-label">Theme</div>
          <div class="theme-toggle-group">
            <button :class="['theme-btn', { active: themeStore.mode === 'dark' }]" @click="themeStore.setMode('dark')" title="Dark Mode"><Moon :size="16" /></button>
            <button :class="['theme-btn', { active: themeStore.mode === 'light' }]" @click="themeStore.setMode('light')" title="Light Mode"><Sun :size="16" /></button>
            <button :class="['theme-btn', { active: themeStore.mode === 'auto' }]" @click="themeStore.setMode('auto')" title="System Setting"><Monitor :size="16" /></button>
          </div>
        </div>
      </nav>
      <div class="sidebar-user" @click="handleLogout">
        <div class="user-avatar">{{ userInitial }}</div>
        <div>
          <div class="user-name">{{ authStore.user?.name }}</div>
          <div class="user-plan">{{ authStore.user?.plan }} plan</div>
        </div>
        <LogOut :size="16" class="logout-icon" title="Logout" />
      </div>
    </aside>

    <!-- Main Content -->
    <main class="dashboard-main">
      <!-- Header -->
      <div class="dash-header">
        <div>
          <h1 class="dash-greeting">Hi, {{ firstName }} ⭐️ </h1>
          <p class="dash-sub">จัดการพอร์ตโฟลิโอและดูสถิติได้ที่นี่</p>
        </div>
        <RouterLink to="/portfolios/new" class="btn btn-primary btn-icon-text">
          <Plus :size="16" /> สร้างพอร์ตใหม่
        </RouterLink>
      </div>

      <!-- Stats Row -->
      <div class="stats-grid">
        <StatCard
          label="Total Views"
          :icon="TrendingUp"
          :value="totalViews"
          change="+12%"
          color="indigo"
        />
        <StatCard
          label="PDF Downloads"
          :icon="FileDown"
          :value="totalPDF"
          change="+8%"
          color="cyan"
        />
        <StatCard
          label="Hire Me Clicks"
          :icon="Briefcase"
          :value="totalHire"
          change="+3 วันนี้"
          color="pink"
        />
        <StatCard
          label="Portfolios"
          :icon="LayoutGrid"
          :value="portfolioStore.portfolios.length"
          :change="`${activeCount} active`"
          color="green"
        />
      </div>

      <!-- Portfolio List -->
      <section class="section-card">
        <div class="section-card-header">
          <h2><Layers :size="20" class="icon-inline" /> My Portfolios</h2>
        </div>

        <div v-if="portfolioStore.loading" class="portfolio-loading">
          <div v-for="i in 3" :key="i" class="skeleton" style="height:72px;border-radius:12px;margin-bottom:8px"></div>
        </div>

        <div v-else-if="portfolioStore.portfolios.length === 0" class="portfolio-empty">
          <p>No portfolios found — <RouterLink to="/portfolios/new" class="link">Create one!</RouterLink></p>
        </div>

        <div v-else class="portfolio-list">
          <PortfolioListItem
            v-for="portfolio in portfolioStore.portfolios"
            :key="portfolio.id"
            :portfolio="portfolio"
            @delete="deletePortfolio(portfolio.id)"
          />
        </div>
      </section>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { usePortfolioStore } from '@/stores/portfolio'
import { useThemeStore } from '@/stores/theme'
import StatCard from '@/components/dashboard/StatCard.vue'
import PortfolioListItem from '@/components/dashboard/PortfolioListItem.vue'
import { BarChart2, Plus, Settings, Moon, Sun, Monitor, LogOut, LayoutGrid, TrendingUp, FileDown, Briefcase, Layers } from 'lucide-vue-next'
import Swal from 'sweetalert2'

const authStore = useAuthStore()
const portfolioStore = usePortfolioStore()
const themeStore = useThemeStore()
const router = useRouter()

const totalViews = ref(0)
const totalPDF = ref(0)
const totalHire = ref(0)

const firstName = computed(() => authStore.user?.name.split(' ')[0] || 'คุณ')
const userInitial = computed(() => authStore.user?.name?.[0]?.toUpperCase() || 'U')
const activeCount = computed(() => portfolioStore.portfolios.filter((p) => p.is_published).length)

onMounted(async () => {
  await portfolioStore.fetchPortfolios()
  // Sum stats from all portfolios
  portfolioStore.portfolios.forEach((p) => {
    totalViews.value += p.view_count || 0
  })
})

async function deletePortfolio(id: string) {
  const { isConfirmed } = await Swal.fire({
    title: 'Delete this portfolio?',
    icon: 'warning',
    showCancelButton: true,
    confirmButtonColor: '#ef4444',
    cancelButtonColor: '#94a3b8',
    confirmButtonText: 'Delete'
  })
  if (isConfirmed) {
    await portfolioStore.deletePortfolio(id)
  }
}

async function handleLogout() {
  await authStore.logout()
  router.push('/login')
}
</script>

<style scoped>
.dashboard-layout {
  display: grid;
  grid-template-columns: 220px 1fr;
  min-height: 100vh;
  background: var(--bg);
}

/* Sidebar */
.sidebar {
  background: var(--sidebar-bg);
  border-right: 1px solid var(--border);
  padding: 24px 16px;
  display: flex; flex-direction: column;
  position: sticky; top: 0; height: 100vh;
}
.sidebar-logo {
  display: flex; align-items: center; gap: 10px;
  margin-bottom: 36px; padding: 0 8px;
}
.logo-pk {
  width: 36px; height: 36px; border-radius: 10px;
  background: linear-gradient(135deg, var(--indigo), var(--purple));
  display: flex; align-items: center; justify-content: center;
  font-family: var(--font-display); font-size: 14px; font-weight: 800; color: #fff;
}
.logo-text { font-family: var(--font-display); font-size: 17px; font-weight: 800; }

.sidebar-nav { flex: 1; display: flex; flex-direction: column; gap: 3px; }
.nav-item {
  display: flex; align-items: center; gap: 10px;
  padding: 10px 12px; border-radius: 10px;
  text-decoration: none; color: var(--muted);
  font-size: 14px; font-weight: 500; transition: all 0.15s;
}
.nav-item:hover, .nav-item.active {
  background: rgba(79, 70, 229, 0.1);
  color: var(--indigo);
}

.sidebar-theme-toggle {
  margin-top: auto;
  padding: 12px;
  border-top: 1px solid var(--border);
  display: flex;
  flex-direction: column;
  gap: 8px;
}
.theme-label {
  font-size: 11px;
  font-weight: 600;
  color: var(--muted);
  text-transform: uppercase;
  letter-spacing: 0.5px;
}
.theme-toggle-group {
  display: flex;
  background: var(--bg2);
  border-radius: 8px;
  padding: 4px;
  gap: 4px;
}
.theme-btn {
  flex: 1;
  background: transparent;
  border: none;
  color: var(--muted);
  border-radius: 6px;
  padding: 6px 0;
  cursor: pointer;
  transition: all 0.2s;
}
.theme-btn:hover { background: var(--surface); color: var(--text); }
.theme-btn.active {
  background: var(--surface);
  color: var(--indigo);
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.sidebar-user {
  display: flex; align-items: center; gap: 10px;
  padding: 12px; border-radius: 12px;
  border: 1px solid var(--border); cursor: pointer;
  transition: all 0.15s; margin-top: 16px;
  background: var(--surface);
}
.sidebar-user:hover { background: var(--bg2); }
.user-avatar {
  width: 36px; height: 36px; border-radius: 50%;
  background: linear-gradient(135deg, var(--indigo), var(--purple));
  display: flex; align-items: center; justify-content: center;
  font-size: 15px; font-weight: 700; color: #fff; flex-shrink: 0;
}
.user-name { font-size: 13px; font-weight: 600; }
.user-plan { font-size: 11px; color: var(--neon-purple); }
.logout-icon { margin-left: auto; color: var(--muted); font-size: 16px; }

/* Main */
.dashboard-main { padding: 36px; overflow-y: auto; }
.dash-header { display: flex; justify-content: space-between; align-items: flex-start; margin-bottom: 28px; }
.dash-greeting { font-size: 26px; font-weight: 800; margin-bottom: 4px; }
.dash-sub { color: var(--muted); font-size: 14px; }

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px; margin-bottom: 28px;
}

.section-card {
  background: var(--surface); border: 1px solid var(--border);
  border-radius: 20px; overflow: hidden;
}
.section-card-header {
  padding: 20px 24px;
  border-bottom: 1px solid var(--border);
}
.section-card-header h2 { font-size: 16px; font-weight: 700; }

.portfolio-loading, .portfolio-empty { padding: 24px; }
.portfolio-empty { color: var(--muted); font-size: 14px; }
.link { color: var(--neon-purple); text-decoration: none; }
.portfolio-list { padding: 12px 16px; display: flex; flex-direction: column; gap: 10px; }

.icon-inline { vertical-align: text-bottom; }
.btn-icon-text { display: inline-flex; align-items: center; justify-content: center; gap: 8px; }

@media (max-width: 768px) {
  .dashboard-layout { grid-template-columns: 1fr; }
  .sidebar { display: none; }
  .dashboard-main { padding: 16px; }
}
</style>
