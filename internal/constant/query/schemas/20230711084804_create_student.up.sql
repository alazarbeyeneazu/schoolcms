CREATE TABLE students
(
  id UUID DEFAULT gen_random_uuid(),
  user_id    UUID   NOT NULL,
  status status NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);

ALTER TABLE students
  ADD CONSTRAINT FK_user_To_students
    FOREIGN KEY (user_id)
    REFERENCES users (id);