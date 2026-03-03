<template>
  <div class="dashboard-layout">
    <AppSidebar />
    <main class="dashboard-main">
      <div class="analytics-container">
        <div class="page-header">
          <div>
            <h1><BarChart2 :size="28" class="icon-inline" /> Analytics</h1>
            <p class="dash-sub">สถิติการเข้าชมและปฏิสัมพันธ์ของพอร์ตโฟลิโอนี้</p>
          </div>
          <RouterLink
            :to="`/portfolios/${route.params.id}/edit`"
            class="btn btn-outline btn-sm"
          >
            <ArrowLeft :size="16" /> Back to Builder
          </RouterLink>
        </div>

        <div v-if="loading" class="loading-grid">
          <div
            v-for="i in 4"
            :key="i"
            class="skeleton"
            style="height: 120px; border-radius: 20px"
          ></div>
        </div>

        <div v-else-if="summary" class="stats-grid">
          <div class="stat-card">
            <div class="stat-label">
              <Eye :size="14" class="icon-inline" /> Total Views
            </div>
            <div class="stat-num">{{ summary.total_views.toLocaleString() }}</div>
            <div class="stat-sub">
              {{ summary.today_views }} วันนี้ · {{ summary.week_views }} สัปดาห์นี้
            </div>
          </div>
          <div class="stat-card">
            <div class="stat-label">
              <Send :size="14" class="icon-inline" /> Hire Me Clicks
            </div>
            <div class="stat-num">{{ summary.hire_clicks }}</div>
            <div class="stat-sub">ความสนใจจ้างงาน</div>
          </div>
          <div class="stat-card">
            <div class="stat-label">
              <Bot :size="14" class="icon-inline" /> AI Chat Sessions
            </div>
            <div class="stat-num">{{ summary.ai_chat_count }}</div>
            <div class="stat-sub">การพูดคุยกับ AI</div>
          </div>
        </div>

        <div class="charts-layout">
          <!-- Daily Views Chart (simple bar) -->
          <div v-if="summary?.daily_views?.length" class="chart-card">
            <h3><LineChart :size="18" class="icon-inline" /> Views รายวัน (7 วันล่าสุด)</h3>
            <div class="bar-chart">
              <div v-for="(dv, i) in summary.daily_views" :key="i" class="bar-col">
                <div
                  class="bar-fill"
                  :style="{ height: barHeight(dv.count) + '%' }"
                  :title="`${dv.count} views`"
                ></div>
                <div class="bar-label">{{ dv.date.slice(5) }}</div>
              </div>
            </div>
          </div>

          <!-- Country stats -->
          <div v-if="summary?.country_stats?.length" class="country-card">
            <h3><Globe :size="18" class="icon-inline" /> ผู้เข้าชมจากประเทศต่างๆ</h3>
            <div class="country-list">
              <div
                v-for="c in summary.country_stats"
                :key="c.country_code"
                class="country-row"
              >
                <span class="country-flag">{{ countryFlag(c.country_code) }}</span>
                <span class="country-name">{{ c.country_code }}</span>
                <div class="country-bar-bg">
                  <div
                    class="country-bar-fill"
                    :style="{ width: (c.count / summary.total_views) * 100 + '%' }"
                  ></div>
                </div>
                <span class="country-count">{{ c.count }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { useRoute, RouterLink } from "vue-router";
import { analyticsAPI } from "@/api";
import type { AnalyticsSummary } from "@/types";
import AppSidebar from "@/components/layout/AppSidebar.vue";
import {
  BarChart2,
  Eye,
  FileDown,
  Bot,
  LineChart,
  Globe,
  ArrowLeft,
  Send,
} from "lucide-vue-next";

const route = useRoute();
const loading = ref(true);
const summary = ref<AnalyticsSummary | null>(null);

const maxViews = computed(() =>
  Math.max(...(summary.value?.daily_views?.map((d) => d.count) || [1])),
);
function barHeight(count: number) {
  return Math.max(4, (count / maxViews.value) * 100);
}
function countryFlag(code: string) {
  return code
    .toUpperCase()
    .replace(/./g, (c) => String.fromCodePoint(0x1f1e6 - 65 + c.charCodeAt(0)));
}

onMounted(async () => {
  try {
    const { data } = await analyticsAPI.summary(route.params.id as string);
    summary.value = data.data;
  } catch (e) {
    console.error("Failed to fetch analytics", e);
  } finally {
    loading.value = false;
  }
});
</script>

<style scoped>
.dashboard-main {
  padding: 36px;
  overflow-y: auto;
}

.analytics-container {
  max-width: 1000px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 32px;
}

.dash-sub {
  color: var(--muted);
  font-size: 14px;
}

h1 {
  font-size: 26px;
  font-weight: 800;
  margin-bottom: 4px;
}

.loading-grid,
.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
  margin-bottom: 24px;
}

.stat-card {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 20px;
  padding: 24px;
  transition: all 0.2s;
}

.stat-card:hover {
  transform: translateY(-2px);
  border-color: var(--indigo);
}

.stat-label {
  font-size: 13px;
  color: var(--muted);
  margin-bottom: 12px;
  font-weight: 600;
}

.stat-num {
  font-family: var(--font-display);
  font-size: 36px;
  font-weight: 800;
  color: var(--neon-cyan);
  margin-bottom: 4px;
}

.stat-sub {
  font-size: 12px;
  color: var(--muted);
}

.charts-layout {
  display: grid;
  grid-template-columns: 1.5fr 1fr;
  gap: 20px;
}

@media (max-width: 1024px) {
  .charts-layout {
    grid-template-columns: 1fr;
  }
}

.chart-card,
.country-card {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 20px;
  padding: 24px;
}

h3 {
  font-size: 16px;
  font-weight: 700;
  margin-bottom: 24px;
}

.bar-chart {
  display: flex;
  align-items: flex-end;
  gap: 10px;
  height: 160px;
  padding-bottom: 10px;
}

.bar-col {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  height: 100%;
  justify-content: flex-end;
}

.bar-fill {
  width: 100%;
  background: linear-gradient(180deg, var(--indigo), var(--purple));
  border-radius: 6px 6px 0 0;
  min-height: 4px;
  transition: height 0.6s cubic-bezier(0.34, 1.56, 0.64, 1);
}

.bar-label {
  font-size: 11px;
  color: var(--muted);
  font-weight: 500;
}

.country-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.country-row {
  display: flex;
  align-items: center;
  gap: 12px;
}

.country-flag {
  font-size: 20px;
  flex-shrink: 0;
}

.country-name {
  font-size: 13px;
  font-weight: 600;
  width: 32px;
  flex-shrink: 0;
}

.country-bar-bg {
  flex: 1;
  height: 8px;
  background: var(--bg2);
  border-radius: 4px;
  overflow: hidden;
}

.country-bar-fill {
  height: 100%;
  background: linear-gradient(90deg, var(--indigo), var(--neon-cyan));
  border-radius: 4px;
  transition: width 1s ease;
}

.country-count {
  font-size: 13px;
  font-weight: 700;
  color: var(--text);
  width: 30px;
  text-align: right;
  flex-shrink: 0;
}

.icon-inline {
  vertical-align: middle;
  margin-top: -3px;
}

.btn-outline {
  background: transparent;
  border: 1px solid var(--border);
  color: var(--muted);
}

.btn-outline:hover {
  background: var(--bg2);
  color: var(--text);
  border-color: var(--indigo);
}

@media (max-width: 768px) {
  .dashboard-main {
    padding: 20px 16px;
  }
}
</style>
