ALTER TABLE user_answer
ADD CONSTRAINT uq_user_answer_task UNIQUE (user_lab_id, task_id);