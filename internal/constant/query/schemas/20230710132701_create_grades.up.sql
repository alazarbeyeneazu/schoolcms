CREATE TABLE grades
(
  id UUID DEFAULT gen_random_uuid(),
  name VARCHAR NOT NULL,
  school_id UUID NOT NULL,
  status status NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);

ALTER TABLE grades
  ADD CONSTRAINT FK_school_To_grade
    FOREIGN KEY (school_id)
    REFERENCES schools (id);