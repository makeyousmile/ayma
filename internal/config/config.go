package config

import "os"

type Config struct {
    Addr         string
    DatabaseURL  string
    CompanyName  string
    City         string
    Phone        string
    Email        string
    Address      string
    WorkHours    string
    AdminUser    string
    AdminPass    string
}

func Load() *Config {
    return &Config{
        Addr:        getEnv("ADDR", ":8080"),
        DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/ayma?sslmode=disable"),
        CompanyName: getEnv("COMPANY_NAME", "Айма Пак"),
        City:        getEnv("CITY", "Могилев"),
        Phone:       getEnv("PHONE", "+375 (XX) XXX-XX-XX"),
        Email:       getEnv("EMAIL", "info@example.by"),
        Address:     getEnv("ADDRESS", "г. Могилев, ул. Примерная, 1"),
        WorkHours:   getEnv("WORK_HOURS", "Пн–Пт 09:00–18:00, Сб 10:00–15:00"),
        AdminUser:   getEnv("ADMIN_USER", "admin"),
        AdminPass:   getEnv("ADMIN_PASS", "admin"),
    }
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}
