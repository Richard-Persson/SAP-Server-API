
ROLLBACK;
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

