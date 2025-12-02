# Echo JWT Starter

Starter project Golang + Echo + JWT + PostgreSQL + Zerolog + .env + Migrate

## 🎯 Fitur

- ✅ Register, Login, Refresh Token
- ✅ JWT Authentication Middleware
- ✅ PostgreSQL tanpa ORM (native SQL)
- ✅ Error handling terpusat
- ✅ Logging dengan zerolog
- ✅ Migrasi database dengan golang-migrate
- ✅ Input validation dengan go-playground/validator
- ✅ CORS, Gzip, Rate Limiting middleware
- ✅ Clean Architecture pattern

## 📋 Prerequisites

- Go 1.20 atau lebih tinggi
- PostgreSQL 14+ (atau gunakan Docker Compose)
- golang-migrate CLI (untuk manual migration)

## 🚀 Quick Start

### 1. Setup Environment

```bash
# Copy .env.example ke .env
cp .env.example .env

# Edit .env sesuai kebutuhan, minimal ubah:
# - DB_NAME (nama database)
# - JWT_SECRET (gunakan secret key yang kuat)
# - X_API_KEY (API key untuk keamanan tambahan)
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Setup Database

**Option A: Gunakan Docker Compose** (dari root project)
```bash
cd ..
docker-compose up -d
cd echo-jwt-starter
```

**Option B: PostgreSQL Manual**
```bash
# Buat database
createdb echo_jwt_starter
```

### 4. Jalankan Migrasi

```bash
make migrate-up
```

### 5. Jalankan Server

```bash
make run
```

Server akan berjalan di `http://localhost:3000` (atau sesuai APP_PORT di .env)

## 📖 API Endpoints

### Authentication

| Method | Endpoint | Deskripsi | Auth Required |
|--------|----------|-----------|---------------|
| POST | `/api/v1/auth/register` | Registrasi user baru | ❌ |
| POST | `/api/v1/auth/login` | Login dan dapatkan token | ❌ |
| POST | `/api/v1/auth/refresh` | Refresh access token | ✅ |
| GET | `/api/v1/auth/me` | Get user profile | ✅ |

### Example: Register

```bash
curl -X POST http://localhost:3000/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "name": "John Doe"
  }'
```

### Example: Login

```bash
curl -X POST http://localhost:3000/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

Response akan berisi `access_token` dan `refresh_token`.

### Example: Access Protected Endpoint

```bash
curl -X GET http://localhost:3000/api/v1/auth/me \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

## 🛠️ Available Commands

```bash
# Jalankan server
make run

# Build binary
make build

# Build untuk Linux
make build-linux

# Build untuk macOS
make build-mac

# Build untuk Windows
make build-windows

# Buat migrasi baru
make migrate-new

# Jalankan migrasi up
make migrate-up

# Rollback migrasi
make migrate-down

# Drop semua migrasi
make migrate-drop
```

## 📁 Struktur Project

```
.
├── cmd/
│   └── server/
│       └── main.go           # Entry point
├── config/
│   └── config.go             # Configuration loader
├── internal/
│   ├── dto/                  # Data Transfer Objects
│   ├── entity/               # Domain entities
│   ├── handler/              # HTTP handlers
│   ├── repository/           # Data access layer
│   ├── routes/               # Route definitions
│   └── service/              # Business logic
├── middleware/
│   └── auth_bearer.go        # JWT authentication middleware
├── migrations/               # Database migrations
├── pkg/
│   ├── config/               # Config utilities
│   ├── db/                   # Database utilities
│   ├── errmsg/               # Error messages
│   ├── jwthandler/           # JWT utilities
│   ├── logging/              # Logger setup
│   ├── response/             # Response formatter
│   ├── utils/                # General utilities
│   └── validator/            # Custom validators
├── .env.example              # Environment template
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## 🔒 Security Notes

- Ganti `JWT_SECRET` di .env dengan key yang kuat dan unik
- Ganti `X_API_KEY` di .env untuk keamanan API
- Gunakan HTTPS di production
- Jangan commit file `.env` ke repository
- Set `APP_ENV=production` di production environment

## 🐛 Troubleshooting

### Migration Error

```bash
# Pastikan database sudah dibuat dan credentials benar
# Cek koneksi database
psql -h localhost -U postgres -d echo_jwt_starter

# Reset migrasi jika perlu
make migrate-drop
make migrate-up
```

### Port Already in Use

Jika port 3000 sudah digunakan, ubah `APP_PORT` di file `.env`.

## 📚 Learn More

- [Echo Framework Documentation](https://echo.labstack.com/)
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [Zerolog](https://github.com/rs/zerolog)

## 🤝 Contributing

Lihat [CONTRIBUTING.md](../CONTRIBUTING.md) di root repository untuk panduan kontribusi.
