# Calculator API

Простой веб-сервис для вычисления арифметических выражений, реализованный на Go.

## Установка

1. Убедитесь, что на вашем компьютере установлен [Golang](https://go.dev/dl/) версии **1.23** или выше.
2. Склонируйте проект из вашего репозитория (или загрузите его вручную).

## Запуск

1. Откройте терминал и выполните команду:

    ```bash
    go run ./cmd/main.go
    ```

    Сервер запустится на порту **8080**.

2. Откройте другой терминал (или командную строку) и отправьте HTTP-запрос с помощью `curl`.

---

## Пример использования

### Успешный запрос

```bash
# Запрос:
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'

# Ответ:
{
  "result": "6"
}
```


Вот исходный README.md файл для вашего проекта:

markdown
Копировать код
# Calculator API

Простой веб-сервис для вычисления арифметических выражений, реализованный на Go.

## Установка

1. Убедитесь, что на вашем компьютере установлен [Golang](https://go.dev/dl/) версии **1.23** или выше.
2. Склонируйте проект из вашего репозитория (или загрузите его вручную).

## Запуск

1. Откройте терминал и выполните команду:

    ```bash
    go run ./cmd/main.go
    ```

    Сервер запустится на порту **8080**.

2. Откройте другой терминал (или командную строку) и отправьте HTTP-запрос с помощью `curl`.

---

## Пример использования

### Успешный запрос

```bash
# Запрос:
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'

# Ответ:
{
  "result": "6"
}
```

### Ошибка 422

```bash
# Запрос:
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*a"
}'


# Ответ:
{
  "error": "Expression is not valid"
}
```

### Ошибка 500

```bash
# Запрос:
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{}'

# Ответ:
{
  "error": "Internal server error"
}
```
