<template>
  <section class="demo-section" id="demo">
    <div class="section-wrap">
      <div class="section-header">
        <div class="section-tag">INTERACTIVE DEMO</div>
        <h2 class="section-h2">เห็นภาพจริงใน 2 นาที</h2>
        <p class="section-sub">ทดลองใช้งานระบบ Builder อัจฉริยะของเรา — ตั้งแต่เลือกเทมเพลตจนถึงการ Publish</p>
      </div>

      <div class="demo-shell-container">
        <div class="demo-shell" ref="demoShell">
          <!-- TOP BAR -->
          <div class="topbar">
            <div class="tb-logo">
              <div class="tb-logomark">
                <img src="/favicon.png" alt="Logo" class="demo-logo-img" />
              </div>
              <div class="tb-logoname">Ploy<em>Kong</em></div>
            </div>

            <div class="tb-steps" id="tbSteps">
              <div v-for="(step, i) in steps" :key="i" 
                   class="tb-step" 
                   :class="{ active: currentScene === i, done: currentScene > i }"
                   @click="goToScene(i)">
                <div class="tb-step-dot"></div>{{ step.name }}
              </div>
            </div>

            <div class="tb-right">
              <div class="tb-badge"><div class="live-dot"></div>Live Demo</div>
            </div>
          </div>

          <!-- MAIN STAGE -->
          <div class="stage">
            <div class="stage-bg">
              <div class="orb orb1"></div>
              <div class="orb orb2"></div>
              <div class="orb orb3"></div>
              <div class="grid-bg"></div>
            </div>

            <!-- SCENES -->
            <div class="scenes-container">
              <!-- SCENE 1: Template Picker -->
              <div class="scene" :class="{ active: currentScene === 0 }" id="demo-scene0">
                <div class="s1-wrap">
                  <div class="s1-left fly-in-l" :class="{ show: sceneAnims[0] }">
                    <div class="s1-eyebrow">STEP 01 — GET STARTED</div>
                    <h2 class="s1-h">เลือก Template<br/><span class="grad-text">ตามสายงาน</span></h2>
                    <p class="s1-p">เลือกจาก 48+ template ที่ออกแบบมาสำหรับแต่ละสายงานโดยเฉพาะ — แค่คลิกเดียว</p>
                    <div class="role-tags">
                      <div v-for="(role, i) in roles" :key="i" 
                           class="role-tag" 
                           :class="{ sel: currentRole === i }"
                           @click="selectRole(i)">
                        {{ role }}
                      </div>
                    </div>
                    <div class="tag-pill tag-accent">✦ 48 Templates Available</div>
                  </div>
                  <div class="s1-right fly-in-r" :class="{ show: sceneAnims[0] }">
                    <div v-for="(tpl, i) in templates" :key="i" 
                         class="tpl-card" 
                         :class="{ sel: currentTpl === i }"
                         @click="selectTpl(i)"
                         :id="'tpl' + i">
                      <div class="tpl-preview" :class="tpl.class">
                        <div class="m-avatar" :style="tpl.avatarBg ? {background: tpl.avatarBg} : {}"></div>
                        <div class="m-line hl" :style="tpl.hlBg ? {background: tpl.hlBg} : {}"></div>
                        <div class="m-line s"></div>
                        <div v-if="tpl.chips" class="m-chips">
                          <div v-for="c in tpl.chips" :key="c" class="m-chip">{{ c }}</div>
                        </div>
                        <div class="tpl-check">✓</div>
                      </div>
                      <div class="tpl-label">{{ tpl.name }}</div>
                    </div>
                  </div>
                </div>
              </div>

              <!-- SCENE 2: Builder + AI -->
              <div class="scene" :class="{ active: currentScene === 1 }" id="demo-scene1">
                <div class="s2-wrap">
                  <div class="s2-sidebar fly-in-l" :class="{ show: sceneAnims[1] }">
                    <div class="s2-sidebar-h">📦 Blocks</div>
                    <div class="block-list">
                      <div v-for="b in blocks" :key="b.id" class="block-item" :id="'block-' + b.id" :class="{ 'dragging-from': draggingBlock === b.id }">
                        <span class="block-icon">{{ b.icon }}</span>{{ b.name }}<span class="block-handle">⠿</span>
                      </div>
                    </div>
                  </div>
                  <div class="s2-canvas fly-in" :class="{ show: sceneAnims[1] }">
                    <div class="s2-canvas-h">
                      <span>🖥️</span>Canvas — pk.ploykong.com
                      <span class="tag-pill tag-green" style="margin-left:auto">Live Preview</span>
                    </div>
                    <div class="canvas-body" id="canvasBody">
                      <div v-for="s in canvasSections" :key="s.id" class="canvas-section" :class="{ 'new-section': s.isNew, 'dropping': s.dropping }">
                        <div class="cs-header"><span class="cs-drag-ico">⠿</span><span>{{ s.title }}</span></div>
                        <div class="cs-content" v-html="s.content"></div>
                      </div>
                      <div class="drop-zone" :class="{ highlight: dropHighlight }" id="dropZone">
                        ⊕ ลาก Block มาวางที่นี่
                      </div>
                    </div>
                  </div>
                  <div class="s2-ai fly-in-r" :class="{ show: sceneAnims[1] }">
                    <div class="ai-panel-h"><span class="ai-spark">✨</span>AI Assistant</div>
                    <div class="ai-body">
                      <div class="ai-score-box">
                        <div class="ai-score-h">📊 Resume Score</div>
                        <div class="ai-score-num">{{ aiScore }}</div>
                        <div class="ai-score-bar"><div class="ai-score-fill" :style="{ width: aiScore + '%' }"></div></div>
                        <div class="ai-score-label">ดี — เพิ่ม Experience เพื่อเพิ่มคะแนน</div>
                      </div>
                      <div class="ai-improve-box">
                        <div class="ai-improve-h"><span>✍️</span>AI Improve Text</div>
                        <div class="ai-text-in">worked on backend projects using go...</div>
                        <div class="ai-arrow">↓ AI กำลังปรับปรุง ✨</div>
                        <div class="ai-text-out">
                          <span>{{ aiTypedText }}</span><span class="typing-cursor" v-if="isTyping"></span>
                        </div>
                      </div>
                      <div class="ai-action"><span>✨</span> Apply AI Suggestion</div>
                    </div>
                  </div>
                </div>
              </div>

              <!-- SCENE 3: Analytics -->
              <div class="scene" :class="{ active: currentScene === 2 }" id="demo-scene2">
                <div class="s3-wrap">
                  <div class="s3-left">
                    <div class="stat-cards">
                      <div v-for="(s, i) in stats" :key="i" class="stat-card fly-in" :class="{ show: sceneAnims[2] }">
                        <div class="sc-label">{{ s.label }}</div>
                        <div class="sc-num c-accent">{{ s.val }}</div>
                        <div class="sc-delta">{{ s.delta }}</div>
                      </div>
                    </div>
                    <div class="chart-card fly-in" :class="{ show: sceneAnims[2] }">
                      <div class="chart-h">
                        <div class="chart-title">📈 Visitor Trend (7 วัน)</div>
                        <div class="chart-legend">
                          <div class="chart-leg"><div class="chart-leg-dot" style="background:var(--accent)"></div>Views</div>
                          <div class="chart-leg"><div class="chart-leg-dot" style="background:var(--cyan)"></div>Clicks</div>
                        </div>
                      </div>
                      <div class="chart-area" ref="chartArea">
                        <svg class="chart-svg" viewBox="0 0 600 160" preserveAspectRatio="none" ref="chartSvg"></svg>
                      </div>
                    </div>
                  </div>
                  <div class="s3-right">
                    <div class="source-card fly-in-r" :class="{ show: sceneAnims[2] }">
                      <div class="sc-head">🔗 Traffic Sources</div>
                      <div v-for="src in trafficSources" :key="src.name" class="source-row">
                        <div class="source-icon" :style="{ background: src.bg }">{{ src.icon }}</div>
                        <div class="source-name">{{ src.name }}</div>
                        <div class="source-bar-wrap"><div class="source-bar" :style="{ width: src.width + '%', background: src.grad }"></div></div>
                        <div class="source-pct" :style="{ color: src.color }">{{ src.width }}%</div>
                      </div>
                    </div>
                    <div class="map-card fly-in-r" :class="{ show: sceneAnims[2] }">
                      <div class="map-head">🌍 Visitor Countries</div>
                      <div class="map-img">
                        <span style="z-index:1">🗺️ World Map</span>
                        <div class="map-dot" style="top:30%;left:72%"></div>
                        <div class="map-dot" style="top:40%;left:78%;animation-delay:0.5s;background:var(--cyan)"></div>
                        <div class="map-dot" style="top:35%;left:50%;animation-delay:1s;background:var(--pink)"></div>
                      </div>
                      <div class="countries-list">
                        <div v-for="c in countries" :key="c.name" class="country-row">
                          <span class="country-flag">{{ c.flag }}</span>
                          <span class="country-name">{{ c.name }}</span>
                          <span class="country-num" :style="{ color: 'var(--accent2)' }">{{ c.num }}</span>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <!-- SCENE 4: Publish -->
              <div class="scene" :class="{ active: currentScene === 3 }" id="demo-scene3">
                <div class="s4-wrap">
                  <div class="s4-left">
                    <div class="s4-eyebrow fly-in" :class="{ show: sceneAnims[3] }">✅ STEP 04 — PUBLISH</div>
                    <h2 class="s4-h fly-in" :class="{ show: sceneAnims[3] }">พร้อม<span class="grad-text">แล้ว!</span><br/>แชร์ได้เลย 🚀</h2>
                    <p class="s4-p fly-in" :class="{ show: sceneAnims[3] }">พอร์ตของคุณ Live แล้ว — URL สวยงาม อัปเดตได้ตลอดเวลา โดยไม่ต้อง Deploy ใหม่</p>
                    <div class="url-pill fly-in" :class="{ show: sceneAnims[3] }">
                      <span class="url-lock">🔒</span><span class="url-domain">ploykong.com/</span><span class="url-slug">pk</span>
                    </div>
                    <div class="publish-btn fly-in active-btn" :class="{ show: sceneAnims[3] }" @click="triggerPublish">🚀 Published!</div>
                    <div class="share-options fly-in" :class="{ show: sceneAnims[3] }">
                      <div v-for="s in shareOptions" :key="s.label" class="share-row">
                        <div class="share-ico">{{ s.icon }}</div><span class="share-label">{{ s.label }}</span><span class="share-val">{{ s.val }}</span>
                      </div>
                    </div>
                  </div>
                  <div class="portfolio-preview fly-in-r" :class="{ show: sceneAnims[3] }">
                    <div class="pp-topbar">
                      <div class="pp-dots"><div class="pp-dot" style="background:#FF5F57"></div><div class="pp-dot" style="background:#FFBD2E"></div><div class="pp-dot" style="background:#28CA41"></div></div>
                      <div style="font-family:'DM Mono',monospace;font-size:10px;color:var(--text2);margin:0 auto">🔒 ploykong.com/pk</div>
                    </div>
                    <div class="pp-hero">
                      <div class="pp-avatar-row">
                        <div class="pp-avatar">PK</div>
                        <div><div class="pp-name">Ploy Kong</div><div class="pp-role">Full-Stack Developer</div></div>
                        <div class="pp-hire-btn">Hire Me ✨</div>
                      </div>
                      <div class="pp-tagline">"Passionate about building high-performance web applications with Go & Vue.js. Open to exciting opportunities."</div>
                      <div style="display:flex;gap:8px;flex-wrap:wrap;">
                        <div class="tag-pill tag-accent">Go</div><div class="tag-pill tag-cyan">Vue.js</div><div class="tag-pill tag-pink">Flutter</div><div class="tag-pill tag-green">Docker</div>
                      </div>
                    </div>
                    <div class="pp-skills">
                      <div class="pp-skill-title">SKILLS</div>
                      <div class="skill-row-p"><span class="sk-name">Go / Fiber</span><div class="sk-track"><div class="sk-fill" style="width:92%"></div></div></div>
                      <div class="skill-row-p"><span class="sk-name">Vue.js 3</span><div class="sk-track"><div class="sk-fill" style="width:85%"></div></div></div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- BOTTOM CONTROLS -->
          <div class="controls">
            <div class="ctrl-info">
              <div class="ctrl-step-name">{{ steps[currentScene].name }}</div>
              <div class="ctrl-step-desc">{{ steps[currentScene].desc }}</div>
            </div>
            <div class="progress-track">
              <div class="progress-fill" :style="{ width: progressValue + '%' }"></div>
            </div>
            <div class="ctrl-btns">
              <select class="speed-select" v-model="speedMult">
                <option :value="1">1× Speed</option>
                <option :value="1.5">1.5× Speed</option>
                <option :value="2">2× Speed</option>
                <option :value="0.5">0.5× Speed</option>
              </select>
              <button class="ctrl-btn" @click="prevScene">
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="15 18 9 12 15 6"/></svg>
                Prev
              </button>
              <button class="ctrl-btn" @click="togglePlay">
                <svg v-if="isPlaying" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><rect x="6" y="4" width="4" height="16"/><rect x="14" y="4" width="4" height="16"/></svg>
                <svg v-else width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polygon points="5 3 19 12 5 21 5 3"/></svg>
                {{ isPlaying ? 'Pause' : 'Play' }}
              </button>
              <button class="ctrl-btn primary" @click="nextScene">
                Next
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="9 18 15 12 9 6"/></svg>
              </button>
            </div>
          </div>

          <!-- Extra layers -->
          <div class="toast-area">
            <TransitionGroup name="toast">
              <div v-for="t in toasts" :key="t.id" class="toast" :style="{ borderLeftColor: t.color, borderLeftWidth: '3px' }">
                <span class="toast-ico">{{ t.icon }}</span><span>{{ t.text }}</span>
              </div>
            </TransitionGroup>
          </div>
          <div class="confetti-wrap" v-if="showConfetti">
            <div v-for="i in 80" :key="i" class="confetti-piece" :style="getConfettiStyle()"></div>
          </div>
          <div class="demo-cursor" :style="{ left: cursorPos.x + 'px', top: cursorPos.y + 'px' }"></div>
          <div class="scene-transition" :class="{ active: isTransitioning }"></div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watch } from 'vue'

const currentScene = ref(0)
const isPlaying = ref(true)
const speedMult = ref(2)
const progressValue = ref(0)
const isTransitioning = ref(false)
const sceneAnims = ref([false, false, false, false])

const steps = [
  { name: 'เลือก Template', desc: 'คลิก template ที่ชอบ หรือรอ auto-play' },
  { name: 'Drag & Drop + AI', desc: 'ลาก Block เพิ่ม ให้ AI ช่วยปรับปรุงเนื้อหา' },
  { name: 'Analytics Dashboard', desc: 'ดูสถิติ Real-time บนมือถือหรือเว็บ' },
  { name: 'Publish & Share', desc: 'URL พร้อม! แชร์ให้ HR ได้เลยทันที' }
]

const roles = ['💻 Programmer', '🎨 Designer', '💼 Business', '🌱 First Jobber']
const currentRole = ref(0)
const currentTpl = ref(0)
const templates = [
  { name: 'Dark Neon Pro', class: 'tpl-a', avatarBg: 'linear-gradient(135deg,#7C6FE0,#E879F9)', hlBg: 'rgba(124,111,224,0.5)', chips: ['Go', 'Vue'] },
  { name: 'Tech Minimal', class: 'tpl-b', avatarBg: 'linear-gradient(135deg,#34D399,#22D3EE)', hlBg: 'rgba(52,211,153,0.5)', chips: ['React'] },
  { name: 'Creative Studio', class: 'tpl-c', avatarBg: 'linear-gradient(135deg,#E879F9,#F4C842)', hlBg: 'rgba(232,121,249,0.5)', chips: ['Design'] },
  { name: 'Clean Corporate', class: 'tpl-d', avatarBg: 'linear-gradient(135deg,#22D3EE,#7C6FE0)', hlBg: 'rgba(34,211,238,0.5)' },
  { name: 'Bold Yellow', class: 'tpl-e', avatarBg: 'linear-gradient(135deg,#F4C842,#E879F9)', hlBg: 'rgba(244,200,66,0.5)' },
  { name: 'Warm Gradient', class: 'tpl-f', avatarBg: 'linear-gradient(135deg,#F87171,#F4C842)', hlBg: 'rgba(248,113,113,0.5)' }
]

const blocks = [
  { id: 'hero', name: 'Hero / Profile', icon: '👤' },
  { id: 'skills', name: 'Skills', icon: '💡' },
  { id: 'exp', name: 'Experience', icon: '💼' },
  { id: 'proj', name: 'Projects', icon: '🚀' },
  { id: 'edu', name: 'Education', icon: '🎓' },
  { id: 'contact', name: 'Contact', icon: '📬' }
]

const canvasSections = ref([
  { id: 'cs-hero', title: '👤 Hero / Profile', content: '<div class="hero-preview"><div class="hp-avatar">NK</div><div><div class="hp-name">Ploy Kong</div><div class="hp-role">Full-Stack Developer</div></div><div class="hp-hire">Hire Me ✨</div></div>' },
  { id: 'cs-skills', title: '💡 Skills', content: '<div class="skill-row-p"><span class="sk-name">Go / Fiber</span><div class="sk-track"><div class="sk-fill" style="width:92%"></div></div><span style="font-size:11px;color:var(--text2)">92%</span></div><div class="skill-row-p"><span class="sk-name">Vue.js 3</span><div class="sk-track"><div class="sk-fill" style="width:85%"></div></div><span style="font-size:11px;color:var(--text2)">85%</span></div>' }
])

const draggingBlock = ref(null)
const dropHighlight = ref(false)
const aiScore = ref(74)
const aiTypedText = ref('')
const isTyping = ref(false)

const stats = ref([
  { label: '👁️ Total Views', val: 0, delta: '↑ +23% this week' },
  { label: '⏱ Avg. Time', val: '0:00', delta: '↑ Good engagement' },
  { label: '💼 Hire Me Clicks', val: 0, delta: '↑ +5 today' },
  { label: '📄 CV Downloads', val: 0, delta: '↑ +12 this week' }
])

const trafficSources = [
  { name: 'LinkedIn', icon: 'in', bg: 'rgba(10,102,194,0.15)', width: 52, grad: 'linear-gradient(90deg,#0A66C2,#22D3EE)', color: '#22D3EE' },
  { name: 'Resume PDF', icon: '📄', bg: 'rgba(59,130,246,0.1)', width: 28, width: 28, grad: 'linear-gradient(90deg,#3B82F6,#7C6FE0)', color: '#7C6FE0' },
  { name: 'Facebook', icon: '📘', bg: 'rgba(24,119,242,0.1)', width: 20, grad: 'linear-gradient(90deg,#1877F2,#E879F9)', color: '#E879F9' }
]

const countries = [
  { flag: '🇹🇭', name: 'Thailand', num: 847 },
  { flag: '🇸🇬', name: 'Singapore', num: 234 },
  { flag: '🇺🇸', name: 'United States', num: 153 }
]

const shareOptions = [
  { icon: '💼', label: 'LinkedIn:', val: 'ploykong.com/pk' },
  { icon: '📧', label: 'Email Signature:', val: 'ploykong.com/pk' },
  { icon: '📄', label: 'Resume PDF:', val: 'QR Code สำหรับพอร์ตน้องใหม่' }
]

const toasts = ref([])
const showConfetti = ref(false)
const cursorPos = ref({ x: 0, y: 0 })
const demoShell = ref(null)

let progressInterval = null
let sceneSequenceTimeout = null

const goToScene = (idx) => {
  if (isTransitioning.value) return
  isTransitioning.value = true
  clearTimers()
  
  setTimeout(() => {
    currentScene.value = idx
    sceneAnims.value = sceneAnims.value.map((_, i) => i === idx)
    
    setTimeout(() => {
      isTransitioning.value = false
      if (isPlaying.value) startProgress()
      runSceneSpecificLogic(idx)
    }, 350)
  }, 350)
}

const nextScene = () => goToScene((currentScene.value + 1) % 4)
const prevScene = () => goToScene((currentScene.value + 3) % 4)

const togglePlay = () => {
  isPlaying.value = !isPlaying.value
  if (isPlaying.value) startProgress()
  else clearProgress()
}

const clearTimers = () => {
  clearProgress()
  clearTimeout(sceneSequenceTimeout)
  if (sceneTimers.length) {
    sceneTimers.forEach(clearTimeout)
    sceneTimers = []
  }
}

const clearProgress = () => {
  clearInterval(progressInterval)
  progressValue.value = 0
}

const startProgress = () => {
  clearProgress()
  const duration = 9000 / speedMult.value
  const step = 50
  const increment = (step / duration) * 100
  
  progressInterval = setInterval(() => {
    progressValue.value += increment
    if (progressValue.value >= 100) {
      clearInterval(progressInterval)
      sceneSequenceTimeout = setTimeout(nextScene, 300)
    }
  }, step)
}

let sceneTimers = []

const runSceneSpecificLogic = (idx) => {
  if (idx === 0) runScene0()
  if (idx === 1) runScene1()
  if (idx === 2) runScene2()
  if (idx === 3) runScene3()
}

// Scene 0: Template select
const runScene0 = () => {
  const tpls = [1, 2, 3, 0, 4, 5, 0]
  let i = 0
  const loop = () => {
    if (currentScene.value !== 0) return
    const targetIdx = tpls[i % tpls.length]
    currentTpl.value = targetIdx
    moveCursorToEl('tpl' + targetIdx)
    i++
    sceneTimers.push(setTimeout(loop, 1200 / speedMult.value))
  }
  loop()
}

// Scene 1: Drag & Drop
const runScene1 = () => {
  // Reset
  canvasSections.value = canvasSections.value.filter(s => !s.isNew)
  aiTypedText.value = ''
  isTyping.value = false
  aiScore.value = 74
  
  sceneTimers.push(setTimeout(() => {
    draggingBlock.value = 'exp'
    moveCursorToEl('block-exp')
    showToast('🖱️ ลาก Experience มาวาง!', 'var(--accent2)')
  }, 800 / speedMult.value))
  
  sceneTimers.push(setTimeout(() => {
    dropHighlight.value = true
    moveCursorToEl('dropZone')
  }, 1800 / speedMult.value))
  
  sceneTimers.push(setTimeout(() => {
    draggingBlock.value = null
    dropHighlight.value = false
    canvasSections.value.splice(2, 0, {
      id: 'cs-exp',
      title: '💼 Experience',
      isNew: true,
      content: '<div style="font-size:13px;font-weight:700;margin-bottom:4px">KMUTT IT Department</div><div style="font-size:11px;color:var(--text2);margin-bottom:6px">Backend Developer Intern · 2023–2024</div><div style="font-size:11px;color:var(--text2);line-height:1.6">worked on backend projects using go...</div>'
    })
    showToast('✅ เพิ่ม Experience Block แล้ว!', 'var(--green)')
  }, 2600 / speedMult.value))
  
  sceneTimers.push(setTimeout(() => {
    animateCount(aiScore, 74, 85, 1000)
  }, 3400 / speedMult.value))
  
  sceneTimers.push(setTimeout(() => {
    isTyping.value = true
    typeText('"Spearheaded development of high-performance microservices using Go/Fiber, achieving 40% latency reduction..."')
  }, 4200 / speedMult.value))
}

// Scene 2: Analytics
const runScene2 = () => {
  stats.value[0].val = 0
  stats.value[1].val = '0:00'
  stats.value[2].val = 0
  stats.value[3].val = 0
  
  sceneTimers.push(setTimeout(() => animateCountVal(0, 1234, 1200), 300))
  sceneTimers.push(setTimeout(() => animateTime(1, 187), 600))
  sceneTimers.push(setTimeout(() => animateCountVal(2, 23, 800), 900))
  sceneTimers.push(setTimeout(() => animateCountVal(3, 67, 1000), 1100))
  
  sceneTimers.push(setTimeout(() => {
    drawChart()
  }, 500))
}

const runScene3 = () => {
  sceneTimers.push(setTimeout(() => {
    triggerPublish()
  }, 1500 / speedMult.value))
}

const typeText = (text) => {
  let i = 0
  const interval = setInterval(() => {
    if (currentScene.value !== 1) { clearInterval(interval); return }
    aiTypedText.value += text[i]
    i++
    if (i >= text.length) {
      clearInterval(interval)
      isTyping.value = false
    }
  }, 40 / speedMult.value)
}

const animateCount = (targetRef, from, to, duration) => {
  const start = performance.now()
  const tick = (now) => {
    const t = Math.min((now-start)/duration, 1)
    targetRef.value = Math.round(from + (to-from)*t)
    if (t < 1) requestAnimationFrame(tick)
  }
  requestAnimationFrame(tick)
}

const animateCountVal = (idx, to, duration) => {
  const start = performance.now()
  const from = 0
  const tick = (now) => {
    const t = Math.min((now-start)/duration, 1)
    const val = Math.round(from + (to-from)*t)
    stats.value[idx].val = val.toLocaleString()
    if (t < 1) requestAnimationFrame(tick)
  }
  requestAnimationFrame(tick)
}

const animateTime = (idx, totalSec) => {
  let s = 0
  const interval = setInterval(() => {
    s += 3
    if (s > totalSec) { clearInterval(interval); s = totalSec }
    const m = Math.floor(s/60), sec = s%60
    stats.value[idx].val = `${m}:${String(sec).padStart(2,'0')}`
    if (s >= totalSec) clearInterval(interval)
  }, 60)
}

const showToast = (text, color) => {
  const id = Date.now()
  const icon = text.match(/^[\p{Emoji}]+/u)?.[0] || 'ℹ️'
  const cleanText = text.replace(/^[\p{Emoji}]+\s*/u, '')
  toasts.value.push({ id, text: cleanText, icon, color })
  setTimeout(() => {
    toasts.value = toasts.value.filter(t => t.id !== id)
  }, 4000)
}

const triggerPublish = () => {
  showConfetti.value = true
  showToast('🚀 Published! ploykong.com/pk', 'var(--green)')
  setTimeout(() => showConfetti.value = false, 4000)
}

const getConfettiStyle = () => {
  const colors = ['#7C6FE0', '#E879F9', '#22D3EE', '#F4C842', '#34D399', '#F87171']
  return {
    left: Math.random() * 100 + 'vw',
    background: colors[Math.floor(Math.random() * colors.length)],
    width: (Math.random() * 8 + 4) + 'px',
    height: (Math.random() * 8 + 4) + 'px',
    borderRadius: Math.random() > 0.5 ? '50%' : '2px',
    animationDuration: (Math.random() * 2 + 2) + 's',
    animationDelay: (Math.random() * 1) + 's'
  }
}

const moveCursorToEl = (id) => {
  const el = document.getElementById(id)
  if (!el || !demoShell.value) return
  const rect = el.getBoundingClientRect()
  const shellRect = demoShell.value.getBoundingClientRect()
  cursorPos.value = {
    x: rect.left - shellRect.left + rect.width / 2,
    y: rect.top - shellRect.top + rect.height / 2
  }
}

const chartArea = ref(null)
const chartSvg = ref(null)

const drawChart = () => {
  if (!chartSvg.value) return
  const svg = chartSvg.value
  svg.innerHTML = ''
  const viewData = [30, 55, 45, 80, 62, 90, 75, 95, 65, 80, 55, 100]
  const clickData = [10, 22, 18, 35, 25, 40, 30, 42, 28, 35, 22, 50]
  const W = 600, H = 160, pad = 20

  const dataToPath = (data, filled = false) => {
    const pts = data.map((v, i) => {
      const x = pad + (i / (data.length - 1)) * (W - pad * 2)
      const y = H - pad - (v / 100) * (H - pad * 2)
      return `${x},${y}`
    })
    let d = `M ${pts.join(' L ')}`
    if (filled) d += ` L ${W - pad},${H - pad} L ${pad},${H - pad} Z`
    return d
  }

  const defs = document.createElementNS('http://www.w3.org/2000/svg', 'defs')
  defs.innerHTML = `<linearGradient id="cg1" x1="0" y1="0" x2="0" y2="1"><stop offset="0%" stop-color="#7C6FE0" stop-opacity="0.35"/><stop offset="100%" stop-color="#7C6FE0" stop-opacity="0.02"/></linearGradient><linearGradient id="cg2" x1="0" y1="0" x2="0" y2="1"><stop offset="0%" stop-color="#22D3EE" stop-opacity="0.25"/><stop offset="100%" stop-color="#22D3EE" stop-opacity="0.02"/></linearGradient>`
  svg.appendChild(defs)

  const area1 = document.createElementNS('http://www.w3.org/2000/svg', 'path')
  area1.setAttribute('d', dataToPath(viewData, true)); area1.setAttribute('fill', 'url(#cg1)'); svg.appendChild(area1)
  const area2 = document.createElementNS('http://www.w3.org/2000/svg', 'path')
  area2.setAttribute('d', dataToPath(clickData, true)); area2.setAttribute('fill', 'url(#cg2)'); svg.appendChild(area2)

  const animatePath = (data, color, delay = 0) => {
    const path = document.createElementNS('http://www.w3.org/2000/svg', 'path')
    path.setAttribute('d', dataToPath(data)); path.setAttribute('fill', 'none'); path.setAttribute('stroke', color); path.setAttribute('stroke-width', '2.5'); path.setAttribute('stroke-linecap', 'round')
    const len = 800; path.setAttribute('stroke-dasharray', len); path.setAttribute('stroke-dashoffset', len)
    svg.appendChild(path)
    setTimeout(() => { path.style.transition = 'stroke-dashoffset 1.5s ease-out'; path.setAttribute('stroke-dashoffset', '0') }, delay)
  }
  animatePath(viewData, '#7C6FE0', 200); animatePath(clickData, '#22D3EE', 500)
}

onMounted(() => {
  setTimeout(() => goToScene(0), 500)
  
  const handleResize = () => {
    // recalibrate cursor if needed
  }
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  clearTimers()
})
</script>

<style scoped>
.demo-section {
  padding: 30px 20px 50px 20px;
  background: var(--bg);
  overflow: hidden;
}
@media (max-width: 600px) {
  .demo-section { padding: 40px 10px; }
}

.section-wrap {
  max-width: 1200px;
  margin: 0 auto;
}

.section-header {
  text-align: center;
  margin-bottom: 60px;
}

.demo-shell-container {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 24px;
  padding: 12px;
  box-shadow: var(--shadow-xl);
  position: relative;
}

.demo-shell {
  display: grid;
  grid-template-rows: 56px 1fr 80px;
  height: 600px;
  width: 100%;
  border-radius: 16px;
  overflow: hidden;
  background: #07080F;
  color: #EEF0FF;
  font-family: 'Plus Jakarta Sans', 'Anuphan', system-ui, sans-serif;
  position: relative;
  --bg: #07080F;
  --surf: #0F1120;
  --surf2: #141629;
  --surf3: #1A1E35;
  --border: rgba(255,255,255,0.07);
  --border2: rgba(255,255,255,0.12);
  --text2: #7B83B4;
  --text3: #3E4468;
  --accent: #7C6FE0;
  --accent2: #A899F8;
  --accentg: rgba(124,111,224,0.18);
  --cyan: #22D3EE;
  --pink: #E879F9;
  --gold: #F4C842;
  --green: #34D399;
  --grad: linear-gradient(135deg,#7C6FE0 0%,#E879F9 50%,#22D3EE 100%);
  --grad2: linear-gradient(135deg,#5B4FD0,#8B5CF6 60%,#7C6FE0);
}

.topbar {
  display: flex; align-items: center; gap: 16px;
  padding: 0 20px; background: var(--surf); border-bottom: 1px solid var(--border); z-index: 10;
}
@media (max-width: 640px) {
  .topbar { padding: 0 12px; height: 50px; }
  .tb-right { display: none; }
}
.tb-logo { display: flex; align-items: center; gap: 10px; }
.demo-logo-img { width: 24px; height: 24px; border-radius: 6px; }
.tb-logoname { font-family: 'Syne', sans-serif; font-size: 16px; font-weight: 800; letter-spacing: -0.3px; }
.tb-logoname em { font-style: normal; background: var(--grad); -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text; }

.tb-steps { 
  display: flex; align-items: center; gap: 4px; margin: 0 auto;
  overflow-x: auto; -webkit-overflow-scrolling: touch;
  scrollbar-width: none;
}
.tb-steps::-webkit-scrollbar { display: none; }
.tb-step {
  display: flex; align-items: center; gap: 6px; padding: 6px 10px; border-radius: 100px;
  font-size: 10px; font-weight: 600; color: var(--text3); cursor: pointer; transition: all 0.3s;
  white-space: nowrap;
}
.tb-step.active { background: var(--accentg); color: var(--accent2); }
.tb-step.done { color: var(--green); }
.tb-step-dot { width: 6px; height: 6px; border-radius: 50%; background: currentColor; opacity: 0.5; }
.tb-step.active .tb-step-dot { opacity: 1; box-shadow: 0 0 8px var(--accent2); }

.tb-badge {
  display: flex; align-items: center; gap: 6px; padding: 5px 12px; border-radius: 100px;
  font-size: 10px; font-weight: 700; background: rgba(52,211,153,0.1); border: 1px solid rgba(52,211,153,0.25); color: var(--green);
}
.live-dot { width: 6px; height: 6px; border-radius: 50%; background: var(--green); animation: pulse 1.5s infinite; }
@keyframes pulse { 0% { opacity: 1; } 50% { opacity: 0.4; } 100% { opacity: 1; } }

.stage { position: relative; overflow: hidden; background: var(--bg); flex: 1; }
.stage-bg { position: absolute; inset: 0; pointer-events: none; }
.orb { position: absolute; border-radius: 50%; filter: blur(80px); opacity: 0.4; }
.orb1 { width: 400px; height: 400px; background: var(--accent); top: -100px; left: -100px; }
.orb2 { width: 300px; height: 300px; background: var(--pink); bottom: -100px; right: -50px; }
.grid-bg { position: absolute; inset: 0; background-image: radial-gradient(var(--border) 1px, transparent 1px); background-size: 32px 32px; opacity: 0.5; }

.scenes-container { position: relative; width: 100%; height: 100%; }
.scene {
  position: absolute; inset: 0; display: flex; align-items: center; justify-content: center;
  padding: 20px; opacity: 0; pointer-events: none; transition: opacity 0.5s;
  overflow-y: auto;
}
.scene.active { opacity: 1; pointer-events: auto; }

.s1-wrap { display: grid; grid-template-columns: 1fr 1.2fr; gap: 40px; width: 100%; max-width: 900px; }
.s1-h { 
  font-family: 'Space Grotesk', 'Anuphan', sans-serif; 
  font-size: 42px; 
  font-weight: 700; 
  line-height: 1.15; 
  margin-bottom: 14px; 
  letter-spacing: -0.5px;
}
.s1-eyebrow { font-size: 10px; font-weight: 700; letter-spacing: 2.5px; color: var(--text2); text-transform: uppercase; margin-bottom: 10px; }
.s1-p { font-size: 13px; color: var(--text2); margin-bottom: 20px; line-height: 1.7; font-weight: 400; }
.role-tags { display: flex; flex-wrap: wrap; gap: 8px; margin-bottom: 20px; }
.role-tag { 
  padding: 6px 14px; 
  border-radius: 100px; 
  font-size: 12px; 
  font-weight: 600;
  background: var(--surf2); 
  border: 1px solid var(--border); 
  cursor: pointer; 
  transition: 0.2s;
  letter-spacing: 0.1px;
}
.role-tag.sel { background: var(--accentg); border-color: var(--accent); color: var(--accent2); }

.s1-right { display: grid; grid-template-columns: repeat(3, 1fr); gap: 10px; }
.tpl-card { border-radius: 12px; overflow: hidden; border: 2px solid var(--border); cursor: pointer; transition: 0.3s; }
.tpl-card.sel { border-color: var(--accent); transform: translateY(-2px); }
.tpl-preview { height: 80px; padding: 10px; position: relative; }
.tpl-label { padding: 6px; background: var(--surf2); font-size: 10px; font-weight: 700; text-align: center; }
.tpl-a { background: #0D0D2B; } .tpl-b { background: #0D1F1A; } .tpl-c { background: #1A0D20; }
.tpl-d { background: #0D1A2A; } .tpl-e { background: #1A1A0D; } .tpl-f { background: #1A0D0D; }
.m-avatar { width: 16px; height: 16px; border-radius: 50%; margin-bottom: 4px; }
.m-line { height: 3px; border-radius: 2px; background: rgba(255,255,255,0.1); margin-bottom: 3px; width: 80%; }
.m-line.hl { width: 50%; }
.tpl-check { position: absolute; top: 5px; right: 5px; width: 14px; height: 14px; border-radius: 50%; background: var(--accent); font-size: 9px; display: flex; align-items: center; justify-content: center; opacity: 0; }
.tpl-card.sel .tpl-check { opacity: 1; }

.s2-wrap { display: grid; grid-template-columns: 180px 1fr 200px; gap: 16px; width: 100%; height: 400px; }
.s2-sidebar, .s2-canvas, .s2-ai { background: var(--surf); border: 1px solid var(--border); border-radius: 12px; overflow: hidden; display: flex; flex-direction: column; }
.s2-sidebar-h, .s2-canvas-h, .ai-panel-h { padding: 10px; border-bottom: 1px solid var(--border); font-size: 11px; font-weight: 700; color: var(--text2); }
.block-item { padding: 8px; font-size: 11px; display: flex; align-items: center; gap: 8px; border-bottom: 1px solid var(--border); }
.block-item.dragging-from { background: var(--accentg); color: var(--accent2); }
.canvas-body { padding: 12px; flex: 1; overflow: auto; display: flex; flex-direction: column; gap: 8px; }
.canvas-section { background: var(--surf2); border: 1px solid var(--border); border-radius: 8px; overflow: hidden; }
.canvas-section.new-section { animation: slideIn 0.4s ease-out; }
@keyframes slideIn { from { opacity: 0; transform: translateY(-10px); } to { opacity: 1; transform: translateY(0); } }
.cs-header { padding: 6px 10px; border-bottom: 1px solid var(--border); font-size: 10px; font-weight: 700; color: var(--text2); }
.cs-content { padding: 8px 10px; }
.drop-zone { border: 2px dashed var(--border2); border-radius: 8px; padding: 12px; text-align: center; font-size: 10px; color: var(--text3); }
.drop-zone.highlight { border-color: var(--accent); background: var(--accentg); color: var(--accent2); }

.ai-body { padding: 12px; display: flex; flex-direction: column; gap: 12px; }
.ai-score-box { background: var(--surf2); padding: 10px; border-radius: 8px; }
.ai-score-num { font-size: 28px; font-weight: 700; color: var(--accent2); font-family: 'Space Grotesk', sans-serif; letter-spacing: -1px; }
.ai-score-bar { height: 4px; background: var(--surf3); border-radius: 2px; margin: 6px 0; overflow: hidden; }
.ai-score-fill { height: 100%; background: var(--grad); transition: width 1s; }
.ai-improve-box { background: var(--surf2); border-radius: 8px; font-size: 10px; }
.ai-text-in { padding: 8px; color: var(--text2); border-bottom: 1px solid var(--border); }
.ai-text-out { padding: 8px; color: var(--accent2); min-height: 40px; }
.ai-action { background: var(--grad2); padding: 8px; border-radius: 6px; text-align: center; font-size: 11px; font-weight: 700; }

.stat-cards { display: grid; grid-template-columns: repeat(2, 1fr); gap: 10px; margin-bottom: 16px; }
.stat-card { background: var(--surf); border: 1px solid var(--border); border-radius: 12px; padding: 12px; }
.sc-label { font-size: 10px; color: var(--text2); font-weight: 500; }
.sc-num { font-size: 20px; font-weight: 700; margin: 4px 0; font-family: 'Space Grotesk', sans-serif; letter-spacing: -0.5px; }
.sc-delta { font-size: 10px; color: var(--green); }
.chart-card { background: var(--surf); border: 1px solid var(--border); border-radius: 12px; padding: 16px; }
.chart-area { height: 120px; position: relative; }
.source-card, .map-card { background: var(--surf); border: 1px solid var(--border); border-radius: 12px; padding: 12px; }
.source-row { display: flex; align-items: center; gap: 8px; margin-bottom: 8px; font-size: 10px; }
.source-icon { width: 20px; height: 20px; border-radius: 4px; display: flex; align-items: center; justify-content: center; }
.source-bar-wrap { flex: 1; height: 4px; background: var(--surf3); border-radius: 2px; }
.source-bar { height: 100%; border-radius: 2px; transition: width 1.5s; }

.s3-wrap { 
  display: grid; 
  grid-template-columns: 1fr 1fr; 
  gap: 20px; 
  width: 100%; 
  height: 100%; 
  max-height: 400px;
  align-content: start;
  overflow: auto;
}
.s3-left { display: flex; flex-direction: column; gap: 12px; overflow: auto; }
.s3-right { display: flex; flex-direction: column; gap: 12px; overflow: auto; }
.publish-btn { display: inline-block; padding: 12px 24px; background: var(--grad2); border-radius: 12px; font-weight: 700; margin: 20px 0; cursor: pointer; width: 100%; text-align: center; font-family: 'Plus Jakarta Sans', sans-serif; letter-spacing: 0.2px; }
.url-pill { display: inline-flex; align-items: center; gap: 8px; padding: 8px 16px; background: var(--surf2); border-radius: 100px; font-family: 'JetBrains Mono', 'Fira Code', monospace; font-size: 11px; max-width: 100%; overflow: hidden; }
.portfolio-preview { background: var(--surf); border: 1px solid var(--border2); border-radius: 16px; overflow: hidden; box-shadow: 0 20px 40px rgba(0,0,0,0.4); }
.s4-wrap { display: grid; grid-template-columns: 1fr 1fr; gap: 24px; align-items: center; width: 100%; max-width: 900px; }
.pp-hero { padding: 16px; background: #0D0D2B; }
.pp-avatar { width: 40px; height: 40px; border-radius: 50%; background: var(--grad); display: flex; align-items: center; justify-content: center; font-weight: 900; }

.controls {
  display: flex; align-items: center; gap: 20px; padding: 0 28px; background: var(--surf); border-top: 1px solid var(--border);
}
.ctrl-info { flex: 1; }
.ctrl-step-name { font-weight: 700; font-size: 13px; font-family: 'Space Grotesk', sans-serif; letter-spacing: -0.2px; }
.ctrl-step-desc { font-size: 10px; color: var(--text2); margin-top: 2px; }
.progress-track { flex: 2; height: 4px; background: var(--surf3); border-radius: 2px; overflow: hidden; }
.progress-fill { height: 100%; background: var(--grad); transition: width 0.1s linear; }
.ctrl-btns { display: flex; gap: 8px; align-items: center; }
.ctrl-btn { padding: 8px 12px; border-radius: 8px; background: var(--surf2); border: 1px solid var(--border); font-size: 11px; font-weight: 600; cursor: pointer; display: flex; align-items: center; gap: 4px; color: #fff; }
.ctrl-btn.primary { background: var(--grad2); border: none; }
.speed-select { background: var(--surf2); border: 1px solid var(--border); color: #fff; padding: 6px; border-radius: 6px; font-size: 11px; }

.demo-cursor {
  position: absolute; width: 20px; height: 20px; background: rgba(124,111,224,0.4); border: 2px solid var(--accent2);
  border-radius: 50%; pointer-events: none; z-index: 100; transform: translate(-50%, -50%); transition: left 0.4s, top 0.4s;
}
.demo-cursor::after { content: ''; position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%); width: 4px; height: 4px; background: #fff; border-radius: 50%; }

.fly-in, .fly-in-l, .fly-in-r { opacity: 0; transition: all 0.6s cubic-bezier(0.34, 1.56, 0.64, 1); }
.fly-in { transform: translateY(20px); }
.fly-in-l { transform: translateX(-20px); }
.fly-in-r { transform: translateX(20px); }
.fly-in.show, .fly-in-l.show, .fly-in-r.show { opacity: 1; transform: translate(0); }

.toast-area { position: absolute; top: 20px; right: 20px; display: flex; flex-direction: column; gap: 8px; z-index: 1000; max-width: calc(100% - 40px); }
.toast { background: var(--surf); border: 1px solid var(--border2); padding: 10px 16px; border-radius: 10px; font-size: 11px; display: flex; align-items: center; gap: 8px; backdrop-filter: blur(10px); box-shadow: 0 4px 12px rgba(0,0,0,0.2); }

@media (max-width: 600px) {
  .toast-area { top: auto; bottom: 80px; left: 50%; right: auto; transform: translateX(-50%); align-items: center; width: 100%; }
  .toast { width: fit-content; max-width: 280px; }
}

.confetti-wrap { position: absolute; inset: 0; pointer-events: none; z-index: 99; }
.confetti-piece { position: absolute; top: -10px; animation: fall linear both; }
@keyframes fall { to { transform: translateY(600px) rotate(720deg); opacity: 0; } }

.scene-transition { position: absolute; inset: 0; background: #07080F; z-index: 500; opacity: 0; pointer-events: none; transition: opacity 0.35s; }
.scene-transition.active { opacity: 1; pointer-events: auto; }

.toast-enter-active, .toast-leave-active { transition: all 0.4s; }
.toast-enter-from { opacity: 0; transform: translateX(30px); }
.toast-leave-to { opacity: 0; transform: translateX(20px); }

.typing-cursor { display: inline-block; width: 2px; height: 12px; background: var(--accent2); animation: blink 1s infinite; vertical-align: middle; }
@keyframes blink { 50% { opacity: 0; } }

@media (max-width: 900px) {
  .demo-shell { height: auto; min-height: 560px; grid-template-rows: auto 1fr auto; overflow: hidden; }
  .s1-wrap, .s2-wrap, .s4-wrap { 
    grid-template-columns: 1fr; 
    gap: 20px; 
    padding: 20px 16px; 
    justify-items: center;
    text-align: center;
    width: 100%;
  }
  .s3-wrap { grid-template-columns: 1fr; gap: 12px; padding: 16px; }
  .s1-right { grid-template-columns: repeat(3, 1fr); gap: 8px; }
  .s2-canvas { width: 100%; max-width: 100%; }
  .s1-h { font-size: 26px; line-height: 1.2; text-align: center; }
  .s1-p { text-align: center; font-size: 11px; }
  .s1-eyebrow { text-align: center; }
  .role-tags { justify-content: center; }
  .s2-sidebar, .s2-ai { display: none; }
  .s4-wrap { align-items: start; }
  .portfolio-preview { display: none; } /* Hide preview card — too cramped */
  .s4-left { width: 100%; }
  .controls { padding: 10px 12px; flex-wrap: wrap; height: auto; justify-content: center; gap: 8px; }
  .ctrl-info { width: 100%; text-align: center; order: 1; }
  .progress-track { order: 3; width: 100%; flex: none; margin-top: 4px; }
  .ctrl-btns { order: 2; }
  .speed-select { display: none; } /* hide speed selector on small screens */
  .scene { padding: 16px 12px; align-items: flex-start; overflow-y: auto; }
}
@media (max-width: 640px) {
  .demo-shell { border-radius: 12px; min-height: 500px; }
  .s1-h { font-size: 22px; }
  .s1-right { grid-template-columns: repeat(2, 1fr); gap: 6px; }
  .tpl-preview { height: 64px; }
  .tpl-label { font-size: 9px; padding: 4px; }
  .stat-cards { grid-template-columns: repeat(2, 1fr); gap: 6px; }
  .stat-card { padding: 8px; }
  .sc-num { font-size: 16px; }
  .chart-card { padding: 10px; }
  .chart-area { height: 80px; }
  .ctrl-btn { padding: 6px 8px; font-size: 10px; }
  .demo-shell-container { padding: 6px; border-radius: 18px; }
  .source-card { padding: 8px; }

  /* Scene 4 on small screens */
  .share-options { font-size: 11px; }
  .url-pill { font-size: 10px; padding: 6px 10px; }
}
</style>
