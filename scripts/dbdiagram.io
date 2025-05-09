Table users {
  usersid SERIAL [pk]
  role VARCHAR(100) [not null]
  first_name VARCHAR(100) [not null]
  last_name VARCHAR(100) [not null]
  email VARCHAR(100) [not null, unique, note: 'Format: %@%.%']
  father_name VARCHAR(100) [not null]
  password VARCHAR(100) [not null]
  passwordsalt VARCHAR(6) [not null]
  date_registration DATE [not null]
}

Table teacher {
  teacherid SERIAL [pk]
  usersid INT [not null, ref: > users.usersid]
}

Table groups {
  groups_id SERIAL [pk] [not null]
  groupsname VARCHAR(10)
}

Table students {
  student_id SERIAL [pk]
  usersid INT [not null, ref: > users.usersid]
  groupsid INT [not null, ref: > groups.groups_id]
}

Table admins {
  admin_id SERIAL [pk]
  usersid INT [not null, ref: > users.usersid]
}

Table labs {
  lab_id SERIAL [pk]
  name VARCHAR(100) [not null]
  description TEXT
  duration INTERVAL [not null]
  registration_date DATE [not null]
  teacher_id INT [not null, ref: > teacher.teacherid]
}

Table user_lab {
  user_lab_id SERIAL [pk]
  user_id INT [not null, ref: > users.usersid]
  lab_id INT [not null, ref: > labs.lab_id]
  assignment_date DATE [not null]
  start_time TIMESTAMP [not null]
  teacher_id INT [not null, ref: > teacher.teacherid]
  deadline TIMESTAMP [not null]
  score INT [not null]
}

Table modules {
  module_id SERIAL [pk]
  type VARCHAR(100) [not null]
  description TEXT
}

Table module_lab {
  module_lab_id SERIAL [pk]
  weight NUMERIC(3, 2) [not null]
  lab_id INT [not null, ref: > labs.lab_id]
  module_id INT [not null, ref: > modules.module_id]
}

Table tasks {
  task_id SERIAL [pk]
  user_lab_id INT [not null, ref: > user_lab.user_lab_id]
  module_id INT [not null, ref: > modules.module_id]
  payload TEXT [not null]
}

Table user_answer {
  user_answer_id SERIAL [pk]
  user_lab_id INT [not null, ref: > user_lab.user_lab_id]
  task_id INT [not null, ref: > tasks.task_id]
  answer TEXT [not null]
  score INT [not null]
}

Table teachergroup {
  teachergroupid SERIAL [pk]
  teacherid INT [not null, ref: > teacher.teacherid]
  groupsid INT [not null, ref: > groups.groups_id]
}