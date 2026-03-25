package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go-backend/config"
	"go-backend/models"
)

var ErrNotFound = errors.New("item not found")
var ErrForbidden = errors.New("forbidden")

func CreateItem(item models.Item) error {
	_, err := config.DB.Exec(
		"INSERT INTO items (name, user_id, remarks) VALUES (?, ?, ?)",
		item.Name, item.UserID, item.Remarks,
	)
	return err
}

func GetAllItems(name string, userID string) ([]models.Item, error) {
	query := "SELECT id, name, user_id, remarks FROM items WHERE 1=1"
	args := []interface{}{}

	if name != "" {
		query += " AND name LIKE ?"
		args = append(args, "%"+name+"%")
	}

	if userID != "" {
		query += " AND user_id = ?"
		args = append(args, userID)
	}

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		var remarks sql.NullString
		err := rows.Scan(&item.ID, &item.Name, &item.UserID, &remarks)
		if err != nil {
			return nil, err
		}
		if remarks.Valid {
			item.Remarks = remarks.String
		}
		items = append(items, item)
	}

	return items, nil
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
	result, err := config.DB.Exec(
		"DELETE FROM items WHERE id=? AND user_id=?",
		id, userID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("item not found or not owned by user")
	}
	return nil
}