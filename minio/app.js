// app.js
const express = require('express');
const multer = require('multer');
const { Client } = require('minio');
const path = require('path');
const fs = require('fs');

const app = express();
const port = 3000;

// Setup Multer (local temp storage)
const upload = multer({ dest: 'uploads/' });

// Setup MinIO Client
const minioClient = new Client({
  endPoint: 'localhost',
  port: 9000,
  useSSL: false,
  accessKey: 'minioadmin',
  secretKey: 'minioadmin123',
});

const BUCKET_NAME = 'uploads';

// Pastikan bucket sudah ada
async function ensureBucket() {
  const exists = await minioClient.bucketExists(BUCKET_NAME);
  if (!exists) {
    await minioClient.makeBucket(BUCKET_NAME);
    console.log(`Bucket "${BUCKET_NAME}" created`);
  } else {
    console.log(`Bucket "${BUCKET_NAME}" exists`);
  }
}
ensureBucket();

// Upload endpoint
app.post('/upload', upload.single('file'), async (req, res) => {
  const filePath = req.file.path;
  const fileName = req.file.originalname;

  try {
    await minioClient.fPutObject(BUCKET_NAME, fileName, filePath, {
      'Content-Type': req.file.mimetype,
    });
    fs.unlinkSync(filePath); // optional: hapus file temp

    res.json({ message: 'File uploaded to MinIO', file: fileName });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
});

app.listen(port, () => {
  console.log(`Server running at http://localhost:${port}`);
});
