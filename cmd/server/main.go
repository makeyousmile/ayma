package main

import (
    "context"
    "log"
    "net/http"
    "time"

    "ayma/internal/config"
    "ayma/internal/db"
    "ayma/internal/handlers"
    "ayma/internal/templates"
)

func main() {
    cfg := config.Load()

    dbConn, err := db.Open(context.Background(), cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("db open: %v", err)
    }
    defer dbConn.Close()

    tmpl, err := templates.New("web/templates")
    if err != nil {
        log.Fatalf("templates: %v", err)
    }

    site := handlers.NewSiteHandler(dbConn, tmpl, cfg)
    admin := handlers.NewAdminHandler(dbConn, tmpl, cfg)

    mux := http.NewServeMux()
    mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

    mux.HandleFunc("/", site.Home)
    mux.HandleFunc("/catalog", site.Catalog)
    mux.HandleFunc("/catalog/", site.Category)
    mux.HandleFunc("/contacts", site.Contacts)

    mux.HandleFunc("/admin", admin.WithAuth(admin.Dashboard))
    mux.HandleFunc("/admin/categories", admin.WithAuth(admin.Categories))
    mux.HandleFunc("/admin/categories/create", admin.WithAuth(admin.CreateCategory))
    mux.HandleFunc("/admin/categories/update", admin.WithAuth(admin.UpdateCategory))
    mux.HandleFunc("/admin/categories/delete", admin.WithAuth(admin.DeleteCategory))

    mux.HandleFunc("/admin/products", admin.WithAuth(admin.Products))
    mux.HandleFunc("/admin/products/create", admin.WithAuth(admin.CreateProduct))
    mux.HandleFunc("/admin/products/update", admin.WithAuth(admin.UpdateProduct))
    mux.HandleFunc("/admin/products/delete", admin.WithAuth(admin.DeleteProduct))

    srv := &http.Server{
        Addr:              cfg.Addr,
        Handler:           logging(mux),
        ReadHeaderTimeout: 5 * time.Second,
    }

    log.Printf("listening on %s", cfg.Addr)
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatalf("server: %v", err)
    }
}

func logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
    })
}
