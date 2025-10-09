


-- Billing Codes
CREATE TABLE IF NOT EXISTS billing_codes (
  id   BIGSERIAL PRIMARY KEY,
  salary DECIMAL NOT NULL
);

-- Users
CREATE TABLE IF NOT EXISTS users (
  id               BIGSERIAL PRIMARY KEY,
  email            TEXT NOT NULL UNIQUE,
  first_name       TEXT NOT NULL,
  last_name        TEXT NOT NULL,
  mobile           TEXT NOT NULL,
  password         TEXT NOT NULL,
  billing_code_id  BIGINT REFERENCES billing_codes(id) ON DELETE CASCADE
);

-- Activities
CREATE TABLE IF NOT EXISTS activities (
  id   BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS time_entries (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
  activity_id BIGINT REFERENCES activities(id) ON DELETE CASCADE,
  date DATE  DEFAULT now(),
  start_time TIME NOT NULL,
  end_time TIME NOT NULL,
  total_hours float
);

CREATE TABLE IF NOT EXISTS user_month_hours (
  user_id BIGINT NOT NULL,
  year SMALLINT NOT NULL,
  month SMALLINT NOT NULL,
  total_hours DECIMAL NOT NULL,
  PRIMARY KEY (user_id, year, month)
);

