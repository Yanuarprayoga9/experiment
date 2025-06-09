-- =====================================================
-- SCHEMA BLOG PLATFORM (SQL Server)
-- =====================================================

CREATE TABLE Pengguna (
    Id              UNIQUEIDENTIFIER NOT NULL 
                    CONSTRAINT PK_Pengguna PRIMARY KEY 
                    DEFAULT NEWID(),
    Nip             VARCHAR(100)      NOT NULL UNIQUE,
    Nama            NVARCHAR(100)     NOT NULL,
    Email           NVARCHAR(100)     NOT NULL,
    KataSandi       NVARCHAR(100)     NOT NULL,
    Peran           VARCHAR(100)      NULL,   -- Contoh: 'SME','Reviewer','Admin','SuperUser'
    Jabatan         VARCHAR(100)      NULL,
    UnitKerja       VARCHAR(100)      NULL,
    DihapusPada     DATETIME2(7)      NULL,
    DibuatPada      DATETIME2(7)      NOT NULL DEFAULT SYSUTCDATETIME()
);

CREATE TABLE PostinganBlog (
    Id              UNIQUEIDENTIFIER NOT NULL 
                    CONSTRAINT PK_PostinganBlog PRIMARY KEY 
                    DEFAULT NEWID(),
    Judul           NVARCHAR(255)     NOT NULL,
    Slug            NVARCHAR(255)     NOT NULL UNIQUE,
    Isi             NVARCHAR(MAX)     NOT NULL,
    PenulisId       UNIQUEIDENTIFIER  NOT NULL,
    StatusReview    VARCHAR(50)       NOT NULL
                    CONSTRAINT CK_PostinganBlog_Status 
                        CHECK (StatusReview IN ('draf','menunggu_review','diterbitkan','ditolak')),
    Visibilitas     VARCHAR(50)       NOT NULL
                    CONSTRAINT CK_PostinganBlog_Visibilitas 
                        CHECK (Visibilitas IN ('internal','publik','terbatas')),
    JumlahTayang    INT               NOT NULL DEFAULT 0,
    UrlThumbnail    NVARCHAR(255)     NULL,
    CreatedAt       DATETIME2(7)      NOT NULL DEFAULT SYSUTCDATETIME(),
    UpdatedAt       DATETIME2(7)      NOT NULL DEFAULT SYSUTCDATETIME(),
    DeletedAt       DATETIME2(7)      NULL,

    CONSTRAINT FK_PostinganBlog_Penulis FOREIGN KEY (PenulisId)
        REFERENCES Pengguna(Id)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION
);

CREATE INDEX IX_PostinganBlog_Status_PenulisId 
    ON PostinganBlog (StatusReview, PenulisId);

CREATE TABLE PostinganLike (
    Id             UNIQUEIDENTIFIER NOT NULL 
                   CONSTRAINT PK_PostinganLike PRIMARY KEY 
                   DEFAULT NEWID(),
    PostinganId    UNIQUEIDENTIFIER NOT NULL,
    PenggunaId     UNIQUEIDENTIFIER NOT NULL,

    CONSTRAINT FK_PostinganLike_Postingan FOREIGN KEY (PostinganId)
        REFERENCES PostinganBlog(Id)
        ON DELETE CASCADE,
    CONSTRAINT FK_PostinganLike_Pengguna FOREIGN KEY (PenggunaId)
        REFERENCES Pengguna(Id)
        ON DELETE NO ACTION,
    CONSTRAINT UQ_PostinganLike UNIQUE (PostinganId, PenggunaId)
);

CREATE TABLE Tag (
    Id           UNIQUEIDENTIFIER NOT NULL 
                 CONSTRAINT PK_Tag PRIMARY KEY 
                 DEFAULT NEWID(),
    NamaTag      NVARCHAR(100)     NOT NULL
);

CREATE TABLE PostinganTag (
    PostinganId  UNIQUEIDENTIFIER NOT NULL,
    TagId        UNIQUEIDENTIFIER NOT NULL,

    CONSTRAINT PK_PostinganTag PRIMARY KEY (PostinganId, TagId),
    CONSTRAINT FK_PostinganTag_Postingan FOREIGN KEY (PostinganId)
        REFERENCES PostinganBlog(Id)
        ON DELETE CASCADE,
    CONSTRAINT FK_PostinganTag_Tag FOREIGN KEY (TagId)
        REFERENCES Tag(Id)
        ON DELETE CASCADE
);

CREATE TABLE KolaboratorBlog (
    Id                   UNIQUEIDENTIFIER  NOT NULL 
                         CONSTRAINT PK_KolaboratorBlog PRIMARY KEY 
                         DEFAULT NEWID(),
    PostinganId          UNIQUEIDENTIFIER  NOT NULL,
    PenggunaId           UNIQUEIDENTIFIER  NOT NULL,
    PeranDalamPostingan  VARCHAR(50)       NOT NULL
                         CONSTRAINT CK_KolaboratorBlog_Peran 
                            CHECK (PeranDalamPostingan IN ('penulisBersama','editor')),
    CreatedAt            DATETIME2(7)      NOT NULL DEFAULT SYSUTCDATETIME(),
    DeletedAt            DATETIME2(7)      NULL,

    CONSTRAINT FK_KolaboratorBlog_Postingan FOREIGN KEY (PostinganId)
        REFERENCES PostinganBlog(Id)
        ON DELETE CASCADE,
    CONSTRAINT FK_KolaboratorBlog_Pengguna FOREIGN KEY (PenggunaId)
        REFERENCES Pengguna(Id)
        ON DELETE NO ACTION
);

CREATE TABLE KomentarBlog (
    Id               UNIQUEIDENTIFIER  NOT NULL 
                     CONSTRAINT PK_KomentarBlog PRIMARY KEY 
                     DEFAULT NEWID(),
    PostinganId      UNIQUEIDENTIFIER  NOT NULL,
    PenggunaId       UNIQUEIDENTIFIER  NOT NULL,
    KomentarIndukId  UNIQUEIDENTIFIER  NULL,
    Konten           NVARCHAR(MAX)     NOT NULL,
    CreatedAt        DATETIME2(7)      NOT NULL DEFAULT SYSUTCDATETIME(),
    DeletedAt        DATETIME2(7)      NULL,

    CONSTRAINT FK_KomentarBlog_Postingan FOREIGN KEY (PostinganId)
        REFERENCES PostinganBlog(Id)
        ON DELETE CASCADE,
    CONSTRAINT FK_KomentarBlog_Pengguna FOREIGN KEY (PenggunaId)
        REFERENCES Pengguna(Id)
        ON DELETE NO ACTION,
    CONSTRAINT FK_KomentarBlog_ParentComment FOREIGN KEY (KomentarIndukId)
        REFERENCES KomentarBlog(Id)
        ON DELETE NO ACTION
);

CREATE TABLE KomentarLike (
    Id             UNIQUEIDENTIFIER NOT NULL 
                   CONSTRAINT PK_KomentarLike PRIMARY KEY 
                   DEFAULT NEWID(),
    KomentarId     UNIQUEIDENTIFIER NOT NULL,
    PenggunaId     UNIQUEIDENTIFIER NOT NULL,
    CreatedAt      DATETIME2(7)     NOT NULL DEFAULT SYSUTCDATETIME(),

    CONSTRAINT FK_KomentarLike_Komentar FOREIGN KEY (KomentarId)
        REFERENCES KomentarBlog(Id)
        ON DELETE CASCADE,
    CONSTRAINT FK_KomentarLike_Pengguna FOREIGN KEY (PenggunaId)
        REFERENCES Pengguna(Id)
        ON DELETE NO ACTION,
    CONSTRAINT UQ_KomentarLike UNIQUE (KomentarId, PenggunaId)
);

CREATE TABLE StatistikBlog (
    Id               UNIQUEIDENTIFIER  NOT NULL 
                     CONSTRAINT PK_StatistikBlog PRIMARY KEY 
                     DEFAULT NEWID(),
    PostinganId      UNIQUEIDENTIFIER  NOT NULL,
    TanggalStat      DATE              NOT NULL,
    JumlahTayang     INT               NOT NULL DEFAULT 0,
    JumlahKomentar   INT               NOT NULL DEFAULT 0,
    JumlahSuka       INT               NOT NULL DEFAULT 0,

    CONSTRAINT UQ_StatistikBlog_PostinganTanggal UNIQUE (PostinganId, TanggalStat),
    CONSTRAINT FK_StatistikBlog_Postingan FOREIGN KEY (PostinganId)
        REFERENCES PostinganBlog(Id)
        ON DELETE CASCADE
);


-- 1. Pengguna
INSERT INTO Pengguna (Id, Nip, Nama, Email, KataSandi, Peran, Jabatan, UnitKerja)
VALUES 
('00000000-0000-0000-0000-000000000001', '123456', 'Andi', 'andi@example.com', 'password123', 'SME', 'Dosen', 'Informatika'),
('00000000-0000-0000-0000-000000000002', '234567', 'Budi', 'budi@example.com', 'password123', 'Reviewer', 'Dosen', 'Informatika'),
('00000000-0000-0000-0000-000000000003', '345678', 'Citra', 'citra@example.com', 'password123', 'Admin', 'Kaprodi', 'Sistem Informasi'),
('00000000-0000-0000-0000-000000000004', '456789', 'Dewi', 'dewi@example.com', 'password123', 'SuperUser', 'Dosen', 'Teknik Sipil'),
('00000000-0000-0000-0000-000000000005', '567890', 'Eka', 'eka@example.com', 'password123', 'SME', 'Mahasiswa', 'Informatika');

-- 2. PostinganBlog
INSERT INTO PostinganBlog (Id, Judul, Slug, Isi, PenulisId, StatusReview, Visibilitas)
VALUES 
('10000000-0000-0000-0000-000000000001', 'Belajar Golang Dasar', 'belajar-golang-dasar', 'Ini adalah isi artikel Golang.', '00000000-0000-0000-0000-000000000001', 'diterbitkan', 'publik'),
('10000000-0000-0000-0000-000000000002', 'Tutorial SQL Server', 'tutorial-sql-server', 'Tutorial lengkap SQL Server.', '00000000-0000-0000-0000-000000000002', 'menunggu_review', 'internal');

-- 3. Tag
INSERT INTO Tag (Id, NamaTag)
VALUES 
('20000000-0000-0000-0000-000000000001', 'Golang'),
('20000000-0000-0000-0000-000000000002', 'SQL'),
('20000000-0000-0000-0000-000000000003', 'Pemrograman'),
('20000000-0000-0000-0000-000000000004', 'Database'),
('20000000-0000-0000-0000-000000000005', 'Tutorial');

-- 4. PostinganTag
INSERT INTO PostinganTag (PostinganId, TagId)
VALUES 
('10000000-0000-0000-0000-000000000001', '20000000-0000-0000-0000-000000000001'),
('10000000-0000-0000-0000-000000000001', '20000000-0000-0000-0000-000000000003'),
('10000000-0000-0000-0000-000000000002', '20000000-0000-0000-0000-000000000002'),
('10000000-0000-0000-0000-000000000002', '20000000-0000-0000-0000-000000000004'),
('10000000-0000-0000-0000-000000000002', '20000000-0000-0000-0000-000000000005');

-- 5. PostinganLike
INSERT INTO PostinganLike (Id, PostinganId, PenggunaId)
VALUES 
(NEWID(), '10000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000002'),
(NEWID(), '10000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000003'),
(NEWID(), '10000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000004'),
(NEWID(), '10000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000005'),
(NEWID(), '10000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000001');

-- 6. KomentarBlog
INSERT INTO KomentarBlog (Id, PostinganId, PenggunaId, KomentarIndukId, Konten)
VALUES 
('30000000-0000-0000-0000-000000000001', '10000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000002', NULL, 'Artikel ini sangat membantu!'),
('30000000-0000-0000-0000-000000000002', '10000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000003', NULL, 'Terima kasih untuk tutorialnya.'),
('30000000-0000-0000-0000-000000000003', '10000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000004', NULL, 'Ini sangat jelas dan bermanfaat.'),
('30000000-0000-0000-0000-000000000004', '10000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000005', NULL, 'Saya akan mencobanya sekarang.'),
('30000000-0000-0000-0000-000000000005', '10000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000001', NULL, 'Mantap sekali!');

-- 7. KomentarLike
INSERT INTO KomentarLike (Id, KomentarId, PenggunaId)
VALUES 
(NEWID(), '30000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000001'),
(NEWID(), '30000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000004'),
(NEWID(), '30000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000005'),
(NEWID(), '30000000-0000-0000-0000-000000000003', '00000000-0000-0000-0000-000000000002'),
(NEWID(), '30000000-0000-0000-0000-000000000004', '00000000-0000-0000-0000-000000000003');

-- =====================================================
-- QUERY MENGHITUNG LIKE DAN KOMENTAR BLOG
-- =====================================================

-- 1. MENAMPILKAN LIST POSTINGAN DENGAN JUMLAH LIKE DAN KOMENTAR
SELECT 
    pb.Id,
    pb.Judul,
    pb.Slug,
    pb.StatusReview,
    pb.Visibilitas,
    pb.JumlahTayang,
    p.Nama AS Penulis,
    p.Peran AS PeranPenulis,
    pb.CreatedAt,
    pb.UpdatedAt,
    -- Hitung jumlah like
    COALESCE(like_count.JumlahLike, 0) AS JumlahLike,
    -- Hitung jumlah komentar (tidak termasuk yang dihapus)
    COALESCE(comment_count.JumlahKomentar, 0) AS JumlahKomentar,
    -- Hitung total interaksi
    (COALESCE(like_count.JumlahLike, 0) + COALESCE(comment_count.JumlahKomentar, 0)) AS TotalInteraksi
FROM PostinganBlog pb
INNER JOIN Pengguna p ON pb.PenulisId = p.Id
-- Join untuk menghitung like
LEFT JOIN (
    SELECT 
        PostinganId,
        COUNT(*) AS JumlahLike
    FROM PostinganLike
    GROUP BY PostinganId
) like_count ON pb.Id = like_count.PostinganId
-- Join untuk menghitung komentar
LEFT JOIN (
    SELECT 
        PostinganId,
        COUNT(*) AS JumlahKomentar
    FROM KomentarBlog
    WHERE DeletedAt IS NULL
    GROUP BY PostinganId
) comment_count ON pb.Id = comment_count.PostinganId
WHERE pb.DeletedAt IS NULL
ORDER BY TotalInteraksi DESC, pb.CreatedAt DESC;

-- =====================================================

-- 2. DETAIL POSTINGAN DENGAN LIKE DAN KOMENTAR LENGKAP
SELECT 
    pb.Id AS PostinganId,
    pb.Judul,
    pb.Slug,
    pb.Isi,
    pb.StatusReview,
    pb.Visibilitas,
    pb.JumlahTayang,
    pb.UrlThumbnail,
    p.Nama AS Penulis,
    p.Email AS EmailPenulis,
    p.Jabatan,
    p.UnitKerja,
    pb.CreatedAt,
    pb.UpdatedAt,
    -- Statistik
    COALESCE(stats.JumlahLike, 0) AS JumlahLike,
    COALESCE(stats.JumlahKomentar, 0) AS JumlahKomentar,
    COALESCE(stats.JumlahKomentarAktif, 0) AS JumlahKomentarAktif
FROM PostinganBlog pb
INNER JOIN Pengguna p ON pb.PenulisId = p.Id
LEFT JOIN (
    SELECT 
        pb2.Id,
        COUNT(DISTINCT pl.Id) AS JumlahLike,
        COUNT(DISTINCT kb.Id) AS JumlahKomentar,
        COUNT(DISTINCT CASE WHEN kb.DeletedAt IS NULL THEN kb.Id END) AS JumlahKomentarAktif
    FROM PostinganBlog pb2
    LEFT JOIN PostinganLike pl ON pb2.Id = pl.PostinganId
    LEFT JOIN KomentarBlog kb ON pb2.Id = kb.PostinganId
    GROUP BY pb2.Id
) stats ON pb.Id = stats.Id
WHERE pb.DeletedAt IS NULL;

-- =====================================================

-- 3. MENAMPILKAN POSTINGAN DENGAN DAFTAR YANG MEMBERIKAN LIKE
SELECT 
    pb.Id AS PostinganId,
    pb.Judul,
    pb.Slug,
    p_penulis.Nama AS Penulis,
    -- Daftar yang memberikan like
    GROUP_CONCAT(
        CONCAT(p_liker.Nama, ' (', p_liker.Jabatan, ')')
        ORDER BY p_liker.Nama
        SEPARATOR '; '
    ) AS DaftarPemberianLike,
    COUNT(pl.Id) AS JumlahLike
FROM PostinganBlog pb
INNER JOIN Pengguna p_penulis ON pb.PenulisId = p_penulis.Id
LEFT JOIN PostinganLike pl ON pb.Id = pl.PostinganId
LEFT JOIN Pengguna p_liker ON pl.PenggunaId = p_liker.Id
WHERE pb.DeletedAt IS NULL
GROUP BY pb.Id, pb.Judul, pb.Slug, p_penulis.Nama
ORDER BY JumlahLike DESC;

-- =====================================================

-- 4. MENAMPILKAN POSTINGAN DENGAN DAFTAR KOMENTAR
SELECT 
    pb.Id AS PostinganId,
    pb.Judul,
    pb.Slug,
    p_penulis.Nama AS Penulis,
    kb.Id AS KomentarId,
    p_komentator.Nama AS Komentator,
    p_komentator.Jabatan AS JabatanKomentator,
    kb.Konten AS IsiKomentar,
    kb.CreatedAt AS WaktuKomentar,
    CASE 
        WHEN kb.KomentarIndukId IS NULL THEN 'Komentar Utama'
        ELSE 'Balasan Komentar'
    END AS TipeKomentar,
    -- Hitung like pada komentar ini
    COALESCE(kl_count.JumlahLikeKomentar, 0) AS JumlahLikeKomentar
FROM PostinganBlog pb
INNER JOIN Pengguna p_penulis ON pb.PenulisId = p_penulis.Id
LEFT JOIN KomentarBlog kb ON pb.Id = kb.PostinganId AND kb.DeletedAt IS NULL
LEFT JOIN Pengguna p_komentator ON kb.PenggunaId = p_komentator.Id
LEFT JOIN (
    SELECT 
        KomentarId,
        COUNT(*) AS JumlahLikeKomentar
    FROM KomentarLike
    GROUP BY KomentarId
) kl_count ON kb.Id = kl_count.KomentarId
WHERE pb.DeletedAt IS NULL
ORDER BY pb.CreatedAt DESC, kb.CreatedAt ASC;

-- =====================================================

-- 5. STATISTIK LENGKAP PER POSTINGAN
SELECT 
    pb.Id AS PostinganId,
    pb.Judul,
    pb.Slug,
    p.Nama AS Penulis,
    pb.StatusReview,
    pb.Visibilitas,
    pb.JumlahTayang,
    pb.CreatedAt,
    
    -- Statistik Like
    COALESCE(like_stats.JumlahLike, 0) AS JumlahLike,
    
    -- Statistik Komentar
    COALESCE(comment_stats.JumlahKomentar, 0) AS JumlahKomentar,
    COALESCE(comment_stats.JumlahKomentarUtama, 0) AS JumlahKomentarUtama,
    COALESCE(comment_stats.JumlahBalasan, 0) AS JumlahBalasan,
    
    -- Statistik Like pada Komentar
    COALESCE(comment_like_stats.TotalLikeKomentar, 0) AS TotalLikeKomentar,
    
    -- Total Engagement
    (COALESCE(like_stats.JumlahLike, 0) + 
     COALESCE(comment_stats.JumlahKomentar, 0) + 
     COALESCE(comment_like_stats.TotalLikeKomentar, 0)) AS TotalEngagement,
    
    -- Engagement Rate (berdasarkan views)
    CASE 
        WHEN pb.JumlahTayang > 0 THEN 
            ROUND(
                ((COALESCE(like_stats.JumlahLike, 0) + 
                  COALESCE(comment_stats.JumlahKomentar, 0) + 
                  COALESCE(comment_like_stats.TotalLikeKomentar, 0)) * 100.0 / pb.JumlahTayang), 2
            )
        ELSE 0 
    END AS EngagementRate
    
FROM PostinganBlog pb
INNER JOIN Pengguna p ON pb.PenulisId = p.Id

-- Statistik Like Postingan
LEFT JOIN (
    SELECT 
        PostinganId,
        COUNT(*) AS JumlahLike
    FROM PostinganLike
    GROUP BY PostinganId
) like_stats ON pb.Id = like_stats.PostinganId

-- Statistik Komentar
LEFT JOIN (
    SELECT 
        PostinganId,
        COUNT(*) AS JumlahKomentar,
        COUNT(CASE WHEN KomentarIndukId IS NULL THEN 1 END) AS JumlahKomentarUtama,
        COUNT(CASE WHEN KomentarIndukId IS NOT NULL THEN 1 END) AS JumlahBalasan
    FROM KomentarBlog
    WHERE DeletedAt IS NULL
    GROUP BY PostinganId
) comment_stats ON pb.Id = comment_stats.PostinganId

-- Statistik Like pada Komentar
LEFT JOIN (
    SELECT 
        kb.PostinganId,
        COUNT(kl.Id) AS TotalLikeKomentar
    FROM KomentarBlog kb
    INNER JOIN KomentarLike kl ON kb.Id = kl.KomentarId
    WHERE kb.DeletedAt IS NULL
    GROUP BY kb.PostinganId
) comment_like_stats ON pb.Id = comment_like_stats.PostinganId

WHERE pb.DeletedAt IS NULL
ORDER BY TotalEngagement DESC, pb.CreatedAt DESC;

-- =====================================================

-- 6. TOP POSTINGAN BERDASARKAN ENGAGEMENT
SELECT 
    RANK() OVER (ORDER BY TotalEngagement DESC) AS Ranking,
    PostinganId,
    Judul,
    Penulis,
    StatusReview,
    JumlahTayang,
    JumlahLike,
    JumlahKomentar,
    TotalLikeKomentar,
    TotalEngagement,
    EngagementRate
FROM (
    SELECT 
        pb.Id AS PostinganId,
        pb.Judul,
        p.Nama AS Penulis,
        pb.StatusReview,
        pb.JumlahTayang,
        COALESCE(like_stats.JumlahLike, 0) AS JumlahLike,
        COALESCE(comment_stats.JumlahKomentar, 0) AS JumlahKomentar,
        COALESCE(comment_like_stats.TotalLikeKomentar, 0) AS TotalLikeKomentar,
        (COALESCE(like_stats.JumlahLike, 0) + 
         COALESCE(comment_stats.JumlahKomentar, 0) + 
         COALESCE(comment_like_stats.TotalLikeKomentar, 0)) AS TotalEngagement,
        CASE 
            WHEN pb.JumlahTayang > 0 THEN 
                ROUND(
                    ((COALESCE(like_stats.JumlahLike, 0) + 
                      COALESCE(comment_stats.JumlahKomentar, 0) + 
                      COALESCE(comment_like_stats.TotalLikeKomentar, 0)) * 100.0 / pb.JumlahTayang), 2
                )
            ELSE 0 
        END AS EngagementRate
        
    FROM PostinganBlog pb
    INNER JOIN Pengguna p ON pb.PenulisId = p.Id
    LEFT JOIN (
        SELECT PostinganId, COUNT(*) AS JumlahLike
        FROM PostinganLike GROUP BY PostinganId
    ) like_stats ON pb.Id = like_stats.PostinganId
    LEFT JOIN (
        SELECT PostinganId, COUNT(*) AS JumlahKomentar
        FROM KomentarBlog WHERE DeletedAt IS NULL GROUP BY PostinganId
    ) comment_stats ON pb.Id = comment_stats.PostinganId
    LEFT JOIN (
        SELECT kb.PostinganId, COUNT(kl.Id) AS TotalLikeKomentar
        FROM KomentarBlog kb INNER JOIN KomentarLike kl ON kb.Id = kl.KomentarId
        WHERE kb.DeletedAt IS NULL GROUP BY kb.PostinganId
    ) comment_like_stats ON pb.Id = comment_like_stats.PostinganId
    
    WHERE pb.DeletedAt IS NULL AND pb.StatusReview = 'diterbitkan'
) AS engagement_data
ORDER BY Ranking;