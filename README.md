# 🚀 Go Payments Platform (Monolith → Microservices)


A production-grade **payment platform backend** built in **Go**, inspired by Stripe/Adyen architecture.
Designed as a **monolith-first system**, with a clear roadmap to evolve into microservices.

This repository contains only the backend service. No frontend is included.

This project focuses on **correctness, security, scalability, and clean architecture**, not just features.

It was started to deeply understand how modern payment gateways are designed and built in production.

---

## ✨ Features (Current & Planned)

### ✅ Current

* Clean architecture folder structure
* Gin HTTP server
* Structured logging with Zap
* Request tracing (X-Request-ID)
* Stripe-style API responses
* Health check endpoint
* PostgreSQL integration (pgx)
* Database migrations
* User authentication (JWT + bcrypt)
* JWT authorization middleware
* Protected routes



### 🚧 In Progress

* Wallet system
* Transaction ledger (double entry)

### 🔮 Planned

* Payment orchestration
* Fraud detection hooks
* Webhooks
* Event-driven architecture (Kafka/NATS)
* Microservices split
* Rust optimization for hot paths
* ML-based risk scoring

---

## 🏗 Architecture

```bash
Monolith First → Microservices Later

                HTTP API
                   ↓
            Delivery Layer (Gin)
                   ↓
            Application Services
                   ↓
              Domain Models
                   ↓
              Repositories
                   ↓
            PostgreSQL / Redis
```

Design principles:

* Clean Architecture
* Domain-driven design
* Idempotent APIs
* Audit-safe transaction flows

---

## 📁 Project Structure

```bash
/cmd/api
/internal
  /delivery/http
  /config
  /app
  /domain
  /repository
/pkg
  /logger
  /response
/migrations
```

---

## 🔐 Security Philosophy

* No card data stored (tokenization approach)
* JWT authentication
* Bcrypt password hashing
* Request idempotency
* Audit-friendly logging
* PCI-DSS inspired architectural decisions

---

## ⚙️ Tech Stack

| Layer      | Technology                   |
| ---------- | ---------------------------- |
| Language   | Go                           |
| Framework  | Gin                          |
| Database   | PostgreSQL                   |
| Cache      | Redis                        |
| Logging    | Zap                          |
| Config     | Viper                        |
| Auth       | JWT                          |
| Migrations | golang-migrate (current), Goose / Atlas (planned)      |
| Infra      | Docker, Kubernetes (planned) |
| Messaging  | Kafka / NATS (planned)       |

---

## 🚀 Getting Started

### Prerequisites

* Go 1.22+
* PostgreSQL
* Redis

### Run locally

```bash
git clone https://github.com/your-username/go-payments-platform.git
cd go-payments-platform
go run ./cmd/api
```

Health check:

```
GET http://localhost:8080/health
```

---

## 🔑 Environment Variables

Create `.env` file:

```env
APP_ENV=development
SERVER_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_NAME=payments
DB_USER=postgres
DB_PASS=password
JWT_SECRET=change_me
```

---

## 📌 API Response Format (Stripe Style)

```json
{
  "object": "response",
  "success": true,
  "data": {},
  "error": null,
  "trace_id": "uuid"
}
```

---

## ❌ Non-Goals

* Building a frontend UI
* Processing real card data
* Competing with real PSPs

This project is purely educational and architectural.

---

## 🛣 Roadmap

* [x] Project structure
* [x] Logging + middleware
* [x] Standard API responses
* [x] PostgreSQL integration
* [x] User auth
* [ ] Wallet module
* [ ] Transactions
* [ ] Webhooks
* [ ] Fraud scoring
* [ ] Event streaming
* [ ] Microservices migration

---

## 📖 Learning Goals

This project is built to understand:

* Real payment system architecture
* Financial transaction safety
* Distributed systems concepts
* Secure API design
* Scalable backend engineering

---

## 🤝 Contributing

This project is currently in active development.
Contributions, suggestions, and discussions are welcome.

---

## 📜 License

[MIT License](LICENSE)

---

### ⭐ If you like this project, consider starring it!
