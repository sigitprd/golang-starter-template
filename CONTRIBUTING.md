# Contributing to Golang Starter Templates

Terima kasih atas minat Anda untuk berkontribusi! 🎉

## Cara Berkontribusi

### 1. Fork & Clone Repository
```bash
git clone https://github.com/your-username/golang-starter-template.git
cd golang-starter-template
```

### 2. Buat Branch Baru
```bash
git checkout -b feature/nama-fitur-anda
```

### 3. Setup Development Environment

#### Jalankan PostgreSQL dengan Docker
```bash
docker-compose up -d
```

#### Install Dependencies (untuk starter yang akan diubah)
```bash
cd echo-jwt-starter  # atau starter lainnya
go mod download
```

#### Copy Environment File
```bash
cp .env.example .env
```

### 4. Lakukan Perubahan

- Pastikan kode mengikuti style guide Go
- Tambahkan test jika menambahkan fitur baru
- Update dokumentasi jika diperlukan

### 5. Test Perubahan Anda

```bash
# Run tests
go test ./...

# Build untuk memastikan tidak ada error
go build ./cmd/server/main.go

# Format code
go fmt ./...
```

### 6. Commit & Push

```bash
git add .
git commit -m "feat: deskripsi singkat perubahan"
git push origin feature/nama-fitur-anda
```

### 7. Buat Pull Request

Buka Pull Request di GitHub dengan deskripsi yang jelas tentang perubahan yang Anda buat.

## Guidelines

### Code Style
- Gunakan `go fmt` untuk format code
- Follow [Effective Go](https://golang.org/doc/effective_go) guidelines
- Berikan nama variable dan function yang deskriptif

### Commit Messages
Gunakan conventional commits:
- `feat:` untuk fitur baru
- `fix:` untuk bug fix
- `docs:` untuk perubahan dokumentasi
- `refactor:` untuk refactoring code
- `test:` untuk menambahkan test
- `chore:` untuk maintenance tasks

### Pull Request
- Berikan judul yang jelas dan deskriptif
- Jelaskan apa yang berubah dan kenapa
- Link ke issue terkait jika ada
- Pastikan semua test passing

## Struktur Project

Setiap starter memiliki struktur yang konsisten:
```
├── cmd/              # Entry point aplikasi
├── config/           # Konfigurasi
├── internal/         # Business logic (handler, service, repository)
├── middleware/       # Middleware (hanya di versi JWT)
├── migrations/       # Database migrations
├── pkg/              # Shared utilities
└── .env.example      # Template environment variables
```

## Pertanyaan?

Jika ada pertanyaan, silakan buka issue di GitHub.

Terima kasih atas kontribusi Anda! 🙏
