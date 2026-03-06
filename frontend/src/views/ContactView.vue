<template>
  <div class="contact-page">
    <div class="contact-container">
      <RouterLink to="/" class="back-link">← กลับหน้าหลัก</RouterLink>
      <header class="contact-header">
        <h1>ติดต่อเรา (Contact Us)</h1>
        <p>หากคุณมีคำถาม ข้อสงสัย หรือข้อเสนอแนะ ทีมงาน PloyKong ยินดีรับฟังครับ</p>
      </header>

      <div class="contact-grid">
        <!-- Contact Info -->
        <div class="contact-info">
          <div class="info-card">
            <div class="info-icon">📧</div>
            <div class="info-body">
              <h3>อีเมล</h3>
              <p>fair2708@gmail.com</p>
            </div>
          </div>
          <div class="info-card">
            <div class="info-icon">💻</div>
            <div class="info-body">
              <h3>GitHub</h3>
              <a href="https://github.com/WarisaTT" target="_blank">github.com/WarisaTT</a>
            </div>
          </div>
          <div class="info-card">
            <div class="info-icon">📍</div>
            <div class="info-body">
              <h3>ที่อยู่</h3>
              <p>Bangkok, Thailand</p>
            </div>
          </div>
        </div>

        <!-- Contact Form (Mockup) -->
        <div class="contact-form-card">
          <form @submit.prevent="handleSubmit">
            <div class="form-group">
              <label>ชื่อของคุณ</label>
              <input type="text" v-model="form.name" placeholder="ชื่อ-นามสกุล" required />
            </div>
            <div class="form-group">
              <label>อีเมลติดต่อ</label>
              <input type="email" v-model="form.email" placeholder="example@gmail.com" required />
            </div>
            <div class="form-group">
              <label>หัวข้อ</label>
              <input type="text" v-model="form.subject" placeholder="ระบุหัวข้อที่ต้องการติดต่อ" required />
            </div>
            <div class="form-group">
              <label>ข้อความ</label>
              <textarea v-model="form.message" rows="5" placeholder="พิมพ์ข้อความของคุณที่นี่..." required></textarea>
            </div>
            <button type="submit" class="submit-btn" :disabled="loading">
              {{ loading ? 'กำลังส่ง...' : 'ส่งข้อความ' }}
            </button>
            <p v-if="success" class="success-msg">ส่งข้อความเรียบร้อยแล้ว! เราจะติดต่อกลับโดยเร็วที่สุด</p>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
import { RouterLink } from 'vue-router';

const form = reactive({
  name: '',
  email: '',
  subject: '',
  message: ''
});

const loading = ref(false);
const success = ref(false);

const handleSubmit = () => {
  loading.value = true;
  // Simulate API call
  setTimeout(() => {
    loading.value = false;
    success.value = true;
    form.name = '';
    form.email = '';
    form.subject = '';
    form.message = '';
    setTimeout(() => success.value = false, 5000);
  }, 1500);
};
</script>

<style scoped>
.contact-page {
  min-height: 100vh;
  background: var(--bg);
  color: var(--text);
  padding: 80px 24px;
}
.contact-container {
  max-width: 1000px;
  margin: 0 auto;
}
.back-link {
  color: var(--primary);
  text-decoration: none;
  font-weight: 600;
  display: inline-block;
  margin-bottom: 32px;
}
.contact-header {
  margin-bottom: 60px;
  text-align: center;
}
.contact-header h1 {
  font-size: 36px;
  margin-bottom: 12px;
}
.contact-header p {
  color: var(--muted);
}

.contact-grid {
  display: grid;
  grid-template-columns: 1fr 1.5fr;
  gap: 40px;
}

.contact-info {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.info-card {
  display: flex;
  align-items: center;
  gap: 20px;
  background: var(--surface);
  padding: 24px;
  border-radius: 16px;
  border: 1px solid var(--border);
}

.info-icon {
  font-size: 24px;
  width: 48px;
  height: 48px;
  background: var(--primary-glow);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.info-body h3 {
  font-size: 16px;
  margin-bottom: 4px;
}
.info-body p, .info-body a {
  color: var(--muted);
  text-decoration: none;
  font-size: 14px;
}

.contact-form-card {
  background: var(--surface);
  padding: 40px;
  border-radius: 24px;
  border: 1px solid var(--border);
}

.form-group {
  margin-bottom: 20px;
}
.form-group label {
  display: block;
  font-size: 14px;
  font-weight: 600;
  margin-bottom: 8px;
}
.form-group input, .form-group textarea {
  width: 100%;
  padding: 12px 16px;
  background: var(--bg);
  border: 1px solid var(--border);
  border-radius: 10px;
  color: var(--text);
  font-family: inherit;
}
.form-group input:focus, .form-group textarea:focus {
  outline: none;
  border-color: var(--primary);
}

.submit-btn {
  width: 100%;
  padding: 14px;
  background: var(--primary);
  color: white;
  border: none;
  border-radius: 12px;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.2s;
}
.submit-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px var(--primary-glow);
}
.submit-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  transform: none;
}

.success-msg {
  color: #10b981;
  font-size: 14px;
  margin-top: 16px;
  text-align: center;
}

@media (max-width: 768px) {
  .contact-grid {
    grid-template-columns: 1fr;
  }
  .contact-header h1 {
    font-size: 28px;
  }
}
</style>
