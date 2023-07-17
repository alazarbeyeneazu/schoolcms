CREATE TABLE families
(
  id UUID DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL,
  family_type VARCHAR NOT NULL,
  status status NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);

ALTER TABLE families 
ADD CONSTRAINT FK_user_TO_familes
FOREIGN KEY (user_id)
REFERENCES users (id);