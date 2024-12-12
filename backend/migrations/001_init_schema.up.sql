CREATE TABLE IF NOT EXISTS hackathons (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS profiles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    academic_group VARCHAR(50),
    telegram_handle VARCHAR(50),
    desired_role VARCHAR(100),
    skills TEXT[],
    about_me TEXT,
    achievements TEXT[],
    status VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS hackathon_profiles (
    hackathon_id INTEGER REFERENCES hackathons(id),
    profile_id INTEGER REFERENCES profiles(id),
    PRIMARY KEY (hackathon_id, profile_id)
); 