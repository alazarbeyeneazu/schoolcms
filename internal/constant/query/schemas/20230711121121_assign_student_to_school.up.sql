CREATE TABLE school_students
(
  id UUID DEFAULT gen_random_uuid(),
  user_id    UUID   NOT NULL ,
  school_id UUID NOT NULL,
  grade_id UUID NOT NULL,
  status status NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);

ALTER TABLE school_students
  ADD CONSTRAINT FK_user_To_school_students
    FOREIGN KEY (user_id)
    REFERENCES users (id);
ALTER TABLE school_students
  ADD CONSTRAINT FK_school_To_school_students
    FOREIGN KEY (school_id)
    REFERENCES schools (id);
ALTER TABLE school_students
  ADD CONSTRAINT FK_grade_To_school_students
    FOREIGN KEY (grade_id)
    REFERENCES grades (id);