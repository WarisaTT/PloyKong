// ─── User ─────────────────────────────────────────────────────────────────────
export interface User {
  id: string;
  email: string;
  name: string;
  avatar_url?: string;
  plan: "free" | "pro" | "team";
  is_verified: boolean;
  created_at: string;
}

// ─── Theme ────────────────────────────────────────────────────────────────────
export interface ThemeConfig {
  mode: "dark" | "light" | "system";
  primary_color: string;
  secondary_color?: string;
  palette?: string;
  template?: string;
  bg_color?: string;
  border_color?: string;
  font: string;
  layout: "centered" | "left" | "split";
  show_divider?: boolean;
}

// ─── Portfolio ────────────────────────────────────────────────────────────────
export interface Portfolio {
  id: string;
  user_id: string;
  slug: string;
  title: string;
  description?: string;
  theme: ThemeConfig;
  seo_title?: string;
  seo_desc?: string;
  og_image_url?: string;
  is_published: boolean;
  has_password: boolean;
  expires_at?: string;
  view_count: number;
  sections?: Section[];
  url?: string;
  created_at: string;
  updated_at: string;
}

// ─── Section ──────────────────────────────────────────────────────────────────
export type SectionType =
  | "hero"
  | "experience"
  | "skills"
  | "projects"
  | "education"
  | "contact"
  | "ai_chat"
  | "custom_text"
  | "stats"
  | "social"
  | "certificates";

export interface Section {
  id: string;
  portfolio_id: string;
  type: SectionType;
  position: number;
  data: any;
  is_visible: boolean;
  column_span?: 'full' | 'half'; // half = 1/2 width, sits side-by-side with next half section
  created_at?: string;
}

// ─── Section Data Types ───────────────────────────────────────────────────────
export interface HeroData {
  name: string;
  role: string;
  tagline: string;
  avatar_url: string;
  show_hire_me: boolean;
  hire_me_link?: string;
  show_resume?: boolean;
  resume_url?: string;
  show_contact?: boolean;
  contact_link?: string;
}

export interface ExperienceItem {
  company: string;
  position: string;
  start_date: string;
  end_date: string;
  is_current: boolean;
  description: string;
  location?: string;
  image_urls?: string[]; // newly added for flexible experience block
}

export interface CertificateItem {
  title: string;
  issuer: string;
  date: string;
  description?: string;
  image_url?: string;
}

export interface SkillItem {
  name: string;
  level: number; // 0-100
  category?: string;
}

export interface ProjectItem {
  title: string;
  description: string;
  image_url?: string;
  live_url?: string;
  github_url?: string;
  tags: string[];
}

export interface EducationItem {
  school: string;
  degree: string;
  field: string;
  start_year: string;
  end_year: string;
  gpa?: string;
}

export interface CustomContact {
  id: string; // unique string (e.g., Date.now().toString())
  platform: string; // 'facebook', 'message-circle' (line), 'phone', 'link'
  label: string; // 'Facebook', 'Line ID', 'Phone'
  value: string; // '@username', '081-xxx-xxxx'
  link: string; // actual URL or mailto:/tel: link
}

export interface ContactData {
  email?: string;
  phone?: string;
  linkedin?: string;
  github?: string;
  website?: string;
  location?: string;
  custom_items?: CustomContact[];
}

export interface StatItem {
  label: string;
  value: string;
}

export interface SocialLinkItem {
  platform: string;
  url: string;
  label?: string;
}

export interface AIChatData {
  prompt_hint?: string;
  example_questions?: string[];
}

// ─── Analytics ────────────────────────────────────────────────────────────────
export interface AnalyticsSummary {
  total_views: number;
  today_views: number;
  week_views: number;
  month_views: number;
  pdf_downloads: number;
  hire_clicks: number;
  ai_chat_count: number;
  country_stats: { country_code: string; count: number }[];
  daily_views: { date: string; count: number }[];
}

import { markRaw } from "vue";
import {
  User,
  Briefcase,
  Wrench,
  Rocket,
  GraduationCap,
  Mail,
  Bot,
  Type,
  BarChart2,
  Globe,
  Award,
} from "lucide-vue-next";

// ─── Block Types (for the builder sidebar) ────────────────────────────────────
export interface BlockType {
  type: SectionType;
  label: string;
  icon: any;
  description: string;
}

export const BLOCK_TYPES: BlockType[] = [
  {
    type: "hero",
    icon: markRaw(User),
    label: "Hero Section",
    description: "ชื่อ ตำแหน่ง รูปโปรไฟล์",
  },
  {
    type: "experience",
    icon: markRaw(Briefcase),
    label: "Experience",
    description: "ประสบการณ์การทำงาน",
  },
  {
    type: "skills",
    icon: markRaw(Wrench),
    label: "Skills",
    description: "ทักษะและความเชี่ยวชาญ",
  },
  {
    type: "projects",
    icon: markRaw(Rocket),
    label: "Projects",
    description: "แสดงผลงาน",
  },
  {
    type: "education",
    icon: markRaw(GraduationCap),
    label: "Education",
    description: "ประวัติการศึกษา",
  },
  {
    type: "contact",
    icon: markRaw(Mail),
    label: "Contact",
    description: "ข้อมูลติดต่อ",
  },
  {
    type: "certificates",
    icon: markRaw(Award),
    label: "Certificates",
    description: "ใบประกาศณียบัตรและรางวัล",
  },
  {
    type: "ai_chat",
    icon: markRaw(Bot),
    label: "AI Chat",
    description: "ระบบ AI ตอบคำถาม (ทดลอง)",
  },
  {
    type: "custom_text",
    icon: markRaw(Type),
    label: "Custom Text",
    description: "ส่วนข้อความอิสระ",
  },
  {
    type: "stats",
    icon: markRaw(BarChart2),
    label: "Stats Counter",
    description: "ตัวเลขผลงาน",
  },
  {
    type: "social",
    icon: markRaw(Globe),
    label: "Social Links",
    description: "Social Media Links",
  },
];
