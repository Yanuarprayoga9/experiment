-- =====================================================
-- SAMPLE DATA - 10 POSTS WITH TAGS AND LIKES (FIXED IDs)
-- =====================================================

-- 1. Pengguna (Users) - 10 users
INSERT INTO Pengguna (Id, Nip, Nama, Email, KataSandi, Peran, Jabatan, UnitKerja) VALUES 
('00000000-0000-0000-0000-000000000011', 'NIP011', 'Ahmad Rizki', 'ahmad.rizki@company.com', 'password123', 'SME', 'Senior Developer', 'IT Development'),
('00000000-0000-0000-0000-000000000012', 'NIP012', 'Sari Dewi', 'sari.dewi@company.com', 'password123', 'Reviewer', 'Tech Lead', 'IT Development'),
('00000000-0000-0000-0000-000000000013', 'NIP013', 'Budi Santoso', 'budi.santoso@company.com', 'password123', 'SME', 'Database Administrator', 'IT Infrastructure'),
('00000000-0000-0000-0000-000000000014', 'NIP014', 'Maya Putri', 'maya.putri@company.com', 'password123', 'SME', 'Frontend Developer', 'IT Development'),
('00000000-0000-0000-0000-000000000015', 'NIP015', 'Andi Pratama', 'andi.pratama@company.com', 'password123', 'Admin', 'Project Manager', 'IT Development'),
('00000000-0000-0000-0000-000000000016', 'NIP016', 'Rina Sari', 'rina.sari@company.com', 'password123', 'SME', 'DevOps Engineer', 'IT Infrastructure'),
('00000000-0000-0000-0000-000000000017', 'NIP017', 'Doni Kurniawan', 'doni.kurniawan@company.com', 'password123', 'SME', 'Backend Developer', 'IT Development'),
('00000000-0000-0000-0000-000000000018', 'NIP018', 'Lisa Anggraeni', 'lisa.anggraeni@company.com', 'password123', 'Reviewer', 'UI/UX Designer', 'IT Development'),
('00000000-0000-0000-0000-000000000019', 'NIP019', 'Fajar Wijaya', 'fajar.wijaya@company.com', 'password123', 'SME', 'Mobile Developer', 'IT Development'),
('00000000-0000-0000-0000-000000000020', 'NIP020', 'Nina Rahmawati', 'nina.rahmawati@company.com', 'password123', 'SuperUser', 'CTO', 'IT Management');

-- 2. PostinganBlog - 10 posts
INSERT INTO PostinganBlog (Id, Judul, Slug, Isi, PenulisId, StatusReview, Visibilitas, JumlahTayang, UrlThumbnail) VALUES 
('10000000-0000-0000-0000-000000000011', 'Panduan Lengkap Belajar Golang dari Nol', 'panduan-lengkap-belajar-golang-dari-nol', 'Artikel ini membahas panduan lengkap untuk mempelajari bahasa pemrograman Go (Golang) dari dasar hingga mahir. Termasuk setup environment, syntax dasar, dan best practices.', '00000000-0000-0000-0000-000000000011', 'diterbitkan', 'publik', 150, 'https://example.com/thumbnails/golang-guide.jpg'),

('10000000-0000-0000-0000-000000000012', 'Optimasi Database SQL Server untuk Performa Tinggi', 'optimasi-database-sql-server-untuk-performa-tinggi', 'Tips dan trik untuk mengoptimalkan performa database SQL Server, termasuk indexing, query optimization, dan maintenance strategies.', '00000000-0000-0000-0000-000000000013', 'diterbitkan', 'internal', 89, 'https://example.com/thumbnails/sql-optimization.jpg'),

('10000000-0000-0000-0000-000000000013', 'Membangun RESTful API dengan Node.js dan Express', 'membangun-restful-api-dengan-nodejs-dan-express', 'Tutorial step-by-step untuk membuat RESTful API menggunakan Node.js dan Express framework, termasuk authentication dan error handling.', '00000000-0000-0000-0000-000000000017', 'diterbitkan', 'publik', 234, 'https://example.com/thumbnails/nodejs-api.jpg'),

('10000000-0000-0000-0000-000000000014', 'Design System: Konsep dan Implementasi', 'design-system-konsep-dan-implementasi', 'Penjelasan tentang design system, komponen-komponennya, dan bagaimana mengimplementasikannya dalam proyek web dan mobile.', '00000000-0000-0000-0000-000000000018', 'menunggu_review', 'publik', 45, 'https://example.com/thumbnails/design-system.jpg'),

('10000000-0000-0000-0000-000000000015', 'Containerization dengan Docker: Best Practices', 'containerization-dengan-docker-best-practices', 'Panduan praktis menggunakan Docker untuk containerization aplikasi, termasuk Dockerfile optimization dan multi-stage builds.', '00000000-0000-0000-0000-000000000016', 'diterbitkan', 'internal', 67, 'https://example.com/thumbnails/docker-guide.jpg'),

('10000000-0000-0000-0000-000000000016', 'React Hooks: useState dan useEffect Deep Dive', 'react-hooks-usestate-dan-useeffect-deep-dive', 'Pembahasan mendalam tentang React Hooks, fokus pada useState dan useEffect, termasuk patterns dan common pitfalls.', '00000000-0000-0000-0000-000000000014', 'diterbitkan', 'publik', 178, 'https://example.com/thumbnails/react-hooks.jpg'),

('10000000-0000-0000-0000-000000000017', 'Flutter vs React Native: Perbandingan Lengkap', 'flutter-vs-react-native-perbandingan-lengkap', 'Analisis komprehensif perbandingan antara Flutter dan React Native untuk pengembangan aplikasi mobile cross-platform.', '00000000-0000-0000-0000-000000000019', 'diterbitkan', 'publik', 312, 'https://example.com/thumbnails/flutter-vs-rn.jpg'),

('10000000-0000-0000-0000-000000000018', 'Microservices Architecture: Implementasi dan Challenges', 'microservices-architecture-implementasi-dan-challenges', 'Panduan implementasi arsitektur microservices, termasuk service communication, data management, dan monitoring.', '00000000-0000-0000-0000-000000000015', 'draf', 'terbatas', 12, NULL),

('10000000-0000-0000-0000-000000000019', 'CI/CD Pipeline dengan GitLab dan Jenkins', 'cicd-pipeline-dengan-gitlab-dan-jenkins', 'Tutorial setup CI/CD pipeline menggunakan GitLab CI dan Jenkins, termasuk automated testing dan deployment strategies.', '00000000-0000-0000-0000-000000000016', 'menunggu_review', 'internal', 28, 'https://example.com/thumbnails/cicd-pipeline.jpg'),

('10000000-0000-0000-0000-000000000020', 'Machine Learning dengan Python: Pemula hingga Mahir', 'machine-learning-dengan-python-pemula-hingga-mahir', 'Panduan komprehensif belajar machine learning dengan Python, dari konsep dasar hingga implementasi algoritma kompleks.', '00000000-0000-0000-0000-000000000012', 'ditolak', 'publik', 0, 'https://example.com/thumbnails/ml-python.jpg');

-- 3. Tag - 20 tags (increased from 15)
INSERT INTO Tag (Id, NamaTag) VALUES 
('20000000-0000-0000-0000-000000000021', 'Golang'),
('20000000-0000-0000-0000-000000000022', 'SQL Server'),
('20000000-0000-0000-0000-000000000023', 'Pemrograman'),
('20000000-0000-0000-0000-000000000024', 'Database'),
('20000000-0000-0000-0000-000000000025', 'Tutorial'),
('20000000-0000-0000-0000-000000000026', 'Node.js'),
('20000000-0000-0000-0000-000000000027', 'API'),
('20000000-0000-0000-0000-000000000028', 'Design'),
('20000000-0000-0000-0000-000000000029', 'UI/UX'),
('20000000-0000-0000-0000-000000000030', 'Docker'),
('20000000-0000-0000-0000-000000000031', 'DevOps'),
('20000000-0000-0000-0000-000000000032', 'React'),
('20000000-0000-0000-0000-000000000033', 'Frontend'),
('20000000-0000-0000-0000-000000000034', 'Mobile'),
('20000000-0000-0000-0000-000000000035', 'Flutter'),
('20000000-0000-0000-0000-000000000036', 'React Native'),
('20000000-0000-0000-0000-000000000037', 'Microservices'),
('20000000-0000-0000-0000-000000000038', 'CI/CD'),
('20000000-0000-0000-0000-000000000039', 'Machine Learning'),
('20000000-0000-0000-0000-000000000040', 'Python');

-- 4. PostinganTag - Relasi posts dengan tags
INSERT INTO PostinganTag (PostinganId, TagId) VALUES 
-- Post 1: Golang Guide
('10000000-0000-0000-0000-000000000011', '20000000-0000-0000-0000-000000000021'), -- Golang
('10000000-0000-0000-0000-000000000011', '20000000-0000-0000-0000-000000000023'), -- Pemrograman
('10000000-0000-0000-0000-000000000011', '20000000-0000-0000-0000-000000000025'), -- Tutorial

-- Post 2: SQL Server Optimization
('10000000-0000-0000-0000-000000000012', '20000000-0000-0000-0000-000000000022'), -- SQL Server
('10000000-0000-0000-0000-000000000012', '20000000-0000-0000-0000-000000000024'), -- Database
('10000000-0000-0000-0000-000000000012', '20000000-0000-0000-0000-000000000025'), -- Tutorial

-- Post 3: Node.js API
('10000000-0000-0000-0000-000000000013', '20000000-0000-0000-0000-000000000026'), -- Node.js
('10000000-0000-0000-0000-000000000013', '20000000-0000-0000-0000-000000000027'), -- API
('10000000-0000-0000-0000-000000000013', '20000000-0000-0000-0000-000000000023'), -- Pemrograman
('10000000-0000-0000-0000-000000000013', '20000000-0000-0000-0000-000000000025'), -- Tutorial

-- Post 4: Design System
('10000000-0000-0000-0000-000000000014', '20000000-0000-0000-0000-000000000028'), -- Design
('10000000-0000-0000-0000-000000000014', '20000000-0000-0000-0000-000000000029'), -- UI/UX
('10000000-0000-0000-0000-000000000014', '20000000-0000-0000-0000-000000000033'), -- Frontend

-- Post 5: Docker Best Practices
('10000000-0000-0000-0000-000000000015', '20000000-0000-0000-0000-000000000030'), -- Docker
('10000000-0000-0000-0000-000000000015', '20000000-0000-0000-0000-000000000031'), -- DevOps
('10000000-0000-0000-0000-000000000015', '20000000-0000-0000-0000-000000000025'), -- Tutorial

-- Post 6: React Hooks
('10000000-0000-0000-0000-000000000016', '20000000-0000-0000-0000-000000000032'), -- React
('10000000-0000-0000-0000-000000000016', '20000000-0000-0000-0000-000000000033'), -- Frontend
('10000000-0000-0000-0000-000000000016', '20000000-0000-0000-0000-000000000023'), -- Pemrograman

-- Post 7: Flutter vs React Native
('10000000-0000-0000-0000-000000000017', '20000000-0000-0000-0000-000000000034'), -- Mobile
('10000000-0000-0000-0000-000000000017', '20000000-0000-0000-0000-000000000035'), -- Flutter
('10000000-0000-0000-0000-000000000017', '20000000-0000-0000-0000-000000000036'), -- React Native

-- Post 8: Microservices
('10000000-0000-0000-0000-000000000018', '20000000-0000-0000-0000-000000000037'), -- Microservices
('10000000-0000-0000-0000-000000000018', '20000000-0000-0000-0000-000000000027'), -- API
('10000000-0000-0000-0000-000000000018', '20000000-0000-0000-0000-000000000023'), -- Pemrograman

-- Post 9: CI/CD Pipeline
('10000000-0000-0000-0000-000000000019', '20000000-0000-0000-0000-000000000038'), -- CI/CD
('10000000-0000-0000-0000-000000000019', '20000000-0000-0000-0000-000000000031'), -- DevOps
('10000000-0000-0000-0000-000000000019', '20000000-0000-0000-0000-000000000025'), -- Tutorial

-- Post 10: Machine Learning
('10000000-0000-0000-0000-000000000020', '20000000-0000-0000-0000-000000000039'), -- Machine Learning
('10000000-0000-0000-0000-000000000020', '20000000-0000-0000-0000-000000000040'), -- Python
('10000000-0000-0000-0000-000000000020', '20000000-0000-0000-0000-000000000023'); -- Pemrograman

-- 5. PostinganLike - Likes untuk posts (distribusi realistis)
INSERT INTO PostinganLike (Id, PostinganId, PenggunaId) VALUES 
-- Post 1 (Golang Guide) - 5 likes
(NEWID(), '10000000-0000-0000-0000-000000000011', '00000000-0000-0000-0000-000000000012'),
(NEWID(), '10000000-0000-0000-0000-000000000011', '00000000-0000-0000-0000-000000000013'),
(NEWID(), '10000000-0000-0000-0000-000000000011', '00000000-0000-0000-0000-000000000014'),
(NEWID(), '10000000-0000-0000-0000-000000000011', '00000000-0000-0000-0000-000000000017'),
(NEWID(), '10000000-0000-0000-0000-000000000011', '00000000-0000-0000-0000-000000000019'),

-- Post 2 (SQL Server) - 3 likes
(NEWID(), '10000000-0000-0000-0000-000000000012', '00000000-0000-0000-0000-000000000011'),
(NEWID(), '10000000-0000-0000-0000-000000000012', '00000000-0000-0000-0000-000000000015'),
(NEWID(), '10000000-0000-0000-0000-000000000012', '00000000-0000-0000-0000-000000000016'),

-- Post 3 (Node.js API) - 7 likes
(NEWID(), '10000000-0000-0000-0000-000000000013', '00000000-0000-0000-0000-000000000011'),
(NEWID(), '10000000-0000-0000-0000-000000000013', '00000000-0000-0000-0000-000000000012'),
(NEWID(), '10000000-0000-0000-0000-000000000013', '00000000-0000-0000-0000-000000000014'),
(NEWID(), '10000000-0000-0000-0000-000000000013', '00000000-0000-0000-0000-000000000015'),
(NEWID(), '10000000-0000-0000-0000-000000000013', '00000000-0000-0000-0000-000000000016'),
(NEWID(), '10000000-0000-0000-0000-000000000013', '00000000-0000-0000-0000-000000000018'),
(NEWID(), '10000000-0000-0000-0000-000000000013', '00000000-0000-0000-0000-000000000020'),

-- Post 4 (Design System) - 2 likes
(NEWID(), '10000000-0000-0000-0000-000000000014', '00000000-0000-0000-0000-000000000014'),
(NEWID(), '10000000-0000-0000-0000-000000000014', '00000000-0000-0000-0000-000000000019'),

-- Post 5 (Docker) - 4 likes
(NEWID(), '10000000-0000-0000-0000-000000000015', '00000000-0000-0000-0000-000000000011'),
(NEWID(), '10000000-0000-0000-0000-000000000015', '00000000-0000-0000-0000-000000000013'),
(NEWID(), '10000000-0000-0000-0000-000000000015', '00000000-0000-0000-0000-000000000017'),
(NEWID(), '10000000-0000-0000-0000-000000000015', '00000000-0000-0000-0000-000000000020'),

-- Post 6 (React Hooks) - 6 likes
(NEWID(), '10000000-0000-0000-0000-000000000016', '00000000-0000-0000-0000-000000000011'),
(NEWID(), '10000000-0000-0000-0000-000000000016', '00000000-0000-0000-0000-000000000012'),
(NEWID(), '10000000-0000-0000-0000-000000000016', '00000000-0000-0000-0000-000000000017'),
(NEWID(), '10000000-0000-0000-0000-000000000016', '00000000-0000-0000-0000-000000000018'),
(NEWID(), '10000000-0000-0000-0000-000000000016', '00000000-0000-0000-0000-000000000019'),
(NEWID(), '10000000-0000-0000-0000-000000000016', '00000000-0000-0000-0000-000000000020'),

-- Post 7 (Flutter vs RN) - 8 likes (paling populer)
(NEWID(), '10000000-0000-0000-0000-000000000017', '00000000-0000-0000-0000-000000000011'),
(NEWID(), '10000000-0000-0000-0000-000000000017', '00000000-0000-0000-0000-000000000012'),
(NEWID(), '10000000-0000-0000-0000-000000000017', '00000000-0000-0000-0000-000000000013'),
(NEWID(), '10000000-0000-0000-0000-000000000017', '00000000-0000-0000-0000-000000000014'),
(NEWID(), '10000000-0000-0000-0000-000000000017', '00000000-0000-0000-0000-000000000015'),
(NEWID(), '10000000-0000-0000-0000-000000000017', '00000000-0000-0000-0000-000000000016'),
(NEWID(), '10000000-0000-0000-0000-000000000017', '00000000-0000-0000-0000-000000000018'),
(NEWID(), '10000000-0000-0000-0000-000000000017', '00000000-0000-0000-0000-000000000020'),

-- Post 8 (Microservices) - 1 like (masih draft)
(NEWID(), '10000000-0000-0000-0000-000000000018', '00000000-0000-0000-0000-000000000020'),

-- Post 9 (CI/CD) - 2 likes
(NEWID(), '10000000-0000-0000-0000-000000000019', '00000000-0000-0000-0000-000000000015'),
(NEWID(), '10000000-0000-0000-0000-000000000019', '00000000-0000-0000-0000-000000000016');

-- Post 10 (ML Python) - 0 likes (ditolak)

-- 6. KomentarBlog - Sample comments
INSERT INTO KomentarBlog (Id, PostinganId, PenggunaId, KomentarIndukId, Konten) VALUES 
-- Comments untuk Post 1 (Golang)
(NEWID(), '10000000-0000-0000-0000-000000000011', '00000000-0000-0000-0000-000000000012', NULL, 'Artikel yang sangat membantu! Saya baru mulai belajar Golang dan ini sangat comprehensive.'),
(NEWID(), '10000000-0000-0000-0000-000000000011', '00000000-0000-0000-0000-000000000013', NULL, 'Bisa tolong tambahkan contoh untuk goroutines juga?'),

-- Comments untuk Post 3 (Node.js API)
(NEWID(), '10000000-0000-0000-0000-000000000013', '00000000-0000-0000-0000-000000000014', NULL, 'Implementasi middleware-nya sangat elegant. Thanks for sharing!'),
(NEWID(), '10000000-0000-0000-0000-000000000013', '00000000-0000-0000-0000-000000000018', NULL, 'Ada rekomendasi untuk testing framework yang cocok dengan setup ini?'),

-- Comments untuk Post 7 (Flutter vs RN)
(NEWID(), '10000000-0000-0000-0000-000000000017', '00000000-0000-0000-0000-000000000014', NULL, 'Perbandingan yang objektif. Saya pribadi lebih prefer Flutter untuk performance.'),
(NEWID(), '10000000-0000-0000-0000-000000000017', '00000000-0000-0000-0000-000000000018', NULL, 'React Native masih unggul dari sisi ekosistem dan hiring pool menurut saya.');
