package handlers

import (
	"database/sql"
	"net/http"

	"ayma/internal/config"
	"ayma/internal/models"
	"ayma/internal/templates"
)

type SiteHandler struct {
	db        *sql.DB
	templates *templates.Templates
	config    *config.Config
}

type SiteData struct {
	Title      string
	Config     *config.Config
	Categories []models.Category
	Products   []models.Product
	Category   *models.Category
	IsHome     bool
	Theme      string
}

func NewSiteHandler(db *sql.DB, tmpl *templates.Templates, cfg *config.Config) *SiteHandler {
	return &SiteHandler{db: db, templates: tmpl, config: cfg}
}

func (h *SiteHandler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	categories, err := listCategories(h.db)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	data := SiteData{
		Title:      "Главная",
		Config:     h.config,
		Categories: categories,
		IsHome:     true,
		Theme:      h.theme(),
	}
	renderSite(w, h.templates, h.templateName(r, "home"), data)
}

func (h *SiteHandler) Catalog(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/catalog" {
		http.NotFound(w, r)
		return
	}

	categories, err := listCategories(h.db)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	data := SiteData{
		Title:      "Каталог",
		Config:     h.config,
		Categories: categories,
		Theme:      h.theme(),
	}
	renderSite(w, h.templates, h.templateName(r, "catalog"), data)
}

func (h *SiteHandler) Category(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[len("/catalog/"):]
	if slug == "" || slug == "/" {
		http.NotFound(w, r)
		return
	}

	category, err := getCategoryBySlug(h.db, slug)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	products, err := listProductsByCategory(h.db, category.ID)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	data := SiteData{
		Title:    category.Name,
		Config:   h.config,
		Category: category,
		Products: products,
		Theme:    h.theme(),
	}
	renderSite(w, h.templates, h.templateName(r, "category"), data)
}

func (h *SiteHandler) Contacts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contacts" {
		http.NotFound(w, r)
		return
	}

	data := SiteData{
		Title:  "Контакты",
		Config: h.config,
		Theme:  h.theme(),
	}
	renderSite(w, h.templates, h.templateName(r, "contacts"), data)
}

func renderSite(w http.ResponseWriter, tmpl *templates.Templates, name string, data SiteData) {
	page, ok := tmpl.Site[name]
	if !ok {
		http.Error(w, "template not found", http.StatusInternalServerError)
		return
	}
	if err := page.ExecuteTemplate(w, "layout", data); err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
	}
}

func (h *SiteHandler) theme() string {
	value, err := getSiteSetting(h.db, "theme")
	if err == nil && value != "" {
		return value
	}
	return defaultTheme
}

func (h *SiteHandler) templateName(r *http.Request, base string) string {
	variant := r.URL.Query().Get("variant")
	if variant == "alt" || variant == "alt2" || variant == "alt3" {
		return base + "_" + variant
	}
	return base
}
