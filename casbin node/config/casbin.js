// config/casbin.js
const casbin = require('casbin');
const { pool } = require('./database');

// Custom MySQL Adapter untuk Casbin
class MySQLAdapter {
  constructor(pool) {
    this.pool = pool;
  }

  async loadPolicy(model) {
    const [rows] = await this.pool.execute('SELECT * FROM casbin_rule');
    
    for (const row of rows) {
      const line = [row.ptype];
      if (row.v0) line.push(row.v0);
      if (row.v1) line.push(row.v1);
      if (row.v2) line.push(row.v2);
      if (row.v3) line.push(row.v3);
      if (row.v4) line.push(row.v4);
      if (row.v5) line.push(row.v5);
      
      model.addPolicy('p', row.ptype, line.slice(1));
    }
  }

  async savePolicy(model) {
    await this.pool.execute('DELETE FROM casbin_rule');
    
    const policies = [];
    
    // Get all policies
    const policyMap = model.model.get('p');
    if (policyMap) {
      for (const [key, ast] of policyMap) {
        for (const rule of ast.policy) {
          policies.push(['p', key, ...rule]);
        }
      }
    }

    // Get all role inheritances
    const roleMap = model.model.get('g');
    if (roleMap) {
      for (const [key, ast] of roleMap) {
        for (const rule of ast.policy) {
          policies.push(['g', key, ...rule]);
        }
      }
    }

    // Insert policies
    for (const policy of policies) {
      const [ptype, sec, ...params] = policy;
      await this.pool.execute(
        'INSERT INTO casbin_rule (ptype, v0, v1, v2, v3, v4, v5) VALUES (?, ?, ?, ?, ?, ?, ?)',
        [
          sec,
          params[0] || null,
          params[1] || null,
          params[2] || null,
          params[3] || null,
          params[4] || null,
          params[5] || null
        ]
      );
    }
  }

  async addPolicy(sec, ptype, rule) {
    await this.pool.execute(
      'INSERT INTO casbin_rule (ptype, v0, v1, v2, v3, v4, v5) VALUES (?, ?, ?, ?, ?, ?, ?)',
      [
        ptype,
        rule[0] || null,
        rule[1] || null,
        rule[2] || null,
        rule[3] || null,
        rule[4] || null,
        rule[5] || null
      ]
    );
  }

  async removePolicy(sec, ptype, rule) {
    const conditions = ['ptype = ?'];
    const params = [ptype];
    
    rule.forEach((param, index) => {
      if (param) {
        conditions.push(`v${index} = ?`);
        params.push(param);
      }
    });

    const query = `DELETE FROM casbin_rule WHERE ${conditions.join(' AND ')}`;
    await this.pool.execute(query, params);
  }

  async removeFilteredPolicy(sec, ptype, fieldIndex, ...fieldValues) {
    const conditions = ['ptype = ?'];
    const params = [ptype];
    
    fieldValues.forEach((value, index) => {
      if (value) {
        conditions.push(`v${fieldIndex + index} = ?`);
        params.push(value);
      }
    });

    const query = `DELETE FROM casbin_rule WHERE ${conditions.join(' AND ')}`;
    await this.pool.execute(query, params);
  }
}

// Membuat enforcer dengan model RBAC
async function createEnforcer() {
  const model = casbin.newModel(`
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
  `);

  const adapter = new MySQLAdapter(pool);
  const enforcer = await casbin.newEnforcer(model, adapter);
  
  return enforcer;
}

module.exports = { createEnforcer, MySQLAdapter };
