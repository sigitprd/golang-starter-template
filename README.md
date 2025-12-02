# 🚀 Golang All Starter Template

Kumpulan starter project backend dengan Golang untuk berbagai kebutuhan REST API. Tersedia dalam varian Echo dan Fiber, dengan dan tanpa autentikasi JWT.

## 📁 Daftar Starter

| Folder              | Deskripsi                                         |
|---------------------|--------------------------------------------------|
| `echo-lite-starter` | Starter ringan menggunakan Echo (tanpa auth)     |
| `echo-jwt-starter`  | Starter Echo dengan sistem autentikasi JWT       |
| `fiber-lite-starter`| Starter ringan menggunakan Fiber (tanpa auth)    |
| `fiber-jwt-starter` | Starter Fiber dengan sistem autentikasi JWT      |

---

## 🧰 Fitur Umum

- Struktur modular dan clean
- Middleware (CORS, logging, recover, gzip)
- Custom validator dengan `go-playground/validator`
- Logging menggunakan Zerolog
- PostgreSQL tanpa ORM
- Konfigurasi menggunakan `.env`
- Migrasi database (opsional)

---

## ⚙️ Cara Menggunakan

### Quick Start

1. **Clone repository ini**
```bash
git clone https://github.com/sigitprd/golang-starter-template.git
cd golang-starter-template
```

2. **Jalankan PostgreSQL dengan Docker Compose** (opsional)
```bash
docker-compose up -d
```
Ini akan menjalankan PostgreSQL di port 5432 dan Adminer (database UI) di port 8080.

3. **Pilih starter yang diinginkan**
```bash
cd echo-jwt-starter  # atau echo-lite-starter, fiber-jwt-starter, fiber-lite-starter
```

4. **Setup environment**
```bash
cp .env.example .env
# Edit .env sesuai kebutuhan (database credentials, JWT secret, dll)
```

5. **Install dependencies**
```bash
go mod download
```

6. **Jalankan migrasi database**
```bash
make migrate-up
```

7. **Jalankan aplikasi**
```bash
make run
# atau: go run ./cmd/server/main.go
```

### Akses Adminer (Database UI)
Jika menggunakan Docker Compose, buka browser: http://localhost:8080
- **System**: PostgreSQL
- **Server**: postgres
- **Username**: postgres
- **Password**: postgres
- **Database**: sesuai DB_NAME di .env Anda

---

## 🏗️ Struktur Umum Starter
```gotemplate
.
├── .github/          # Workflow CI/CD
├── cmd/              # Entry point aplikasi
├── config/           # Konfigurasi aplikasi
├── internal/         # Handler, service, repo, dsb.
├── logs/             # File log aplikasi
├── middleware/       # Middleware (CORS, logging, dsb.)
├── migrations/       # Migrasi dan seed SQL
├── pkg/              # Helper, util, logger, response
├── storage/          # File statis (jika ada)
├── go.mod / go.sum   # Dependency management
├── .env              # File konfigurasi environment
├── .env.deploy       # File konfigurasi untuk deploy
└── Makefile          # Skrip build dan run
```

---

## 📌 Tech Stack
- **Golang**: Bahasa pemrograman utama
- **Echo/Fiber**: Web framework
- **JWT**: Autentikasi token
- **PostgreSQL**: Database relasional
- **Zerolog**: Logging
- **Validator**: Validasi input
- **Migrate**: Migrasi database
- **CORS**: Cross-Origin Resource Sharing
- **Gzip**: Kompresi response
- **Middleware**: Error handling, logging, dsb.
- **.env**: Konfigurasi environment
- **Makefile**: Skrip build dan run

---

## 🧱 Cocok Digunakan Untuk
- Membuat REST API backend sederhana hingga menengah
- Belajar best practice structuring project Golang
- Membuat microservice ringan dengan kebutuhan JWT
- Prototype atau MVP development
- Base template untuk production-ready application

---

## 🎯 Pilih Starter yang Tepat

| Kebutuhan | Starter yang Cocok |
|-----------|-------------------|
| Simple REST API tanpa auth | `echo-lite-starter` atau `fiber-lite-starter` |
| REST API dengan JWT auth | `echo-jwt-starter` atau `fiber-jwt-starter` |
| Prefer Echo framework | `echo-lite-starter` atau `echo-jwt-starter` |
| Prefer Fiber framework (lebih cepat) | `fiber-lite-starter` atau `fiber-jwt-starter` |

---

## 🤝 Contributing

Kami menyambut kontribusi! Silakan baca [CONTRIBUTING.md](CONTRIBUTING.md) untuk panduan lengkap.

---

## 📄 License

Project ini menggunakan [MIT License](LICENSE).

---