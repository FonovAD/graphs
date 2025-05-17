package storage

const (
	SelectUserByEmail = `
	SELECT usersid, role, first_name, last_name, email, father_name, password, passwordsalt, date_registration
	FROM users
	WHERE users.email = $1;
	`
)
