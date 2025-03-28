# REST API на Go

## Описание

Этот проект представляет собой REST API для управления пользователями. API позволяет создавать, получать, обновлять и удалять пользователей в базе данных PostgreSQL. Разработка ведётся на языке Go с использованием фреймворка Gin и ORM GORM.

## Функциональность

- **Создание пользователя** (`POST /users`)
- **Получение информации о пользователе** (`GET /users/:id`)
- **Обновление данных пользователя** (`PUT /users/:id`)
- **Удаление пользователя** (`DELETE /users/:id`)

## Технологии

- **Язык программирования**: Go
- **Фреймворк**: Gin
- **База данных**: PostgreSQL
- **ORM**: GORM
- **Логирование**: Zap
- **Контейнеризация**: Docker
- **Развертывание**: Task (опционально)

## Установка и запуск

### Требования
- Go 1.20+
- Docker и Docker Compose
- PostgreSQL

### Запуск локально

1. Склонируйте репозиторий:
   ```sh
   git clone https://github.com/Sergey-Polishchenko/simple-api.git
   cd simple-api
   ```

2. Установите зависимости:
   ```sh
   go mod tidy
   ```

3. Собирите проект используя Docker:
   ```sh
   task build
   ```

4. Запустите сервер:
   ```sh
   task up
   ```

### Запуск с Docker

1. Соберите и запустите контейнеры:
   ```sh
   task build
   task up
   ```

## Переменные окружения

Приложение использует переменные окружения, которые можно задать в файле `.env`:

```ini
PORT=8080

DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_database
DB_HOST=localhost
DB_PORT=5432
```
**Можете просто скопировать переменные окружения из примера**:
```sh
cp .env.example .env
```

## Тестирование

Для запуска тестов выполните команду:
```sh
 go test ./...
```

## API Документация

Этот API документирован с помощью OpenAPI.  
Спецификация доступна в файле [`api.yml`](./api/openapi/api.yml). 

### 1. Создать пользователя
**POST** `/users`
#### Запрос:
```json
{
  "name": "Иван Иванов"
}
```
#### Ответ:
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "name": "Иван Иванов"
}
```

### 2. Получить пользователя
**GET** `/users/:id`
#### Ответ:
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "name": "Иван Иванов"
}
```

### 3. Обновить пользователя
**PUT** `/users/:id`
#### Запрос:
```json
{
  "name": "Пётр Петров"
}
```
#### Ответ:
```json
{
  "message": "user updated"
}
```

### 4. Удалить пользователя
**DELETE** `/users/:id`
#### Ответ:
```json
{
  "message": "user removed"
}
```

## TODO

- [ ] Увеличить покрытие тестами.
- [ ] Добавить gRPC api.
- [ ] Добавить кэширование с использование Redis.

## Лицензия

Проект распространяется под лицензией [MIT](LICENSE).
