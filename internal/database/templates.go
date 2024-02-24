package database

const (
	initRequest = `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY ,
			email TEXT,
			password TEXT,
			login TEXT
		);
	`

	authUser = `
		SELECT id, password FROM users 
		WHERE email = $1;
`
	insertIntoUsers = `INSERT INTO users (role, first_name, last_name, email, father_name, password, passwordsalt, date_registration)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)  RETURNING usersid;
`

	SelectUserByEmail = `
SELECT role, first_name, last_name, email, father_name, password, passwordsalt, date_registration
FROM users
WHERE users.email = $1;
`

	SelectAllAvailableTests = `SELECT tests.testsid, tests.test_name, tests.date_start, tests.date_end
FROM tests;
`

	SelectAllTasksInTest = `SELECT tests.test_name, tests.time, task.task_name, task.max_grade 
FROM tests
JOIN test_task ON test_task.taskid = tests.testsid
JOIN task ON test_task.taskid = task.taskid
GROUP BY test_name
WHERE tests.testsid = $1;
`

	InsertIntoResult = `INSERT INTO result (time_start, time_end, sum_grade, studentid, testid)
VALUES ($1, $2, $3, $4, $5);
`
)
