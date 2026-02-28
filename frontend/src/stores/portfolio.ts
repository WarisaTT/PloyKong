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

  // show friendly alert when an error is set
  watch(error, (v) => {
    if (!v) return;
    try {
      const friendly = v;
      showError(String(friendly));
    } catch (e) {
      // eslint-disable-next-line no-console
      console.error("Alert failed", e);
    }
  });

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
      error.value = e.response?.data?.error || "Failed to save";
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

  // ─── Section CRUD ─────────────────────────────────────────────────────────
  async function addSection(type: string) {
    if (!activePortfolio.value) return;

    // Ensure we always append to the very bottom, beneath any existing blocks
    const maxPos = sections.value.length > 0
      ? Math.max(...sections.value.map(s => s.position))
      : -1;
    const position = maxPos + 1;

    const { data } = await sectionAPI.create(activePortfolio.value.id, {
      type,
      position,
      data: getDefaultSectionData(type),
    });
    sections.value.push({
      id: data.data.id,
      portfolio_id: activePortfolio.value.id,
      type,
      position,
      data: getDefaultSectionData(type),
      is_visible: true,
    } as Section);
    return data.data.id;
  }

  async function updateSection(sectionId: string, sectionData: any) {
    if (!activePortfolio.value) return;
    saving.value = true;
    try {
      await sectionAPI.update(sectionId, { data: sectionData });
      const idx = sections.value.findIndex((s) => s.id === sectionId);
      if (idx !== -1) sections.value[idx].data = sectionData;
    } finally {
      saving.value = false;
    }
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

  function toggleSectionVisibility(sectionId: string) {
    const section = sections.value.find((s) => s.id === sectionId);
    if (!section || !activePortfolio.value) return;
    section.is_visible = !section.is_visible;
    sectionAPI.update(sectionId, { is_visible: section.is_visible });
  }

  function toggleSectionColumnSpan(sectionId: string) {
    const section = sections.value.find((s) => s.id === sectionId);
    if (!section || !activePortfolio.value) return;
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
  };
});
