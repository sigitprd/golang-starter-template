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

1. Pilih folder starter yang kamu inginkan
2. Salin `.env.example` menjadi `.env`, lalu sesuaikan
3. Jalankan database (PostgreSQL)
4. Jalankan aplikasi:

```bash
go run ./cmd/server/main.go
```

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

---