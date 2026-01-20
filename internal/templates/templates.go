package templates

import (
    "fmt"
    "html/template"
    "path/filepath"
)

type Templates struct {
    Site  *template.Template
    Admin *template.Template
}

func New(baseDir string) (*Templates, error) {
    funcMap := template.FuncMap{
        "formatPrice": func(value float64) string {
            return fmt.Sprintf("%.2f", value)
        },
    }

    site, err := template.New("layout.html").Funcs(funcMap).ParseGlob(filepath.Join(baseDir, "*.html"))
    if err != nil {
        return nil, err
    }

    admin, err := template.New("admin_layout.html").Funcs(funcMap).ParseGlob(filepath.Join(baseDir, "admin_*.html"))
    if err != nil {
        return nil, err
    }

    return &Templates{Site: site, Admin: admin}, nil
}
