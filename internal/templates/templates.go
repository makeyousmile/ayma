package templates

import (
	"fmt"
	"html/template"
	"path/filepath"
)

type Templates struct {
	Site  map[string]*template.Template
	Admin map[string]*template.Template
}

func New(baseDir string) (*Templates, error) {
	funcMap := template.FuncMap{
		"formatPrice": func(value float64) string {
			return fmt.Sprintf("%.2f", value)
		},
	}

	site := make(map[string]*template.Template)
	admin := make(map[string]*template.Template)

	siteLayout := filepath.Join(baseDir, "layout.html")
	adminLayout := filepath.Join(baseDir, "admin_layout.html")

	var err error
	site["home"], err = template.New("layout.html").Funcs(funcMap).ParseFiles(siteLayout, filepath.Join(baseDir, "home.html"))
	if err != nil {
		return nil, err
	}
	site["home_alt"], err = template.New("layout.html").Funcs(funcMap).ParseFiles(siteLayout, filepath.Join(baseDir, "home_alt.html"))
	if err != nil {
		return nil, err
	}
	site["home_alt2"], err = template.New("layout.html").Funcs(funcMap).ParseFiles(siteLayout, filepath.Join(baseDir, "home_alt2.html"))
	if err != nil {
		return nil, err
	}
	site["home_alt3"], err = template.New("layout.html").Funcs(funcMap).ParseFiles(siteLayout, filepath.Join(baseDir, "home_alt3.html"))
	if err != nil {
		return nil, err
	}
	site["catalog"], err = template.New("layout.html").Funcs(funcMap).ParseFiles(siteLayout, filepath.Join(baseDir, "catalog.html"))
	if err != nil {
		return nil, err
	}
	site["catalog_alt"], err = template.New("layout.html").Funcs(funcMap).ParseFiles(siteLayout, filepath.Join(baseDir, "catalog_alt.html"))
	if err != nil {
		return nil, err
	}
	site["catalog_alt2"], err = template.New("layout.html").Funcs(funcMap).ParseFiles(siteLayout, filepath.Join(baseDir, "catalog_alt2.html"))
	if err != nil {
		return nil, err
	}
	site["catalog_alt3"], err = template.New("layout.html").Funcs(funcMap).ParseFiles(siteLayout, filepath.Join(baseDir, "catalog_alt3.html"))
	if err != nil {
		return nil, err
	}
	site["category"], err = template.New("layout.html").Funcs(funcMap).ParseFiles(siteLayout, filepath.Join(baseDir, "category.html"))
	if err != nil {
		return nil, err
	}
	site["category_alt"], err = template.New("layout.html").Funcs(funcMap).ParseFiles(siteLayout, filepath.Join(baseDir, "category_alt.html"))
	if err != nil {
		return nil, err
	}
	site["category_alt2"], err = template.New("layout.html").Funcs(funcMap).ParseFiles(siteLayout, filepath.Join(baseDir, "category_alt2.html"))
	if err != nil {
		return nil, err
	}
	site["category_alt3"], err = template.New("layout.html").Funcs(funcMap).ParseFiles(siteLayout, filepath.Join(baseDir, "category_alt3.html"))
	if err != nil {
		return nil, err
	}
	site["contacts"], err = template.New("layout.html").Funcs(funcMap).ParseFiles(siteLayout, filepath.Join(baseDir, "contacts.html"))
	if err != nil {
		return nil, err
	}
	site["contacts_alt"], err = template.New("layout.html").Funcs(funcMap).ParseFiles(siteLayout, filepath.Join(baseDir, "contacts_alt.html"))
	if err != nil {
		return nil, err
	}
	site["contacts_alt2"], err = template.New("layout.html").Funcs(funcMap).ParseFiles(siteLayout, filepath.Join(baseDir, "contacts_alt2.html"))
	if err != nil {
		return nil, err
	}
	site["contacts_alt3"], err = template.New("layout.html").Funcs(funcMap).ParseFiles(siteLayout, filepath.Join(baseDir, "contacts_alt3.html"))
	if err != nil {
		return nil, err
	}

	admin["categories"], err = template.New("admin_layout.html").Funcs(funcMap).ParseFiles(adminLayout, filepath.Join(baseDir, "admin_categories.html"))
	if err != nil {
		return nil, err
	}
	admin["products"], err = template.New("admin_layout.html").Funcs(funcMap).ParseFiles(adminLayout, filepath.Join(baseDir, "admin_products.html"))
	if err != nil {
		return nil, err
	}
	admin["settings"], err = template.New("admin_layout.html").Funcs(funcMap).ParseFiles(adminLayout, filepath.Join(baseDir, "admin_settings.html"))
	if err != nil {
		return nil, err
	}

	return &Templates{Site: site, Admin: admin}, nil
}
