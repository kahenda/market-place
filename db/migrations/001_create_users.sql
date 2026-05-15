CREATE TABLE IF NOT EXISTS users (
  id            SERIAL PRIMARY KEY,
  name          VARCHAR(100) NOT NULL,
  email         VARCHAR(150) UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  location      VARCHAR(100),
  avatar_url    TEXT,
  created_at    TIMESTAMP DEFAULT NOW()
);
