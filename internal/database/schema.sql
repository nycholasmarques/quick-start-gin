CREATE TABLE users (
  user_id        BIGSERIAL PRIMARY KEY,
  public_id      UUID NOT NULL DEFAULT gen_random_uuid(),

  first_name     VARCHAR(150) NOT NULL,
  last_name      VARCHAR(150) NOT NULL,

  email          VARCHAR(255) NOT NULL UNIQUE,
  password_hash  VARCHAR(255) NOT NULL,
  phone          VARCHAR(20) NOT NULL UNIQUE,

  is_active      BOOLEAN NOT NULL DEFAULT TRUE,
  email_verified BOOLEAN NOT NULL DEFAULT FALSE,
  phone_verified BOOLEAN NOT NULL DEFAULT FALSE,

  last_login_at  TIMESTAMPTZ,

  created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at     TIMESTAMPTZ
);

CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_public_id ON users (public_id);
CREATE INDEX idx_users_is_active ON users (is_active);
