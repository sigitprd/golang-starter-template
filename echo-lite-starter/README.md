# Echo Lite Starter

Starter project Golang + Echo + PostgreSQL + Zerolog + .env + Migrate (tanpa JWT authentication)

## 🎯 Fitur

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
cd echo-lite-starter
```

**Option B: PostgreSQL Manual**
```bash
# Buat database
createdb echo_lite_starter
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

Starter ini adalah template dasar untuk REST API. Anda bisa menambahkan endpoint sesuai kebutuhan di `internal/routes/routes.go`.

### Example Endpoints (perlu diimplementasikan)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/v1/health` | Health check |
| GET | `/api/v1/items` | List items |
| POST | `/api/v1/items` | Create item |
| GET | `/api/v1/items/:id` | Get item by ID |
| PUT | `/api/v1/items/:id` | Update item |
| DELETE | `/api/v1/items/:id` | Delete item |

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
├── migrations/               # Database migrations
├── pkg/
│   ├── config/               # Config utilities
│   ├── db/                   # Database utilities
│   ├── errmsg/               # Error messages
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

## 🚀 Next Steps

Setelah setup, Anda bisa:

1. Menambahkan entity baru di `internal/entity/`
2. Membuat repository untuk data access di `internal/repository/`
3. Implementasi business logic di `internal/service/`
4. Tambahkan HTTP handlers di `internal/handler/`
5. Daftarkan routes di `internal/routes/routes.go`
6. Buat database migrations di `migrations/`

## 🔒 Security Notes

- Ganti `X_API_KEY` di .env untuk keamanan API
- Gunakan HTTPS di production
- Jangan commit file `.env` ke repository
- Set `APP_ENV=production` di production environment
- Jika butuh authentication, pertimbangkan gunakan `echo-jwt-starter`

## 🐛 Troubleshooting

### Migration Error

```bash
# Pastikan database sudah dibuat dan credentials benar
# Cek koneksi database
psql -h localhost -U postgres -d echo_lite_starter

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
