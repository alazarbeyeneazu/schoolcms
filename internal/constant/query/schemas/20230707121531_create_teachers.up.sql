CREATE TABLE teachers (
    id UUID DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    title VARCHAR NOT NULL,
    status status DEFAULT 'PENDING' NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

ALTER TABLE teachers
  ADD CONSTRAINT FK_users_TO_teachers
    FOREIGN KEY (user_id)
    REFERENCES users (id);