# 📡 Kafka Golang Example

A simple Go project that demonstrates basic work with Apache Kafka: the producer sends messages, and the consumer reads them.

---

## Technologies used

- **Language:** Go 1.21+
- **Message Broker:** Apache Kafka
- **Zookeeper** (for Kafka)
- **Containerization:** Docker, Docker Compose
- **Lib:** `github.com/confluentinc/confluent-kafka-go`

---

## Launch in Docker
Download repository using command:
```
git clone https://github.com/spookysploit/kafka-golang.git
```
Open the repository:
```
cd ./kafka-golang
```
Launch command:
```
docker compose up --build
```
After the containers are launched, the web interface will also be available at:\
__http://localhost:9020/__

---

## Launching producer and consumer
Download all the dependencies:
```
go mod tidy
```
Producer working example (sending messages):
```
go run ./cmd/producer/main.go
```
![image](https://github.com/user-attachments/assets/aa0c0e2c-426c-48a1-8622-79a1be064feb)

--- 

Consumer working example (reading messages):
```
go run ./cmd/consumer/main.go
```
![image](https://github.com/user-attachments/assets/d2e1f11f-947f-4121-9192-9ef24eea6941)
