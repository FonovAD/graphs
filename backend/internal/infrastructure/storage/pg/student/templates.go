package storage

const (
	insertIntoTaskResult = `INSERT INTO taskresult (task_type, usersid, grade)
	VALUES ($1, $2, $3) ON CONFLICT DO NOTHING RETURNING task_type;
	`

	selectTasksByUserID = `
	SELECT t.task_id, t.payload, ua.score
	FROM tasks t
	JOIN modules m ON t.module_id = m.module_id
	join user_task ut on t.task_id = ut.task_id
	JOIN user_lab ul ON ut.user_lab_id  = ul.user_lab_id
	JOIN users u ON ul.user_id = u.usersid
	LEFT JOIN user_answer ua ON t.task_id = ua.task_id
	WHERE u.usersid = $1 AND m.module_id = $2;
	`

	selectStudent = `
	SELECT s.student_id
	FROM users u
	INNER JOIN students s ON u.usersid = s.usersid
	WHERE u.usersid = :usersid;
	`
)
