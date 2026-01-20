package handlers

import (
	"database/sql"
	"regexp"
	"strings"

	"ayma/internal/models"
)

var slugCleaner = regexp.MustCompile(`[^\\pL\\pN-]+`)

func slugify(given string, fallback string) string {
	base := strings.TrimSpace(given)
	if base == "" {
		base = fallback
	}
	base = strings.ToLower(base)
	base = strings.ReplaceAll(base, " ", "-")
	base = slugCleaner.ReplaceAllString(base, "-")
	base = strings.Trim(base, "-")
	if base == "" {
		return "category"
	}
	return base
}

func listCategories(db *sql.DB) ([]models.Category, error) {
	rows, err := db.Query(`
        SELECT id, name, slug, description, sort_order, is_active
        FROM categories
        ORDER BY sort_order, name
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Slug, &c.Description, &c.SortOrder, &c.IsActive); err != nil {
			return nil, err
		}
		result = append(result, c)
	}
	return result, rows.Err()
}

func getCategoryBySlug(db *sql.DB, slug string) (*models.Category, error) {
	var c models.Category
	err := db.QueryRow(`
        SELECT id, name, slug, description, sort_order, is_active
        FROM categories
        WHERE slug = $1
    `, slug).Scan(&c.ID, &c.Name, &c.Slug, &c.Description, &c.SortOrder, &c.IsActive)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func getCategoryByID(db *sql.DB, id int) (*models.Category, error) {
	var c models.Category
	err := db.QueryRow(`
        SELECT id, name, slug, description, sort_order, is_active
        FROM categories
        WHERE id = $1
    `, id).Scan(&c.ID, &c.Name, &c.Slug, &c.Description, &c.SortOrder, &c.IsActive)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func createCategory(db *sql.DB, category models.Category) error {
	_, err := db.Exec(`
        INSERT INTO categories (name, slug, description, sort_order, is_active)
        VALUES ($1, $2, $3, $4, $5)
    `, category.Name, category.Slug, category.Description, category.SortOrder, category.IsActive)
	return err
}

func updateCategory(db *sql.DB, category models.Category) error {
	_, err := db.Exec(`
        UPDATE categories
        SET name = $1, slug = $2, description = $3, is_active = $4
        WHERE id = $5
    `, category.Name, category.Slug, category.Description, category.IsActive, category.ID)
	return err
}

func deleteCategory(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM categories WHERE id = $1`, id)
	return err
}

func listProducts(db *sql.DB) ([]models.Product, error) {
	rows, err := db.Query(`
        SELECT id, category_id, name, description, unit, price, is_active
        FROM products
        ORDER BY id DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.CategoryID, &p.Name, &p.Description, &p.Unit, &p.Price, &p.IsActive); err != nil {
			return nil, err
		}
		result = append(result, p)
	}
	return result, rows.Err()
}

func listProductsByCategory(db *sql.DB, categoryID int) ([]models.Product, error) {
	rows, err := db.Query(`
        SELECT id, category_id, name, description, unit, price, is_active
        FROM products
        WHERE category_id = $1 AND is_active = true
        ORDER BY id DESC
    `, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.CategoryID, &p.Name, &p.Description, &p.Unit, &p.Price, &p.IsActive); err != nil {
			return nil, err
		}
		result = append(result, p)
	}
	return result, rows.Err()
}

func getProductByID(db *sql.DB, id int) (*models.Product, error) {
	var p models.Product
	err := db.QueryRow(`
        SELECT id, category_id, name, description, unit, price, is_active
        FROM products
        WHERE id = $1
    `, id).Scan(&p.ID, &p.CategoryID, &p.Name, &p.Description, &p.Unit, &p.Price, &p.IsActive)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func createProduct(db *sql.DB, product models.Product) error {
	_, err := db.Exec(`
        INSERT INTO products (category_id, name, description, unit, price, is_active)
        VALUES ($1, $2, $3, $4, $5, $6)
    `, product.CategoryID, product.Name, product.Description, product.Unit, product.Price, product.IsActive)
	return err
}

func updateProduct(db *sql.DB, product models.Product) error {
	_, err := db.Exec(`
        UPDATE products
        SET category_id = $1, name = $2, description = $3, unit = $4, price = $5, is_active = $6
        WHERE id = $7
    `, product.CategoryID, product.Name, product.Description, product.Unit, product.Price, product.IsActive, product.ID)
	return err
}

func deleteProduct(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM products WHERE id = $1`, id)
	return err
}

func getSiteSetting(db *sql.DB, key string) (string, error) {
	var value string
	err := db.QueryRow(`SELECT value FROM site_settings WHERE key = $1`, key).Scan(&value)
	if err != nil {
		return "", err
	}
	return value, nil
}

func setSiteSetting(db *sql.DB, key string, value string) error {
	_, err := db.Exec(`
		INSERT INTO site_settings (key, value)
		VALUES ($1, $2)
		ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value
	`, key, value)
	return err
}
