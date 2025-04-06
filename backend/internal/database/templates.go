package database

const (
	insertIntoUsers = `INSERT INTO users (role, first_name, last_name, email, father_name, password, passwordsalt, date_registration)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)  RETURNING usersid;
`

	SelectUserByEmail = `
SELECT usersid, role, first_name, last_name, email, father_name, password, passwordsalt, date_registration
FROM users
WHERE users.email = $1;
`

	SelectAllAvailableTests = `SELECT tests.testsid, tests.test_name, tests.date_start, tests.date_end
FROM tests;
`

	InsertIntoResult = `INSERT INTO result (time_start, time_end, sum_grade, max_grade, usersid, testsid)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING resultid;
`
	InsertIntoTaskResult = `INSERT INTO taskresult (task_type, usersid, grade)
VALUES ($1, $2, $3) ON CONFLICT DO NOTHING RETURNING task_type;
`

	SelectResultsByUserID = `
SELECT resultid, time_start, time_end, sum_grade, max_grade, usersid, testsid FROM result WHERE usersid=$1; 
`

	SelectTasksByTestID = `
SELECT taskid, task_name, answer, data, max_grade, description FROM task WHERE testsid=$1;
`

	InsertTest = `
INSERT INTO tests (test_name, description, time, date_start, date_end)
VALUES ($1, $2, $3, $4, $5) RETURNING testsid;
`

	InsertTask = `
INSERT INTO task (testsid, task_name, answer, data, max_grade, description)
Values ($1, $2, $3, $4, $5, $6) RETURNING taskid;
`
)
