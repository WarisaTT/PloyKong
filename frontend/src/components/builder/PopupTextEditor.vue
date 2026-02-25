<template>
  <div class="popup-text-editor-wrapper">
    <button type="button" class="btn btn-secondary btn-sm" style="width: 100%; justify-content: flex-start; text-align: left;" @click.prevent="open">
      <Edit3 :size="14" /> {{ buttonText || `Edit ${title}` }}
    </button>

    <Teleport to="body">
      <Transition name="fade">
        <div v-if="isOpen" class="modal-overlay" @mousedown.self="close">
          <div class="modal-content animate-slide-up">
            <div class="modal-header">
              <h2 class="layered-title" :data-text="title" style="font-size: 24px;">{{ title }}</h2>
              <button type="button" class="icon-btn" style="background: transparent; border: none; cursor: pointer; color: var(--text);" @click.prevent="close">
                <X :size="24" />
              </button>
            </div>
            <div class="modal-body">
              <textarea
                v-model="tempValue"
                class="form-input"
                :rows="rows || 8"
                :placeholder="placeholder"
                style="resize: vertical; font-size: 16px; padding: 16px;"
              ></textarea>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-secondary" @click.prevent="close">Cancel</button>
              <button type="button" class="btn btn-primary" @click.prevent="save">Save Changes</button>
            </div>
          </div>
        </div>
      </Transition>
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
.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.6);
  backdrop-filter: blur(8px);
  z-index: 999999;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
}

.modal-content {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 20px;
  width: 100%;
  max-width: 800px;
  max-height: 90vh;
  display: flex;
  flex-direction: column;
  box-shadow: 0 32px 64px rgba(0, 0, 0, 0.4);
}

.modal-header {
  padding: 24px 32px 16px;
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  border-bottom: 1px solid var(--border);
}

.modal-body {
  padding: 32px;
  overflow-y: auto;
  flex: 1;
}

.modal-footer {
  padding: 24px 32px;
  display: flex;
  justify-content: flex-end;
  gap: 16px;
  border-top: 1px solid var(--border);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
