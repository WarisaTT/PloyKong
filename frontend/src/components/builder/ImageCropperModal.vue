<template>
  <Transition name="fade">
    <div v-if="show" class="cropper-modal-backdrop" @click.self="cancel">
      <div class="cropper-modal-card">
        <div class="cropper-modal-header">
          <h3 class="cropper-modal-title">Crop Image</h3>
          <button type="button" class="close-btn" @click="cancel">
            <X :size="20" />
          </button>
        </div>
        
        <div class="cropper-modal-body">
          <Cropper
            ref="cropperRef"
            class="cropper-instance"
            :src="imageSrc"
            :stencil-props="{
              aspectRatio: aspectRatio,
              grid: true
            }"
            image-restriction="stencil"
          />
        </div>
        
        <div class="cropper-modal-footer">
          <button type="button" class="btn btn-secondary" @click="cancel">
            Cancel
          </button>
          <button type="button" class="btn btn-primary" @click="confirm">
            <Check :size="16" /> Done
          </button>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { Cropper } from 'vue-advanced-cropper';
import 'vue-advanced-cropper/dist/style.css';
import { X, Check } from 'lucide-vue-next';

const props = defineProps<{
  show: boolean;
  imageSrc: string;
  aspectRatio?: number;
}>();

const emit = defineEmits<{
  (e: 'confirm', blob: Blob): void;
  (e: 'cancel'): void;
}>();

const cropperRef = ref<any>(null);

function confirm() {
  if (cropperRef.value) {
    const { canvas } = cropperRef.value.getResult();
    if (canvas) {
      canvas.toBlob((blob: Blob) => {
        if (blob) {
          emit('confirm', blob);
        }
      }, 'image/jpeg', 0.9);
    }
  }
}

function cancel() {
  emit('cancel');
}
</script>

<style scoped>
.cropper-modal-backdrop {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.85);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 10000;
  backdrop-filter: blur(5px);
}

.cropper-modal-card {
  background: var(--surface);
  width: 90%;
  max-width: 600px;
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5);
  border: 1px solid var(--border);
}

.cropper-modal-header {
  padding: 16px 20px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-bottom: 1px solid var(--border);
}

.cropper-modal-title {
  margin: 0;
  font-size: 18px;
  font-weight: 700;
  color: var(--text);
}

.close-btn {
  background: transparent;
  border: none;
  color: var(--muted);
  cursor: pointer;
  padding: 4px;
  border-radius: 50%;
  display: flex;
  transition: all 0.2s;
}

.close-btn:hover {
  background: rgba(255, 255, 255, 0.1);
  color: var(--text);
}

.cropper-modal-body {
  padding: 0;
  height: 400px;
  background: #000;
  display: flex;
  align-items: center;
  justify-content: center;
}

.cropper-instance {
  width: 100%;
  height: 100%;
}

.cropper-modal-footer {
  padding: 16px 20px;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 12px;
  border-top: 1px solid var(--border);
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style>
