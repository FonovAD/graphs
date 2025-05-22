package storage

const (
	SelectUserByEmail = `
	SELECT usersid, role, first_name, last_name, email, father_name, password, passwordsalt, date_registration
	FROM users
	WHERE users.email = $1;
	`

	insertIntoUsers = `
	INSERT INTO users (role, first_name, last_name, email, father_name, password, passwordsalt, date_registration)
	VALUES (:role, :first_name, :last_name, :email, :father_name, :password, :passwordsalt, :date_registration)  RETURNING usersid;`
)
