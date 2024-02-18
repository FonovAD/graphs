package database

const (
	initRequest = `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY ,
			email TEXT,
			password TEXT,
			login TEXT
		);
	`

	createUser = `
		INSERT INTO users(email, password, login)
			VALUES ($1, $2, $3) RETURNING id;
`

	authUser = `
		SELECT id, password FROM users 
		WHERE email=$1;
`
)
