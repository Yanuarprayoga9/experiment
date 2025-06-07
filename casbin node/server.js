

// server.js
const express = require('express');
const cors = require('cors');
require('dotenv').config();

const { createCasbinTable } = require('./config/database.js');
const { createEnforcer } = require('./config/casbin.js');

const app = express();
const PORT = process.env.PORT || 3000;

// Middleware
app.use(cors());
app.use(express.json());

let enforcer;

// Initialize Casbin enforcer
async function initializeCasbin() {
  try {
    await createCasbinTable();
    enforcer = await createEnforcer();
    await enforcer.loadPolicy();
    console.log('Casbin enforcer berhasil diinisialisasi');
  } catch (error) {
    console.error('Error menginisialisasi Casbin:', error);
  }
}

// Routes

// 1. Tambah Policy
app.post('/api/policies', async (req, res) => {
  try {
    const { subject, object, action } = req.body;
    
    if (!subject || !object || !action) {
      return res.status(400).json({
        success: false,
        message: 'Subject, object, dan action harus diisi'
      });
    }

    const added = await enforcer.addPolicy(subject, object, action);
    await enforcer.savePolicy();

    res.json({
      success: true,
      message: added ? 'Policy berhasil ditambahkan' : 'Policy sudah ada',
      data: { subject, object, action }
    });
  } catch (error) {
    res.status(500).json({
      success: false,
      message: 'Error menambah policy',
      error: error.message
    });
  }
});

// 2. Hapus Policy
app.delete('/api/policies', async (req, res) => {
  try {
    const { subject, object, action } = req.body;
    
    if (!subject || !object || !action) {
      return res.status(400).json({
        success: false,
        message: 'Subject, object, dan action harus diisi'
      });
    }

    const removed = await enforcer.removePolicy(subject, object, action);
    await enforcer.savePolicy();

    res.json({
      success: true,
      message: removed ? 'Policy berhasil dihapus' : 'Policy tidak ditemukan',
      data: { subject, object, action }
    });
  } catch (error) {
    res.status(500).json({
      success: false,
      message: 'Error menghapus policy',
      error: error.message
    });
  }
});

// 3. Lihat semua Policy
app.get('/api/policies', async (req, res) => {
  try {
    const policies = await enforcer.getPolicy();
    const groupingPolicies = await enforcer.getGroupingPolicy();

    res.json({
      success: true,
      data: {
        policies: policies.map(p => ({
          subject: p[0],
          object: p[1],
          action: p[2]
        })),
        roles: groupingPolicies.map(g => ({
          user: g[0],
          role: g[1]
        }))
      }
    });
  } catch (error) {
    res.status(500).json({
      success: false,
      message: 'Error mengambil policies',
      error: error.message
    });
  }
});

// 4. Tambah Role untuk User
app.post('/api/roles', async (req, res) => {
  try {
    const { user, role } = req.body;
    
    if (!user || !role) {
      return res.status(400).json({
        success: false,
        message: 'User dan role harus diisi'
      });
    }

    const added = await enforcer.addRoleForUser(user, role);
    await enforcer.savePolicy();

    res.json({
      success: true,
      message: added ? 'Role berhasil ditambahkan untuk user' : 'Role sudah ada untuk user',
      data: { user, role }
    });
  } catch (error) {
    res.status(500).json({
      success: false,
      message: 'Error menambah role',
      error: error.message
    });
  }
});

// 5. Hapus Role dari User
app.delete('/api/roles', async (req, res) => {
  try {
    const { user, role } = req.body;
    
    if (!user || !role) {
      return res.status(400).json({
        success: false,
        message: 'User dan role harus diisi'
      });
    }

    const removed = await enforcer.deleteRoleForUser(user, role);
    await enforcer.savePolicy();

    res.json({
      success: true,
      message: removed ? 'Role berhasil dihapus dari user' : 'Role tidak ditemukan untuk user',
      data: { user, role }
    });
  } catch (error) {
    res.status(500).json({
      success: false,
      message: 'Error menghapus role',
      error: error.message
    });
  }
});

// 6. Check Permission (Enforce)
app.post('/api/check', async (req, res) => {
  try {
    const { subject, object, action } = req.body;
    
    if (!subject || !object || !action) {
      return res.status(400).json({
        success: false,
        message: 'Subject, object, dan action harus diisi'
      });
    }

    const allowed = await enforcer.enforce(subject, object, action);

    res.json({
      success: true,
      allowed,
      message: allowed ? 'Akses diizinkan' : 'Akses ditolak',
      data: { subject, object, action }
    });
  } catch (error) {
    res.status(500).json({
      success: false,
      message: 'Error memeriksa permission',
      error: error.message
    });
  }
});

// 7. Get Roles untuk User
app.get('/api/roles/:user', async (req, res) => {
  try {
    const { user } = req.params;
    const roles = await enforcer.getRolesForUser(user);

    res.json({
      success: true,
      data: {
        user,
        roles
      }
    });
  } catch (error) {
    res.status(500).json({
      success: false,
      message: 'Error mengambil roles user',
      error: error.message
    });
  }
});

// 8. Get Users dengan Role tertentu
app.get('/api/users/:role', async (req, res) => {
  try {
    const { role } = req.params;
    const users = await enforcer.getUsersForRole(role);

    res.json({
      success: true,
      data: {
        role,
        users
      }
    });
  } catch (error) {
    res.status(500).json({
      success: false,
      message: 'Error mengambil users untuk role',
      error: error.message
    });
  }
});

// Health check
app.get('/health', (req, res) => {
  res.json({
    success: true,
    message: 'API berjalan dengan baik',
    timestamp: new Date().toISOString()
  });
});

// Start server
app.listen(PORT, async () => {
  console.log(`Server berjalan di port ${PORT}`);
  await initializeCasbin();
});
