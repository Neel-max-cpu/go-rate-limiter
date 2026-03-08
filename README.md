# Go API Rate Limiter

A backend service written in **Go** that provides **API key management and rate limiting** using **Redis** and **PostgreSQL**.

Users can register to receive an **API key**, and all API requests are **rate limited using Redis**.

This project demonstrates a **clean Go project structure**, **middleware**, **database usage**, and **containerized development using Docker**.

---

# Project Idea

Build a backend service where:

* Users register and receive an **API key**
* Every API request is **rate limited**
* **Redis** is used to track request counts
* **PostgreSQL** stores users and API keys

---

# Tech Stack

| Technology | Purpose                 |
| ---------- | ----------------------- |
| Go         | Backend language        |
| Gin        | HTTP API framework      |
| PostgreSQL | Persistent data storage |
| Redis      | Rate limiting           |
| Docker     | Containerization        |

---

# Go Packages Used

Install dependencies:

```
go get github.com/gin-gonic/gin
go get github.com/redis/go-redis/v9
go get github.com/jackc/pgx/v5
go get github.com/google/uuid
go get github.com/joho/godotenv
go get go.uber.org/zap
```

| Package  | Purpose                    |
| -------- | -------------------------- |
| Gin      | HTTP API framework         |
| pgx      | PostgreSQL driver          |
| go-redis | Redis client               |
| uuid     | Generate API keys          |
| godotenv | Load environment variables |
| zap      | Structured logging         |

---

# Initialize Project

Install go mod

```
go mod init github.com/<username>/<repo-name>
```

Example:

```
go mod init github.com/Neel-max-cpu/go-rate-limiter
```

If dependencies break:

```
go mod tidy
```

---

# Project Structure

```
cmd/
   server/
       main.go

internal/
   config/
   db/
   redis/
   handlers/
   middleware/
   models/
   services/
   utils/
```

Explanation:

| Folder     | Purpose                     |
| ---------- | --------------------------- |
| cmd/server | Entry point of application  |
| config     | Environment configuration   |
| db         | PostgreSQL connection       |
| redis      | Redis client setup          |
| handlers   | HTTP handlers (controllers) |
| middleware | Rate limiting middleware    |
| models     | Data structures             |
| services   | Business logic              |
| utils      | Helper utilities            |

---

# Application Flow

```
Client
   ↓
Router
   ↓
Middleware
   ↓
Handler
   ↓
Service
   ↓
Database / Redis
```

---

# Running the Application

### Run locally

Start dependencies only:

```
docker compose up postgres redis -d
```

Then run Go server:

```
go run cmd/server/main.go
```

Optional:

```
go run cmd/server/main.go --env=local
```

---

# Running Fully in Docker

Build and start everything:

```
docker compose up --build
```

---

# Docker Overview

* **Dockerfile** → recipe to build ONE container
* **docker-compose.yml** → recipe to run MANY containers together

---

# Docker Commands Cheat Sheet

Start services:

```
docker compose up -d
```

Check running containers:

```
docker ps
```

Stop services:

```
docker compose stop
```

Restart stopped services:

```
docker compose start
```

Remove containers but keep data:

```
docker compose down
```

Full reset (delete data):

```
docker compose down -v
```

---

# Recommended Daily Docker Workflow

### start

```
docker compose up -d
```

### pause

```
docker compose stop
docker compose start
```

### shutdown

```
docker compose down
```


### delete

```
docker compose down -v
```

---

# Gin Basics

Creating a router:

```go
r := gin.Default()
```

This automatically attaches two middlewares:

* **Logger**
* **Recovery** (prevents server crash)

---

# Go Visibility Rules

Go does not use `public/private` keywords.

Instead:

| Naming                 | Visibility        |
| ---------------------- | ----------------- |
| Uppercase first letter | Exported (public) |
| Lowercase first letter | Private           |

Example:

```
type User struct {}   // public
type user struct {}   // private
```

---

# Development Workflow

### Development Mode

Run dependencies only:

```
docker compose up postgres redis -d
```

Run Go server locally:

```
go run cmd/server/main.go
```

### Production-like Mode

Run everything via Docker:

```
docker compose up --build
```

---

# Learning Outcomes

This project teaches:

* REST API design
* Go project architecture
* Middleware usage
* Rate limiting algorithms
* Redis usage
* PostgreSQL integration
* Docker containerization
* Clean service-layer architecture

---

# Author

Built as a learning project for backend development using Go.
