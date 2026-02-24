<!-- NotFoundView.vue -->
<template>
  <div class="notfound-page">
    <div style="margin-bottom:16px; color: var(--muted)"><Ghost :size="72" :stroke-width="1.5" /></div>
    <h1>404 — Page not found</h1>
    <p> The page you are looking for might have been removed, had its name changed, or is temporarily unavailable. </p>
    <RouterLink to="/dashboard" class="btn btn-primary" style="margin-top:24px">Back to dashboard</RouterLink>
  </div>
</template>
<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { Ghost } from 'lucide-vue-next'
</script>
<style scoped>
.notfound-page { display:flex;flex-direction:column;align-items:center;justify-content:center;min-height:100vh;text-align:center;padding:24px; }
h1 { font-size:28px;font-weight:800;margin-bottom:8px; }
p { color:var(--muted); }
</style>
