<template>
  <div class="section-editor">
    <!-- Hero Editor -->
    <template v-if="section.type === 'hero'">
      <div class="editor-field">
        <label class="form-label">ชื่อ</label>
        <input v-model="form.name" class="form-input" placeholder="Your Name" @input="emit('update', form)" />
      </div>
      <div class="editor-field">
        <label class="form-label">ตำแหน่ง / Role</label>
        <input v-model="form.role" class="form-input" placeholder="Full-Stack Developer" @input="emit('update', form)" />
      </div>
      <div class="editor-field">
        <label class="form-label">Tagline</label>
        <textarea v-model="form.tagline" class="form-input" rows="2" placeholder="Building cool things..." @input="emit('update', form)" style="resize:vertical"></textarea>
      </div>
      <div class="editor-field">
        <label class="form-label">รูป Avatar URL</label>
        <input v-model="form.avatar_url" class="form-input" placeholder="https://..." @input="emit('update', form)" />
      </div>
      <div class="editor-field toggle-row">
        <label class="form-label" style="margin:0">แสดงปุ่ม Hire Me</label>
        <button class="toggle" :class="{ on: form.show_hire_me }" @click="form.show_hire_me = !form.show_hire_me; emit('update', form)">
          {{ form.show_hire_me ? 'ON' : 'OFF' }}
        </button>
      </div>
      <button class="btn btn-secondary btn-sm ai-btn" @click="improveTagline">✨ AI Improve Tagline</button>
    </template>

    <!-- Skills Editor -->
    <template v-else-if="section.type === 'skills'">
      <div class="skills-editor">
        <div v-for="(skill, i) in form.items" :key="i" class="skill-row">
          <input v-model="skill.name" class="form-input" style="flex:1" placeholder="Skill name" @input="emit('update', form)" />
          <input v-model.number="skill.level" type="range" min="0" max="100" style="flex:1" @input="emit('update', form)" />
          <span class="skill-pct">{{ skill.level }}%</span>
          <button class="icon-btn danger" @click="removeSkill(i)">×</button>
        </div>
        <button class="btn btn-secondary btn-sm" style="width:100%;margin-top:8px" @click="addSkill">+ เพิ่ม Skill</button>
      </div>
    </template>

    <!-- Experience Editor -->
    <template v-else-if="section.type === 'experience'">
      <div v-for="(item, i) in form.items" :key="i" class="exp-editor-item">
        <div class="exp-editor-header">
          <span style="font-size:12px;font-weight:700">งาน #{{ i + 1 }}</span>
          <button class="icon-btn danger" style="margin-left:auto" @click="removeExp(i)">× ลบ</button>
        </div>
        <input v-model="item.company" class="form-input" placeholder="Company" @input="emit('update', form)" />
        <input v-model="item.position" class="form-input" placeholder="Position" @input="emit('update', form)" />
        <div style="display:grid;grid-template-columns:1fr 1fr;gap:6px">
          <input v-model="item.start_date" class="form-input" placeholder="2022-01" @input="emit('update', form)" />
          <input v-model="item.end_date" class="form-input" placeholder="2024-01" :disabled="item.is_current" @input="emit('update', form)" />
        </div>
        <textarea v-model="item.description" class="form-input" rows="2" placeholder="Describe your role..." @input="emit('update', form)" style="resize:vertical"></textarea>
      </div>
      <button class="btn btn-secondary btn-sm" style="width:100%;margin-top:8px" @click="addExp">+ เพิ่มประสบการณ์</button>
    </template>

    <!-- Projects Editor -->
    <template v-else-if="section.type === 'projects'">
      <div v-for="(proj, i) in form.items" :key="i" class="proj-editor-item">
        <div class="exp-editor-header">
          <span style="font-size:12px;font-weight:700">Project #{{ i + 1 }}</span>
          <button class="icon-btn danger" style="margin-left:auto" @click="removeProj(i)">× ลบ</button>
        </div>
        <input v-model="proj.title" class="form-input" placeholder="Project Title" @input="emit('update', form)" />
        <textarea v-model="proj.description" class="form-input" rows="2" placeholder="Description" @input="emit('update', form)" style="resize:vertical"></textarea>
        <input v-model="proj.live_url" class="form-input" placeholder="Live URL (optional)" @input="emit('update', form)" />
        <input v-model="proj.github_url" class="form-input" placeholder="GitHub URL (optional)" @input="emit('update', form)" />
        <input
          :value="(proj.tags || []).join(', ')"
          class="form-input"
          placeholder="Tags: Go, Vue.js, Docker"
          @input="proj.tags = ($event.target as HTMLInputElement).value.split(',').map(s => s.trim()); emit('update', form)"
        />
      </div>
      <button class="btn btn-secondary btn-sm" style="width:100%;margin-top:8px" @click="addProj">+ เพิ่ม Project</button>
    </template>

    <!-- Contact Editor -->
    <template v-else-if="section.type === 'contact'">
      <div class="editor-field">
        <label class="form-label">Email</label>
        <input v-model="form.email" class="form-input" placeholder="you@example.com" @input="emit('update', form)" />
      </div>
      <div class="editor-field">
        <label class="form-label">LinkedIn</label>
        <input v-model="form.linkedin" class="form-input" placeholder="linkedin.com/in/..." @input="emit('update', form)" />
      </div>
      <div class="editor-field">
        <label class="form-label">GitHub</label>
        <input v-model="form.github" class="form-input" placeholder="github.com/..." @input="emit('update', form)" />
      </div>
      <div class="editor-field">
        <label class="form-label">Location</label>
        <input v-model="form.location" class="form-input" placeholder="Bangkok, Thailand" @input="emit('update', form)" />
      </div>
    </template>

    <!-- Custom Text Editor -->
    <template v-else-if="section.type === 'custom_text'">
      <div class="editor-field">
        <label class="form-label">Content</label>
        <textarea v-model="form.content" class="form-input" rows="5" placeholder="เขียนอะไรก็ได้..." @input="emit('update', form)" style="resize:vertical"></textarea>
      </div>
    </template>

    <!-- Fallback -->
    <template v-else>
      <p style="font-size:12px;color:var(--muted);text-align:center;padding:12px 0">
        Editor สำหรับ "{{ section.type }}" อยู่ระหว่างพัฒนา
      </p>
    </template>
  </div>
</template>

<script setup lang="ts">
import { reactive, watch } from 'vue'
import type { Section } from '@/types'
import { aiAPI } from '@/api'

const props = defineProps<{ section: Section }>()
const emit = defineEmits<{ update: [data: any] }>()

// Deep clone section data into reactive form
const form = reactive(JSON.parse(JSON.stringify(props.section.data || {})))

watch(() => props.section.id, () => {
  Object.assign(form, JSON.parse(JSON.stringify(props.section.data || {})))
})

// Skills helpers
function addSkill() {
  if (!form.items) form.items = []
  form.items.push({ name: '', level: 70, category: '' })
  emit('update', form)
}
function removeSkill(i: number) {
  form.items.splice(i, 1)
  emit('update', form)
}

// Experience helpers
function addExp() {
  if (!form.items) form.items = []
  form.items.push({ company: '', position: '', start_date: '', end_date: '', is_current: false, description: '' })
  emit('update', form)
}
function removeExp(i: number) {
  form.items.splice(i, 1)
  emit('update', form)
}

// Project helpers
function addProj() {
  if (!form.items) form.items = []
  form.items.push({ title: '', description: '', live_url: '', github_url: '', tags: [] })
  emit('update', form)
}
function removeProj(i: number) {
  form.items.splice(i, 1)
  emit('update', form)
}

// AI improve tagline
async function improveTagline() {
  if (!form.tagline) return
  try {
    const { data } = await aiAPI.improveText(form.tagline, 'portfolio tagline for ' + form.role)
    form.tagline = data.data.improved_text
    emit('update', form)
  } catch {}
}
</script>

<style scoped>
.section-editor { display: flex; flex-direction: column; gap: 10px; }
.editor-field { display: flex; flex-direction: column; }
.toggle-row { flex-direction: row; align-items: center; justify-content: space-between; }
.toggle {
  padding: 5px 14px; border-radius: 100px; border: 1px solid var(--border);
  background: rgba(255,255,255,0.05); color: var(--muted);
  font-size: 12px; font-weight: 700; cursor: pointer; transition: all 0.15s;
}
.toggle.on { background: rgba(79,70,229,0.2); border-color: var(--indigo); color: #818cf8; }

.skills-editor { display: flex; flex-direction: column; gap: 8px; }
.skill-row { display: flex; align-items: center; gap: 6px; }
.skill-pct { font-size: 11px; color: var(--muted); width: 30px; text-align: right; flex-shrink: 0; }

.exp-editor-item, .proj-editor-item {
  background: rgba(255,255,255,0.02); border: 1px solid var(--border);
  border-radius: 10px; padding: 12px; display: flex; flex-direction: column; gap: 7px;
}
.exp-editor-header { display: flex; align-items: center; margin-bottom: 4px; }

.ai-btn { margin-top: 4px; background: rgba(168,85,247,0.1); border-color: rgba(168,85,247,0.3); color: var(--neon-purple); }

.icon-btn {
  width: 26px; height: 26px; border-radius: 6px; border: none;
  background: rgba(255,255,255,0.05); color: var(--muted);
  font-size: 14px; cursor: pointer; transition: all 0.15s;
  display: flex; align-items: center; justify-content: center; flex-shrink: 0;
}
.icon-btn.danger:hover { background: rgba(239,68,68,0.15); color: var(--danger); }
</style>
