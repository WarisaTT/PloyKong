<template>
  <div class="portfolio-item">
    <div class="portfolio-info">
      <div class="portfolio-dot" :class="{ live: portfolio.is_published }"></div>
      <div>
        <div class="portfolio-title">{{ portfolio.title }}</div>
        <div class="portfolio-slug">{{ portfolio.slug }}.ploykong.com</div>
      </div>
    </div>

    <div class="portfolio-stats">
      <span class="stat-chip">👁️ {{ portfolio.view_count.toLocaleString() }}</span>
      <span class="badge" :class="portfolio.is_published ? 'badge-live' : 'badge-draft'">
        {{ portfolio.is_published ? '● Live' : '○ Draft' }}
      </span>
    </div>

    <div class="portfolio-actions">
      <RouterLink :to="`/portfolios/${portfolio.id}/edit`" class="btn btn-secondary btn-sm">
        ✏️ Edit
      </RouterLink>
      <RouterLink :to="`/portfolios/${portfolio.id}/analytics`" class="btn btn-secondary btn-sm">
        📊
      </RouterLink>
      <a :href="`/p/${portfolio.slug}`" target="_blank" class="btn btn-secondary btn-sm">
        🔗
      </a>
      <button class="btn btn-danger btn-sm" @click="$emit('delete')">🗑️</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
import type { Portfolio } from '@/types'

defineProps<{ portfolio: Portfolio }>()
defineEmits<{ delete: [] }>()
</script>

<style scoped>
.portfolio-item {
  display: flex; align-items: center; justify-content: space-between;
  padding: 14px 16px; background: rgba(255,255,255,0.02);
  border: 1px solid var(--border); border-radius: 12px;
  gap: 16px; flex-wrap: wrap; transition: all 0.15s;
}
.portfolio-item:hover { border-color: rgba(79,70,229,0.3); background: rgba(79,70,229,0.04); }

.portfolio-info { display: flex; align-items: center; gap: 12px; flex: 1; min-width: 160px; }
.portfolio-dot {
  width: 8px; height: 8px; border-radius: 50%;
  background: var(--muted); flex-shrink: 0;
}
.portfolio-dot.live { background: var(--success); box-shadow: 0 0 6px var(--success); }
.portfolio-title { font-size: 14px; font-weight: 700; margin-bottom: 2px; }
.portfolio-slug { font-size: 11px; color: var(--neon-cyan); }

.portfolio-stats { display: flex; align-items: center; gap: 10px; }
.stat-chip { font-size: 12px; color: var(--muted); }

.portfolio-actions { display: flex; gap: 6px; }
</style>
