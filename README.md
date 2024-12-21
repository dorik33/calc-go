Описание:
Веб-сервис для вычисления арифметических выражений. Пользователь отправляет выражение в формате JSON через POST-запрос, а сервис возвращает результат вычисления или сообщение об ошибке, также в формате JSON, с соответствующим HTTP-кодом статуса.
//expression - строка-выражение состоящее из односимвольных идентификаторов и знаков арифметических действий


Запуск проекта:
go run cmd/main/main.go
Сервер запускается по адресу: localhost:8081

Запуск тестов:
go test ./tests/handler_test.go  


Примеры использования:
Верное выражение:
curl -X POST -i -H "Content-Type: application/json" -d "{\"expression\": \"2+2*2\"}" localhost:8081/api/v1/calculate
Ответ:
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 21 Dec 2024 19:29:34 GMT
Content-Length: 19
{"result":"6.0000"}

Сложное верное выражение:
curl -X POST -i -H "Content-Type: application/json" -d "{\"expression\": \"(2+2)*((4*2)-1)\"}" localhost:8081/api/v1/calculate
Ответ:
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 21 Dec 2024 19:37:50 GMT
Content-Length: 20
{"result":"28.0000"}

Неверное выражение:
curl -i -X POST -H "Content-Type: application/json" -d "{\"expression\": \"2+2*a\"}" localhost:8081/api/v1/calculate
Ответ:
HTTP/1.1 422 Unprocessable Entity
Content-Type: application/json
Date: Sat, 21 Dec 2024 19:30:34 GMT
Content-Length: 25
{"error":"Invalid input"}

Пустое выражение:
curl -i -X POST -H "Content-Type: application/json" -d "{\"expression\": \"\"}" localhost:8081/api/v1/calculate
Ответ:
HTTP/1.1 422 Unprocessable Entity
Content-Type: application/json
Date: Sat, 21 Dec 2024 19:35:12 GMT
Content-Length: 25
{"error":"Invalid input"}

Деление на ноль:
curl -i -X POST -H "Content-Type: application/json" -d "{\"expression\": \"10/0\"}" localhost:8081/api/v1/calculate
Ответ:
HTTP/1.1 422 Unprocessable Entity
Content-Type: application/json
Date: Sat, 21 Dec 2024 19:36:21 GMT
Content-Length: 28
{"error":"Division by zero"}

Неверный json:
curl -i -X POST -H "Content-Type: application/json" -d '{}' localhost:8081/api/v1/calculate
Ответ:
HTTP/1.1 500 Internal Server Error
Content-Type: application/json
Date: Sat, 21 Dec 2024 19:40:54 GMT
Content-Length: 86
{"error":"Error parsing JSON: invalid character '\\'' looking for beginning of value"}


