<template>
  <div
    class="section-block"
    :class="[
      `type-${section.type}`,
      { selected: isSelected, hidden: !section.is_visible },
    ]"
    @click="$emit('select')"
    :style="themeVars"
  >
    <!-- Section Header -->
    <div class="section-header">
      <div class="section-left">
        <span class="drag-handle" title="ลาก">⠿</span>
        <component :is="icon" class="section-icon" :size="14" />
        <span class="section-type-label">{{ label }}</span>
        <span v-if="!section.is_visible" class="visibility-badge">Hidden</span>
      </div>
      <div class="section-actions" @click.stop>
        <!-- Move Up -->
        <button class="icon-btn" title="Move Up" @click="$emit('move-up')">
          <ArrowUp :size="12" />
        </button>

        <!-- Toggle List/Grid Column Span -->
        <button
          v-if="section.type !== 'hero'"
          class="icon-btn"
          :title="section.column_span === 'half' ? 'Make Full Width' : 'Make Half Width'"
          @click="$emit('toggle-column-span')"
        >
          <Columns v-if="section.column_span === 'half'" :size="12" />
          <Square v-else :size="12" />
        </button>

        <!-- Move Down -->
        <button class="icon-btn" title="Move Down" @click="$emit('move-down')">
          <ArrowDown :size="12" />
        </button>

        <!-- Toggle Visibility -->
        <button
          class="icon-btn"
          :title="section.is_visible ? 'Hide' : 'Show'"
          @click="$emit('toggle-visibility')"
        >
          <Eye v-if="section.is_visible" :size="12" />
          <EyeOff v-else :size="12" />
        </button>

        <!-- Duplicate -->
        <button class="icon-btn" title="Duplicate" @click="$emit('duplicate')">
          <Copy :size="12" />
        </button>

        <!-- Delete -->
        <button class="icon-btn danger" title="Delete" @click="confirmDelete">
          <Trash2 :size="12" />
        </button>
      </div>
    </div>

    <!-- Section Preview -->
    <div
      class="section-preview"
      :class="[themeClass, templateClass, { 'is-half-split': isHalfSplit }]"
      :style="section.data?.section_bg_color ? `pointer-events: none; background: ${section.data.section_bg_color} !important; --item-bg: ${section.data.section_bg_color};` : 'pointer-events: none; background: color-mix(in srgb, var(--theme-bg, var(--bg)) 60%, transparent);'"
    >
      <div 
        class="section-preview-inner"
        :class="{ 'pub-container': section.type !== 'hero' && !isHalfSplit }"
      >
        <PublicHero v-if="section.type === 'hero'" :data="section.data" />
        <PublicSkills
          v-else-if="section.type === 'skills'"
          :data="section.data"
        />
        <PublicProjects
          v-else-if="section.type === 'projects'"
          :data="section.data"
        />
        <PublicExperience
          v-else-if="section.type === 'experience'"
          :data="section.data"
        />
        <PublicContact
          v-else-if="section.type === 'contact'"
          :data="section.data"
        />
        <PublicCustomText
          v-else-if="section.type === 'custom_text'"
          :data="section.data"
        />
        <PublicEducation
          v-else-if="section.type === 'education'"
          :data="section.data"
        />
        <PublicStats
          v-else-if="section.type === 'stats'"
          :data="section.data"
        />
        <PublicSocial
          v-else-if="section.type === 'social'"
          :data="section.data"
        />
        <PublicCertificates
          v-else-if="section.type === 'certificates'"
          :data="section.data"
        />
        <GenericPreview v-else :section="section" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import type { Section } from "@/types";
import { BLOCK_TYPES } from "@/types";
import {
  Package,
  Eye,
  EyeOff,
  ArrowUp,
  ArrowDown,
  Trash2,
  Rows,
  Wrench,
  Trophy,
  Briefcase,
  Mail,
  User,
  GraduationCap,
  Sparkles,
  Link,
  Award,
  Columns,
  Square,
  Copy,
} from "lucide-vue-next";
import Swal from "sweetalert2";
import PublicHero from "../portfolio/PublicHero.vue";
import PublicSkills from "../portfolio/PublicSkills.vue";
import PublicProjects from "../portfolio/PublicProjects.vue";
import PublicExperience from "../portfolio/PublicExperience.vue";
import PublicContact from "../portfolio/PublicContact.vue";
import PublicCustomText from "../portfolio/PublicCustomText.vue";
import PublicEducation from "../portfolio/PublicEducation.vue";
import PublicStats from "../portfolio/PublicStats.vue";
import PublicSocial from "../portfolio/PublicSocial.vue";
import PublicCertificates from "../portfolio/PublicCertificates.vue";
import GenericPreview from "./previews/GenericPreview.vue";

const props = defineProps<{
  section: Section;
  isSelected: boolean;
  isHalfSplit?: boolean;
  themeClass?: string;
  templateClass?: string;
  themeVars?: Record<string, string>;
}>();

const emit = defineEmits([
  "select",
  "delete",
  "toggle-visibility",
  "toggle-column-span",
  "move-up",
  "move-down",
  "duplicate",
]);

const blockMeta = computed(() =>
  BLOCK_TYPES.find((b) => b.type === props.section.type),
);
const icon = computed(() => blockMeta.value?.icon || Package);
const label = computed(() => blockMeta.value?.label || props.section.type);

async function confirmDelete() {
  const { isConfirmed } = await Swal.fire({
    title: "Delete this section?",
    icon: "warning",
    showCancelButton: true,
    confirmButtonColor: "#ef4444",
    cancelButtonColor: "#94a3b8",
    confirmButtonText: "Delete",
  });
  if (isConfirmed) {
    emit("delete");
  }
}
</script>

<style scoped>
.section-block {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  background: color-mix(in srgb, var(--theme-bg, var(--bg)) 60%, transparent);
  border: 1.5px solid var(--section-border);
  border-radius: 20px;
  overflow: hidden;
  position: relative;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  cursor: pointer;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.03);
}

.section-block:hover {
  border-color: var(--primary);
  transform: translateY(-4px);
  box-shadow: 0 12px 24px var(--primary-glow);
}

.section-block.selected {
  border-color: var(--primary);
  border-width: 2px;
  box-shadow: 0 0 0 4px var(--primary-glow), 0 12px 24px rgba(0, 0, 0, 0.1);
}

:global(.theme-dark) .section-block {
  background: color-mix(in srgb, var(--theme-bg, var(--bg)) 60%, transparent);
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
  padding: 10px 16px;
  background: var(--sidebar-bg);
  border: 1px solid var(--section-border);
  z-index: 10;
  opacity: 1;
  transition: opacity 0.3s ease;
}

.section-header.section-left {
  color: var(--text);
}


.section-left {
  display: flex;
  align-items: center;
  gap: 8px;
}

.drag-handle {
  color: var(--muted);
  cursor: grab;
  font-size: 14px;
  padding: 2px;
}

.section-icon {
  color: var(--primary);
  opacity: 0.8;
}

.section-type-label {
  font-size: 11px;
  font-weight: 700;
  color: var(--text);
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.visibility-badge {
  font-size: 9px;
  padding: 1px 6px;
  border-radius: 100px;
  background: rgba(var(--primary-glow), 0.1);
  color: var(--primary);
  border: 1px solid rgba(var(--primary-glow), 0.2);
}

.section-actions {
  display: flex;
  gap: 4px;
}

.icon-btn {
  width: 24px;
  height: 24px;
  border-radius: 6px;
  border: 1px solid transparent;
  background: transparent;
  color: var(--muted);
  cursor: pointer;
  transition: all 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
}

.icon-btn:hover {
  background: rgba(var(--primary-glow), 0.1);
  color: var(--primary);
  border-color: rgba(var(--primary-glow), 0.2);
}

.icon-btn.danger:hover {
  background: rgba(239, 68, 68, 0.1);
  color: #ef4444;
  border-color: rgba(239, 68, 68, 0.2);
}

.section-preview {
  flex: 1;
  background: var(--bg);
  border-radius: 0 0 16px 16px;
  color: var(--text);
  font-family: var(--font-body);
  padding: 0;
  min-height: 60px;
  display: flex;
  flex-direction: column;
  position: relative;
}

.tpl-business .pub-section.is-row-start{
  background: var(--primary) !important;
  border-radius: 20px !important;
}

/* Business: no longer transparent */

.tpl-firstjobber .section-preview > div {
  display: flex;
  flex-direction: column;
  color: var(--text);
}

.tpl-firstjobber .section-preview > div > section {
  flex: 1;
  color: var(--text);
}

.pub-container {
  margin: 0 auto;
  padding: 0 24px;
}

/* Business template: match others */

.tpl-business.section-preview.inner-section {
  background: var(--bg) !important;
  color: var(--text);
}
.tpl-business .pub-container {
  padding: 5px !important;
  margin: 5px !important;
}

/* Force all inner content to fill width */
.tpl-business :deep(.pub-section),
.tpl-business :deep(.pub-section-content),
.tpl-business :deep(.experience-list),
.tpl-business :deep(.contact-grid),
.tpl-business :deep(.skills-groups),
.tpl-business :deep(.skills-grid),
.tpl-business :deep(.proj-grid),
.tpl-business :deep(.education-list),
.tpl-business :deep(.custom-content) {
  max-width: 100% !important;
  width: 100% !important;
}

</style>
