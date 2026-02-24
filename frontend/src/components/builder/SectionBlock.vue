<template>
  <div class="section-block" :class="[`type-${section.type}`, { selected: isSelected, hidden: !section.is_visible }]"
    @click="$emit('select')">
    <!-- Section Header -->
    <div class="section-header">
      <div class="section-left">
        <span class="drag-handle" title="ลาก">⠿</span>
        <component :is="icon" class="section-icon" :size="16" />
        <span class="section-type-label">{{ label }}</span>
        <span v-if="!section.is_visible" class="visibility-badge">Hidden</span>
      </div>
      <div class="section-actions" @click.stop>
        <!-- Move Up -->
        <button class="icon-btn" title="Move Up" @click="$emit('move-up')">
          <ArrowUp :size="14" />
        </button>

        <!-- Move Down -->
        <button class="icon-btn" title="Move Down" @click="$emit('move-down')">
          <ArrowDown :size="14" />
        </button>

        <!-- Toggle Visibility -->
        <button class="icon-btn" :title="section.is_visible ? 'Hide' : 'Show'" @click="$emit('toggle-visibility')">
          <Eye v-if="section.is_visible" :size="14" />
          <EyeOff v-else :size="14" />
        </button>

        <!-- Delete -->
        <button class="icon-btn danger" title="Delete" @click="confirmDelete">
          <Trash2 :size="14" />
        </button>
      </div>
    </div>

    <!-- Section Preview -->
    <div class="section-preview" :class="themeClass" :style="[themeVars || {}, { pointerEvents: 'none' }]">
      <PublicHero v-if="section.type === 'hero'" :data="section.data" />
      <PublicSkills v-else-if="section.type === 'skills'" :data="section.data" />
      <PublicProjects v-else-if="section.type === 'projects'" :data="section.data" />
      <PublicExperience v-else-if="section.type === 'experience'" :data="section.data" />
      <PublicContact v-else-if="section.type === 'contact'" :data="section.data" />
      <GenericPreview v-else :section="section" />
    </div>

    <!-- Selected Indicator -->
    <div v-if="isSelected" class="selected-indicator">
      <span>✏️ Click Properties panel to edit</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { Section } from '@/types'
import { BLOCK_TYPES } from '@/types'
import { Package, Eye, EyeOff, ArrowUp, ArrowDown, Trash2 } from 'lucide-vue-next'
import Swal from 'sweetalert2'
import PublicHero from '../portfolio/PublicHero.vue'
import PublicSkills from '../portfolio/PublicSkills.vue'
import PublicProjects from '../portfolio/PublicProjects.vue'
import PublicExperience from '../portfolio/PublicExperience.vue'
import PublicContact from '../portfolio/PublicContact.vue'
import GenericPreview from './previews/GenericPreview.vue'

const props = defineProps<{
  section: Section
  isSelected: boolean
  themeClass?: string
  themeVars?: Record<string, string>
}>()

const emit = defineEmits<{
  select: []
  delete: []
  'toggle-visibility': []
  'move-up': []
  'move-down': []
}>()

const blockMeta = computed(() => BLOCK_TYPES.find((b) => b.type === props.section.type))
const icon = computed(() => blockMeta.value?.icon || Package)
const label = computed(() => blockMeta.value?.label || props.section.type)

async function confirmDelete() {
  const { isConfirmed } = await Swal.fire({
    title: 'Delete this section?',
    icon: 'warning',
    showCancelButton: true,
    confirmButtonColor: '#ef4444',
    cancelButtonColor: '#94a3b8',
    confirmButtonText: 'Delete'
  })
  if (isConfirmed) {
    emit('delete')
  }
}
</script>

<style scoped>
.section-block {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 12px;
  overflow: hidden;
  cursor: pointer;
  transition: all 0.2s;
}

.section-block:hover {
  border-color: rgba(79, 70, 229, 0.4);
}

.section-block.selected {
  border-color: var(--purple);
  box-shadow: 0 0 0 2px rgba(168, 85, 247, 0.2);
}

.section-block.hidden {
  opacity: 0.5;
  filter: grayscale(1);
  border: 1px dashed var(--muted);
}

.section-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px 14px;
  background: var(--bg2);
  border-bottom: 1px solid var(--border);
}

.section-left {
  display: flex;
  align-items: center;
  gap: 10px;
}

.drag-handle {
  color: var(--muted);
  cursor: grab;
  font-size: 16px;
}

.section-icon {
  font-size: 16px;
}

.section-type-label {
  font-size: 13px;
  font-weight: 600;
}

.visibility-badge {
  font-size: 10px;
  padding: 2px 8px;
  border-radius: 100px;
  background: rgba(148, 163, 184, 0.1);
  color: var(--muted);
  border: 1px solid var(--border);
}

.section-actions {
  display: flex;
  gap: 4px;
}

.icon-btn {
  width: 28px;
  height: 28px;
  border-radius: 6px;
  border: none;
  background: var(--surface);
  color: var(--muted);
  font-size: 12px;
  cursor: pointer;
  transition: all 0.15s;
  display: flex;
  align-items: center;
  justify-content: center;
}

.icon-btn:hover {
  background: var(--bg);
  color: var(--text);
}

.icon-btn.danger:hover {
  background: rgba(239, 68, 68, 0.15);
  color: var(--danger);
}

.section-preview {
  background: var(--bg);
  color: var(--text);
  font-family: var(--font-body);
  padding: 14px;
  min-height: 60px;
}

.selected-indicator {
  padding: 6px 14px;
  background: rgba(168, 85, 247, 0.08);
  /* fine for both themes as an accent */
  border-top: 1px solid rgba(168, 85, 247, 0.2);
  font-size: 11px;
  color: var(--neon-purple);
  text-align: center;
}
</style>
