package repository

import (
	"database/sql"
	"go-backend/config"
	"go-backend/models"
)

func CreateItem(item models.Item) error {
	_, err := config.DB.Exec(
		"INSERT INTO items (name, user_id, remarks) VALUES (?, ?, ?)",
		item.Name, item.UserID, item.Remarks,
	)
	return err
}

func GetItemsByUser(userID int) ([]models.Item, error) {
	rows, err := config.DB.Query(
		"SELECT id, name, user_id, remarks FROM items WHERE user_id=?",
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item

	for rows.Next() {
		var item models.Item
		var remarks sql.NullString // Gunakan NullString jika db bisa NULL
		rows.Scan(&item.ID, &item.Name, &item.UserID, &remarks)
		if remarks.Valid {
			item.Remarks = remarks.String
		}
		items = append(items, item)
	}

	return items, nil
}

func DeleteItem(id int, userID int) error {
	_, err := config.DB.Exec(
		"DELETE FROM items WHERE id=? AND user_id=?",
		id, userID,
	)
	return err
}