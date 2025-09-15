# Go Backend Lite (Free Version) 🚀

Simple Go backend boilerplate dengan **Gin**, **PostgreSQL**, JWT authentication, dan protected routes.

![demo-gif](https://via.placeholder.com/600x200?text=Demo+GIF)  

---

## **Features**
- ✅ Gin HTTP server  
- ✅ PostgreSQL database integration  
- ✅ User registration & login  
- ✅ JWT authentication  
- ✅ Protected route example  

---

## **Requirements**
- Go 1.21+  
- PostgreSQL 17+  
- Postman (optional, untuk test API)  

---

## **Quick Start**

1. Clone repo:  
```bash
git clone https://github.com/username/go-backend-boilerplate-lite.git
cd go-backend-boilerplate-lite
```
2. Copy .env.example → .env dan update values:
```bash
cp .env.example .env
```
3. Run server:
```bash
go run cmd/server/main.go
```
4. Test API endpoints:
- /auth/register → POST user register
- /auth/login → POST login, dapat JWT token
- /users/ → GET protected route, pakai Authorization: Bearer <token>

---

## **Demo**
- Health check: /health →
```json
{
  "status": "ok",
  "message": "Go Backend Lite is running 🚀"
}
```

---

## **Version Pro (Coming Soon)**
- Multi-role access / admin dashboard
- Rate limiting / analytics / logging
- Extra endpoints ready-to-deploy
- Get it on Gumroad → Gumroad link placeholder

---

## **Contribute / Feedback**
- Feel free to open an issue / pull request
- Your feedback is welcome 😎

---

## **License**
MIT License © 2025 [alifslhdn]
