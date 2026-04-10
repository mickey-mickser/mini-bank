DROP INDEX IF EXISTS idx_accounts_user_id;
DROP INDEX IF EXISTS uq_users_login_lower;

DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS account_status;
DROP TYPE IF EXISTS user_status;
DROP TYPE IF EXISTS user_role;

-- DROP EXTENSION IF EXISTS pgcrypto;