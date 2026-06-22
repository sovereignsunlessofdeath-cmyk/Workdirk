CREATE TABLE sessions (
    id BINARY(16) NOT NULL,
    user_id BINARY(16) NOT NULL,
    token VARCHAR(255) NOT NULL,
    expires_at DATETIME NOT NULL,

    PRIMARY KEY (id),
    CONSTRAINT uk_session_token UNIQUE (token),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);