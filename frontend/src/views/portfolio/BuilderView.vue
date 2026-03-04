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

    <!-- Mobile Backdrop -->
    <Transition name="fade-backdrop">
      <div 
        v-if="(!isLeftCollapsed || !isRightCollapsed) && isMobile" 
        class="mobile-backdrop"
        @click="isLeftCollapsed = true; isRightCollapsed = true"
      />
    </Transition>

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
          <div class="canvas-content-limiter" :class="[templateClass, themeClass]" :style="{ 'background-color': 'transparent !important', border: 'none !important' }">
          <!-- Canvas Loading and Empty states -->
          <div v-if="store.loading" class="canvas-loading">...</div>
          <div v-else-if="store.sections.length === 0" class="canvas-empty">...</div>

          <TransitionGroup
            tag="div"
            name="section-list"
            v-else
            class="sections-list sections-grid"
            :class="{ 'has-divider': theme.show_divider }"
            ref="listRef"
          >
            <div
              v-for="section in gridSections"
              :key="section.id"
              :data-id="section.id"
              class="section-wrapper drag-item"
              :style="themeVars"
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
                :theme-vars="themeVars"
                :is-half-split="section.column_span === 'half'"
                :class="{ 'block-highlight': highlightedBlocks.has(section.id) }"
                @select="selectSection(section.id)"
                @delete="deleteSection(section.id)"
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
          <div class="url-display copyable-url" @click="copyPublicLink" :title="copied ? 'Copied!' : 'Click to copy'">
            <span class="url-prefix">pk.io/</span>
            <span class="url-slug">{{ store.activePortfolio?.slug }}</span>
            <Copy v-if="!copied" :size="12" style="margin-left: auto; opacity: 0.5" />
            <Check v-else :size="12" style="margin-left: auto; color: var(--success)" />
          </div>
        </div>

        <!-- AI Actions — above Section Editor, with inline preview -->
        <div class="ai-panel">
          <div class="ai-panel-title">
            <Bot :size="14" class="icon-inline" /> AI Assistant
          </div>
          
          <template v-if="!aiLoading && !aiResult">
            <!-- Magic Fill Button -->
            <button
              class="btn btn-primary btn-sm btn-icon-text magic-fill-btn magic-fill-btn-inline"
              style="width: 100%; margin-bottom: 12px"
              :disabled="!isHeroComplete"
              @click="aiMagicFill"
            >
              <Zap :size="14" /> Auto-fill Section 
            </button>
            <p v-if="!isHeroComplete" 
               style="font-size: 10px; color: var(--danger); margin-bottom: 12px; text-align: center;">
              *กรุณากรอกข้อมูลที่ส่วน Hero ให้ครบก่อน
            </p>

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
          </template>

          <!-- AI Loading Animation -->
          <div v-if="aiLoading" class="ai-loading">
            <div class="ai-pulse">
              <Sparkles :size="24" class="sparkle-icon" />
              <div class="pulse-ring"></div>
            </div>
            <p>AI กำลังวิเคราะห์พอร์ตของคุณ...</p>
          </div>

          <template v-else-if="aiResult">
            <!-- Score result -->
            <template v-if="isJson(aiResult) && lastAIAction === 'score'">
              <div class="ai-result-wrap">
                <AIAnalysisResult :result="aiResult" />
                <div style="display: flex; gap: 8px; margin-top: 12px;">
                  <button @click="aiResult = null" class="btn btn-outline btn-xs" style="width: 100%">Close</button>
                </div>
              </div>
            </template>

            <!-- Improve / Magic — full diff preview inline in AI panel -->
            <template v-else-if="lastAIAction === 'improve' || lastAIAction === 'magic'">
              <Transition name="ai-preview">
                <div class="ai-improve-preview">
                  <div class="aip-header">
                    <div class="aip-title">
                      <Sparkles :size="14" class="icon-inline" /> AI Suggestion Preview
                    </div>
                    <span class="aip-badge">{{ lastAIAction === 'magic' ? 'Auto-fill' : 'Improved' }}</span>
                  </div>

                  <div class="aip-body">
                    <div class="aip-col">
                      <div class="aip-col-label">
                        {{ lastAIAction === 'improve' ? '📄 ข้อความเดิม' : '📄 ข้อความเดิม / ตัวอย่างจากที่อื่นๆ' }}
                      </div>
                      <div class="aip-text aip-original" v-html="renderHumanReadable(aiSourceData)" />
                    </div>
                    <div class="aip-divider-v">
                      <div class="aip-arrow">→</div>
                    </div>
                    <div class="aip-col">
                      <div class="aip-col-label">✨ AI แนะนำ</div>
                      <div class="aip-text aip-new" v-html="renderHumanReadable(parseAIResult(aiResult))" />
                    </div>
                  </div>

                  <div class="aip-actions">
                    <button @click="aiResult = null; lastAIAction = null" class="btn btn-secondary">
                      Discard
                    </button>
                    <button @click="applyAIResult" class="btn btn-primary aip-apply-btn">
                      <Check :size="16" /> Apply Changes
                    </button>
                  </div>
                </div>
              </Transition>
            </template>

            <!-- Fallback text result -->
            <template v-else>
              <div class="ai-result-wrap">
                <div class="ai-text-result">{{ aiResult }}</div>
                <div style="display: flex; gap: 8px; margin-top: 12px;">
                  <button @click="aiResult = null" class="btn btn-outline btn-xs" style="width: 100%">Close</button>
                </div>
              </div>
            </template>
          </template>
        </div>

        <!-- Section Editor (when section selected) -->
        <div v-if="selectedSection" class="section-editor" ref="editorRef">
          <div class="editor-divider"></div>
          <div class="props-title" style="margin-bottom: 12px;">
            <Pencil :size="12" class="icon-inline" /> Edit:
            {{ selectedSection.type }}
          </div>

          <!-- Section Toolbar -->
          <div class="section-toolbar">
            <div class="toolbar-group">
              <button class="toolbar-btn" @click="moveSection(selectedSection.id, 'up')" title="Move Up">
                <ChevronUp :size="16" />
              </button>
              <button class="toolbar-btn" @click="moveSection(selectedSection.id, 'down')" title="Move Down">
                <ChevronDown :size="16" />
              </button>
            </div>
            
            <div class="toolbar-divider"></div>
            
            <div class="toolbar-group">
              <button 
                class="toolbar-btn" 
                @click="store.toggleSectionColumnSpan(selectedSection.id)"
                :title="selectedSection.column_span === 'half' ? 'Full Width' : 'Half Width'"
                v-if="selectedSection.type !== 'hero'"
              >
                <Columns v-if="selectedSection.column_span === 'half'" :size="16" />
                <Square v-else :size="16" />
              </button>
              <button class="toolbar-btn" @click="handleDuplicate(selectedSection)" title="Duplicate">
                <Copy :size="16" />
              </button>
            </div>

            <div class="toolbar-divider"></div>

            <button class="toolbar-btn btn-danger-soft" @click="deleteSection(selectedSection.id)" title="Delete">
              <Trash2 :size="16" />
            </button>
          </div>

          <!-- Column Divider Toggle: Only show for half-width sections -->
          <div v-if="selectedSection.column_span === 'half' && theme.template !== 'firstjobber'" class="props-section" style="background: rgba(var(--primary-glow), 0.05); padding: 12px; border-radius: 12px; border: 1px dashed var(--border); margin-bottom: 20px;">
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
            :is-hero-complete="isHeroComplete"
            @update="handleSectionUpdate"
            @magic-fill="aiMagicFill"
          />
        </div>
      </div>
    </aside>
    </div>

    <!-- ─── Mobile Bottom Navigation Bar ─── -->
    <nav class="mobile-bottom-nav">
      <button 
        class="mob-nav-btn" 
        :class="{ active: !isLeftCollapsed }"
        @click="openSidebar('left')"
      >
        <Blocks :size="22" />
        <span>Blocks</span>
      </button>
      <button 
        class="mob-nav-btn" 
        :class="{ active: store.activePortfolio?.is_published }"
        @click="previewPortfolio"
        :disabled="!store.activePortfolio?.is_published"
      >
        <Eye :size="22" />
        <span>Preview</span>
      </button>
      <button 
        class="mob-nav-btn" 
        :class="{ active: !isRightCollapsed && !selectedSectionId }"
        @click="openSidebar('right')"
      >
        <Palette :size="22" />
        <span>Design</span>
      </button>
      <button 
        v-if="selectedSectionId"
        class="mob-nav-btn pulse-hint" 
        :class="{ active: !isRightCollapsed && !!selectedSectionId }"
        @click="scrollToEditor"
      >
        <Pencil :size="22" />
        <span>Edit</span>
      </button>
      <button 
        class="mob-nav-btn"
        :class="store.activePortfolio?.is_published ? 'btn-danger-mobile' : 'btn-primary-mobile'"
        @click="togglePublish"
      >
        <Play v-if="store.activePortfolio?.is_published" :size="22" />
        <Zap v-else :size="22" />
        <span>{{ store.activePortfolio?.is_published ? 'Unpublish' : 'Publish' }}</span>
      </button>
    </nav>
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
  ChevronDown,
  Columns,
  Square,
  Trash2,
  X,
  Copy,
  Check
} from 'lucide-vue-next'
import SectionBlock from '@/components/builder/SectionBlock.vue'
import SectionEditor from '@/components/builder/SectionEditor.vue'
import { debounce } from 'lodash'
import Sortable from 'sortablejs'
import AIAnalysisResult from '@/components/builder/AIAnalysisResult.vue'
import { toastError, toastSuccess } from '@/utils/alert'

const route = useRoute()
const store = usePortfolioStore()

const selectedSectionId = ref<string | null>(null)
const aiResult = ref<string | null>(null)
const aiSourceData = ref<any>(null)      // snapshot of data when AI call was made
const aiSourceSectionId = ref<string | null>(null) // section id (null = all sections)
const lastAIAction = ref<'score' | 'improve' | 'magic' | null>(null)
const aiLoading = ref(false)

function isJson(str: string) {
  try {
    const trimmed = str.trim();
    // Detect Object or Array, also handle markdown-wrapped if needed
    const isWrapped = trimmed.startsWith('```json') && trimmed.endsWith('```');
    const content = isWrapped ? trimmed.replace(/^```json/, '').replace(/```$/, '').trim() : trimmed;
    
    return (content.startsWith('{') && content.endsWith('}')) || 
           (content.startsWith('[') && content.endsWith(']'));
  } catch (e) { return false; }
}
const canvasRef = ref<HTMLElement | null>(null)
const listRef = ref<any>(null)
const editorRef = ref<HTMLElement | null>(null)
const propsScrollRef = ref<HTMLElement | null>(null)
const highlightedBlocks = ref<Set<string>>(new Set())

const isLeftCollapsed = ref(true)
const isRightCollapsed = ref(true)
const leftSidebarWidth = ref(260)
const rightSidebarWidth = ref(320)
const copied = ref(false)

// Resizing logic
const isResizing = ref<'left' | 'right' | null>(null)
const isMobile = ref(window.innerWidth <= 768)

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
  font: 'Sarabun',
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

const isHeroComplete = computed(() => {
  const hero = store.sections.find(s => s.type === 'hero');
  if (!hero || !hero.data) return false;
  return !!hero.data.name && !!hero.data.role; // tagline is optional
});

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
    '--primary-glow': hexToRgb(primary),
    '--secondary': secondary ? secondary : primary,
    '--font-display': `"${font}", sans-serif`,
    '--font-body': `"${font}", sans-serif`,
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
    vars['--section-border'] = theme.border_color
  } else {
    vars['--section-border'] = `rgba(${hexToRgb(primary)}, 0.2)`
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
  // Lock Hero and AI Chat: only one allowed per portfolio
  if (type === 'hero' || type === 'ai_chat') {
    const exists = store.sections.some(s => s.type === type);
    if (exists) {
      toastError(`สามารถเพิ่มส่วน ${type === 'hero' ? 'Hero' : 'AI Chat'} ได้เพียงอันเดียวต่อพอร์ตโฟลิโอ`);
      return;
    }
  }

  const newSectionId = await store.addSection(type, selectedSectionId.value)
  if (!newSectionId) return

  // Auto-scroll to bottom only if no section was selected (appending to end)
  if (!selectedSectionId.value) {
    await nextTick()
    if (canvasRef.value) {
      canvasRef.value.scrollTo({
        top: canvasRef.value.scrollHeight,
        behavior: 'smooth'
      })
    }
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

  // Use the API's Public URL to trigger the SEO/Crawler handler
  const apiBase = import.meta.env.VITE_API_URL || '/api/v1'
  const url = `${apiBase}/public/p/${slug}`
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

function deleteSection(id: string) {
  debouncedSave.cancel()
  store.deleteSection(id)
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

function handleDuplicate(section: any) {
  if (!section) return;
  if (section.type === 'hero' || section.type === 'ai_chat') {
    toastError(`สามารถมีส่วน ${section.type === 'hero' ? 'Hero' : 'AI Chat'} ได้เพียงอันเดียวต่อพอร์ตโฟลิโอ`);
    return;
  }
  store.duplicateSection(section.id);
}

async function aiScoreResume() {
  aiLoading.value = true
  aiResult.value = null
  lastAIAction.value = 'score'
  const content = JSON.stringify(store.sections.map((s) => s.data))
  try {
    const { data } = await aiAPI.scoreResume(content, 'Software Developer')
    aiResult.value = data.data.analysis
  } catch {
    aiResult.value = 'AI unavailable. Please check API key.'
  } finally {
    aiLoading.value = false
  }
}

async function aiImproveContent() {
  aiLoading.value = true
  aiResult.value = null
  lastAIAction.value = 'improve'

  const selected = store.sections.find((s) => s.id === selectedSectionId.value)

  if (selected) {
    // Improve selected section
    aiSourceSectionId.value = selected.id
    aiSourceData.value = JSON.parse(JSON.stringify(selected.data)) // snapshot
    const text = JSON.stringify(selected.data)
    try {
      const { data } = await aiAPI.improveText(text, `${selected.type} section`)
      aiResult.value = data.data.improved_text
    } catch {
      aiResult.value = 'AI unavailable.'
    } finally {
      aiLoading.value = false
    }
  } else {
    // No section selected — improve entire portfolio
    aiSourceSectionId.value = null
    const allData = store.sections.map(s => ({ type: s.type, data: s.data }))
    aiSourceData.value = allData // snapshot of all
    const text = JSON.stringify(allData)
    try {
      const { data } = await aiAPI.improveText(text, 'full portfolio resume')
      aiResult.value = data.data.improved_text
    } catch {
      aiResult.value = 'AI unavailable.'
    } finally {
      aiLoading.value = false
    }
  }
}

const MAGIC_SKIP_TYPES = new Set(['ai_chat', 'custom_text', 'hero'])

// Helper to clean AI response (remove markdown blocks)
function getCleanJson(raw: string): string {
  let clean = raw.trim();
  if (clean.startsWith('```')) {
    clean = clean.replace(/^```(json)?/, '').replace(/```$/, '').trim();
  }
  return clean;
}

// Helper to ensure list-based sections get an object { items: [] } if result is an array
function wrapIfNeeded(data: any, sectionType: string): any {
  if (Array.isArray(data)) {
    const listTypes = ['skills', 'experience', 'projects', 'certificates', 'education', 'stats', 'social'];
    if (listTypes.includes(sectionType)) {
      return { items: data };
    }
  }
  return data;
}

async function aiMagicFill() {
  const selected = store.sections.find((s) => s.id === selectedSectionId.value)

  // Context data always needed
  const previousData = JSON.stringify(
    store.portfolios
      .filter(p => p.id !== store.activePortfolio?.id)
      .map(p => ({ title: p.title, sections: p.sections?.map(s => s.data) }))
      .slice(0, 3)
  )
  const heroData = JSON.stringify(store.sections.find(s => s.type === 'hero')?.data || {})
  const template = store.activePortfolio?.theme?.template || 'classic'

  // ── Case 1: selected section that can't be auto-filled ──
  if (selected && MAGIC_SKIP_TYPES.has(selected.type)) {
    toastError(`ไม่สามารถ Auto-fill Section ประเภท "${selected.type}" ได้`)
    return
  }

  // ── Case 2: single section fill (with preview) ──
  if (selected) {
    aiLoading.value = true
    aiResult.value = null
    lastAIAction.value = 'magic'
    aiSourceSectionId.value = selected.id
    aiSourceData.value = JSON.parse(JSON.stringify(selected.data))

    try {
      const { data } = await aiAPI.magicFill({ section_type: selected.type, previous_data: previousData, hero_data: heroData, template })
      aiResult.value = typeof data.data === 'string' ? data.data : JSON.stringify(data.data, null, 2)
    } catch (e: any) {
      aiResult.value = 'AI Failed: ' + (e.response?.data?.error || e.message)
    } finally {
      aiLoading.value = false
    }
    return
  }

  // ── Case 3: no section selected → fill ALL fillable sections directly ──
  const fillable = store.sections.filter(s => !MAGIC_SKIP_TYPES.has(s.type))
  if (fillable.length === 0) {
    toastError('ไม่มี Section ที่สามารถ Auto-fill ได้')
    return
  }

  aiLoading.value = true
  aiResult.value = null
  lastAIAction.value = null // no preview for batch fill
  aiSourceSectionId.value = null
  aiSourceData.value = null

  let successCount = 0
  try {
    for (const sec of fillable) {
      try {
        const { data } = await aiAPI.magicFill({ section_type: sec.type, previous_data: previousData, hero_data: heroData, template })
        const rawResult = typeof data.data === 'string' ? data.data : JSON.stringify(data.data, null, 2)

        // Parse and apply immediately
        let clean = getCleanJson(rawResult);
        let parsed: any;
        try { parsed = JSON.parse(clean) } catch { parsed = clean }

        // Wrap if it's a list section but result is a raw array
        const finalData = wrapIfNeeded(parsed, sec.type);

        // Apply to section directly
        const secRef = store.sections.find(s => s.id === sec.id)
        if (secRef && typeof finalData === 'object' && finalData !== null) {
          secRef.data = finalData;
          debouncedSave(sec.id, finalData);
          successCount++;
        }
      } catch {
        // skip failed sections silently
      }
    }
    toastSuccess(`✨ Auto-fill สำเร็จ ${successCount}/${fillable.length} Section`)
  } finally {
    aiLoading.value = false
  }
}

function applyAIResult() {
  if (!aiResult.value) return;
  
  try {
    let raw = aiResult.value;
    let clean = getCleanJson(raw);
    let resultToApply: any;

    const targetSec = aiSourceSectionId.value ? store.sections.find(s => s.id === aiSourceSectionId.value) : null;

    if (isJson(clean)) {
      let parsed = JSON.parse(clean);
      
      // Smart field mapping for Stats (AI often uses 'name' or 'description' instead of 'label')
      if (targetSec?.type === 'stats' || targetSec?.type === 'skills') {
        const items = Array.isArray(parsed) ? parsed : (parsed.items || []);
        if (Array.isArray(items)) {
          items.forEach(it => {
            if (!it.label && it.name) it.label = it.name;
            if (!it.label && it.title) it.label = it.title;
            if (!it.value && it.number) it.value = it.number;
            if (!it.value && it.count) it.value = it.count;
            if (!it.value && it.amount) it.value = it.amount;
          });
          if (!Array.isArray(parsed)) parsed.items = items;
          else parsed = items;
        }
      }

      // Merge with existing data to prevent losing fields like 'items' if AI only returns 'title'
      if (aiSourceData.value && typeof aiSourceData.value === 'object' && !Array.isArray(parsed)) {
        resultToApply = { ...aiSourceData.value, ...parsed };
      } else {
        resultToApply = wrapIfNeeded(parsed, targetSec?.type || '');
      }
    } else {
      // Plain text result — merge into existing section data
      const sourceSection = aiSourceSectionId.value 
        ? store.sections.find(s => s.id === aiSourceSectionId.value)
        : null;
      
      if (sourceSection && typeof aiSourceData.value === 'object' && aiSourceData.value !== null) {
        const sectionType = sourceSection.type;
        const currentData = { ...aiSourceData.value };

        if (sectionType === 'custom_text') {
          resultToApply = { ...currentData, content: clean };
        } else if (sectionType === 'hero') {
          resultToApply = { ...currentData, bio: clean };
        } else if (['stats', 'skills', 'experience', 'projects'].includes(sectionType)) {
          // List-based section got plain text: don't overwrite the structure
          console.warn(`[AI] Section ${sectionType} received plain text. Keeping structure.`);
          resultToApply = currentData; 
        } else {
          const descField = ['description', 'content', 'bio', 'tagline', 'text'].find(f => f in currentData);
          if (descField) {
            resultToApply = { ...currentData, [descField]: clean };
          } else {
            resultToApply = currentData; 
          }
        }
      } else {
        resultToApply = clean;
      }
    }

    if (aiSourceSectionId.value) {
      // Apply to specific section
      selectedSectionId.value = aiSourceSectionId.value;
      handleSectionUpdate(resultToApply);
    } else if (Array.isArray(resultToApply)) {
      // All-portfolio improve — apply each section
      resultToApply.forEach((item: any) => {
        const sec = store.sections.find(s => s.type === item.type);
        if (sec && item.data) {
          const sec2 = store.sections.find(s => s.id === sec.id);
          const existingData = sec2 ? { ...sec2.data } : {};
          const mergedData = (typeof item.data === 'object' && !Array.isArray(item.data)) 
            ? { ...existingData, ...item.data } 
            : wrapIfNeeded(item.data, sec.type);
            
          if (sec2) sec2.data = mergedData;
          debouncedSave(sec.id, mergedData);
        }
      });
    } else {
      // Fallback: apply to currently selected section
      if (!selectedSectionId.value) return;
      handleSectionUpdate(resultToApply);
    }

    aiResult.value = null;
    aiSourceData.value = null;
    aiSourceSectionId.value = null;
    lastAIAction.value = null;
  } catch (e) {
    console.error('Failed to parse AI result:', e);
  }
}

// Parse AI result (might be JSON string or plain text)
function parseAIResult(raw: string | null): any {
  if (!raw) return null;
  try {
    let clean = raw.trim();
    if (clean.startsWith('```json')) clean = clean.replace(/^```json/, '').replace(/```$/, '').trim();
    else if (clean.startsWith('```')) clean = clean.replace(/^```/, '').replace(/```$/, '').trim();
    return JSON.parse(clean);
  } catch {
    return raw; // return as plain string if not JSON
  }
}

// Convert section data to readable HTML rows
function renderHumanReadable(data: any, depth = 0): string {
  if (data === null || data === undefined) return '<span class="aip-empty">—</span>';
  if (typeof data === 'string') {
    if (!data.trim()) return '<span class="aip-empty">—</span>';
    return `<span class="aip-value">${escHtml(data)}</span>`;
  }
  if (typeof data === 'number' || typeof data === 'boolean') {
    return `<span class="aip-value">${data}</span>`;
  }
  if (Array.isArray(data)) {
    if (data.length === 0) return '<span class="aip-empty">— (ว่าง)</span>';
    return data.map((item, i) => {
      if (typeof item === 'object' && item !== null) {
        return `<div class="aip-array-item"><div class="aip-item-num">${i + 1}</div><div class="aip-item-body">${renderHumanReadable(item, depth + 1)}</div></div>`;
      }
      return `<div class="aip-array-item"><div class="aip-item-num">${i + 1}</div><span class="aip-value">${escHtml(String(item))}</span></div>`;
    }).join('');
  }
  if (typeof data === 'object') {
    const SKIP_KEYS = new Set(['id', 'sectionId', 'prompt_hint', 'avatar_url']);
    const LABELS: Record<string, string> = {
      name: 'ชื่อ', role: 'ตำแหน่ง', bio: 'Bio', title: 'หัวข้อ',
      description: 'คำอธิบาย', company: 'บริษัท', period: 'ช่วงเวลา',
      location: 'สถานที่', skills: 'ทักษะ', projects: 'โปรเจกต์',
      experiences: 'ประสบการณ์', items: 'รายการ', content: 'เนื้อหา',
      url: 'URL', link: 'ลิงก์', email: 'อีเมล', phone: 'เบอร์โทร',
      tech: 'เทคโนโลยี', level: 'ระดับ', category: 'หมวดหมู่',
      degree: 'วุฒิการศึกษา', school: 'สถาบัน', gpa: 'เกรด',
      tagline: 'Tagline', label: 'ป้ายข้อความ', value: 'ค่า',
      highlight: 'ไฮไลท์', achievements: 'ผลงาน', responsibilities: 'หน้าที่'
    };
    const entries = Object.entries(data).filter(([k]) => !SKIP_KEYS.has(k));
    if (entries.length === 0) return '<span class="aip-empty">—</span>';
    return entries.map(([key, val]) => {
      const label = LABELS[key] || key;
      if (val === null || val === undefined || val === '') return '';
      const isComplex = typeof val === 'object' || Array.isArray(val);
      return `<div class="aip-field ${depth > 0 ? 'aip-field-nested' : ''}">
        <div class="aip-field-key">${escHtml(label)}</div>
        <div class="aip-field-val">${renderHumanReadable(val, depth + 1)}</div>
      </div>`;
    }).filter(Boolean).join('');
  }
  return escHtml(String(data));
}

function escHtml(str: string): string {
  return str.replace(/&/g,'&amp;').replace(/</g,'&lt;').replace(/>/g,'&gt;').replace(/"/g,'&quot;');
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

.magic-fill-btn-inline {
  background: var(--primary);
  color: var(--text);
  border: 1px solid var(--border);
  border-radius: 10px;
  padding: 10px 16px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: 
    background 0.2s ease,
    transform 0.15s ease,
    box-shadow 0.2s ease,
    opacity 0.2s ease;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

/* Hover */
.magic-fill-btn-inline:hover:not(:disabled) {
  background: var(--primary-glow);
  transform: translateY(-2px);
  box-shadow: 0 6px 18px rgba(0, 0, 0, 0.15);
}

/* Active (กดลง) */
.magic-fill-btn-inline:active:not(:disabled) {
  transform: translateY(0);
  box-shadow: 0 3px 10px rgba(0, 0, 0, 0.12);
}

/* Disabled */
.magic-fill-btn-inline:disabled {
  background: var(--surface-muted, rgba(128,128,128,0.15));
  color: var(--text-muted, rgba(128,128,128,0.7));
  border-color: var(--border-muted, rgba(128,128,128,0.25));
  cursor: not-allowed;
  transform: none;
  box-shadow: none;
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
  font-family: var(--font-body) !important;
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
/* Redundant rules removed — now handled by template-styles.css */

@media (max-width: 768px) {
  .sections-grid.tpl-business {
    grid-template-columns: 1fr !important;
    gap: 16px !important;
    background: transparent !important;
  }
  
  .sections-grid.tpl-business > .section-wrapper.is-row-start:not(.span-full),
  .sections-grid.tpl-business > .section-wrapper.is-right:not(.span-full) {
    width: 100% !important;
  }

  .sections-grid.tpl-business > .section-wrapper.span-full {
    background: var(--bg) !important; /* Stacked, no need for horizontal gradient */
  }

  .sections-grid.tpl-business > .section-wrapper.span-full :deep(.pub-section) {
    margin-left: 0 !important;
    width: 100% !important;
    padding: 24px 16px !important;
  }

  /* Reset layout-split on mobile */
  .sections-grid.tpl-business > .section-wrapper :deep(.pub-section.layout-split) {
    flex-direction: column !important;
  }
  .sections-grid.tpl-business > .section-wrapper :deep(.pub-section.layout-split .section-header-wrapper) {
    max-width: 100% !important;
    margin-bottom: 20px !important;
  }
}

/* Business Builder: style cards */
.tpl-business.sections-grid > .section-wrapper {
  border-radius: 12px !important;
  overflow: hidden;
}
.tpl-business.sections-grid > .section-wrapper :deep(.section-block) {
  border-radius: 12px !important;
}
.tpl-business.sections-grid > .section-wrapper :deep(.section-preview) {
  border-radius: 0 0 12px 12px !important;
}

/* Business Builder: make section-block transparent in sidebar area to show wrapper bg */
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.section-block),
.tpl-business.sections-grid > .section-wrapper.span-full :deep(.section-block) {
  background: transparent !important;
}

/* 
   REMOVED: Portfolio theme dependent section-header styling. 
   SectionBlock now manages its own stable editor chrome colors.
*/


/* Business Builder: span-full → content in right 70% */
.tpl-business.sections-grid > .section-wrapper.span-full :deep(.pub-section) {
  margin-left: 30% !important;
  width: 70% !important;
  padding: 40px 40px 40px 40px !important;
  background: transparent !important;
}

/* Business Builder: is-row-start → white title */
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.layered-title) {
  color: #fff !important;
  -webkit-text-fill-color: #fff !important;
}

/* Business Builder: is-right → primary title */
.tpl-business.sections-grid > .section-wrapper.is-right:not(.span-full) :deep(.layered-title) {
  color: var(--primary) !important;
  -webkit-text-fill-color: var(--primary) !important;
}

/* Business Builder: sidebar contact color */
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.contact-val),
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.contact-label),
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.custom-content) {
  color: #000 !important;
}

/* Business Builder: force all content to full width */
.tpl-business.sections-grid > .section-wrapper :deep(.pub-section-content),
.tpl-business.sections-grid > .section-wrapper :deep(.experience-list),
.tpl-business.sections-grid > .section-wrapper :deep(.contact-grid),
.tpl-business.sections-grid > .section-wrapper :deep(.skills-groups),
.tpl-business.sections-grid > .section-wrapper :deep(.skills-grid),
.tpl-business.sections-grid > .section-wrapper :deep(.proj-grid),
.tpl-business.sections-grid > .section-wrapper :deep(.education-list),
.tpl-business.sections-grid > .section-wrapper :deep(.custom-content) {
  max-width: 100% !important;
  width: 100% !important;
  margin: 0 !important;
}

/* Business Builder: 30% column - skills single column for bars */
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.skills-grid:not(.grid-tags)) {
  grid-template-columns: 1fr !important;
}

/* Business Builder: 30% column - tags flex wrap */
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.grid-tags) {
  display: flex !important;
  flex-wrap: wrap !important;
  justify-content: center !important;
  gap: 8px !important;
}

.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.skill-tag) {
  width: auto !important;
  flex: none !important;
}

/* Business Builder: 30% column - skills-groups single column */
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.skills-groups) {
  grid-template-columns: 1fr !important;
}

/* Business Builder: 30% column - pub-section padding */
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.pub-section) {
  padding: 24px 16px !important;
  overflow: hidden !important;
}

/* Business Builder: 30% column - experience single column & centered */
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.experience-item) {
  grid-template-columns: 1fr !important;
  gap: 4px !important;
  text-align: center !important;
}
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.tl-header) {
  align-items: center !important;
}
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.timeline-content) {
  text-align: center !important;
}

/* Business Builder: 30% column - white text for contrast */
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.tl-position),
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.tl-company),
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.tl-desc),
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.tl-date),
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.group-title),
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.stat-label),
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.contact-label),
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.contact-value),
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.custom-content) {
  color: rgba(255,255,255,0.85) !important;
}
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.skill-tag) {
  border-color: rgba(255,255,255,0.3) !important;
  background: rgba(255,255,255,0.1) !important;
  color: #fff !important;
}
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.contact-card) {
  border-color: rgba(255,255,255,0.2) !important;
  background: rgba(255,255,255,0.08) !important;
}
.tpl-business.sections-grid > .section-wrapper.is-row-start:not(.span-full) :deep(.contact-icon) {
  color: #fff !important;
}

/* Business Builder: layout-split inside sections */
.tpl-business.sections-grid > .section-wrapper :deep(.pub-section.layout-split) {
  display: flex !important;
  flex-direction: row !important;
  gap: 16px !important;
  align-items: flex-start !important;
}
.tpl-business.sections-grid > .section-wrapper :deep(.pub-section.layout-split .section-header-wrapper) {
  flex: 0 0 auto !important;
  max-width: 30% !important;
}
.tpl-business.sections-grid > .section-wrapper :deep(.pub-section.layout-split .pub-section-content) {
  flex: 1 !important;
  min-width: 0 !important;
}

.section-wrapper.span-full {
  grid-column: 1 / -1;
  width: 100%;
}
.section-wrapper.span-half {
  grid-column: span 1;
}
.tpl-business.sections-grid > .section-wrapper.span-half {
  overflow: hidden !important;
  min-width: 0 !important;
  max-width: 100% !important;
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
.copyable-url {
  cursor: pointer;
  transition: all 0.2s;
  display: flex;
  align-items: center;
}
.copyable-url:hover {
  background: rgba(79, 70, 229, 0.05);
  border-color: var(--indigo);
}
.magic-fill-btn {
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  border: none;
  animation: shimmer 2s infinite linear;
  background-size: 200% auto;
}
@keyframes shimmer {
  to { background-position: 200% center; }
}

.section-editor {
  margin-bottom: 3px;
  position: relative;
}

.section-toolbar {
  display: flex;
  align-items: center;
  gap: 4px;
  background: var(--surface);
  border: 1px solid var(--border);
  padding: 6px;
  border-radius: 12px;
  margin-bottom: 20px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.05);
}

.toolbar-group {
  display: flex;
  gap: 4px;
}

.toolbar-divider {
  width: 1px;
  height: 20px;
  background: var(--border);
  margin: 0 4px;
}

.toolbar-btn {
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  border: none;
  background: transparent;
  color: var(--muted);
  cursor: pointer;
  transition: all 0.2s;
}

.toolbar-btn:hover {
  background: rgba(var(--primary-glow), 0.1);
  color: var(--primary);
}

.toolbar-btn.btn-danger-soft {
  color: #ef4444;
}

.toolbar-btn.btn-danger-soft:hover {
  background: rgba(239, 68, 68, 0.1);
  color: #ef4444;
}

/* ─── AI Improve Preview ────────────────────────── */
.ai-improve-preview {
  margin-top: 16px;
  border: 1.5px solid var(--indigo);
  border-radius: 16px;
  background: var(--surface);
  overflow: hidden;
  box-shadow: 0 0 24px rgba(79,70,229,0.12);
  container-type: inline-size; /* enables @container queries on children */
}

.aip-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 16px;
  background: rgba(79,70,229,0.08);
  border-bottom: 1px solid var(--border);
}

.aip-title {
  font-size: 12px;
  font-weight: 700;
  color: var(--indigo);
  display: flex;
  align-items: center;
  gap: 6px;
  letter-spacing: 0.3px;
}

.aip-badge {
  font-size: 10px;
  font-weight: 800;
  padding: 2px 8px;
  border-radius: 100px;
  background: rgba(79,70,229,0.15);
  color: var(--indigo);
  letter-spacing: 0.5px;
}

.aip-body {
  display: grid;
  grid-template-columns: 1fr 24px 1fr;
  gap: 0;
  min-height: 200px;
  max-height: 420px;
  overflow: hidden; /* only col scrolls individually */
}

.aip-col {
  display: flex;
  flex-direction: column;
  overflow: hidden;
  max-height: 420px;
}

.aip-col-label {
  font-size: 10px;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  padding: 8px 12px;
  color: var(--muted);
  border-bottom: 1px solid var(--border);
  flex-shrink: 0;
}

.aip-text {
  flex: 1;
  overflow-y: auto;
  padding: 10px 12px;
  font-size: 11px;
  line-height: 1.6;
  max-height: 360px;
  scrollbar-width: thin;
  scrollbar-color: var(--border) transparent;
}
.aip-text::-webkit-scrollbar {
  width: 4px;
}
.aip-text::-webkit-scrollbar-thumb {
  background: var(--border);
  border-radius: 4px;
}

.aip-text pre {
  white-space: pre-wrap;
  word-break: break-word;
  font-family: 'JetBrains Mono', 'Fira Code', monospace;
  font-size: 10.5px;
  margin: 0;
}

/* Human-readable field rendering */
.aip-text :deep(.aip-field) {
  padding: 8px 0;
  border-bottom: 1px solid var(--border);
  display: flex;
  flex-direction: column;
  gap: 4px;
}
.aip-text :deep(.aip-field:last-child) {
  border-bottom: none;
}
.aip-text :deep(.aip-field-key) {
  font-size: 10px;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: var(--muted);
}
.aip-text :deep(.aip-field-val) {
  font-size: 12px;
  color: inherit;
  line-height: 1.6;
}
.aip-text :deep(.aip-value) {
  font-size: 12px;
  line-height: 1.6;
  white-space: pre-wrap;
  word-break: break-word;
}
.aip-text :deep(.aip-empty) {
  opacity: 0.35;
  font-style: italic;
  font-size: 11px;
}
.aip-text :deep(.aip-array-item) {
  display: flex;
  gap: 8px;
  padding: 6px 0;
  border-bottom: 1px solid var(--border);
  align-items: flex-start;
}
.aip-text :deep(.aip-array-item:last-child) {
  border-bottom: none;
}
.aip-text :deep(.aip-item-num) {
  font-size: 10px;
  font-weight: 800;
  width: 18px;
  height: 18px;
  border-radius: 50%;
  background: rgba(79,70,229,0.15);
  color: var(--indigo);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  margin-top: 1px;
}
.aip-text :deep(.aip-item-body) {
  flex: 1;
  min-width: 0;
}
.aip-text :deep(.aip-field-nested) {
  padding: 4px 0 4px 8px;
  border-bottom: none;
  border-left: 2px solid var(--border);
  margin-left: 4px;
}

.aip-original {
  background: rgba(239,68,68,0.04);
  color: var(--muted);
  border-right: 1px solid var(--border);
}

.aip-new {
  background: rgba(79,70,229,0.05);
  color: var(--text);
}

.aip-divider-v {
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--bg2);
  border-left: 1px solid var(--border);
  border-right: 1px solid var(--border);
}

.aip-arrow {
  font-size: 16px;
  color: var(--indigo);
  font-weight: 700;
}

.aip-actions {
  display: flex;
  gap: 10px;
  padding: 12px 16px;
  border-top: 1px solid var(--border);
  background: var(--surface);
}

.aip-actions .btn {
  flex: 1;
  font-size: 13px;
  padding: 10px 16px;
  border-radius: 12px;
}

.aip-apply-btn {
  background: linear-gradient(135deg, var(--indigo), var(--purple));
  font-weight: 700;
  box-shadow: 0 4px 16px rgba(79,70,229,0.3);
}

.aip-apply-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 6px 20px rgba(79,70,229,0.45);
}

/* Hint in mini panel */
.ai-preview-hint {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 12px;
  color: var(--muted);
  padding: 10px 12px;
  background: rgba(79,70,229,0.06);
  border-radius: 10px;
  border: 1px dashed rgba(79,70,229,0.3);
}

/* Transition */
.ai-preview-enter-active, .ai-preview-leave-active {
  transition: all 0.35s cubic-bezier(0.16, 1, 0.3, 1);
}
.ai-preview-enter-from, .ai-preview-leave-to {
  opacity: 0;
  transform: translateY(-8px) scale(0.98);
}

/* Stack columns vertically when panel is too narrow (container query) */
@container (max-width: 400px) {
  .aip-body {
    grid-template-columns: 1fr;
    max-height: none;
    overflow: visible;
  }
  .aip-divider-v {
    height: 32px;
    border-top: 1px solid var(--border);
    border-left: none;
    border-right: none;
    justify-content: center;
    align-items: center;
    background: var(--bg2);
  }
  .aip-arrow {
    transform: rotate(90deg);
    font-size: 14px;
  }
  .aip-original {
    border-right: none;
    border-bottom: 1px solid var(--border);
  }
  .aip-col {
    max-height: 260px;
    overflow: hidden;
  }
  .aip-col-label {
    font-size: 9px;
    padding: 6px 10px;
  }
  .aip-text {
    max-height: 220px;
    font-size: 10.5px;
  }
}

/* Fallback for browsers without container query support */
@media (max-width: 460px) {
  .aip-body {
    grid-template-columns: 1fr;
  }
  .aip-divider-v {
    display: none;
  }
  .aip-original {
    border-right: none;
    border-bottom: 1px solid var(--border);
  }
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

/* Mobile Bottom Nav */
.builder-mobile-tabs {
  display: none; /* Hidden on all sizes — replaced by bottom nav on mobile */
}

.mobile-bottom-nav {
  display: none;
}

@media (max-width: 768px) {
  .mobile-bottom-nav {
    display: flex;
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    background: var(--sidebar-bg);
    border-top: 1px solid var(--border);
    padding: 8px 4px;
    padding-bottom: calc(8px + env(safe-area-inset-bottom));
    z-index: 500;
    gap: 4px;
    box-shadow: 0 -4px 20px rgba(0,0,0,0.15);
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
    bottom: 60px; /* account for bottom nav */
    width: 85vw !important;
    max-width: 320px;
    z-index: 600;
    margin: 0 !important;
    transform: none !important;
    box-shadow: 0 0 50px rgba(0,0,0,0.4);
    background: var(--surface) !important;
  }
  
  .block-palette {
    left: 0;
  }
  
  .props-panel {
    right: 0;
  }

  /* Overlay hidden state */
  .block-palette.collapsed {
    transform: translateX(-100%) !important;
  }
  .props-panel.collapsed {
    transform: translateX(100%) !important;
  }

  .builder-canvas {
    padding: 12px;
    padding-bottom: 80px; /* space for bottom nav */
  }

  .sections-grid,
  .sections-grid.tpl-business {
    grid-template-columns: 1fr !important;
    gap: 12px !important;
  }

  .section-wrapper.span-half,
  .sections-grid.tpl-business .section-wrapper.span-half,
  .sections-grid.tpl-business .section-wrapper.span-full {
    grid-column: 1 / -1 !important;
    width: 100% !important;
  }

  /* SectionBlock mobile header compacting */
  :deep(.section-header) {
    padding: 6px 10px !important;
  }
  :deep(.section-actions) {
    gap: 4px !important;
  }
  :deep(.icon-btn) {
    padding: 4px !important;
  }
  
  /* Reset span-full sidebar area on mobile */
  .sections-grid.tpl-business > .section-wrapper.span-full :deep(.pub-section) {
    margin-left: 0 !important;
    width: 100% !important;
    padding: 20px 12px !important;
  }
  .sections-grid.tpl-business > .section-wrapper.span-full :deep(.section-header-wrapper) {
    text-align: center !important;
    padding-left: 0 !important;
  }
  .sections-grid.tpl-business > .section-wrapper.span-full :deep(.pub-section-content) {
    padding-left: 0 !important;
  }

  /* Mobile Backdrop */
  .mobile-backdrop {
    position: fixed;
    inset: 0;
    bottom: 60px;
    background: rgba(0,0,0,0.5);
    z-index: 599;
    backdrop-filter: blur(2px);
  }
  .fade-backdrop-enter-active, .fade-backdrop-leave-active { transition: opacity 0.3s; }
  .fade-backdrop-enter-from, .fade-backdrop-leave-to { opacity: 0; }
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

/* AI Panel Styles */
.ai-panel {
  padding: 16px;
  border-top: 1px solid var(--border);
  margin-top: auto;
}
.ai-panel-title {
  font-size: 11px;
  font-weight: 700;
  color: var(--muted);
  text-transform: uppercase;
  letter-spacing: 0.5px;
  margin-bottom: 12px;
}

.ai-loading {
  text-align: center;
  padding: 20px 10px;
}
.ai-loading p {
  font-size: 13px;
  color: var(--muted);
  margin-top: 12px;
}
.ai-pulse {
  position: relative;
  width: 48px;
  height: 48px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: center;
}
.sparkle-icon {
  color: var(--indigo);
  z-index: 2;
}
.pulse-ring {
  position: absolute;
  width: 100%;
  height: 100%;
  border-radius: 50%;
  background: var(--indigo);
  opacity: 0.3;
  animation: ai-pulse-anim 1.5s infinite ease-out;
}

@keyframes ai-pulse-anim {
  0% { transform: scale(0.8); opacity: 0.5; }
  100% { transform: scale(2.2); opacity: 0; }
}

.ai-text-result {
  font-size: 13px;
  line-height: 1.5;
  color: var(--text);
  background: var(--bg2);
  padding: 12px;
  border-radius: 10px;
  white-space: pre-wrap;
}

/* Hide handles on mobile */
@media (max-width: 768px) {
  .resize-handle {
    display: none !important;
  }
}

/* Mobile Bottom Nav Buttons */
.mob-nav-btn {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 3px;
  padding: 6px 4px;
  background: transparent;
  border: none;
  color: var(--muted);
  font-size: 10px;
  font-weight: 600;
  cursor: pointer;
  border-radius: 8px;
  transition: all 0.2s;
  min-width: 0;
}
.mob-nav-btn.active {
  color: var(--indigo);
}
.mob-nav-btn.btn-primary-mobile {
  color: var(--indigo);
}
.mob-nav-btn.btn-danger-mobile {
  color: var(--danger);
}
.mob-nav-btn:disabled {
  opacity: 0.3;
}

/* ── Business Hero: ล้าง builder chrome ให้ hero เต็มความกว้าง ── */
.tpl-business.sections-grid > .section-wrapper.span-full :deep(.section-block) {
    border-radius: 0 !important;
    border: none !important;
    overflow: visible !important;
}

/* Hero section-block ไม่มี border/radius */
.tpl-business.sections-grid > .section-wrapper .type-hero :deep(.section-block),
.tpl-business.sections-grid > .section-wrapper .type-hero {
    border-radius: 0 !important;
    border: 1px solid var(--border) !important;
    overflow: hidden !important;
}

/* Hero preview เต็ม width ไม่มี padding */
.tpl-business.sections-grid > .section-wrapper .type-hero :deep(.section-preview) {
    border-radius: 0 !important;
    padding: 0 !important;
}

.tpl-business.sections-grid > .section-wrapper .type-hero :deep(.section-preview-inner) {
    padding: 0 !important;
    margin: 0 !important;
    width: 100% !important;
}

/* pub-hero-inner ต้องเต็ม width ใน builder ด้วย */
.tpl-business.sections-grid > .section-wrapper .type-hero :deep(.pub-hero-inner) {
    width: 100% !important;
    max-width: 100% !important;
    margin: 0 !important;
    padding: 0 !important;
    min-height: 380px;
}

/* hero-avatar-wrap ใน builder: fixed 30% */
.tpl-business.sections-grid > .section-wrapper .type-hero :deep(.hero-avatar-wrap) {
    width: 30% !important;
    min-width: 30% !important;
    flex-shrink: 0 !important;
    min-height: 380px !important;
}

/* hero-text ใน builder: 70% */
.tpl-business.sections-grid > .section-wrapper .type-hero :deep(.hero-text) {
    width: 70% !important;
    flex: 1 !important;
    min-height: 380px !important;
    padding: 48px 40px !important;
}

/* avatar image ย่อลงนิดนึงใน builder view */
.tpl-business.sections-grid > .section-wrapper .type-hero :deep(.hero-img),
.tpl-business.sections-grid > .section-wrapper .type-hero :deep(.hero-avatar-fallback) {
    width: 140px !important;
    height: 140px !important;
    font-size: 48px !important;
}

/* hero-name ย่อลงใน builder */
.tpl-business.sections-grid > .section-wrapper .type-hero :deep(.hero-name) {
    font-size: 28px !important;
}

</style>
