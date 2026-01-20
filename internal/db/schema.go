package db

import "database/sql"

const schemaSQL = `
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    slug TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    sort_order INT NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    category_id INT NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    unit TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS site_settings (
    key TEXT PRIMARY KEY,
    value TEXT NOT NULL
);
`

const seedSQL = `
INSERT INTO categories (name, slug, description, sort_order, is_active)
SELECT *
FROM (
    VALUES
    ('Пакеты', 'pakety', 'Пакеты всех форматов для магазинов, кафе и доставки.', 1, true),
    ('Средства индивидуальной защиты', 'siz', 'Перчатки, маски, одноразовые халаты.', 2, true),
    ('Бумажно-гигиеническая продукция', 'gigi', 'Салфетки, полотенца, туалетная бумага.', 3, true),
    ('Бумажная одноразовая посуда и упаковка', 'paper', 'Стаканы, коробки, тарелки из бумаги.', 4, true),
    ('Барные украшения', 'bar', 'Мешалки, трубочки, декор для напитков.', 5, true),
    ('Пластиковая одноразовая посуда и упаковка', 'plastic', 'Контейнеры, крышки, ложки, вилки.', 6, true),
    ('Бытовая химия', 'chem', 'Чистящие средства и расходники.', 7, true),
    ('Чековая лента, термоэтикетки, этикет-лента, бумага А4 и другое', 'paper-rolls', 'Ленты, этикетки, офисные расходники.', 8, true)
) AS seed(name, slug, description, sort_order, is_active)
WHERE NOT EXISTS (SELECT 1 FROM categories);

INSERT INTO products (category_id, name, description, unit, price, is_active)
SELECT *
FROM (
    VALUES
    (1, 'Пакеты майка 30x55', 'Усиленные пакеты для розницы.', 'уп.', 6.50, true),
    (4, 'Стакан бумажный 250 мл', 'Белый стакан для горячих напитков.', 'уп.', 9.90, true),
    (6, 'Контейнер PP 500 мл', 'Прозрачный контейнер с крышкой.', 'уп.', 12.80, true),
    (2, 'Перчатки виниловые', 'Одноразовые перчатки размера M.', 'уп.', 7.20, true),
    (7, 'Средство для пола 5л', 'Концентрированное средство.', 'шт.', 14.00, true)
) AS seed(category_id, name, description, unit, price, is_active)
WHERE NOT EXISTS (SELECT 1 FROM products);

INSERT INTO site_settings (key, value)
VALUES ('theme', 'terra')
ON CONFLICT (key) DO NOTHING;
`

func EnsureSchema(db *sql.DB) error {
	if _, err := db.Exec(schemaSQL); err != nil {
		return err
	}
	if _, err := db.Exec(seedSQL); err != nil {
		return err
	}
	return nil
}
