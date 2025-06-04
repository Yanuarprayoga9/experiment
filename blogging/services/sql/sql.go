package sqlquery

const CASBIN_CREATE = `
-- Casbin Rules
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[CasbinRules]') AND type = 'U')
BEGIN
    CREATE TABLE CasbinRules(
        p_type NVARCHAR(32) DEFAULT '' NOT NULL,
        v0 NVARCHAR(255) DEFAULT '' NOT NULL,
        v1 NVARCHAR(255) DEFAULT '' NOT NULL,
        v2 NVARCHAR(255) DEFAULT '' NOT NULL,
        v3 NVARCHAR(255) DEFAULT '' NOT NULL,
        v4 NVARCHAR(255) DEFAULT '' NOT NULL,
        v5 NVARCHAR(255) DEFAULT '' NOT NULL
    );
    CREATE INDEX idx_CasbinRules ON CasbinRules (p_type, v0, v1);
END;

-- Users
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Users]') AND type = 'U')
BEGIN
    CREATE TABLE Users (
        Id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
        Email NVARCHAR(100) NOT NULL UNIQUE,
        Name NVARCHAR(100) NOT NULL,
        Password NVARCHAR(100) NOT NULL,
        DeletedAt DATETIME2 NULL
    );
END;

-- Users Attributes
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[UsersAttributes]') AND type = 'U')
BEGIN
    CREATE TABLE UsersAttributes (
        Id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
        UserId UNIQUEIDENTIFIER NOT NULL,
        AttributeKey NVARCHAR(100) NOT NULL,
        AttributeValue NVARCHAR(100) NOT NULL,
        DeletedAt DATETIME2 NULL,
    );
END;


-- User Groups
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[UserGroups]') AND type = 'U')
BEGIN
    CREATE TABLE UserGroups (
        Id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
        Name NVARCHAR(100) NOT NULL,
        Description NVARCHAR(255) NULL,
        DeletedAt DATETIME2 NULL
    );
END;

-- User-Group Memberships
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[UserGroupMemberships]') AND type = 'U')
BEGIN
    CREATE TABLE UserGroupMemberships (
        UserId UNIQUEIDENTIFIER NOT NULL,
        GroupId UNIQUEIDENTIFIER NOT NULL,
        DeletedAt DATETIME2 NULL,
        PRIMARY KEY (UserId, GroupId)
    );
END;

-- Permissions
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Permissions]') AND type = 'U')
BEGIN
    CREATE TABLE Permissions (
        Id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
        Action NVARCHAR(50) NOT NULL,
        Resource NVARCHAR(200) NOT NULL,
        DeletedAt DATETIME2 NULL,
        CONSTRAINT UQ_Permissions UNIQUE (Action, Resource)
    );
END;

-- Group-Permission Mapping
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[GroupPermissions]') AND type = 'U')
BEGIN
    CREATE TABLE GroupPermissions (
        GroupId UNIQUEIDENTIFIER NOT NULL,
        PermissionId UNIQUEIDENTIFIER NOT NULL,
        DeletedAt DATETIME2 NULL,
        PRIMARY KEY (GroupId, PermissionId)
    );
END;

-- Coaching
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Coaching]') AND type = 'U')
BEGIN
    CREATE TABLE Coaching (
        Id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
        Judul NVARCHAR(255) NOT NULL,
        Tujuan NVARCHAR(MAX) NULL,
        JamMulai DATETIME2 NOT NULL,
        Coach UNIQUEIDENTIFIER,
        Coachee UNIQUEIDENTIFIER,
        CreatedAt DATETIME2 NOT NULL DEFAULT SYSDATETIME(),
        DeletedAt DATETIME2 NULL,
        Status INT NOT NULL DEFAULT 0,
        Approved BIT NOT NULL DEFAULT 0,
        ApprovedBy UNIQUEIDENTIFIER NULL,
        ApprovedAt DATETIME2 NULL
    );
END;


-- TODO: Fix this structure
-- Coaching
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Mentoring]') AND type = 'U')
BEGIN
    CREATE TABLE Mentoring (
        Id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
        Judul NVARCHAR(255) NOT NULL,
        Tujuan NVARCHAR(MAX) NULL,
        JamMulai DATETIME2 NOT NULL,
        Coach UNIQUEIDENTIFIER,
        Coachee UNIQUEIDENTIFIER,
        CreatedAt DATETIME2 NOT NULL DEFAULT SYSDATETIME(),
        DeletedAt DATETIME2 NULL,
        Status INT NOT NULL DEFAULT 0,
        Approved BIT NOT NULL DEFAULT 0,
        ApprovedBy UNIQUEIDENTIFIER NULL,
        ApprovedAt DATETIME2 NULL
    );
END;

-- Coaching
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[CoachingFormPendaftaran]') AND type = 'U')
BEGIN
    CREATE TABLE CoachingFormPendaftaran (
        Id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
        CoachingId UNIQUEIDENTIFIER,
        Pertanyaan UNIQUEIDENTIFIER,
        Jawaban TEXT
    );
END;


-- Mentoring
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[MentoringFormPendaftaran]') AND type = 'U')
BEGIN
    CREATE TABLE MentoringFormPendaftaran (
        Id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
        CoachingId UNIQUEIDENTIFIER,
        Pertanyaan UNIQUEIDENTIFIER,
        Jawaban TEXT
    );
END;


-- Coaching Sesi
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[CoachingSesi]') AND type = 'U')
BEGIN
    CREATE TABLE CoachingSesi (
        Id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
        CoachingId UNIQUEIDENTIFIER NOT NULL,
        [Order] INT NOT NULL DEFAULT 1,
        JamMulai DATETIME2 NOT NULL,
        JamSelesai DATETIME2 NOT NULL,
        ApprovedAt DATETIME2 NULL,
        ApprovedBy UNIQUEIDENTIFIER NULL
    );
END;

-- Coaching Sesi Catatan
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[CoachingSesiCatatan]') AND type = 'U')
BEGIN
    CREATE TABLE CoachingSesiCatatan (
        Id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
        SesiCoachingId UNIQUEIDENTIFIER NOT NULL,
        FormId UNIQUEIDENTIFIER NOT NULL,
        [Key] NVARCHAR(255) NOT NULL,
        Value NVARCHAR(255) NULL,
        CreatedAt DATETIME2 NOT NULL DEFAULT SYSDATETIME(),
        CreatedBy UNIQUEIDENTIFIER NOT NULL,
        Approved BIT NOT NULL DEFAULT 0,
        ApprovedBy UNIQUEIDENTIFIER NULL,
        ApprovedAt DATETIME2 NULL
    );
END;

-- Mentoring
-- IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Mentoring]') AND type = 'U')
-- BEGIN
    -- CREATE TABLE Mentoring (
        -- Id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
        -- Judul NVARCHAR(255) NOT NULL,
        -- Tujuan NVARCHAR(MAX) NULL,
        -- Coach UNIQUEIDENTIFIER NOT NULL,
        -- JamMulai DATETIME2 NOT NULL,
        -- JamSelesai DATETIME2 NOT NULL,
        -- Approved BIT NOT NULL DEFAULT 0,
        -- ApprovedBy UNIQUEIDENTIFIER NULL,
        -- CreatedAt DATETIME2 NOT NULL DEFAULT SYSDATETIME(),
        -- DeletedAt DATETIME2 NULL
    -- );
-- END;

-- Attachments
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[Attachments]') AND type = 'U')
BEGIN
    CREATE TABLE Attachments (
        Id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
        Resource NVARCHAR(255) NOT NULL,
        ResourceId NVARCHAR(255) NOT NULL,
        Link NVARCHAR(255) NOT NULL,
        Filename NVARCHAR(255) NOT NULL,
        DeletedAt DATETIME2 NULL,
        CreatedAt DATETIME2 NOT NULL DEFAULT SYSDATETIME()
    );
END;

-- Master Data
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[MasterData]') AND type = 'U')
BEGIN
    CREATE TABLE MasterData (
        Id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
        Resource NVARCHAR(255) NOT NULL,
        [Key] NVARCHAR(255) NOT NULL,
        Description NVARCHAR(255) NOT NULL,
        DeletedAt DATETIME2 NULL,
        CreatedAt DATETIME2 NOT NULL DEFAULT SYSDATETIME()
    );
END;

-- Master Form
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[MasterForm]') AND type = 'U')
BEGIN
    CREATE TABLE MasterForm (
        Id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
        Resource NVARCHAR(255) NOT NULL,
        [Key] NVARCHAR(255) NOT NULL,
        Label NVARCHAR(255) NOT NULL,
        [Type] NVARCHAR(255) NOT NULL,
        [Order] INT NOT NULL,
        DeletedAt DATETIME2 NULL
    );
END;

-- Master Module
IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[MasterModule]') AND type = 'U')
BEGIN
    CREATE TABLE MasterModule (
        Id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
        Name NVARCHAR(255) NOT NULL,
        URL NVARCHAR(255) NOT NULL,
        ParentId UNIQUEIDENTIFIER NULL,
        [Order] INT NOT NULL,
        DeletedAt DATETIME2 NULL,
		AdditionalActions NVARCHAR(MAX) NULL
    );
END;
`

const CASBIN_VIEW = ``
// const CASBIN_VIEW = `

// EXECUTE('CREATE OR ALTER VIEW dbo.CoachingWithUsersView AS
// SELECT
//     ISNULL(c.Id, ''00000000-0000-0000-0000-000000000000'') AS CoachingId,
//     ISNULL(c.Judul, ''-'') AS Judul,
//     ISNULL(c.Tujuan, ''-'') AS Tujuan,
//     ISNULL(c.Coach, ''00000000-0000-0000-0000-000000000000'') AS CoachId,
//     ISNULL(u.Name, ''-'') AS CoachName,
//     ISNULL(u.Email, ''-'') AS CoachEmail,
//     ISNULL(c.Coachee, ''00000000-0000-0000-0000-000000000000'') AS CoacheeId,
//     ISNULL(uc.Name, ''-'') AS CoacheeName,
//     ISNULL(uc.Email, ''-'') AS CoacheeEmail,
//     c.CreatedAt AS CreatedAt,
//     c.DeletedAt AS CoachingDeletedAt,
//     u.DeletedAt AS CoachDeletedAt,
//     c.JamMulai as JamMulai,
//     c.status
// FROM
//     dbo.Coaching c
// LEFT JOIN
//     KMS.dbo.Users u ON c.Coach = u.Id AND u.DeletedAt IS NULL
// LEFT JOIN
//     KMS.dbo.Users uc ON c.Coachee = uc.Id AND u.DeletedAt IS NULL
// WHERE
//     c.DeletedAt IS NULL;
// ');

// EXECUTE('CREATE OR ALTER VIEW dbo.MentoringWithUsersView AS
// SELECT
//     ISNULL(c.Id, ''00000000-0000-0000-0000-000000000000'') AS CoachingId,
//     ISNULL(c.Judul, ''-'') AS Judul,
//     ISNULL(c.Tujuan, ''-'') AS Tujuan,
//     ISNULL(c.Coach, ''00000000-0000-0000-0000-000000000000'') AS CoachId,
//     ISNULL(u.Name, ''-'') AS CoachName,
//     ISNULL(u.Email, ''-'') AS CoachEmail,
//     ISNULL(c.Coachee, ''00000000-0000-0000-0000-000000000000'') AS CoacheeId,
//     ISNULL(uc.Name, ''-'') AS CoacheeName,
//     ISNULL(uc.Email, ''-'') AS CoacheeEmail,
//     c.CreatedAt AS CreatedAt,
//     c.DeletedAt AS CoachingDeletedAt,
//     u.DeletedAt AS CoachDeletedAt,
//     c.JamMulai as JamMulai,
//     c.status
// FROM
//     dbo.Mentoring c
// LEFT JOIN
//     KMS.dbo.Users u ON c.Coach = u.Id AND u.DeletedAt IS NULL
// LEFT JOIN
//     KMS.dbo.Users uc ON c.Coachee = uc.Id AND u.DeletedAt IS NULL
// WHERE
//     c.DeletedAt IS NULL;
// ');

// `

const CASBIN_INIT = ``

// const CASBIN_INIT = `
// INSERT INTO KMS.dbo.Users (Id, Email, Name, DeletedAt, Password)
// VALUES
// (NEWID(), 'andi.wijaya@example.com', 'Andi Wijaya', NULL, 'TCE9C1KbHekY8A$B159f3lDyB5e1Eghz5GjohTZBIqmB0EkzN3ZPlBIFQI'),
// (NEWID(), 'sri.lestari@example.com', 'Sri Lestari', NULL, 'Npq3hRR/UxoZqA$pU/67bzsBOgw1huWyYA0ibxiHZsbXuMa4ybVtZ1dtZM'),
// (NEWID(), 'budi.santoso@example.com', 'Budi Santoso', NULL, 'BrB5i60GhsdriA$2/TkBs4MVMQRBIJaqMWotPBsC1+CjQtGwpxfk53BzBs');

// -- Ambil UserId berdasarkan Email
// DECLARE @AndiId uniqueidentifier = (SELECT Id FROM KMS.dbo.Users WHERE Email = 'andi.wijaya@example.com');
// DECLARE @SriId uniqueidentifier = (SELECT Id FROM KMS.dbo.Users WHERE Email = 'sri.lestari@example.com');
// DECLARE @BudiId uniqueidentifier = (SELECT Id FROM KMS.dbo.Users WHERE Email = 'budi.santoso@example.com');

// -- Tambahkan atribut JABATAN dan UNIT_KERJA
// INSERT INTO KMS.dbo.UsersAttributes (Id, UserId, AttributeKey, AttributeValue, DeletedAt)
// VALUES
// (NEWID(), @AndiId, 'JABATAN', 'Kepala Sub Bagian', NULL),
// (NEWID(), @AndiId, 'UNIT_KERJA', 'Bagian Umum', NULL),

// (NEWID(), @SriId, 'JABATAN', 'Staf Administrasi', NULL),
// (NEWID(), @SriId, 'UNIT_KERJA', 'Bagian Keuangan', NULL),

// (NEWID(), @BudiId, 'JABATAN', 'Analis Data', NULL),
// (NEWID(), @BudiId, 'UNIT_KERJA', 'Bagian Teknologi Informasi', NULL);

// `

// Query dashboard
// SELECT 
// 	Coach,
// 	Coachee,
// 	ISNULL([Menunggu pra-coaching], 0) as MenungguPraCoaching,
// 	ISNULL([Menunggu Persetujuan], 0) as MenungguPersetujuan,
// 	ISNULL([Terjadwal], 0) as Terjadwal,
// 	ISNULL([Sedang Berlangsung], 0) as SedangBerlangsung,
// 	ISNULL([Selesai], 0) as Selesai,
// 	ISNULL([Revisi], 0) as Revisi
// FROM
// 	(
// 	SELECT
// 		u.Name as Coachee,
// 		v.Name as Coach,
// 		md.Description as Status
// 	FROM
// 		Coaching c
// 	left join Users u on
// 		u.Id = c.Coachee
// 	left join Users v on
// 		v.Id = c.Coach
// 	left join MasterData md on
// 		md.[Key] = c.Status
// 		and md.Resource = 'status coaching'
// ) as SourceTable
// PIVOT (
// 	Count(Status)
// 		FOR Status IN ([Menunggu pra-coaching], [Sedang Berlangsung], [Menunggu Persetujuan], [Terjadwal], [Selesai], [Revisi])) as pt
// order by
// Coach,	
// Coachee


// WITH CoacheeCount AS (
// SELECT
// 	v.Name AS Coach,
// 	COUNT(DISTINCT u.Id) AS TotalCoachee
// FROM
// 	Coaching c
// LEFT JOIN Users u ON
// 	u.Id = c.Coachee
// LEFT JOIN Users v ON
// 	v.Id = c.Coach
// GROUP BY
// 	v.Name
// )
// SELECT 
// 	pt.Coach,
// 	cc.TotalCoachee,
// 	ISNULL([Menunggu pra-coaching], 0) as MenungguPraCoaching,
// 	ISNULL([Menunggu Persetujuan], 0) as MenungguPersetujuan,
// 	ISNULL([Terjadwal], 0) as Terjadwal,
// 	ISNULL([Sedang Berlangsung], 0) as SedangBerlangsung,
// 	ISNULL([Selesai], 0) as Selesai,
// 	ISNULL([Revisi], 0) as Revisi
// FROM
// 	(
// 	SELECT
// 		v.Name as Coach,
// 		md.Description as Status
// 	FROM
// 		Coaching c
// 	left join Users u on
// 		u.Id = c.Coachee
// 	left join Users v on
// 		v.Id = c.Coach
// 	left join MasterData md on
// 		md.[Key] = c.Status
// 		and md.Resource = 'status coaching'
// ) as SourceTable
// PIVOT (
// 	Count(Status)
// 	FOR Status IN ([Menunggu pra-coaching], [Sedang Berlangsung], [Menunggu Persetujuan], [Terjadwal], [Selesai], [Revisi])) as pt
// left join CoacheeCount cc on
// 	pt.Coach = cc.Coach
// order by
// 	pt.Coach
