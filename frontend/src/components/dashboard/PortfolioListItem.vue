<template>
  <div class="portfolio-item">
    <div class="portfolio-info">
      <div
        class="portfolio-dot"
        :class="{ live: portfolio.is_published }"
      ></div>
      <div class="portfolio-text">
        <div class="portfolio-title">{{ portfolio.title }}</div>
        <div class="portfolio-slug">{{ portfolio.slug }}</div>
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
      <button
        class="btn btn-secondary btn-sm"
        title="Duplicate Portfolio"
        @click="$emit('duplicate')"
      >
        <Files :size="16" />
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
      <!-- View Public Link Button (SEO Friendly) -->
      <a
        v-if="portfolio.is_published"
        :href="publicUrl"
        target="_blank"
        class="btn btn-secondary btn-sm"
        title="View Public Link"
      >
        <ExternalLink :size="16" />
      </a>
      <button
        v-if="portfolio.is_published"
        class="btn btn-secondary btn-sm"
        title="Copy SEO Link"
        @click="copyLink"
      >
        <Check v-if="isCopied" :size="16" class="text-success" />
        <Link v-else :size="16" />
      </button>
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
import { ref, computed } from "vue";
import { RouterLink } from "vue-router";
import type { Portfolio } from "@/types";
import { portfolioAPI } from "@/api";
import { showSuccess, toastError } from "@/utils/alert";
import {
  Eye,
  EyeOff,
  Pencil,
  BarChart2,
  ExternalLink,
  Trash2,
  FileDown,
  Loader2,
  Files,
  Link,
  Check,
} from "lucide-vue-next";

const props = defineProps<{ portfolio: Portfolio }>();
defineEmits<{ delete: []; duplicate: [] }>();

const isDownloading = ref(false);
const isCopied = ref(false);

const publicUrl = computed(() => {
  const apiBase = import.meta.env.VITE_API_URL || '/api/v1'
  return `${apiBase}/public/p/${props.portfolio.slug}`
});

function copyLink() {
  navigator.clipboard.writeText(publicUrl.value).then(() => {
    isCopied.value = true;
    setTimeout(() => {
      isCopied.value = false;
    }, 2000);
  });
}

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
    toastError("Failed to generate PDF. Please try again.");
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
  padding: 16px;
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 12px;
  gap: 16px;
  transition: all 0.2s ease;
}

.portfolio-item:hover {
  border-color: var(--primary);
  background: var(--bg2);
  transform: translateY(-2px);
  box-shadow: var(--shadow-md);
}

.portfolio-info {
  display: flex;
  align-items: center;
  gap: 12px;
  flex: 1;
  min-width: 0;
}

.portfolio-dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: var(--muted);
  flex-shrink: 0;
}

.portfolio-dot.live {
  background: var(--success);
  box-shadow: 0 0 8px var(--success);
}

.portfolio-text {
  min-width: 0;
  flex: 1;
}

.portfolio-title {
  font-size: 15px;
  font-weight: 700;
  margin-bottom: 2px;
  white-space: nowrap;
  overflow: hidden;
  text-anchor: middle;
  text-overflow: ellipsis;
}

.portfolio-slug {
  font-size: 12px;
  color: var(--muted);
  opacity: 0.8;
}

.portfolio-stats {
  display: flex;
  align-items: center;
  gap: 12px;
}

.stat-chip {
  font-size: 12px;
  color: var(--text-2);
  display: inline-flex;
  align-items: center;
  gap: 6px;
  background: var(--surface-2);
  padding: 4px 10px;
  border-radius: 8px;
}

.portfolio-actions {
  display: flex;
  gap: 6px;
  flex-wrap: wrap;
  justify-content: flex-end;
}

.portfolio-actions .btn {
  padding: 8px 10px;
  border-radius: 8px;
}

.icon-btn-labeled {
  padding-left: 14px !important;
  padding-right: 14px !important;
}

@media (max-width: 768px) {
  .portfolio-item {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
  }
  .portfolio-info {
    width: 100%;
  }
  .portfolio-stats {
    width: 100%;
    justify-content: flex-start;
  }
  .portfolio-actions {
    width: 100%;
    justify-content: flex-start;
    border-top: 1px solid var(--border);
    padding-top: 12px;
  }
}
</style>
