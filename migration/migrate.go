package migration

import (
	"log"

	"go-backend/config"
)

func Migrate() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100),
		email VARCHAR(100) UNIQUE,
		password VARCHAR(255)
	);
	
	CREATE TABLE IF NOT EXISTS items (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100),
		user_id INT,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);
	`

	_, err := config.DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Users table ready")
}