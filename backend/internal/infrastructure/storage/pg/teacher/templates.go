package storage

const (
	insertIntoUsers = `
	INSERT INTO users (role, first_name, last_name, email, father_name, password, passwordsalt, date_registration)
	VALUES (:role, :first_name, :last_name, :email, :father_name, :password, :passwordsalt, :date_registration)  RETURNING usersid;`

	insertIntoStudent = `
	INSERT INTO students (usersid, groups_id) 
	VALUES (:usersid, :groups_id)  RETURNING student_id;
	`

	selectAllModules = `SELECT module_id, type FROM modules;`

	createLab = `
	INSERT INTO labs(name, description, duration, registration_date, teacher_id) 
	VALUES (:name, :description, :duration, :registration_date, :teacher_id) RETURNING lab_id;
	`

	addModuleToLab = `
	INSERT INTO module_lab(weight, lab_id, module_id) 
	VALUES (:weight, :lab_id, :module_id) ON CONFLICT (lab_id, module_id) DO NOTHING
	RETURNING module_lab_id;
	`

	selectModulesFromLab = `
	SELECT l.lab_id, ml.module_lab_id, ml.weight, m.module_id, m.type
	FROM labs l 
	INNER JOIN module_lab ml ON l.lab_id = ml.lab_id
	INNER JOIN modules m ON ml.module_id = m.module_id
	WHERE l.lab_id = $1;
	`

	removeModuleFromLab = `
	DELETE FROM module_lab 
	WHERE lab_id = :lab_id AND module_id = :module_id
	RETURNING module_lab_id;
	`

	selectLabInfo = `
	SELECT 
    l.lab_id, 
    l.name, 
    l.description, 
    l.duration, 
    l.registration_date, 
    CONCAT_WS(' ', u.first_name, u.last_name, u.father_name) AS teacher_fio
	FROM labs l 
	INNER JOIN teacher t ON t.teacherid = l.teacher_id 
	INNER JOIN users u ON u.usersid = t.usersid
	WHERE l.lab_id = :lab_id;
	`

	removeLabFromUserLab = `
	DELETE FROM user_lab 
	WHERE user_id = :user_id AND lab_id = :lab_id
	RETURNING lab_id;
	`

	updateLabInfo = `
	UPDATE labs 
	SET name = :name, description = :description, duration = :duration
	WHERE lab_id = :lab_id;
	`

	insertUserLab = `
	INSERT INTO user_lab(user_id, lab_id, assignment_date, start_time, teacher_id, deadline) 
	VALUES (:user_id, :lab_id, :assignment_date, :start_time, :teacher_id, :deadline) RETURNING user_lab_id;
	`

	insertLabToStudentGroup = `
	INSERT INTO user_lab (
    user_id, 
    lab_id, 
    assignment_date, 
    start_time, 
    teacher_id, 
    deadline,
    score
	)
	SELECT s.usersid, :lab_id, :assignment_date, :start_time, :teacher_id, :deadline, NULL                       
	FROM students s
	WHERE s.groupsid = :groups_id
	RETURNING lab_id;
	`

	selectNonExistingUserLabs = `
	SELECT l.lab_id, l.name 
	FROM labs l 
	LEFT OUTER JOIN user_lab ul ON l.lab_id = ul.lab_id 
	WHERE ul.lab_id IS NULL
	LIMIT $1 OFFSET $2;
	`

	selectExistingUserLabs = `
	SELECT ul.user_lab_id, l.lab_id, l.name 
	FROM labs l 
	INNER JOIN user_lab ul ON l.lab_id = ul.lab_id 
	LIMIT $1 OFFSET $2;
	`

	selectTeacher = `
	SELECT t.teacherid
	FROM users u
	INNER JOIN teacher t ON u.usersid = t.usersid
	WHERE u.usersid = :usersid;
	`

	selectGroups = `
	SELECT g.groups_id, groupsname
	FROM groups g;
	`

	insertTask = `
	INSERT INTO tasks (task_id, module_id, payload, answer) VALUES (:task_id, :module_id, :payload, :answer) 
	ON CONFLICT (task_id) DO UPDATE
	SET
		module_id = EXCLUDED.module_id,
		payload = EXCLUDED.payload,
		answer = EXCLUDED.answer
	RETURNING task_id;
	`
)
