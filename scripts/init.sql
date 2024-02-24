CREATE TABLE IF NOT EXISTS users
(
  usersid SERIAL,
  role VARCHAR(100) NOT NULL,
  first_name VARCHAR(100) NOT NULL,
  last_name VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL,
  father_name VARCHAR(100) NOT NULL,
  password VARCHAR(100) NOT NULL,
  passwordsalt VARCHAR(6) NOT NULL,
  date_registration DATE NOT NULL,
  UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS teacher
(
    teacherid INT NOT NULL,
    usersid INT NOT NULL,
    PRIMARY KEY (teacherid),
    FOREIGN KEY (usersid) REFERENCES users(usersid)
);

CREATE TABLE IF NOT EXISTS groups
(
    groupsid SERIAL,
    groupsname VARCHAR(10)
);

CREATE TABLE IF NOT EXISTS student
(
    studentid INT NOT NULL,
    usersid INT NOT NULL,
    groupsid INT NOT NULL,
    PRIMARY KEY (studentid),
    FOREIGN KEY (usersid) REFERENCES users(usersid),
    FOREIGN KEY (groupsid) REFERENCES groups(groupsid)
);

CREATE TABLE IF NOT EXISTS tests
(
    testsid INT NOT NULL,
    test_name varchar (100) NOT NULL,
    date_start DATE NOT NULL,
    time INTERVAL HOUR TO SECOND NOT NULL,
    date_end DATE NOT NULL,
    PRIMARY KEY (testsid)
);

CREATE TABLE IF NOT EXISTS teachergroup
(
    teachergroupid SERIAL,
    teacherid INT NOT NULL,
    groupsid INT NOT NULL,
    FOREIGN KEY (teacherid) REFERENCES teacher(teacherid),
    FOREIGN KEY (groupsid) REFERENCES groups(groupsid)
);

CREATE TABLE IF NOT EXISTS result
(
    time_start TIMESTAMP NOT NULL,
    time_end  TIMESTAMP NOT NULL,
    resultid SERIAL,
    sum_grade INT NOT NULL,
    studentid INT NOT NULL,
    testsid INT NOT NULL,
    FOREIGN KEY (studentid) REFERENCES student(studentid),
    FOREIGN KEY (testsid) REFERENCES tests(testsid)
);

CREATE TABLE IF NOT EXISTS task
(
    taskid SERIAL,
    task_name VARCHAR(100) NOT NULL,
    answer JSONB NOT NULL,
    data JSONB NOT NULL,
    max_grade INT NOT NULL,
    description VARCHAR (1000) NOT NULL
);

CREATE TABLE IF NOT EXISTS test_task
(
    testtaskid SERIAL,
    testsid INT NOT NULL,
    taskid INT NOT NULL,
    FOREIGN KEY (testsid) REFERENCES tests(testsid),
    FOREIGN KEY (taskid) REFERENCES task(taskid)
);

CREATE TABLE grade
(
    gradeid SERIAL,
    grade INT NOT NULL,
    resultid INT NOT NULL,
    FOREIGN KEY (resultid) REFERENCES result(resultid)
);