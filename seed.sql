CREATE TABLE IF NOT EXISTS users (
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL
);

-- CREATE TABLE IF NOT EXISTS quests (
--   id SERIAL PRIMARY KEY,
--   title VARCHAR(255) NOT NULL,
--   description VARCHAR(255) NOT NULL,
--   user_id INT NOT NULL REFERENCES users(id)
-- );