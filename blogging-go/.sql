CREATE TABLE casbin_rule (
    id INT PRIMARY KEY IDENTITY,
    ptype VARCHAR(100) NOT NULL,
    v0 VARCHAR(100),
    v1 VARCHAR(100),
    v2 VARCHAR(100),
    v3 VARCHAR(100),
    v4 VARCHAR(100),
    v5 VARCHAR(100)
);


-- Tabel Users (sinkron dari LDAP)
CREATE TABLE users (
    id INT PRIMARY KEY IDENTITY,
    ldap_uid VARCHAR(100) UNIQUE NOT NULL,
    username VARCHAR(100),
    email VARCHAR(100)
);

-- Tabel User Groups
CREATE TABLE user_groups (
    id INT PRIMARY KEY IDENTITY,
    name VARCHAR(100) UNIQUE NOT NULL
);

-- Relasi: user ke grup
CREATE TABLE user_group_membership (
    id INT PRIMARY KEY IDENTITY,
    user_id INT NOT NULL,
    group_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES user_groups(id) ON DELETE CASCADE
);

-- Tabel Permissions
CREATE TABLE permissions (
    id INT PRIMARY KEY IDENTITY,
    name VARCHAR(100) NOT NULL,
    object VARCHAR(100) NOT NULL,
    action VARCHAR(100) NOT NULL,
    UNIQUE (object, action)
);

-- Relasi: grup ke izin
CREATE TABLE group_permissions (
    id INT PRIMARY KEY IDENTITY,
    group_id INT NOT NULL,
    permission_id INT NOT NULL,
    FOREIGN KEY (group_id) REFERENCES user_groups(id) ON DELETE CASCADE,
    FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE
);


-- Tambahkan grup
INSERT INTO user_groups (name) VALUES ('mahasiswa'), ('informatika');

-- Tambahkan permission
INSERT INTO permissions (name, object, action)
VALUES 
  ('post:create', 'post', 'create'),
  ('comment:reply', 'comment', 'reply');

-- Tambahkan user (dari LDAP)
INSERT INTO users (ldap_uid, username, email)
VALUES ('12345', 'budi', 'budi@example.com');

-- Tambahkan membership user ke grup
INSERT INTO user_group_membership (user_id, group_id)
VALUES (1, 1); -- budi → mahasiswa

-- Tambahkan izin ke grup
INSERT INTO group_permissions (group_id, permission_id)
VALUES 
  (1, 1), -- mahasiswa → post:create
  (1, 2); -- mahasiswa → comment:reply
