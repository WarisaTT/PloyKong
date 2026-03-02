<template>
  <div class="template-previews-container">
    <!-- CLASSIC -->
    <div v-if="layout === 'classic'" class="tpl-classic" :style="{ background: th.bg, minHeight: '500px', fontFamily: 'system-ui, sans-serif' }">
      <!-- Hero Area (Gradient Header) -->
      <div :style="{ background: `radial-gradient(ellipse at top, ${th.primary}40, transparent 70%)`, padding: '40px 24px 24px', textAlign: 'center', borderBottom: `1px solid ${th.border}` }">
        <div :style="{ width: '80px', height: '80px', borderRadius: '50%', background: `linear-gradient(135deg, ${th.primary}, ${th.secondary})`, margin: '0 auto 16px', display: 'flex', alignItems: 'center', justifyContent: 'center', fontSize: '36px', boxShadow: th.isDark ? `0 0 24px ${th.primary}50` : `0 12px 24px ${th.primary}30`, border: `4px solid ${th.bg}` }">✨</div>
        <div :style="{ color: th.text, fontSize: '24px', fontWeight: 900, letterSpacing: '-0.5px', marginBottom: '8px' }">{{ h.name }}</div>
        <div :style="{ background: `linear-gradient(135deg, ${th.primary}, ${th.secondary})`, WebkitBackgroundClip: 'text', WebkitTextFillColor: 'transparent', fontSize: '14px', fontWeight: 800, marginBottom: '12px', display: 'inline-block' }">{{ h.role }}</div>
        <div :style="{ color: th.muted, fontSize: '12px', lineHeight: 1.6, maxWidth: '300px', margin: '0 auto', marginBottom: '20px' }">{{ h.tagline.split('\n')[0] }}</div>
        
        <div style="display: flex; gap: 8px; justify-content: center; flex-wrap: wrap;">
          <div :style="{ background: `linear-gradient(135deg, ${th.primary}, ${th.secondary})`, color: '#fff', fontSize: '10px', fontWeight: 800, padding: '8px 16px', borderRadius: '100px', boxShadow: `0 4px 12px ${th.primary}40` }">Hire Me</div>
          <div :style="{ background: th.surface, color: th.text, border: `1px solid ${th.primary}50`, fontSize: '10px', fontWeight: 800, padding: '8px 16px', borderRadius: '100px' }">Download CV</div>
        </div>
      </div>
      
      <!-- Body Section -->
      <div style="padding: 24px;">
        <div style="display: grid; grid-template-columns: 1fr; gap: 20px;">
          <!-- Skills -->
          <div :style="{ background: th.surface, borderRadius: '16px', padding: '20px', border: `1px solid ${th.border}`, boxShadow: th.isDark ? '0 4px 20px rgba(0,0,0,0.4)' : '0 4px 20px rgba(0,0,0,0.05)' }">
            <div :style="{ display: 'flex', alignItems: 'center', gap: '8px', marginBottom: '16px' }">
              <div :style="{ background: th.primary + '20', color: th.primary, padding: '6px', borderRadius: '8px', display: 'flex', alignItems: 'center', justifyContent: 'center' }">⭐</div>
              <span :style="{ color: th.text, fontSize: '16px', fontWeight: 800 }">Skills</span>
            </div>
            <div style="display: grid; grid-template-columns: 1fr; gap: 12px;">
              <div v-for="(s, i) in skills.slice(0,3)" :key="i">
                <div style="display:flex; justify-content:space-between; margin-bottom:6px;">
                  <span :style="{ color: th.text, fontSize: '12px', fontWeight: 600 }">{{ s.name }}</span>
                  <span :style="{ color: th.primary, fontSize: '11px', fontWeight: 700 }">{{ s.level }}%</span>
                </div>
                <div :style="{ background: th.border, borderRadius: '100px', height: '6px', overflow: 'hidden' }">
                  <div :style="{ width: s.level + '%', height: '100%', background: `linear-gradient(90deg, ${th.primary}, ${th.secondary})`, borderRadius: '100px' }"></div>
                </div>
              </div>
            </div>
          </div>

          <!-- Experience -->
          <div :style="{ background: th.surface, borderRadius: '16px', padding: '20px', border: `1px solid ${th.border}`, boxShadow: th.isDark ? '0 4px 20px rgba(0,0,0,0.4)' : '0 4px 20px rgba(0,0,0,0.05)' }">
            <div :style="{ display: 'flex', alignItems: 'center', gap: '8px', marginBottom: '16px' }">
              <div :style="{ background: th.primary + '20', color: th.primary, padding: '6px', borderRadius: '8px', display: 'flex', alignItems: 'center', justifyContent: 'center' }">💼</div>
              <span :style="{ color: th.text, fontSize: '16px', fontWeight: 800 }">Experience</span>
            </div>
            <div style="position: relative; padding-left: 14px; border-left: 2px solid var(--border);">
              <div v-for="(e, i) in exps" :key="i" :style="{ marginBottom: Number(i) < exps.length - 1 ? '16px' : '0', position: 'relative' }">
                <div :style="{ position: 'absolute', left: '-20px', top: '4px', width: '10px', height: '10px', borderRadius: '50%', background: th.bg, border: `2px solid ${th.primary}` }"></div>
                <div :style="{ color: th.text, fontSize: '13px', fontWeight: 800 }">{{ e.position }}</div>
                <div style="display: flex; gap: 8px; align-items: center; margin-top: 4px; margin-bottom: 6px;">
                  <span :style="{ color: th.primary, fontSize: '11px', fontWeight: 600 }">{{ e.company }}</span>
                  <span :style="{ color: th.muted, fontSize: '10px', background: th.border, padding: '2px 8px', borderRadius: '100px' }">{{ e.start_date }} – {{ e.is_current ? 'Present' : e.end_date }}</span>
                </div>
                <div :style="{ color: th.muted, fontSize: '11px', lineHeight: 1.6 }">{{ e.description }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- PROGRAMMER -->
    <div v-if="layout === 'programmer'" class="tpl-programmer" :style="{ backgroundColor: th.bg, minHeight: '500px', fontFamily: '\'Courier New\', monospace', display: 'flex', flexDirection: 'column' }">
      <!-- Top bar -->
      <div :style="{ background: th.isDark ? th.primary + '12' : th.card, borderBottom: `1px solid ${th.primary}40`, padding: '12px 16px', display: 'flex', alignItems: 'center', justifyContent: 'space-between' }">
        <div style="display:flex; gap:6px;">
          <div v-for="c in ['#FF5F57','#FFBD2E','#28CA41']" :key="c" :style="{ width: '10px', height: '10px', borderRadius: '50%', background: c }"></div>
        </div>
        <span :style="{ color: th.primary, fontSize: '10px', fontWeight: 700, letterSpacing: '2px' }">portfolio.exe — {{ h.name }}</span>
        <span :style="{ color: th.muted, fontSize: '9px' }">●</span>
      </div>
      <!-- Body -->
      <div style="display:grid; grid-template-columns:1fr 1fr; gap:0; flex:1;">
        <!-- LEFT -->
        <div :style="{ borderRight: `1px solid ${th.border}`, padding: '14px' }">
          <div style="margin-bottom:14px;">
            <div :style="{ color: th.primary, fontSize: '9px', fontWeight: 700, letterSpacing: '2px', marginBottom: '6px' }">// IDENTITY</div>
            <div :style="{ width: '44px', height: '44px', borderRadius: '6px', background: `linear-gradient(135deg, ${th.primary}, ${th.secondary})`, display: 'flex', alignItems: 'center', justifyContent: 'center', fontSize: '20px', marginBottom: '8px', boxShadow: th.isDark ? `0 0 14px ${th.primary}40` : 'none' }">👤</div>
            <div :style="{ color: th.text, fontSize: '14px', fontWeight: 800, letterSpacing: '-0.5px' }">{{ h.name }}</div>
            <div :style="{ color: th.primary, fontSize: '10px', fontWeight: 700, marginBottom: '4px' }">{{ h.role }}</div>
            <div :style="{ color: th.muted, fontSize: '9px', lineHeight: 1.6 }">{{ h.tagline.split('\n')[0] }}</div>
          </div>
          <div>
            <div :style="{ color: th.primary, fontSize: '9px', fontWeight: 700, letterSpacing: '2px', marginBottom: '8px' }">// SKILLS</div>
            <div v-for="(s, i) in skills" :key="i" style="margin-bottom: 8px;">
              <div style="display:flex; justify-content:space-between; margin-bottom:4px;">
                <span :style="{ color: th.text, fontSize: '11px', fontWeight: 600 }">{{ s.name }}</span>
                <span :style="{ color: th.primary, fontSize: '10px', fontWeight: 700 }">{{ s.level }}%</span>
              </div>
              <div :style="{ background: th.border, borderRadius: '100px', height: '4px', overflow: 'hidden' }">
                <div :style="{ width: s.level + '%', height: '100%', background: `linear-gradient(90deg, ${th.primary}80, ${th.primary})`, borderRadius: '100px', boxShadow: th.isDark ? `0 0 6px ${th.primary}60` : 'none' }"></div>
              </div>
            </div>
          </div>
        </div>
        <!-- RIGHT -->
        <div style="padding:14px; display:flex; flex-direction:column; gap:12px;">
          <div>
            <div :style="{ color: th.secondary, fontSize: '9px', fontWeight: 700, letterSpacing: '2px', marginBottom: '8px' }">// EXPERIENCE</div>
            <div v-for="(e, i) in exps" :key="i" :style="{ background: th.isDark ? th.secondary + '08' : th.card, border: `1px solid ${th.border}`, borderLeft: `2px solid ${th.primary}`, borderRadius: '6px', padding: '8px 10px', marginBottom: '6px' }">
              <div :style="{ color: th.text, fontSize: '11px', fontWeight: 700 }">{{ e.position }}</div>
              <div :style="{ color: th.primary, fontSize: '9px', fontWeight: 600 }">{{ e.company }} · {{ e.is_current ? 'Now' : e.end_date }}</div>
              <div :style="{ color: th.muted, fontSize: '9px', marginTop: '3px', lineHeight: 1.5 }">{{ e.description }}</div>
            </div>
          </div>
          <div>
            <div :style="{ color: th.secondary, fontSize: '9px', fontWeight: 700, letterSpacing: '2px', marginBottom: '8px' }">// PROJECTS</div>
            <div v-for="(p, i) in projects" :key="i" :style="{ background: th.card, border: `1px solid ${th.border}`, borderRadius: '6px', padding: '8px 10px', marginBottom: '6px' }">
              <div :style="{ color: th.text, fontSize: '11px', fontWeight: 700, marginBottom: '3px' }">{{ p.title }}</div>
              <div style="display:flex; gap:4px; flex-wrap:wrap;">
                <span v-for="(t, j) in p.tags" :key="j" :style="{ background: th.primary + '18', color: th.primary, border: `1px solid ${th.primary}35`, fontSize: '9px', fontWeight: 700, padding: '2px 7px', borderRadius: '100px', letterSpacing: '0.5px' }">{{ t }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- DESIGNER -->
    <div v-else-if="layout === 'designer'" class="tpl-designer" :style="{ background: th.bg, minHeight: '500px', fontFamily: 'Georgia, \'Times New Roman\', serif' }">
      <div :style="{ position: 'relative', overflow: 'hidden', background: th.isDark ? `linear-gradient(135deg, ${th.primary}20 0%, ${th.bg} 60%)` : `linear-gradient(135deg, ${th.primary}10 0%, ${th.surface} 60%)`, borderBottom: `1px solid ${th.border}`, padding: '24px 20px 20px' }">
        <div :style="{ position: 'absolute', right: '-40px', top: '-40px', width: '160px', height: '160px', borderRadius: '50%', background: `radial-gradient(circle, ${th.primary}30, transparent 70%)`, pointerEvents: 'none' }"></div>
        <div :style="{ color: th.muted, fontSize: '9px', fontWeight: 700, letterSpacing: '3px', textTransform: 'uppercase', marginBottom: '12px' }">Portfolio</div>
        <div :style="{ fontSize: '28px', fontWeight: 900, lineHeight: 1, marginBottom: '6px', color: th.text, letterSpacing: '-1.5px' }">{{ h.name }}</div>
        <div :style="{ fontSize: '13px', fontWeight: 400, fontStyle: 'italic', background: `linear-gradient(90deg, ${th.primary}, ${th.secondary})`, WebkitBackgroundClip: 'text', WebkitTextFillColor: 'transparent', marginBottom: '12px' }">{{ h.role }}</div>
        <div v-if="h.show_hire_me" :style="{ display: 'inline-block', background: `linear-gradient(135deg, ${th.primary}, ${th.secondary})`, color: th.isDark ? '#000' : '#fff', fontSize: '10px', fontWeight: 800, letterSpacing: '1.5px', textTransform: 'uppercase', padding: '7px 16px', borderRadius: '2px', boxShadow: th.isDark ? `0 0 16px ${th.primary}50` : `0 4px 12px ${th.primary}30` }">Hire Me →</div>
      </div>
      <div style="padding:16px 20px; display:grid; grid-template-columns:3fr 2fr; gap:16px;">
        <div>
          <div :style="{ color: th.primary, fontSize: '9px', fontWeight: 800, letterSpacing: '3px', textTransform: 'uppercase', marginBottom: '12px' }">Selected Work</div>
          <div v-for="(p, i) in projects" :key="i" :style="{ borderTop: `2px solid ${i===0 ? th.primary : th.border}`, paddingTop: '10px', paddingBottom: '10px' }">
            <div :style="{ color: th.text, fontSize: '13px', fontWeight: 700, marginBottom: '4px' }">{{ p.title }}</div>
            <div :style="{ color: th.muted, fontSize: '10px', lineHeight: 1.6, marginBottom: '6px' }">{{ p.description }}</div>
            <div style="display:flex; gap:4px; flex-wrap:wrap;">
              <span v-for="(t, j) in p.tags" :key="j" :style="{ background: th.primary + '18', color: th.primary, border: `1px solid ${th.primary}35`, fontSize: '9px', fontWeight: 700, padding: '2px 7px', borderRadius: '100px', letterSpacing: '0.5px' }">{{ t }}</span>
            </div>
          </div>
        </div>
        <div>
          <div :style="{ color: th.primary, fontSize: '9px', fontWeight: 800, letterSpacing: '3px', textTransform: 'uppercase', marginBottom: '12px' }">Expertise</div>
          <div v-for="(s, i) in skills" :key="i" style="margin-bottom: 8px;">
            <div style="display:flex; justify-content:space-between; margin-bottom:4px;">
              <span :style="{ color: th.text, fontSize: '11px', fontWeight: 600 }">{{ s.name }}</span>
              <span :style="{ color: th.primary, fontSize: '10px', fontWeight: 700 }">{{ s.level }}%</span>
            </div>
            <div :style="{ background: th.border, borderRadius: '100px', height: '4px', overflow: 'hidden' }">
              <div :style="{ width: s.level + '%', height: '100%', background: `linear-gradient(90deg, ${th.primary}80, ${th.primary})`, borderRadius: '100px', boxShadow: th.isDark ? `0 0 6px ${th.primary}60` : 'none' }"></div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- FIRST JOBBER -->
    <div v-else-if="layout === 'firstjobber'" class="tpl-firstjobber" :style="{ background: th.bg, minHeight: '500px', fontFamily: '\'Plus Jakarta Sans\', sans-serif', padding: '10px 16px', display: 'flex', flexDirection: 'column', gap: '12px' }">
      <!-- Hero Card -->
      <div :style="{ background: th.surface, border: `1px solid ${th.border}`, borderRadius: '18px', padding: '24px', boxShadow: th.isDark ? '0 6px 24px rgba(0,0,0,0.3)' : '0 6px 24px rgba(0,0,0,0.08)', display: 'flex', gap: '20px', alignItems: 'center' }">
        <!-- Avatar -->
        <div :style="{ width: '72px', height: '72px', borderRadius: '14px', background: `linear-gradient(135deg, ${th.primary}, ${th.secondary})`, display: 'flex', alignItems: 'center', justifyContent: 'center', fontSize: '32px', flexShrink: 0 }">👤</div>
        <!-- Text Grid -->
        <div style="flex:1; min-width:0; display:grid; grid-template-columns:1fr auto; row-gap:2px; column-gap:12px; align-items:center;">
          <!-- Badge row 1 -->
          <div :style="{ gridColumn: '1/2', gridRow: '1' }">
            <span :style="{ background: th.primary + '20', color: th.primary, fontSize: '9px', fontWeight: 600, padding: '3px 10px', borderRadius: '999px', display: 'inline-block' }">{{ h.role }}</span>
          </div>
          <!-- Name row 2 -->
          <div :style="{ gridColumn: '1/2', gridRow: '2', color: th.text, fontSize: '18px', fontWeight: 800, letterSpacing: '-0.5px', lineHeight: 1.2 }">{{ h.name }}</div>
          <!-- CTA row 1-2 right -->
          <div :style="{ gridColumn: '2/3', gridRow: '1/3', display: 'flex', flexDirection: 'column', alignItems: 'flex-end', justifyContent: 'center', gap: '4px' }">
            <div v-if="h.show_hire_me" :style="{ background: `linear-gradient(135deg, ${th.primary}, ${th.secondary})`, color: '#fff', fontSize: '8px', fontWeight: 800, padding: '5px 12px', borderRadius: '100px', boxShadow: `0 3px 10px ${th.primary}40`, whiteSpace: 'nowrap' }">Hire Me!</div>
            <div v-if="h.show_resume" :style="{ background: th.surface, color: th.text, border: `1px solid ${th.border}`, fontSize: '8px', fontWeight: 700, padding: '4px 10px', borderRadius: '100px', whiteSpace: 'nowrap' }">Resume</div>
          </div>
          <!-- Tagline row 3 -->
          <div :style="{ gridColumn: '1/3', gridRow: '3', color: th.muted, fontSize: '10px', marginTop: '8px', lineHeight: 1.6, borderTop: `1px solid ${th.border}`, paddingTop: '10px' }">{{ h.tagline.split('\n')[0] }}</div>
        </div>
      </div>

      <!-- Skills + Experience side by side -->
      <div style="display:grid; grid-template-columns:1fr 1fr; gap:12px;">
        <!-- Skills Card -->
        <div :style="{ background: th.surface, border: `1px solid ${th.border}`, borderRadius: '16px', padding: '16px', boxShadow: th.isDark ? '0 4px 20px rgba(0,0,0,0.3)' : '0 2px 12px rgba(0,0,0,0.07)' }">
          <div :style="{ color: th.primary, fontSize: '9px', fontWeight: 800, letterSpacing: '1.5px', textTransform: 'uppercase', marginBottom: '12px' }">Skills</div>
          <div v-for="(s, i) in skills.slice(0,3)" :key="i" style="margin-bottom: 8px;">
            <div style="display:flex; justify-content:space-between; margin-bottom:4px;">
              <span :style="{ color: th.text, fontSize: '10px', fontWeight: 600 }">{{ s.name }}</span>
              <span :style="{ color: th.primary, fontSize: '9px', fontWeight: 700 }">{{ s.level }}%</span>
            </div>
            <div :style="{ background: th.border, borderRadius: '100px', height: '4px', overflow: 'hidden' }">
              <div :style="{ width: s.level + '%', height: '100%', background: `linear-gradient(90deg, ${th.primary}80, ${th.primary})`, borderRadius: '100px' }"></div>
            </div>
          </div>
        </div>

        <!-- Experience Card (flat list with dividers) -->
        <div :style="{ background: th.surface, border: `1px solid ${th.border}`, borderRadius: '16px', padding: '16px', boxShadow: th.isDark ? '0 4px 20px rgba(0,0,0,0.3)' : '0 2px 12px rgba(0,0,0,0.07)' }">
          <div :style="{ color: th.primary, fontSize: '9px', fontWeight: 800, letterSpacing: '1.5px', textTransform: 'uppercase', marginBottom: '12px' }">Experience</div>
          <div v-for="(e, i) in exps" :key="i" :style="{ paddingBottom: Number(i) < exps.length - 1 ? '10px' : '0', marginBottom: Number(i) < exps.length - 1 ? '10px' : '0', borderBottom: Number(i) < exps.length - 1 ? `1px solid ${th.border}` : 'none' }">
            <div :style="{ color: th.text, fontSize: '11px', fontWeight: 800 }">{{ e.position }}</div>
            <div :style="{ color: th.primary, fontSize: '9px', fontWeight: 600 }">{{ e.company }}</div>
            <div :style="{ color: th.muted, fontSize: '9px' }">{{ e.is_current ? 'Present' : e.end_date }}</div>
          </div>
        </div>
      </div>

      <!-- Projects Card (2-col grid with dividers) -->
      <div :style="{ background: th.surface, border: `1px solid ${th.border}`, borderRadius: '16px', padding: '16px', boxShadow: th.isDark ? '0 4px 20px rgba(0,0,0,0.3)' : '0 2px 12px rgba(0,0,0,0.07)' }">
        <div :style="{ color: th.primary, fontSize: '9px', fontWeight: 800, letterSpacing: '1.5px', textTransform: 'uppercase', marginBottom: '12px' }">Projects 🚀</div>
        <div style="display:grid; grid-template-columns:1fr 1fr; gap:0;">
          <div v-for="(p, i) in projects" :key="i" :style="{ padding: '8px 10px', borderBottom: Number(i) < projects.length - 2 ? `1px solid ${th.border}` : 'none', borderLeft: Number(i) % 2 === 1 ? `1px solid ${th.border}` : 'none' }">
            <div :style="{ color: th.text, fontSize: '11px', fontWeight: 700, marginBottom: '4px' }">{{ p.title }}</div>
            <div style="display:flex; gap:3px; flex-wrap:wrap;">
              <span v-for="(t, j) in p.tags.slice(0,2)" :key="j" :style="{ background: th.primary + '18', color: th.primary, border: `1px solid ${th.primary}35`, fontSize: '8px', fontWeight: 700, padding: '2px 7px', borderRadius: '100px' }">{{ t }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- BUSINESS -->
    <div v-else-if="layout === 'business'" class="tpl-business" :style="{ display: 'flex', background: th.bg, minHeight: '500px', fontFamily: 'serif' }">
      <div :style="{ width: '130px', flexShrink: 0, background: th.isDark ? th.primary + '18' : th.primary + '08', borderRight: `2px solid ${th.primary}40`, padding: '24px 12px', display: 'flex', flexDirection: 'column', alignItems: 'center', gap: '20px' }">
        <div :style="{ width: '64px', height: '64px', borderRadius: '8px', background: `linear-gradient(135deg, ${th.primary}, ${th.secondary})`, display: 'flex', alignItems: 'center', justifyContent: 'center', fontSize: '28px', boxShadow: `0 4px 16px ${th.primary}40`, border: `2px solid ${th.bg}` }">👤</div>
        <div style="text-align:center;">
          <div :style="{ color: th.text, fontSize: '13px', fontWeight: 900, textTransform: 'uppercase', letterSpacing: '1px' }">{{ h.name.split(' ')[0] }}</div>
          <div :style="{ color: th.primary, fontSize: '9px', fontWeight: 700, marginTop: '2px', opacity: 0.8 }">{{ h.role.split(' ')[0] }}</div>
        </div>
        <div style="width: 100%; display: flex; flex-direction: column; gap: 4px;">
           <div v-for="(item, i) in ['PROFILES','COMPETENCIES','HISTORY']" :key="i" :style="{ width: '100%', padding: '8px 4px', borderRadius: '2px', color: i===0 ? th.primary : th.muted, fontSize: '8px', fontWeight: 800, textAlign: 'center', borderBottom: `1px solid ${i===0 ? th.primary+'40' : 'transparent'}`, letterSpacing: '1px' }">{{ item }}</div>
        </div>
      </div>
      <div style="flex:1; padding:32px 24px; display:flex; flex-direction:column; gap:24px; overflow:hidden;">
        <div :style="{ borderBottom: `2px solid ${th.primary}`, paddingBottom: '20px' }">
          <div :style="{ color: th.primary, fontSize: '10px', fontWeight: 800, letterSpacing: '3px', textTransform: 'uppercase', marginBottom: '8px' }">Executive Summary</div>
          <div :style="{ color: th.text, fontSize: '24px', fontWeight: 900, letterSpacing: '-1px', marginBottom: '4px' }">{{ h.name }}</div>
          <div :style="{ color: th.muted, fontSize: '11px', fontWeight: 700, letterSpacing: '1px', marginBottom: '12px' }">{{ h.role }}</div>
          <div :style="{ color: th.muted, fontSize: '11px', lineHeight: 1.8 }">{{ h.tagline.split('\n')[0] }}</div>
        </div>
        <div style="display: grid; grid-template-columns: 1fr; gap: 20px;">
           <div>
              <div :style="{ color: th.primary, fontSize: '10px', fontWeight: 800, letterSpacing: '2px', textTransform: 'uppercase', marginBottom: '12px' }">Core Skills</div>
              <div style="display: flex; flex-wrap: wrap; gap: 8px;">
                <div v-for="(s, i) in skills.slice(0,4)" :key="i" :style="{ border: `1px solid ${th.border}`, padding: '6px 14px', borderRadius: '4px', fontSize: '10px', fontWeight: 700, color: th.text }">{{ s.name }}</div>
              </div>
           </div>
           <div>
              <div :style="{ color: th.primary, fontSize: '10px', fontWeight: 800, letterSpacing: '2px', textTransform: 'uppercase', marginBottom: '12px' }">Professional Experience</div>
              <div v-for="(e, i) in exps.slice(0,1)" :key="i" style="display: flex; gap: 16px;">
                <div :style="{ color: th.primary, fontSize: '10px', fontWeight: 800, width: '80px', flexShrink: 0 }">{{ e.start_date.split('-')[0] }} - Pres</div>
                <div>
                  <div :style="{ color: th.text, fontSize: '13px', fontWeight: 800 }">{{ e.position }}</div>
                  <div :style="{ color: th.muted, fontSize: '11px', marginTop: '2px' }">{{ e.company }}</div>
                </div>
              </div>
           </div>
        </div>
      </div>
    </div>

    <!-- CREATIVE -->
    <div v-else-if="layout === 'creative'" class="tpl-creative" :style="{ background: th.bg, minHeight: '500px', position: 'relative', overflow: 'hidden', fontFamily: '\'Impact\', \'Arial Black\', sans-serif' }">
      <div :style="{ position: 'absolute', top: '-60px', right: '-60px', width: '200px', height: '200px', borderRadius: '50%', background: `radial-gradient(circle, ${th.primary}25, transparent 70%)`, pointerEvents: 'none' }"></div>
      <div :style="{ position: 'absolute', bottom: '-40px', left: '-40px', width: '140px', height: '140px', borderRadius: '50%', background: `radial-gradient(circle, ${th.secondary}20, transparent 70%)`, pointerEvents: 'none' }"></div>
      <div style="position:relative; padding:20px 20px 16px; z-index:1;">
        <div :style="{ fontSize: '32px', fontWeight: 900, lineHeight: 0.9, letterSpacing: '-2px', textTransform: 'uppercase', color: th.isDark ? th.text : th.bg, WebkitTextStroke: th.isDark ? `1px ${th.primary}` : 'none', background: th.isDark ? 'none' : `linear-gradient(135deg, ${th.primary}, ${th.secondary})`, WebkitBackgroundClip: th.isDark ? 'unset' : 'text', WebkitTextFillColor: th.isDark ? th.text : 'transparent', marginBottom: '4px' }">{{ h.name.split(' ')[0] }}</div>
        <div :style="{ fontSize: '12px', fontWeight: 400, letterSpacing: '3px', textTransform: 'uppercase', color: th.primary, marginBottom: '2px', fontFamily: 'sans-serif' }">{{ h.role }}</div>
        <div style="display:flex; gap:4px; flex-wrap:wrap; margin:10px 0 18px;">
          <span v-for="(l, i) in ['Open to work', 'Bangkok', 'Gen-Z Dev']" :key="i" :style="{ background: i===0 ? `linear-gradient(135deg, ${th.primary}, ${th.secondary})` : th.primary + '15', color: i===0 ? (th.isDark?'#000':'#fff') : th.primary, fontSize: '8px', fontWeight: 800, letterSpacing: '1px', textTransform: 'uppercase', padding: '3px 8px', borderRadius: '2px' }">{{ l }}</span>
        </div>
        <div style="display:grid; grid-template-columns:3fr 2fr; gap:14px;">
          <div>
            <div v-for="(p, i) in projects" :key="i" :style="{ position: 'relative', overflow: 'hidden', background: th.primary + (i===0?'18':'10'), border: `1px solid ${th.primary}${i===0?'50':'25'}`, borderRadius: i===0 ? '14px' : '8px', padding: '12px 14px', marginBottom: '8px' }">
              <div v-if="i===0" :style="{ position: 'absolute', top: 0, left: 0, right: 0, height: '2px', background: `linear-gradient(90deg, ${th.primary}, ${th.secondary})` }"></div>
              <div :style="{ color: th.text, fontSize: '12px', fontWeight: 700, marginBottom: '4px' }">{{ p.title }}</div>
              <div :style="{ color: th.muted, fontSize: '9px', lineHeight: 1.5, marginBottom: '6px' }">{{ p.description }}</div>
              <div style="display:flex; gap:4px; flex-wrap:wrap;">
                <span v-for="(t, j) in p.tags" :key="j" :style="{ background: th.primary + '18', color: th.primary, border: `1px solid ${th.primary}35`, fontSize: '9px', fontWeight: 700, padding: '2px 7px', borderRadius: '100px', letterSpacing: '0.5px' }">{{ t }}</span>
              </div>
            </div>
          </div>
          <div>
            <div :style="{ writingMode: 'vertical-rl', textOrientation: 'mixed', color: th.primary, fontSize: '9px', fontWeight: 900, letterSpacing: '3px', textTransform: 'uppercase', marginBottom: '10px', transform: 'rotate(180deg)' }">SKILLS</div>
            <div v-for="(s, i) in skills" :key="i" :style="{ background: th.primary + Math.round((s.level/100)*30+8).toString(16).padStart(2,'0'), border: `1px solid ${th.primary}30`, borderRadius: '6px', padding: '5px 8px', marginBottom: '5px', display: 'flex', justifyContent: 'space-between', alignItems: 'center' }">
              <span :style="{ color: th.text, fontSize: '9px', fontWeight: 700 }">{{ s.name.split(' ')[0] }}</span>
              <span :style="{ color: th.primary, fontSize: '9px', fontWeight: 800 }">{{ s.level }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- OTHER -->
    <div v-else-if="layout === 'other'" class="tpl-other" :style="{ background: th.bg, minHeight: '500px', fontFamily: '\'Garamond\',\'Book Antiqua\',Georgia,serif', padding: '28px 24px' }">
      <div :style="{ height: '2px', background: `linear-gradient(90deg, ${th.primary}, ${th.secondary}, transparent)`, marginBottom: '24px' }"></div>
      <div style="text-align:center; margin-bottom:28px;">
        <div :style="{ color: th.muted, fontSize: '9px', letterSpacing: '4px', textTransform: 'uppercase', marginBottom: '8px' }">{{ h.role }}</div>
        <div :style="{ color: th.text, fontSize: '22px', fontWeight: 400, letterSpacing: '-0.5px', marginBottom: '8px' }">{{ h.name }}</div>
        <div :style="{ width: '40px', height: '1px', background: `linear-gradient(90deg, transparent, ${th.primary}, transparent)`, margin: '0 auto 10px' }"></div>
        <div :style="{ color: th.muted, fontSize: '10px', lineHeight: 1.8, fontStyle: 'italic', maxWidth: '280px', margin: '0 auto' }">{{ h.tagline.split('\n')[0] }}</div>
        <div v-if="h.show_hire_me" :style="{ marginTop: '12px', display: 'inline-block', border: `1px solid ${th.primary}`, color: th.primary, fontSize: '9px', fontWeight: 600, letterSpacing: '2px', textTransform: 'uppercase', padding: '6px 18px', borderRadius: 0 }">Available for Hire</div>
      </div>
      <div style="display:grid; grid-template-columns:1fr 1fr 1fr; gap:20px;">
        <div>
          <div :style="{ color: th.primary, fontSize: '8px', fontWeight: 700, letterSpacing: '3px', textTransform: 'uppercase', borderBottom: `1px solid ${th.border}`, paddingBottom: '6px', marginBottom: '10px' }">Skills</div>
          <div v-for="(s, i) in skills" :key="i" style="margin-bottom:8px;">
            <div style="display:flex; justify-content:space-between;">
              <span :style="{ color: th.text, fontSize: '10px' }">{{ s.name }}</span>
              <span :style="{ color: th.muted, fontSize: '9px' }">{{ s.level }}%</span>
            </div>
            <div :style="{ height: '1px', background: th.border, marginTop: '4px', position: 'relative' }">
              <div :style="{ position: 'absolute', left: 0, top: 0, height: '1px', width: s.level + '%', background: th.primary }"></div>
            </div>
          </div>
        </div>
        <div>
          <div :style="{ color: th.primary, fontSize: '8px', fontWeight: 700, letterSpacing: '3px', textTransform: 'uppercase', borderBottom: `1px solid ${th.border}`, paddingBottom: '6px', marginBottom: '10px' }">Experience</div>
          <div v-for="(e, i) in exps" :key="i" style="margin-bottom:10px;">
            <div :style="{ color: th.text, fontSize: '11px', fontWeight: 700, fontStyle: 'italic' }">{{ e.position }}</div>
            <div :style="{ color: th.primary, fontSize: '9px' }">{{ e.company }}</div>
            <div :style="{ color: th.muted, fontSize: '8px' }">{{ e.start_date }} – {{ e.is_current ? 'Present' : e.end_date }}</div>
          </div>
        </div>
        <div>
          <div :style="{ color: th.primary, fontSize: '8px', fontWeight: 700, letterSpacing: '3px', textTransform: 'uppercase', borderBottom: `1px solid ${th.border}`, paddingBottom: '6px', marginBottom: '10px' }">Projects</div>
          <div v-for="(p, i) in projects" :key="i" style="margin-bottom:10px;">
            <div :style="{ color: th.text, fontSize: '11px', fontWeight: 700 }">{{ p.title }}</div>
            <div :style="{ color: th.muted, fontSize: '9px', lineHeight: 1.5, marginBottom: '4px' }">{{ p.description }}</div>
            <div style="display:flex; gap:3px; flex-wrap:wrap;">
              <span v-for="(t, j) in p.tags.slice(0,2)" :key="j" :style="{ background: th.primary + '18', color: th.primary, border: `1px solid ${th.primary}35`, fontSize: '9px', fontWeight: 700, padding: '2px 7px', borderRadius: '100px', letterSpacing: '0.5px' }">{{ t }}</span>
            </div>
          </div>
        </div>
      </div>
      <div :style="{ marginTop: '20px', paddingTop: '14px', borderTop: `1px solid ${th.border}`, display: 'flex', justifyContent: 'center', gap: '20px' }">
        <span v-for="(v, i) in [ct.email, ct.location, ct.github].filter(Boolean)" :key="i" :style="{ color: th.muted, fontSize: '9px', letterSpacing: '0.5px' }">{{ v }}</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps<{
  layout: string;
  theme: any;
  mockData?: any;
}>();

// Ensure 'th' is always bound
const th = computed(() => props.theme);

// Fallback to internal MOCK if not provided
const h = computed(() => props.mockData?.hero || {
  name: "Ploy Kong",
  role: "Full-Stack Developer",
  tagline: "Building cool things with code.\nTurning ideas into products that matter.",
  show_hire_me: true
});
const skills = computed(() => props.mockData?.skills?.items || [
  { name: "Vue.js / React", level: 92 },
  { name: "Go / Node.js", level: 80 },
  { name: "UI/UX Design", level: 68 },
  { name: "Docker / DevOps", level: 74 },
]);
const exps = computed(() => props.mockData?.experience?.items || [
  { company: "TechStart Co.", position: "Frontend Engineer", start_date: "2023-01", end_date: "", is_current: true, description: "Led Vue.js migration and improved performance by 40%." },
  { company: "Creative Lab", position: "Junior Developer", start_date: "2022-01", end_date: "2022-12", is_current: false, description: "Built responsive UIs for 10+ client projects." },
]);
const projects = computed(() => props.mockData?.projects?.items || [
  { title: "PloyKong Builder", description: "No-code portfolio platform for Gen-Z", tags: ["Vue.js", "Go", "Docker"] },
  { title: "AI Resume Score", description: "LLM-powered resume analyzer", tags: ["Python", "RAG", "OpenAI"] },
]);
const ct = computed(() => props.mockData?.contact || {
  email: "alex@example.com",
  location: "Bangkok, Thailand",
  github: "github.com/alex"
});
</script>

<style scoped>
.template-previews-container {
  width: 100%;
}
</style>
