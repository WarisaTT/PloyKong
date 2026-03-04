<template>
  <div class="public-portfolio" :class="[themeClass, templateClass]" :style="themeVars">
    <!-- Loading -->
    <div v-if="loading" class="pub-loading">
      <div class="pub-spinner"></div>
      <p>Loading...</p>
    </div>

    <!-- Password protected -->
    <div v-else-if="requiresPassword" class="pub-password-wall">
      <div class="pw-card">
        <div class="pw-icon">
          <Lock :size="48" style="stroke-width: 1.5px" />
        </div>
        <h2>This portfolio is locked</h2>
        <p>Please enter the password to continue.</p>
        <input
          v-model="password"
          type="password"
          class="form-input"
          placeholder="ใส่รหัสผ่าน"
          @keyup.enter="loadWithPassword"
        />
        <button
          class="btn btn-primary"
          style="width: 100%; margin-top: 8px"
          @click="loadWithPassword"
        >
          ปลดล็อก
        </button>
        <p v-if="pwError" class="pw-error">{{ pwError }}</p>
      </div>
    </div>

    <!-- 404 -->
    <div v-else-if="notFound" class="pub-404">
      <div class="not-found-content">
        <div style="margin-bottom: 16px">
          <Search :size="64" style="stroke-width: 1.5px" />
        </div>
        <h1>Portfolio Not Found</h1>
        <p>This portfolio may have been deleted or the URL is incorrect.</p>
        <a
          href="https://app.ploykong.com"
          class="btn btn-primary"
          style="margin-top: 24px"
          >Back to Home</a
        >
      </div>
    </div>

    <!-- Portfolio Content -->
    <div v-else-if="portfolio">
      <!-- SEO Meta (set via document.title) -->

      <!-- Render sections: Grid layout for half-width sections side by side -->
      <!-- Render sections: Dynamic grouping for full-bleed vs grid-constrained -->
      <div v-for="(group, idx) in sectionGroups" :key="idx" class="section-group-wrap">
        <!-- Full-bleed Hero Group -->
        <div v-if="group.type === 'full'" class="full-bleed-section">
          <component
            :is="getSectionComponent(group.sections[0].type)"
            :data="group.sections[0].data"
            :theme="portfolio.theme"
            :portfolio="portfolio"
            class="is-full-bleed"
          />
        </div>

        <!-- Constrained Grid Group -->
        <div v-else class="sections-container">
          <TransitionGroup 
            tag="div" 
            name="section-list" 
            class="sections-grid" 
            :class="{ 'has-divider': portfolio.theme.show_divider }"
          >
            <div
              v-for="section in group.sections"
              :key="section.id"
              :data-section-type="section.type"
              class="section-wrapper"
              :class="{ 
                'span-half': section.column_span === 'half', 
                'span-full': section.column_span !== 'half', 
                'is-right': section.isRight,
                'is-row-start': section.isRowStart,
                'is-paired': section.isPaired,
                'hide-divider': !!section.data.hide_divider
              }"
            >
              <component
                :is="getSectionComponent(section.type)"
                :data="section.data"
                :theme="portfolio.theme"
                :portfolio="portfolio"
                :class="{ 'is-half-split': section.column_span === 'half' }"
              />
            </div>
          </TransitionGroup>
        </div>
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
            <span
              ><Bot :size="14" class="icon-inline" /> AI Interview Coach</span
            >
            <span class="chat-sub">ถามฉันเกี่ยวกับ {{ portfolio.title }}</span>
          </div>
          <div class="chat-messages" ref="chatMessagesRef">
            <div
              v-for="(msg, i) in chatMessages"
              :key="i"
              class="chat-msg"
              :class="msg.role"
            >
              <span class="chat-bubble" v-html="renderMarkdown(msg.content)"></span>
            </div>
            <div v-if="chatLoading" class="chat-msg assistant">
              <span class="chat-bubble chat-typing"
                ><Loader2 :size="14" class="spin icon-inline" />
                กำลังคิด...</span
              >
            </div>
          </div>
          
          <!-- Example Questions -->
          <div v-if="exampleQuestions.length > 0 && chatMessages.length <= 1" class="chat-suggestions">
            <button 
              v-for="(q, i) in exampleQuestions" 
              :key="i"
              class="suggestion-btn"
              @click="chatInput = q; sendChat()"
            >
              {{ q }}
            </button>
          </div>

          <div class="chat-input-row">
            <input
              v-model="chatInput"
              class="form-input"
              placeholder="ถามอะไรก็ได้เกี่ยวกับ portfolio นี้..."
              @keyup.enter="sendChat"
            />
            <button
              class="btn btn-primary btn-sm"
              @click="sendChat"
              :disabled="chatLoading"
            >
              ส่ง
            </button>
          </div>
        </div>
      </Transition>

      <!-- Watermark / Powered by -->
      <div class="pk-watermark">
        <a href="https://ploy-kong.vercel.app/" target="_blank">
          Made with <strong>PloyKong</strong>
          <Zap :size="12" class="icon-inline" />
        </a>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, nextTick } from "vue";
import { useRoute } from "vue-router";
import { publicAPI } from "@/api";
import type { Portfolio } from "@/types";
import { Lock, Search, Bot, X, Loader2, Zap } from "lucide-vue-next";
import { marked } from "marked";

const renderMarkdown = (text: string) => {
  return marked.parse(text);
};

function hexToRgb(hex: string) {
  if (!hex) return "0, 0, 0";
  const r = parseInt(hex.slice(1, 3), 16);
  const g = parseInt(hex.slice(3, 5), 16);
  const b = parseInt(hex.slice(5, 7), 16);
  return `${r}, ${g}, ${b}`;
}

// Public section renderers
import PublicHero from "@/components/portfolio/PublicHero.vue";
import PublicSkills from "@/components/portfolio/PublicSkills.vue";
import PublicProjects from "@/components/portfolio/PublicProjects.vue";
import PublicExperience from "@/components/portfolio/PublicExperience.vue";
import PublicContact from "@/components/portfolio/PublicContact.vue";
import PublicCustomText from "@/components/portfolio/PublicCustomText.vue";
import PublicCertificates from "@/components/portfolio/PublicCertificates.vue";
import PublicEducation from "@/components/portfolio/PublicEducation.vue";
import PublicStats from "@/components/portfolio/PublicStats.vue";
import PublicSocial from "@/components/portfolio/PublicSocial.vue";

const route = useRoute();
const slug = computed(() => route.params.slug as string);

const portfolio = ref<Portfolio | null>(null);
const loading = ref(true);
const notFound = ref(false);
const requiresPassword = ref(false);
const password = ref("");
const pwError = ref("");

const chatOpen = ref(false);
const chatInput = ref("");
const chatLoading = ref(false);
const chatMessages = ref<{ role: "user" | "assistant"; content: string }[]>([
  {
    role: "assistant",
    content:
      "สวัสดีครับ! ผมเป็น AI ตัวแทนของเจ้าของพอร์ตนี้ ถามอะไรก็ได้เลยครับ :)",
  },
]);
const chatMessagesRef = ref<HTMLElement | null>(null);
const sessionId = ref(Math.random().toString(36).slice(2));

const themeClass = computed(() => {
  const mode = portfolio.value?.theme?.mode || "system";
  return `theme-${mode}`;
});


const templateClass = computed(() => {
  const t = portfolio.value?.theme?.template || 'classic';
  return `tpl-${t}`;
});

const themeVars = computed(() => {
  const primary = portfolio.value?.theme?.primary_color || "#4F46E5";
  const secondary = portfolio.value?.theme?.secondary_color || "";
  const font = portfolio.value?.theme?.font || "Plus Jakarta Sans";
  const isLight = portfolio.value?.theme?.mode === "light";
  const glowHex = isLight ? "25" : "40";
  const vars: Record<string, string> = {
    "--primary": primary,
    "--primary-glow": `${primary}${glowHex}`,
    "--secondary": secondary ? secondary : `${primary}${glowHex}`,
    "--font-display": `"${font}", sans-serif`,
    "--font-body": `"${font}", sans-serif`
  };

  // Explicitly set text colors for portfolio isolation
  // Use system preference if mode is 'system'
  const isSystemDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
  const mode = portfolio.value?.theme?.mode || "system";
  const effectiveMode = (mode as any) === "system" ? (isSystemDark ? "dark" : "light") : mode;

  if (effectiveMode === "light") {
    vars["--text"] = "#000000";
    vars["--muted"] = "#6b7280";
  } else {
    vars["--text"] = "#ffffff";
    vars["--muted"] = "#94a3b8";
  }
  
  const bg = portfolio.value?.theme?.bg_color;
  const border = portfolio.value?.theme?.border_color;
  
  if (bg) {
    vars["--theme-bg"] = bg;
    vars["--bg-rgb"] = hexToRgb(bg);
  }

  if (border && border !== "#000000") {
    vars["--avatar-border"] = border;
  }

  return vars;
});

const visibleSections = computed(() =>
  (portfolio.value?.sections || [])
    .filter((s) => s.is_visible)
    .sort((a, b) => a.position - b.position),
);

const nonAIChatSections = computed(() =>
  visibleSections.value.filter((s) => s.type !== "ai_chat"),
);

// Flattened sections for CSS Grid, calculating left/right placement for the divider
const gridSections = computed(() => {
  const result: (typeof visibleSections.value[0] & { isRight?: boolean, isRowStart?: boolean, isPaired?: boolean })[] = [];
  const sections = nonAIChatSections.value;
  
  for (let i = 0; i < sections.length; i++) {
    const s = sections[i];
    const prev = i > 0 ? sections[i-1] : null;
    const next = i < sections.length - 1 ? sections[i+1] : null;
    
    if (s.column_span === 'half') {
      const isRight = prev?.column_span === 'half' && !result[i-1]?.isRight;
      const isRowStart = !isRight;
      const isPaired = isRowStart ? (next?.column_span === 'half') : true;
      
      result.push({ ...s, isRight, isRowStart, isPaired });
    } else {
      result.push({ ...s, isRight: false, isRowStart: true, isPaired: false });
    }
  }
  return result;
});

// Dynamic grouping for full-bleed sections (Hero) vs grid sections
const sectionGroups = computed(() => {
  const groups: { type: 'full' | 'grid'; sections: any[] }[] = [];
  const sections = gridSections.value;
  if (!sections.length) return groups;

  let i = 0;
  while (i < sections.length) {
    const s = sections[i];
    // Hero is full-bleed if NOT using firstjobber template
    const isHeroFullBleed = s.type === 'hero' && portfolio.value?.theme.template !== 'firstjobber';

    if (isHeroFullBleed) {
      groups.push({ type: 'full', sections: [s] });
      i++;
    } else {
      const gridItems: any[] = [];
      while (i < sections.length) {
        const next = sections[i];
        const nextIsFull = next.type === 'hero' && portfolio.value?.theme.template !== 'firstjobber';
        if (nextIsFull) break;
        gridItems.push(next);
        i++;
      }
      groups.push({ type: 'grid', sections: gridItems });
    }
  }
  return groups;
});

const hasAIChat = computed(() =>
  (portfolio.value?.sections || []).some((s) => s.type === "ai_chat" && s.is_visible),
);

const exampleQuestions = computed(() => {
  const aiChatSection = (portfolio.value?.sections || []).find((s) => s.type === "ai_chat");
  if (aiChatSection && aiChatSection.data && aiChatSection.data.example_questions) {
    return aiChatSection.data.example_questions.filter((q: string) => q.trim() !== "");
  }
  return [];
});

function getSectionComponent(type: string) {
  const map: Record<string, any> = {
    hero: PublicHero,
    skills: PublicSkills,
    projects: PublicProjects,
    experience: PublicExperience,
    contact: PublicContact,
    custom_text: PublicCustomText,
    certificates: PublicCertificates,
    education: PublicEducation,
    stats: PublicStats,
    social: PublicSocial,
  };
  return map[type] || PublicCustomText;
}

async function loadPortfolio(pw?: string) {
  loading.value = true;
  try {
    const { data } = await publicAPI.view(slug.value, pw);
    portfolio.value = data.data;
    // SEO Title update
    setTimeout(() => {
      const title = data.data.seo_title?.String || data.data.title;
      document.title = (typeof title === 'string' ? title : data.data.title) + " | PloyKong";
    }, 100);
    // Track view
    publicAPI.track(slug.value, "view", document.referrer || "direct");
  } catch (e: any) {
    if (e.response?.data?.requires_password) {
      requiresPassword.value = true;
    } else {
      notFound.value = true;
    }
  } finally {
    loading.value = false;
  }
}

async function loadWithPassword() {
  pwError.value = "";
  try {
    const { data } = await publicAPI.view(slug.value, password.value);
    portfolio.value = data.data;
    requiresPassword.value = false;
  } catch {
    pwError.value = "รหัสผ่านไม่ถูกต้อง";
  }
}

async function sendChat() {
  if (!chatInput.value.trim() || chatLoading.value) return;
  const msg = chatInput.value.trim();
  chatInput.value = "";
  chatMessages.value.push({ role: "user", content: msg });
  chatLoading.value = true;

  await nextTick();
  if (chatMessagesRef.value) {
    chatMessagesRef.value.scrollTop = chatMessagesRef.value.scrollHeight;
  }

  try {
    const { data } = await publicAPI.chat(slug.value, msg, sessionId.value);
    chatMessages.value.push({ role: "assistant", content: data.data.response });
    publicAPI.track(slug.value, "ai_chat");
  } catch {
    chatMessages.value.push({
      role: "assistant",
      content: "ขออภัย เกิดข้อผิดพลาด กรุณาลองใหม่",
    });
  } finally {
    chatLoading.value = false;
    await nextTick();
    if (chatMessagesRef.value) {
      chatMessagesRef.value.scrollTop = chatMessagesRef.value.scrollHeight;
    }
  }
}

onMounted(() => loadPortfolio());
</script>

<style scoped>
.public-portfolio {
  min-height: 100vh;
  position: relative;
  color: var(--text);
  background: var(--theme-bg, var(--bg)) fixed !important; /* Fixed background for premium feel */
  overflow-x: hidden;
  width: 100%;
}

/* Loading */
.pub-spinner {
  width: 40px;
  height: 40px;
  border: 3px solid var(--border);
  border-top-color: var(--purple);
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}
@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

/* Password wall, 404, Loading */
.pub-password-wall,
.pub-404,
.pub-loading {
  min-height: 100vh;
  gap: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  background: transparent !important; /* Let body background show through */
  color: var(--text);
}
:global(.theme-light) .pub-password-wall,
:global(.theme-light) .pub-404,
:global(.theme-light) .pub-loading {
  background: var(--bg) !important;
}
.pub-loading {
  gap: 16px;
  color: var(--muted);
}
.pub-password-wall {
  padding: 24px;
}

.pw-card {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 24px;
  padding: 44px;
  width: 100%;
  max-width: 380px;
  text-align: center;
  backdrop-filter: blur(20px);
}
.pw-icon {
  font-size: 48px;
  margin-bottom: 16px;
}
.pw-card h2 {
  font-size: 20px;
  margin-bottom: 20px;
}
.pw-error {
  color: var(--danger);
  font-size: 13px;
  margin-top: 8px;
}

.btn-sm {
  background: linear-gradient(to top left, var(--primary), var(--secondary));
  box-shadow: 0 8px 24px color-mix(in srgb, var(--primary), rgba(0, 0, 0, 0.2));
  border: 1px solid var(--primary);
}

.btn-sm:hover {
  background: color-mix(in srgb, var(--primary), var(--secondary));
  transform: translateY(-2px);
  box-shadow: 0 8px 24px color-mix(in srgb, var(--primary), rgba(0, 0, 0, 0.2));
  border: 1px solid var(--secondary);
}

/* 404 */
.pub-404 {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
}
.not-found-content {
  text-align: center;
}
.not-found-content h1 {
  font-size: 28px;
  margin-bottom: 8px;
}
.not-found-content p {
  color: var(--muted);
}

/* Sections */
/* --- CSS Grid Layout --- */
.section-group-wrap {
  width: 100%;
  height: 100%;
}
.full-bleed-section {
  width: 100%;
  height: 100%;
}
.sections-container {
  width: 100%;
  display: flex;
  justify-content: center;
  padding: 0 12px; /* Fixed small padding for mobile, grid max-width handles desktop */
}

.sections-grid {
  display: grid;
  grid-template-columns: minmax(0, 1fr) minmax(0, 1fr);
  gap: 20px;
  width: 100%;
  max-width: 1200px;
  padding-bottom: 40px;
  align-items: stretch;
}
.section-wrapper.span-full {
  grid-column: 1 / -1;
}
.section-wrapper.span-half {
  grid-column: span 1;
  padding: 0 !important;
}

/* Business template: 30/70 split instead of default 50/50 */
.tpl-business .sections-grid {
  grid-template-columns: 3fr 7fr !important;
  gap: 0 !important;
  max-width: 100% !important;
  padding-bottom: 0 !important;
}
@media (max-width: 768px) {
  .tpl-business .sections-grid {
    grid-template-columns: 1fr !important;
    gap: 16px !important;
  }
}

/* Animations */
.section-list-move,
.section-list-enter-active,
.section-list-leave-active {
  transition: all 0.4s ease;
}
.section-list-enter-from,
.section-list-leave-to {
  opacity: 0;
  transform: translateY(20px);
}
.section-list-leave-active {
  position: absolute;
  width: 100%;
  left: 0;
  box-sizing: border-box;
  opacity: 0;
}

/* Divider Implementation */
.sections-grid.has-divider .section-wrapper.span-half.is-right {
  position: relative;
}
.sections-grid.has-divider .section-wrapper.span-half.is-right::before {
  content: "";
  position: absolute;
  top: 10%;
  bottom: 10%;
  left: -10px; /* Offset to exactly half of the 20px grid gap */
  width: 1px;
  background-color: var(--border);
}

@media (max-width: 640px) {
  .sections-grid {
    grid-template-columns: minmax(0, 1fr);
  }
  .section-wrapper.span-half {
    grid-column: 1 / -1;
  }
  .sections-grid.has-divider .section-wrapper.span-half.is-right {
    padding-left: 0;
  }
  .sections-grid.has-divider .section-wrapper.span-half.is-right::before {
    display: none;
  }
}

.public-portfolio {
  padding-top: 3px;
  padding-bottom: 3px;
}

/* AI Chat FAB */
.chat-fab {
  position: fixed;
  bottom: 28px;
  right: 28px;
  width: 56px;
  height: 56px;
  border-radius: 50%;
  color: rgba(255, 255, 255, 1);
  background: linear-gradient(135deg, var(--primary), var(--secondary));
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 22px;
  cursor: pointer;
  z-index: 100;
  box-shadow: 0 8px 24px color-mix(in srgb, var(--primary), rgba(0, 0, 0, 0.2));
  transition: transform 0.2s;
}
.chat-fab:hover {
  transform: scale(1.1);
}

/* Chat Panel */
.chat-panel {
  position: fixed;
  bottom: 96px;
  right: 28px;
  width: 340px;
  max-height: 480px;
  background: color-mix(in srgb, var(--primary), rgba(168, 105, 105, 0.6));
  border: 1px solid var(--secondary);
  border-radius: 20px;
  z-index: 100;
  display: flex;
  flex-direction: column;
  box-shadow: 0 24px 60px rgba(0, 0, 0, 0.5);
  backdrop-filter: blur(20px);
}
:global(.theme-dark) .chat-panel {
  background: rgba(10, 14, 30, 1);
}
:global(.theme-light) .chat-panel {
  box-shadow: 0 24px 60px rgba(255, 255, 255, 0.98);
}

.chat-header {
  padding: 16px;
  border-bottom: 1px solid var(--border);
}
.chat-header span:first-child {
  font-size: 14px;
  font-weight: 700;
  display: block;
  color: rgba(var(--text-rgb), 1);
}
.chat-sub {
  font-size: 11px;
  color: rgba(var(--text-rgb), 1);
}

:global(.theme-dark) .chat-sub {
  --text-rgb: 255, 255, 255;
}

:global(.theme-light) .chat-sub {
  --text-rgb: 0, 0, 0;
} 
.chat-messages {
  flex: 1;
  overflow-y: auto;
  padding: 14px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  min-height: 0;
}
.chat-msg {
  display: flex;
}
.chat-msg.user {
  justify-content: flex-end;
}
.chat-bubble {
  max-width: 100%;
  padding: 10px 14px;
  border-radius: 14px;
  font-size: 13px;
  line-height: 1.5;
  word-wrap: break-word;
}
.chat-bubble :first-child { margin-top: 0; }
.chat-bubble :last-child { margin-bottom: 0; }
.chat-bubble p { margin-bottom: 4px; }
.chat-bubble strong { font-weight: 700; }
.chat-msg.user .chat-bubble {
  background-color: var(--primary);
  color: #fff;
}
.chat-msg.assistant .chat-bubble {
  background-color: color-mix(in srgb, var(--primary), rgba(0, 0, 0, 0.6));
  color: #fff;
}
.chat-typing {
  color: var(--muted);
  font-style: italic;
}

.chat-suggestions {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  padding: 0 14px 12px;
}
.suggestion-btn {
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid var(--border);
  border-radius: 12px;
  padding: 4px 10px;
  font-size: 11px;
  color: var(--text);
  cursor: pointer;
  transition: all 0.2s;
  text-align: left;
}
.suggestion-btn:hover {
  background: var(--primary);
  color: #fff;
  border-color: var(--primary);
}

.chat-input-row {
  padding: 12px;
  border-top: 1px solid var(--border);
  display: flex;
  gap: 8px;
}
.chat-input-row .form-input {
  flex: 1;
  padding: 8px 12px;
  font-size: 13px;
}

/* Chat slide transition */
.chat-slide-enter-active,
.chat-slide-leave-active {
  transition: all 0.25s ease;
}
.chat-slide-enter-from,
.chat-slide-leave-to {
  opacity: 0;
  transform: translateY(16px) scale(0.96);
}

/* Watermark */
.pk-watermark {
  text-align: center;
  padding: 40px 24px 60px 24px;
}
.pk-watermark a {
  font-size: 12px;
  color: var(--watercolor);
  text-decoration: none;
}
.pk-watermark a:hover {
  color: var(--primary);
}
.pk-watermark strong {
  color: var(--primary);
}
.pk-watermark a .theme-light {
  --watercolor: #575757;
}

:global(.theme-dark) .pk-watermark a {
  --watercolor: #7b7a7a;
}
/* Hide Builder Placeholder hints on Public View */
.public-portfolio :deep(.generic-hint) {
  display: none !important;
}

.icon-inline {
  vertical-align: text-bottom;
  margin-right: 2px;
}
@keyframes spin {
  100% {
    transform: rotate(360deg);
  }
}
.spin {
  animation: spin 1s linear infinite;
}

@media (max-width: 480px) {
  .chat-panel {
    right: 12px;
    left: 12px;
    width: auto;
    bottom: 80px; /* Bring it down a bit to not feel too high */
  }
  .chat-fab {
    bottom: 16px;
    right: 16px;
  }
  /* Fix send button on mobile */
  .chat-input-row {
    padding: 10px;
    gap: 8px;
  }
  .chat-input-row .form-input {
    flex: 1;
    min-width: 0; /* Important for flex-grow on mobile */
    height: 40px;
  }
  .chat-input-row .btn {
    width: auto !important;
    padding: 0 20px !important;
    height: 40px !important;
    flex-shrink: 0;
    margin: 0 !important;
  }
}
</style>
