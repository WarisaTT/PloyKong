<template>
  <div class="builder-layout">
    <!-- ─── Top Bar ─────────────────────────────────────────────────────── -->
    <header class="builder-topbar">
      <div class="topbar-left">
        <RouterLink to="/dashboard" class="back-btn" title="Dashboard">
          <ArrowLeft class="icon" :size="20" />
        </RouterLink>
      </div>

      <div class="topbar-center mobile-hide">
        <div class="slug-container">
          <span class="portfolio-slug">
            <Globe :size="14" class="icon-inline" />
            {{ store.activePortfolio?.slug }}.ploykong.com
          </span>
          <button
            class="copy-link-btn-indigo"
            :disabled="!store.activePortfolio?.is_published"
            title="Copy Public Link"
            @click="copyPublicLink"
          >
            <Check v-if="copied" :size="14" />
            <Copy v-else :size="14" />
            <span class="copy-text">{{ copied ? 'Copied!' : 'Copy' }}</span>
          </button>
        </div>
        <span v-if="store.saving" class="saving-indicator">
          <Save :size="12" class="icon-inline" /> Saving...
        </span>
      </div>

      <!-- MOBILE TOP NAV SELECTOR -->
      <div class="builder-mobile-tabs @768:flex">
        <button 
          class="mobile-tab-btn" 
          :class="{ active: !isLeftCollapsed }"
          @click="openSidebar('left')"
          title="Blocks"
        >
          <Blocks :size="18" />
        </button>
        <button 
          class="mobile-tab-btn" 
          :class="{ active: !isRightCollapsed && !selectedSectionId }"
          @click="openSidebar('right')"
          title="Design"
        >
          <Palette :size="18" />
        </button>
        <button 
          v-if="selectedSectionId"
          class="mobile-tab-btn pulse-hint" 
          :class="{ active: !isRightCollapsed && selectedSectionId }"
          @click="scrollToEditor"
          title="Edit Section"
        >
          <Pencil :size="18" />
        </button>
        <button 
          class="mobile-tab-btn"
          :disabled="!store.activePortfolio?.is_published"
          @click="copyPublicLink"
          title="Copy Link"
        >
          <Check v-if="copied" :size="18" class="text-success" />
          <Copy v-else :size="18" />
        </button>
        <button 
          v-if="!isLeftCollapsed || !isRightCollapsed"
          class="mobile-tab-btn hide-btn"
          @click="isLeftCollapsed = true; isRightCollapsed = true"
          title="Close Sidebar"
        >
          <X :size="18" />
        </button>
      </div>

      <div class="topbar-actions">
        <button
          class="btn btn-secondary btn-sm btn-icon-only-mobile"
          :disabled="!store.activePortfolio?.is_published"
          @click="previewPortfolio"
        >
          <Eye v-if="store.activePortfolio?.is_published" :size="14" />
          <EyeOff v-else :size="14" />
          <span class="mobile-hide">Preview</span>
        </button>
        <button
          class="btn btn-sm btn-icon-only-mobile publish-btn"
          :class="store.activePortfolio?.is_published ? 'btn-danger' : 'btn-primary'"
          @click="togglePublish"
        >
          <Play v-if="store.activePortfolio?.is_published" :size="14" />
          <Zap v-else :size="14" />
          <span class="mobile-hide">{{ store.activePortfolio?.is_published ? 'Unpublish' : 'Publish' }}</span>
        </button>
      </div>
    </header>

    <!-- ─── Main Builder Area ──────────────────────────────────────────── -->
    <div
      class="builder-main"
      :class="{
        'left-collapsed': isLeftCollapsed,
        'right-collapsed': isRightCollapsed,
      }"
    >
      <!-- Left: Block Palette -->
      <aside 
        class="block-palette" 
        :class="{ collapsed: isLeftCollapsed }" 
        :style="{ 
          width: leftSidebarWidth + 'px', 
          '--sidebar-width': leftSidebarWidth + 'px' 
        }"
      >
        <div class="resize-handle right-handle" @mousedown="startResizing('left', $event)"></div>
        <div class="palette-header">
          <span class="palette-title" :style="{paddingTop: '10px'}">
            <Blocks :size="18" class="icon-inline" :style="{paddingTop: '5px'}" /> Blocks
          </span>
          <button class="mobile-close-sidebar @768:flex" @click="isLeftCollapsed = true">
            <X :size="18" />
          </button>
          <span class="palette-hint mobile-hide" :style="{paddingTop: '10px'}">Click to add</span>
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
      <div class="canvas-wrapper">
        <!-- Floating Sidebar Toggles - Hidden on mobile -->
        <button
          class="sidebar-toggle left-toggle mobile-hide"
          @click="isLeftCollapsed = !isLeftCollapsed"
          :title="isLeftCollapsed ? 'Expand Blocks' : 'Collapse Blocks'"
        >
          <ChevronRight v-if="isLeftCollapsed" :size="16" />
          <ChevronLeft v-else :size="16" />
        </button>

        <button
          class="sidebar-toggle right-toggle mobile-hide"
          @click="isRightCollapsed = !isRightCollapsed"
          :title="isRightCollapsed ? 'Expand Properties' : 'Collapse Properties'"
        >
          <ChevronLeft v-if="isRightCollapsed" :size="16" />
          <ChevronRight v-else :size="16" />
        </button>

        <main class="builder-canvas" ref="canvasRef">
          <div class="canvas-content-limiter">
          <!-- Canvas Loading and Empty states -->
          <div v-if="store.loading" class="canvas-loading">...</div>
          <div v-else-if="store.sections.length === 0" class="canvas-empty">...</div>

          <TransitionGroup
            tag="div"
            name="section-list"
            v-else
            class="sections-list sections-grid"
            :class="[templateClass, { 'has-divider': theme.show_divider }]"
            ref="listRef"
          >
            <div
              v-for="section in gridSections"
              :key="section.id"
              :data-id="section.id"
              class="section-wrapper drag-item"
              :class="{ 
                'span-half': section.column_span === 'half', 
                'span-full': section.column_span !== 'half', 
                'is-right': section.isRight,
                'is-row-start': section.isRowStart,
                'is-paired': section.isPaired,
                'builder-half-row': section.column_span === 'half',
                'builder-full-row': section.column_span !== 'half'
              }"
            >
              <SectionBlock
                :section="section"
                :is-selected="selectedSectionId === section.id"
                :theme-class="themeClass"
                :template-class="templateClass"
                :is-half-split="section.column_span === 'half'"
                :theme-vars="themeVars"
                :class="{ 'block-highlight': highlightedBlocks.has(section.id) }"
                @select="selectSection(section.id)"
                @delete="store.deleteSection(section.id)"
                @toggle-visibility="store.toggleSectionVisibility(section.id)"
                @toggle-column-span="store.toggleSectionColumnSpan(section.id)"
                @move-up="moveSection(section.id, 'up')"
                @move-down="moveSection(section.id, 'down')"
                @duplicate="store.duplicateSection(section.id)"
              />
            </div>
          </TransitionGroup>
          </div>
        </main>
      </div>

      <!-- Right: Properties Panel -->
      <aside 
        class="props-panel" 
        :class="{ collapsed: isRightCollapsed }" 
        :style="{ 
          width: rightSidebarWidth + 'px', 
          '--sidebar-width': rightSidebarWidth + 'px' 
        }"
      >
        <div class="resize-handle left-handle" @mousedown="startResizing('right', $event)"></div>
        <div class="props-inner-scroll" ref="propsScrollRef">
          <div class="props-header">
          <span class="props-title">
            <Palette :size="14" class="icon-inline" /> Properties
          </span>
          <button class="mobile-close-sidebar @768:flex" @click="isRightCollapsed = true">
            <X :size="18" />
          </button>
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
              v-for="(color, idx) in currentColors"
              :key="idx"
              class="color-swatch"
              :style="{ background: color }"
              :class="{ active: theme.primary_color === color }"
              :title="`Color ${idx + 1}`"
              @click="
                updateTheme({ primary_color: color, secondary_color: '' })
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
          <div class="props-label">Portfolio URL</div>
          <div class="url-display">
            <span class="url-prefix">pk.io/</span>
            <span class="url-slug">{{ store.activePortfolio?.slug }}</span>
          </div>
        </div>

        <!-- Section Editor (when section selected) -->
        <div v-if="selectedSection" class="section-editor" ref="editorRef">
          <div class="editor-divider"></div>
          <div class="props-title" style="margin-bottom: 12px;">
            <Pencil :size="12" class="icon-inline" /> Edit:
            {{ selectedSection.type }}
          </div>

          <!-- Column Divider Toggle: Only show for half-width sections -->
          <div v-if="selectedSection.column_span === 'half'" class="props-section" style="background: rgba(var(--primary-glow), 0.05); padding: 12px; border-radius: 12px; border: 1px dashed var(--border); margin-bottom: 20px;">
            <label class="toggle-row" style="cursor: pointer;">
              <span class="props-label" style="margin: 0; font-size: 11px;">
                <LayoutGrid :size="12" class="icon-inline" /> Show Vertical Divider
              </span>
              <div class="toggle-switch">
                <input
                  type="checkbox"
                  v-model="theme.show_divider"
                  @change="updateTheme({ show_divider: theme.show_divider })"
                />
                <span class="slider round"></span>
              </div>
            </label>
            <p style="font-size: 10px; color: var(--muted); margin-top: 6px; line-height: 1.3;">
              แสดงเส้นคั่นแนวตั้งระหว่าง Section ที่วางคู่กัน
            </p>
          </div>
          <SectionEditor
            :section="selectedSection"
            :template="store.activePortfolio?.theme?.template || 'classic'"
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
  Download,
  ChevronLeft,
  ChevronRight,
  ChevronUp,
  X,
  Copy,
  Check
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
const editorRef = ref<HTMLElement | null>(null)
const propsScrollRef = ref<HTMLElement | null>(null)
const highlightedBlocks = ref<Set<string>>(new Set())

const isLeftCollapsed = ref(false)
const isRightCollapsed = ref(false)
const leftSidebarWidth = ref(260)
const rightSidebarWidth = ref(320)
const copied = ref(false)

// Resizing logic
const isResizing = ref<'left' | 'right' | null>(null)

function startResizing(side: 'left' | 'right', e: MouseEvent) {
  isResizing.value = side
  document.addEventListener('mousemove', handleMouseMove)
  document.addEventListener('mouseup', stopResizing)
  document.body.style.cursor = 'col-resize'
  document.body.style.userSelect = 'none'
}

function handleMouseMove(e: MouseEvent) {
  if (!isResizing.value) return
  if (isResizing.value === 'left') {
    const newWidth = e.clientX
    if (newWidth > 180 && newWidth < 500) {
      leftSidebarWidth.value = newWidth
    }
  } else {
    const newWidth = window.innerWidth - e.clientX
    if (newWidth > 250 && newWidth < 600) {
      rightSidebarWidth.value = newWidth
    }
  }
}

function stopResizing() {
  isResizing.value = null
  document.removeEventListener('mousemove', handleMouseMove)
  document.removeEventListener('mouseup', stopResizing)
  document.body.style.cursor = ''
  document.body.style.userSelect = ''
}

function hexToRgb(hex: string) {
  if (!hex) return "0, 0, 0";
  // Remove # if present
  const cleanHex = hex.startsWith('#') ? hex.slice(1) : hex;
  if (cleanHex.length !== 6) return "0, 0, 0";
  const r = parseInt(cleanHex.slice(0, 2), 16);
  const g = parseInt(cleanHex.slice(2, 4), 16);
  const b = parseInt(cleanHex.slice(4, 6), 16);
  return `${r}, ${g}, ${b}`;
}

const PALETTES = [
  { id:"classic", label:"Classic Colors", colors:["#4F46E5","#9333EA","#DB2777","#06B6D4","#10B981","#F97316"] },
  { id:"neon", label:"Neon Cyber", colors:["#00F5FF","#7B61FF","#FF00C8","#39FF14","#FFEA00","#FF3366"] },
  { id:"sunset", label:"Sunset Warm", colors:["#FF5E3A","#FF9500","#FF2A55","#D35DD8","#FFB800","#FF7043"] },
  { id:"forest", label:"Forest Nature", colors:["#10B981","#059669","#34D399","#A3E635","#06B6D4","#3B82F6"] },
  { id:"gold", label:"Luxury Gold", colors:["#D4AF37","#FBBF24","#FCD34D","#B45309","#92400E","#F59E0B"] }
];

const FONTS = [
  'Inter',
  'Outfit',
  'Space Grotesk',
  'Plus Jakarta Sans',
  'Syne',
  'Playfair Display',
  'Prompt',
  'Kanit',
  'Noto Sans Thai',
  'Sarabun',
  'Mitr',
  'Chakra Petch',
  'IBM Plex Sans Thai'
]

const theme = reactive({
  mode: 'dark' as 'dark' | 'light',
  primary_color: '#4F46E5',
  secondary_color: '',
  palette: 'classic',
  template: 'classic',
  bg_color: '',
  border_color: '',
  font: 'Plus Jakarta Sans',
  layout: 'centered' as 'centered' | 'left' | 'split',
  show_divider: false
})

const currentColors = computed(() => {
  const p = PALETTES.find(p => p.id === theme.palette)
  return p ? p.colors : PALETTES[0].colors
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

  // Mobile optimization: collapse sidebars on small screens
  if (window.innerWidth < 1100) {
    isLeftCollapsed.value = true
    isRightCollapsed.value = true
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
    // handle: '.drag-handle', // NOW REMOVED TO ALLOW DRAGGING THE WHOLE SECTION
    filter: 'input, button, .icon-btn, .color-swatch, .form-input, select', // Prevent dragging when clicking UI items
    preventOnFilter: false,
    animation: 250, // Smooth slide animation
    delay: 50, // Short delay to allow clicking
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

// Flattened sections for CSS Grid, calculating left/right placement for the divider
const gridSections = computed(() => {
  const result: (typeof sortedSections.value[0] & { isRight?: boolean, isRowStart?: boolean, isPaired?: boolean })[] = [];
  const items = sortedSections.value;
  
  for (let i = 0; i < items.length; i++) {
    const s = items[i];
    const prev = i > 0 ? items[i-1] : null;
    const next = i < items.length - 1 ? items[i+1] : null;
    
    if (s.column_span === 'half') {
      const isRight = prev?.column_span === 'half' && !result[i-1]?.isRight;
      const isRowStart = !isRight;
      const isPaired = isRowStart ? (next?.column_span === 'half') : true;
      
      result.push({ ...s, isRight, isRowStart, isPaired });
    } else {
      result.push({ ...s, isRight: false, isRowStart: true, isPaired: false });
    }
  }
  return result;
});

const themeClass = computed(() => {
  return `theme-${theme.mode || 'dark'}`
})

const templateClass = computed(() => {
  return `tpl-${theme.template || 'classic'}`
})

const themeVars = computed(() => {
  const primary = theme.primary_color || '#4F46E5'
  const secondary = theme.secondary_color || ''
  const font = theme.font || 'Prompt'
  const isLight = theme.mode === 'light'
  const glowHex = isLight ? '25' : '40'
  const vars: Record<string, string> = {
    '--primary': primary,
    '--primary-glow': `${primary}${glowHex}`,
    '--secondary': secondary ? secondary : `${primary}${glowHex}`,
    '--font-display': font,
    '--font-body': font,
    '--text': isLight ? '#0f172a' : '#ffffff',
    '--muted': isLight ? '#64748b' : '#94a3b8',
    '--surface': isLight ? '#ffffff' : '#111827',
    '--border': isLight ? '#e2e8f0' : '#1f2937'
  }
  if (theme.bg_color) {
    vars['--bg'] = theme.bg_color
    vars['--bg-rgb'] = hexToRgb(theme.bg_color)
  } else {
    vars['--bg'] = isLight ? '#faf7ff' : '#050814'
    vars['--bg-rgb'] = isLight ? '250, 247, 255' : '5, 8, 20'
  }
  if (theme.border_color) {
    vars['--avatar-border'] = theme.border_color
  }
  return vars
})

function selectSection(id: string) {
  const isNewSelection = selectedSectionId.value !== id;
  selectedSectionId.value = isNewSelection ? id : null;
  
  if (selectedSectionId.value) {
    // Scroll to editor
    isRightCollapsed.value = false;
    if (window.innerWidth <= 1100) isLeftCollapsed.value = true;
    nextTick(() => {
      scrollToEditor();
    });
  } else {
    // Scroll back to top (Theme Settings)
    nextTick(() => {
      if (propsScrollRef.value) {
        propsScrollRef.value.scrollTo({ top: 0, behavior: 'smooth' });
      }
    });
  }
}

function openSidebar(side: 'left' | 'right') {
  if (side === 'left') {
    isLeftCollapsed.value = false
    if (window.innerWidth <= 1100) isRightCollapsed.value = true
  } else {
    isRightCollapsed.value = false
    if (window.innerWidth <= 1100) isLeftCollapsed.value = true
  }
}

function scrollToEditor() {
  isRightCollapsed.value = false
  isLeftCollapsed.value = true
  
  // Wait for sidebar to expand before scrolling
  setTimeout(() => {
    if (editorRef.value) {
      editorRef.value.scrollIntoView({ behavior: 'smooth', block: 'start' })
    }
  }, 300)
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

function copyPublicLink() {
  const slug = store.activePortfolio?.slug
  if (!slug) return
  
  const url = `${window.location.origin}/p/${slug}`
  navigator.clipboard.writeText(url).then(() => {
    copied.value = true
    setTimeout(() => {
      copied.value = false
    }, 2000)
  })
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

.builder-topbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 24px;
  background: var(--sidebar-bg);
  border-bottom: 1px solid var(--border);
  flex-shrink: 0;
  z-index: 50;
  position: relative;
}
.topbar-left, .topbar-actions {
  flex: 1;
  display: flex;
}
.topbar-actions {
  justify-content: flex-end;
}
.topbar-center {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  align-items: center;
  gap: 16px;
}
.slug-container {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 4px 12px;
}
.portfolio-slug {
  font-size: 13px;
  color: var(--indigo);
  font-weight: 600;
  display: inline-flex;
  align-items: center;
  gap: 6px;
}
.copy-link-btn-indigo {
  background: transparent;
  border: none;
  color: var(--indigo);
  display: flex;
  align-items: center;
  gap: 4px;
  cursor: pointer;
  padding: 4px 8px;
  border-radius: 6px;
  font-size: 12px;
  font-weight: 600;
  transition: all 0.2s;
  margin-left: 4px;
}
.copy-link-btn-indigo:hover {
  background: rgba(79, 70, 229, 0.15);
}
.copy-link-btn-indigo:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}
.copy-text {
  font-size: 11px;
}
@media (max-width: 1100px) {
  .topbar-center {
    position: static !important;
    transform: none !important;
    left: auto !important;
  }
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
  display: flex;
  flex: 1;
  overflow: hidden;
  position: relative;
}

.block-palette {
  position: relative;
  width: 240px;
  flex-shrink: 0;
  background: var(--sidebar-bg);
  border-right: 1px solid var(--border);
  overflow: hidden; /* Hide content during resize/collapse */
  display: flex;
  flex-direction: column;
  transition: margin-left 0.4s cubic-bezier(0.4, 0, 0.2, 1), opacity 0.3s ease;
  z-index: 10;
}
.block-palette.collapsed {
  margin-left: calc(-1 * var(--sidebar-width, 240px));
  opacity: 0;
  pointer-events: none;
}

/* Resize Handles */
.resize-handle {
  position: absolute;
  top: 0;
  bottom: 0;
  width: 8px;
  cursor: col-resize;
  z-index: 1000;
  transition: background 0.2s;
}
.resize-handle:hover {
  background: var(--primary);
  opacity: 0.5;
}
.right-handle {
  right: 0;
}
.left-handle {
  left: 0;
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
.palette-list {
  flex: 1;
  overflow-y: auto;
}
.palette-item:hover .palette-add {
  opacity: 1;
}

/* Scrollable Inner Props */
.props-inner-scroll {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow-y: auto;
  overflow-x: hidden;
  height: 100%;
  padding: 20px 16px;
}

/* Canvas */
.canvas-wrapper {
  flex: 1;
  position: relative;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.builder-canvas {
  flex: 1;
  position: relative;
  overflow-y: auto;
  padding: 32px 16px;
  padding-bottom: 10px; /* Reduced from 500px for better balance */
  background: var(--bg) fixed;
}

.canvas-content-limiter {
  max-width: 1100px;
  margin: 0 auto;
  width: 100%;
}

.sidebar-toggle {
  position: absolute;
  top: 16px;
  background: var(--surface);
  border: 1px solid var(--border);
  width: 28px;
  height: 28px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--muted);
  cursor: pointer;
  z-index: 100; /* Ensure it stays above sections */
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  transition: all 0.2s ease;
}
.sidebar-toggle:hover {
  background: var(--bg);
  color: var(--text);
  transform: scale(1.1);
}
.left-toggle {
  left: 16px;
}
.right-toggle {
  right: 16px;
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

/* Layout and Grid */
.sections-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 40px;
  width: 100% !important;
  align-items: stretch; /* FORCE EQUAL HEIGHTS ON ROWS */
}
.tpl-firstjobber.sections-grid {
  gap: 20px !important;
}
.section-wrapper.span-full {
  grid-column: 1 / -1;
  width: 100%;
}
.section-wrapper.span-half {
  grid-column: span 1;
}

/* Animations */
.section-list-move,
.section-list-enter-active,
.section-list-leave-active {
  transition: all 0.4s ease;
}
.section-list-enter-from,
.section-list-leave-to {
  opacity: 0;
  transform: translateY(20px);
}
.section-list-leave-active {
  position: absolute;
  width: auto; 
}

/* Divider Implementation */
.sections-grid.has-divider .section-wrapper.span-half.is-right {
  position: relative;
  padding-left: 0; 
}
.sections-grid.has-divider .section-wrapper.span-half.is-right::before {
  content: "";
  position: absolute;
  top: 10%;
  bottom: 10%;
  left: -20px; /* Offset to visually span the gap */
  width: 1px;
  background-color: var(--border);
}

@media (max-width: 640px) {
  .sections-grid {
    grid-template-columns: 1fr;
  }
  .section-wrapper.span-half {
    grid-column: 1 / -1;
  }
  .sections-grid.has-divider .section-wrapper.span-half.is-right {
    padding-left: 0;
  }
  .sections-grid.has-divider .section-wrapper.span-half.is-right::before {
    display: none;
  }
}

/* Props Panel */
.props-panel {
  position: relative;
  width: 320px;
  flex-shrink: 0;
  background: var(--sidebar-bg);
  border-left: 1px solid var(--border);
  overflow: hidden;
  padding: 0; /* Padding moved to inner scroll */
  transition: margin-right 0.4s cubic-bezier(0.4, 0, 0.2, 1), opacity 0.3s ease;
  z-index: 10;
  display: flex;
  flex-direction: column;
}
.props-panel.collapsed {
  margin-right: calc(-1 * var(--sidebar-width, 320px));
  opacity: 0;
  pointer-events: none;
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
  margin-bottom: 3px;
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
  margin-top: 20px;
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
  margin-bottom: 2px;
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
  justify-content: center;
  gap: 6px;
}

/* Toggle Switch css */
.toggle-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  cursor: pointer;
}
.toggle-switch {
  position: relative;
  display: inline-block;
  width: 40px;
  height: 22px;
}
.toggle-switch input {
  opacity: 0;
  width: 0;
  height: 0;
}
.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: var(--border);
  transition: .4s;
  border-radius: 34px;
}
.slider:before {
  position: absolute;
  content: "";
  height: 16px;
  width: 16px;
  left: 3px;
  bottom: 3px;
  background-color: white;
  transition: .4s;
  border-radius: 50%;
}
input:checked + .slider {
  background-color: var(--primary);
}
input:checked + .slider:before {
  transform: translateX(18px);
}

/* Mobile Tabs at Top */
.builder-mobile-tabs {
  display: none;
  gap: 4px;
  background: var(--surface);
  padding: 4px;
  border-radius: 12px;
  border: 1px solid var(--border);
}

@media (max-width: 768px) {
  .builder-mobile-tabs {
    display: flex;
  }
  .topbar-center {
    display: none !important;
  }
}

.mobile-tab-btn {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  border: none;
  background: transparent;
  color: var(--muted);
  cursor: pointer;
  transition: all 0.2s ease;
}

.mobile-tab-btn.active {
  background: var(--primary);
  color: #fff;
}

.mobile-tab-btn.hide-btn {
  color: var(--muted);
  margin-left: 8px;
  border-left: 1px solid var(--border);
  border-radius: 0;
  width: 32px;
}

.mobile-close-sidebar {
  display: none;
  background: transparent;
  border: none;
  color: var(--muted);
  cursor: pointer;
  padding: 4px;
}

@media (max-width: 768px) {
  .mobile-close-sidebar {
    display: flex;
  }
}

/* Responsive Utilities */
.mobile-hide {
  display: flex;
}

@media (max-width: 768px) {
  .mobile-hide {
    display: none !important;
  }
}

.btn-icon-only-mobile {
  padding: 8px !important;
  min-width: 40px !important;
}

@media (max-width: 768px) {
  .btn-icon-only-mobile {
    padding: 8px !important;
  }
}

/* Sidebar Overlays on Mobile */
@media (max-width: 768px) {
  .block-palette, .props-panel {
    position: fixed !important;
    top: 0;
    bottom: 0;
    width: 280px !important;
    z-index: 1000;
    margin: 0 !important;
    transform: none !important;
    box-shadow: 0 0 50px rgba(0,0,0,0.5);
    background: var(--surface) !important;
  }
  
  .block-palette {
    left: 0;
  }
  
  .props-panel {
    right: 0;
    width: 320px !important;
  }

  /* Overlay hidden state */
  .block-palette.collapsed {
    transform: translateX(-100%) !important;
  }
  .props-panel.collapsed {
    transform: translateX(100%) !important;
  }

  .builder-canvas {
    padding: 16px;
  }

  .sections-grid {
    grid-template-columns: 1fr !important;
    gap: 20px !important;
  }
}

.pulse-hint {
  color: var(--primary);
  animation: pulse-soft 2s infinite;
}

@keyframes pulse-soft {
  0% { transform: scale(1); opacity: 1; }
  50% { transform: scale(1.05); opacity: 0.8; }
  100% { transform: scale(1); opacity: 1; }
}

.builder-layout.theme-dark {
  --divider-color: rgba(255, 255, 255, 0.15);
}

.builder-layout.theme-light {
  --divider-color: rgba(0, 0, 0, 0.15);
}

/* Hide handles on mobile */
@media (max-width: 768px) {
  .resize-handle {
    display: none !important;
  }
}
</style>
