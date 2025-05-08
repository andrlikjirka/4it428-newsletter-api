-- Create the schema
CREATE SCHEMA IF NOT EXISTS user_service;

-- Create the users table inside user_service schema
CREATE TABLE IF NOT EXISTS user_service.users (
  id UUID PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  firstname TEXT NOT NULL,
  lastname TEXT NOT NULL,
  firebase_uid TEXT NOT NULL
);
