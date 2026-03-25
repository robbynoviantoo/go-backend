package repository

import (
	"go-backend/config"
	"go-backend/models"
)

func CreateUser(user models.User) error {
	_, error := config.DB.Exec(
		"INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		user.Name, user.Email, user.Password,
	)
	return error
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	row := config.DB.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return user, err
}

func GetAllUsers() ([]models.User, error) {
	rows, err := config.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name, &u.Email)
		users = append(users, u)
	}

	return users, nil
}

func UpdateUser(id int, user models.User) error {
	_, err := config.DB.Exec(
		"UPDATE users SET name=?, email=? WHERE id=?",
		user.Name, user.Email, id,
	)
	return err
}

func Me(id int) (models.User, error) {
	var user models.User

	row := config.DB.QueryRow("SELECT id, name FROM users WHERE id = ?", id)

	err := row.Scan(&user.ID, &user.Name)
	return user, err
}

func DeleteUser(id int) error {
	_, err := config.DB.Exec("DELETE FROM users WHERE id=?", id)
	return err
}
