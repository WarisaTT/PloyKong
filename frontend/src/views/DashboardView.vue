<template>
  <div class="dashboard-layout">
    <AppSidebar />

    <!-- Main Content -->
    <main class="dashboard-main">
      <!-- Header -->
      <div class="dash-header">
        <div>
          <h1 class="dash-greeting">Hi, {{ firstName }} ⭐️</h1>
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
          :change="todayViews > 0 ? `+${todayViews} วันนี้` : '0 วันนี้'"
          color="indigo"
        />
        <StatCard
          label="Hire Me Clicks"
          :icon="Briefcase"
          :value="totalHire"
          change=""
          color="pink"
        />
        <StatCard
          label="AI Answered"
          :icon="Sparkles"
          :value="totalAIChat"
          change="จากทุกพอร์ต"
          color="purple"
        />
        <StatCard
          label="Knowledge Gaps"
          :icon="Lightbulb"
          :value="gapCount"
          :change="gapCount > 0 ? 'รอการสอนเพิ่ม' : 'ครบถ้วน'"
          color="amber"
          @click="router.push('/ai-center')"
          style="cursor: pointer"
        />
        <StatCard
          label="Active Portfolios"
          :icon="LayoutGrid"
          :value="portfolioStore.portfolios.length"
          :change="`${activeCount} published`"
          color="green"
        />
      </div>

      <!-- Portfolio List -->
      <section class="section-card">
        <div class="section-card-header">
          <h2><Layers :size="20" class="icon-inline" /> My Portfolios</h2>
        </div>

        <div v-if="portfolioStore.loading" class="portfolio-loading">
          <div
            v-for="i in 3"
            :key="i"
            class="skeleton"
            style="height: 72px; border-radius: 12px; margin-bottom: 8px"
          ></div>
        </div>

        <div
          v-else-if="portfolioStore.portfolios.length === 0"
          class="portfolio-empty"
        >
          <p>
            No portfolios found —
            <RouterLink to="/portfolios/new" class="link"
              >Create one!</RouterLink
            >
          </p>
        </div>

        <TransitionGroup tag="div" name="list" v-else class="portfolio-list">
          <PortfolioListItem
            v-for="portfolio in portfolioStore.portfolios"
            :key="portfolio.id"
            :portfolio="portfolio"
            @delete="deletePortfolio(portfolio.id)"
            @duplicate="duplicatePortfolio(portfolio.id)"
          />
        </TransitionGroup>
      </section>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { RouterLink, useRouter } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { usePortfolioStore } from "@/stores/portfolio";
import { useThemeStore } from "@/stores/theme";
import StatCard from "@/components/dashboard/StatCard.vue";
import PortfolioListItem from "@/components/dashboard/PortfolioListItem.vue";
import {
  BarChart2,
  Plus,
  Settings,
  Moon,
  Sun,
  Monitor,
  LogOut,
  LayoutGrid,
  TrendingUp,
  FileDown,
  Briefcase,
  Layers,
  Sparkles,
  Lightbulb,
} from "lucide-vue-next";
import Swal from "sweetalert2";
import { portfolioAPI, analyticsAPI } from "@/api";
import { showSuccess, toastError } from "@/utils/alert";
import AppSidebar from "@/components/layout/AppSidebar.vue";

const authStore = useAuthStore();
const portfolioStore = usePortfolioStore();
const themeStore = useThemeStore();
const router = useRouter();

const totalViews = ref(0);
const todayViews = ref(0);
const totalPDF = ref(0);
const totalHire = ref(0);
const totalAIChat = ref(0);
const gapCount = computed(() => portfolioStore.gapCount);

const firstName = computed(() => authStore.user?.name.split(" ")[0] || "คุณ");
const activeCount = computed(
  () => portfolioStore.portfolios.filter((p) => p.is_published).length,
);

onMounted(async () => {
  await portfolioStore.fetchPortfolios();
  // Sum basic stats
  portfolioStore.portfolios.forEach((p) => {
    totalViews.value += p.view_count || 0;
  });

  // Fetch gap count for the stat card
  portfolioStore.fetchGapCount();

  // Fetch detailed analytics (PDFs, Hire Me)
  await Promise.allSettled(
    portfolioStore.portfolios.map(async (p) => {
      try {
        const res = await analyticsAPI.summary(p.id);
        if (res.data?.success && res.data?.data) {
          totalPDF.value += res.data.data.pdf_downloads || 0;
          totalHire.value += res.data.data.hire_clicks || 0;
          todayViews.value += res.data.data.today_views || 0;
          totalAIChat.value += res.data.data.ai_chat_count || 0;
        }
      } catch (e) {
        console.error("Failed to load analytics for", p.id);
      }
    }),
  );
});

async function deletePortfolio(id: string) {
  const { isConfirmed } = await Swal.fire({
    title: "Delete this portfolio?",
    icon: "warning",
    showCancelButton: true,
    confirmButtonColor: "#ef4444",
    cancelButtonColor: "#94a3b8",
    confirmButtonText: "Delete",
  });
  if (isConfirmed) {
    await portfolioStore.deletePortfolio(id);
  }
}

async function duplicatePortfolio(id: string) {
  try {
    await portfolioStore.duplicatePortfolio(id);
    showSuccess("Portfolio duplicated successfully!");
  } catch (e: any) {
    toastError("Failed to duplicate portfolio");
  }
}

async function handleLogout() {
  await authStore.logout();
  router.push("/login");
}
</script>

<style scoped>
/* Main */
.dashboard-main {
  padding: 36px;
  overflow-y: auto;
}
.dash-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 28px;
}
.dash-greeting {
  font-size: 26px;
  font-weight: 800;
  margin-bottom: 4px;
}
.dash-sub {
  color: var(--muted);
  font-size: 14px;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
  margin-bottom: 28px;
}

.section-card {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 20px;
  overflow: hidden;
}
.section-card-header {
  padding: 20px 24px;
  border-bottom: 1px solid var(--border);
}
.section-card-header h2 {
  font-size: 16px;
  font-weight: 700;
}

.portfolio-loading,
.portfolio-empty {
  padding: 24px;
}
.portfolio-empty {
  color: var(--muted);
  font-size: 14px;
}
.link {
  color: var(--neon-purple);
  text-decoration: none;
}
.portfolio-list {
  padding: 12px 16px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

/* List Transitions */
.list-move,
.list-enter-active,
.list-leave-active {
  transition: all 0.4s ease;
}
.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateX(-20px);
}
.list-leave-active {
  position: absolute;
  width: calc(100% - 32px); /* keep width stable while absolute */
}

.icon-inline {
  vertical-align: text-bottom;
}
.btn-icon-text {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

@media (max-width: 768px) {
  .logout-btn {
    color: var(--danger, #f87171);
  }
  .dashboard-main {
    padding: 20px 16px;
  }
  .dash-header {
    flex-direction: column;
    gap: 16px;
  }
  .dash-header .btn {
    width: 100%;
  }
}

@media (min-width: 769px) {
  .mobile-nav {
    display: none;
  }
}
</style>
