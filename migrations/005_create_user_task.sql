CREATE TABLE IF NOT EXISTS user_task (
    user_task_id SERIAL PRIMARY KEY,
    user_lab_id INTEGER NOT NULL REFERENCES user_lab(user_lab_id),
    task_id INTEGER NOT NULL REFERENCES tasks(task_id),
    assigned_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (user_lab_id, task_id) 
);
