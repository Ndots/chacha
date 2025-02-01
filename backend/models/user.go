package models

import "database/sql"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Role     string `json:"role"`
}

func CreateUser(db *sql.DB, user *User) error {
	query := `INSERT INTO users (name, email, password, address, role) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return db.QueryRow(query, user.Name, user.Email, user.Password, user.Address, user.Role).Scan(&user.ID)
}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	user := &User{}
	query := `SELECT id, name, email, password, address, role FROM users WHERE email = $1`
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Address, &user.Role)
	return user, err
}
