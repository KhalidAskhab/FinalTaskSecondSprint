# Калькулятор на Go
Веб-сервис для вычисления математических выражений


## Ссылка для скачивания
https://github.com/KhalidAskhab/FinalTaskSecondSprint.git

## Запуск
Для запсука веб-сервиса проделайте следующие операции:

1. Запустите веб-сервис командой
    ```
    go run ./cmd/calc_service/main.go
    ```

2. После запуска сервис будет доступен по адресу [http://localhost:8080/api/v1/calculate](http://localhost:8080/api/v1/calculate)

## Функционал
Данный веб-сервис принимает POST запрос к эндпоинту `/api/v1/calculate` содержащий математическое выражение и возвращает результат его выполнения.


## Запросы на Windows

### Для успешного вычисления арифметического выражения, например, 2 + 2, используйте следующий curl запрос:
```
curl --location "http://localhost:8080/api/v1/calculate" ^
--header "Content-Type: application/json" ^
--data "{\"expression\": \"2 + 2\"}"
```
Ожидаемый ответ:
```
{"result": "4"}
```
### Ошибка 422 (Expression is not valid)
Если пользователь введет некорректное выражение, например, 2 + a, то сервис должен вернуть ошибку 422:
```
curl --location "http://localhost:8080/api/v1/calculate" ^
--header "Content-Type: application/json" ^
--data "{\"expression\": \"2 + a\"}"
```

Ожидаемый ответ:
```
{"error": "Expression is not valid"}
```
### Ошибка 500 (Internal server error)
Для имитации внутренней ошибки сервиса, например, деление на ноль, используйте:
```
curl --location "http://localhost:8080/api/v1/calculate" ^
--header "Content-Type: application/json" ^
--data "{\"expression\": \"1 / 0\"}"
```

Ожидаемый ответ:
```
{"error": "Internal server error"}
```

###  Ошибка 400 (Некорректный JSON)
Если отправить некорректный JSON, например, без кавычек вокруг ключа:
```
curl --location "http://localhost:8080/api/v1/calculate" ^
--header "Content-Type: application/json" ^
--data "{expression: \"2 + 2\"}"
```
Ожидаемый ответ:
```
{"error": "Некорректный JSON"}
```

## Запросы для MacOS и Linux

### Для успешного вычисления арифметического выражения, например, 2 + 2, используйте следующий curl запрос:
```
curl --location "http://localhost:8080/api/v1/calculate" \
--header "Content-Type: application/json" \
--data "{\"expression\": \"2 + 2\"}"
```
Ожидаемый ответ:
```
{"result": "4"}
```
### Ошибка 422 (Expression is not valid)
Если пользователь введет некорректное выражение, например, 2 + a, то сервис должен вернуть ошибку 422:
```
curl --location "http://localhost:8080/api/v1/calculate" \
--header "Content-Type: application/json" \
--data "{\"expression\": \"2 + a\"}"
```

Ожидаемый ответ:
```
{"error": "Expression is not valid"}
```
### Ошибка 500 (Internal server error)
Для имитации внутренней ошибки сервиса, например, деление на ноль, используйте:
```
curl --location "http://localhost:8080/api/v1/calculate" \
--header "Content-Type: application/json" \
--data "{\"expression\": \"1 / 0\"}"
```

Ожидаемый ответ:
```
{"error": "Internal server error"}
```

###  Ошибка 400 (Некорректный JSON)
Если отправить некорректный JSON, например, без кавычек вокруг ключа:
```
curl --location "http://localhost:8080/api/v1/calculate" \
--header "Content-Type: application/json" \
--data "{expression: \"2 + 2\"}"
```
Ожидаемый ответ:
```
{"error": "Некорректный JSON"}
```
