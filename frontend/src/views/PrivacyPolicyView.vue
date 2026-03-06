<template>
  <div class="policy-page">
    <div class="policy-container">
      <RouterLink to="/" class="back-link">← กลับหน้าหลัก</RouterLink>
      <header class="policy-header">
        <h1>นโยบายความเป็นส่วนตัว (Privacy Policy)</h1>
        <p>อัปเดตล่าสุด: 6 มีนาคม 2569</p>
      </header>

      <section class="policy-content">
        <h2>1. ข้อมูลที่เราเก็บรวบรวม</h2>
        <p>เราเก็บรวบรวมข้อมูลที่คุณเลือกที่จะแบ่งปันกับเราเมื่อคุณใช้งาน PloyKong:</p>
        <ul>
          <li><strong>ข้อมูลบัญชี:</strong> ชื่อ, อีเมล, และข้อมูลโปรไฟล์เมื่อคุณสมัครสมาชิก</li>
          <li><strong>ข้อมูลพอร์ตโฟลิโอ:</strong> ข้อมูลทั้งหมดที่คุณใส่ใน Resume เช่น ประวัติการทำงาน, การศึกษา, ทักษะ และรูปภาพ</li>
          <li><strong>ข้อมูลการใช้งาน:</strong> เราเก็บสถิติการเข้าชมแบบไม่ระบุตัวตน (IP Address, ประเภทเบราว์เซอร์) เพื่อวิเคราะห์ประสิทธิภาพของเว็บไซต์</li>
        </ul>

        <h2>2. การใช้ข้อมูล</h2>
        <p>เราใช้ข้อมูลของคุณเพื่อ:</p>
        <ul>
          <li>ให้บริการและรักษาความปลอดภัยของบัญชีผู้ใช้</li>
          <li>แสดงผลพอร์ตโฟลิโอของคุณต่อสาธารณะตามที่คุณตั้งค่าไว้</li>
          <li>พัฒนาและปรับปรุงฟีเจอร์ AI ของเรา (โดยไม่ใช้ข้อมูลส่วนบุคคลระบุตัวตนไปเทรนโมเดลภายนอก)</li>
        </ul>

        <h2>3. การแบ่งปันข้อมูล</h2>
        <p>เราจะไม่ขายข้อมูลส่วนบุคคลของคุณให้กับบุคคลที่สาม เราจะแบ่งปันข้อมูลเฉพาะเมื่อ:</p>
        <ul>
          <li>คุณเลือกที่จะ Publish พอร์ตโฟลิโอเป็นสาธารณะ</li>
          <li>มีความจำเป็นตามกฎหมาย</li>
        </ul>

        <h2>4. ความปลอดภัย</h2>
        <p>เราใช้มาตรการรักษาความปลอดภัยที่ได้มาตรฐานสากล รวมถึงการเข้ารหัสรหัสผ่าน (Hashing) และการสื่อสารผ่านโปรโตคอล HTTPS เพื่อปกป้องข้อมูลของคุณ</p>

        <h2>5. การติดต่อเรา</h2>
        <p>หากคุณมีคำถามเกี่ยวกับนโยบายความเป็นส่วนตัว สามารถติดต่อเราได้ที่ช่องทางโซเชียลมีเดียหรือ GitHub ของผู้พัฒนา <br> E-mail: fair2708@gmail.com</p>
      </section>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router';
</script>

<style scoped>
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
