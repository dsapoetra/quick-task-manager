-- Tasks table with assignee and assigner columns and foreign keys to users table
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(255) NOT NULL,
    assignee_id INT,
    assigner_id INT,
    priority INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_assignee_id ON tasks(assignee_id);
CREATE INDEX idx_assigner_id ON tasks(assigner_id);
CREATE INDEX idx_priority ON tasks(priority);
CREATE INDEX idx_id ON tasks(id);
