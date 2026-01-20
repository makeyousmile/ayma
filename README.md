# Ayma — сайт поставщика одноразовой посуды

## Запуск

1. Установите PostgreSQL и создайте базу данных:

```
createdb ayma
```

2. Примените миграции:

```
psql $env:DATABASE_URL -f migrations/001_init.sql
psql $env:DATABASE_URL -f migrations/002_seed.sql
```

3. Запустите сервер:

```
go run ./cmd/server
```

## Запуск через Docker Compose

```
docker compose up --build
```

Сайт будет доступен на `http://localhost:8080`. База данных и сиды
инициализируются при первом запуске контейнера PostgreSQL.

## Переменные окружения

- `ADDR` (по умолчанию `:8080`)
- `DATABASE_URL` (по умолчанию `postgres://postgres:postgres@localhost:5432/ayma?sslmode=disable`)
- `COMPANY_NAME`
- `CITY`
- `PHONE`
- `EMAIL`
- `ADDRESS`
- `WORK_HOURS`
- `ADMIN_USER` (по умолчанию `admin`)
- `ADMIN_PASS` (по умолчанию `admin`)

## Страницы

- `/` — главная
- `/catalog` — каталог
- `/catalog/{slug}` — категория
- `/contacts` — контакты
- `/admin` — админка (Basic Auth)
