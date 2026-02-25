<template>
  <div class="popup-text-editor-wrapper">
    <button type="button" class="btn btn-secondary btn-sm" style="width: 100%; justify-content: flex-start; text-align: left;" @click.prevent="open">
      <Edit3 :size="14" /> {{ buttonText || `Edit ${title}` }}
    </button>

    <Teleport to="body">
      <div v-if="isOpen" class="raw-modal-overlay" @mousedown.self="close">
        <div class="raw-modal-content">
          <div class="raw-modal-header">
            <h2 class="layered-title" :data-text="title" style="font-size: 24px; margin: 0;">{{ title }}</h2>
            <button type="button" @click.prevent="close" style="background: transparent; border: none; cursor: pointer; color: var(--text);">
              <X :size="24" />
            </button>
          </div>
          <div class="raw-modal-body">
            <textarea
              v-model="tempValue"
              class="form-input"
              :rows="rows || 8"
              :placeholder="placeholder"
              style="resize: vertical; font-size: 16px; padding: 16px; width: 100%;"
            ></textarea>
          </div>
          <div class="raw-modal-footer">
            <button type="button" class="btn btn-secondary" @click.prevent="close">Cancel</button>
            <button type="button" class="btn btn-primary" @click.prevent="save">Save Changes</button>
          </div>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from "vue";
import { Edit3, X } from "lucide-vue-next";

const props = defineProps<{
  modelValue: string;
  title: string;
  placeholder?: string;
  buttonText?: string;
  rows?: number;
}>();

const emit = defineEmits<{
  "update:modelValue": [string];
}>();

const isOpen = ref(false);
const tempValue = ref(props.modelValue || "");

watch(
  () => props.modelValue,
  (newVal) => {
    if (!isOpen.value) {
      tempValue.value = newVal || "";
    }
  }
);

watch(isOpen, (newVal) => {
  if (newVal) {
    tempValue.value = props.modelValue || "";
    document.body.style.overflow = "hidden";
  } else {
    document.body.style.overflow = "";
  }
});

function open() {
  isOpen.value = true;
}

function close() {
  isOpen.value = false;
}

function save() {
  emit("update:modelValue", tempValue.value);
  close();
}
</script>

<style scoped>
.raw-modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: rgba(0, 0, 0, 0.75);
  backdrop-filter: blur(8px);
  z-index: 2147483647; /* Maximum visible z-index */
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
}

.raw-modal-content {
  background: var(--surface, #1e1e2f);
  border: 1px solid var(--border, #333);
  border-radius: 16px;
  width: 100%;
  max-width: 800px;
  display: flex;
  flex-direction: column;
  box-shadow: 0 20px 50px rgba(0,0,0,0.5);
}

.raw-modal-header {
  padding: 24px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid var(--border, #333);
}

.raw-modal-body {
  padding: 24px;
}

.raw-modal-footer {
  padding: 16px 24px;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  border-top: 1px solid var(--border, #333);
}
</style>
