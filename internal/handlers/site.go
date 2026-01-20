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
    ContentTemplate string
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
        ContentTemplate: "home_content",
    }
    renderSite(w, h.templates, "layout", data)
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
        ContentTemplate: "catalog_content",
    }
    renderSite(w, h.templates, "layout", data)
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
        ContentTemplate: "category_content",
    }
    renderSite(w, h.templates, "layout", data)
}

func (h *SiteHandler) Contacts(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/contacts" {
        http.NotFound(w, r)
        return
    }

    data := SiteData{
        Title:  "Контакты",
        Config: h.config,
        ContentTemplate: "contacts_content",
    }
    renderSite(w, h.templates, "layout", data)
}

func renderSite(w http.ResponseWriter, tmpl *templates.Templates, name string, data SiteData) {
    if err := tmpl.Site.ExecuteTemplate(w, name, data); err != nil {
        http.Error(w, "template error", http.StatusInternalServerError)
    }
}
