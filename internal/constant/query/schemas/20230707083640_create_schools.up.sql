CREATE TABLE schools (
    id UUID DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    logo VARCHAR,
    status status DEFAULT 'PENDING',
    created_at VARCHAR,
    updated_at VARCHAR,
    deleted_at VARCHAR,
    PRIMARY KEY (id)
);