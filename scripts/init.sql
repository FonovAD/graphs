CREATE TABLE IF NOT EXISTS users
(
    usersid SERIAL NOT NULL,
    role VARCHAR(100) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) CHECK (email LIKE '%@%.%') NOT NULL,
    father_name VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    passwordsalt VARCHAR(6) NOT NULL,
    date_registration DATE NOT NULL,
    PRIMARY KEY (usersid),
    UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS teacher
(
    teacherid SERIAL NOT NULL,
    usersid INT NOT NULL,
    PRIMARY KEY (teacherid),
    FOREIGN KEY (usersid) REFERENCES users(usersid)
);

CREATE TABLE IF NOT EXISTS students
(
    student_id SERIAL NOT NULL,
    usersid INT NOT NULL,
    groupsid INT NOT NULL,
    PRIMARY KEY (student_id),
    FOREIGN KEY (usersid) REFERENCES users(usersid),
    FOREIGN KEY (groupsid) REFERENCES groups(groupsid)
);

CREATE TABLE IF NOT EXISTS admins
(
    admin_id SERIAL NOT NULL,
    usersid INT NOT NULL,
    PRIMARY KEY (admin_id),
    FOREIGN KEY (usersid) REFERENCES users(usersid)
);

CREATE TABLE IF NOT EXISTS groups
(
    groups_id SERIAL NOT NULL,
    groupsname VARCHAR(10) NOT NULL,
    PRIMARY KEY (groupsid)
);

CREATE TABLE IF NOT EXISTS labs
(
    lab_id SERIAL NOT NULL,
    name varchar (100) NOT NULL,
    description text,
    duration INTERVAL HOUR TO SECOND NOT NULL,
    registration_date DATE NOT NULL,
    teacher_id INT NOT NULL,
    PRIMARY KEY (lab_id),
    FOREIGN KEY (teacher_id) REFERENCES teacher(teacherid)
);

CREATE TABLE IF NOT EXISTS user_lab
(
    user_lab_id SERIAL NOT NULL,
    user_id INT NOT NULL,
    lab_id INT NOT NULL,
    assignment_date DATE NOT NULL,
    start_time TIMESTAMP NOT NULL,
    teacher_id INT NOT NULL,
    deadline TIMESTAMP NOT NULL,
    score INT NOT NULL,
    PRIMARY KEY (user_lab_id),
    FOREIGN KEY (teacher_id) REFERENCES teacher(teacherid),
    FOREIGN KEY (user_id) REFERENCES users(usersid),
    FOREIGN KEY (lab_id) REFERENCES labs(lab_id)
);

CREATE TABLE IF NOT EXISTS modules
(
    module_id SERIAL NOT NULL,
    type varchar (100) NOT NULL,
    description text,
    PRIMARY KEY (module_id)
);

CREATE TABLE IF NOT EXISTS module_lab
(
    module_lab_id SERIAL NOT NULL,
    weight NUMERIC(3, 2) NOT NULL,
    lab_id INT NOT NULL,
    module_id INT NOT NULL,
    PRIMARY KEY (module_lab_id),
    FOREIGN KEY (lab_id) REFERENCES labs(lab_id),
    FOREIGN KEY (module_id) REFERENCES modules(module_id)
);

CREATE TABLE IF NOT EXISTS tasks
(
    task_id SERIAL NOT NULL,
    user_lab_id INT NOT NULL,
    module_id INT NOT NULL,
    payload TEXT NOT NULL,
    PRIMARY KEY (task_id),
    FOREIGN KEY (user_lab_id) REFERENCES user_lab(user_lab_id),
    FOREIGN KEY (module_id) REFERENCES modules(module_id)
);

CREATE TABLE IF NOT EXISTS user_answer
(
    user_answer_id SERIAL NOT NULL,
    user_lab_id INT NOT NULL,
    task_id INT NOT NULL,
    answer TEXT NOT NULL,
    score INT NOT NULL,
    PRIMARY KEY (user_answer_id),
    FOREIGN KEY (user_lab_id) REFERENCES user_lab(user_lab_id),
    FOREIGN KEY (task_id) REFERENCES tasks(task_id)
);

CREATE TABLE IF NOT EXISTS teachergroup
(
    teachergroupid SERIAL NOT NULL,
    teacherid INT NOT NULL,
    groupsid INT NOT NULL,
    PRIMARY KEY (teachergroupid),
    FOREIGN KEY (teacherid) REFERENCES teacher(teacherid),
    FOREIGN KEY (groupsid) REFERENCES groups(groupsid)
);


