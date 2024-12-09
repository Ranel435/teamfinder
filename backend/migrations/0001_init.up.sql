CREATE EXTENSION IF NOT EXISTS dblink;

-- Create our application role if it doesn't exist
-- DO
-- $function$
-- BEGIN
--     IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'teamfinder_user') THEN
--         CREATE ROLE teamfinder_user WITH LOGIN PASSWORD 'teamfinder_password';
--     END IF;
-- END;
-- $function$;

-- Create database if it doesn't exist
-- DO
-- $function$
-- BEGIN
--     IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'teamfinder_db') THEN
--         PERFORM dblink_exec('dbname=' || current_database(), 'CREATE DATABASE teamfinder_db');
--     END IF;
-- END;
-- $function$;

-- GRANT ALL PRIVILEGES ON DATABASE teamfinder_db TO teamfinder_user;

-- Create tables
-- Таблица пользователей
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    telegram_id VARCHAR(50) UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Таблица профилей пользователей
CREATE TABLE IF NOT EXISTS profiles (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100),
    academic_group VARCHAR(50),
    telegram_handle VARCHAR(50),
    desired_role VARCHAR(50),
    skills TEXT[] DEFAULT '{}',
    about_me VARCHAR(300),
    achievements TEXT[] DEFAULT '{}',
    status VARCHAR(20) CHECK (status IN ('active', 'inactive', 'pending')) DEFAULT 'active',
    wishlist TEXT[] DEFAULT '{}',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Таблица команд
CREATE TABLE IF NOT EXISTS teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    captain_id INT NOT NULL REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE (name, captain_id)
);

-- Таблица участников команд
CREATE TABLE IF NOT EXISTS team_members (
    id SERIAL PRIMARY KEY,
    team_id INT NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role VARCHAR(50),
    joined_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Таблица уведомлений
CREATE TABLE IF NOT EXISTS notifications (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type VARCHAR(50) CHECK (type IN ('info', 'warning', 'error')),
    message TEXT,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

