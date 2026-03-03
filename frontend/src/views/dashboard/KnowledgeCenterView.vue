<template>
  <div class="dashboard-layout">
    <AppSidebar />

    <!-- Main Content -->
    <main class="dashboard-main">
      <div class="dash-header">
        <div>
          <h1 class="dash-greeting">AI Knowledge Center <Sparkles :size="30" class="icon-inline" :style="{paddingTop: '5px'}"/></h1>
          <p class="dash-sub">ตอบคำถามที่ AI ของคุณยังไม่รู้ เพื่อฉลาดขึ้นในทุกๆ พอร์ต</p>
        </div>
        <div class="header-actions">
          <button @click="showTips = !showTips" class="btn btn-outline btn-sm">
            <Lightbulb :size="16" /> {{ showTips ? 'Hide Tips' : 'Training Tips' }}
          </button>
          <button @click="fetchGaps" class="btn btn-secondary btn-sm" :disabled="loading">
             <RefreshCw :size="16" :class="{ spin: loading }" /> Refresh
          </button>
        </div>
      </div>

      <!-- Training Tips Card -->
      <Transition name="fade">
        <div v-if="showTips" class="tips-card">
          <h3><Sparkles :size="18" /> ทริคการสอน AI ให้เก่งขึ้น</h3>
          <div class="tips-grid">
            <div class="tip-item">
              <div class="tip-icon">🎯</div>
              <h4>ตอบเป็นกลาง (Neutral)</h4>
              <p>ถ้าคำถามใช้ได้กับทุกพอร์ต ให้ใช้คำว่า <strong>"ในภาพรวม..."</strong> AI จะเข้าใจว่าเป็นข้อมูลพื้นฐานของคุณครับ</p>
            </div>
            <div class="tip-item">
              <div class="tip-icon">🏷️</div>
              <h4>แยกตามประเภท</h4>
              <p>ระบุประเภทข้อมูล เช่น <strong>[Skill]</strong> หรือ <strong>[Experience]</strong> นำหน้าคำตอบเพื่อให้ AI จัดหมวดหมู่ได้ดีขึ้น</p>
            </div>
            <div class="tip-item">
              <div class="tip-icon">🤖</div>
              <h4>สอนแบบเจาะจง</h4>
              <p>ถ้าต้องการให้ AI ตอบเฉพาะพอร์ตนี้เท่านั้น ให้เริ่มด้วย <strong>"สำหรับโปรเจกต์นี้..."</strong> ครับ</p>
            </div>
            <div class="tip-item">
              <div class="tip-icon">🚫</div>
              <h4>เจอคำถามที่ไม่อยากตอบ</h4>
              <p>ขอไม่ตอบคำถามนี้<strong> "ด้วยเหตุผลส่วนตัว...." </strong>ครับ</p>
            </div>
          </div>
        </div>
      </Transition>

      <div v-if="loading" class="loading-state">
        <div v-for="i in 3" :key="i" class="skeleton gap-skeleton"></div>
      </div>

      <div v-else-if="gaps.length === 0" class="empty-state">
        <div class="empty-icon">✨</div>
        <h3>ยอดเยี่ยมมาก! AI ของคุณตอบได้ทุกคำถาม</h3>
        <p>ยังไม่มีคำถามที่แสดงเป็น Knowledge Gap ในขณะนี้</p>
      </div>

      <div v-else class="gaps-list">
        <div v-for="(gap, i) in gaps" :key="i" class="gap-card">
          <div class="gap-card-header">
            <div class="portfolio-badge">
              <Layers :size="12" /> {{ gap.portfolio_title }}
            </div>
            <span class="gap-time">{{ formatDate(gap.time) }}</span>
          </div>
          
          <div class="gap-content">
            <div class="gap-q">
              <strong>Question:</strong>
              <p>{{ gap.question }}</p>
            </div>
            
            <div class="gap-a-input">
              <label>เพิ่มข้อมูลสำหรับตอบคำถามนี้:</label>
              <div class="input-row">
                <textarea 
                  v-model="gap.answer" 
                  class="form-input" 
                  rows="2"
                  placeholder="เช่น 'ฉันถนัด Go และ Vue.js เป็นพิเศษ...'"
                ></textarea>
                <button 
                  @click="learn(gap)" 
                  class="btn btn-primary" 
                  :disabled="!gap.answer || gap.saving"
                >
                  <template v-if="gap.saving"><Loader2 :size="18" class="spin" /></template>
                  <template v-else>Learn & Solve</template>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { portfolioAPI, sectionAPI } from '@/api';
import { useAuthStore } from '@/stores/auth';
import { useThemeStore } from '@/stores/theme';
import { usePortfolioStore } from '@/stores/portfolio';
import { useRouter } from 'vue-router';
import AppSidebar from '@/components/layout/AppSidebar.vue';
import { 
  BarChart2, Sparkles, Plus, Settings, LogOut, 
  Layers, RefreshCw, Loader2, CheckCircle, Lightbulb
} from 'lucide-vue-next';
import dayjs from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime';
import utc from 'dayjs/plugin/utc';
import { toastSuccess, toastError } from '@/utils/alert';

dayjs.extend(relativeTime);
dayjs.extend(utc);

const authStore = useAuthStore();
const themeStore = useThemeStore();
const portfolioStore = usePortfolioStore();
const router = useRouter();

const gaps = ref<any[]>([]);
const loading = ref(false);
const showTips = ref(true);

const userInitial = computed(() => authStore.user?.name?.[0]?.toUpperCase() || 'U');

async function fetchGaps() {
  loading.value = true;
  try {
    const { data } = await portfolioAPI.listGaps();
    gaps.value = data.data.map((g: any) => ({ ...g, answer: '', saving: false }));
    // Update global gap count
    portfolioStore.fetchGapCount();
  } catch (err) {
    console.error(err);
    toastError("ไม่สามารถโหลดข้อมูลได้");
  } finally {
    loading.value = false;
  }
}

async function learn(gap: any) {
  gap.saving = true;
  try {
    // 1. Find the AI Chat section ID for this portfolio
    const sectionsRes = await sectionAPI.list(gap.portfolio_id);
    const aiSection = sectionsRes.data.data.find((s: any) => s.type === 'ai_chat');
    
    if (!aiSection) {
      toastError("ไม่พบ AI Chat Section ในพอร์ตนี้");
      return;
    }

    // 2. Append training data to prompt_hint
    const currentData = aiSection.data || {};
    const trainingEntry = `\nQ: ${gap.question}\nA: ${gap.answer}`;
    if (!currentData.prompt_hint) currentData.prompt_hint = "";
    currentData.prompt_hint += trainingEntry;

    // 3. Save update
    await sectionAPI.update(aiSection.id, { data: currentData });
    
    toastSuccess("AI เรียนรู้ข้อมูลใหม่เรียบร้อยแล้ว!");
    
    // Remove from local list
    gaps.value = gaps.value.filter(g => g !== gap);
    // Update global gap count
    portfolioStore.gapCount--;
  } catch (err) {
    console.error(err);
    toastError("เกิดข้อผิดพลาดในการบันทึกข้อมูล");
  } finally {
    gap.saving = false;
  }
}

function formatDate(date: string) {
  if (!date) return "";
  // Assume backend sends UTC. Convert to local for display.
  // If backend uses default MySQL timestamp, it might not have the 'Z'.
  // dayjs.utc(date) will force it to be treated as UTC.
  return dayjs.utc(date).local().fromNow();
}

async function handleLogout() {
  await authStore.logout();
  router.push('/login');
}

onMounted(fetchGaps);
</script>

<style scoped>
.user-name { font-size:12px; font-weight:700; }
.user-plan { font-size:10px; color:var(--neon-purple); text-transform:uppercase; }

/* Content */
.dashboard-main { padding: 40px; }
.dash-header { display:flex; justify-content:space-between; align-items:flex-start; margin-bottom:32px; }
.header-actions { display:flex; gap:12px; }

.tips-card {
  background: linear-gradient(135deg, rgba(79, 70, 229, 0.08), rgba(168, 85, 247, 0.08));
  border: 1px dashed var(--indigo);
  border-radius: 16px;
  padding: 24px;
  margin-bottom: 32px;
}
.tips-card h3 { display:flex; align-items:center; gap:8px; font-size:16px; margin-bottom:16px; color:var(--indigo); }
.tips-grid { display:grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap:20px; }
.tip-item h4 { font-size:14px; margin: 8px 0 4px 0; color:var(--text); }
.tip-item p { font-size:12px; color:var(--muted); line-height:1.5; }
.tip-icon { font-size:20px; }

.fade-enter-active, .fade-leave-active { transition: all 0.3s ease; max-height: 400px; opacity: 1; }
.fade-enter-from, .fade-leave-to { max-height: 0; opacity: 0; padding: 0; margin: 0; overflow: hidden; }

.dash-greeting { font-size:28px; font-weight:800; margin-bottom:4px; }
.dash-sub { color:var(--muted); font-size:15px; }

.gaps-list { display:flex; flex-direction:column; gap:20px; }
.gap-card {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 16px;
  padding: 20px;
  transition: transform 0.2s;
}
.gap-card:hover { transform: translateY(-2px); box-shadow: 0 8px 24px rgba(0,0,0,0.1); }

.gap-card-header {
  display:flex; justify-content:space-between; align-items:center; margin-bottom:16px;
}
.portfolio-badge {
  display:inline-flex; align-items:center; gap:6px;
  background: var(--bg2); color: var(--indigo);
  padding: 4px 10px; border-radius:8px; font-size:12px; font-weight:700;
}
.gap-time { font-size:12px; color:var(--muted); }

.gap-q { margin-bottom:20px; }
.gap-q strong { display:block; font-size:12px; text-transform:uppercase; color:var(--muted); margin-bottom:6px; }
.gap-q p { font-size:16px; color:var(--text); font-weight:500; }

.gap-a-input label { display:block; font-size:12px; font-weight:700; margin-bottom:8px; color:var(--text); }
.input-row { display:flex; gap:12px; align-items: flex-end; }
.input-row textarea { flex:1; border-radius:12px; padding:12px; resize: none; }
.input-row .btn { height: 48px; min-width: 140px; }

.empty-state {
  text-align:center; padding: 60px 0; color:var(--muted);
}
.empty-icon { font-size: 48px; margin-bottom: 20px; }
.empty-state h3 { color:var(--text); margin-bottom:8px; }

.skeleton { background: var(--bg2); border-radius:12px; }
.gap-skeleton { height: 180px; margin-bottom: 20px; animation: pulse 1.5s infinite ease-in-out; }

@keyframes pulse {
  0% { opacity: 0.6; }
  50% { opacity: 1; }
  100% { opacity: 0.6; }
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}
.spin { animation: spin 1s linear infinite; }

@media (max-width: 768px) {
  .dashboard-main { padding: 20px; }
  .dash-header { flex-direction: column; gap: 16px; }
  .header-actions { width: 100%; display: flex; gap: 8px; }
  .header-actions .btn { flex: 1; padding: 10px; font-size: 13px; }
  .input-row { flex-direction:column; align-items: stretch; }
  .input-row .btn { width: 100%; }
  .tips-card { padding: 16px; }
  .tips-grid { grid-template-columns: 1fr; }
}
</style>
