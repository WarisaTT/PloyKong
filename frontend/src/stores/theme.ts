import { defineStore } from "pinia";
import { useColorMode } from "@vueuse/core";
import { computed } from "vue";

export const useThemeStore = defineStore("theme", () => {
  // useColorMode automatically handles localStorage and system preference sync
  // Default to 'dark' for the initial app style, but it will read from localStorage if previously set.
  const colorMode = useColorMode({
    emitAuto: true,
    modes: {
      light: "theme-light",
      dark: "theme-dark",
    },
    attribute: "class",
    initialValue: "auto", // default PloyKong app theme
  });

  const mode = computed({
    get: () => colorMode.value,
    set: (v) => {
      colorMode.value = v;
    },
  });

  // Expose the current resolved theme class for the app wrapper
  const themeClass = computed(() => {
    if (colorMode.value === "auto") {
      return "theme-system"; // our CSS handles the actual colors using prefers-color-scheme
    }
    return `theme-${colorMode.value}`;
  });

  const resolvedDark = computed(() => {
    if (colorMode.value === "auto") {
      return window.matchMedia("(prefers-color-scheme: dark)").matches;
    }
    return colorMode.value === "dark";
  });

  function setMode(newMode: "light" | "dark" | "auto") {
    colorMode.value = newMode;
  }

  return {
    mode,
    themeClass,
    resolvedDark,
    setMode,
  };
});


