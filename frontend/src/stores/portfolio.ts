import { defineStore } from "pinia";
import { ref, watch } from "vue";
import { showError } from "@/utils/alert";
import { portfolioAPI, sectionAPI } from "@/api";
import type { Portfolio, Section } from "@/types";

export const usePortfolioStore = defineStore("portfolio", () => {
  // ─── State ────────────────────────────────────────────────────────────────
  const portfolios = ref<Portfolio[]>([]);
  const activePortfolio = ref<Portfolio | null>(null);
  const sections = ref<Section[]>([]);
  const loading = ref(false);
  const saving = ref(false);
  const error = ref<string | null>(null);
  const gapCount = ref(0);


  // ─── Portfolio CRUD ───────────────────────────────────────────────────────
  async function fetchPortfolios() {
    loading.value = true;
    try {
      const { data } = await portfolioAPI.list();
      portfolios.value = data.data || [];
    } catch (e: any) {
      error.value = e.response?.data?.error || "Failed to load portfolios";
    } finally {
      loading.value = false;
    }
  }

  async function loadPortfolio(id: string) {
    loading.value = true;
    try {
      const { data } = await portfolioAPI.getById(id);
      activePortfolio.value = data.data;
      sections.value = data.data.sections || [];
    } catch (e: any) {
      error.value = e.response?.data?.error || "Failed to load portfolio";
    } finally {
      loading.value = false;
    }
  }

  async function createPortfolio(payload: {
    slug: string;
    title: string;
    theme?: any;
  }) {
    loading.value = true;
    try {
      const { data } = await portfolioAPI.create(payload);
      portfolios.value.unshift(data.data);
      return data.data;
    } catch (e: any) {
      error.value = e.response?.data?.error || "Failed to create portfolio";
      return null;
    } finally {
      loading.value = false;
    }
  }

  async function savePortfolio(updates: Partial<Portfolio>) {
    if (!activePortfolio.value) return;
    saving.value = true;
    try {
      await portfolioAPI.update(activePortfolio.value.id, updates);
      Object.assign(activePortfolio.value, updates);
    } catch (e: any) {
      const msg = e.response?.data?.error || "Failed to save";
      error.value = msg;
      showError(msg);
    } finally {
      saving.value = false;
    }
  }

  async function publishPortfolio() {
    if (!activePortfolio.value) return;
    await portfolioAPI.publish(activePortfolio.value.id);
    activePortfolio.value.is_published = true;
  }

  async function unpublishPortfolio() {
    if (!activePortfolio.value) return;
    await portfolioAPI.unpublish(activePortfolio.value.id);
    activePortfolio.value.is_published = false;
  }

  async function duplicatePortfolio(id: string) {
    loading.value = true;
    try {
      await portfolioAPI.duplicate(id);
      await fetchPortfolios(); // Refresh list to show the copied portfolio
    } catch (e: any) {
      error.value = e.response?.data?.error || "Failed to duplicate portfolio";
    } finally {
      loading.value = false;
    }
  }

  async function deletePortfolio(id: string) {
    await portfolioAPI.delete(id);
    portfolios.value = portfolios.value.filter((p) => p.id !== id);
    if (activePortfolio.value?.id === id) {
      activePortfolio.value = null;
    }
  }

  async function fetchGapCount() {
    try {
      const { data } = await portfolioAPI.listGaps();
      gapCount.value = data.data?.length || 0;
    } catch (e) {
      console.error("Failed to load gaps");
    }
  }

  // ─── Section CRUD ─────────────────────────────────────────────────────────
  async function addSection(type: string, afterSectionId?: string | null) {
    if (!activePortfolio.value) return;

    let position = 0;
    if (afterSectionId) {
      const selectedIdx = sections.value.findIndex(s => s.id === afterSectionId);
      if (selectedIdx !== -1) {
        position = sections.value[selectedIdx].position + 1;
      } else {
        const maxPos = sections.value.length > 0
          ? Math.max(...sections.value.map(s => s.position))
          : -1;
        position = maxPos + 1;
      }
    } else {
      const maxPos = sections.value.length > 0
        ? Math.max(...sections.value.map(s => s.position))
        : -1;
      position = maxPos + 1;
    }

    const { data } = await sectionAPI.create(activePortfolio.value.id, {
      type,
      position,
      data: getDefaultSectionData(type),
    });

    // Reload portfolio to sync all shifted positions from backend
    await loadPortfolio(activePortfolio.value.id);
    return data.data.id;
  }

  // Debounce for updateSection to prevent lag on every keystroke
  let updateTimeout: any = null;
  const pendingUpdates = new Map<string, Partial<Section>>();

  async function updateSection(sectionId: string, updates: Partial<Section>) {
    if (!activePortfolio.value) return;

    // Update local state immediately for snappy UI
    const idx = sections.value.findIndex((s) => s.id === sectionId);
    if (idx !== -1) {
      Object.assign(sections.value[idx], updates);
    }

    // Merge updates for this section
    const currentPending = pendingUpdates.get(sectionId) || {};
    pendingUpdates.set(sectionId, { ...currentPending, ...updates });

    if (updateTimeout) clearTimeout(updateTimeout);

    saving.value = true;
    updateTimeout = setTimeout(async () => {
      try {
        const mergedUpdates = pendingUpdates.get(sectionId);
        if (mergedUpdates) {
          pendingUpdates.delete(sectionId);
          await sectionAPI.update(sectionId, mergedUpdates);
        }
      } catch (e: any) {
        if (e.response?.status === 404) {
          // Section missing from DB — try to re-persist it
          const lost = sections.value.find(s => s.id === sectionId);
          if (lost && activePortfolio.value) {
            try {
              console.warn(`[store] Re-creating lost section: ${sectionId} (${lost.type})`);
              const { data } = await sectionAPI.create(activePortfolio.value.id, {
                type: lost.type,
                position: lost.position,
                data: lost.data, // use latest local data
                hide_title: lost.hide_title,
                hide_divider: lost.hide_divider,
                include_in_resume: lost.include_in_resume,
              });
              // Update the local section to use the new DB-assigned ID
              const idx = sections.value.findIndex(s => s.id === sectionId);
              if (idx !== -1) {
                sections.value[idx].id = data.data.id;
              }
            } catch (createErr) {
              console.error('[store] Failed to re-create section:', createErr);
            }
          }
        } else {
          showError("Failed to save changes");
        }
      } finally {
        if (pendingUpdates.size === 0) {
          saving.value = false;
        }
      }
    }, 1000); // 1 second debounce
  }

  async function deleteSection(sectionId: string) {
    if (!activePortfolio.value) return;
    await sectionAPI.delete(sectionId);
    sections.value = sections.value.filter((s) => s.id !== sectionId);
  }

  async function duplicateSection(sectionId: string) {
    if (!activePortfolio.value) return;
    loading.value = true;
    try {
      await sectionAPI.duplicate(sectionId);
      await loadPortfolio(activePortfolio.value.id); // Reload to reflect shifted positions and new section
    } catch (e: any) {
      error.value = e.response?.data?.error || "Failed to duplicate section";
    } finally {
      loading.value = false;
    }
  }

  async function reorderSections(newOrder: Section[]) {
    if (!activePortfolio.value) return;
    sections.value = newOrder.map((s, i) => ({ ...s, position: i }));
    const order = sections.value.map((s) => ({
      id: s.id,
      position: s.position,
    }));
    await sectionAPI.reorder(order);
  }

  async function toggleSectionVisibility(sectionId: string) {
    const section = sections.value.find((s) => s.id === sectionId);
    if (!section || !activePortfolio.value) return;
    section.is_visible = !section.is_visible;
    try {
      await sectionAPI.update(sectionId, { is_visible: section.is_visible });
    } catch (e: any) {
      if (e.response?.status === 404) return;
      throw e;
    }
  }

  async function toggleSectionColumnSpan(sectionId: string) {
    const section = sections.value.find((s) => s.id === sectionId);
    if (!section || !activePortfolio.value) return;

    // HERO is ALWAYS FULL WIDTH
    if (section.type === 'hero') return;

    const next = (section.column_span || 'full') === 'full' ? 'half' : 'full';
    section.column_span = next;
    sectionAPI.update(sectionId, { column_span: next });
  }

  // ─── Helpers ──────────────────────────────────────────────────────────────
  function getDefaultSectionData(type: string): any {
    const defaults: Record<string, any> = {
      hero: {
        name: "",
        role: "",
        tagline: "",
        avatar_url: "",
        show_hire_me: true,
      },
      experience: { items: [] },
      skills: { items: [] },
      projects: { items: [] },
      education: { items: [] },
      contact: { email: "", linkedin: "", github: "", location: "" },
      ai_chat: {
        enabled: true,
        greeting: "Hi! Ask me anything about my work.",
      },
      custom_text: { content: "", alignment: "left" },
      stats: { items: [] },
      social: { links: [] },
    };
    return defaults[type] || {};
  }

  return {
    portfolios,
    activePortfolio,
    sections,
    loading,
    saving,
    error,
    fetchPortfolios,
    loadPortfolio,
    createPortfolio,
    savePortfolio,
    publishPortfolio,
    unpublishPortfolio,
    duplicatePortfolio,
    deletePortfolio,
    addSection,
    updateSection,
    duplicateSection,
    deleteSection,
    reorderSections,
    toggleSectionVisibility,
    toggleSectionColumnSpan,
    gapCount,
    fetchGapCount,
  };
});
