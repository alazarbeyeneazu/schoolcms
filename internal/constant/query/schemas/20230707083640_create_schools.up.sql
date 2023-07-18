CREATE TABLE schools (
    id UUID DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    logo VARCHAR,
    phone VARCHAR NOT NULL UNIQUE,
    status status DEFAULT 'PENDING',
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);