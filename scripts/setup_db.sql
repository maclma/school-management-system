-- Create database
CREATE DATABASE school_management;

-- Create user (optional)
CREATE USER school_admin WITH PASSWORD 'school_password';

-- Grant privileges
GRANT ALL PRIVILEGES ON DATABASE school_management TO school_admin;

-- Connect to database and create extensions
\c school_management;

-- Create user roles enum type
CREATE TYPE user_role AS ENUM ('admin', 'teacher', 'student', 'parent');

-- Create tables (these will be auto-migrated, but here's the SQL for reference)
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    role user_role NOT NULL,
    date_of_birth DATE,
    address TEXT,
    profile_image VARCHAR(255),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);