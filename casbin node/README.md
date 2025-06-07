# Casbin API with Node.js and MySQL

API authorization menggunakan Casbin dengan Node.js, Express, dan MySQL untuk mengelola permissions dan roles dengan model RBAC (Role-Based Access Control).

## 📋 Fitur

- ✅ **RBAC Model** - Role-Based Access Control
- ✅ **MySQL Integration** - Custom adapter untuk menyimpan policies di database
- ✅ **REST API** - Endpoints lengkap untuk mengelola policies dan roles
- ✅ **Real-time Sync** - Sinkronisasi otomatis antara aplikasi dan database
- ✅ **Flexible Permissions** - Support untuk complex authorization rules

## 🚀 Quick Start

### Prerequisites

- Node.js (v14 atau lebih baru)
- MySQL (v5.7 atau lebih baru)
- npm atau yarn

### Installation

1. **Clone atau copy project files**

2. **Install dependencies:**
```bash
npm install
```

3. **Setup environment:**
```bash
cp .env.example .env
```

Edit file `.env`:
```env
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=casbin_db
PORT=3000
```

4. **Buat database MySQL:**
```sql
CREATE DATABASE casbin_db;
```

5. **Jalankan aplikasi:**
```bash
# Development mode
npm run dev

# Production mode  
npm start
```

Server akan berjalan di `http://localhost:3000`

## 📚 API Endpoints

### 🔐 Policy Management

#### Tambah Policy
```http
POST /api/policies
Content-Type: application/json

{
  "subject": "admin",
  "object": "users", 
  "action": "read"
}
```

#### Hapus Policy
```http
DELETE /api/policies
Content-Type: application/json

{
  "subject": "admin",
  "object": "users",
  "action": "read"
}
```

#### Lihat Semua Policies
```http
GET /api/policies
```

### 👥 Role Management

#### Tambah Role untuk User
```http
POST /api/roles
Content-Type: application/json

{
  "user": "john",
  "role": "admin"
}
```

#### Hapus Role dari User
```http
DELETE /api/roles
Content-Type: application/json

{
  "user": "john", 
  "role": "admin"
}
```

#### Get Roles untuk User
```http
GET /api/roles/john
```

#### Get Users dengan Role Tertentu
```http
GET /api/users/admin
```

### ✅ Permission Check

#### Check Permission (Enforce)
```http
POST /api/check
Content-Type: application/json

{
  "subject": "john",
  "object": "users",
  "action": "read"
}
```

### 🏥 Health Check
```http
GET /health
```

## 💾 Database Schema

Aplikasi otomatis membuat tabel `casbin_rule`:

```sql
CREATE TABLE casbin_rule (
  id INT AUTO_INCREMENT PRIMARY KEY,
  ptype VARCHAR(255) NOT NULL,    -- Policy type (p, g)
  v0 VARCHAR(255),                -- Parameter 1
  v1 VARCHAR(255),                -- Parameter 2
  v2 VARCHAR(255),                -- Parameter 3
  v3 VARCHAR(255),                -- Parameter 4
  v4 VARCHAR(255),                -- Parameter 5
  v5 VARCHAR(255),                -- Parameter 6
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### Penjelasan Kolom v0-v5

#### Untuk Policy Rules (ptype = 'p'):
| Kolom | Isi | Contoh | Keterangan |
|-------|-----|--------|------------|
| **v0** | Subject | `alice`, `admin` | User atau role yang melakukan aksi |
| **v1** | Object | `users`, `posts` | Resource yang diakses |
| **v2** | Action | `read`, `write` | Aksi yang dilakukan |
| **v3** | Effect | `allow`, `deny` | Efek policy (opsional) |
| **v4** | Domain | `tenant1` | Domain untuk multi-tenancy (opsional) |
| **v5** | Condition | `owner` | Kondisi tambahan (opsional) |

#### Untuk Role Rules (ptype = 'g'):
| Kolom | Isi | Contoh | Keterangan |
|-------|-----|--------|------------|
| **v0** | User | `alice`, `bob` | User yang mendapat role |
| **v1** | Role | `admin`, `editor` | Role yang diberikan |
| **v2** | Domain | `tenant1` | Domain (opsional) |

## 🎯 Contoh Penggunaan

### Scenario: Blog Management System

1. **Setup Roles dan Policies:**
```bash
# Tambah policy untuk admin
curl -X POST http://localhost:3000/api/policies \
  -H "Content-Type: application/json" \
  -d '{"subject": "admin", "object": "posts", "action": "write"}'

# Tambah policy untuk editor  
curl -X POST http://localhost:3000/api/policies \
  -H "Content-Type: application/json" \
  -d '{"subject": "editor", "object": "posts", "action": "read"}'

# Assign role admin ke user john
curl -X POST http://localhost:3000/api/roles \
  -H "Content-Type: application/json" \
  -d '{"user": "john", "role": "admin"}'
```

2. **Check Permissions:**
```bash
# Check apakah john bisa write posts
curl -X POST http://localhost:3000/api/check \
  -H "Content-Type: application/json" \
  -d '{"subject": "john", "object": "posts", "action": "write"}'

# Response: {"success": true, "allowed": true, "message": "Akses diizinkan"}
```

3. **View Current Setup:**
```bash
# Lihat semua policies dan roles
curl http://localhost:3000/api/policies

# Lihat roles untuk user john
curl http://localhost:3000/api/roles/john
```

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Client App    │────│   Express API   │────│   MySQL DB      │
│                 │    │                 │    │                 │
│ - Web App       │    │ - REST Routes   │    │ - casbin_rule   │
│ - Mobile App    │    │ - Casbin Logic  │    │   table         │
│ - Other Service │    │ - Custom Adapter│    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

### Components:

- **Express Server**: REST API endpoints
- **Casbin Enforcer**: Authorization engine dengan RBAC model
- **Custom MySQL Adapter**: Menghubungkan Casbin dengan MySQL
- **MySQL Database**: Menyimpan policies dan roles secara persistent

## 🔧 Konfigurasi Model

Model RBAC yang digunakan:

```ini
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
```

### Kustomisasi Model

Anda bisa mengubah model di `config/casbin.js` untuk kebutuhan yang lebih kompleks:

```javascript
// Contoh model ABAC (Attribute-Based Access Control)
const model = casbin.newModel(`
  [request_definition]
  r = sub, obj, act, env

  [policy_definition]
  p = sub, obj, act, env

  [matchers]
  m = r.sub == p.sub && r.obj == p.obj && r.act == p.act && r.env == p.env
`);
```

## 🧪 Testing

### Manual Testing

```bash
# Test policy creation
curl -X POST http://localhost:3000/api/policies \
  -H "Content-Type: application/json" \
  -d '{"subject": "testuser", "object": "testresource", "action": "read"}'

# Test permission check  
curl -X POST http://localhost:3000/api/check \
  -H "Content-Type: application/json" \
  -d '{"subject": "testuser", "object": "testresource", "action": "read"}'
```

### Database Verification

```sql
-- Lihat semua policies
SELECT ptype, v0 as subject, v1 as object, v2 as action 
FROM casbin_rule 
WHERE ptype = 'p';

-- Lihat semua role assignments
SELECT ptype, v0 as user, v1 as role 
FROM casbin_rule 
WHERE ptype = 'g';
```

## 🔍 Troubleshooting

### Common Issues:

1. **Database Connection Error:**
   - Pastikan MySQL service berjalan
   - Check kredensial di file `.env`
   - Pastikan database sudah dibuat

2. **Policy Tidak Tersimpan:**
   - Check response dari API calls
   - Verify dengan query database langsung
   - Pastikan format request sesuai

3. **Permission Check Selalu False:**
   - Pastikan policy sudah ditambahkan
   - Check apakah user memiliki role yang sesuai
   - Verify model matcher logic

### Debug Mode:

Tambahkan logging untuk debugging:

```javascript
// Di server.js, tambah sebelum enforce check
console.log('Checking permission:', { subject, object, action });
console.log('User roles:', await enforcer.getRolesForUser(subject));
console.log('All policies:', await enforcer.getPolicy());
```

