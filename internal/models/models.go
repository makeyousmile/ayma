package models

type Category struct {
    ID          int
    Name        string
    Slug        string
    Description string
    SortOrder   int
    IsActive    bool
}

type Product struct {
    ID          int
    CategoryID  int
    Name        string
    Description string
    Unit        string
    Price       float64
    IsActive    bool
}
