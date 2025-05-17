package storage

const (
	insertIntoUsers = `INSERT INTO users (role, first_name, last_name, email, father_name, password, passwordsalt, date_registration)
	VALUES (:role, :first_name, :last_name, :email, :father_name, :password, :passwordsalt, :date_registration)  RETURNING usersid;
	`
)
