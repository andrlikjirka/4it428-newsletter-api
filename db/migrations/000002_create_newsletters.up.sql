-- Create the schema
CREATE SCHEMA IF NOT EXISTS newsletter_service;

-- Create table
CREATE TABLE IF NOT EXISTS newsletter_service.newsletters (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL
);
