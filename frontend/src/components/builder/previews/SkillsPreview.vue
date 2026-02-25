<template>
  <div class="skills-preview">
    <div v-if="!data.items || data.items.length === 0" class="empty-msg">
      ยังไม่มี skills — คลิก Properties เพื่อเพิ่ม
    </div>
    <div v-else class="skills-list">
      <div
        v-for="(skill, i) in data.items.slice(0, 5)"
        :key="i"
        class="skill-row"
      >
        <div class="skill-name">{{ skill.name }}</div>
        <div class="skill-bar-wrap">
          <div class="skill-bar" :style="{ width: skill.level + '%' }"></div>
        </div>
        <div class="skill-level">{{ skill.level }}%</div>
      </div>
      <div v-if="data.items.length > 5" class="more-hint">
        +{{ data.items.length - 5 }} more skills
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
defineProps<{ data: any }>();
</script>
<style scoped>
.skills-preview {
  display: flex;
  flex-direction: column;
  gap: 8px;
}
.empty-msg {
  font-size: 12px;
  color: var(--muted);
  font-style: italic;
}
.skills-list {
  display: flex;
  flex-direction: column;
  gap: 7px;
}
.skill-row {
  display: flex;
  align-items: center;
  gap: 10px;
}
.skill-name {
  font-size: 12px;
  font-weight: 600;
  width: 80px;
  flex-shrink: 0;
}
.skill-bar-wrap {
  flex: 1;
  height: 4px;
  background: rgba(255, 255, 255, 0.08);
  border-radius: 2px;
  overflow: hidden;
}
:global(.theme-light) .skill-bar-wrap {
  background: rgba(0, 0, 0, 0.15);
}
.skill-bar {
  height: 100%;
  background: linear-gradient(90deg, var(--indigo), var(--neon-cyan));
  border-radius: 2px;
  transition: width 0.5s ease;
}
:global(.theme-light) .skill-bar {
  background: linear-gradient(
    90deg,
    var(--indigo),
    color-mix(in srgb, var(--indigo) 60%, black)
  );
}
.skill-level {
  font-size: 11px;
  color: var(--muted);
  width: 32px;
  text-align: right;
  flex-shrink: 0;
}
.more-hint {
  font-size: 11px;
  color: var(--muted);
  text-align: center;
  margin-top: 2px;
}
:global(.theme-light) .skill-name {
  color: #111827;
}
:global(.theme-dark) .skill-name {
  color: #ffffff;
}
:global(.theme-light) .skill-level {
  color: #4b5563;
}
:global(.theme-light) .empty-msg,
:global(.theme-light) .more-hint {
  color: #6b7280;
}
</style>
