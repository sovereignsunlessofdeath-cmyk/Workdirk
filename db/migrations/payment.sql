CREATE TABLE IF NOT EXISTS payments (
    id BINARY(16) NOT NULL,
    job_id BINARY(16) NOT NULL,
    customer_id BINARY(16) NOT NULL,
    worker_id BINARY(16) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) DEFAULT 'Held_Escrow',

    PRIMARY KEY (id),
    CONSTRAINT uk_payment_job UNIQUE (job_id),
    FOREIGN KEY (job_id) REFERENCES jobs(id),
    FOREIGN KEY (customer_id) REFERENCES users(id),
    FOREIGN KEY (worker_id) REFERENCES users(id)
);