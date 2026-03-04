<template>
  <div class="contact-preview" :class="'layout-' + (data.layout || 'centered')" :style="data.section_bg_color ? { background: data.section_bg_color + ' !important' } : {}">
    <div v-if="data.email" class="contact-row">
      <Mail :size="16" class="icon-inline" /> {{ data.email }}
    </div>
    <div v-if="data.linkedin" class="contact-row">
      <Linkedin :size="16" class="icon-inline" /> {{ data.linkedin }}
    </div>
    <div v-if="data.github" class="contact-row">
      <Github :size="16" class="icon-inline" /> {{ data.github }}
    </div>
    <div v-if="data.location" class="contact-row">
      <MapPin :size="16" class="icon-inline" /> {{ data.location }}
    </div>

    <!-- Custom Contacts -->
    <div
      v-for="contact in data.custom_items"
      :key="contact.id"
      class="contact-row"
    >
      <component
        :is="getIcon(contact.platform)"
        :size="16"
        class="icon-inline"
      />
      {{ contact.label }}: {{ contact.value }}
    </div>

    <div
      v-if="
        !data.email &&
        !data.linkedin &&
        !data.github &&
        (!data.custom_items || data.custom_items.length === 0)
      "
      class="empty-msg"
    >
      ยังไม่มีข้อมูลติดต่อ
    </div>
  </div>
</template>
<script setup lang="ts">
import {
  Mail,
  Linkedin,
  Github,
  MapPin,
  Facebook,
  Twitter,
  Instagram,
  Youtube,
  MessageCircle,
  Phone,
  Link,
} from "lucide-vue-next";

defineProps<{ data: any }>();

function getIcon(platform: string) {
  switch (platform) {
    case "facebook":
      return Facebook;
    case "twitter":
      return Twitter;
    case "instagram":
      return Instagram;
    case "youtube":
      return Youtube;
    case "message-circle":
      return MessageCircle;
    case "phone":
      return Phone;
    default:
      return Link;
  }
}
</script>
<style scoped>
.contact-preview {
  display: flex;
  flex-direction: column;
  gap: 6px;
}
.contact-row {
  font-size: 12px;
  color: var(--muted);
  display: flex;
  align-items: center;
  gap: 8px;
}
.icon-inline {
  vertical-align: middle;
}
.empty-msg {
  font-size: 12px;
  color: var(--muted);
  font-style: italic;
}
</style>
