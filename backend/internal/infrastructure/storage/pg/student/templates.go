package pg

const (
	InsertIntoTaskResult = `INSERT INTO taskresult (task_type, usersid, grade)
	VALUES ($1, $2, $3) ON CONFLICT DO NOTHING RETURNING task_type;
	`
)
