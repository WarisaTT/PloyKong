<template>
  <div class="theme-switcher" :class="{ open: showMenu }" @click.stop>
    <button class="theme-btn" @click="showMenu = !showMenu" :title="`Theme: ${themeStore.mode}`">
      <SunIcon v-if="!themeStore.resolvedDark" :size="16" />
      <MoonIcon v-else :size="16" />
    </button>
    <Transition name="menu-pop">
      <div v-if="showMenu" class="theme-menu">
        <button
          v-for="opt in options"
          :key="opt.value"
          class="theme-opt"
          :class="{ active: themeStore.mode === opt.value }"
          @click="select(opt.value as 'light' | 'dark' | 'auto')"
        >
          <component :is="opt.icon" :size="14" />
          {{ opt.label }}
          <CheckIcon v-if="themeStore.mode === opt.value" :size="12" class="check" />
        </button>
      </div>
    </Transition>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue";
import { useThemeStore } from "@/stores/theme";
import { Sun as SunIcon, Moon as MoonIcon, Monitor, Check as CheckIcon } from "lucide-vue-next";

const themeStore = useThemeStore();
const showMenu = ref(false);

const options = [
  { value: "light", label: "Light", icon: SunIcon },
  { value: "dark", label: "Dark", icon: MoonIcon },
  { value: "auto", label: "System", icon: Monitor }, // Map 'system' to our Store's 'auto'
];

function select(val: "light" | "dark" | "auto") {
  themeStore.setMode(val);
  showMenu.value = false;
}

function onClickOutside() {
  showMenu.value = false;
}
onMounted(() => document.addEventListener("click", onClickOutside));
onUnmounted(() => document.removeEventListener("click", onClickOutside));
</script>

<style scoped>
.theme-switcher {
  position: relative;
}

.theme-btn {
  width: 36px;
  height: 36px;
  border-radius: 10px;
  background: var(--surface-2);
  border: 1px solid var(--border);
  color: var(--text-2);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
}
.theme-btn:hover {
  background: var(--surface-3);
  color: var(--text);
}

.theme-menu {
  position: absolute;
  top: calc(100% + 8px);
  right: 0;
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: 12px;
  padding: 6px;
  min-width: 140px;
  box-shadow: var(--shadow-lg);
  z-index: 200;
}

.theme-opt {
  width: 100%;
  padding: 8px 12px;
  border-radius: 8px;
  background: none;
  border: none;
  cursor: pointer;
  color: var(--text-2);
  font-size: 13px;
  font-weight: 500;
  display: flex;
  align-items: center;
  gap: 8px;
  transition: all 0.15s;
  text-align: left;
}
.theme-opt:hover {
  background: var(--surface-2);
  color: var(--text);
}
.theme-opt.active {
  background: var(--accent-subtle);
  color: var(--accent);
}
.check {
  margin-left: auto;
}

.menu-pop-enter-active,
.menu-pop-leave-active {
  transition: all 0.18s cubic-bezier(0.34, 1.56, 0.64, 1);
}
.menu-pop-enter-from,
.menu-pop-leave-to {
  opacity: 0;
  transform: scale(0.9) translateY(-6px);
}
</style>
