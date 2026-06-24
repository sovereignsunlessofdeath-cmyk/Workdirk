CREATE TABLE IF NOT EXISTS users (
    id BINARY(16) NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password_hash BINARY(60) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    role VARCHAR(10) DEFAULT 'User',
    PRIMARY KEY (id),
    CONSTRAINT uk_user_email UNIQUE (email),
    CONSTRAINT uk_user_phone UNIQUE (phone_number)
);