package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"ayma/internal/config"
	"ayma/internal/models"
	"ayma/internal/templates"
)

type AdminHandler struct {
	db        *sql.DB
	templates *templates.Templates
	config    *config.Config
}

type AdminData struct {
	Title        string
	Config       *config.Config
	Categories   []models.Category
	Products     []models.Product
	EditCategory *models.Category
	EditProduct  *models.Product
}

func NewAdminHandler(db *sql.DB, tmpl *templates.Templates, cfg *config.Config) *AdminHandler {
	return &AdminHandler{db: db, templates: tmpl, config: cfg}
}

func (h *AdminHandler) WithAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != h.config.AdminUser || pass != h.config.AdminPass {
			w.Header().Set("WWW-Authenticate", `Basic realm="Admin"`)
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func (h *AdminHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/admin/categories", http.StatusFound)
}

func (h *AdminHandler) Categories(w http.ResponseWriter, r *http.Request) {
	categories, err := listCategories(h.db)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	var edit *models.Category
	if idStr := r.URL.Query().Get("edit"); idStr != "" {
		if id, err := strconv.Atoi(idStr); err == nil {
			edit, _ = getCategoryByID(h.db, id)
		}
	}

	data := AdminData{
		Title:        "Категории",
		Config:       h.config,
		Categories:   categories,
		EditCategory: edit,
	}
	renderAdmin(w, h.templates, "categories", data)
}

func (h *AdminHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}

	category := models.Category{
		Name:        r.FormValue("name"),
		Slug:        slugify(r.FormValue("slug"), r.FormValue("name")),
		Description: r.FormValue("description"),
		IsActive:    r.FormValue("is_active") == "on",
	}

	if err := createCategory(h.db, category); err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/categories", http.StatusFound)
}

func (h *AdminHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	category := models.Category{
		ID:          id,
		Name:        r.FormValue("name"),
		Slug:        slugify(r.FormValue("slug"), r.FormValue("name")),
		Description: r.FormValue("description"),
		IsActive:    r.FormValue("is_active") == "on",
	}

	if err := updateCategory(h.db, category); err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/categories", http.StatusFound)
}

func (h *AdminHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := deleteCategory(h.db, id); err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/categories", http.StatusFound)
}

func (h *AdminHandler) Products(w http.ResponseWriter, r *http.Request) {
	products, err := listProducts(h.db)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	categories, err := listCategories(h.db)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	var edit *models.Product
	if idStr := r.URL.Query().Get("edit"); idStr != "" {
		if id, err := strconv.Atoi(idStr); err == nil {
			edit, _ = getProductByID(h.db, id)
		}
	}

	data := AdminData{
		Title:       "Товары",
		Config:      h.config,
		Categories:  categories,
		Products:    products,
		EditProduct: edit,
	}
	renderAdmin(w, h.templates, "products", data)
}

func (h *AdminHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}

	categoryID, err := strconv.Atoi(r.FormValue("category_id"))
	if err != nil {
		http.Error(w, "invalid category", http.StatusBadRequest)
		return
	}

	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		http.Error(w, "invalid price", http.StatusBadRequest)
		return
	}

	product := models.Product{
		CategoryID:  categoryID,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Unit:        r.FormValue("unit"),
		Price:       price,
		IsActive:    r.FormValue("is_active") == "on",
	}

	if err := createProduct(h.db, product); err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/products", http.StatusFound)
}

func (h *AdminHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	categoryID, err := strconv.Atoi(r.FormValue("category_id"))
	if err != nil {
		http.Error(w, "invalid category", http.StatusBadRequest)
		return
	}

	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		http.Error(w, "invalid price", http.StatusBadRequest)
		return
	}

	product := models.Product{
		ID:          id,
		CategoryID:  categoryID,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Unit:        r.FormValue("unit"),
		Price:       price,
		IsActive:    r.FormValue("is_active") == "on",
	}

	if err := updateProduct(h.db, product); err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/products", http.StatusFound)
}

func (h *AdminHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := deleteProduct(h.db, id); err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/products", http.StatusFound)
}

func renderAdmin(w http.ResponseWriter, tmpl *templates.Templates, name string, data AdminData) {
	page, ok := tmpl.Admin[name]
	if !ok {
		http.Error(w, "template not found", http.StatusInternalServerError)
		return
	}
	if err := page.ExecuteTemplate(w, "admin_layout", data); err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
	}
}
