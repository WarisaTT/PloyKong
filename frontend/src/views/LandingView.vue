<template>
  <div class="landing">
    <!-- Navbar -->
    <nav class="navbar">
      <div class="nav-inner">
        <div class="nav-logo">
          <span class="logo-pk">PK</span>
          <span class="logo-name grad-text">PloyKong</span>
        </div>
        <div class="nav-links">
          <a href="#features" class="nav-link">Features</a>
          <a href="#pricing" class="nav-link">Pricing</a>
          <RouterLink to="/login" class="btn btn-secondary btn-sm">Login</RouterLink>
          <RouterLink to="/register" class="btn btn-primary btn-sm">เริ่มฟรี 🚀</RouterLink>
        </div>
      </div>
    </nav>

    <!-- Hero -->
    <section class="hero-section">
      <div class="hero-glow"></div>
      <div class="hero-content animate-slide-up" v-intersect>
        <div class="hero-badge">✨ No-Code · AI-Powered · Gen-Z Ready</div>
        <h1 class="hero-headline">
          มีดีต้อง<span class="grad-text">ปล่อยของ</span><br />
          สร้างพอร์ตสวยง่ายๆ<br />ในไม่กี่คลิก
        </h1>
        <p class="hero-sub">
          แพลตฟอร์ม No-Code Portfolio Builder สำหรับ Gen-Z Professionals<br />
          ลากวาง ปรับแต่ง เผยแพร่ ในไม่ถึง 5 นาที
        </p>
        <div class="hero-cta">
          <RouterLink to="/register" class="btn btn-primary btn-lg neon-glow">
            🚀 เริ่มสร้างพอร์ตฟรี
          </RouterLink>
          <a href="#demo" class="btn btn-secondary btn-lg">▶ ดู Demo</a>
        </div>
        <div class="hero-stats">
          <div class="hero-stat"><span class="stat-num">5 นาที</span><span class="stat-label">สร้างพอร์ตเสร็จ</span></div>
          <div class="stat-divider"></div>
          <div class="hero-stat"><span class="stat-num">AI</span><span class="stat-label">ช่วยเขียนเนื้อหา</span></div>
          <div class="stat-divider"></div>
          <div class="hero-stat"><span class="stat-num">∞</span><span class="stat-label">Templates ฟรี</span></div>
        </div>
      </div>
      <!-- Floating portfolio cards -->
      <div class="hero-visual" aria-hidden="true">
        <div class="float-card card-1 glass">
          <div class="fc-avatar">NP</div>
          <div><div class="fc-name">Napat K.</div><div class="fc-role">Full-Stack Dev</div></div>
          <div class="fc-live">● Live</div>
        </div>
        <div class="float-card card-2 glass">
          <div class="fc-stat">👁️ 1,234 views this week</div>
          <div class="fc-bar"><div class="fc-bar-fill" style="width:75%"></div></div>
        </div>
        <div class="float-card card-3 glass">
          <div class="fc-ai">🤖 AI Improved!</div>
          <div class="fc-text">"Led development of a high-performance..."</div>
        </div>
      </div>
    </section>

    <!-- Features -->
    <section id="features" class="features-section">
      <div class="section-inner animate-slide-up" v-intersect>
        <div class="section-tag">FEATURES</div>
        <h2 class="section-title">ครบทุกอย่างที่ต้องใช้</h2>
        <div class="features-grid">
          <div v-for="(f, i) in features" :key="f.title" class="feature-card glass animate-slide-up" v-intersect :style="{ transitionDelay: `${i * 50}ms` }">
            <div class="feat-icon">{{ f.icon }}</div>
            <div class="feat-title">{{ f.title }}</div>
            <div class="feat-desc">{{ f.desc }}</div>
          </div>
        </div>
      </div>
    </section>

    <!-- How it works -->
    <section id="demo" class="how-section">
      <div class="section-inner animate-slide-up" v-intersect>
        <div class="section-tag">HOW IT WORKS</div>
        <h2 class="section-title">3 ขั้นตอน สร้างพอร์ตมืออาชีพ</h2>
        <div class="steps-row">
          <div v-for="(step, i) in steps" :key="i" class="how-step animate-slide-up" v-intersect :style="{ transitionDelay: `${i * 100}ms` }">
            <div class="how-num">{{ i + 1 }}</div>
            <div class="how-icon">{{ step.icon }}</div>
            <div class="how-title">{{ step.title }}</div>
            <div class="how-desc">{{ step.desc }}</div>
          </div>
        </div>
      </div>
    </section>

    <!-- Pricing -->
    <section id="pricing" class="pricing-section">
      <div class="section-inner animate-slide-up" v-intersect>
        <div class="section-tag">PRICING</div>
        <h2 class="section-title">ราคาที่เป็นมิตร</h2>
        <div class="pricing-grid">
          <div v-for="(plan, i) in plans" :key="plan.name" class="pricing-card animate-slide-up" :class="{ featured: plan.featured }" v-intersect :style="{ transitionDelay: `${i * 100}ms` }">
            <div v-if="plan.featured" class="plan-badge">⭐ ยอดนิยม</div>
            <div class="plan-name">{{ plan.name }}</div>
            <div class="plan-price">
              <span class="price-num">{{ plan.price }}</span>
              <span class="price-per">{{ plan.per }}</span>
            </div>
            <ul class="plan-features">
              <li v-for="f in plan.features" :key="f">✓ {{ f }}</li>
            </ul>
            <RouterLink to="/register" class="btn" :class="plan.featured ? 'btn-primary' : 'btn-secondary'" style="width:100%">
              เริ่มใช้งาน
            </RouterLink>
          </div>
        </div>
      </div>
    </section>

    <!-- CTA -->
    <section class="cta-section">
      <div class="cta-glow"></div>
      <div class="cta-content animate-slide-up" v-intersect>
        <h2>พร้อมปล่อยของแล้วหรือยัง? 🚀</h2>
        <p>สมัครฟรี ไม่ต้องใส่บัตรเครดิต</p>
        <RouterLink to="/register" class="btn btn-primary btn-lg neon-glow">เริ่มสร้างพอร์ตฟรี</RouterLink>
      </div>
    </section>

    <!-- Footer -->
    <footer class="footer">
      <div class="footer-inner">
        <div class="footer-brand">
          <span class="logo-pk small">PK</span>
          <span style="font-weight:700;">PloyKong</span>
          <span style="color:var(--muted);font-size:13px;">— มีดีต้องปล่อยของ</span>
        </div>
        <div class="footer-links">
          <a href="#">Privacy</a>
          <a href="#">Terms</a>
          <a href="https://github.com" target="_blank">GitHub</a>
        </div>
        <div style="font-size:12px;color:var(--muted)">© 2025 PloyKong · Built by KMUTT IT</div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'

const features = [
  { icon: '🎨', title: 'Drag & Drop Builder', desc: 'ลากวางส่วนประกอบได้อิสระ แก้ไขแล้วเห็นผลทันที (WYSIWYG)' },
  { icon: '🤖', title: 'AI Content Writer', desc: 'ให้ AI ช่วยเกลาข้อความให้ดูเป็นมืออาชีพ วิเคราะห์คะแนน Resume' },
  { icon: '📊', title: 'Analytics Dashboard', desc: 'ดูสถิติผู้เข้าชม แหล่งที่มา และพิกัดประเทศผ่านแอป Flutter' },
  { icon: '🔒', title: 'Privacy Controls', desc: 'ล็อกรหัสผ่านเฉพาะส่วน หรือสร้างลิงก์ชั่วคราวที่หมดอายุเอง' },
  { icon: '⚡', title: 'One-Click Publish', desc: 'อัปเดตเนื้อหาเข้าลิงก์เดิมได้ทันที ไม่ต้อง Deploy ใหม่' },
  { icon: '🌐', title: 'Custom URL', desc: 'ได้ URL เป็นของตัวเอง เช่น yourname.ploykong.com ทันที' },
  { icon: '💬', title: 'AI Interview Coach', desc: 'แชทบอทอัจฉริยะที่ HR ถามข้อมูลเกี่ยวกับคุณได้โดยตรง' },
  { icon: '📱', title: 'Mobile Dashboard', desc: 'รับการแจ้งเตือนมือถือเมื่อมีคนกด Hire Me หรือดาวน์โหลด CV' },
]

const steps = [
  { icon: '🎭', title: 'เลือก Template', desc: 'เลือกสาย Programmer, Designer หรือ First Jobber' },
  { icon: '✍️', title: 'เพิ่มเนื้อหา', desc: 'ลากวาง Block และให้ AI ช่วยเขียนให้' },
  { icon: '🚀', title: 'Publish!', desc: 'กด Publish แล้วแชร์ลิงก์ได้เลย' },
]

const plans = [
  {
    name: 'Free',
    price: '฿0',
    per: '/ เดือน',
    features: ['1 Portfolio', 'Template พื้นฐาน', 'Analytics 7 วัน', 'ploykong.com/p/slug'],
    featured: false,
  },
  {
    name: 'Pro',
    price: '฿149',
    per: '/ เดือน',
    features: ['ไม่จำกัด Portfolios', 'AI Content Writer', 'Custom Domain', 'Analytics ไม่จำกัด', 'PDF Export', 'AI Interview Coach', 'Priority Support'],
    featured: true,
  },
  {
    name: 'Team',
    price: '฿499',
    per: '/ เดือน',
    features: ['ทุกอย่างใน Pro', 'สูงสุด 10 สมาชิก', 'Team Analytics', 'White-label'],
    featured: false,
  },
]
</script>

<style scoped>
/* --- Layout --- */
.landing { min-height: 100vh; }

/* --- Navbar --- */
.navbar { position: fixed; top: 0; left: 0; right: 0; z-index: 100; padding: 0 24px; border-bottom: 1px solid var(--border); background: var(--bg); backdrop-filter: blur(20px); }
.nav-inner { max-width: 1100px; margin: 0 auto; display: flex; align-items: center; height: 60px; gap: 20px; }
.nav-logo { display: flex; align-items: center; gap: 10px; margin-right: auto; text-decoration: none; color: var(--text); }
.logo-pk { width: 32px; height: 32px; border-radius: 8px; background: linear-gradient(135deg, var(--indigo), var(--purple)); display: flex; align-items: center; justify-content: center; font-family: var(--font-display); font-size: 13px; font-weight: 800; color: #fff; flex-shrink: 0; }
.logo-pk.small { width: 24px; height: 24px; font-size: 10px; border-radius: 6px; }
.logo-name { font-family: var(--font-display); font-size: 18px; font-weight: 800; }
.nav-links { display: flex; align-items: center; gap: 16px; }
.nav-link { font-size: 14px; color: var(--muted); text-decoration: none; transition: color 0.15s; }
.nav-link:hover { color: var(--text); }

/* --- Hero --- */
.hero-section { position: relative; min-height: 100vh; display: flex; align-items: center; justify-content: center; padding: 100px 24px 60px; overflow: hidden; }
.hero-glow { position: absolute; top: 10%; left: 50%; transform: translateX(-50%); width: 700px; height: 500px; background: radial-gradient(ellipse, rgba(79,70,229,0.3) 0%, rgba(168,85,247,0.15) 40%, transparent 70%); pointer-events: none; }
.hero-content { position: relative; z-index: 1; text-align: center; max-width: 680px; }
.hero-badge { display: inline-block; padding: 6px 18px; border-radius: 100px; background: rgba(79,70,229,0.12); border: 1px solid rgba(79,70,229,0.3); font-size: 13px; font-weight: 600; color: #818cf8; margin-bottom: 24px; }
.hero-headline { font-size: clamp(38px, 6vw, 68px); font-weight: 800; line-height: 1.15; margin-bottom: 20px; }
.hero-sub { font-size: 16px; color: var(--muted); line-height: 1.8; margin-bottom: 32px; }
.hero-cta { display: flex; gap: 14px; justify-content: center; flex-wrap: wrap; margin-bottom: 40px; }
.hero-stats { display: flex; align-items: center; justify-content: center; gap: 24px; flex-wrap: wrap; }
.hero-stat { display: flex; flex-direction: column; align-items: center; gap: 4px; }
.stat-num { font-family: var(--font-display); font-size: 24px; font-weight: 800; color: var(--neon-cyan); }
.stat-label { font-size: 12px; color: var(--muted); }
.stat-divider { width: 1px; height: 40px; background: var(--border); }

/* Floating cards */
.hero-visual { position: absolute; right: calc(50% - 560px); top: 50%; transform: translateY(-50%); display: none; }
@media (min-width: 1100px) { .hero-visual { display: block; } }
.float-card { position: absolute; border-radius: 14px; padding: 14px 16px; font-size: 13px; animation: float 4s ease-in-out infinite; }
.card-1 { top: -80px; left: 0; width: 200px; display: flex; align-items: center; gap: 10px; animation-delay: 0s; }
.card-2 { top: 30px; left: 30px; width: 220px; animation-delay: 1.5s; }
.card-3 { top: 120px; left: -20px; width: 230px; animation-delay: 0.8s; }
.fc-avatar { width: 36px; height: 36px; border-radius: 50%; background: linear-gradient(135deg, var(--indigo), var(--purple)); display: flex; align-items: center; justify-content: center; font-family: var(--font-display); font-size: 14px; font-weight: 800; flex-shrink: 0; }
.fc-name { font-weight: 700; font-size: 13px; }
.fc-role { font-size: 11px; color: var(--neon-purple); }
.fc-live { margin-left: auto; font-size: 11px; color: var(--success); font-weight: 700; flex-shrink: 0; }
.fc-stat { font-size: 12px; margin-bottom: 8px; }
.fc-bar { height: 4px; background: rgba(255,255,255,0.1); border-radius: 2px; }
.fc-bar-fill { height: 100%; background: linear-gradient(90deg, var(--indigo), var(--neon-cyan)); border-radius: 2px; }
.fc-ai { font-size: 11px; color: var(--neon-purple); font-weight: 700; margin-bottom: 6px; }
.fc-text { font-size: 12px; color: var(--muted); line-height: 1.5; }
@keyframes float { 0%, 100% { transform: translateY(0); } 50% { transform: translateY(-10px); } }

/* --- Sections common --- */
.section-inner { max-width: 1100px; margin: 0 auto; padding: 80px 24px; text-align: center; }
.section-tag { font-size: 12px; font-weight: 700; letter-spacing: 2px; color: #818cf8; text-transform: uppercase; margin-bottom: 12px; }
.section-title { font-size: clamp(26px, 4vw, 40px); font-weight: 800; margin-bottom: 48px; }

/* --- Features --- */
.features-section { background: var(--bg2); }
.features-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(240px, 1fr)); gap: 20px; text-align: left; }
.feature-card { padding: 24px; border-radius: 18px; transition: all 0.25s; background: var(--surface); border: 1px solid var(--border); }
.feature-card:hover { transform: translateY(-4px); border-color: var(--indigo); }
.feat-icon { font-size: 28px; margin-bottom: 14px; }
.feat-title { font-size: 15px; font-weight: 700; margin-bottom: 8px; }
.feat-desc { font-size: 13px; color: var(--muted); line-height: 1.7; }

/* --- How it works --- */
.steps-row { display: flex; gap: 32px; justify-content: center; flex-wrap: wrap; }
.how-step { max-width: 220px; text-align: center; }
.how-num { width: 40px; height: 40px; border-radius: 50%; background: linear-gradient(135deg, var(--indigo), var(--purple)); display: flex; align-items: center; justify-content: center; font-family: var(--font-display); font-size: 16px; font-weight: 800; margin: 0 auto 14px; }
.how-icon { font-size: 32px; margin-bottom: 12px; }
.how-title { font-size: 16px; font-weight: 700; margin-bottom: 8px; }
.how-desc { font-size: 13px; color: var(--muted); line-height: 1.7; }

/* --- Pricing --- */
.pricing-section { background: var(--bg2); }
.pricing-grid { display: flex; gap: 20px; justify-content: center; flex-wrap: wrap; }
.pricing-card { background: var(--surface); border: 1px solid var(--border); border-radius: 24px; padding: 32px 28px; min-width: 240px; max-width: 280px; flex: 1; text-align: left; position: relative; transition: all 0.25s; }
.pricing-card.featured { border-color: var(--indigo); background: var(--bg2); transform: scale(1.03); }
.plan-badge { position: absolute; top: -14px; left: 50%; transform: translateX(-50%); padding: 4px 16px; background: linear-gradient(135deg, var(--indigo), var(--purple)); border-radius: 100px; font-size: 12px; font-weight: 700; white-space: nowrap; color: #fff; }
.plan-name { font-size: 14px; font-weight: 700; color: var(--muted); letter-spacing: 1px; text-transform: uppercase; margin-bottom: 12px; }
.plan-price { display: flex; align-items: baseline; gap: 6px; margin-bottom: 24px; }
.price-num { font-family: var(--font-display); font-size: 40px; font-weight: 800; }
.price-per { font-size: 13px; color: var(--muted); }
.plan-features { list-style: none; margin-bottom: 28px; display: flex; flex-direction: column; gap: 10px; }
.plan-features li { font-size: 13px; color: var(--muted); }

/* --- CTA --- */
.cta-section { position: relative; padding: 80px 24px; text-align: center; overflow: hidden; }
.cta-glow { position: absolute; inset: 0; background: radial-gradient(ellipse 60% 80% at 50% 50%, rgba(79,70,229,0.2) 0%, transparent 70%); pointer-events: none; }
.cta-content { position: relative; z-index: 1; }
.cta-content h2 { font-size: clamp(24px, 4vw, 40px); font-weight: 800; margin-bottom: 12px; }
.cta-content p { font-size: 15px; color: var(--muted); margin-bottom: 28px; }

/* --- Footer --- */
.footer { border-top: 1px solid var(--border); padding: 28px 24px; }
.footer-inner { max-width: 1100px; margin: 0 auto; display: flex; align-items: center; gap: 20px; flex-wrap: wrap; }
.footer-brand { display: flex; align-items: center; gap: 8px; margin-right: auto; }
.footer-links { display: flex; gap: 16px; }
.footer-links a { font-size: 13px; color: var(--muted); text-decoration: none; }
.footer-links a:hover { color: var(--text); }

@media (max-width: 600px) {
  .nav-links .nav-link { display: none; }
  .hero-stats { gap: 14px; }
}
</style>
