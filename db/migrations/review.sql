CREATE TABLE reviews (
    id BINARY(16) NOT NULL,
    job_id BINARY(16) NOT NULL,
    reviewer_id BINARY(16) NOT NULL,
    reviewee_id BINARY(16) NOT NULL,
    rating INT NOT NULL,
    comment TEXT NULL,

    PRIMARY KEY (id),
    CONSTRAINT uk_review_job UNIQUE (job_id),
    FOREIGN KEY (job_id) REFERENCES jobs(id),
    FOREIGN KEY (reviewer_id) REFERENCES users(id),
    FOREIGN KEY (reviewee_id) REFERENCES users(id)
);