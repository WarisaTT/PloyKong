<template>
  <div class="section-editor">
    <div class="generic-settings-card">
      <div v-if="isLayoutAdjustable" class="editor-field" style="margin-bottom: 12px">
        <label class="form-label">Layout (การจัดเรียง)</label>
        <select
          v-model="form.layout"
          class="form-input form-select"
          @change="emitUpdate()"
        >
          <option value="centered">Centered (กึ่งกลาง)</option>
          <option value="left">Left Aligned (ชิดซ้าย)</option>
          <option value="split">Split (แบ่งซ้ายขวา)</option>
        </select>
      </div>

      <div v-if="section.type !== 'hero'" class="editor-field toggle-row" style="margin-bottom: 12px">
        <label class="form-label" style="margin: 0">ซ่อนหัวข้อ</label>
        <button
          class="toggle"
          :class="{ on: !!form.hide_title }"
          @click="form.hide_title = !form.hide_title; emitUpdate()"
        >
          {{ form.hide_title ? "Show" : "Hide" }}
        </button>
      </div>

      <div v-if="section.type !== 'hero'" class="editor-field toggle-row" style="margin-bottom: 12px">
        <label class="form-label" style="margin: 0">ซ่อนขีดคั่น</label>
        <button
          class="toggle"
          :class="{ on: !!form.hide_divider }"
          @click="form.hide_divider = !form.hide_divider; emitUpdate()"
        >
          {{ form.hide_divider ? "Show" : "Hide" }}
        </button>
      </div>

      <div v-if="section.type !== 'hero'" class="editor-field" style="margin-bottom: 12px">
        <label class="form-label">Section Title (หัวข้อ)</label>
        <input
          v-model="form.title"
          class="form-input"
          placeholder="ระบุหัวข้อที่ต้องการ..."
          @input="emitUpdate()"
        />
      </div>

      <div class="editor-field">
        <label class="form-label">Cover Image</label>
        <div class="image-input-row">
          <input
            v-model="form.image_url"
            class="form-input"
            placeholder="https:// หรืออัปโหลดรูป..."
            @input="emitUpdate()"
          />
          <input
            type="file"
            ref="coverInputRef"
            accept="image/*"
            style="display: none"
            @change="uploadCover"
          />
          <button
            type="button"
            class="btn btn-secondary btn-sm"
            @click="coverInputRef?.click()"
            :disabled="uploadingCover"
          >
            <template v-if="uploadingCover"
              ><Loader2 :size="14" class="spin" /> กำลังอัพโหลด...</template
            >
            <template v-else>Browse</template>
          </button>
          <button
            v-if="form.image_url"
            type="button"
            class="btn btn-secondary btn-sm danger-light"
            @click="form.image_url = ''; emitUpdate()"
          >
            <X :size="14" />
          </button>
        </div>
      </div>

      <!-- Section BG Color (firstjobber only) -->
      <div v-if="template === 'firstjobber'" class="editor-field" style="margin-top: 10px">
        <label class="form-label">Section Background Color</label>
        <div style="display: flex; align-items: center; gap: 8px">
          <input
            type="color"
            v-model="form.section_bg_color"
            style="width: 36px; height: 28px; border: 1px solid var(--border); border-radius: 6px; cursor: pointer; padding: 0"
            @input="emitUpdate()"
          />
          <input
            v-model="form.section_bg_color"
            class="form-input"
            style="flex: 1; font-size: 12px"
            placeholder="#1e1b2e"
            @input="emitUpdate()"
          />
          <button
            v-if="form.section_bg_color"
            class="btn btn-secondary btn-sm"
            style="font-size: 14px; padding: 11px 16px;"
            @click="form.section_bg_color = ''; emitUpdate()"
          >
            Reset
          </button>
        </div>
      </div>
    </div>
    
    <div class="editor-divider" style="margin: 2px 2;"></div>

    <!-- Hero Editor -->
    <template v-if="section.type === 'hero'">
      <div class="editor-field">
        <label class="form-label">Name</label>
        <AutoResizeTextarea
          v-model="form.name"
          placeholder="Your Name"
          @input="emitUpdate()"
        />
      </div>
      <div class="editor-field">
        <label class="form-label">Position</label>
        <AutoResizeTextarea
          v-model="form.role"
          placeholder="Full-Stack Developer"
          @input="emitUpdate()"
        />
      </div>
      <div class="editor-field">
        <label class="form-label">Tagline</label>
        <PopupTextEditor
          v-model="form.tagline"
          title="Tagline"
          placeholder="Building cool things..."
          :rows="4"
          @update:modelValue="emitUpdate()"
        />
      </div>
      <div class="editor-field">
        <label class="form-label">Profile Picture</label>
        <div class="image-input-row">
          <input
            v-model="form.avatar_url"
            class="form-input"
            placeholder="https://..."
            @input="emitUpdate()"
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
          <button
            v-if="form.avatar_url"
            type="button"
            class="btn btn-secondary btn-sm danger-light"
            @click="form.avatar_url = ''; emitUpdate()"
          >
            <X :size="14" />
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
            emitUpdate();
          "
        >
          {{ form.show_hire_me ? "ON" : "OFF" }}
        </button>
      </div>
      <div class="editor-field toggle-row">
        <label class="form-label" style="margin: 0">Resume / CV Button</label>
        <button
          class="toggle"
          :class="{ on: form.show_resume }"
          @click="
            form.show_resume = !form.show_resume;
            emitUpdate();
          "
        >
          {{ form.show_resume ? "ON" : "OFF" }}
        </button>
      </div>


    </template>

    <!-- Skills Editor -->
    <template v-else-if="section.type === 'skills'">
      <div class="skills-editor">
        <div
          v-for="(skill, i) in form.items as any[]"
          :key="i"
          class="skill-row"
        >
          <div style="display: flex; flex-direction: column; gap: 8px; width: 100%">
            <AutoResizeTextarea
              v-model="skill.name"
              :style="{ flex: 2 }"
              placeholder="Skill Name"
              @input="emitUpdate()"
            />
            <AutoResizeTextarea
              v-model="skill.category"
              :style="{ flex: 1 }"
              placeholder="Type"
              @input="emitUpdate()"
            />
            <button class="icon-btn danger" style="margin-left: auto;" @click="removeSkill(i)" title="Remove Skill">
              <X :size="16" />
            </button>
          </div>
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
            <Trash2 :size="14" />
          </button>
        </div>
        <AutoResizeTextarea
          v-model="item.company"
          placeholder="Company"
          @input="emitUpdate()"
        />
        <AutoResizeTextarea
          v-model="item.position"
          placeholder="Position"
          @input="emitUpdate()"
        />
        <AutoResizeTextarea
          v-model="item.location"
          placeholder="Location (e.g. Bangkok, Remote)"
          @input="emitUpdate()"
        />
        <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 6px">
          <input
            v-model="item.start_date"
            class="form-input"
            placeholder="2022-01"
            @input="emitUpdate()"
          />
          <input
            v-model="item.end_date"
            class="form-input"
            placeholder="2024-01"
            :disabled="item.is_current"
            @input="emitUpdate()"
          />
        </div>
        <div class="editor-field toggle-row" style="margin-bottom: 8px">
          <label class="form-label" style="margin: 0; font-size: 11px">ทำงานอยู่ที่นี่ (Current)</label>
          <button
            class="toggle"
            :class="{ on: item.is_current }"
            @click="item.is_current = !item.is_current; emitUpdate()"
          >
            {{ item.is_current ? "YES" : "NO" }}
          </button>
        </div>
        <PopupTextEditor
          v-model="item.description"
          title="Experience Description"
          placeholder="Describe your role..."
          :rows="6"
          @update:modelValue="emitUpdate()"
        />

        <div class="editor-field" style="margin-top: 8px;">
          <label class="form-label">รูปภาพประกอบ (Gallery)</label>
          <div v-if="item.image_urls && item.image_urls.length > 0" class="exp-gallery-preview">
            <div v-for="(img, idx) in item.image_urls" :key="idx" class="exp-img-thumbnail">
              <img :src="img" />
              <button class="icon-btn danger btn-sm delete-btn" @click="removeExpImage(Number(i), Number(idx))">
                <X :size="12" />
              </button>
            </div>
          </div>
          
          <div style="display: flex; gap: 8px; margin-top: 4px;">
            <input type="file" :ref="el => setExpInputRef(el, Number(i))" accept="image/*" style="display: none" @change="e => uploadExpImage(e, Number(i))" />
            <button class="btn btn-secondary btn-sm" @click="triggerExpUpload(Number(i))" :disabled="uploadingExp === Number(i)">
              <template v-if="uploadingExp === Number(i)"><Loader2 :size="14" class="spin" /> กำลังอัพโหลด...</template>
              <template v-else><Plus :size="14" /> เพิ่มรูปภาพ</template>
            </button>
          </div>
        </div>
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
            <Trash2 :size="14" />
          </button>
        </div>
        <AutoResizeTextarea
          v-model="proj.title"
          placeholder="Project Title"
          @input="emitUpdate()"
        />
        <PopupTextEditor
          v-model="proj.description"
          title="Project Description"
          placeholder="Description"
          :rows="6"
          @update:modelValue="emitUpdate()"
        />
        <input
          v-model="proj.live_url"
          class="form-input"
          placeholder="Live URL (optional)"
          @input="emitUpdate()"
        />
        <input
          v-model="proj.github_url"
          class="form-input"
          placeholder="GitHub URL (optional)"
          @input="emitUpdate()"
        />
        <input
          :value="(proj.tags || []).join(', ')"
          class="form-input"
          placeholder="Tags: Go, Vue.js, Docker"
          @input="
            proj.tags = ($event.target as HTMLInputElement).value
              .split(',')
              .map((s) => s.trim());
            emitUpdate();
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

    <!-- Certificates Editor -->
    <template v-else-if="section.type === 'certificates'">
      <div v-for="(cert, i) in form.items as any[]" :key="i" class="proj-editor-item">
        <div class="exp-editor-header">
          <span style="font-size: 12px; font-weight: 700">Certificate #{{ i + 1 }}</span>
          <button class="icon-btn danger btn-icon-text" style="margin-left: auto; width: auto; padding: 0 8px" @click="removeCert(Number(i))">
            <Trash2 :size="14" />
          </button>
        </div>
        <AutoResizeTextarea v-model="cert.title" placeholder="Certificate Title" @input="emitUpdate()" />
        <AutoResizeTextarea v-model="cert.issuer" placeholder="Issuer (e.g. Coursera, Google)" @input="emitUpdate()" />
        <input v-model="cert.date" class="form-input" placeholder="Issue Date (e.g. 2023)" @input="emitUpdate()" />
        <PopupTextEditor v-model="cert.description" title="Description (Optional)" placeholder="Additional details..." :rows="3" @update:modelValue="emitUpdate()" />
        
        <div class="editor-field" style="margin-top: 8px;">
          <label class="form-label">รูปใบประกาศ</label>
          <div class="image-input-row">
            <input v-model="cert.image_url" class="form-input" placeholder="https://..." @input="emitUpdate()" />
            <input type="file" :ref="el => setCertInputRef(el, Number(i))" accept="image/*" style="display: none" @change="e => uploadCertImage(e, Number(i))" />
            <button type="button" class="btn btn-secondary btn-sm" @click="triggerCertUpload(Number(i))" :disabled="uploadingCert === Number(i)">
              <template v-if="uploadingCert === Number(i)"><Loader2 :size="14" class="spin" /></template>
              <template v-else>Browse</template>
            </button>
            <button
              v-if="cert.image_url"
              type="button"
              class="btn btn-secondary btn-sm danger-light"
              @click="cert.image_url = ''; emitUpdate()"
            >
              <X :size="14" />
            </button>
          </div>
        </div>
      </div>
      <button class="btn btn-secondary btn-sm btn-icon-text" style="width: 100%; margin-top: 8px" @click="addCert">
        <Plus :size="16" /> Add Certificate
      </button>
    </template>

    <!-- Contact Editor -->
    <template v-else-if="section.type === 'contact'">
      <div class="editor-field">
        <label class="form-label">Email</label>
        <input
          v-model="form.email"
          class="form-input"
          placeholder="ploykong@example.com"
          @input="emitUpdate()"
        />
      </div>
      <div class="editor-field">
        <label class="form-label">LinkedIn</label>
        <input
          v-model="form.linkedin"
          class="form-input"
          placeholder="linkedin.com/in/..."
          @input="emitUpdate()"
        />
      </div>
      <div class="editor-field">
        <label class="form-label">GitHub</label>
        <input
          v-model="form.github"
          class="form-input"
          placeholder="github.com/..."
          @input="emitUpdate()"
        />
      </div>
      <div class="editor-field" style="margin-bottom: 24px">
        <label class="form-label">Location</label>
        <input
          v-model="form.location"
          class="form-input"
          placeholder="Bangkok, Thailand"
          @input="emitUpdate()"
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
            @change="emitUpdate()"
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
              @input="emitUpdate()"
            />
          </div>
          <div class="editor-field">
            <label class="form-label">Value</label>
            <input
              v-model="contact.value"
              class="form-input"
              placeholder="Value"
              @input="emitUpdate()"
            />
          </div>
        </div>

        <div class="editor-field" style="margin-top: 6px">
          <label class="form-label">URL / Link Action</label>
          <input
            v-model="contact.link"
            class="form-input"
            placeholder="https://... or tel:0812345678"
            @input="emitUpdate()"
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
    
    <!-- AI Chat Editor -->
    <template v-else-if="section.type === 'ai_chat'">
      <div class="editor-field">
        <label class="form-label">Custom Info / Training Data (สิ่งที่ AI ควรรู้เกี่ยวกับคุณ)</label>
        <PopupTextEditor
          v-model="form.prompt_hint"
          title="Training Data"
          placeholder="ข้อมูลนี้จะถูกส่งไปให้ AI เพื่อให้ตอบคำถามได้แม่นยำขึ้น เช่น 'ฉันชอบแมว', 'ฉันทำงานที่ Google มาก่อน', 'ฉันเชี่ยวชาญด้าน React เป็นพิเศษ'..."
          :rows="6"
          @update:modelValue="emitUpdate()"
        />
      </div>

      <div class="editor-field" style="margin-top: 12px">
        <label class="form-label" style="display: flex; justify-content: space-between; align-items: center;">
          <span>AI Knowledge Center (คำถามที่ AI ตอบไม่ได้)</span>
          <button class="btn-text btn-sm" @click="fetchGaps" :disabled="loadingGaps">
            <Loader2 v-if="loadingGaps" :size="12" class="spin" />
            <RefreshCw v-else :size="12" /> Refresh
          </button>
        </label>
        
        <div class="gaps-container" style="background: rgba(0,0,0,0.2); border-radius: 12px; padding: 12px; border: 1px dashed var(--border);">
          <div v-if="gaps.length === 0" style="text-align: center; color: var(--muted); padding: 12px; font-size: 13px;">
            <CheckCircle :size="24" style="margin-bottom: 8px; opacity: 0.5;" />
            <p>ยังไม่มีคำถามที่ AI ตอบไม่ได้ เยี่ยมมาก!</p>
          </div>
          <div v-for="(gap, i) in gaps" :key="i" class="gap-item" style="margin-bottom: 12px; padding-bottom: 12px; border-bottom: 1px solid rgba(255,255,255,0.05);">
            <div style="font-size: 13px; font-weight: 600; color: var(--primary);">Q: {{ gap.question }}</div>
            <div style="display: flex; gap: 8px; margin-top: 8px;">
               <input v-model="gap.answer" class="form-input" style="font-size: 12px;" placeholder="เพิ่มข้อมูลสำหรับตอบคำถามนี้..." />
               <button class="btn btn-primary btn-sm" @click="addTrainingData(gap)" :disabled="!gap.answer">Learn</button>
            </div>
          </div>
        </div>
      </div>

      <div class="editor-field" style="margin-top: 16px">
        <label class="form-label">Example Questions (คำถามตัวอย่าง)</label>
        <div v-for="(q, i) in (form.example_questions || [])" :key="i" class="q-row" style="display: flex; gap: 8px; margin-bottom: 6px">
          <input
            v-model="form.example_questions[i]"
            class="form-input"
            placeholder="เช่น 'คุณถนัดเทคโนโลยีอะไรบ้าง?'"
            @input="emitUpdate()"
          />
          <button class="icon-btn danger" @click="removeExampleQuestion(Number(i))">
            <X :size="14" />
          </button>
        </div>
        <button class="btn btn-secondary btn-sm" style="width: 100%" @click="addExampleQuestion">
          <Plus :size="14" /> Add Question
        </button>
      </div>
    </template>

    <!-- Custom Text Editor -->
    <template v-else-if="section.type === 'custom_text'">
      <div class="editor-field" style="margin-bottom: 8px;">
        <label class="form-label">หัวข้อ (Layered-title / Optional)</label>
        <AutoResizeTextarea
          v-model="form.title"
          placeholder="ใส่ชื่อหัวข้อที่ต้องการ..."
          @input="emitUpdate()"
        />
      </div>
      <div class="editor-field">
        <label class="form-label">เนื้อหาบทความ (Content)</label>
        <PopupTextEditor
          v-model="form.content"
          title="Custom Text Content"
          placeholder="เขียนอะไรก็ได้..."
          :rows="14"
          @update:modelValue="emitUpdate()"
        />
      </div>
    </template>

    <!-- Education Editor -->
    <template v-else-if="section.type === 'education'">
      <div v-for="(edu, i) in form.items as any[]" :key="i" class="proj-editor-item">
        <div class="exp-editor-header">
          <span style="font-size: 12px; font-weight: 700">Education #{{ i + 1 }}</span>
          <button class="icon-btn danger btn-icon-text" style="margin-left: auto; width: auto; padding: 0 8px" @click="removeEdu(Number(i))">
            <Trash2 :size="14" />
          </button>
        </div>
        <AutoResizeTextarea v-model="edu.school" placeholder="School / University" @input="emitUpdate()" />
        <div style="display: flex; gap: 8px;">
          <AutoResizeTextarea v-model="edu.degree" :style="{ flex: 1 }" placeholder="Degree (e.g. B.S.)" @input="emitUpdate()" />
          <AutoResizeTextarea v-model="edu.field" :style="{ flex: 2 }" placeholder="Field of Study" @input="emitUpdate()" />
        </div>
        <div style="display: flex; gap: 8px;">
          <input v-model="edu.start_year" class="form-input" style="flex: 1" placeholder="Start Year" @input="emitUpdate()" />
          <input v-model="edu.end_year" class="form-input" style="flex: 1" placeholder="End Year (or Present)" @input="emitUpdate()" />
          <input v-model="edu.gpa" class="form-input" style="flex: 1" placeholder="GPA (Optional)" @input="emitUpdate()" />
        </div>
      </div>
      <button class="btn btn-secondary btn-sm btn-icon-text" style="width: 100%; margin-top: 8px" @click="addEdu">
        <Plus :size="16" /> Add Education
      </button>
    </template>

    <!-- Stats Editor -->
    <template v-else-if="section.type === 'stats'">
      <div v-for="(stat, i) in form.items as any[]" :key="i" class="proj-editor-item">
        <div class="exp-editor-header" style="margin-bottom:0">
          <div style="display: flex; gap: 8px; width: 100%; align-items: center;">
            <input v-model="stat.value" class="form-input" style="flex: 1" placeholder="Value (e.g. 1M+)" @input="emitUpdate()" />
            <input v-model="stat.label" class="form-input" style="flex: 2" placeholder="Label (e.g. Users Served)" @input="emitUpdate()" />
            <button class="icon-btn danger" style="flex-shrink:0" @click="removeStat(Number(i))" title="Remove Stat"><Trash2 :size="14" /></button>
          </div>
        </div>
      </div>
      <button class="btn btn-secondary btn-sm btn-icon-text" style="width: 100%; margin-top: 8px" @click="addStat">
        <Plus :size="16" /> Add Stat
      </button>
    </template>

    <!-- Social Editor -->
    <template v-else-if="section.type === 'social'">
      <div v-for="(soc, i) in form.items as any[]" :key="i" class="proj-editor-item">
        <div class="exp-editor-header" style="margin-bottom:0">
          <div style="display: flex; gap: 8px; width: 100%; align-items: center;">
            <select v-model="soc.platform" class="form-input form-select" style="flex: 1; padding: 6px;" @change="emitUpdate()">
              <option value="facebook">Facebook</option>
              <option value="twitter">X / Twitter</option>
              <option value="instagram">Instagram</option>
              <option value="youtube">YouTube</option>
              <option value="linkedin">LinkedIn</option>
              <option value="github">GitHub</option>
              <option value="line">Line</option>
              <option value="website">Website</option>
            </select>
            <input v-model="soc.label" class="form-input" style="flex: 1" placeholder="Label (Optional)" @input="emitUpdate()" />
            <button class="icon-btn danger" style="flex-shrink:0" @click="removeSocial(Number(i))" title="Remove"><Trash2 :size="14" /></button>
          </div>
        </div>
        <input v-model="soc.url" class="form-input" style="margin-top: 6px;" placeholder="URL (https://...)" @input="emitUpdate()" />
      </div>
      <button class="btn btn-secondary btn-sm btn-icon-text" style="width: 100%; margin-top: 8px" @click="addSocial">
        <Plus :size="16" /> Add Social Link
      </button>
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
import { reactive, watch, ref, computed } from "vue";
import type { Section } from "@/types";
import { aiAPI, uploadAPI } from "@/api";
import PopupTextEditor from "./PopupTextEditor.vue";
import AutoResizeTextarea from "./AutoResizeTextarea.vue";
import { X, Plus, Trash2, Loader2, RefreshCw, CheckCircle } from "lucide-vue-next";
import { toastError, toastSuccess } from "@/utils/alert";
import { useRoute } from "vue-router";
import { portfolioAPI } from "@/api";

const props = defineProps<{ 
  section: Section, 
  template?: string,
  isHeroComplete?: boolean
}>();
const emit = defineEmits<{ 
  update: [data: any],
  'magic-fill': []
}>();
const route = useRoute();

// Wrapper to prevent emitting the Vue Proxy directly, which causes infinite update loops in the parent
function emitUpdate() {
  emit('update', JSON.parse(JSON.stringify(form)));
}

// Check if the current section + template combo locks the layout via CSS
const isLayoutAdjustable = computed(() => {
  const currentTemplate = props.template || 'classic';
  if (props.section.type === 'hero') {
    if (['firstjobber', 'programmer', 'business'].includes(currentTemplate)) {
      return false; // CSS uses !important to force layout, preventing Vue layout classes from working visually
    }
  }
  return true;
});

function getDefaultTitle(type: string) {
  const defaults: Record<string, string> = {
    skills: "Skills",
    experience: "Experience",
    projects: "Projects",
    education: "Education",
    contact: "Contact",
    certificates: "Certificates",
    stats: "Stats",
    social: "Connect",
  };
  return defaults[type] || "";
}

// Deep clone section data into reactive form
const initialData = JSON.parse(JSON.stringify(props.section.data || {}));
if (props.section.type === 'hero') {
  if (initialData.show_resume === undefined) initialData.show_resume = false;
  if (initialData.show_hire_me === undefined) initialData.show_hire_me = false;
  if (!initialData.name) initialData.name = "";
  if (!initialData.role) initialData.role = "";
  if (!initialData.tagline) initialData.tagline = "";
}

const form = reactive(initialData);

// --- AI Knowledge Center ---
const gaps = ref<any[]>([]);
const loadingGaps = ref(false);

async function fetchGaps() {
  if (props.section.type !== 'ai_chat') return;
  const portfolioId = route.params.id as string;
  if (!portfolioId) return;
  
  loadingGaps.value = true;
  try {
    const { data } = await portfolioAPI.getKnowledgeGaps(portfolioId);
    gaps.value = data.data.map((g: any) => ({ ...g, answer: "" }));
  } catch (e) {
    console.error("Failed to fetch gaps", e);
  } finally {
    loadingGaps.value = false;
  }
}

function addTrainingData(gap: any) {
  if (!gap.answer) return;
  
  const trainingEntry = `\nQ: ${gap.question}\nA: ${gap.answer}`;
  if (!form.prompt_hint) form.prompt_hint = "";
  form.prompt_hint += trainingEntry;
  
  // Remove from local list
  gaps.value = gaps.value.filter(g => g.question !== gap.question);
  toastSuccess("เพิ่มข้อมูลการตอบคำถามเรียบร้อย!");
  emitUpdate();
}

// Fetch gaps on load if it's an AI Chat section
if (props.section.type === 'ai_chat') {
  fetchGaps();
}

// Only overwrite local form if the selected section ID actually changes
watch(
  () => props.section.id,
  (newId, oldId) => {
    if (newId === oldId) return; // Prevent parent save cycle from clobbering local edit state

    const newData = JSON.parse(JSON.stringify(props.section.data || {}));
    if (newData.hide_percentage === undefined) newData.hide_percentage = false;
    if (newData.hide_title === undefined) newData.hide_title = false;
    if (newData.hide_divider === undefined) newData.hide_divider = false;
    if (props.section.type === 'hero') {
      if (newData.show_resume === undefined) newData.show_resume = false;
      if (newData.show_hire_me === undefined) newData.show_hire_me = false;
      if (!newData.name) newData.name = "";
      if (!newData.role) newData.role = "";
      if (!newData.tagline) newData.tagline = "";
    } else if (!newData.title) {
      newData.title = getDefaultTitle(props.section.type);
    }
    
    // Clear and replace local form deeply
    Object.keys(form).forEach(key => delete (form as any)[key]);
    Object.assign(form, newData);
  }
);

// Skills helpers
function addSkill() {
  if (!form.items) form.items = [];
  form.items.push({ name: "", level: 70, category: "" });
  emitUpdate();
}
function removeSkill(i: number) {
  form.items.splice(i, 1);
  emitUpdate();
}

// Upload Avatar (for Hero)
const avatarInputRef = ref<HTMLInputElement | null>(null);
const uploadingAvatar = ref(false);

// Upload Cover (for all Generic Sections)
const coverInputRef = ref<HTMLInputElement | null>(null);
const uploadingCover = ref(false);

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
    image_urls: []
  });
  emitUpdate();
}
function removeExp(i: number) {
  form.items.splice(i, 1);
  emitUpdate();
}

// Experience Gallery Upload
const expInputRefs = ref<Record<number, HTMLInputElement | null>>({});
const uploadingExp = ref<number | null>(null);

function setExpInputRef(el: any, i: number) {
  if (el) expInputRefs.value[i] = el;
}
function triggerExpUpload(i: number) {
  expInputRefs.value[i]?.click();
}
const currentUploadingSectionId = ref<string | null>(null);

async function safeUpload(file: File, callback: (url: string) => void) {
  const uploadId = props.section.id;
  try {
    const { data } = await uploadAPI.image(file);
    const apiBase = import.meta.env.VITE_API_URL || "/api/v1";
    let rootUrl = apiBase.replace(/\/api\/v1\/?$/, "");
    const fullUrl = rootUrl + data.data.url;
    
    // ONLY update if we are still on the same section!
    if (props.section.id === uploadId) {
      callback(fullUrl);
      emitUpdate();
    } else {
      console.warn("Upload finished but section changed. Discarding result to prevent corruption.");
    }
  } catch (error) {
    toastError("Failed to upload image");
  }
}

async function uploadAvatar(e: Event) {
  const file = (e.target as HTMLInputElement).files?.[0];
  if (!file) return;
  uploadingAvatar.value = true;
  await safeUpload(file, (url) => { form.avatar_url = url; });
  uploadingAvatar.value = false;
  if (avatarInputRef.value) avatarInputRef.value.value = "";
}

async function uploadCover(e: Event) {
  const file = (e.target as HTMLInputElement).files?.[0];
  if (!file) return;
  uploadingCover.value = true;
  await safeUpload(file, (url) => { form.image_url = url; });
  uploadingCover.value = false;
  if (coverInputRef.value) coverInputRef.value.value = "";
}

async function uploadExpImage(e: Event, i: number) {
  const file = (e.target as HTMLInputElement).files?.[0];
  if (!file) return;
  uploadingExp.value = i;
  await safeUpload(file, (url) => {
    if (!form.items[i].image_urls) form.items[i].image_urls = [];
    form.items[i].image_urls.push(url);
  });
  uploadingExp.value = null;
  if (expInputRefs.value[i]) expInputRefs.value[i]!.value = "";
}
function removeExpImage(expIndex: number, imgIndex: number) {
  form.items[expIndex].image_urls.splice(imgIndex, 1);
  emitUpdate();
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
  emitUpdate();
}
function removeProj(i: number) {
  form.items.splice(i, 1);
  emitUpdate();
}

// Certificate helpers
function addCert() {
  if (!form.items) form.items = [];
  form.items.push({
    title: "",
    issuer: "",
    date: "",
    description: "",
    image_url: ""
  });
  emitUpdate();
}
function removeCert(i: number) {
  form.items.splice(i, 1);
  emitUpdate();
}

const certInputRefs = ref<Record<number, HTMLInputElement | null>>({});
const uploadingCert = ref<number | null>(null);

function setCertInputRef(el: any, i: number) {
  if (el) certInputRefs.value[i] = el;
}
function triggerCertUpload(i: number) {
  certInputRefs.value[i]?.click();
}
async function uploadCertImage(e: Event, i: number) {
  const file = (e.target as HTMLInputElement).files?.[0];
  if (!file) return;
  uploadingCert.value = i;
  await safeUpload(file, (url) => {
    form.items[i].image_url = url;
  });
  uploadingCert.value = null;
  if (certInputRefs.value[i]) certInputRefs.value[i]!.value = "";
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
  emitUpdate();
}
function removeCustomContact(i: number) {
  form.custom_items.splice(i, 1);
  emitUpdate();
}

// Education helpers
function addEdu() {
  if (!form.items) form.items = [];
  form.items.push({ school: "", degree: "", field: "", start_year: "", end_year: "", gpa: "" });
  emitUpdate();
}
function removeEdu(i: number) {
  form.items.splice(i, 1);
  emitUpdate();
}

// Stats helpers
function addStat() {
  if (!form.items) form.items = [];
  form.items.push({ label: "", value: "" });
  emitUpdate();
}
function removeStat(i: number) {
  form.items.splice(i, 1);
  emitUpdate();
}

// Social helpers
function addSocial() {
  if (!form.items) form.items = [];
  form.items.push({ platform: "facebook", url: "", label: "" });
  emitUpdate();
}
function removeSocial(i: number) {
  form.items.splice(i, 1);
  emitUpdate();
}



// AI Chat helpers
function addExampleQuestion() {
  if (!form.example_questions) form.example_questions = [];
  form.example_questions.push("");
  emitUpdate();
}
function removeExampleQuestion(i: number) {
  form.example_questions.splice(i, 1);
  emitUpdate();
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
  padding-bottom: 8px;
}

.toggle {
  padding: 4px 12px; /* Tighter toggle padding */
  border-radius: 100px;
  border: 1px solid var(--border);
  background: var(--surface);
  color: var(--muted);
  font-size: 11px; /* Slightly smaller text */
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
  flex-direction: column;
  gap: 6px;
  background: var(--surface);
  border: 1px solid var(--border);
  padding: 20px;
  border-radius: 8px;
}

.generic-settings-card {
  background: rgba(255, 255, 255, 0.02);
  border: 1px dashed var(--border);
  border-radius: 8px;
  padding: 10px; /* Reduced from 12px */
}

.exp-editor-item,
.proj-editor-item {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 8px; /* Slightly tighter border radius */
  padding: 10px; /* Reduced from 12px */
  display: flex;
  flex-direction: column;
  gap: 6px; /* Reduced from 7px */
}
.exp-editor-header {
  display: flex;
  align-items: center;
  margin-bottom: 4px;
}

.btn-text {
  background: none;
  border: none;
  color: var(--text);
  cursor: pointer;
  font-size: 12px;
  font-weight: 600;
  padding: 4px 8px;
  border-radius: 6px;
  transition: background 0.2s;
  display: flex;
  align-items: center;
  gap: 4px;
}
.btn-text:hover {
  background: rgba(255,255,255,0.1);
  color: var(--text);
  border-color: var(--border);
  border-width: 1px;
  border-style: solid;
}
.btn-text.btn-sm {
  padding: 2px 6px;
  font-size: 11px;
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
.magic-fill-btn-inline {
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  border: none;
  animation: shimmer 2s infinite linear;
  background-size: 200% auto;
  color: #fff;
}
@keyframes shimmer {
  to { background-position: 200% center; }
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

.exp-gallery-preview {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-top: 6px;
  margin-bottom: 6px;
}
.exp-img-thumbnail {
  position: relative;
  width: 60px;
  height: 40px;
  border-radius: 4px;
  overflow: hidden;
  border: 1px solid var(--border);
}
.exp-img-thumbnail img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
.exp-img-thumbnail .delete-btn {
  position: absolute;
  top: -4px;
  right: -4px;
  width: 20px;
  height: 20px;
  padding: 0;
  border-radius: 50%;
  background: var(--danger);
  color: white;
  transform: scale(0.8);
  opacity: 0;
  transition: opacity 0.2s;
}
.exp-img-thumbnail:hover .delete-btn {
  opacity: 1;
}

.danger-light {
  background: rgba(239, 68, 68, 0.1) !important;
  border-color: rgba(239, 68, 68, 0.2) !important;
  color: #ef4444 !important;
}
.danger-light:hover {
  background: rgba(239, 68, 68, 0.2) !important;
}

.image-input-row {
  display: flex;
  gap: 8px;
}

@media (max-width: 480px) {
  .image-input-row {
    flex-wrap: wrap;
  }
  .image-input-row input {
    width: 100%;
    order: 1;
  }
  .image-input-row .btn {
    flex: 1;
    order: 2;
  }
}
</style>
