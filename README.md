# Amartha - Golang Developer Test

This repository focuses on achieving the objectives of the **Amartha Golang Developer Test**.

---

## 📂 Project Structure

```go
.
├── internal/           # Business logic and application layers
│   ├── domain/         # Domain models
│   │   ├── entities/   # Core entities representing business objects
│   │   ├── interfaces/ # Interfaces for defining contracts between layers
│   │   ├── usecase/    # Application use cases implementing business logic
│   ├── infrastructure/ # Infrastructure-related code (e.g., HTTP handlers, database)
│   │   ├── delivery/   # HTTP request handlers
│   │   ├── repository/ # Data access layer interacting with storage
├── test/               # Unit and integration tests
├── go.sum              # Checksum file ensuring module integrity and security
├── go.mod              # Go module dependencies
├── main.go             # Main application entry point
└── README.md           # Project documentation

```

## 📥 Getting Started

### Prerequisites

Ensure you have the following installed:

- [Go](https://go.dev/dl/) (latest stable version)
- [Git](https://git-scm.com/)
- Any required dependencies (listed in `go.mod`)

### 📌 Clone the Repository

Run the following command:

```sh
git clone <repository-url>
cd <repository-folder>
```

### 📦 Install Dependencies

```sh
go mod tidy
```

### 🚀 Running the Application

```sh
go run main.go
```

### 🧪 Running Tests

```sh
go test ./...
```

---

## ⚙️ Features

- Loan management system
- Borrower and payment handling
- Unit tests for business logic and services
- Structured project using Clean Architecture

## 📜 License

This project is for assessment purposes only.
