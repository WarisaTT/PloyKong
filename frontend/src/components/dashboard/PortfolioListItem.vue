<template>
  <div class="portfolio-item">
    <div class="portfolio-info">
      <div
        class="portfolio-dot"
        :class="{ live: portfolio.is_published }"
      ></div>
      <div>
        <div class="portfolio-title">{{ portfolio.title }}</div>
        <div class="portfolio-slug">{{ portfolio.slug }}.ploykong.com</div>
      </div>
    </div>

    <div class="portfolio-stats">
      <span class="stat-chip"
        ><Eye :size="14" class="icon" />
        {{ portfolio.view_count.toLocaleString() }}</span
      >
      <span
        class="badge"
        :class="portfolio.is_published ? 'badge-live' : 'badge-draft'"
      >
        {{ portfolio.is_published ? "● Live" : "○ Draft" }}
      </span>
    </div>

    <div class="portfolio-actions">
      <button
        class="btn btn-secondary btn-sm"
        title="Generate PDF"
        @click="downloadPdf"
        :disabled="isDownloading"
      >
        <Loader2 v-if="isDownloading" class="spin" :size="16" />
        <FileDown v-else :size="16" />
      </button>
      <RouterLink
        :to="`/portfolios/${portfolio.id}/edit`"
        class="btn btn-secondary btn-sm icon-btn-labeled"
      >
        <Pencil :size="14" /> Edit
      </RouterLink>
      <RouterLink
        :to="`/portfolios/${portfolio.id}/analytics`"
        class="btn btn-secondary btn-sm"
        title="Analytics"
      >
        <BarChart2 :size="16" />
      </RouterLink>
      <!-- View Public Link Button -->
      <a
        v-if="portfolio.is_published"
        :href="`/p/${portfolio.slug}`"
        target="_blank"
        class="btn btn-secondary btn-sm"
        title="View Public Link"
      >
        <ExternalLink :size="16" />
      </a>
      <button
        v-else
        disabled
        class="btn btn-secondary btn-sm"
        title="Portfolio is unpublished"
      >
        <EyeOff :size="16" />
      </button>
      <button
        class="btn btn-danger btn-sm"
        @click="$emit('delete')"
        title="Delete"
      >
        <Trash2 :size="16" />
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { RouterLink } from "vue-router";
import type { Portfolio } from "@/types";
import { portfolioAPI } from "@/api";
import { showError, showSuccess } from "@/utils/alert";
import {
  Eye,
  EyeOff,
  Pencil,
  BarChart2,
  ExternalLink,
  Trash2,
  FileDown,
  Loader2,
} from "lucide-vue-next";

const props = defineProps<{ portfolio: Portfolio }>();
defineEmits<{ delete: [] }>();

const isDownloading = ref(false);

async function downloadPdf() {
  if (isDownloading.value) return;
  isDownloading.value = true;
  try {
    const response = await portfolioAPI.exportPdf(props.portfolio.id);
    const url = window.URL.createObjectURL(response.data as Blob);
    const link = document.createElement("a");
    link.href = url;
    link.setAttribute("download", `${props.portfolio.slug}-resume.pdf`);
    document.body.appendChild(link);
    link.click();
    if (link.parentNode) link.parentNode.removeChild(link);
    showSuccess("PDF Downloaded successfully!");
  } catch (err: any) {
    console.error("Failed to download PDF", err);
    showError("Failed to generate PDF. Please try again.");
  } finally {
    isDownloading.value = false;
  }
}
</script>

<style scoped>
.portfolio-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 14px 16px;
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 12px;
  gap: 16px;
  flex-wrap: wrap;
  transition: all 0.15s;
}
.portfolio-item:hover {
  border-color: rgba(79, 70, 229, 0.3);
  background: var(--bg2);
}

.portfolio-info {
  display: flex;
  align-items: center;
  gap: 12px;
  flex: 1;
  min-width: 160px;
}
.portfolio-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: var(--muted);
  flex-shrink: 0;
}
.portfolio-dot.live {
  background: var(--success);
  box-shadow: 0 0 6px var(--success);
}
.portfolio-title {
  font-size: 14px;
  font-weight: 700;
  margin-bottom: 2px;
}
.portfolio-slug {
  font-size: 11px;
  color: var(--neon-cyan);
}

.portfolio-stats {
  display: flex;
  align-items: center;
  gap: 10px;
}
.stat-chip {
  font-size: 12px;
  color: var(--muted);
  display: inline-flex;
  align-items: center;
  gap: 4px;
}
.stat-chip .icon {
  opacity: 0.7;
}

.portfolio-actions {
  display: flex;
  gap: 6px;
}
.portfolio-actions .btn {
  padding: 6px 10px;
}
.icon-btn-labeled {
  display: inline-flex;
  align-items: center;
  gap: 6px;
}
</style>
