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
	WITH inserted_user_labs AS (
    INSERT INTO user_lab (
        user_id, lab_id, assignment_date, start_time, teacher_id, deadline, score
    )
    SELECT 
        s.usersid, :lab_id, :assignment_date, :start_time, :teacher_id, :deadline, NULL                       
    FROM students s
    WHERE s.groupsid = :groups_id
    RETURNING user_lab_id, user_id
	),
	available_tasks AS (
		SELECT 
			task_id,
			ROW_NUMBER() OVER (ORDER BY task_id) as task_row
		FROM tasks
		WHERE module_id IN (
			SELECT module_id FROM module_lab WHERE lab_id = :lab_id
		)
		AND user_lab_id IS NULL
	),
	numbered_students AS (
		SELECT 
			user_lab_id,
			user_id,
			ROW_NUMBER() OVER (ORDER BY user_id) as student_row
		FROM inserted_user_labs
	),
	updated_tasks AS (
		UPDATE tasks t
		SET user_lab_id = ns.user_lab_id
		FROM available_tasks at
		JOIN numbered_students ns ON at.task_row = ns.student_row
		WHERE t.task_id = at.task_id
		RETURNING t.task_id
	)
	SELECT :lab_id AS lab_id;
	`

	selectAvailableTasksCountByModule = `
	with modules_from_lab as (
	select ml.module_id 
	from module_lab ml
	join labs l on l.lab_id = ml.lab_id
	where l.lab_id = $1
	)

	select count (*) as tasks_count
	from tasks t
	join modules_from_lab mfl on mfl.module_id = t.module_id
	where t.user_lab_id is null;
	`

	selectStudentsCountFromGroup = `
	select count (*) as students_count
	from students s
	where s.groupsid  = $1;
	`

	selectNonExistingUserLabs = `
	SELECT l.lab_id, l.name 
	FROM labs l 
	LEFT OUTER JOIN user_lab ul ON l.lab_id = ul.lab_id 
	WHERE ul.lab_id IS NULL
	LIMIT $1 OFFSET $2;
	`

	selectExistingUserLabs = `
	SELECT 
		l.lab_id, 
		l.name as lab_name,
		(
			SELECT json_agg(json_build_object('groupID', g.groups_id, 'groupName', g.groupsname))
			FROM (
				SELECT DISTINCT g.groups_id, g.groupsname
				FROM user_lab ul
				JOIN users u ON ul.user_id = u.usersid
				JOIN students s ON u.usersid = s.usersid
				JOIN "groups" g ON s.groupsid = g.groups_id
				WHERE ul.lab_id = l.lab_id
			) g
		) as groups
	FROM labs l
	WHERE EXISTS (
		SELECT 1 FROM user_lab WHERE lab_id = l.lab_id
	);
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
	INSERT INTO tasks (module_id, payload, answer) 
	VALUES (:module_id, :payload, :answer)
	RETURNING task_id;
	`

	updateTask = `
	UPDATE tasks
	SET
		module_id = :module_id,
		payload = :payload,
		answer = :answer
	WHERE task_id = :task_id
	RETURNING task_id;
	`

	selectTasksByModule = `
	SELECT t.task_id, t.payload
	FROM tasks t
	JOIN modules m ON t.module_id = m.module_id
	WHERE m.module_id = $1;
	`
)
