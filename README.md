# 🚀 PloyKong (ปล่อยของ)
> **The Ultimate No-Code Profile Builder for Gen-Z Professionals**
> *"มีดีต้องปล่อยของ สร้างพอร์ตสวยง่ายๆ ในไม่กี่คลิก"*

## 📁 Project Structure

```
ploykong/
├── backend/          Go Fiber REST API (auth, portfolios, analytics, AI)
├── frontend/         Vue.js 3 + Pinia Web Builder (drag-drop, WYSIWYG)
├── mobile/           Flutter Admin Dashboard (charts, push notifications)
├── database/         MySQL 8 schema, views, stored procedures, seed data
├── nginx/            Wildcard SSL reverse proxy config (*.ploykong.com)
├── docs/             Full HTML deployment guide (Phase 0–9)
├── .github/          GitHub Actions CI/CD pipeline
└── docker-compose.yml  API + MySQL + Redis + Nginx + Certbot
```

## 🛠️ Tech Stack
- **Backend**: Go 1.22 + Fiber Framework — JWT auth, REST API, OpenAI integration
- **Frontend**: Vue.js 3 + Pinia + TypeScript — Drag-drop builder, live preview
- **Mobile**: Flutter 3.16 + Riverpod — Analytics charts, push notifications (iOS/Android)
- **Database**: MySQL 8 + JSON columns — Flexible block storage
- **Cache**: Redis — Session, rate limiting
- **Infra**: Docker Compose, Nginx wildcard SSL, GitHub Actions CI/CD

## 🎯 Key Features
- 🎨 Drag & drop portfolio builder (WYSIWYG, one-click publish)
- 🤖 AI Content Writer, Resume Scorer, Interview Coach (RAG chatbot)
- 📊 Visitor analytics (country, source, daily views, hire clicks)
- 🔔 Real-time push notifications (PDF downloads, Hire Me clicks)
- 🔒 Password-protected sections, self-destructing links

## 🚀 Quick Start
```bash
# 1. Infrastructure
cp backend/.env.example backend/.env && docker compose up -d mysql redis

# 2. Backend (http://localhost:3000)
cd backend && go mod download && go run cmd/main.go

# 3. Frontend (http://localhost:5173)
cd frontend && npm install && npm run dev

# 4. Mobile
cd mobile && flutter pub get && flutter run
```

## 💰 Cost: ~$25/month (VPS $24 + domain $1.25, SSL & Vercel free)

---
> KMUTT Information Technology · Built with 💜 Go, Vue.js 3, Flutter
