ALTER TABLE module_lab
ADD CONSTRAINT uq_module_lab UNIQUE (lab_id, module_id);