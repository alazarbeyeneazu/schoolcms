CREATE TABLE school_teachers
(
   id UUID DEFAULT gen_random_uuid(),
  school_id  UUID    NOT NULL,
  teacher_id UUID    NOT NULL,
  subject    VARCHAR   NOT NULL,
  status status DEFAULT 'PENDING' NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id)
);