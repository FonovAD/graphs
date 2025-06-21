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
	-- Доступные таски для лабы
	available_tasks AS (
		SELECT 
			t.task_id,
			ROW_NUMBER() OVER (ORDER BY t.task_id) - 1 AS task_num  -- Нумерация с 0
		FROM tasks t
		JOIN module_lab ml ON t.module_id = ml.module_id
		WHERE ml.lab_id = :lab_id
	),
	-- Нумерованные студенты
	numbered_students AS (
		SELECT 
			user_lab_id,
			user_id,
			ROW_NUMBER() OVER (ORDER BY user_id) - 1 AS student_num  -- Нумерация с 0
		FROM inserted_user_labs
	),
	-- Общее кол-во вариков
	task_count AS (
		SELECT COUNT(*) AS total FROM available_tasks
	)
	-- Назначение по кругу
	INSERT INTO user_task (user_lab_id, task_id)
	SELECT 
		ns.user_lab_id,
		at.task_id
	FROM numbered_students ns
	JOIN available_tasks at ON at.task_num = (ns.student_num % (SELECT total FROM task_count))
	RETURNING :lab_id AS lab_id;
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

	getLabsResults = `
	WITH student_lab_scores AS (
		SELECT 
			ul.lab_id,
			l.name AS lab_name,  -- Добавляем название лабораторной работы
			s.usersid AS user_id,
			CONCAT(u.last_name, ' ', u.first_name, ' ', u.father_name) AS fio,
			ul.score AS overall_score,
			m.module_id,
			m.type AS module_name,
			COALESCE(ua.score, 0) AS module_score
		FROM user_lab ul
		JOIN labs l ON ul.lab_id = l.lab_id  -- Добавляем соединение с таблицей labs
		JOIN students s ON ul.user_id = s.usersid
		JOIN users u ON s.usersid = u.usersid
		join user_task ut on ul.user_lab_id = ut.user_lab_id
		JOIN tasks t ON ut.task_id = t.task_id
		JOIN modules m ON t.module_id = m.module_id
		LEFT JOIN user_answer ua ON t.task_id = ua.task_id AND ul.user_lab_id = ua.user_lab_id
		WHERE s.groupsid = $1
	)

	SELECT 
		lab_id,
		MAX(lab_name) AS lab_name,  -- Добавляем название лабораторной работы в результат
		json_agg(
			json_build_object(
				'user_id', user_id,
				'fio', fio,
				'overall_score', overall_score,
				'module_results', (
					SELECT json_agg(
						json_build_object(
							'module_id', module_id,
							'module_name', module_name,
							'module_score', module_score
						)
					)
					FROM student_lab_scores s2
					WHERE s2.user_id = s1.user_id AND s2.lab_id = s1.lab_id
				)
			)
		) AS students
	FROM (
		SELECT DISTINCT lab_id, lab_name, user_id, fio, overall_score
		FROM student_lab_scores
	) s1
	GROUP BY lab_id
	ORDER BY lab_id;
	`
)
