


-- Billing Codes
CREATE TABLE IF NOT EXISTS billing_codes (
  id   BIGSERIAL PRIMARY KEY,
  number INT NOT NULL
);

-- Users
CREATE TABLE IF NOT EXISTS users (
  id               BIGSERIAL PRIMARY KEY,
  email            TEXT NOT NULL UNIQUE,
  first_name       TEXT NOT NULL,
  last_name        TEXT NOT NULL,
  mobile           TEXT NOT NULL,
  password         TEXT NOT NULL,
  billing_code_id  BIGINT REFERENCES billing_codes(id)
);

-- Activities
CREATE TABLE IF NOT EXISTS activities (
  id   BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL
);

-- Months
CREATE TABLE IF NOT EXISTS months (
  id          BIGSERIAL PRIMARY KEY,
  year        INT NOT NULL,
  month       INT NOT NULL,
  user_id     BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  total_hours INT NOT NULL DEFAULT 0,
  UNIQUE(user_id, year, month)
);

-- Days
CREATE TABLE IF NOT EXISTS days (
  id          BIGSERIAL PRIMARY KEY,
  date        DATE NOT NULL,
  month_id    BIGINT NOT NULL REFERENCES months(id) ON DELETE CASCADE,
  user_id     BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  total_hours INT NOT NULL DEFAULT 0,
  UNIQUE(date, user_id)
);

-- Time Entries
CREATE TABLE IF NOT EXISTS time_entries (
  id           BIGSERIAL PRIMARY KEY,
  user_id      BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  day_id       BIGINT NOT NULL REFERENCES days(id) ON DELETE CASCADE,
  date         DATE NOT NULL,
  activity_id  BIGINT REFERENCES activities(id) ON DELETE SET NULL
);


