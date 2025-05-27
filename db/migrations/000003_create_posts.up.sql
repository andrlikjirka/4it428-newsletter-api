-- Create the schema
CREATE SCHEMA IF NOT EXISTS newsletter_service;

-- Create table
CREATE TABLE IF NOT EXISTS newsletter_service.posts (
    id UUID PRIMARY KEY,
    newsletter_id UUID REFERENCES newsletter_service.newsletters(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    html_content TEXT NOT NULL,
    published BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ NULL
);