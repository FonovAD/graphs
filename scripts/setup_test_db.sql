-- DROP TABLE IF EXISTS grade CASCADE;
-- DROP TABLE IF EXISTS task CASCADE;
-- DROP TABLE IF EXISTS result CASCADE;
-- DROP TABLE IF EXISTS tests CASCADE;
-- DROP TABLE IF EXISTS users CASCADE;

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

CREATE TABLE IF NOT EXISTS tests
(
    testsid SERIAL NOT NULL,
    test_name text NOT NULL,
    description text,
    date_start DATE NOT NULL,
    time INTERVAL HOUR TO SECOND NOT NULL,
    date_end DATE NOT NULL,
    PRIMARY KEY (testsid)
);

CREATE TABLE IF NOT EXISTS result
(
    time_start TIMESTAMP NOT NULL,
    time_end  TIMESTAMP NOT NULL,
    resultid SERIAL NOT NULL,
    sum_grade INT NOT NULL,
    max_grade INT NOT NULL,
    usersid INT NOT NULL,
    testsid INT NOT NULL,
    PRIMARY KEY (resultid),
    FOREIGN KEY (usersid) REFERENCES users(usersid),
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


-- INSERT INTO tests (test_name, description, time, date_start, date_end)
-- VALUES ('Графы. Раздел 1', 'Тест по первому разделу', '10h', '03/03/2014', '03/03/2024');
-- -- INSERT INTO tests (test_name, description) VALUES ('Графы. Раздел 2', 'Тест по второму разделу');
-- -- INSERT INTO tests (test_name, description) VALUES ('Графы. Раздел 3', 'Тест по третьему разделу');
--
-- INSERT INTO task (testsid, task_name, answer, data, max_grade, description)
-- Values (1, 'Задание 1', '1', 'Тут картиночка графа 1', 10, 'Сколько графов на картинке?');
-- INSERT INTO task (testsid, task_name, answer, data, max_grade, description)
-- Values (1, 'Задание 2', '2', 'Тут картиночка графа 2', 20, 'Сколько ребер у графа?');
-- INSERT INTO task (testsid, task_name, answer, data, max_grade, description)
-- Values (1, 'Задание 3', '3', 'Тут картиночка графа 3', 30, 'Сколько вершин у графа?');