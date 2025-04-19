-- Garante que a extensão pgcrypto está habilitada
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Insere a empresa e pega o ID gerado
WITH inserted_company AS (
  INSERT INTO companies (id, name)
  VALUES (
    gen_random_uuid(),
    'CyberPlace 3D'
  )
  RETURNING id
)
-- Insere o usuário com o ID da empresa
INSERT INTO users (id, company_id, name, email, password_hash, role)
SELECT
  gen_random_uuid(),
  inserted_company.id,
  'Vinícius Guilherme Batista',
  'vinigbatista21@gmail.com',
  crypt('vinii', gen_salt('bf')),
  'admin'
FROM inserted_company;
