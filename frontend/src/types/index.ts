// ─── User ─────────────────────────────────────────────────────────────────────
export interface User {
  id: string
  email: string
  name: string
  avatar_url?: string
  plan: 'free' | 'pro' | 'team'
  is_verified: boolean
  created_at: string
}

// ─── Theme ────────────────────────────────────────────────────────────────────
export interface ThemeConfig {
  mode: 'dark' | 'light'
  primary_color: string
  font: string
  layout: 'centered' | 'left' | 'split'
}

// ─── Portfolio ────────────────────────────────────────────────────────────────
export interface Portfolio {
  id: string
  user_id: string
  slug: string
  title: string
  description?: string
  theme: ThemeConfig
  seo_title?: string
  seo_desc?: string
  og_image_url?: string
  is_published: boolean
  has_password: boolean
  expires_at?: string
  view_count: number
  sections?: Section[]
  url?: string
  created_at: string
  updated_at: string
}

// ─── Section ──────────────────────────────────────────────────────────────────
export type SectionType =
  | 'hero'
  | 'experience'
  | 'skills'
  | 'projects'
  | 'education'
  | 'contact'
  | 'ai_chat'
  | 'custom_text'
  | 'stats'
  | 'social'

export interface Section {
  id: string
  portfolio_id: string
  type: SectionType
  position: number
  data: any
  is_visible: boolean
  created_at?: string
}

// ─── Section Data Types ───────────────────────────────────────────────────────
export interface HeroData {
  name: string
  role: string
  tagline: string
  avatar_url: string
  show_hire_me: boolean
  hire_me_link?: string
}

export interface ExperienceItem {
  company: string
  position: string
  start_date: string
  end_date: string
  is_current: boolean
  description: string
  location?: string
}

export interface SkillItem {
  name: string
  level: number // 0-100
  category?: string
}

export interface ProjectItem {
  title: string
  description: string
  image_url?: string
  live_url?: string
  github_url?: string
  tags: string[]
}

export interface EducationItem {
  school: string
  degree: string
  field: string
  start_year: string
  end_year: string
  gpa?: string
}

export interface ContactData {
  email?: string
  phone?: string
  linkedin?: string
  github?: string
  website?: string
  location?: string
}

// ─── Analytics ────────────────────────────────────────────────────────────────
export interface AnalyticsSummary {
  total_views: number
  today_views: number
  week_views: number
  month_views: number
  pdf_downloads: number
  hire_clicks: number
  ai_chat_count: number
  country_stats: { country_code: string; count: number }[]
  daily_views: { date: string; count: number }[]
}

// ─── Block Types (for the builder sidebar) ────────────────────────────────────
export interface BlockType {
  type: SectionType
  label: string
  icon: string
  description: string
}

export const BLOCK_TYPES: BlockType[] = [
  { type: 'hero',        icon: '👤', label: 'Hero Section',    description: 'ชื่อ ตำแหน่ง รูปโปรไฟล์' },
  { type: 'experience',  icon: '💼', label: 'Experience',      description: 'ประสบการณ์การทำงาน' },
  { type: 'skills',      icon: '⚡', label: 'Skills',          description: 'ทักษะและความเชี่ยวชาญ' },
  { type: 'projects',    icon: '🚀', label: 'Projects',        description: 'แสดงผลงาน' },
  { type: 'education',   icon: '🎓', label: 'Education',       description: 'ประวัติการศึกษา' },
  { type: 'contact',     icon: '📬', label: 'Contact',         description: 'ช่องทางการติดต่อ' },
  { type: 'ai_chat',     icon: '🤖', label: 'AI Interview Coach', description: 'แชทบอทสำหรับ HR' },
  { type: 'custom_text', icon: '✍️', label: 'Custom Text',    description: 'ข้อความกำหนดเอง' },
  { type: 'stats',       icon: '📊', label: 'Stats Counter',   description: 'ตัวเลขผลงาน' },
  { type: 'social',      icon: '🌐', label: 'Social Links',    description: 'Social Media Links' },
]
