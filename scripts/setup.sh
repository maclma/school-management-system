#!/bin/bash

# School Management System Setup Script

set -e

echo "Setting up School Management System..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go 1.21 or higher."
    exit 1
fi

# Check if PostgreSQL is installed
if ! command -v psql &> /dev/null; then
    echo "PostgreSQL is not installed. Please install PostgreSQL."
    exit 1
fi

# Check if .env file exists
if [ ! -f .env ]; then
    echo "Creating .env file from .env.example..."
    cp .env.example .env
    echo "Please update the .env file with your database credentials."
fi

# Load environment variables
source .env

# Create database
echo "Creating database '$DB_NAME'..."
sudo -u postgres createdb "$DB_NAME" 2>/dev/null || true

# Run SQL setup script
echo "Running database setup script..."
psql -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME" -f scripts/setup_db.sql

# Install Go dependencies
echo "Installing Go dependencies..."
go mod download

# Build the application
echo "Building the application..."
go build -o bin/server ./cmd/server

# Create necessary directories
mkdir -p logs
mkdir -p uploads

echo "Setup completed successfully!"
echo ""
echo "To run the application:"
echo "1. Update .env file with your settings"
echo "2. Run: ./bin/server"
echo "3. Or for development with live reload: make dev"
echo ""
echo "Default admin credentials:"
echo "Email: admin@school.com"
echo "Password: admin123"