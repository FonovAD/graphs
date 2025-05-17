package storage

const (
	insertIntoUsers = `INSERT INTO users (role, first_name, last_name, email, father_name, password, passwordsalt, date_registration)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)  RETURNING usersid;
	`
)
