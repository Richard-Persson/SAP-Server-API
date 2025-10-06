
-- Users
CREATE TABLE IF NOT EXISTS users (
  id            BIGSERIAL PRIMARY KEY,
  full_name     TEXT NOT NULL
);

-- Activities
CREATE TABLE IF NOT EXISTS activities (
  id   BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  code TEXT NOT NULL
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

-- TimeEntries
CREATE TABLE IF NOT EXISTS time_entries (
  id               BIGSERIAL PRIMARY KEY,
  user_id          BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  day_id           BIGINT NOT NULL REFERENCES days(id) ON DELETE CASCADE,
  date             DATE NOT NULL,
  activity_id      BIGINT NOT NULL REFERENCES activities(id) ON DELETE SET NULL,
  start_ts         TIMESTAMPTZ,
  end_ts           TIMESTAMPTZ,
  duration_minutes INT,
  project_code     TEXT,
  entry_status     INT,
  created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
  modified_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

