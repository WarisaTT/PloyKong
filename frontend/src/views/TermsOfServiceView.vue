<template>
  <div class="policy-page">
    <div class="policy-container">
      <RouterLink to="/" class="back-link">← กลับหน้าหลัก</RouterLink>
      <header class="policy-header">
        <h1>ข้อกำหนดและเงื่อนไขการใช้งาน (Terms of Service)</h1>
        <p>อัปเดตล่าสุด: 6 มีนาคม 2569</p>
      </header>

      <section class="policy-content">
        <h2>1. การยอมรับข้อกำหนด</h2>
        <p>การสมัครสมาชิกหรือเข้าใช้งาน PloyKong ถือว่าคุณยอมรับข้อกำหนดและเงื่อนไขเหล่านี้ทุกประการ</p>

        <h2>2. บัญชีผู้ใช้</h2>
        <p>คุณมีหน้าที่รับผิดชอบในการรักษาความลับของรหัสผ่านและข้อมูลบัญชีของคุณ และรับผิดชอบต่อกิจกรรมทั้งหมดที่เกิดขึ้นภายใต้บัญชีของคุณ</p>

        <h2>3. เนื้อหาจากผู้ใช้ (User Generated Content)</h2>
        <p>คุณยังคงเป็นเจ้าของลิขสิทธิ์ในเนื้อหาที่คุุณสร้างขึ้น อย่างไรก็ตาม คุณตกลงว่า:</p>
        <ul>
          <li>เนื้อหาต้องไม่ผิดกฎหมาย ไม่เป็นการละเมิดสิทธิผู้อื่น</li>
          <li>ไม่ใช้บริการของเราเพื่อจุดประสงค์ในการหลอกลวง (Phishing) หรือสแปม</li>
          <li>เราขอสงวนสิทธิ์ในการระงับบัญชีหรือลบเนื้อหาที่ละเมิดข้อกำหนดโดยไม่ต้องแจ้งให้ทราบล่วงหน้า</li>
        </ul>

        <h2>4. การใช้ AI</h2>
        <p>บริการของเรามีการใช้เทคโนโลยี AI จากบุคคลที่สาม (เช่น OpenAI) เพื่อช่วยในการสร้างเนื้อหา คุณควรตรวจสอบความถูกต้องของข้อมูลก่อนนำไปใช้งานจริง เราไม่ได้รับผิดชอบต่อความผิดพลาดที่เกิดจากเนื้อหาที่สร้างโดย AI</p>

        <h2>5. การจำกัดความรับผิดชอบ</h2>
        <p>PloyKong ให้บริการ "ตามสภาพที่เป็นอยู่" เราไม่รับรองว่าบริการจะไม่มีข้อบกพร่องหรือการขัดข้อง และจะไม่รับผิดชอบต่อความเสียหายใดๆ ที่เกิดจากการใช้หรือการไม่สามารถเข้าใช้งานบริการได้</p>

        <h2>6. การเปลี่ยนแปลงข้อกำหนด</h2>
        <p>เราอาจปรับปรุงข้อกำหนดเหล่านี้เป็นครั้งคราว การใช้งานต่อหลังจากมีการเปลี่ยนแปลงถือว่าคุณยอมรับข้อกำหนดใหม่นั้น</p>
      </section>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router';
</script>

<style scoped>
/* ใช้สไตล์เดียวกับ PrivacyPolicyView เพื่อความสม่ำเสมอ */
.policy-page {
  min-height: 100vh;
  background: var(--bg);
  color: var(--text);
  padding: 80px 24px;
  line-height: 1.8;
}
.policy-container {
  max-width: 800px;
  margin: 0 auto;
}
.back-link {
  color: var(--primary);
  text-decoration: none;
  font-weight: 600;
  display: inline-block;
  margin-bottom: 32px;
}
.policy-header {
  margin-bottom: 48px;
  border-bottom: 1px solid var(--border);
  padding-bottom: 24px;
}
.policy-header h1 {
  font-size: 32px;
  margin-bottom: 8px;
}
.policy-header p {
  color: var(--muted);
  font-size: 14px;
}
.policy-content h2 {
  margin-top: 40px;
  font-size: 20px;
  color: var(--primary);
}
.policy-content ul {
  padding-left: 20px;
  margin: 16px 0;
}
.policy-content li {
  margin-bottom: 12px;
}
@media (max-width: 640px) {
  .policy-header h1 {
    font-size: 24px;
  }
}
</style>
