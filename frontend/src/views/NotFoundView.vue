<!-- NotFoundView.vue -->
<template>
  <div class="notfound-page">
    <div style="font-size:72px;margin-bottom:16px">🛸</div>
    <h1>404 — ไม่พบหน้านี้</h1>
    <p>หน้าที่คุณตามหาอาจถูกย้ายหรือลบไปแล้ว</p>
    <RouterLink to="/dashboard" class="btn btn-primary" style="margin-top:24px">กลับหน้าหลัก</RouterLink>
  </div>
</template>
<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>
<style scoped>
.notfound-page { display:flex;flex-direction:column;align-items:center;justify-content:center;min-height:100vh;text-align:center;padding:24px; }
h1 { font-size:28px;font-weight:800;margin-bottom:8px; }
p { color:var(--muted); }
</style>
