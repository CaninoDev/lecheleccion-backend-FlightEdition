CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE SCHEMA IF NOT EXISTS lecheleccion;

CREATE TABLE IF NOT EXISTS lecheleccion.users(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT
);

CREATE TABLE IF NOT EXISTS lecheleccion.articles(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  url text,
  urlToImage text,
  source text,
  publication_date timestamp,
  title text,
  body text,
  created_at timestamp,
  updated_at timestamp
);

CREATE TABLE IF NOT EXISTS lecheleccion.votes(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  article_id uuid REFERENCES lecheleccion.articles,
  user_id uuid REFERENCES lecheleccion.users
);

CREATE TABLE IF NOT EXISTS lecheleccion.biases(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  conservative real,
  liberal real,
  libertarian real,
  green real,

  article_id UUID REFERENCES lecheleccion.articles,
  user_id UUID REFERENCES lecheleccion.users,
  vote_id UUID REFERENCES lecheleccion.votes,

  check (
    (
      (article_id IS NOT NULL)::integer +
      (user_id IS NOT NULL)::integer +
      (vote_id IS NOT NULL)::integer
    ) = 1
  )
);

CREATE UNIQUE INDEX ON biases (article_id) WHERE document_id IS NOT NULL;
CREATE UNIQUE INDEX ON biases (user_id) WHERE document_id IS NOT NULL;
CREATE UNIQUE INDEX ON biases (vote_id) WHERE document_id IS NOT NULL;
