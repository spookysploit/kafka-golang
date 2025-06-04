# 📡 Kafka Golang Example

Простой проект на Go, демонстрирующий базовую работу с Apache Kafka: продюсер отправляет сообщения, а консьюмер их читает.

---

## Используемые технологии

- **Язык:** Go 1.21+
- **Брокер сообщений:** Apache Kafka
- **Zookeeper** (для Kafka)
- **Контейнеризация:** Docker, Docker Compose
- **Библиотека:** `github.com/confluentinc/confluent-kafka-go`

---

## Запуск в Docker
Скачайте репозиторий используя команду:
```
git clone https://github.com/spookysploit/kafka-golang.git
```
Перейдите в репозиторий:
```
cd ./kafka-golang
```
Запустите командой:
```
docker compose up --build
```
После запуска контейнеров также будет доступен веб-интерфейс по адресу:\
__http://localhost:9020/__

---

## Запуск producer и consumer
Скачайте все зависимости:
```
go mod tidy
```
Пример работы producer (отправка сообщений):
```
go run ./cmd/producer/main.go
```
![image](https://github.com/user-attachments/assets/aa0c0e2c-426c-48a1-8622-79a1be064feb)

--- 

Пример работы consumer (чтение сообщений):
```
go run ./cmd/consumer/main.go
```
![image](https://github.com/user-attachments/assets/d2e1f11f-947f-4121-9192-9ef24eea6941)
