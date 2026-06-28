-- 000002_create_table.up.sql

ALTER TABLE users RENAME COLUMN password_hash TO password;
ALTER TABLE users RENAME COLUMN display_name TO full_name;
