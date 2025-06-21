ALTER TABLE user_lab
ADD CONSTRAINT uq_user_lab_user UNIQUE (lab_id, user_id);