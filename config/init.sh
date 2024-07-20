#!/bin/bash
set -e

if ! psql -U "$POSTGRES_USER" -d postgres -c "SELECT 1 FROM pg_database WHERE datname = 'rinha-db'" | grep -q 1; then
  psql -U "$POSTGRES_USER" -d postgres -c "CREATE DATABASE \"rinha-db\""
fi

if ! psql -U "$POSTGRES_USER" -d postgres -c "SELECT 1 FROM pg_database WHERE datname = 'test-rinha-db'" | grep -q 1; then
  psql -U "$POSTGRES_USER" -d postgres -c "CREATE DATABASE \"test-rinha-db\""
fi
