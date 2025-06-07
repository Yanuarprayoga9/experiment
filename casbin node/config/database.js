

// config/database.js
const mysql = require('mysql2/promise');
require('dotenv').config();

const dbConfig = {
  host: process.env.DB_HOST,
  user: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  database: process.env.DB_NAME,
  waitForConnections: true,
  connectionLimit: 10,
  queueLimit: 0
};

const pool = mysql.createPool(dbConfig);

// Fungsi untuk membuat tabel casbin_rule jika belum ada
async function createCasbinTable() {
  try {
    const createTableQuery = `
      CREATE TABLE IF NOT EXISTS casbin_rule (
        id INT AUTO_INCREMENT PRIMARY KEY,
        ptype VARCHAR(255) NOT NULL,
        v0 VARCHAR(255),
        v1 VARCHAR(255),
        v2 VARCHAR(255),
        v3 VARCHAR(255),
        v4 VARCHAR(255),
        v5 VARCHAR(255),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
      )
    `;
    
    await pool.execute(createTableQuery);
    console.log('Tabel casbin_rule berhasil dibuat atau sudah ada');
  } catch (error) {
    console.error('Error membuat tabel:', error);
  }
}

module.exports = { pool, createCasbinTable };
