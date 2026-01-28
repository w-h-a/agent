CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    session_id TEXT,
    role VARCHAR(50) NOT NULL,
    parts JSONB,
    embedding VECTOR(1536),
    created_at TIMESTAMP DEFAULT NOW()
);