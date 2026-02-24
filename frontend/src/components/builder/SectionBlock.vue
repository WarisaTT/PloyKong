<template>
  <div
    class="section-block"
    :class="[`type-${section.type}`, { selected: isSelected, hidden: !section.is_visible }]"
    @click="$emit('select')"
  >
    <!-- Section Header -->
    <div class="section-header">
      <div class="section-left">
        <span class="drag-handle" title="ลาก">⠿</span>
        <span class="section-icon">{{ icon }}</span>
        <span class="section-type-label">{{ label }}</span>
        <span v-if="!section.is_visible" class="visibility-badge">Hidden</span>
      </div>
      <div class="section-actions" @click.stop>
        <button class="icon-btn" title="Move Up" @click="$emit('move-up')">↑</button>
        <button class="icon-btn" title="Move Down" @click="$emit('move-down')">↓</button>
        <button
          class="icon-btn"
          :title="section.is_visible ? 'Hide' : 'Show'"
          @click="$emit('toggle-visibility')"
        >{{ section.is_visible ? '👁️' : '🙈' }}</button>
        <button class="icon-btn danger" title="Delete" @click="confirmDelete">🗑️</button>
      </div>
    </div>

    <!-- Section Preview -->
    <div class="section-preview">
      <HeroPreview v-if="section.type === 'hero'" :data="section.data" />
      <SkillsPreview v-else-if="section.type === 'skills'" :data="section.data" />
      <ProjectsPreview v-else-if="section.type === 'projects'" :data="section.data" />
      <ExperiencePreview v-else-if="section.type === 'experience'" :data="section.data" />
      <ContactPreview v-else-if="section.type === 'contact'" :data="section.data" />
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
import HeroPreview from './previews/HeroPreview.vue'
import SkillsPreview from './previews/SkillsPreview.vue'
import ProjectsPreview from './previews/ProjectsPreview.vue'
import ExperiencePreview from './previews/ExperiencePreview.vue'
import ContactPreview from './previews/ContactPreview.vue'
import GenericPreview from './previews/GenericPreview.vue'

const props = defineProps<{
  section: Section
  isSelected: boolean
}>()

const emit = defineEmits<{
  select: []
  delete: []
  'toggle-visibility': []
  'move-up': []
  'move-down': []
}>()

const blockMeta = computed(() => BLOCK_TYPES.find((b) => b.type === props.section.type))
const icon = computed(() => blockMeta.value?.icon || '📦')
const label = computed(() => blockMeta.value?.label || props.section.type)

function confirmDelete() {
  if (confirm('ลบ section นี้หรือไม่?')) {
    emit('delete')
  }
}
</script>

<style scoped>
.section-block {
  background: rgba(15, 20, 40, 0.6);
  border: 1px solid var(--border);
  border-radius: 12px;
  overflow: hidden;
  cursor: pointer;
  transition: all 0.2s;
}
.section-block:hover { border-color: rgba(79, 70, 229, 0.4); }
.section-block.selected {
  border-color: var(--purple);
  box-shadow: 0 0 0 2px rgba(168, 85, 247, 0.2);
}
.section-block.hidden { opacity: 0.4; }

.section-header {
  display: flex; align-items: center; justify-content: space-between;
  padding: 10px 14px;
  background: rgba(255, 255, 255, 0.03);
  border-bottom: 1px solid var(--border);
}
.section-left { display: flex; align-items: center; gap: 10px; }
.drag-handle { color: var(--muted); cursor: grab; font-size: 16px; }
.section-icon { font-size: 16px; }
.section-type-label { font-size: 13px; font-weight: 600; }
.visibility-badge {
  font-size: 10px; padding: 2px 8px; border-radius: 100px;
  background: rgba(148,163,184,0.1); color: var(--muted); border: 1px solid var(--border);
}

.section-actions { display: flex; gap: 4px; }
.icon-btn {
  width: 28px; height: 28px; border-radius: 6px; border: none;
  background: rgba(255,255,255,0.05); color: var(--muted);
  font-size: 12px; cursor: pointer; transition: all 0.15s;
  display: flex; align-items: center; justify-content: center;
}
.icon-btn:hover { background: rgba(255,255,255,0.1); color: #fff; }
.icon-btn.danger:hover { background: rgba(239,68,68,0.15); color: var(--danger); }

.section-preview { padding: 14px; min-height: 60px; }
.selected-indicator {
  padding: 6px 14px;
  background: rgba(168, 85, 247, 0.08);
  border-top: 1px solid rgba(168, 85, 247, 0.2);
  font-size: 11px; color: var(--neon-purple); text-align: center;
}
</style>
