# Go Backend Cheatsheet

Cheatsheet ini buat kamu yang terbiasa dengan Node.js/npm dan sedang belajar backend Go.

## Analogi Node.js dan Go

| Node.js | Go |
| --- | --- |
| `npm init` | `go mod init nama-module` |
| `npm install express` | `go get github.com/gin-gonic/gin` |
| `npm install` | `go mod download` atau `go mod tidy` |
| `package.json` | `go.mod` |
| `package-lock.json` | `go.sum` |
| `node_modules` | module cache Go, tidak disimpan di project |
| `npm uninstall package` | hapus import lalu jalankan `go mod tidy` |
| `npm run dev` | biasanya `go run cmd/main.go` |

## Command Go Yang Sering Dipakai

### Membuat Project Go Baru

```bash
go mod init nama-project
```

### Menambahkan Library

```bash
go get nama/package
```

Contoh:

```bash
go get github.com/gin-gonic/gin
```

### Menambahkan Library Versi Terbaru

```bash
go get github.com/gin-gonic/gin@latest
```

### Menambahkan Library Versi Tertentu

```bash
go get github.com/gin-gonic/gin@v1.12.0
```

### Merapikan Dependency

Jalankan ini setelah menambah atau menghapus library:

```bash
go mod tidy
```

### Download Semua Dependency

```bash
go mod download
```

### Menjalankan Backend

Untuk project ini:

```bash
go run cmd/main.go
```

### Build Project

```bash
go build ./...
```

### Menjalankan Test

```bash
go test ./...
```

## `go get` vs `go install`

Gunakan `go get` kalau library dipakai di kode project.

```bash
go get github.com/gin-gonic/gin
```

Gunakan `go install` kalau ingin install command line tool ke komputer.

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Patokan gampang:

```text
Mau pakai library di kode?  go get
Mau install CLI/tool?       go install
Mau rapikan dependency?     go mod tidy
Mau jalanin backend?        go run cmd/main.go
```

## File Dependency Di Go

### `go.mod`

File utama dependency project Go.

Contoh:

```go
module go-backend

go 1.25.4

require (
    github.com/gin-gonic/gin v1.12.0
    github.com/go-sql-driver/mysql v1.9.3
)
```

### `go.sum`

File checksum dependency. Anggap mirip lock file.

Biasanya kamu tidak perlu edit manual file ini.

## Cara Tambah Stack Di Go

### Tambah Gin

Gin dipakai untuk membuat REST API.

```bash
go get github.com/gin-gonic/gin
go mod tidy
```

Import di kode:

```go
import "github.com/gin-gonic/gin"
```

### Tambah CORS Untuk Gin

CORS dipakai supaya frontend bisa akses backend.

```bash
go get github.com/gin-contrib/cors
go mod tidy
```

Import di kode:

```go
import "github.com/gin-contrib/cors"
```

### Tambah MySQL Driver

Driver MySQL dipakai agar `database/sql` bisa konek ke MySQL.

```bash
go get github.com/go-sql-driver/mysql
go mod tidy
```

Import di kode:

```go
import _ "github.com/go-sql-driver/mysql"
```

Tanda `_` artinya package di-load untuk efek samping. Dalam kasus ini, efeknya adalah mendaftarkan driver MySQL ke `database/sql`.

### Tambah JWT

JWT dipakai untuk login token/authentication.

```bash
go get github.com/golang-jwt/jwt/v5
go mod tidy
```

Import di kode:

```go
import "github.com/golang-jwt/jwt/v5"
```

### Tambah Goose Migration

Goose dipakai untuk migration database.

Sebagai library project:

```bash
go get github.com/pressly/goose/v3
go mod tidy
```

Sebagai CLI tool:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

## Stack Project Ini

Project ini memakai stack berikut:

| Bagian | Stack |
| --- | --- |
| Backend language | Go |
| HTTP framework | Gin |
| Database | MySQL |
| DB access | `database/sql` |
| MySQL driver | `github.com/go-sql-driver/mysql` |
| Authentication | JWT |
| Migration | Goose |
| Frontend | React |
| Frontend build tool | Vite |
| HTTP client frontend | Axios |
| Frontend routing | React Router DOM |

## Alur Backend Project Ini

```text
frontend React
    |
routes
    |
handler
    |
service
    |
repository
    |
MySQL
```

Contoh alur login:

```text
POST /login
    |
handler.Login
    |
service.Login
    |
repository mencari user berdasarkan email
    |
validasi password
    |
generate JWT token
    |
response ke frontend
```

## Pola Aman Saat Tambah Library

Gunakan urutan ini:

```bash
go get nama-library
go mod tidy
go run cmd/main.go
```

Kalau ada error, baca pesan error-nya dari atas ke bawah. Biasanya Go cukup jelas memberi tahu package mana yang kurang atau import mana yang tidak dipakai.

## Tiga Command Yang Wajib Hafal

Kalau masih awal belajar Go backend, cukup hafal ini dulu:

```bash
go get nama-library
go mod tidy
go run cmd/main.go
```

Itu sudah cukup untuk mulai menambah stack, merapikan dependency, dan menjalankan backend.
