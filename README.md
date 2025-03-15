# go-chi-sqlite-jwt-starter

A robust Go backend starter project built with Chi router, SQLite database, and JWT authentication.

[![Go Build and Check](https://github.com/belgiannoise/go-chi-sqlite-jwt-starter/actions/workflows/build.yml/badge.svg)](https://github.com/belgiannoise/go-chi-sqlite-jwt-starter/actions/workflows/build.yml)

## Features

- [x] **Go-Chi Router**: Fast and lightweight HTTP routing
- [x] **SQLite Database**: Embedded database with robust SQL support
- [x] **JWT Authentication**: Secure authentication with role-based access control
- [x] **Protected Routes**: Authorization middleware for secure endpoints
- [x] **User Management**: Registration and login flows
- [x] **CRUD Examples**: Complete CRUD operations for categories and category groups
- [x] **Service Pattern**: Clean architecture with service interfaces and implementations
- [x] **Environment Config**: Structured configuration using environment variables
- [x] **GitHub Actions CI**: Automated build, test and linting workflow

## Project Structure

```
├── .github/workflows   # CI/CD pipeline configurations
├── cmd/                # Application entry points
├── config/             # Configuration management
├── internal/           # Private application code
│   ├── auth/           # Authentication logic and handlers
│   ├── category/       # (CRUD Example) Category domain
│   ├── category-group/ # (CRUD Example) Category group domain
│   ├── database/       # Database connection and schema
│   ├── models/         # Data structures
│   ├── provider/       # Service provider pattern
│   ├── server/         # HTTP server and routing
│   ├── user/           # User domain
│   ├── utils/          # Utility functions
│   └── validation/     # Input validation
└── data/               # Database files (gitignored)
```

## Quick Start

### Prerequisites

- Go 1.21 or higher
- Git

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/go-chi-sqlite-jwt-starter.git
   cd go-chi-sqlite-jwt-starter
   ```

2. Create a `.env` file in the project root:
   ```
   PORT=9494
   DB_FOLDER=./data
   AUTH_PRIVATE_KEY=<your-secret-key>
   ```

3. Run the application:
   ```bash
   go run cmd/server.go
   ```

4. The server will start at http://localhost:9494

### Building

```bash
go build -o go-chi-sqlite-jwt-starter ./cmd/server.go
```

## License

[MIT](LICENSE)
