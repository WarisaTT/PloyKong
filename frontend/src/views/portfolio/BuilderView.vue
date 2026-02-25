<template>
  <div class="builder-layout">
    <!-- ─── Top Bar ─────────────────────────────────────────────────────── -->
    <header class="builder-topbar">
      <RouterLink to="/dashboard" class="back-btn" title="Dashboard">
        <ArrowLeft class="icon" :size="20" />
      </RouterLink>

      <div class="topbar-center">
        <span class="portfolio-slug"
          ><Globe :size="14" class="icon-inline" />
          {{ store.activePortfolio?.slug }}.ploykong.com</span
        >
        <span v-if="store.saving" class="saving-indicator"
          ><Save :size="12" class="icon-inline" /> Saving...</span
        >
      </div>

      <div class="topbar-actions">
        <button
          class="btn btn-secondary btn-sm btn-icon-text"
          :disabled="!store.activePortfolio?.is_published"
          :title="
            !store.activePortfolio?.is_published
              ? 'ไม่สามารถเปิดดูได้เว้นแต่จะ Publish'
              : 'Preview'
          "
          @click="previewPortfolio"
        >
          <Eye v-if="store.activePortfolio?.is_published" :size="14" />
          <EyeOff v-else :size="14" />
          Preview
        </button>
        <button
          class="btn btn-secondary btn-sm btn-icon-text"
          :disabled="!store.activePortfolio?.id || isExporting"
          title="Export as PDF"
          @click="exportPDF"
        >
          <Download :size="14" />
          {{ isExporting ? 'Exporting...' : 'Export' }}
        </button>
        <button
          class="btn btn-sm btn-icon-text publish-btn"
          :class="
            store.activePortfolio?.is_published ? 'btn-danger' : 'btn-primary'
          "
          @click="togglePublish"
        >
          <Play v-if="store.activePortfolio?.is_published" :size="14" />
          <Zap v-else :size="14" />
          {{ store.activePortfolio?.is_published ? 'Unpublish' : 'Publish' }}
        </button>
      </div>
    </header>

    <!-- ─── Main Builder Area ──────────────────────────────────────────── -->
    <div class="builder-main">
      <!-- Left: Block Palette -->
      <aside class="block-palette">
        <div class="palette-header">
          <span class="palette-title"
            ><Blocks :size="14" class="icon-inline" /> Blocks</span
          >
          <span class="palette-hint">Click to add</span>
        </div>
        <div class="palette-list">
          <div
            v-for="block in BLOCK_TYPES"
            :key="block.type"
            class="palette-item"
            @click="addBlock(block.type)"
          >
            <component :is="block.icon" class="palette-icon" :size="18" />
            <div>
              <div class="palette-name">{{ block.label }}</div>
              <div class="palette-desc">{{ block.description }}</div>
            </div>
            <span class="palette-add">+</span>
          </div>
        </div>
      </aside>

      <!-- Center: Canvas -->
      <main class="builder-canvas" ref="canvasRef">
        <div v-if="store.loading" class="canvas-loading">
          <div
            class="skeleton"
            style="height: 80px; margin-bottom: 12px; border-radius: 12px"
          ></div>
          <div
            class="skeleton"
            style="height: 120px; margin-bottom: 12px; border-radius: 12px"
          ></div>
          <div
            class="skeleton"
            style="height: 160px; border-radius: 12px"
          ></div>
        </div>

        <div v-else-if="store.sections.length === 0" class="canvas-empty">
          <div class="empty-icon">
            <Palette :size="48" stroke-width="1.5" />
          </div>
          <h3>Create Your Portfolio!</h3>
          <p>Click Block on the left to add components</p>
        </div>

        <TransitionGroup
          v-else
          name="section-list"
          tag="div"
          class="sections-list"
          ref="listRef"
        >
          <SectionBlock
            v-for="section in sortedSections"
            :key="section.id"
            :section="section"
            :is-selected="selectedSectionId === section.id"
            :theme-class="themeClass"
            :theme-vars="themeVars"
            :class="{ 'block-highlight': highlightedBlocks.has(section.id) }"
            @select="selectSection(section.id)"
            @delete="store.deleteSection(section.id)"
            @toggle-visibility="store.toggleSectionVisibility(section.id)"
            @move-up="moveSection(section.id, 'up')"
            @move-down="moveSection(section.id, 'down')"
          />
        </TransitionGroup>
      </main>

      <!-- Right: Properties Panel -->
      <aside class="props-panel">
        <div class="props-header">
          <span class="props-title"
            ><Palette :size="14" class="icon-inline" /> Properties</span
          >
        </div>

        <!-- Theme Settings -->
        <div class="props-section">
          <div class="props-label">Theme Mode</div>
          <div class="theme-toggle">
            <button
              :class="['toggle-btn', { active: theme.mode === 'dark' }]"
              @click="updateTheme({ mode: 'dark', bg_color: '#000000' })"
            >
              <Moon :size="14" /> Dark
            </button>
            <button
              :class="['toggle-btn', { active: theme.mode === 'light' }]"
              @click="updateTheme({ mode: 'light', bg_color: '#ffffff' })"
            >
              <Sun :size="14" /> Light
            </button>
          </div>
        </div>

        <div class="props-section">
          <div class="props-label">Primary Gradient</div>
          <div
            class="color-palette"
            style="
              display: flex;
              gap: 8px;
              align-items: center;
              margin-bottom: 12px;
            "
          >
            <div
              v-for="color in COLORS"
              :key="color.value"
              class="color-swatch"
              :style="{ background: color.value }"
              :class="{ active: theme.primary_color === color.value }"
              :title="color.name"
              @click="
                updateTheme({ primary_color: color.value, secondary_color: '' })
              "
            />
          </div>
          <div
            style="
              display: flex;
              align-items: center;
              gap: 8px;
              margin-top: 8px;
            "
          >
            <input
              type="color"
              class="form-input"
              style="height: 36px; padding: 2px"
              :value="theme.primary_color"
              @input="
                updateTheme({
                  primary_color: ($event.target as HTMLInputElement).value
                })
              "
            />
            <span style="font-size: 11px; color: var(--muted)">To</span>
            <input
              type="color"
              class="form-input"
              style="height: 36px; padding: 2px"
              :value="theme.secondary_color"
              @input="
                updateTheme({
                  secondary_color: ($event.target as HTMLInputElement).value
                })
              "
            />
          </div>
        </div>

        <div class="props-section">
          <div class="props-label">Background Custom Color</div>
          <input
            type="color"
            class="form-input"
            style="height: 40px; padding: 2px"
            v-model="theme.bg_color"
            @change="updateTheme({ bg_color: theme.bg_color })"
          />
        </div>

        <div class="props-section">
          <div class="props-label">Image Border Color</div>
          <input
            type="color"
            class="form-input"
            style="height: 40px; padding: 2px"
            v-model="theme.border_color"
            @change="updateTheme({ border_color: theme.border_color })"
          />
        </div>

        <div class="props-section">
          <div class="props-label">Font</div>
          <select
            class="form-input form-select"
            v-model="theme.font"
            @change="updateTheme({ font: theme.font })"
          >
            <option v-for="font in FONTS" :key="font" :value="font">
              {{ font }}
            </option>
          </select>
        </div>

        <div class="props-section">
          <div class="props-label">Layout</div>
          <select
            class="form-input form-select"
            v-model="theme.layout"
            @change="updateTheme({ layout: theme.layout })"
          >
            <option value="centered">Centered</option>
            <option value="left">Left Aligned</option>
            <option value="split">Split</option>
          </select>
        </div>

        <div class="props-section">
          <div class="props-label">Portfolio URL</div>
          <div class="url-display">
            <span class="url-prefix">pk.io/</span>
            <span class="url-slug">{{ store.activePortfolio?.slug }}</span>
          </div>
        </div>

        <!-- Section Editor (when section selected) -->
        <div v-if="selectedSection" class="section-editor">
          <div class="editor-divider"></div>
          <div class="props-label">
            <Pencil :size="12" class="icon-inline" /> Edit:
            {{ selectedSection.type }}
          </div>
          <SectionEditor
            :section="selectedSection"
            @update="handleSectionUpdate"
          />
        </div>

        <!-- AI Actions -->
        <div class="ai-panel">
          <div class="ai-panel-title">
            <Bot :size="14" class="icon-inline" /> AI Assistant
          </div>
          <button
            class="btn btn-secondary btn-sm btn-icon-text"
            style="width: 100%; margin-bottom: 8px"
            @click="aiScoreResume"
          >
            <BarChart2 :size="14" /> Score My Resume
          </button>
          <button
            class="btn btn-secondary btn-sm btn-icon-text"
            style="width: 100%"
            @click="aiImproveContent"
          >
            <Sparkles :size="14" /> Improve Content
          </button>
          <div v-if="aiResult" class="ai-result">
            <pre>{{ aiResult }}</pre>
          </div>
        </div>
      </aside>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, reactive, nextTick, watch } from 'vue'
import { useRoute, RouterLink } from 'vue-router'
import { usePortfolioStore } from '@/stores/portfolio'
import { aiAPI, portfolioAPI } from '@/api'
import { BLOCK_TYPES, type SectionType } from '@/types'
import {
  Globe,
  Save,
  Eye,
  EyeOff,
  Zap,
  Blocks,
  Palette,
  Moon,
  Sun,
  Pencil,
  BarChart2,
  Sparkles,
  ArrowLeft,
  Play,
  Download
} from 'lucide-vue-next'
import SectionBlock from '@/components/builder/SectionBlock.vue'
import SectionEditor from '@/components/builder/SectionEditor.vue'
import { debounce } from 'lodash'
import Sortable from 'sortablejs'

const route = useRoute()
const store = usePortfolioStore()

const selectedSectionId = ref<string | null>(null)
const aiResult = ref<string | null>(null)
const canvasRef = ref<HTMLElement | null>(null)
const listRef = ref<any>(null)
const highlightedBlocks = ref<Set<string>>(new Set())
const isExporting = ref(false)

const COLORS = [
  { name: 'Indigo', value: '#4F46E5' },
  { name: 'Purple', value: '#a855f7' },
  { name: 'Pink', value: '#ec4899' },
  { name: 'Cyan', value: '#06b6d4' },
  { name: 'Emerald', value: '#10b981' },
  { name: 'Orange', value: '#f97316' }
]

const FONTS = [
  'Plus Jakarta Sans',
  'Syne',
  'Sarabun',
  'Prompt',
  'Noto Sans Thai'
]

const theme = reactive({
  mode: 'dark' as 'dark' | 'light',
  primary_color: '#4F46E5',
  secondary_color: '',
  bg_color: '',
  border_color: '',
  font: 'Plus Jakarta Sans',
  layout: 'centered' as 'centered' | 'left' | 'split'
})

const heroSection = computed(() =>
  store.sections.find((s) => s.type === 'hero' && s.is_visible !== false)
)

const contentSections = computed(() =>
  [...store.sections]
    .filter((s) => s.type !== 'hero' && s.is_visible !== false)
    .sort((a, b) => a.position - b.position)
)

const sortedSections = computed(() =>
  [...store.sections].sort((a, b) => a.position - b.position)
)

const selectedSection = computed(() =>
  store.sections.find((s) => s.id === selectedSectionId.value)
)

onMounted(async () => {
  const portfolioId = route.params.id as string
  await store.loadPortfolio(portfolioId)
  if (store.activePortfolio?.theme) {
    Object.assign(theme, store.activePortfolio.theme)
    if (theme.mode === ('system' as any)) {
      theme.mode = 'dark'
    }
  }

  // Initialize Drag and Drop if list exists
  if (listRef.value) {
    initSortable()
  } else {
    // If it's initially empty, wait for sections to load/add
    const unwatch = watch(
      () => store.sections.length,
      (newLen) => {
        if (newLen > 0) {
          nextTick(() => {
            if (listRef.value) initSortable()
          })
          unwatch()
        }
      }
    )
  }
})

function initSortable() {
  const el = listRef.value.$el || listRef.value
  if (!el) return

  Sortable.create(el, {
    handle: '.drag-handle', // Drag only by this icon
    animation: 250, // Smooth slide animation
    delay: 150, // Require 150ms hold to start dragging (Great for preventing accidental drags on touch)
    delayOnTouchOnly: true, // Only require hold on touch devices
    ghostClass: 'sortable-ghost', // Class added to the dragged item
    onEnd: (evt: Sortable.SortableEvent) => {
      if (evt.oldIndex === undefined || evt.newIndex === undefined) return
      if (evt.oldIndex === evt.newIndex) return

      const newOrder = [...sortedSections.value]
      const [movedItem] = newOrder.splice(evt.oldIndex, 1)
      newOrder.splice(evt.newIndex, 0, movedItem)

      store.reorderSections(newOrder)
    }
  })
}

const themeClass = computed(() => {
  return `theme-${theme.mode || 'dark'}`
})

const themeVars = computed(() => {
  const primary = theme.primary_color || '#4F46E5'
  const secondary = theme.secondary_color || ''
  const font = theme.font || 'Prompt'
  // In light mode, glow needs to be slightly more transparent so it isn't overpowering, but visible.
  const glowHex = theme.mode === 'light' ? '15' : '40'
  const vars: Record<string, string> = {
    '--primary': primary,
    '--primary-glow': `${primary}${glowHex}`,
    '--secondary': secondary ? secondary : `${primary}${glowHex}`,
    '--font-display': font,
    '--font-body': font
  }
  if (theme.bg_color && theme.bg_color !== '#000000') {
    vars['--bg'] = theme.bg_color
  }
  if (theme.border_color && theme.border_color !== '#000000') {
    vars['--avatar-border'] = theme.border_color
  }
  return vars
})

function selectSection(id: string) {
  selectedSectionId.value = selectedSectionId.value === id ? null : id
}

async function addBlock(type: SectionType) {
  const newSectionId = await store.addSection(type)
  if (!newSectionId) return

  // Auto-scroll to bottom
  await nextTick()
  if (canvasRef.value) {
    canvasRef.value.scrollTo({
      top: canvasRef.value.scrollHeight,
      behavior: 'smooth'
    })
  }

  // Highlight effect
  highlightedBlocks.value.add(newSectionId)
  setTimeout(() => {
    highlightedBlocks.value.delete(newSectionId)
  }, 1000) // Remove class after animation ends
}

function updateTheme(partial: Partial<typeof theme>) {
  Object.assign(theme, partial)
  store.savePortfolio({ theme: { ...theme } })
}

async function togglePublish() {
  if (store.activePortfolio?.is_published) {
    await store.unpublishPortfolio()
  } else {
    await store.publishPortfolio()
  }
}

function previewPortfolio() {
  const slug = store.activePortfolio?.slug
  if (slug) window.open(`/p/${slug}`, '_blank')
}

async function exportPDF() {
  if (!store.activePortfolio?.id) return

  try {
    isExporting.value = true
    const response = await portfolioAPI.exportPdf(store.activePortfolio.id)

    // Create a blob URL and trigger download
    const blob = new Blob([response.data], { type: 'application/pdf' })
    const url = window.URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    link.download = `${store.activePortfolio.slug || 'portfolio'}.pdf`
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
    window.URL.revokeObjectURL(url)
  } catch (error) {
    console.error('Failed to export PDF:', error)
  } finally {
    isExporting.value = false
  }
}

// Debounce the backend API call to save bandwidth
const debouncedSave = debounce(async (id: string, data: any) => {
  await store.updateSection(id, data)
}, 500)

// Update local state instantly so the Preview refreshes immediately in real-time
function handleSectionUpdate(data: any) {
  if (!selectedSectionId.value) return

  // Optimistic update for real-time canvas preview
  const sec = store.sections.find((s) => s.id === selectedSectionId.value)
  if (sec) sec.data = JSON.parse(JSON.stringify(data))

  debouncedSave(selectedSectionId.value, data)
}

function moveSection(id: string, direction: 'up' | 'down') {
  const idx = sortedSections.value.findIndex((s) => s.id === id)
  if (direction === 'up' && idx === 0) return
  if (direction === 'down' && idx === sortedSections.value.length - 1) return

  const newOrder = [...sortedSections.value]
  const swapIdx = direction === 'up' ? idx - 1 : idx + 1
  ;[newOrder[idx], newOrder[swapIdx]] = [newOrder[swapIdx], newOrder[idx]]
  store.reorderSections(newOrder)
}

async function aiScoreResume() {
  const content = JSON.stringify(store.sections.map((s) => s.data))
  try {
    const { data } = await aiAPI.scoreResume(content, 'Software Developer')
    aiResult.value = data.data.analysis
  } catch {
    aiResult.value = 'AI unavailable. Please check API key.'
  }
}

async function aiImproveContent() {
  const heroSection = store.sections.find((s) => s.type === 'hero')
  if (!heroSection || !selectedSectionId.value) {
    aiResult.value = 'Select a section first to improve its content.'
    return
  }
  const text = heroSection.data?.tagline || ''
  try {
    const { data } = await aiAPI.improveText(text, 'portfolio tagline')
    aiResult.value = data.data.improved_text
  } catch {
    aiResult.value = 'AI unavailable.'
  }
}
</script>

<style scoped>
.builder-layout {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background: var(--bg);
  overflow: hidden;
}

/* Topbar */
.builder-topbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 24px;
  background: var(--sidebar-bg);
  border-bottom: 1px solid var(--border);
  flex-shrink: 0;
  z-index: 50;
}
.back-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  color: var(--muted);
  text-decoration: none;
  border-radius: 8px;
  background: var(--surface);
  border: 1px solid var(--border);
  transition: all 0.2s ease;
}

.back-btn:hover {
  color: var(--text);
  border-color: var(--indigo);
  background: rgba(79, 70, 229, 0.1);
  transform: translateY(-1px);
}

.back-btn:hover .icon {
  transform: translateX(-2px);
  transition: transform 0.2s ease;
}
.topbar-center {
  display: flex;
  align-items: center;
  gap: 12px;
}
.portfolio-slug {
  font-size: 13px;
  color: var(--indigo);
  font-weight: 600;
  display: inline-flex;
  align-items: center;
  gap: 6px;
}
.saving-indicator {
  font-size: 12px;
  color: var(--muted);
  animation: pulse 1s infinite;
  display: inline-flex;
  align-items: center;
  gap: 6px;
}
.topbar-actions {
  display: flex;
  gap: 10px;
}
.publish-btn {
  min-width: 115px;
}

/* Main 3-column layout */
.builder-main {
  display: grid;
  grid-template-columns: 240px 1fr 260px;
  flex: 1;
  overflow: hidden;
}

/* Block Palette */
.block-palette {
  background: var(--sidebar-bg);
  border-right: 1px solid var(--border);
  overflow-y: auto;
  padding: 16px 12px;
}

.palette-header {
  display: flex;
  justify-content: space-between;
  align-items: center;

  margin-top: -10px; /* 👈 ขยับขึ้น */
  margin-bottom: 5px;
  padding: 12px 16px;

  border-radius: 2px;
  border-bottom: 1px solid var(--divider-color);
}

.palette-title {
  font-size: 12px;
  font-weight: 700;
  color: var(--muted);
  letter-spacing: 1.5px;
  text-transform: uppercase;
}
.palette-hint {
  font-size: 10px;
  color: rgba(148, 163, 184, 0.5);
}

.palette-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 12px;
  border-radius: 10px;
  cursor: pointer;
  transition: all 0.15s;
  margin-bottom: 3px;
  border: 1px solid transparent;
}
.palette-item:hover {
  background: rgba(79, 70, 229, 0.12);
  border-color: rgba(79, 70, 229, 0.25);
}
.palette-icon {
  font-size: 18px;
  width: 28px;
  text-align: center;
  flex-shrink: 0;
}
.palette-name {
  font-size: 13px;
  font-weight: 600;
}
.palette-desc {
  font-size: 11px;
  color: var(--muted);
}
.palette-add {
  margin-left: auto;
  color: var(--indigo);
  font-size: 18px;
  font-weight: 300;
  opacity: 0;
  transition: opacity 0.15s;
}
.palette-item:hover .palette-add {
  opacity: 1;
}

/* Canvas */
.builder-canvas {
  overflow-y: auto;
  padding: 24px;
  background: var(--bg);
}
.canvas-empty {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 60%;
  text-align: center;
  border: 2px dashed var(--border);
  border-radius: 16px;
  padding: 48px;
}
.empty-icon {
  font-size: 56px;
  margin-bottom: 16px;
}
.canvas-empty h3 {
  font-size: 20px;
  margin-bottom: 8px;
}
.canvas-empty p {
  color: var(--muted);
  font-size: 14px;
}

.sections-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
.section-list-move {
  transition: transform 0.3s;
}
.section-list-enter-active {
  animation: fadeIn 0.3s ease;
}
.section-list-leave-active {
  animation: fadeIn 0.2s ease reverse;
}

/* Props Panel */
.props-panel {
  background: var(--sidebar-bg);
  border-left: 1px solid var(--border);
  overflow-y: auto;
  padding: 16px;
}
.props-header {
  margin-bottom: 16px;
}
.props-title {
  font-size: 12px;
  font-weight: 700;
  color: var(--muted);
  letter-spacing: 1.5px;
  text-transform: uppercase;
}
.props-section {
  margin-bottom: 20px;
}
.props-label {
  font-size: 12px;
  font-weight: 600;
  color: var(--muted);
  margin-bottom: 8px;
}

.theme-toggle {
  display: flex;
  gap: 6px;
}
.toggle-btn {
  flex: 1;
  padding: 7px;
  border-radius: 8px;
  border: 1px solid var(--border);
  background: transparent;
  color: var(--muted);
  font-size: 12px;
  cursor: pointer;
  transition: all 0.15s;
}
.toggle-btn.active {
  background: rgba(79, 70, 229, 0.15);
  border-color: var(--indigo);
  color: var(--indigo);
}

.color-palette {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}
.color-swatch {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  cursor: pointer;
  transition: transform 0.15s;
  border: 2px solid transparent;
}
.color-swatch:hover {
  transform: scale(1.2);
}
.color-swatch.active {
  border-color: #fff;
  transform: scale(1.1);
}

.url-display {
  background: var(--bg);
  border: 1px solid var(--border);
  border-radius: 8px;
  padding: 8px 12px;
  font-size: 12px;
}
.url-prefix {
  color: var(--muted);
}
.url-slug {
  color: var(--indigo);
  font-weight: 600;
}

.section-editor {
  margin-bottom: 20px;
}

.editor-divider {
  height: 1px;
  margin: 24px 0;
  border-radius: 999px;
}

/* Dark */
:global(.theme-dark) .editor-divider {
  background: rgba(255, 255, 255, 0.08);
}

/* Light */
:global(.theme-light) .editor-divider {
  background: rgba(0, 0, 0, 0.08);
}

.ai-panel {
  background: var(--primary);
  border: 1px solid var(--border);
  border-radius: 12px;
  padding: 14px;
}
.ai-panel-title {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  font-weight: 700;
  margin-bottom: 12px;
  color: var(--neon-purple);
}
.ai-result {
  margin-top: 12px;
  padding: 10px;
  background: var(--bg);
  border-radius: 8px;
  font-size: 11px;
  color: var(--muted);
  max-height: 120px;
  overflow-y: auto;
  border: 1px solid var(--border);
}
.ai-result pre {
  white-space: pre-wrap;
  word-break: break-word;
}

.icon-inline {
  vertical-align: middle;
  border: 1px solid var(--border);
}

.btn-icon-text {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  justify-content: center;
}

.palette-title {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  font-weight: 700;
  color: var(--muted);
  letter-spacing: 1.5px;
  text-transform: uppercase;
}
.props-title {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  font-weight: 700;
  color: var(--muted);
  letter-spacing: 1.5px;
  text-transform: uppercase;
}
.props-label {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  font-weight: 600;
  color: var(--muted);
  margin-bottom: 8px;
}

.toggle-btn {
  flex: 1;
  padding: 7px;
  border-radius: 8px;
  border: 1px solid var(--border);
  background: transparent;
  color: var(--muted);
  font-size: 12px;
  cursor: pointer;
  transition: all 0.15s;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
}

@media (max-width: 1024px) {
  .builder-main {
    grid-template-columns: 200px 1fr;
  }
  .props-panel {
    display: none;
  }
}
@media (max-width: 768px) {
  .builder-main {
    grid-template-columns: 1fr;
  }
  .block-palette {
    display: none;
  }
}

.builder-layout.theme-dark {
  --divider-color: rgba(255, 255, 255, 0.15);
}

.builder-layout.theme-light {
  --divider-color: rgba(0, 0, 0, 0.15);
}
</style>
