<template>
  <div v-if="parsedData" class="analysis-result">
    <!-- Score Header -->
    <div class="score-section">
      <div class="score-circle-container">
        <svg viewBox="0 0 36 36" class="circular-chart">
          <path class="circle-bg"
            d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831"
          />
          <path class="circle"
            :stroke-dasharray="`${score}, 100`"
            d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831"
          />
          <text x="18" y="20.35" class="percentage">{{ score }}</text>
        </svg>
      </div>
      <div class="score-info">
        <div class="score-label">Portfolio Score</div>
        <div class="score-status">{{ statusText }}</div>
      </div>
    </div>

    <!-- Details Sections -->
    <div class="details-grid">
      <!-- Strengths -->
      <section v-if="parsedData.strengths?.length" class="detail-block">
        <h3><CheckCircle2 :size="14" class="icon-pos" /> จุดแข็ง (Strengths)</h3>
        <ul>
          <li v-for="(s, i) in parsedData.strengths" :key="i">{{ s }}</li>
        </ul>
      </section>

      <!-- Improvements -->
      <section v-if="parsedData.improvements?.length" class="detail-block">
        <h3><AlertCircle :size="14" class="icon-pos" /> สิ่งที่ควรปรับปรุง (Improvements)</h3>
        <ul>
          <li v-for="(s, i) in parsedData.improvements" :key="i">{{ s }}</li>
        </ul>
      </section>

      <!-- Missing Sections -->
      <section v-if="parsedData.missing_sections?.length" class="detail-block">
        <h3><HelpCircle :size="14" class="icon-pos" /> หัวข้อที่ขาดหาย (Missing)</h3>
        <ul>
          <li v-for="(s, i) in parsedData.missing_sections" :key="i">{{ s }}</li>
        </ul>
      </section>
    </div>
  </div>
  <div v-else class="error-msg">
    {{ result }}
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { CheckCircle2, AlertCircle, HelpCircle } from 'lucide-vue-next';

const props = defineProps<{
  result: string | null;
}>();

const parsedData = computed(() => {
  if (!props.result) return null;
  try {
    // If it's already an object, return it. If it's a string, try to parse it.
    // Sometimes AI returns it as a string instead of JSON if it includes other text.
    // But our API usually returns the string analysis.
    let cleanStr = props.result.trim();
    // Strip markdown code blocks if present
    if (cleanStr.startsWith('```json')) {
      cleanStr = cleanStr.replace(/```json|```/g, '').trim();
    } else if (cleanStr.startsWith('```')) {
      cleanStr = cleanStr.replace(/```/g, '').trim();
    }
    return JSON.parse(cleanStr);
  } catch (e) {
    console.warn("Failed to parse AI result JSON", e);
    return null;
  }
});

const score = computed(() => parsedData.value?.score || 0);

const statusText = computed(() => {
  const s = score.value;
  if (s >= 80) return "🎯 ยอดเยี่ยมมาก!";
  if (s >= 50) return "💪 ดีพอใช้ ปรับปรุงอีกนิด";
  return "🚀 เริ่มต้นได้ดี เร่งเครื่องอีกหน่อย";
});

const statusColor = computed(() => {
  const s = score.value;
  if (s >= 80) return "var(--green)";
  if (s >= 50) return "var(--warning)";
  return "var(--danger)";
});
</script>

<style scoped>
.analysis-result {
  margin-top: 16px;
  background: var(--bg2);
  border-radius: 12px;
  padding: 20px;
  animation: slideIn 0.3s ease-out;
}

.score-section {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 24px;
}

.score-circle-container {
  width: 64px;
  height: 64px;
}

.circular-chart {
  display: block;
  margin: 10px auto;
  max-width: 100%;
  max-height: 250px;
}

.circle-bg {
  fill: none;
  stroke: var(--border);
  stroke-width: 2.8;
}

.circle {
  fill: none;
  stroke-width: 2.8;
  stroke-linecap: round;
  transition: stroke-dasharray 1s ease 0s;
  stroke: var(--indigo);
}

.percentage {
  fill: var(--text);
  font-family: var(--font-display);
  font-size: 10px;
  text-anchor: middle;
  font-weight: 800;
}

.score-info {
  display: flex;
  flex-direction: column;
}

.score-label {
  font-size: 11px;
  color: var(--muted);
  text-transform: uppercase;
  font-weight: 600;
  letter-spacing: 0.5px;
}

.score-status {
  font-size: 16px;
  font-weight: 800;
  color: var(--text);
}

.details-grid {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.detail-block h3 {
  font-size: 12px;
  font-weight: 800;
  margin-bottom: 10px;
  display: flex;
  align-items: center;
  gap: 6px;
  color: var(--text);
  text-transform: uppercase;
}

.icon-pos {
  color: var(--indigo);
}

.detail-block ul {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.detail-block li {
  font-size: 13px;
  color: var(--muted);
  position: relative;
  padding-left: 14px;
  line-height: 1.4;
}

.detail-block li::before {
  content: "•";
  position: absolute;
  left: 0;
  color: var(--indigo);
}

.error-msg {
  font-size: 13px;
  color: var(--danger);
  padding: 12px;
}

@keyframes slideIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>
