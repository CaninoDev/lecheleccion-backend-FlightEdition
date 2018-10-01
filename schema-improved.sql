CREATE SCHEMA IF NOT EXISTS lecheleccion;

CREATE TABLE IF NOT EXISTS articles(
  id SERIAL PRIMARY KEY,
  url TEXT,
  urlToImage TEXT,
  source TEXT,
  publication_date TIMESTAMP,
  title TEXT,
  body TEXT,
  external_reference_id SMALLINT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS biases(
  id SERIAL PRIMARY KEY,
  libertarian REAL NOT NULL,
  green REAL NOT NULL,
  liberal REAL NOT NULL, 
  conservative REAL NOT NULL,
  biasable_type TEXT,
  biasable_id SMALLINT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users(
  id SERIAL PRIMARY KEY,
  name TEXT
);

CREATE TABLE IF NOT EXISTS votes(
  id SERIAL PRIMARY KEY,
  article_id INTEGER REFERENCES articles,
  user_id INTEGER REFERENCES users
);

CREATE INDEX ON biases (biasable_type, biasable_id);
CREATE INDEX ON votes (article_id, user_id);
