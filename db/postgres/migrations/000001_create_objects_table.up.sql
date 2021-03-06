CREATE TABLE IF NOT EXISTS objects(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    object_name VARCHAR(1024) NOT NULL,
    object_extension VARCHAR(64) NOT NULL,
    object_path VARCHAR(4096) NOT NULL,
    bucket_name VARCHAR(64) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);