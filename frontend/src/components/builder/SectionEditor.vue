<template>
  <div class="section-editor">
    <!-- Hero Editor -->
    <template v-if="section.type === 'hero'">
      <div class="editor-field">
        <label class="form-label">ชื่อ</label>
        <input
          v-model="form.name"
          class="form-input"
          placeholder="Your Name"
          @input="emit('update', form)"
        />
      </div>
      <div class="editor-field">
        <label class="form-label">ตำแหน่ง / Role</label>
        <input
          v-model="form.role"
          class="form-input"
          placeholder="Full-Stack Developer"
          @input="emit('update', form)"
        />
      </div>
      <div class="editor-field">
        <label class="form-label">Tagline</label>
        <PopupTextEditor
          v-model="form.tagline"
          title="Tagline"
          placeholder="Building cool things..."
          :rows="4"
          @update:modelValue="emit('update', form)"
        />
      </div>
      <div class="editor-field">
        <label class="form-label">Profile Picture</label>
        <div style="display: flex; gap: 8px">
          <input
            v-model="form.avatar_url"
            class="form-input"
            style="flex: 1"
            placeholder="https://..."
            @input="emit('update', form)"
          />
          <input
            type="file"
            ref="avatarInputRef"
            accept="image/*"
            style="display: none"
            @change="uploadAvatar"
          />
          <button
            type="button"
            class="btn btn-secondary btn-sm"
            @click="avatarInputRef?.click()"
            :disabled="uploadingAvatar"
          >
            <template v-if="uploadingAvatar"
              ><Loader2 :size="14" class="spin" /> Uploading...</template
            >
            <template v-else>Upload</template>
          </button>
        </div>
      </div>
      <div class="editor-field toggle-row">
        <label class="form-label" style="margin: 0">Hire Me Button</label>
        <button
          class="toggle"
          :class="{ on: form.show_hire_me }"
          @click="
            form.show_hire_me = !form.show_hire_me;
            emit('update', form);
          "
        >
          {{ form.show_hire_me ? "ON" : "OFF" }}
        </button>
      </div>
      <button class="btn btn-secondary btn-sm ai-btn" @click="improveTagline">
        <Sparkles :size="14" /> AI Improve Tagline
      </button>
    </template>

    <!-- Skills Editor -->
    <template v-else-if="section.type === 'skills'">
      <div class="skills-editor">
        <div
          v-for="(skill, i) in form.items as any[]"
          :key="i"
          class="skill-row"
        >
          <input
            v-model="skill.name"
            class="form-input"
            style="flex: 1"
            placeholder="Skill name"
            @input="emit('update', form)"
          />
          <input
            v-model.number="skill.level"
            type="range"
            min="0"
            max="100"
            style="flex: 1"
            @input="emit('update', form)"
          />
          <span class="skill-pct">{{ skill.level }}%</span>
          <button class="icon-btn danger" @click="removeSkill(i)">
            <X :size="16" />
          </button>
        </div>
        <button
          class="btn btn-secondary btn-sm btn-icon-text"
          style="width: 100%; margin-top: 8px"
          @click="addSkill"
        >
          <Plus :size="16" /> Add Skill
        </button>
      </div>
    </template>

    <!-- Experience Editor -->
    <template v-else-if="section.type === 'experience'">
      <div
        v-for="(item, i) in form.items as any[]"
        :key="i"
        class="exp-editor-item"
      >
        <div class="exp-editor-header">
          <span style="font-size: 12px; font-weight: 700"
            >งาน #{{ i + 1 }}</span
          >
          <button
            class="icon-btn danger btn-icon-text"
            style="margin-left: auto; width: auto; padding: 0 8px"
            @click="removeExp(i)"
          >
            <Trash2 :size="14" /> Delete
          </button>
        </div>
        <input
          v-model="item.company"
          class="form-input"
          placeholder="Company"
          @input="emit('update', form)"
        />
        <input
          v-model="item.position"
          class="form-input"
          placeholder="Position"
          @input="emit('update', form)"
        />
        <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 6px">
          <input
            v-model="item.start_date"
            class="form-input"
            placeholder="2022-01"
            @input="emit('update', form)"
          />
          <input
            v-model="item.end_date"
            class="form-input"
            placeholder="2024-01"
            :disabled="item.is_current"
            @input="emit('update', form)"
          />
        </div>
        <PopupTextEditor
          v-model="item.description"
          title="Experience Description"
          placeholder="Describe your role..."
          :rows="6"
          @update:modelValue="emit('update', form)"
        />
      </div>
      <button
        class="btn btn-secondary btn-sm btn-icon-text"
        style="width: 100%; margin-top: 8px"
        @click="addExp"
      >
        <Plus :size="16" /> Add Experience
      </button>
    </template>

    <!-- Projects Editor -->
    <template v-else-if="section.type === 'projects'">
      <div
        v-for="(proj, i) in form.items as any[]"
        :key="i"
        class="proj-editor-item"
      >
        <div class="exp-editor-header">
          <span style="font-size: 12px; font-weight: 700"
            >Project #{{ i + 1 }}</span
          >
          <button
            class="icon-btn danger btn-icon-text"
            style="margin-left: auto; width: auto; padding: 0 8px"
            @click="removeProj(i)"
          >
            <Trash2 :size="14" /> Delete
          </button>
        </div>
        <input
          v-model="proj.title"
          class="form-input"
          placeholder="Project Title"
          @input="emit('update', form)"
        />
        <PopupTextEditor
          v-model="proj.description"
          title="Project Description"
          placeholder="Description"
          :rows="6"
          @update:modelValue="emit('update', form)"
        />
        <input
          v-model="proj.live_url"
          class="form-input"
          placeholder="Live URL (optional)"
          @input="emit('update', form)"
        />
        <input
          v-model="proj.github_url"
          class="form-input"
          placeholder="GitHub URL (optional)"
          @input="emit('update', form)"
        />
        <input
          :value="(proj.tags || []).join(', ')"
          class="form-input"
          placeholder="Tags: Go, Vue.js, Docker"
          @input="
            proj.tags = ($event.target as HTMLInputElement).value
              .split(',')
              .map((s) => s.trim());
            emit('update', form);
          "
        />
      </div>
      <button
        class="btn btn-secondary btn-sm btn-icon-text"
        style="width: 100%; margin-top: 8px"
        @click="addProj"
      >
        <Plus :size="16" /> Add Project
      </button>
    </template>

    <!-- Contact Editor -->
    <template v-else-if="section.type === 'contact'">
      <div class="editor-field">
        <label class="form-label">Email</label>
        <input
          v-model="form.email"
          class="form-input"
          placeholder="you@example.com"
          @input="emit('update', form)"
        />
      </div>
      <div class="editor-field">
        <label class="form-label">LinkedIn</label>
        <input
          v-model="form.linkedin"
          class="form-input"
          placeholder="linkedin.com/in/..."
          @input="emit('update', form)"
        />
      </div>
      <div class="editor-field">
        <label class="form-label">GitHub</label>
        <input
          v-model="form.github"
          class="form-input"
          placeholder="github.com/..."
          @input="emit('update', form)"
        />
      </div>
      <div class="editor-field" style="margin-bottom: 24px">
        <label class="form-label">Location</label>
        <input
          v-model="form.location"
          class="form-input"
          placeholder="Bangkok, Thailand"
          @input="emit('update', form)"
        />
      </div>

      <!-- Custom Contacts List -->
      <div
        v-for="(contact, i) in (form.custom_items || []) as any[]"
        :key="contact.id"
        class="proj-editor-item"
      >
        <div class="exp-editor-header">
          <span style="font-size: 12px; font-weight: 700"
            >Custom #{{ i + 1 }}</span
          >
          <button
            class="icon-btn danger btn-icon-text"
            style="margin-left: auto; width: auto; padding: 0 8px"
            @click="removeCustomContact(i)"
          >
            <Trash2 :size="14" />
          </button>
        </div>

        <div class="editor-field">
          <label class="form-label">Platform / Icon</label>
          <select
            v-model="contact.platform"
            class="form-input form-select"
            @change="emit('update', form)"
          >
            <option value="facebook">Facebook</option>
            <option value="twitter">X / Twitter</option>
            <option value="instagram">Instagram</option>
            <option value="youtube">YouTube</option>
            <option value="message-circle">Line / Messaging</option>
            <option value="phone">Phone Call</option>
            <option value="link">Website / Link</option>
          </select>
        </div>

        <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 6px">
          <div class="editor-field">
            <label class="form-label">Label</label>
            <input
              v-model="contact.label"
              class="form-input"
              placeholder="Label"
              @input="emit('update', form)"
            />
          </div>
          <div class="editor-field">
            <label class="form-label">Value</label>
            <input
              v-model="contact.value"
              class="form-input"
              placeholder="Value"
              @input="emit('update', form)"
            />
          </div>
        </div>

        <div class="editor-field" style="margin-top: 6px">
          <label class="form-label">URL / Link Action</label>
          <input
            v-model="contact.link"
            class="form-input"
            placeholder="https://... or tel:0812345678"
            @input="emit('update', form)"
          />
        </div>
      </div>

      <button
        class="btn btn-secondary btn-sm btn-icon-text"
        style="width: 100%; margin-top: 8px"
        @click="addCustomContact"
      >
        <Plus :size="16" /> Add Custom Contact
      </button>
    </template>

    <!-- Custom Text Editor -->
    <template v-else-if="section.type === 'custom_text'">
      <div class="editor-field">
        <label class="form-label">Content</label>
        <PopupTextEditor
          v-model="form.content"
          title="Custom Text Content"
          placeholder="เขียนอะไรก็ได้..."
          :rows="14"
          @update:modelValue="emit('update', form)"
        />
      </div>
    </template>

    <!-- Fallback -->
    <template v-else>
      <p
        style="
          font-size: 12px;
          color: var(--muted);
          text-align: center;
          padding: 12px 0;
        "
      >
        Editor สำหรับ "{{ section.type }}" อยู่ระหว่างพัฒนา
      </p>
    </template>
  </div>
</template>

<script setup lang="ts">
import { reactive, watch, ref } from "vue";
import type { Section } from "@/types";
import { aiAPI, uploadAPI } from "@/api";
import PopupTextEditor from "./PopupTextEditor.vue";
import { Sparkles, X, Plus, Trash2, Loader2 } from "lucide-vue-next";
import Swal from "sweetalert2";

const props = defineProps<{ section: Section }>();
const emit = defineEmits<{ update: [data: any] }>();

// Deep clone section data into reactive form
const form = reactive(JSON.parse(JSON.stringify(props.section.data || {})));

watch(
  () => props.section.id,
  () => {
    Object.assign(form, JSON.parse(JSON.stringify(props.section.data || {})));
  },
);

// Skills helpers
function addSkill() {
  if (!form.items) form.items = [];
  form.items.push({ name: "", level: 70, category: "" });
  emit("update", form);
}
function removeSkill(i: number) {
  form.items.splice(i, 1);
  emit("update", form);
}

// Upload Avatar
const avatarInputRef = ref<HTMLInputElement | null>(null);
const uploadingAvatar = ref(false);

async function uploadAvatar(e: Event) {
  const file = (e.target as HTMLInputElement).files?.[0];
  if (!file) return;

  uploadingAvatar.value = true;
  try {
    const { data } = await uploadAPI.image(file);
    // Create base URL to prepend to relative upload paths, similar to how API requests are made.
    // If VITE_API_URL is like 'http://localhost:8082/api/v1', we want 'http://localhost:8082'.
    const apiBase = import.meta.env.VITE_API_URL || "/api/v1";
    let rootUrl = apiBase.replace(/\/api\/v1\/?$/, "");
    form.avatar_url = rootUrl + data.data.url;
    emit("update", form);
  } catch (error) {
    Swal.fire("Error", "Failed to upload image", "error");
  } finally {
    uploadingAvatar.value = false;
    if (avatarInputRef.value) avatarInputRef.value.value = "";
  }
}

// Experience helpers
function addExp() {
  if (!form.items) form.items = [];
  form.items.push({
    company: "",
    position: "",
    start_date: "",
    end_date: "",
    is_current: false,
    description: "",
  });
  emit("update", form);
}
function removeExp(i: number) {
  form.items.splice(i, 1);
  emit("update", form);
}

// Project helpers
function addProj() {
  if (!form.items) form.items = [];
  form.items.push({
    title: "",
    description: "",
    live_url: "",
    github_url: "",
    tags: [],
  });
  emit("update", form);
}
function removeProj(i: number) {
  form.items.splice(i, 1);
  emit("update", form);
}

// Contact helpers
function addCustomContact() {
  if (!form.custom_items) form.custom_items = [];
  form.custom_items.push({
    id: Date.now().toString(),
    platform: "link",
    label: "",
    value: "",
    link: "",
  });
  emit("update", form);
}
function removeCustomContact(i: number) {
  form.custom_items.splice(i, 1);
  emit("update", form);
}

// AI improve tagline
async function improveTagline() {
  if (!form.tagline) return;
  try {
    const { data } = await aiAPI.improveText(
      form.tagline,
      "portfolio tagline for " + form.role,
    );
    form.tagline = data.data.improved_text;
    emit("update", form);
  } catch {}
}
</script>

<style scoped>
.section-editor {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
.editor-field {
  display: flex;
  flex-direction: column;
}
.toggle-row {
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}
.form-label {
  color: var(--text);
  font-size: 12px;
  font-weight: 700;
}

.toggle {
  padding: 5px 14px;
  border-radius: 100px;
  border: 1px solid var(--border);
  background: var(--surface);
  color: var(--muted);
  font-size: 12px;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.15s;
}
.toggle.on {
  background: rgba(79, 70, 229, 0.15);
  border-color: var(--indigo);
  color: var(--indigo);
}

.skills-editor {
  display: flex;
  flex-direction: column;
  gap: 8px;
}
.skill-row {
  display: flex;
  align-items: center;
  gap: 6px;
}
.skill-pct {
  font-size: 11px;
  color: var(--muted);
  width: 30px;
  text-align: right;
  flex-shrink: 0;
}

.exp-editor-item,
.proj-editor-item {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 10px;
  padding: 12px;
  display: flex;
  flex-direction: column;
  gap: 7px;
}
.exp-editor-header {
  display: flex;
  align-items: center;
  margin-bottom: 4px;
}

.ai-btn {
  margin-top: 4px;
  background: rgba(168, 85, 247, 0.1);
  border-color: rgba(168, 85, 247, 0.3);
  color: var(--neon-purple);
  display: inline-flex;
  align-items: center;
  gap: 6px;
}
.btn-icon-text {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  justify-content: center;
}

.icon-btn {
  width: 26px;
  height: 26px;
  border-radius: 6px;
  border: none;
  background: var(--surface);
  color: var(--muted);
  font-size: 14px;
  cursor: pointer;
  transition: all 0.15s;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}
.icon-btn:hover {
  background: var(--bg);
  color: var(--text);
}
.icon-btn.danger:hover {
  background: rgba(239, 68, 68, 0.15);
  color: var(--danger);
}
</style>
