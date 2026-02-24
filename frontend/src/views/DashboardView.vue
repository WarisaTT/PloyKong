<template>
  <div class="dashboard-layout">
    <!-- Sidebar -->
    <aside class="sidebar">
      <div class="sidebar-logo">
        <span class="logo-pk">PK</span>
        <span class="logo-text">PloyKong</span>
      </div>
      <nav class="sidebar-nav">
        <RouterLink to="/dashboard" class="nav-item active">
          <span>📊</span> Dashboard
        </RouterLink>
        <RouterLink to="/portfolios/new" class="nav-item">
          <span>➕</span> สร้างพอร์ตใหม่
        </RouterLink>
        <RouterLink to="/settings" class="nav-item">
          <span>⚙️</span> Settings
        </RouterLink>
      </nav>
      <div class="sidebar-user" @click="authStore.logout">
        <div class="user-avatar">{{ userInitial }}</div>
        <div>
          <div class="user-name">{{ authStore.user?.name }}</div>
          <div class="user-plan">{{ authStore.user?.plan }} plan</div>
        </div>
        <span class="logout-icon" title="Logout">↩</span>
      </div>
    </aside>

    <!-- Main Content -->
    <main class="dashboard-main">
      <!-- Header -->
      <div class="dash-header">
        <div>
          <h1 class="dash-greeting">สวัสดี, {{ firstName }} 👋</h1>
          <p class="dash-sub">จัดการพอร์ตโฟลิโอและดูสถิติได้ที่นี่</p>
        </div>
        <RouterLink to="/portfolios/new" class="btn btn-primary">
          ➕ สร้างพอร์ตใหม่
        </RouterLink>
      </div>

      <!-- Stats Row -->
      <div class="stats-grid">
        <StatCard
          label="📈 Total Views"
          :value="totalViews"
          change="+12%"
          color="indigo"
        />
        <StatCard
          label="📄 PDF Downloads"
          :value="totalPDF"
          change="+8%"
          color="cyan"
        />
        <StatCard
          label="🎯 Hire Me Clicks"
          :value="totalHire"
          change="+3 วันนี้"
          color="pink"
        />
        <StatCard
          label="🗂️ Portfolios"
          :value="portfolioStore.portfolios.length"
          :change="`${activeCount} active`"
          color="green"
        />
      </div>

      <!-- Portfolio List -->
      <section class="section-card">
        <div class="section-card-header">
          <h2>🗂️ พอร์ตโฟลิโอของฉัน</h2>
        </div>

        <div v-if="portfolioStore.loading" class="portfolio-loading">
          <div v-for="i in 3" :key="i" class="skeleton" style="height:72px;border-radius:12px;margin-bottom:8px"></div>
        </div>

        <div v-else-if="portfolioStore.portfolios.length === 0" class="portfolio-empty">
          <p>ยังไม่มีพอร์ตโฟลิโอ — <RouterLink to="/portfolios/new" class="link">สร้างเลย!</RouterLink></p>
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
import { RouterLink } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { usePortfolioStore } from '@/stores/portfolio'
import StatCard from '@/components/dashboard/StatCard.vue'
import PortfolioListItem from '@/components/dashboard/PortfolioListItem.vue'

const authStore = useAuthStore()
const portfolioStore = usePortfolioStore()

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
  if (confirm('ลบพอร์ตโฟลิโอนี้หรือไม่?')) {
    await portfolioStore.deletePortfolio(id)
  }
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
  background: rgba(8, 13, 31, 0.95);
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
  background: rgba(79, 70, 229, 0.15);
  color: #fff;
}

.sidebar-user {
  display: flex; align-items: center; gap: 10px;
  padding: 12px; border-radius: 12px;
  border: 1px solid var(--border); cursor: pointer;
  transition: all 0.15s; margin-top: 16px;
}
.sidebar-user:hover { background: rgba(255,255,255,0.05); }
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

@media (max-width: 768px) {
  .dashboard-layout { grid-template-columns: 1fr; }
  .sidebar { display: none; }
  .dashboard-main { padding: 16px; }
}
</style>
