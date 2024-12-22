# 🎉 Финальная задача
**Кто читает данный файл, желаю хорошего настроения и успешного перехода на следующий модуль!**

---

# 📊 О проекте **yandex_calculate**

**yandex_calculate** - это веб-сервис для подсчета арифметических выражений. Сервис принимает POST-запросы с телом, содержащим арифметическое выражение, и возвращает результат вычисления. Данная программа помогает пользователю быстро и удобно вычислять арифметические выражения.

---



## 📁 Структура проекта
```bash
yandex_calculate/
│
├── cmd/ # Директория с основной точкой входа приложения
│   └── main.go
│
├── internal/ # Внутренние пакеты и логики
│   ├── application/ # Логика приложения
│   │   └── application.go
│   │
│   ├── ... # Другие внутренние компоненты
│   
├── pkg/ # Внешние библиотеки и утилиты
│   ├── calculation/ # Пакет для вычислений
│   │   ├── calculation.go
│   │   ├── calculation_test.go
│   │   └── errors.go
│   │
│   ├── ... # Другие внешние компоненты
│
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

## 📁 Основные команды проекта

| Command | Description |
| --- | --- |
|``` go run ./yandex_calculate/cmd/main.go ``` | **Запуск проекта** |
|```curl --location 'localhost:8080/api/v1/calculate'``` | **Отправка POST-запроса серверу** |


---

## 📌 Примеры запросов

| Example | Description |
|---------|-------------|
| **Успешный запрос** | Returns the result of the expression. |
|  | ```bash |
|  | curl --location 'localhost:8080/api/v1/calculate' \ |
|  | --header 'Content-Type: application/json' \ |
|  | --data '{ |
|  | "expression": "2+2*2" |
|  | }' |
|  | ``` |
| **Запрос, возвращающий ошибку с кодом 500** | Returns an error due to division by zero. |
|  | ```bash |
|  | curl --location 'localhost:8080/api/v1/calculate' \ |
|  | --header 'Content-Type: application/json' \ |
|  | --data '{ |
|  | "expression": "2/0" |
|  | }' |
|  | ``` |
| **Запрос, возвращающий ошибку с кодом 422** | Returns an error due to invalid syntax. |
|  | ```bash |
|  | curl --location 'localhost:8080/api/v1/calculate' \ |
|  | --header 'Content-Type: application/json' \ |
|  | --data '{ |
|  | "expression": "2+2/*" |
|  | }' |
|  | ``` |

---

# 📊 Для вашего удобства**

**postman** - можете про тестировать отправку запросов с клиента на сервер и получить ответ от сервера
**скачать postman** - https://www.postman.com/downloads/

# 🙌 Спасибо за использование **yandex_calculate**!
