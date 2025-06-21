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

	selectModuleTypeByTask = `
	SELECT CONCAT(m."type", ' ', COALESCE(t.subtype, '')) AS task_type
	FROM user_task ut 
	JOIN tasks t ON ut.task_id = t.task_id
	JOIN modules m ON m.module_id = t.module_id
	JOIN user_lab ul ON ut.user_lab_id = ul.user_lab_id 
	JOIN users u ON u.usersid = ul.user_id
	WHERE ut.task_id = :task_id AND ul.user_id = :user_id;
	`

	selectModuleTypeByLab = `
	select CONCAT(m."type", ' ', COALESCE(t.subtype, '')) AS task_type
	from user_lab ul 
	join user_task ut on ul.user_lab_id = ut.user_lab_id
	join tasks t on ut.task_id = t.task_id
	join modules m on t.module_id = m.module_id
	where ul.user_id = :user_id and ul.lab_id = :lab_id;
	`

	selectScore = `
	select ua.score
	from user_lab ul
	join user_answer ua on ul.user_lab_id = ua.user_lab_id 
	where ul.lab_id = :lab_id and ul.user_id = :user_id;
	`
)
