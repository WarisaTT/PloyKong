<template>
  <div class="projects-preview">
    <div v-if="!data.items || !data.items.length" class="empty-msg">ยังไม่มี projects</div>
    <div v-else class="proj-grid">
      <div v-for="(proj, i) in data.items.slice(0, 3)" :key="i" class="proj-card">
        <div class="proj-title">{{ proj.title }}</div>
        <div class="proj-desc">{{ proj.description }}</div>
        <div class="proj-tags">
          <span v-for="tag in (proj.tags || []).slice(0,3)" :key="tag" class="proj-tag">{{ tag }}</span>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
defineProps<{ data: any }>()
</script>
<style scoped>
.projects-preview {}
.empty-msg { font-size: 12px; color: var(--muted); font-style: italic; }
.proj-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(160px, 1fr)); gap: 8px; }
.proj-card { background: rgba(255,255,255,0.03); border: 1px solid var(--border); border-radius: 8px; padding: 10px; }
.proj-title { font-size: 12px; font-weight: 700; margin-bottom: 4px; }
.proj-desc { font-size: 11px; color: var(--muted); margin-bottom: 6px; line-height: 1.4; overflow: hidden; display: -webkit-box; -webkit-line-clamp: 2; -webkit-box-orient: vertical; }
.proj-tags { display: flex; flex-wrap: wrap; gap: 4px; }
.proj-tag { font-size: 10px; padding: 2px 7px; border-radius: 100px; background: rgba(79,70,229,0.15); color: #818cf8; border: 1px solid rgba(79,70,229,0.25); }
</style>
