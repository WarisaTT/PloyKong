<template>
  <div id="app" :class="[themeStore.themeClass]">
    <RouterView v-slot="{ Component }">
      <Transition name="page" mode="out-in">
        <component :is="Component" />
      </Transition>
    </RouterView>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, watch } from "vue";
import { RouterView } from "vue-router";
import { useAuthStore } from "@/stores/auth";
import { useThemeStore } from "@/stores/theme";

const authStore = useAuthStore();
const themeStore = useThemeStore();

// Sync theme to body for teleported components
watch(() => themeStore.themeClass, (newClass, oldClass) => {
  if (oldClass) document.body.classList.remove(oldClass);
  if (newClass) document.body.classList.add(newClass);
}, { immediate: true });

onMounted(async () => {
  if (authStore.isAuthenticated && !authStore.user) {
    await authStore.fetchMe();
  }
});
</script>

<style>
#app {
  min-height: 100vh;
  /* Sync background with main theme variables */
  background: var(--bg);
  background-color: var(--bg-color);
  color: var(--text);
  transition: background-color 0.3s ease;
}

.page-enter-active,
.page-leave-active {
  transition: opacity 0.2s ease;
}
.page-enter-from,
.page-leave-to {
  opacity: 0;
}
</style>
