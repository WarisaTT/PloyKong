<template>
  <div class="public-portfolio" :class="themeClass" :style="themeVars">
    <!-- Loading -->
    <div v-if="loading" class="pub-loading">
      <div class="pub-spinner"></div>
      <p>Loading...</p>
    </div>

    <!-- Password protected -->
    <div v-else-if="requiresPassword" class="pub-password-wall">
      <div class="pw-card">
        <div class="pw-icon"><Lock :size="48" style="stroke-width: 1.5px" /></div>
        <h2>This portfolio is locked</h2>
        <p>Please enter the password to continue.</p>
        <input v-model="password" type="password" class="form-input" placeholder="ใส่รหัสผ่าน" @keyup.enter="loadWithPassword" />
        <button class="btn btn-primary" style="width:100%;margin-top:8px" @click="loadWithPassword">ปลดล็อก</button>
        <p v-if="pwError" class="pw-error">{{ pwError }}</p>
      </div>
    </div>

    <!-- 404 -->
    <div v-else-if="notFound" class="pub-404">
      <div class="not-found-content">
        <div style="margin-bottom:16px"><Search :size="64" style="stroke-width: 1.5px" /></div>
        <h1>Portfolio Not Found</h1>
        <p>This portfolio may have been deleted or the URL is incorrect.</p>
        <a href="https://app.ploykong.com" class="btn btn-primary" style="margin-top:24px">Back to Home</a>
      </div>
    </div>

    <!-- Portfolio Content -->
    <div v-else-if="portfolio">
      <!-- SEO Meta (set via document.title) -->

      <!-- Render each section -->
      <div class="sections-container">
        <template v-for="section in visibleSections" :key="section.id">
          <div :class="{ 'pub-container': section.type !== 'hero' }">
            <component
              :is="getSectionComponent(section.type)"
              :data="section.data"
              :theme="portfolio.theme"
              :portfolio="portfolio"
            />
          </div>
        </template>
      </div>

      <!-- AI Chat FAB -->
      <div v-if="hasAIChat" class="chat-fab" @click="chatOpen = !chatOpen">
        <X v-if="chatOpen" :size="24" />
        <Bot v-else :size="24" />
      </div>

      <!-- AI Chat Panel -->
      <Transition name="chat-slide">
        <div v-if="chatOpen" class="chat-panel">
          <div class="chat-header">
            <span><Bot :size="14" class="icon-inline" /> AI Interview Coach</span>
            <span class="chat-sub">ถามฉันเกี่ยวกับ {{ portfolio.title }}</span>
          </div>
          <div class="chat-messages" ref="chatMessagesRef">
            <div
              v-for="(msg, i) in chatMessages"
              :key="i"
              class="chat-msg"
              :class="msg.role"
            >
              <span class="chat-bubble">{{ msg.content }}</span>
            </div>
            <div v-if="chatLoading" class="chat-msg assistant">
              <span class="chat-bubble chat-typing"><Loader2 :size="14" class="spin icon-inline" /> กำลังคิด...</span>
            </div>
          </div>
          <div class="chat-input-row">
            <input
              v-model="chatInput"
              class="form-input"
              placeholder="ถามอะไรก็ได้เกี่ยวกับ portfolio นี้..."
              @keyup.enter="sendChat"
            />
            <button class="btn btn-primary btn-sm" @click="sendChat" :disabled="chatLoading">ส่ง</button>
          </div>
        </div>
      </Transition>

      <!-- Watermark / Powered by -->
      <div class="pk-watermark">
        <a href="https://app.ploykong.com" target="_blank">
          Made with <strong>PloyKong</strong> <Zap :size="12" class="icon-inline" />
        </a>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, nextTick } from 'vue'
import { useRoute } from 'vue-router'
import { publicAPI } from '@/api'
import type { Portfolio } from '@/types'
import { Lock, Search, Bot, X, Loader2, Zap } from 'lucide-vue-next'

// Public section renderers
import PublicHero from '@/components/portfolio/PublicHero.vue'
import PublicSkills from '@/components/portfolio/PublicSkills.vue'
import PublicProjects from '@/components/portfolio/PublicProjects.vue'
import PublicExperience from '@/components/portfolio/PublicExperience.vue'
import PublicContact from '@/components/portfolio/PublicContact.vue'
import PublicCustomText from '@/components/portfolio/PublicCustomText.vue'

const route = useRoute()
const slug = computed(() => route.params.slug as string)

const portfolio = ref<Portfolio | null>(null)
const loading = ref(true)
const notFound = ref(false)
const requiresPassword = ref(false)
const password = ref('')
const pwError = ref('')

const chatOpen = ref(false)
const chatInput = ref('')
const chatLoading = ref(false)
const chatMessages = ref<{ role: 'user' | 'assistant'; content: string }[]>([
  { role: 'assistant', content: 'สวัสดีครับ! ผมเป็น AI ตัวแทนของเจ้าของพอร์ตนี้ ถามอะไรก็ได้เลยครับ :)' }
])
const chatMessagesRef = ref<HTMLElement | null>(null)
const sessionId = ref(Math.random().toString(36).slice(2))

const themeClass = computed(() => {
  const mode = portfolio.value?.theme?.mode || 'system'
  return `theme-${mode}`
})
const themeVars = computed(() => {
  const primary = portfolio.value?.theme?.primary_color || '#4F46E5'
  const secondary = portfolio.value?.theme?.secondary_color || ''
  const font = portfolio.value?.theme?.font || 'Plus Jakarta Sans'
  const isLight = portfolio.value?.theme?.mode === 'light'
  const glowHex = isLight ? '15' : '40'
  const vars: Record<string, string> = {
    '--primary': primary,
    '--primary-glow': `${primary}${glowHex}`,
    '--secondary': secondary ? secondary : `${primary}${glowHex}`,
    '--font-display': font,
    '--font-body': font
  }
  const bg = portfolio.value?.theme?.bg_color
  const border = portfolio.value?.theme?.border_color
  if (bg && bg !== '#000000') vars['--bg'] = bg
  if (border && border !== '#000000') vars['--avatar-border'] = border
  return vars
})

const visibleSections = computed(() =>
  (portfolio.value?.sections || [])
    .filter((s) => s.is_visible)
    .sort((a, b) => a.position - b.position)
)

const hasAIChat = computed(() =>
  visibleSections.value.some((s) => s.type === 'ai_chat')
)

function getSectionComponent(type: string) {
  const map: Record<string, any> = {
    hero: PublicHero,
    skills: PublicSkills,
    projects: PublicProjects,
    experience: PublicExperience,
    contact: PublicContact,
    custom_text: PublicCustomText,
  }
  return map[type] || PublicCustomText
}

async function loadPortfolio(pw?: string) {
  loading.value = true
  try {
    const { data } = await publicAPI.view(slug.value, pw)
    portfolio.value = data.data
    document.title = `${data.data.title} | PloyKong`
    // Track view
    publicAPI.track(slug.value, 'view', document.referrer || 'direct')
  } catch (e: any) {
    if (e.response?.data?.requires_password) {
      requiresPassword.value = true
    } else {
      notFound.value = true
    }
  } finally {
    loading.value = false
  }
}

async function loadWithPassword() {
  pwError.value = ''
  try {
    const { data } = await publicAPI.view(slug.value, password.value)
    portfolio.value = data.data
    requiresPassword.value = false
  } catch {
    pwError.value = 'รหัสผ่านไม่ถูกต้อง'
  }
}

async function sendChat() {
  if (!chatInput.value.trim() || chatLoading.value) return
  const msg = chatInput.value.trim()
  chatInput.value = ''
  chatMessages.value.push({ role: 'user', content: msg })
  chatLoading.value = true

  await nextTick()
  if (chatMessagesRef.value) {
    chatMessagesRef.value.scrollTop = chatMessagesRef.value.scrollHeight
  }

  try {
    const { data } = await publicAPI.chat(slug.value, msg, sessionId.value)
    chatMessages.value.push({ role: 'assistant', content: data.data.response })
    publicAPI.track(slug.value, 'ai_chat')
  } catch {
    chatMessages.value.push({ role: 'assistant', content: 'ขออภัย เกิดข้อผิดพลาด กรุณาลองใหม่' })
  } finally {
    chatLoading.value = false
    await nextTick()
    if (chatMessagesRef.value) {
      chatMessagesRef.value.scrollTop = chatMessagesRef.value.scrollHeight
    }
  }
}

onMounted(() => loadPortfolio())
</script>

<style scoped>
.public-portfolio {
  min-height: 100vh;
  position: relative;
  color: var(--text);
  background: var(--bg);
}
.public-portfolio.theme-light {
  background: linear-gradient(to bottom right, #faf7ff, #f3e8ff);
}

/* Loading */
.pub-spinner { width: 40px; height: 40px; border: 3px solid var(--border); border-top-color: var(--purple); border-radius: 50%; animation: spin 0.8s linear infinite; }
@keyframes spin { to { transform: rotate(360deg); } }

/* Password wall, 404, Loading */
.pub-password-wall,
.pub-404,
.pub-loading {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  background: var(--bg);
  color: var(--text);
}
:global(.theme-light) .pub-password-wall,
:global(.theme-light) .pub-404,
:global(.theme-light) .pub-loading {
  background: linear-gradient(to bottom right, #faf7ff, #f3e8ff);
}
.pub-loading { gap: 16px; color: var(--muted); }
.pub-password-wall { padding: 24px; }

.pw-card { background: var(--surface); border: 1px solid var(--border); border-radius: 24px; padding: 44px; width: 100%; max-width: 380px; text-align: center; backdrop-filter: blur(20px); }
.pw-icon { font-size: 48px; margin-bottom: 16px; }
.pw-card h2 { font-size: 20px; margin-bottom: 20px; }
.pw-error { color: var(--danger); font-size: 13px; margin-top: 8px; }

/* 404 */
.pub-404 { display: flex; align-items: center; justify-content: center; min-height: 100vh; }
.not-found-content { text-align: center; }
.not-found-content h1 { font-size: 28px; margin-bottom: 8px; }
.not-found-content p { color: var(--muted); }

/* Sections */
.sections-container { padding-bottom: 80px; }
.pub-container { max-width: 900px; margin: 0 auto; padding: 0 24px; }

/* AI Chat FAB */
.chat-fab {
  position: fixed; bottom: 28px; right: 28px;
  width: 56px; height: 56px; border-radius: 50%;
  background: linear-gradient(135deg, var(--indigo), var(--purple));
  display: flex; align-items: center; justify-content: center;
  font-size: 22px; cursor: pointer; z-index: 100;
  box-shadow: 0 8px 24px rgba(79, 70, 229, 0.5);
  transition: transform 0.2s;
}
.chat-fab:hover { transform: scale(1.1); }

/* Chat Panel */
.chat-panel {
  position: fixed; bottom: 96px; right: 28px;
  width: 340px; max-height: 480px;
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 20px; z-index: 100;
  display: flex; flex-direction: column;
  box-shadow: 0 24px 60px rgba(0,0,0,0.5);
  backdrop-filter: blur(20px);
}
:global(.theme-dark) .chat-panel {
  background: rgba(10, 14, 30, 0.98);
}
:global(.theme-light) .chat-panel {
  box-shadow: 0 24px 60px rgba(0,0,0,0.15);
}
.chat-header { padding: 16px; border-bottom: 1px solid var(--border); }
.chat-header span:first-child { font-size: 14px; font-weight: 700; display: block; }
.chat-sub { font-size: 11px; color: var(--muted); }
.chat-messages { flex: 1; overflow-y: auto; padding: 14px; display: flex; flex-direction: column; gap: 10px; min-height: 0; }
.chat-msg { display: flex; }
.chat-msg.user { justify-content: flex-end; }
.chat-bubble {
  max-width: 80%; padding: 9px 14px; border-radius: 14px;
  font-size: 13px; line-height: 1.5;
}
.chat-msg.user .chat-bubble { background: linear-gradient(135deg, var(--indigo), var(--purple)); color: #fff; }
.chat-msg.assistant .chat-bubble { background: var(--bg2); border: 1px solid var(--border); }
.chat-typing { color: var(--muted); font-style: italic; }
.chat-input-row { padding: 12px; border-top: 1px solid var(--border); display: flex; gap: 8px; }
.chat-input-row .form-input { flex: 1; padding: 8px 12px; font-size: 13px; }

/* Chat slide transition */
.chat-slide-enter-active, .chat-slide-leave-active { transition: all 0.25s ease; }
.chat-slide-enter-from, .chat-slide-leave-to { opacity: 0; transform: translateY(16px) scale(0.96); }

/* Watermark */
.pk-watermark { text-align: center; padding: 24px; }
.pk-watermark a { font-size: 12px; color: var(--muted); text-decoration: none; }
.pk-watermark a:hover { color: var(--neon-purple); }
.pk-watermark strong { color: var(--neon-purple); }


.icon-inline { vertical-align: text-bottom; margin-right: 2px; }
@keyframes spin { 100% { transform: rotate(360deg); } }
.spin { animation: spin 1s linear infinite; }

@media (max-width: 480px) {
  .chat-panel { right: 12px; left: 12px; width: auto; }
  .chat-fab { bottom: 16px; right: 16px; }
}
</style>
