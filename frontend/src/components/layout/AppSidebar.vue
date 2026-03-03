<template>
  <!-- Mobile Header -->
  <header class="mobile-nav">
    <div class="sidebar-logo">
      <div class="logo-pk">
        <img src="/favicon.png" alt="PloyKong Logo" class="logo-img" />
      </div>
      <span class="logo-text">Ploy<em>Kong</em></span>
    </div>
    <div class="mobile-nav-actions">
      <button 
        class="mobile-icon-btn menu-toggle" 
        @click="isMenuOpen = !isMenuOpen"
        :title="isMenuOpen ? 'Close Menu' : 'Open Menu'"
      >
        <Menu v-if="!isMenuOpen" :size="24" />
        <X v-else :size="24" />
      </button>
    </div>
  </header>

  <!-- Mobile Menu Overlay -->
  <Transition name="slide-up">
    <div v-if="isMenuOpen" class="mobile-menu-overlay" @click.self="isMenuOpen = false">
      <div class="mobile-menu-card">
        <nav class="mobile-menu-nav">
          <RouterLink
            to="/dashboard"
            class="mobile-nav-item"
            exact-active-class="active"
            @click="isMenuOpen = false"
          >
            <BarChart2 :size="20" /> Dashboard
          </RouterLink>
          <RouterLink
            to="/ai-center"
            class="mobile-nav-item"
            exact-active-class="active"
            @click="isMenuOpen = false"
          >
            <Sparkles :size="20" /> AI Knowledge Center
            <span v-if="gapCount > 0" class="nav-badge">{{ gapCount }}</span>
          </RouterLink>
          <RouterLink
            to="/portfolios/new"
            class="mobile-nav-item"
            exact-active-class="active"
            @click="isMenuOpen = false"
          >
            <Plus :size="20" /> สร้างพอร์ตใหม่
          </RouterLink>
          <RouterLink 
            to="/settings" 
            class="mobile-nav-item" 
            exact-active-class="active"
            @click="isMenuOpen = false"
          >
            <Settings :size="20" /> Settings
          </RouterLink>
          
          <div class="mobile-theme-row">
            <span class="theme-label">Theme</span>
            <div class="theme-toggle-group">
              <button @click="themeStore.setMode('light')" :class="{ active: themeStore.mode === 'light' }"><Sun :size="16" /></button>
              <button @click="themeStore.setMode('dark')" :class="{ active: themeStore.mode === 'dark' }"><Moon :size="16" /></button>
              <button @click="themeStore.setMode('auto')" :class="{ active: themeStore.mode === 'auto' }"><Monitor :size="16" /></button>
            </div>
          </div>
        </nav>

        <div class="mobile-menu-footer">
          <div class="sidebar-user" @click="handleLogout">
            <div class="user-avatar">{{ userInitial }}</div>
            <div class="user-meta">
              <div class="user-name">{{ authStore.user?.name }}</div>
              <div class="user-plan">{{ authStore.user?.plan }} plan</div>
            </div>
            <LogOut :size="18" class="logout-icon" />
          </div>
        </div>
      </div>
    </div>
  </Transition>

  <!-- Desktop Sidebar -->
  <aside class="sidebar">
    <div class="sidebar-logo">
      <div class="logo-pk">
        <img src="/favicon.png" alt="PloyKong Logo" class="logo-img" />
      </div>
      <span class="logo-text">Ploy<em>Kong</em></span>
    </div>
    
    <nav class="sidebar-nav">
      <RouterLink
        to="/dashboard"
        class="nav-item"
        exact-active-class="active"
      >
        <BarChart2 :size="16" /> Dashboard
      </RouterLink>
      <RouterLink
        to="/ai-center"
        class="nav-item"
        exact-active-class="active"
      >
        <Sparkles :size="16" /> AI Knowledge Center
        <span v-if="gapCount > 0" class="nav-badge">{{ gapCount }}</span>
      </RouterLink>
      <RouterLink
        to="/portfolios/new"
        class="nav-item"
        exact-active-class="active"
      >
        <Plus :size="16" /> สร้างพอร์ตใหม่
      </RouterLink>
      <RouterLink to="/settings" class="nav-item" exact-active-class="active">
        <Settings :size="16" /> Settings
      </RouterLink>

      <div class="sidebar-theme-toggle">
        <div class="theme-label">Theme</div>
        <div class="theme-toggle-group">
          <button
            :class="['theme-btn', { active: themeStore.mode === 'dark' }]"
            @click="themeStore.setMode('dark')"
            title="Dark Mode"
          >
            <Moon :size="16" />
          </button>
          <button
            :class="['theme-btn', { active: themeStore.mode === 'light' }]"
            @click="themeStore.setMode('light')"
            title="Light Mode"
          >
            <Sun :size="16" />
          </button>
          <button
            :class="['theme-btn', { active: themeStore.mode === 'auto' }]"
            @click="themeStore.setMode('auto')"
            title="System Setting"
          >
            <Monitor :size="16" />
          </button>
        </div>
      </div>
    </nav>

    <div class="sidebar-user" @click="handleLogout">
      <div class="user-avatar">{{ userInitial }}</div>
      <div class="user-meta">
        <div class="user-name">{{ authStore.user?.name }}</div>
        <div class="user-plan">{{ authStore.user?.plan }} plan</div>
      </div>
      <LogOut :size="16" class="logout-icon" title="Logout" />
    </div>
  </aside>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { RouterLink, useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth';
import { useThemeStore } from '@/stores/theme';
import { usePortfolioStore } from '@/stores/portfolio';
import { portfolioAPI } from '@/api';
import { 
  BarChart2, Sparkles, Plus, Settings, LogOut, 
  Moon, Sun, Monitor, Menu, X 
} from 'lucide-vue-next';

const isMenuOpen = ref(false);

const authStore = useAuthStore();
const themeStore = useThemeStore();
const portfolioStore = usePortfolioStore();
const router = useRouter();

const userInitial = computed(() => authStore.user?.name?.[0]?.toUpperCase() || 'U');
const gapCount = computed(() => portfolioStore.gapCount);

async function handleLogout() {
  await authStore.logout();
  router.push('/login');
}

onMounted(() => {
  portfolioStore.fetchGapCount();
});
</script>

<style scoped>
/* Sidebar Styles */
.sidebar {
  background: var(--sidebar-bg);
  border-right: 1px solid var(--border);
  padding: 24px 16px;
  display: flex;
  flex-direction: column;
  position: sticky;
  top: 0;
  height: 100vh;
  z-index: 100;
}

.sidebar-logo {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 32px;
  padding-left: 8px;
}

.logo-pk {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  background: linear-gradient(135deg, var(--indigo), var(--purple));
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.logo-text {
  font-family: 'Syne', sans-serif; font-size: 18px; font-weight: 800;
  letter-spacing: -0.3px;
}

.logo-text em{ font-style: normal; background: var(--grad-pk); -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text; }

.sidebar-nav {
  display: flex;
  flex-direction: column;
  gap: 6px;
  flex: 1;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 14px;
  border-radius: 12px;
  color: var(--muted);
  text-decoration: none;
  font-size: 14px;
  font-weight: 600;
  transition: all 0.2s;
  position: relative;
}

.nav-item:hover {
  background: var(--bg2);
  color: var(--text);
}

.nav-item.active {
  background: rgba(79, 70, 229, 0.1);
  color: var(--indigo);
}

.nav-badge {
  position: absolute;
  right: 12px;
  background: var(--danger, #ef4444);
  color: white;
  font-size: 10px;
  padding: 2px 6px;
  border-radius: 10px;
  font-weight: 800;
  box-shadow: 0 2px 8px rgba(239, 68, 68, 0.3);
}

.sidebar-theme-toggle {
  margin-top: 24px;
  padding-top: 24px;
  border-top: 1px solid var(--border);
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.theme-label {
  font-size: 11px;
  font-weight: 600;
  color: var(--muted);
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.theme-toggle-group {
  display: flex;
  background: var(--bg2);
  border-radius: 8px;
  padding: 4px;
  gap: 4px;
}

.theme-btn {
  flex: 1;
  background: transparent;
  border: none;
  color: var(--muted);
  border-radius: 6px;
  padding: 6px 0;
  cursor: pointer;
  transition: all 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
}

.theme-btn:hover {
  background: var(--surface);
  color: var(--text);
}

.theme-btn.active {
  background: var(--surface);
  color: var(--indigo);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.sidebar-user {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 12px;
  border-radius: 12px;
  border: 1px solid var(--border);
  cursor: pointer;
  transition: all 0.15s;
  margin-top: 16px;
  background: var(--surface);
}

.sidebar-user:hover {
  background: var(--bg2);
}

.user-avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background: linear-gradient(135deg, var(--indigo), var(--purple));
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 15px;
  font-weight: 700;
  color: #fff;
  flex-shrink: 0;
}

.user-meta {
  flex: 1;
  min-width: 0;
}

.user-name {
  font-size: 13px;
  font-weight: 600;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.user-plan {
  font-size: 11px;
  color: var(--neon-purple);
}

.logout-icon {
  color: var(--muted);
  flex-shrink: 0;
}

/* Mobile Nav */
.mobile-nav {
  display: none;
}

@media (max-width: 768px) {
  .sidebar {
    display: none;
  }
  
  .mobile-nav {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 12px 20px;
    background: var(--surface);
    border-bottom: 1px solid var(--border);
    position: sticky;
    top: 0;
    z-index: 50;
  }
  
  .mobile-nav-actions {
    display: flex;
    align-items: center;
    gap: 12px;
  }
  
  .mobile-icon-btn {
    background: transparent;
    border: none;
    color: var(--muted);
    cursor: pointer;
    padding: 8px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 8px;
    transition: all 0.2s;
  }
  
  .mobile-icon-btn:hover {
    background: var(--bg-2);
    color: var(--text);
  }
  
  .logout-btn {
    color: var(--danger, #ef4444);
  }

  /* Mobile menu overlay */
  .mobile-menu-overlay {
    position: fixed;
    inset: 0;
    background: rgba(0,0,0,0.4);
    backdrop-filter: blur(4px);
    z-index: 1000;
    display: flex;
    flex-direction: column;
    justify-content: flex-end;
  }
  
  .mobile-menu-card {
    background: var(--surface);
    border-radius: 20px 20px 0 0;
    padding: 24px 20px 40px;
    box-shadow: 0 -10px 40px rgba(0,0,0,0.2);
    border-top: 1px solid var(--border);
  }

  .mobile-menu-nav {
    display: flex;
    flex-direction: column;
    gap: 8px;
    margin-bottom: 24px;
  }

  .mobile-nav-item {
    display: flex;
    align-items: center;
    gap: 16px;
    padding: 14px 18px;
    border-radius: 12px;
    color: var(--text);
    text-decoration: none;
    font-size: 16px;
    font-weight: 700;
    background: var(--bg2);
    position: relative;
  }
  
  .mobile-nav-item.active {
    background: var(--indigo);
    color: #fff;
  }

  .mobile-theme-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px 18px;
    margin-top: 10px;
  }

  .mobile-menu-footer .sidebar-user {
    margin-top: 0;
    background: var(--bg2);
  }

  /* Transition */
  .slide-up-enter-active, .slide-up-leave-active {
    transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  }
  .slide-up-enter-from, .slide-up-leave-to {
    transform: translateY(100%);
    opacity: 0;
  }
}
</style>
