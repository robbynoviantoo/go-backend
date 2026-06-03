package repository

import (
	"database/sql"
	"fmt"

	"go-backend/config"
	"go-backend/models"
)

func CreateProductCategory(category models.ProductCategory) error {
	_, err := config.DB.Exec(
		"INSERT INTO product_categories (name) VALUES (?)",
		category.Name,
	)
	return err
}

func GetProductCategories() ([]models.ProductCategory, error) {
	rows, err := config.DB.Query("SELECT id, name FROM product_categories ORDER BY name ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.ProductCategory
	for rows.Next() {
		var category models.ProductCategory
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func UpdateProductCategory(id int, category models.ProductCategory) error {
	result, err := config.DB.Exec(
		"UPDATE product_categories SET name=? WHERE id=?",
		category.Name, id,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("category not found")
	}

	return nil
}

func DeleteProductCategory(id int) error {
	result, err := config.DB.Exec("DELETE FROM product_categories WHERE id=?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("category not found")
	}

	return nil
}

func CreateProduct(product models.Product) (int64, error) {
	result, err := config.DB.Exec(
		`INSERT INTO products (user_id, category_id, name, description, price, stock)
		VALUES (?, ?, ?, ?, ?, ?)`,
		product.UserID,
		product.CategoryID,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func GetProducts(name string, categoryID string, userID string) ([]models.Product, error) {
	query := `
		SELECT
			p.id,
			p.user_id,
			p.category_id,
			p.name,
			p.description,
			p.price,
			p.stock,
			u.name AS user_name,
			c.name AS category_name
		FROM products p
		JOIN users u ON u.id = p.user_id
		JOIN product_categories c ON c.id = p.category_id
		WHERE 1=1`
	args := []interface{}{}

	if name != "" {
		query += " AND p.name LIKE ?"
		args = append(args, "%"+name+"%")
	}

	if categoryID != "" {
		query += " AND p.category_id = ?"
		args = append(args, categoryID)
	}

	if userID != "" {
		query += " AND p.user_id = ?"
		args = append(args, userID)
	}

	query += " ORDER BY p.id DESC"

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		product, err := scanProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func GetProductByID(id int) (models.Product, error) {
	row := config.DB.QueryRow(
		`SELECT
			p.id,
			p.user_id,
			p.category_id,
			p.name,
			p.description,
			p.price,
			p.stock,
			u.name AS user_name,
			c.name AS category_name
		FROM products p
		JOIN users u ON u.id = p.user_id
		JOIN product_categories c ON c.id = p.category_id
		WHERE p.id = ?`,
		id,
	)

	return scanProduct(row)
}

func UpdateProduct(id int, userID int, product models.Product) error {
	result, err := config.DB.Exec(
		`UPDATE products
		SET category_id=?, name=?, description=?, price=?, stock=?
		WHERE id=? AND user_id=?`,
		product.CategoryID,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		id,
		userID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("product not found or not owned by user")
	}

	return nil
}

func DeleteProduct(id int, userID int) error {
	result, err := config.DB.Exec(
		"DELETE FROM products WHERE id=? AND user_id=?",
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
		return fmt.Errorf("product not found or not owned by user")
	}

	return nil
}

type productScanner interface {
	Scan(dest ...interface{}) error
}

func scanProduct(scanner productScanner) (models.Product, error) {
	var product models.Product
	var description sql.NullString

	err := scanner.Scan(
		&product.ID,
		&product.UserID,
		&product.CategoryID,
		&product.Name,
		&description,
		&product.Price,
		&product.Stock,
		&product.UserName,
		&product.CategoryName,
	)
	if err != nil {
		return product, err
	}

	if description.Valid {
		product.Description = description.String
	}

	return product, nil
}
