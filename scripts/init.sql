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

CREATE TABLE IF NOT EXISTS groups
(
    groupsid SERIAL NOT NULL,
    groupsname VARCHAR(10),
    PRIMARY KEY (groupsid)
);

CREATE TABLE IF NOT EXISTS student
(
    studentid SERIAL NOT NULL,
    usersid INT NOT NULL,
    groupsid INT NOT NULL,
    PRIMARY KEY (studentid),
    FOREIGN KEY (usersid) REFERENCES users(usersid),
    FOREIGN KEY (groupsid) REFERENCES groups(groupsid)
);

CREATE TABLE IF NOT EXISTS tests
(
    testsid SERIAL NOT NULL,
    test_name varchar (100) NOT NULL,
    description text,
    date_start DATE NOT NULL,
    time INTERVAL HOUR TO SECOND NOT NULL,
    date_end DATE NOT NULL,
    PRIMARY KEY (testsid)
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

CREATE TABLE IF NOT EXISTS result
(
    time_start TIMESTAMP NOT NULL,
    time_end  TIMESTAMP NOT NULL,
    resultid SERIAL NOT NULL,
    sum_grade INT NOT NULL,
    studentid INT NOT NULL,
    testsid INT NOT NULL,
    PRIMARY KEY (resultid),
    FOREIGN KEY (studentid) REFERENCES student(studentid),
    FOREIGN KEY (testsid) REFERENCES tests(testsid)
);

CREATE TABLE IF NOT EXISTS task
(
    testsid INT NOT NULL,
    taskid SERIAL NOT NULL,
    task_name VARCHAR(100) NOT NULL,
    answer TEXT NOT NULL,
    data TEXT NOT NULL,
    max_grade INT NOT NULL,
    description TEXT NOT NULL,
    PRIMARY KEY (taskid),
    FOREIGN KEY (testsid) REFERENCES tests(testsid)
);

CREATE TABLE IF NOT EXISTS grade
(
    gradeid SERIAL NOT NULL,
    grade INT NOT NULL,
    resultid INT NOT NULL,
    PRIMARY KEY (gradeid),
    FOREIGN KEY (resultid) REFERENCES result(resultid)
);