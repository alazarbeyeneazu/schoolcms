CREATE TABLE users
(
  id   UUID  DEFAULT gen_random_uuid(), 
  first_name  VARCHAR  ,
  middle_name VARCHAR  ,
  last_name   VARCHAR  ,
  gender VARCHAR NOT NULL,
  phone  VARCHAR  NULL UNIQUE,
  profile     VARCHAR  ,
  status status NOT NULL,
  created_at  TIMESTAMP NOT NULL DEFAULT now(),
  updated_at  TIMESTAMP,
  deleted_at  TIMESTAMP,
  PRIMARY KEY (id)
);
