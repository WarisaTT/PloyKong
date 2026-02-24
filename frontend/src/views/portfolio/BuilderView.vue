<template>
  <div class="builder-layout">
    <!-- ─── Top Bar ─────────────────────────────────────────────────────── -->
    <header class="builder-topbar">
      <RouterLink to="/dashboard" class="back-btn">
        ← Dashboard
      </RouterLink>

      <div class="topbar-center">
        <span class="portfolio-slug">🌐 {{ store.activePortfolio?.slug }}.ploykong.com</span>
        <span v-if="store.saving" class="saving-indicator">💾 Saving...</span>
      </div>

      <div class="topbar-actions">
        <button class="btn btn-secondary btn-sm" @click="previewPortfolio">
          👁️ Preview
        </button>
        <button
          class="btn btn-sm"
          :class="store.activePortfolio?.is_published ? 'btn-danger' : 'btn-primary'"
          @click="togglePublish"
        >
          {{ store.activePortfolio?.is_published ? '⏸ Unpublish' : '⚡ Publish' }}
        </button>
      </div>
    </header>

    <!-- ─── Main Builder Area ──────────────────────────────────────────── -->
    <div class="builder-main">
      <!-- Left: Block Palette -->
      <aside class="block-palette">
        <div class="palette-header">
          <span class="palette-title">🧩 Blocks</span>
          <span class="palette-hint">คลิกเพื่อเพิ่ม</span>
        </div>
        <div class="palette-list">
          <div
            v-for="block in BLOCK_TYPES"
            :key="block.type"
            class="palette-item"
            @click="addBlock(block.type)"
          >
            <span class="palette-icon">{{ block.icon }}</span>
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
          <div class="skeleton" style="height:80px;margin-bottom:12px;border-radius:12px"></div>
          <div class="skeleton" style="height:120px;margin-bottom:12px;border-radius:12px"></div>
          <div class="skeleton" style="height:160px;border-radius:12px"></div>
        </div>

        <div v-else-if="store.sections.length === 0" class="canvas-empty">
          <div class="empty-icon">🎨</div>
          <h3>เริ่มสร้างพอร์ตของคุณ!</h3>
          <p>คลิก Block ด้านซ้ายเพื่อเพิ่มส่วนประกอบ</p>
        </div>

        <TransitionGroup v-else name="section-list" tag="div" class="sections-list">
          <SectionBlock
            v-for="section in sortedSections"
            :key="section.id"
            :section="section"
            :is-selected="selectedSectionId === section.id"
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
          <span class="props-title">🎨 Properties</span>
        </div>

        <!-- Theme Settings -->
        <div class="props-section">
          <div class="props-label">Theme Mode</div>
          <div class="theme-toggle">
            <button
              :class="['toggle-btn', { active: theme.mode === 'dark' }]"
              @click="updateTheme({ mode: 'dark' })"
            >🌙 Dark</button>
            <button
              :class="['toggle-btn', { active: theme.mode === 'light' }]"
              @click="updateTheme({ mode: 'light' })"
            >☀️ Light</button>
          </div>
        </div>

        <div class="props-section">
          <div class="props-label">Primary Color</div>
          <div class="color-palette">
            <div
              v-for="color in COLORS"
              :key="color.value"
              class="color-swatch"
              :style="{ background: color.value }"
              :class="{ active: theme.primary_color === color.value }"
              :title="color.name"
              @click="updateTheme({ primary_color: color.value })"
            />
          </div>
        </div>

        <div class="props-section">
          <div class="props-label">Font</div>
          <select class="form-input" v-model="theme.font" @change="updateTheme({ font: theme.font })">
            <option v-for="font in FONTS" :key="font" :value="font">{{ font }}</option>
          </select>
        </div>

        <div class="props-section">
          <div class="props-label">Layout</div>
          <select class="form-input" v-model="theme.layout" @change="updateTheme({ layout: theme.layout })">
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
          <div class="props-label">✏️ Edit: {{ selectedSection.type }}</div>
          <SectionEditor
            :section="selectedSection"
            @update="handleSectionUpdate"
          />
        </div>

        <!-- AI Actions -->
        <div class="ai-panel">
          <div class="ai-panel-title">🤖 AI Assistant</div>
          <button class="btn btn-secondary btn-sm" style="width:100%;margin-bottom:8px" @click="aiScoreResume">
            📊 Score My Resume
          </button>
          <button class="btn btn-secondary btn-sm" style="width:100%" @click="aiImproveContent">
            ✨ Improve Content
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
import { ref, computed, onMounted, reactive } from 'vue'
import { useRoute, RouterLink } from 'vue-router'
import { usePortfolioStore } from '@/stores/portfolio'
import { aiAPI } from '@/api'
import { BLOCK_TYPES, type SectionType } from '@/types'
import SectionBlock from '@/components/builder/SectionBlock.vue'
import SectionEditor from '@/components/builder/SectionEditor.vue'

const route = useRoute()
const store = usePortfolioStore()

const selectedSectionId = ref<string | null>(null)
const aiResult = ref<string | null>(null)
const canvasRef = ref<HTMLElement | null>(null)

const COLORS = [
  { name: 'Indigo', value: '#4F46E5' },
  { name: 'Purple', value: '#a855f7' },
  { name: 'Pink', value: '#ec4899' },
  { name: 'Cyan', value: '#06b6d4' },
  { name: 'Emerald', value: '#10b981' },
  { name: 'Orange', value: '#f97316' },
]

const FONTS = ['Syne', 'Plus Jakarta Sans', 'Sarabun', 'Prompt', 'Noto Sans Thai']

const theme = reactive({
  mode: 'dark' as 'dark' | 'light',
  primary_color: '#4F46E5',
  font: 'Syne',
  layout: 'centered' as 'centered' | 'left' | 'split',
})

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
  }
})

function selectSection(id: string) {
  selectedSectionId.value = selectedSectionId.value === id ? null : id
}

async function addBlock(type: SectionType) {
  await store.addSection(type)
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

async function handleSectionUpdate(data: any) {
  if (!selectedSectionId.value) return
  await store.updateSection(selectedSectionId.value, data)
}

function moveSection(id: string, direction: 'up' | 'down') {
  const idx = sortedSections.value.findIndex((s) => s.id === id)
  if (direction === 'up' && idx === 0) return
  if (direction === 'down' && idx === sortedSections.value.length - 1) return

  const newOrder = [...sortedSections.value]
  const swapIdx = direction === 'up' ? idx - 1 : idx + 1;
  [newOrder[idx], newOrder[swapIdx]] = [newOrder[swapIdx], newOrder[idx]]
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
  background: rgba(8, 13, 31, 0.95);
  border-bottom: 1px solid var(--border);
  flex-shrink: 0;
  z-index: 50;
}
.back-btn { color: var(--muted); text-decoration: none; font-size: 14px; transition: color 0.2s; }
.back-btn:hover { color: #fff; }
.topbar-center { display: flex; align-items: center; gap: 12px; }
.portfolio-slug { font-size: 13px; color: var(--neon-cyan); font-weight: 600; }
.saving-indicator { font-size: 12px; color: var(--muted); animation: pulse 1s infinite; }
.topbar-actions { display: flex; gap: 10px; }

/* Main 3-column layout */
.builder-main {
  display: grid;
  grid-template-columns: 240px 1fr 260px;
  flex: 1;
  overflow: hidden;
}

/* Block Palette */
.block-palette {
  background: rgba(8, 13, 31, 0.9);
  border-right: 1px solid var(--border);
  overflow-y: auto;
  padding: 16px 12px;
}
.palette-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 12px; padding: 0 4px; }
.palette-title { font-size: 12px; font-weight: 700; color: var(--muted); letter-spacing: 1.5px; text-transform: uppercase; }
.palette-hint { font-size: 10px; color: rgba(148,163,184,0.5); }

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
.palette-icon { font-size: 18px; width: 28px; text-align: center; flex-shrink: 0; }
.palette-name { font-size: 13px; font-weight: 600; }
.palette-desc { font-size: 11px; color: var(--muted); }
.palette-add { margin-left: auto; color: var(--indigo); font-size: 18px; font-weight: 300; opacity: 0; transition: opacity 0.15s; }
.palette-item:hover .palette-add { opacity: 1; }

/* Canvas */
.builder-canvas {
  overflow-y: auto;
  padding: 24px;
  background: var(--bg);
}
.canvas-empty {
  display: flex; flex-direction: column; align-items: center; justify-content: center;
  height: 60%; text-align: center;
  border: 2px dashed var(--border); border-radius: 16px; padding: 48px;
}
.empty-icon { font-size: 56px; margin-bottom: 16px; }
.canvas-empty h3 { font-size: 20px; margin-bottom: 8px; }
.canvas-empty p { color: var(--muted); font-size: 14px; }

.sections-list { display: flex; flex-direction: column; gap: 10px; }
.section-list-move { transition: transform 0.3s; }
.section-list-enter-active { animation: fadeIn 0.3s ease; }
.section-list-leave-active { animation: fadeIn 0.2s ease reverse; }

/* Props Panel */
.props-panel {
  background: rgba(8, 13, 31, 0.9);
  border-left: 1px solid var(--border);
  overflow-y: auto;
  padding: 16px;
}
.props-header { margin-bottom: 16px; }
.props-title { font-size: 12px; font-weight: 700; color: var(--muted); letter-spacing: 1.5px; text-transform: uppercase; }
.props-section { margin-bottom: 20px; }
.props-label { font-size: 12px; font-weight: 600; color: var(--muted); margin-bottom: 8px; }

.theme-toggle { display: flex; gap: 6px; }
.toggle-btn {
  flex: 1; padding: 7px; border-radius: 8px; border: 1px solid var(--border);
  background: transparent; color: var(--muted); font-size: 12px; cursor: pointer; transition: all 0.15s;
}
.toggle-btn.active {
  background: rgba(79, 70, 229, 0.2); border-color: var(--indigo); color: #fff;
}

.color-palette { display: flex; gap: 8px; flex-wrap: wrap; }
.color-swatch {
  width: 24px; height: 24px; border-radius: 50%;
  cursor: pointer; transition: transform 0.15s;
  border: 2px solid transparent;
}
.color-swatch:hover { transform: scale(1.2); }
.color-swatch.active { border-color: #fff; transform: scale(1.1); }

.url-display {
  background: rgba(255,255,255,0.04); border: 1px solid var(--border);
  border-radius: 8px; padding: 8px 12px; font-size: 12px;
}
.url-prefix { color: var(--muted); }
.url-slug { color: var(--neon-cyan); font-weight: 600; }

.section-editor { margin-bottom: 20px; }
.editor-divider { height: 1px; background: var(--border); margin: 20px 0; }

.ai-panel { background: rgba(79,70,229,0.08); border: 1px solid rgba(79,70,229,0.2); border-radius: 12px; padding: 14px; }
.ai-panel-title { font-size: 13px; font-weight: 700; margin-bottom: 12px; color: var(--neon-purple); }
.ai-result { margin-top: 12px; padding: 10px; background: rgba(0,0,0,0.3); border-radius: 8px; font-size: 11px; color: var(--muted); max-height: 120px; overflow-y: auto; }
.ai-result pre { white-space: pre-wrap; word-break: break-word; }

@media (max-width: 1024px) {
  .builder-main { grid-template-columns: 200px 1fr; }
  .props-panel { display: none; }
}
@media (max-width: 768px) {
  .builder-main { grid-template-columns: 1fr; }
  .block-palette { display: none; }
}
</style>
