CREATE TABLE family_to_students
(
  id UUID DEFAULT gen_random_uuid(),
  student_id UUID NOT NULL,
  family_id UUID NOT NULL,
  family_type VARCHAR NOT NULL,
  status status NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);

ALTER TABLE family_to_students 
ADD CONSTRAINT FK_student_To_family_to_students
FOREIGN KEY (student_id)
REFERENCES students (id);

ALTER TABLE family_to_students 
ADD CONSTRAINT FK_family_To_family_to_students
FOREIGN KEY (family_id)
REFERENCES families (id);