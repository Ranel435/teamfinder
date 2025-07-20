package database

import (
	"context"
	"fmt"
	"log"
)

func RunMigrations() error {
	if Pool == nil {
		return fmt.Errorf("database pool is not initialized")
	}

	migrations := []string{
		createUsersTable,
		createHackathonsTable,
		createProfilesTable,
		createIndexes,
	}

	for i, migration := range migrations {
		log.Printf("Running migration %d...", i+1)
		if _, err := Pool.Exec(context.Background(), migration); err != nil {
			return fmt.Errorf("failed to run migration %d: %w", i+1, err)
		}
	}

	log.Println("All migrations completed successfully")
	return nil
}

const createUsersTable = `
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createHackathonsTable = `
CREATE TABLE IF NOT EXISTS hackathons (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createProfilesTable = `
CREATE TABLE IF NOT EXISTS profiles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    hackathon_id INTEGER REFERENCES hackathons(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    academic_group VARCHAR(100),
    telegram_handle VARCHAR(255),
    desired_role VARCHAR(255),
    skills TEXT[],
    about_me TEXT,
    achievements TEXT[],
    status VARCHAR(50) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, hackathon_id)
);
`

const createIndexes = `
CREATE INDEX IF NOT EXISTS idx_profiles_user_id ON profiles(user_id);
CREATE INDEX IF NOT EXISTS idx_profiles_hackathon_id ON profiles(hackathon_id);
CREATE INDEX IF NOT EXISTS idx_profiles_status ON profiles(status);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
`
