package storage

const (
	insertIntoTaskResult = `INSERT INTO taskresult (task_type, usersid, grade)
	VALUES ($1, $2, $3) ON CONFLICT DO NOTHING RETURNING task_type;
	`

	selectTasksByUserID = `
	SELECT t.task_id, t.payload
	FROM tasks t
	JOIN modules m ON t.module_id = m.module_id
	JOIN user_lab ul ON t.user_lab_id  = ul.user_lab_id
	JOIN users u ON ul.user_id = u.usersid
	WHERE u.usersid = $1 AND m.module_id = $2;
	`

	selectStudent = `
	SELECT s.student_id
	FROM users u
	INNER JOIN students s ON u.usersid = s.usersid
	WHERE u.usersid = :usersid;
	`
)
