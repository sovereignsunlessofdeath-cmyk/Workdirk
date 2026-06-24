CREATE TABLE IF NOT EXISTS jobs (
    id BINARY(16) NOT NULL,
    customer_id BINARY(16) NOT NULL,
    worker_id BINARY(16) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(50) DEFAULT 'Pending',
    PRIMARY KEY (id),
    FOREIGN KEY (customer_id) REFERENCES users (id),
    FOREIGN KEY (worker_id) REFERENCES users (id)
);