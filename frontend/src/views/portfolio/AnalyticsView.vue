<!-- AnalyticsView.vue -->
<template>
  <div class="analytics-page">
    <div class="page-header">
      <RouterLink :to="`/portfolios/${route.params.id}/edit`" class="back-btn">← กลับ Builder</RouterLink>
      <h1>📊 Analytics</h1>
    </div>

    <div v-if="loading" class="loading-grid">
      <div v-for="i in 4" :key="i" class="skeleton" style="height:100px;border-radius:16px"></div>
    </div>

    <div v-else-if="summary" class="stats-grid">
      <div class="stat-card">
        <div class="stat-label">👁️ Total Views</div>
        <div class="stat-num">{{ summary.total_views.toLocaleString() }}</div>
        <div class="stat-sub">{{ summary.today_views }} วันนี้ · {{ summary.week_views }} สัปดาห์นี้</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">📄 PDF Downloads</div>
        <div class="stat-num">{{ summary.pdf_downloads }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">🎯 Hire Me Clicks</div>
        <div class="stat-num">{{ summary.hire_clicks }}</div>
      </div>
      <div class="stat-card">
        <div class="stat-label">🤖 AI Chat Sessions</div>
        <div class="stat-num">{{ summary.ai_chat_count }}</div>
      </div>
    </div>

    <!-- Daily Views Chart (simple bar) -->
    <div v-if="summary?.daily_views?.length" class="chart-card">
      <h3>📈 Views รายวัน (7 วัน)</h3>
      <div class="bar-chart">
        <div v-for="(dv, i) in summary.daily_views" :key="i" class="bar-col">
          <div class="bar-fill" :style="{ height: barHeight(dv.count) + '%' }" :title="`${dv.count} views`"></div>
          <div class="bar-label">{{ dv.date.slice(5) }}</div>
        </div>
      </div>
    </div>

    <!-- Country stats -->
    <div v-if="summary?.country_stats?.length" class="country-card">
      <h3>🌍 ผู้เข้าชมจากประเทศ</h3>
      <div class="country-list">
        <div v-for="c in summary.country_stats" :key="c.country_code" class="country-row">
          <span class="country-flag">{{ countryFlag(c.country_code) }}</span>
          <span class="country-name">{{ c.country_code }}</span>
          <div class="country-bar-bg">
            <div class="country-bar-fill" :style="{ width: (c.count / summary.total_views * 100) + '%' }"></div>
          </div>
          <span class="country-count">{{ c.count }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute, RouterLink } from 'vue-router'
import { analyticsAPI } from '@/api'
import type { AnalyticsSummary } from '@/types'

const route = useRoute()
const loading = ref(true)
const summary = ref<AnalyticsSummary | null>(null)

const maxViews = computed(() => Math.max(...(summary.value?.daily_views?.map(d => d.count) || [1])))
function barHeight(count: number) { return Math.max(4, (count / maxViews.value) * 100) }
function countryFlag(code: string) {
  return code.toUpperCase().replace(/./g, c => String.fromCodePoint(0x1F1E6 - 65 + c.charCodeAt(0)))
}

onMounted(async () => {
  const { data } = await analyticsAPI.summary(route.params.id as string)
  summary.value = data.data
  loading.value = false
})
</script>

<style scoped>
.analytics-page { max-width: 900px; margin: 0 auto; padding: 36px 24px; }
.page-header { display: flex; align-items: center; gap: 16px; margin-bottom: 32px; }
.back-btn { color: var(--muted); text-decoration: none; font-size: 14px; }
h1 { font-size: 24px; font-weight: 800; }
.loading-grid, .stats-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(180px, 1fr)); gap: 16px; margin-bottom: 24px; }
.stat-card { background: var(--surface); border: 1px solid var(--border); border-radius: 16px; padding: 22px; }
.stat-label { font-size: 13px; color: var(--muted); margin-bottom: 8px; font-weight: 600; }
.stat-num { font-family: var(--font-display); font-size: 36px; font-weight: 800; color: var(--neon-cyan); margin-bottom: 4px; }
.stat-sub { font-size: 12px; color: var(--muted); }
.chart-card, .country-card { background: var(--surface); border: 1px solid var(--border); border-radius: 20px; padding: 24px; margin-bottom: 20px; }
h3 { font-size: 16px; font-weight: 700; margin-bottom: 20px; }
.bar-chart { display: flex; align-items: flex-end; gap: 8px; height: 120px; }
.bar-col { flex: 1; display: flex; flex-direction: column; align-items: center; gap: 6px; height: 100%; justify-content: flex-end; }
.bar-fill { width: 100%; background: linear-gradient(180deg, var(--indigo), var(--purple)); border-radius: 4px 4px 0 0; min-height: 4px; transition: height 0.5s ease; }
.bar-label { font-size: 10px; color: var(--muted); }
.country-list { display: flex; flex-direction: column; gap: 10px; }
.country-row { display: flex; align-items: center; gap: 10px; }
.country-flag { font-size: 18px; flex-shrink: 0; }
.country-name { font-size: 13px; font-weight: 600; width: 40px; flex-shrink: 0; }
.country-bar-bg { flex: 1; height: 6px; background: rgba(255,255,255,0.08); border-radius: 3px; overflow: hidden; }
.country-bar-fill { height: 100%; background: linear-gradient(90deg, var(--indigo), var(--neon-cyan)); border-radius: 3px; }
.country-count { font-size: 12px; color: var(--muted); width: 30px; text-align: right; flex-shrink: 0; }
</style>
