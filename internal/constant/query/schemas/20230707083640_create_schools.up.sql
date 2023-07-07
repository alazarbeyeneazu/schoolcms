CREATE TABLE schools (
    id UUID DEFAULT gen_random_uuid(),
    name VARCHAR,
    logo VARCHAR,
    created_at VARCHAR,
    updated_at VARCHAR,
    deleted_at VARCHAR,
    PRIMARY KEY (id)
);