# Blog & Forum API Documentation

Sistem blogging dan forum dengan fitur:

- **Blog Postingan (Post)**
- **Komentar & Balasan** (dengan parent_id untuk nested comment)
- **Like** untuk Post dan Komentar
- **Forum** (forum, forum comment, dan comment ordering)
- Like pada Forum dan Forum Comments

---

## Endpoints

### Blog Posts

- **GET /posts**: List posts (support pagination & search)  
- **POST /posts**: Buat posting baru  
- **GET /posts/:postId**: Detail posting  
- **PUT /posts/:postId**: Update posting  
- **DELETE /posts/:postId**: Hapus posting  
- **GET /posts//top-contributors**: Top Contributors
- **GET /posts/popular-topic**: Popular topic by tag


### Blog Comments

- **GET /posts/:postId/comments**: List komentar post  
- **POST /posts/:postId/comments**: Tambah komentar pada post  
- **PUT /comments/:commentId**: Update komentar  
- **DELETE /comments/:commentId**: Hapus komentar  

### Blog Likes

- **POST /posts/:postId/likes**: Like posting  
- **DELETE /posts/:postId/likes**: Unlike posting  
- **POST /comments/:commentId/likes**: Like komentar  
- **DELETE /comments/:commentId/likes**: Unlike komentar  

---

### Query Parameters (Support)

- Pagination: `GET /posts?page=1&limit=10`  
- Search: `GET /posts?search=golang`

---

## Forum Endpoints

### Forum

- **GET /forums**: List forum  
- **POST /forums**: Buat forum baru  
- **GET /forums/:forumId**: Detail forum  
- **PUT /forums/:forumId**: Update forum  
- **DELETE /forums/:forumId**: Hapus forum  
- **GET /forums/top-contributors**: Top Contributors
- **GET /forums/popular-topic**: Popular topic by tag
- **GET /forums/stats**: Popular topic by tag


### Forum Comments

- **GET /forums/:forumId/comments**: List komentar forum  
- **POST /forums/:forumId/comments**: Tambah komentar pada forum  
- **PUT /forum-comments/:commentId**: Update komentar forum  
- **DELETE /forum-comments/:commentId**: Hapus komentar forum  

### Comment Ordering (Up/Down)

- **POST /forum-comments/:commentId/move-up**: Pindah komentar ke atas  
- **POST /forum-comments/:commentId/move-down**: Pindah komentar ke bawah  

### Forum Likes

- **POST /forums/:forumId/likes**: Like forum  
- **DELETE /forums/:forumId/likes**: Unlike forum  
- **POST /forum-comments/:commentId/likes**: Like komentar forum  
- **DELETE /forum-comments/:commentId/likes**: Unlike komentar forum  
