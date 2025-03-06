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
│   ├── agent
│       └── main.go
│   ├── orchestrator
│   │   └── main.go
│   │
│   ├── main.go
├── internal/ # Внутренние пакеты и логики
│   ├── application/ # Логика приложения
│   │   └── agent
          └── agent.go
           └── client.go
             └── worker.go
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
└── README.md
```

## 📁 Основные команды проекта

| Command | Description |
| --- | --- |
|``` go run ./yandex_calculate/cmd/main.go ``` | **Запуск проекта** |
|```go run ./cmd/main.go``` | **Запуск оркестратора** |
|```COMPUTING_POWER=4 go run ./cmd/main.go``` | **Запуск агента** |



---

## 📌 Примеры запросов

📌 Примеры запросов
<h3>1. Добавление выражения</h3>

Запрос:
bash
Copy

curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2 + 2 * 2"
}'

Ответ:
json
Copy

{
  "id": 1
}

<h3>2. Получение списка выражений</h3>

Запрос:
bash
Copy

curl --location 'localhost:8080/api/v1/expressions'

Ответ:
json
Copy

{
  "expressions": [
    {
      "id": 1,
      "status": "pending",
      "result": 0
    }
  ]
}

<h3>3. Получение выражения по ID</h3>

Запрос:
bash
Copy

curl --location 'localhost:8080/api/v1/expressions/1'

Ответ:
json
Copy

{
  "expression": {
    "id": 1,
    "status": "completed",
    "result": 6
  }
}

<h3>4. Получение задачи для агента</h3>

Запрос:
bash
Copy

curl --location 'localhost:8080/internal/task'

Ответ:
json
Copy

{
  "task": {
    "id": 1,
    "arg1": 2,
    "arg2": 2,
    "operation": "multiplication",
    "operation_time": 1000
  }
}

<h3>5. Отправка результата вычисления</h3>

Запрос:
bash
Copy

curl --location 'localhost:8080/internal/task/result' \
--header 'Content-Type: application/json' \
--data '{
  "id": 1,
  "result": 4
}'

Ответ:
json
Copy

{}



---
Схема взаимодействия

    Пользователь отправляет POST-запрос с выражением на /api/v1/calculate.

    Оркестратор разбивает выражение на задачи и сохраняет их.

    Агент запрашивает задачи у оркестратора через /internal/task.

    Агент выполняет задачи и отправляет результаты на /internal/task/result.

    Пользователь может запросить статус выражения через /api/v1/expressions/:id.

📦 Использование Postman

Для тестирования API можно использовать Postman. Примеры запросов:

    Добавление выражения.

    Получение списка выражений.

    Получение задачи для агента.

    Отправка результата вычисления.

🙌 Спасибо за использование yandex_calculate!

Если у вас есть вопросы или предложения, создайте issue в репозитории. Удачи в обучении! 🚀
# 📊 Для вашего удобства**

**postman** - можете про тестировать отправку запросов с клиента на сервер и получить ответ от сервера
**скачать postman** - https://www.postman.com/downloads/

# 🙌 Спасибо за использование **yandex_calculate**!
