<template>
  <div class="landing" :class="{ 'theme-dark': isDark, 'theme-light': !isDark }">

    <!-- ══════════════════════════════════════════
         NAVBAR
    ════════════════════════════════════════════ -->
    <nav class="navbar" :class="{ scrolled: scrolled }">
      <div class="nav-inner">
        <!-- Logo -->
        <RouterLink to="/" class="nav-logo">
          <div class="logo-mark">
            <img src="/favicon.png" alt="PloyKong Logo" class="logo-img" />
          </div>
          <span class="logo-text">Ploy<em>Kong</em></span>
        </RouterLink>

        <!-- Nav links (desktop) -->
        <div class="nav-center">
          <a href="#features" class="nav-link">Features</a>
          <a href="#how" class="nav-link">How it works</a>
          <a href="#pricing" class="nav-link">Pricing</a>
        </div>

        <!-- Nav right -->
        <div class="nav-right">
          <ThemeToggle />
          <RouterLink to="/login" class="btn-ghost">Login</RouterLink>
          <RouterLink to="/register" class="btn-cta">
            <span>เริ่มฟรี</span>
            <ArrowRight :size="14" />
          </RouterLink>
        </div>

        <!-- Mobile hamburger -->
        <button class="hamburger" @click="mobileOpen = !mobileOpen">
          <Menu v-if="!mobileOpen" :size="20" />
          <X v-else :size="20" />
        </button>
      </div>

      <!-- Mobile menu -->
      <Transition name="mobile-menu">
        <div v-if="mobileOpen" class="mobile-menu">
          <a href="#features" class="mobile-link" @click="mobileOpen=false">Features</a>
          <a href="#how" class="mobile-link" @click="mobileOpen=false">How it works</a>
          <a href="#pricing" class="mobile-link" @click="mobileOpen=false">Pricing</a>
          <div class="mobile-divider"></div>
          <RouterLink to="/login" class="mobile-link" @click="mobileOpen=false">Login</RouterLink>
          <RouterLink to="/register" class="btn-cta w-full" @click="mobileOpen=false">เริ่มฟรีเลย 🚀</RouterLink>
        </div>
      </Transition>
    </nav>

    <!-- ══════════════════════════════════════════
         HERO
    ════════════════════════════════════════════ -->
    <section class="hero">
      <!-- BG mesh -->
      <div class="hero-mesh" aria-hidden="true">
        <div class="mesh-orb orb-1"></div>
        <div class="mesh-orb orb-2"></div>
        <div class="mesh-orb orb-3"></div>
        <div class="grid-overlay"></div>
      </div>

      <div class="hero-inner">
        <!-- Left content -->
        <div class="hero-content">
          <div class="hero-pill">
            <span class="pill-dot"></span>
            No-Code · AI-Powered · Gen-Z Ready
          </div>

          <h1 class="hero-h1">
            <span class="h1-line">มีดีต้อง</span>
            <span class="h1-accent">ปล่อยของ</span>
            <span class="h1-line h1-sub">Portfolio Builder</span>
          </h1>

          <p class="hero-body">
            สร้างเว็บ Portfolio ระดับมืออาชีพด้วยระบบ Drag & Drop<br/>
            ให้ AI ช่วยเขียนเนื้อหา — เสร็จใน 5 นาที ไม่ต้องรู้โค้ด
          </p>

          <div class="hero-actions">
            <RouterLink to="/register" class="btn-hero-primary">
              <Rocket :size="16" />
              เริ่มสร้างพอร์ตฟรี
            </RouterLink>
            <a href="#demo" class="btn-hero-ghost">
              <Play :size="14" />
              ดู Demo 2 นาที
            </a>
          </div>

          <!-- Social proof -->
          <div class="hero-proof">
            <div class="proof-avatars">
              <div class="proof-av" v-for="i in 4" :key="i" :style="{ background: avColors[i-1] }">
                {{ avNames[i-1] }}
              </div>
            </div>
            <div class="proof-text">
              <span class="proof-num">2,400+</span> คนสร้างพอร์ตสัปดาห์นี้
            </div>
          </div>
        </div>

        <!-- Right: Browser mockup -->
        <div class="hero-visual" aria-hidden="true">
          <div class="browser-frame">
            <div class="browser-bar">
              <div class="browser-dots">
                <span></span><span></span><span></span>
              </div>
              <div class="browser-url">
                <Lock2 :size="10" />
                pk.ploykong.com
              </div>
            </div>
            <div class="browser-body">
              <!-- Mini portfolio preview -->
              <div class="mini-portfolio">
                <div class="mp-hero">
                  <div class="mp-avatar">PK</div>
                  <div>
                    <div class="mp-name">Ploy K.</div>
                    <div class="mp-role">Full-Stack Developer</div>
                  </div>
                  <div class="mp-hire">Hire Me ✨</div>
                </div>
                <div class="mp-skills">
                  <div class="mp-skill-bar" v-for="s in skills" :key="s.name">
                    <span class="mp-skill-name">{{ s.name }}</span>
                    <div class="mp-skill-track">
                      <div class="mp-skill-fill" :style="{ width: s.pct + '%' }"></div>
                    </div>
                  </div>
                </div>
                <div class="mp-projects">
                  <div class="mp-proj" v-for="p in projects" :key="p">
                    <div class="mp-proj-dot"></div>
                    <div class="mp-proj-name">{{ p }}</div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Floating cards -->
          <div class="float-card fc-ai">
            <Sparkles :size="14" class="fc-icon" />
            <div>
              <div class="fc-label">AI Improved!</div>
              <div class="fc-val">"Led development of a high-impact..."</div>
            </div>
          </div>
          <div class="float-card fc-stat">
            <TrendingUp :size="14" class="fc-icon" />
            <div>
              <div class="fc-label">This Week</div>
              <div class="fc-val">1,234 views 👁️</div>
            </div>
          </div>
          <div class="float-card fc-notify">
            <Bell :size="14" class="fc-icon" />
            <div class="fc-label">HR กด Hire Me! 🎉</div>
          </div>
        </div>
      </div>

      <!-- Stats bar -->
      <div class="hero-stats">
        <div class="hstat" v-for="s in heroStats" :key="s.label">
          <span class="hstat-num">{{ s.num }}</span>
          <span class="hstat-label">{{ s.label }}</span>
        </div>
      </div>
    </section>

    <!-- ══════════════════════════════════════════
         MARQUEE LOGOS / TECH STRIP
    ════════════════════════════════════════════ -->
    <div class="marquee-strip">
      <div class="marquee-label">Built with</div>
      <div class="marquee-track">
        <div class="marquee-inner">
          <span v-for="t in [...techs, ...techs, ...techs]" :key="t+Math.random()" class="marquee-item">{{ t }}</span>
        </div>
      </div>
    </div>

    <!-- ══════════════════════════════════════════
         FEATURES
    ════════════════════════════════════════════ -->
    <section id="features" class="features-section">
      <div class="section-wrap">
        <div class="section-header">
          <div class="section-tag">FEATURES</div>
          <h2 class="section-h2">ครบทุกอย่างที่พอร์ตดีต้องมี</h2>
          <p class="section-sub">จากการสร้างจนถึงการ Publish — ทุกฟีเจอร์ถูกออกแบบมาเพื่อให้คุณโดดเด่น</p>
        </div>

        <!-- Big feature cards -->
        <div class="feat-bento">
          <!-- Card 1: Big - Drag & Drop -->
          <div class="bento-card bento-lg feat-builder">
            <div class="bento-glow"></div>
            <div class="feat-tag">CORE</div>
            <h3 class="bento-h3">Drag & Drop Builder</h3>
            <p class="bento-p">ลากวางส่วนประกอบได้อิสระ แก้ไขแล้วเห็นผลทันที — WYSIWYG ที่ลื่นไหลที่สุด</p>
            <!-- Mini drag UI mockup -->
            <div class="drag-demo">
              <div class="drag-block drag-b1">📋 ประสบการณ์การทำงาน</div>
              <div class="drag-block drag-b2 dragging">🚀 โปรเจกต์ <span class="drag-handle">⠿</span></div>
              <div class="drag-block drag-b3">💡 ทักษะ</div>
            </div>
          </div>

          <!-- Card 2: AI -->
          <div class="bento-card bento-sm feat-ai">
            <div class="bento-glow pink"></div>
            <div class="feat-tag pink">AI</div>
            <h3 class="bento-h3">AI Content Writer</h3>
            <p class="bento-p">ให้ AI เกลาข้อความให้ดูเป็น Pro</p>
            <div class="ai-bubble">
              <div class="ai-in">worked on some projects...</div>
              <div class="ai-arrow">↓ AI ✨</div>
              <div class="ai-out">"Spearheaded development of enterprise-scale applications..."</div>
            </div>
          </div>

          <!-- Card 3: Analytics -->
          <div class="bento-card bento-sm feat-analytics">
            <div class="bento-glow cyan"></div>
            <div class="feat-tag cyan">ANALYTICS</div>
            <h3 class="bento-h3">Real-time Dashboard</h3>
            <p class="bento-p">ดูสถิติผู้เข้าชมบนมือถือ แบบ Real-time</p>
            <div class="analytics-bars">
              <div v-for="(h,i) in chartData" :key="i" class="analytics-bar" :style="{ height: h + '%' }"></div>
            </div>
          </div>

          <!-- Card 4: Privacy -->
          <div class="bento-card bento-sm feat-privacy">
            <div class="bento-glow gold"></div>
            <div class="feat-tag gold">PRIVACY</div>
            <h3 class="bento-h3">Smart Privacy</h3>
            <p class="bento-p">ล็อกรหัสผ่านเฉพาะส่วน สร้างลิงก์ชั่วคราว</p>
            <div class="lock-demo">
              <Lock2 :size="32" class="lock-icon" />
              <div class="lock-text">NDA-protected section</div>
              <div class="lock-timer">⏱ หมดอายุ 24h</div>
            </div>
          </div>

          <!-- Card 7: Templates (Moved here to fill Row 2) -->
          <div class="bento-card bento-sm feat-templates">
            <div class="bento-glow purple"></div>
            <div class="feat-tag purple">DESIGN</div>
            <h3 class="bento-h3">Premium Templates</h3>
            <p class="bento-p">เลือกเทมเพลตที่เหมาะกับสายงานคุณ</p>
            <div class="tpl-mini-grid">
              <div class="tpl-dot gold"></div>
              <div class="tpl-dot cyan"></div>
              <div class="tpl-dot pink"></div>
            </div>
          </div>

          <!-- Card 5: AI Interview Coach - big -->
          <div class="bento-card bento-md feat-coach">
            <div class="bento-glow"></div>
            <div class="feat-tag">AI COACH</div>
            <h3 class="bento-h3">AI Interview Coach</h3>
            <p class="bento-p">แชทบอทอัจฉริยะในหน้าพอร์ตสาธารณะ — HR พิมพ์ถามข้อมูลเกี่ยวกับคุณได้โดยตรง</p>
            <div class="chat-demo">
              <div class="chat-msg hr">
                <div class="chat-avatar hr-av">HR</div>
                <div class="chat-bubble">What's Ploykong's experience with microservices?</div>
              </div>
              <div class="chat-msg ai">
                <div class="chat-avatar ai-av">🤖</div>
                <div class="chat-bubble">Ploykong has 3 years of experience building microservices with Go, having led the architecture of...</div>
              </div>
            </div>
          </div>

          <!-- Card 6: Publish -->
          <div class="bento-card bento-sm feat-publish">
            <div class="bento-glow cyan"></div>
            <div class="feat-tag cyan">DEPLOY</div>
            <h3 class="bento-h3">One-Click Publish</h3>
            <p class="bento-p">URL สวย แชร์ได้ทันที อัปเดตเข้าลิงก์เดิม</p>
            <div class="url-demo">
              <span class="url-base">ploykong.com/</span><span class="url-slug">pk</span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- ══════════════════════════════════════════
         HOW IT WORKS
    ════════════════════════════════════════════ -->
    <section id="how" class="how-section">
      <div class="section-wrap">
        <div class="section-header">
          <div class="section-tag">HOW IT WORKS</div>
          <h2 class="section-h2">3 ขั้นตอน สร้างพอร์ตมืออาชีพ</h2>
        </div>

        <div class="steps-grid">
          <div class="step-card" v-for="(step, i) in steps" :key="i">
            <div class="step-number">{{ String(i + 1).padStart(2, '0') }}</div>
            <div class="step-icon-wrap">
              <span class="step-icon">{{ step.icon }}</span>
            </div>
            <h3 class="step-title">{{ step.title }}</h3>
            <p class="step-desc">{{ step.desc }}</p>
            <div class="step-connector" v-if="i < 2"></div>
          </div>
        </div>

        <!-- Timeline visual -->
        <div class="timeline">
          <div class="timeline-bar">
            <div class="timeline-fill"></div>
          </div>
          <div class="timeline-labels">
            <span>เลือก Template</span>
            <span>Customize ด้วย AI</span>
            <span>Publish & แชร์</span>
          </div>
        </div>
      </div>
    </section>

    <!-- Interactive Demo Section -->
    <LandingDemo />

    <!-- ══════════════════════════════════════════
         PRICING
    ════════════════════════════════════════════ -->
    <section id="pricing" class="pricing-section">
      <div class="section-wrap">
        <div class="section-header">
          <div class="section-tag">PRICING</div>
          <h2 class="section-h2">ราคาที่เป็นมิตร เริ่มฟรีได้เลย</h2>
          <!-- Billing toggle -->
          <div class="billing-toggle">
            <span :class="{ active: !annual }">รายเดือน</span>
            <button class="toggle-track" :class="{ on: annual }" @click="annual = !annual">
              <div class="toggle-thumb"></div>
            </button>
            <span :class="{ active: annual }">รายปี <span class="save-badge">ประหยัด 30%</span></span>
          </div>
        </div>

        <div class="pricing-grid">
          <div
            class="price-card"
            v-for="(plan, i) in plans"
            :key="plan.name"
            :class="{ featured: plan.featured }"
          >
            <div v-if="plan.featured" class="price-badge">⭐ ยอดนิยม</div>
            <div class="price-header">
              <div class="price-name">{{ plan.name }}</div>
              <div class="price-amount">
                <span class="price-num">{{ annual ? plan.priceAnnual : plan.price }}</span>
                <span class="price-per">฿ / เดือน</span>
              </div>
              <p class="price-desc">{{ plan.desc }}</p>
            </div>
            <ul class="price-features">
              <li v-for="f in plan.features" :key="f">
                <Check :size="14" class="feat-check" />
                {{ f }}
              </li>
            </ul>
            <RouterLink to="/register" class="price-btn" :class="{ 'price-btn-primary': plan.featured }">
              {{ plan.cta }}
              <ArrowRight :size="14" />
            </RouterLink>
          </div>
        </div>
      </div>
    </section>

    <!-- ══════════════════════════════════════════
         TESTIMONIALS / SOCIAL PROOF
    ════════════════════════════════════════════ -->
    <section class="testi-section">
      <div class="section-wrap">
        <div class="section-header">
          <div class="section-tag">STORIES</div>
          <h2 class="section-h2">คนที่ปล่อยของไปแล้ว</h2>
        </div>
        <div class="testi-grid">
          <div class="testi-card" v-for="t in testimonials" :key="t.name">
            <div class="testi-stars">★★★★★</div>
            <p class="testi-quote">{{ t.quote }}</p>
            <div class="testi-author">
              <div class="testi-av" :style="{ background: t.color }">{{ t.initials }}</div>
              <div>
                <div class="testi-name">{{ t.name }}</div>
                <div class="testi-role">{{ t.role }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- ══════════════════════════════════════════
         CTA BANNER
    ════════════════════════════════════════════ -->
    <section class="cta-section">
      <div class="cta-mesh">
        <div class="cta-orb cta-orb-1"></div>
        <div class="cta-orb cta-orb-2"></div>
      </div>
      <div class="cta-inner">
        <div class="cta-badge">🚀 เริ่มได้ฟรี · ไม่ต้องใส่บัตรเครดิต</div>
        <h2 class="cta-h2">พร้อมปล่อยของ<br/><span class="grad-text">แล้วหรือยัง?</span></h2>
        <p class="cta-body">เข้าร่วมกับนักศึกษาและ First Jobber กว่า 2,400 คน<br/>ที่เลือก PloyKong เป็น Portfolio Builder</p>
        <div class="cta-actions">
          <RouterLink to="/register" class="btn-cta-hero">
            <Rocket :size="16" />
            เริ่มสร้างพอร์ตฟรี — ใช้ได้เลย
          </RouterLink>
        </div>
      </div>
    </section>

    <!-- ══════════════════════════════════════════
         FOOTER
    ════════════════════════════════════════════ -->
    <footer class="footer">
      <div class="footer-inner">
        <div class="footer-brand">
          <div class="logo-mark sm">
            <img src="/favicon.png" alt="PloyKong Logo" class="logo-img" />
          </div>
          <span class="footer-name">Ploy<em>Kong</em></span>
          <span class="footer-slogan">— มีดีต้องปล่อยของ</span>
        </div>
        <div class="footer-links">
          <a href="#">Privacy</a>
          <a href="#">Terms</a>
          <a href="https://github.com/WarisaTT" target="_blank">GitHub</a>
          <a href="#">Contact</a>
        </div>
        <div class="footer-copy">
          © 2026 PloyKong · Built by Garfairdummn
        </div>
      </div>
    </footer>

  </div>
</template>

<script setup>
import LandingDemo from '@/components/LandingDemo.vue'
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { RouterLink } from 'vue-router'
import { useThemeStore } from '@/stores/theme'
import ThemeToggle from '@/components/ThemeToggle.vue'

import {
  ArrowRight, Rocket, Play, Sparkles, TrendingUp, Bell,
  Lock as Lock2, Menu, X, Check
} from 'lucide-vue-next'

const themeStore = useThemeStore()
const isDark = computed(() => themeStore.resolvedDark)

// Navbar scroll
const scrolled = ref(false)
const mobileOpen = ref(false)
const annual = ref(false)

function onScroll() { scrolled.value = window.scrollY > 20 }
onMounted(() => window.addEventListener('scroll', onScroll))
onUnmounted(() => window.removeEventListener('scroll', onScroll))

const avColors = ['linear-gradient(135deg,#7C6FE0,#E879F9)', 'linear-gradient(135deg,#22D3EE,#7C6FE0)', 'linear-gradient(135deg,#E879F9,#F4C842)', 'linear-gradient(135deg,#34D399,#22D3EE)']
const avNames = ['NK', 'PP', 'AT', 'MJ']

const heroStats = [
  { num: '< 5 นาที', label: 'สร้างพอร์ตเสร็จ' },
  { num: '48+', label: 'Templates' },
  { num: '100%', label: 'No-Code' },
  { num: '2,400+', label: 'ผู้ใช้งาน' },
]

const skills = [
  { name: 'Go / Fiber', pct: 92 },
  { name: 'Vue.js 3', pct: 85 },
  { name: 'Flutter', pct: 78 },
]
const projects = ['PloyKong', 'E-Commerce API', 'Mobile Dashboard']
const chartData = [40, 60, 45, 80, 65, 90, 75, 95, 70, 85, 60, 100]

const techs = ['Go · Fiber', 'Vue.js 3', 'Flutter', 'MySQL', 'Docker', 'OpenAI API', 'Vercel', 'JWT Auth', 'Wildcard DNS']

const steps = [
  { icon: '🎭', title: 'เลือก Template', desc: 'เลือกสาย Programmer, Designer หรือ First Jobber จาก 48 template' },
  { icon: '✨', title: 'เพิ่มเนื้อหาด้วย AI', desc: 'ลากวาง Block ตามต้องการ แล้วให้ AI ช่วยเกลาข้อความให้ดูเป็น Pro' },
  { icon: '🚀', title: 'Publish & แชร์', desc: 'กด Publish รับ URL สวย ส่งให้ HR ได้เลยทันที อัปเดตได้ตลอด' },
]

const plans = [
  {
    name: 'Free',
    price: '0',
    priceAnnual: '0',
    desc: 'เริ่มต้นฟรี ไม่มีค่าใช้จ่าย',
    features: ['1 Portfolio', 'Template พื้นฐาน', 'Analytics 7 วัน', 'URL: ploykong.com/p/slug'],
    cta: 'เริ่มใช้ฟรี',
    featured: false,
  },
  {
    name: 'Pro',
    price: '149',
    priceAnnual: '99',
    desc: 'สำหรับคนที่จริงจังกับการสมัครงาน',
    features: ['Portfolios ไม่จำกัด', 'AI Content Writer', 'Custom Domain', 'Analytics ไม่จำกัด', 'PDF Export', 'AI Interview Coach', 'Mobile App Dashboard', 'Priority Support'],
    cta: 'เริ่ม Pro ทันที',
    featured: true,
  },
  {
    name: 'Team',
    price: '499',
    priceAnnual: '349',
    desc: 'สำหรับกลุ่ม บริษัท หรือสถาบัน',
    features: ['ทุกอย่างใน Pro', 'สูงสุด 10 สมาชิก', 'Team Analytics', 'White-label', 'Custom Branding'],
    cta: 'ติดต่อทีม',
    featured: false,
  },
]

const testimonials = [
  {
    quote: 'สมัครงานได้เร็วขึ้นมาก HR ทักมาเลยว่าพอร์ตสวยมาก ใช้เวลาสร้างแค่ 10 นาที',
    name: 'Thanaphon H.',
    role: 'Car Engineer @ Grab',
    initials: 'NK',
    color: 'linear-gradient(135deg,#7C6FE0,#E879F9)',
  },
  {
    quote: 'AI ช่วยเขียนได้ดีมาก ข้อความที่ออกมาดูเป็น Professional มากกว่าที่เขียนเองอีก',
    name: 'Inrita W.',
    role: 'Data Engineer @ Siam Piwat',
    initials: 'IW',
    color: 'linear-gradient(135deg,#E879F9,#F4C842)',
  },
  {
    quote: 'ระบบ Analytics ทำให้รู้ว่า HR ดูพอร์ตนานแค่ไหน ปรับปรุงได้ตรงจุดมากๆ',
    name: 'Promptpol P.',
    role: 'Frontend Develop @ CP-ALL',
    initials: 'PP',
    color: 'linear-gradient(135deg,#22D3EE,#34D399)',
  },
]
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Noto+Sans+Thai:wght@400;500;600;700;800&display=swap');

:root {
  --font-display: "Syne", "Anuphan", "Noto Sans Thai", sans-serif;
  --font-body: "Inter", "Anuphan", "Noto Sans Thai", sans-serif;
}

/* ══════════════════════════════════════════
   BASE
════════════════════════════════════════════ */
.landing {
  min-height: 100vh;
  font-family: var(--font-body);    
  background: var(--bg);
  color: var(--text);
  line-height: 1.6;
}

.grad-text {
  background: var(--grad-pk);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

/* Marquee Styles */
.marquee-strip {
  border-top: 1px solid var(--border);
  border-bottom: 1px solid var(--border);
  background: var(--surface);
  padding: 12px 0;
  display: flex;
  align-items: center;
  overflow: hidden;
  white-space: nowrap;
}

.marquee-label {
  padding: 0 24px;
  font-size: 11px;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 2px;
  color: var(--text-3);
  background: var(--surface);
  z-index: 2;
  position: relative;
}

.marquee-track {
  display: flex;
  overflow: hidden;
  mask-image: linear-gradient(to right, transparent, black 10%, black 90%, transparent);
}

.marquee-inner {
  display: flex;
  gap: 32px;
  animation: marquee 30s linear infinite;
  padding-left: 32px;
}

.marquee-item {
  font-size: 13px;
  font-weight: 600;
  color: var(--text-2);
}

@keyframes marquee {
  0% { transform: translateX(0); }
  100% { transform: translateX(-50%); }
}

@media (max-width: 640px) {
  .marquee-label { display: none; }
  .marquee-inner { animation-duration: 20s; }
}

/* ══════════════════════════════════════════
   NAVBAR
════════════════════════════════════════════ */
.navbar {
  position: fixed; top: 0; left: 0; right: 0; z-index: 100;
  padding: 0 24px;
  transition: all 0.3s;
}
.navbar.scrolled {
  background: var(--surface);
  border-bottom: 1px solid var(--border);
  backdrop-filter: blur(20px);
  box-shadow: var(--shadow-sm);
}

.nav-inner {
  max-width: 1200px; margin: 0 auto;
  display: flex; align-items: center; height: 64px; gap: 16px;
}

.nav-logo {
  display: flex; align-items: center; gap: 10px;
  text-decoration: none; color: var(--text);
  margin-right: 32px;
}
.logo-mark {
  width: 34px; height: 34px; border-radius: 10px;
  background: var(--surface-2); border: 1px solid var(--border-2);
  display: flex; align-items: center; justify-content: center;
  box-shadow: 0 0 0 0 var(--accent-glow);
  transition: box-shadow 0.3s;
}
.logo-mark:hover { box-shadow: 0 0 16px var(--accent-glow); }
.logo-mark.sm { width: 26px; height: 26px; border-radius: 7px; }

.logo-text {
  font-family: 'Syne', sans-serif; font-size: 18px; font-weight: 800;
  letter-spacing: -0.3px;
}

.logo-text em{ font-style: normal; background: var(--grad-pk); -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text; }

.nav-center {
  display: flex; align-items: center; gap: 6px; flex: 1;
}
.nav-link {
  padding: 6px 14px; border-radius: 8px; font-size: 14px; font-weight: 500;
  color: var(--text-2); text-decoration: none;
  transition: all 0.15s;
}
.nav-link:hover { background: var(--surface-2); color: var(--text); }

.nav-right {
  display: flex; align-items: center; gap: 10px; margin-left: auto;
}

.btn-ghost {
  padding: 8px 16px; border-radius: 10px; font-size: 14px; font-weight: 600;
  color: var(--text-2); text-decoration: none;
  transition: all 0.15s;
}
.btn-ghost:hover { color: var(--text); }

.btn-cta {
  display: inline-flex; align-items: center; gap: 6px;
  padding: 9px 18px; border-radius: 10px; font-size: 14px; font-weight: 700;
  background: var(--grad-primary); color: #fff; text-decoration: none;
  box-shadow: 0 4px 16px var(--accent-glow);
  transition: all 0.2s;
}
.btn-cta:hover { transform: translateY(-1px); box-shadow: 0 8px 24px var(--accent-glow); }
.w-full { width: 100%; justify-content: center; }

.hamburger {
  display: none; align-items: center; justify-content: center;
  width: 40px; height: 40px; border-radius: 10px;
  background: var(--surface-2); border: 1px solid var(--border);
  color: var(--text); cursor: pointer; margin-left: auto;
}
@media (max-width: 768px) {
  .hamburger { display: flex; }
  .nav-center, .nav-right { display: none; }
}

.mobile-menu {
  padding: 16px; border-top: 1px solid var(--border);
  background: var(--surface); display: flex; flex-direction: column; gap: 8px;
}
.mobile-link {
  padding: 12px 16px; border-radius: 10px; font-size: 15px; font-weight: 500;
  color: var(--text-2); text-decoration: none;
  transition: all 0.15s;
}
.mobile-link:hover { background: var(--surface-2); color: var(--text); }
.mobile-divider { height: 1px; background: var(--border); margin: 4px 0; }

.mobile-menu-enter-active, .mobile-menu-leave-active { transition: all 0.2s ease; }
.mobile-menu-enter-from, .mobile-menu-leave-to { opacity: 0; transform: translateY(-10px); }

/* ══════════════════════════════════════════
   HERO
════════════════════════════════════════════ */
.hero {
  position: relative; min-height: 100vh;
  display: flex; flex-direction: column; justify-content: center;
  padding: 100px 24px 60px; overflow: hidden;
}

.hero-mesh {
  position: absolute; inset: 0; pointer-events: none;
}
.mesh-orb {
  position: absolute; border-radius: 50%;
  filter: blur(80px); opacity: 0.6;
}
.orb-1 {
  width: 600px; height: 600px;
  background: radial-gradient(circle, var(--accent-glow), transparent 70%);
  top: -100px; left: -100px;
  animation: orbFloat 20s ease-in-out infinite alternate;
}
.orb-2 {
  width: 500px; height: 500px;
  background: radial-gradient(circle, rgba(232,121,249,0.12), transparent 70%);
  bottom: -80px; right: -60px;
  animation: orbFloat 16s ease-in-out infinite alternate-reverse;
}
.orb-3 {
  width: 350px; height: 350px;
  background: radial-gradient(circle, rgba(34,211,238,0.08), transparent 70%);
  top: 40%; right: 20%;
  animation: orbFloat 24s ease-in-out infinite alternate;
}

.grid-overlay {
  position: absolute; inset: 0;
  background-image:
    linear-gradient(var(--border) 1px, transparent 1px),
    linear-gradient(90deg, var(--border) 1px, transparent 1px);
  background-size: 50px 50px;
  mask-image: radial-gradient(ellipse 80% 80% at 50% 40%, black 20%, transparent 80%);
}

@keyframes orbFloat {
  from { transform: translate(0,0) scale(1); }
  to   { transform: translate(40px, 30px) scale(1.15); }
}

.hero-inner {
  position: relative; z-index: 1;
  max-width: 1200px; margin: 0 auto; width: 100%;
  display: grid; grid-template-columns: 1fr 1fr; gap: 60px; align-items: center;
}
@media (max-width: 960px) {
  .hero-inner { grid-template-columns: 1fr; gap: 40px; text-align: center; }
  .hero-content { align-items: center; }
  .hero-body { margin: 0 auto; }
  .hero-actions { justify-content: center; }
  .hero-visual { max-width: 500px; margin: 0 auto; }
}
@media (max-width: 480px) {
  .hero { padding-top: 80px; }
  .hero-h1 { font-size: 38px; }
  .h1-sub { font-size: 24px; }
  .hero-actions { flex-direction: column; width: 100%; }
  .btn-hero-primary, .btn-hero-ghost { width: 100%; justify-content: center; }
}

.hero-content { display: flex; flex-direction: column; gap: 24px; }

.hero-pill {
  display: inline-flex; align-items: center; gap: 8px;
  padding: 7px 16px; border-radius: 100px; width: fit-content;
  background: var(--accent-subtle); border: 1px solid var(--accent-glow);
  font-size: 12px; font-weight: 600; color: var(--accent-2);
  letter-spacing: 0.3px;
}
.pill-dot {
  width: 7px; height: 7px; border-radius: 50%;
  background: var(--accent-2);
  box-shadow: 0 0 8px var(--accent-2);
  animation: pulse 2s ease-in-out infinite;
}
@keyframes pulse { 0%,100%{ box-shadow: 0 0 6px var(--accent-2); } 50%{ box-shadow: 0 0 16px var(--accent-2); } }

.hero-h1 {
  font-family: 'Syne', sans-serif;
  font-size: clamp(44px, 5vw, 72px);
  font-weight: 900; line-height: 1.05; letter-spacing: -2px;
  display: flex; flex-direction: column; gap: 4px;
}
.h1-accent {
  background: var(--grad-pk);
  -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text;
}
.h1-sub {
  font-size: clamp(28px, 3vw, 44px);
  color: var(--text-2);
}

.hero-body {
  font-size: 16px; color: var(--text-2); line-height: 1.8;
  font-weight: 300; max-width: 480px;
}

.hero-actions {
  display: flex; align-items: center; gap: 14px; flex-wrap: wrap;
}

.btn-hero-primary {
  display: inline-flex; align-items: center; gap: 8px;
  padding: 14px 28px; border-radius: 14px; font-size: 15px; font-weight: 700;
  background: var(--grad-primary); color: #fff; text-decoration: none;
  box-shadow: 0 8px 32px var(--accent-glow);
  transition: all 0.25s;
}
.btn-hero-primary:hover { transform: translateY(-2px); box-shadow: 0 14px 40px var(--accent-glow); }

.btn-hero-ghost {
  display: inline-flex; align-items: center; gap: 8px;
  padding: 13px 22px; border-radius: 14px; font-size: 15px; font-weight: 600;
  background: var(--surface-2); border: 1px solid var(--border-2);
  color: var(--text-2); text-decoration: none;
  transition: all 0.2s;
}
.btn-hero-ghost:hover { background: var(--surface-3); color: var(--text); }

.hero-proof {
  display: flex; align-items: center; gap: 14px;
}
.proof-avatars {
  display: flex;
}
.proof-av {
  width: 34px; height: 34px; border-radius: 50%;
  font-size: 11px; font-weight: 700; color: #fff;
  display: flex; align-items: center; justify-content: center;
  border: 2px solid var(--bg);
  margin-left: -8px;
}
.proof-av:first-child { margin-left: 0; }
.proof-text { font-size: 13px; color: var(--text-2); }
.proof-num { font-weight: 700; color: var(--text); }

/* Browser mockup */
.hero-visual { position: relative; }

.browser-frame {
  background: var(--surface); border: 1px solid var(--border-2);
  border-radius: 16px; overflow: hidden;
  box-shadow: var(--shadow-lg), 0 0 0 1px var(--border);
}
.browser-bar {
  display: flex; align-items: center; gap: 12px;
  padding: 12px 16px; border-bottom: 1px solid var(--border);
  background: var(--surface-2);
}
.browser-dots {
  display: flex; gap: 5px;
}
.browser-dots span {
  width: 10px; height: 10px; border-radius: 50%;
  background: var(--border-2);
}
.browser-dots span:nth-child(1) { background: #FF5F57; }
.browser-dots span:nth-child(2) { background: #FFBD2E; }
.browser-dots span:nth-child(3) { background: #28CA41; }

.browser-url {
  display: flex; align-items: center; gap: 6px;
  padding: 5px 12px; border-radius: 8px; flex: 1; max-width: 220px;
  background: var(--surface); border: 1px solid var(--border);
  font-size: 11px; color: var(--text-2); font-family: monospace;
}

.browser-body { padding: 20px; }

.mini-portfolio { display: flex; flex-direction: column; gap: 16px; }

.mp-hero {
  display: flex; align-items: center; gap: 12px;
  padding: 12px; border-radius: 12px; background: var(--surface-2);
}
.mp-avatar {
  width: 42px; height: 42px; border-radius: 50%;
  background: var(--grad-pk);
  display: flex; align-items: center; justify-content: center;
  font-size: 14px; font-weight: 800; color: #fff; flex-shrink: 0;
}
.mp-name { font-size: 14px; font-weight: 700; }
.mp-role { font-size: 11px; color: var(--text-2); }
.mp-hire {
  margin-left: auto; padding: 5px 12px; border-radius: 20px; flex-shrink: 0;
  background: var(--grad-primary); color: #fff; font-size: 11px; font-weight: 700;
}

.mp-skills { display: flex; flex-direction: column; gap: 8px; }
.mp-skill-bar { display: flex; align-items: center; gap: 10px; }
.mp-skill-name { font-size: 11px; color: var(--text-2); min-width: 70px; flex-shrink: 0; }
.mp-skill-track {
  flex: 1; height: 5px; border-radius: 3px; background: var(--surface-3);
}
.mp-skill-fill {
  height: 100%; border-radius: 3px;
  background: var(--grad-pk);
  transition: width 1s ease;
}

.mp-projects { display: flex; flex-direction: column; gap: 6px; }
.mp-proj { display: flex; align-items: center; gap: 8px; }
.mp-proj-dot {
  width: 6px; height: 6px; border-radius: 50%; flex-shrink: 0;
  background: var(--accent);
}
.mp-proj-name { font-size: 12px; color: var(--text-2); }

/* Floating cards */
.float-card {
  position: absolute;
  background: var(--surface); border: 1px solid var(--border-2);
  border-radius: 12px; padding: 10px 14px;
  display: flex; align-items: center; gap: 10px;
  box-shadow: var(--shadow-md);
  font-size: 12px;
}
.fc-icon { color: var(--accent-2); flex-shrink: 0; }
.fc-label { font-size: 10px; font-weight: 700; color: var(--text-2); }
.fc-val { font-size: 12px; color: var(--text); font-weight: 600; }

.fc-ai {
  bottom: 30px; left: -60px;
  animation: floatAnim 4s ease-in-out infinite;
}
.fc-stat {
  top: 40px; right: -40px;
  animation: floatAnim 5s ease-in-out infinite 1s;
}
.fc-notify {
  top: -20px; left: 10%;
  animation: floatAnim 3.5s ease-in-out infinite 0.5s;
}

@keyframes floatAnim {
  0%,100%{ transform: translateY(0); }
  50%{ transform: translateY(-8px); }
}

/* Stats bar */
.hero-stats {
  position: relative; z-index: 1;
  max-width: 1200px; margin: 48px auto 0;
  display: grid; grid-template-columns: repeat(4, 1fr);
  gap: 1px; background: var(--border);
  border: 1px solid var(--border); border-radius: 16px; overflow: hidden;
}
@media (max-width: 600px) {
  .hero-stats { grid-template-columns: repeat(2, 1fr); margin-top: 32px; }
  .hstat { padding: 12px; }
  .hstat-num { font-size: 18px; }
}
@media (max-width: 400px) {
  .hero-stats { grid-template-columns: 1fr !important; }
}

.hstat {
  display: flex; flex-direction: column; align-items: center;
  padding: 20px 16px; gap: 4px;
  background: var(--surface);
}
.hstat-num {
  font-family: 'Syne', sans-serif;
  font-size: 22px; font-weight: 800;
  background: var(--grad-pk);
  -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text;
}
.hstat-label { font-size: 12px; color: var(--text-2); }

/* ══════════════════════════════════════════
   MARQUEE
════════════════════════════════════════════ */
.marquee-strip {
  border-top: 1px solid var(--border); border-bottom: 1px solid var(--border);
  padding: 14px 0; background: var(--surface-2);
  display: flex; align-items: center; gap: 24px; overflow: hidden;
}
.marquee-label {
  font-size: 11px; font-weight: 700; letter-spacing: 2px; text-transform: uppercase;
  color: var(--text-3); padding: 0 24px; white-space: nowrap;
}
.marquee-track { flex: 1; overflow: hidden; }
.marquee-inner {
  display: flex; gap: 32px; white-space: nowrap;
  animation: marqueeScroll 20s linear infinite;
}
.marquee-item {
  font-size: 13px; font-weight: 600; color: var(--text-2);
  padding: 4px 16px; border-radius: 100px;
  background: var(--surface-3); border: 1px solid var(--border);
  white-space: nowrap;
}

@keyframes marqueeScroll {
  from { transform: translateX(0); }
  to { transform: translateX(-50%); }
}

/* ══════════════════════════════════════════
   SECTION SHARED
════════════════════════════════════════════ */
.section-wrap { 
  width: 100%; 
  max-width: 1200px; 
  margin: 0 auto; 
  padding: 100px 0; 
}

@media (max-width: 768px) {
  .section-wrap { width: 90%; padding: 60px 0; }
}
.section-header { text-align: center; margin-bottom: 60px; }
.section-tag {
  font-size: 11px; font-weight: 700; letter-spacing: 3px; text-transform: uppercase;
  color: var(--accent-2); margin-bottom: 14px;
}
.section-h2 {
  font-family: 'Syne', sans-serif;
  font-size: clamp(28px, 4vw, 44px); font-weight: 800; letter-spacing: -1px;
  margin-bottom: 16px;
}
.section-sub { font-size: 15px; color: var(--text-2); max-width: 520px; margin: 0 auto; line-height: 1.7; }

/* ══════════════════════════════════════════
   FEATURES BENTO
════════════════════════════════════════════ */
.features-section { background: var(--bg-2); }

.feat-bento {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-template-rows: auto auto;
  gap: 16px;
}

@media (max-width: 900px) { 
  .feat-bento { grid-template-columns: repeat(2, 1fr); } 
  .bento-lg, .bento-md { grid-column: span 2; }
}
@media (max-width: 600px) { 
  .feat-bento { grid-template-columns: 1fr; }
  .bento-lg, .bento-md, .bento-sm { grid-column: span 1; }
  .section-wrap { padding: 60px 20px; }
}

.bento-card {
  position: relative; overflow: hidden;
  background: var(--surface); border: 1px solid var(--border);
  border-radius: 20px; padding: 28px;
  transition: all 0.3s;
}
.bento-card:hover { border-color: var(--border-2); transform: translateY(-3px); box-shadow: var(--shadow-md); }

.bento-lg { grid-column: span 2; grid-row: span 1; }
.bento-md { grid-column: span 2; }
.bento-sm { grid-column: span 1; }

.bento-glow {
  position: absolute; top: -60px; right: -60px;
  width: 200px; height: 200px; border-radius: 50%;
  background: radial-gradient(circle, var(--accent-glow), transparent 70%);
  pointer-events: none;
}
.bento-glow.pink { background: radial-gradient(circle, rgba(232,121,249,0.15), transparent 70%); }
.bento-glow.cyan { background: radial-gradient(circle, rgba(34,211,238,0.12), transparent 70%); }
.bento-glow.gold { background: radial-gradient(circle, rgba(244,200,66,0.12), transparent 70%); }

.feat-tag {
  display: inline-block; padding: 4px 10px; border-radius: 6px; margin-bottom: 12px;
  font-size: 9px; font-weight: 700; letter-spacing: 2px; text-transform: uppercase;
  background: var(--accent-subtle); color: var(--accent-2); border: 1px solid var(--accent-glow);
}
.feat-tag.pink { background: rgba(232,121,249,0.08); color: #E879F9; border-color: rgba(232,121,249,0.2); }
.feat-tag.cyan { background: rgba(34,211,238,0.08); color: var(--cyan); border-color: rgba(34,211,238,0.2); }
.feat-tag.gold { background: rgba(244,200,66,0.08); color: var(--gold); border-color: rgba(244,200,66,0.2); }

.bento-h3 { font-size: 20px; font-weight: 700; margin-bottom: 8px; }
.bento-p { font-size: 13px; color: var(--text-2); line-height: 1.7; }

/* Drag demo */
.drag-demo { margin-top: 20px; display: flex; flex-direction: column; gap: 8px; max-width: 320px; }
.drag-block {
  padding: 10px 14px; border-radius: 10px; font-size: 13px;
  background: var(--surface-2); border: 1px solid var(--border);
  cursor: grab; user-select: none; transition: all 0.2s;
}
.drag-block.dragging {
  border-color: var(--accent); background: var(--accent-subtle);
  box-shadow: 0 4px 20px var(--accent-glow);
  transform: rotate(-1deg) scale(1.02);
}
.drag-handle { float: right; color: var(--text-3); cursor: grab; }

/* AI bubble */
.ai-bubble { margin-top: 16px; display: flex; flex-direction: column; gap: 8px; }
.ai-in, .ai-out {
  padding: 8px 12px; border-radius: 10px; font-size: 11px; line-height: 1.5;
}
.ai-in { background: var(--surface-2); color: var(--text-2); border: 1px solid var(--border); }
.ai-out { background: var(--accent-subtle); color: var(--accent-2); border: 1px solid var(--accent-glow); }
.ai-arrow { font-size: 11px; color: var(--accent-2); text-align: center; font-weight: 700; }

/* Analytics bars */
.analytics-bars {
  margin-top: 20px; display: flex; align-items: flex-end; gap: 4px; height: 70px;
}
.analytics-bar {
  flex: 1; border-radius: 4px 4px 0 0; min-height: 15%;
  background: var(--grad-pk); opacity: 0.7;
  transition: opacity 0.2s;
}
.analytics-bar:hover { opacity: 1; }

/* Lock demo */
.lock-demo {
  margin-top: 20px; display: flex; flex-direction: column; align-items: center; gap: 8px;
}
.lock-icon { color: var(--gold); }
.lock-text { font-size: 12px; color: var(--text-2); font-weight: 600; }
.lock-timer {
  font-size: 11px; padding: 4px 12px; border-radius: 20px;
  background: rgba(244,200,66,0.1); color: var(--gold);
  border: 1px solid rgba(244,200,66,0.2);
}

/* Chat demo */
.chat-demo { margin-top: 16px; display: flex; flex-direction: column; gap: 10px; }
.chat-msg { display: flex; gap: 8px; align-items: flex-start; }
.chat-msg.ai { flex-direction: row-reverse; }
.chat-avatar {
  width: 28px; height: 28px; border-radius: 50%; flex-shrink: 0;
  display: flex; align-items: center; justify-content: center; font-size: 11px; font-weight: 700;
}
.hr-av { background: var(--surface-3); color: var(--text-2); }
.ai-av { background: var(--accent-subtle); font-size: 14px; }
.chat-bubble {
  padding: 8px 12px; border-radius: 12px; font-size: 11px; line-height: 1.5;
  max-width: 85%;
}
.chat-msg.hr .chat-bubble { background: var(--surface-2); color: var(--text); }
.chat-msg.ai .chat-bubble { background: var(--accent-subtle); color: var(--accent-2); }

/* URL demo */
.url-demo {
  margin-top: 16px; padding: 10px 16px; border-radius: 10px;
  background: var(--surface-2); border: 1px solid var(--border);
  font-family: monospace; font-size: 13px;
}
.url-base { color: var(--text-2); }
.url-slug { color: var(--accent-2); font-weight: 700; }

/* ══════════════════════════════════════════
   HOW IT WORKS
════════════════════════════════════════════ */
.steps-grid {
  display: grid; grid-template-columns: repeat(3, 1fr); gap: 32px;
  position: relative;
}
@media (max-width: 768px) { .steps-grid { grid-template-columns: 1fr; } }

.step-card {
  display: flex; flex-direction: column; align-items: center; text-align: center;
  position: relative;
}
.step-number {
  font-family: 'Syne', sans-serif; font-size: 60px; font-weight: 900; line-height: 1;
  background: var(--grad-pk);
  -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text;
  opacity: 0.2; margin-bottom: -10px;
}
.step-icon-wrap {
  width: 64px; height: 64px; border-radius: 18px;
  background: var(--surface-2); border: 1px solid var(--border-2);
  display: flex; align-items: center; justify-content: center;
  font-size: 28px; margin-bottom: 16px;
  box-shadow: var(--shadow-sm);
}
.step-title { font-size: 18px; font-weight: 700; margin-bottom: 10px; }
.step-desc { font-size: 13px; color: var(--text-2); line-height: 1.7; }
.step-connector {
  position: absolute; top: 90px; right: -16px;
  width: 32px; height: 2px; background: var(--border-2);
}

.timeline {
  margin-top: 48px; display: flex; flex-direction: column; gap: 12px;
}
.timeline-bar {
  height: 3px; background: var(--surface-3); border-radius: 2px; overflow: hidden;
}
.timeline-fill {
  height: 100%; width: 100%; background: var(--grad-pk);
  border-radius: 2px;
  animation: fillBar 3s ease-in-out infinite alternate;
}
@keyframes fillBar { from { width: 33%; } to { width: 100%; } }

.timeline-labels {
  display: flex; justify-content: space-between;
  font-size: 12px; color: var(--text-3);
}

/* ══════════════════════════════════════════
   PRICING
════════════════════════════════════════════ */
.pricing-section { background: var(--bg-2); }

.billing-toggle {
  display: inline-flex; align-items: center; gap: 12px; margin-top: 20px;
  font-size: 14px; font-weight: 600;
}
.billing-toggle span { color: var(--text-3); transition: color 0.2s; }
.billing-toggle span.active { color: var(--text); }

.toggle-track {
  width: 46px; height: 26px; border-radius: 100px;
  background: var(--surface-3); border: 1px solid var(--border-2);
  cursor: pointer; position: relative; transition: background 0.2s;
}
.toggle-track.on { background: var(--accent); }
.toggle-thumb {
  position: absolute; top: 3px; left: 3px;
  width: 18px; height: 18px; border-radius: 50%; background: #fff;
  transition: transform 0.2s; box-shadow: 0 2px 4px rgba(0,0,0,0.2);
}
.toggle-track.on .toggle-thumb { transform: translateX(20px); }

.save-badge {
  font-size: 10px; padding: 2px 8px; border-radius: 6px;
  background: rgba(52,211,153,0.12); color: var(--success);
  border: 1px solid rgba(52,211,153,0.2); margin-left: 4px;
}

.pricing-grid {
  display: grid; grid-template-columns: repeat(3, 1fr); gap: 20px;
}
@media (max-width: 900px) { .pricing-grid { grid-template-columns: 1fr; max-width: 400px; margin: 0 auto; } }

.price-card {
  background: var(--surface); border: 1px solid var(--border);
  border-radius: 24px; padding: 32px 28px;
  position: relative; transition: all 0.3s;
  display: flex; flex-direction: column;
}
.price-card:hover { transform: translateY(-4px); box-shadow: var(--shadow-md); }
.price-card.featured {
  border-color: var(--accent); background: var(--surface-2);
  box-shadow: 0 0 0 1px var(--accent-glow), var(--shadow-md);
}

.price-badge {
  position: absolute; top: -14px; left: 50%; transform: translateX(-50%);
  padding: 5px 16px; border-radius: 100px;
  background: var(--grad-primary); color: #fff;
  font-size: 12px; font-weight: 700; white-space: nowrap;
  box-shadow: 0 4px 12px var(--accent-glow);
}

.price-header { margin-bottom: 24px; }
.price-name {
  font-size: 12px; font-weight: 700; letter-spacing: 2px; text-transform: uppercase;
  color: var(--text-2); margin-bottom: 14px;
}
.price-amount { display: flex; align-items: baseline; gap: 4px; margin-bottom: 8px; }
.price-num { font-family: 'Syne', 'Anuphan', sans-serif; font-size: 48px; font-weight: 700; line-height: 1; }
.price-per { font-size: 14px; color: var(--text-2); }
.price-desc { font-size: 13px; color: var(--text-2); }

.price-features {
  list-style: none; display: flex; flex-direction: column; gap: 10px; flex: 1;
  margin-bottom: 28px;
}
.price-features li {
  display: flex; align-items: center; gap: 8px; font-size: 13px; color: var(--text-2);
}
.feat-check { color: var(--success); flex-shrink: 0; }

.price-btn {
  display: flex; align-items: center; justify-content: center; gap: 8px;
  padding: 12px; border-radius: 12px; font-size: 14px; font-weight: 700;
  background: var(--surface-2); border: 1px solid var(--border-2);
  color: var(--text); text-decoration: none;
  transition: all 0.2s;
}
.price-btn:hover { background: var(--surface-3); }
.price-btn-primary {
  background: var(--grad-primary); color: #fff; border-color: transparent;
  box-shadow: 0 4px 16px var(--accent-glow);
}
.price-btn-primary:hover { box-shadow: 0 8px 24px var(--accent-glow); background: var(--grad-primary); }

/* ══════════════════════════════════════════
   TESTIMONIALS
════════════════════════════════════════════ */
.testi-grid {
  display: grid; grid-template-columns: repeat(3, 1fr); gap: 20px;
}
@media (max-width: 900px) { .testi-grid { grid-template-columns: 1fr; max-width: 500px; margin: 0 auto; } }

.testi-card {
  background: var(--surface); border: 1px solid var(--border);
  border-radius: 20px; padding: 28px;
  display: flex; flex-direction: column; gap: 16px;
  transition: all 0.3s;
}
.testi-card:hover { border-color: var(--border-2); transform: translateY(-3px); }

.testi-stars { color: var(--gold); font-size: 14px; letter-spacing: 2px; }
.testi-quote { font-size: 14px; color: var(--text-2); line-height: 1.7; flex: 1; }
.testi-author { display: flex; align-items: center; gap: 12px; }
.testi-av {
  width: 40px; height: 40px; border-radius: 50%;
  display: flex; align-items: center; justify-content: center;
  font-size: 14px; font-weight: 700; color: #fff; flex-shrink: 0;
}
.testi-name { font-size: 14px; font-weight: 700; }
.testi-role { font-size: 12px; color: var(--text-2); }

/* ══════════════════════════════════════════
   CTA
════════════════════════════════════════════ */
.cta-section {
  position: relative; padding: 120px 24px; text-align: center; overflow: hidden;
  background: var(--bg);
}
.cta-mesh { position: absolute; inset: 0; pointer-events: none; }
.cta-orb {
  position: absolute; border-radius: 50%; filter: blur(80px);
}
.cta-orb-1 {
  width: 600px; height: 600px;
  background: radial-gradient(circle, var(--accent-glow), transparent 70%);
  top: 50%; left: 50%; transform: translate(-50%, -50%);
}
.cta-orb-2 {
  width: 400px; height: 400px;
  background: radial-gradient(circle, rgba(232,121,249,0.1), transparent 70%);
  bottom: -100px; right: -50px;
}

.cta-inner { position: relative; z-index: 1; max-width: 600px; margin: 0 auto; }
.cta-badge {
  display: inline-block; padding: 7px 18px; border-radius: 100px;
  background: var(--accent-subtle); border: 1px solid var(--accent-glow);
  font-size: 13px; font-weight: 600; color: var(--accent-2); margin-bottom: 24px;
}
.cta-h2 {
  font-family: 'Syne', sans-serif;
  font-size: clamp(36px, 5vw, 60px); font-weight: 900; letter-spacing: -2px;
  margin-bottom: 16px; line-height: 1.1;
}
.cta-body { font-size: 15px; color: var(--text-2); margin-bottom: 36px; line-height: 1.8; }
.cta-actions { display: flex; justify-content: center; }

.btn-cta-hero {
  display: inline-flex; align-items: center; gap: 10px;
  padding: 16px 36px; border-radius: 16px; font-size: 16px; font-weight: 700;
  background: var(--grad-primary); color: #fff; text-decoration: none;
  box-shadow: 0 8px 40px var(--accent-glow);
  transition: all 0.25s;
}
.btn-cta-hero:hover { transform: translateY(-3px); box-shadow: 0 16px 60px var(--accent-glow); }

/* ══════════════════════════════════════════
   FOOTER
════════════════════════════════════════════ */
.footer {
  border-top: 1px solid var(--border);
  padding: 32px 24px; background: var(--surface);
}
.footer-inner {
  max-width: 1200px; margin: 0 auto;
  display: flex; align-items: center; gap: 24px; flex-wrap: wrap;
}
.footer-brand { display: flex; align-items: center; gap: 8px; margin-right: auto; }
.footer-name { font-family: 'Syne', sans-serif; font-size: 16px; font-weight: 700; }
.footer-slogan { font-size: 13px; color: var(--text-3); }

.footer-links {
  display: flex; gap: 20px;
}
.footer-links a {
  font-size: 13px; color: var(--text-2); text-decoration: none; transition: color 0.15s;
}
.footer-links a:hover { color: var(--text); }
.footer-copy { font-size: 12px; color: var(--text-3); }
</style>
