# Quick Task Manager Backend

Welcome to the **Quick Task Manager Backend** repository! This project serves as the backend for the Quick Task Manager application, providing APIs and core functionality for task management and user authentication.

## Table of Contents

- [About](#about)
- [Features](#features)
- [Technologies](#technologies)
- [Setup](#setup)
- [Usage](#usage)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)

## About

The Quick Task Manager backend is built using Go (Golang) and provides RESTful APIs for managing tasks, handling user authentication, and integrating with a PostgreSQL database for persistent storage.

## Features

- User authentication using JWT (JSON Web Tokens)
- CRUD operations for tasks
- Database migrations for schema management
- Middleware for authentication and request validation
- Comprehensive unit and integration tests

## Technologies

- **Language:** Go (Golang)
- **Framework:** Fiber (for routing)
- **Database:** PostgreSQL
- **Authentication:** JWT
- **Environment Configuration:** Viper
- **Testing:** Mock-based testing

## Setup

### Prerequisites

Ensure you have the following installed on your machine:

- Go (v1.19 or later)
- PostgreSQL
- Make (optional, for automation)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/dsapoetra/quick-task-manager.git
   cd quick-task-manager/backend
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up the environment variables by creating a `.env` file in the project root:

   ```env
    DB_HOST=
    DB_PORT=
    DB_USER=
    DB_PASSWORD=
    DB_NAME=
    JWT_SECRET=
   ```

4. Apply database migrations:

   ```bash
   make migrate-up
   ```
### Running the Server

Start the server using:

```bash
make run
```

Or directly:

```bash
go run cmd/server/main.go
```

The server will run at [http://localhost:8080](http://localhost:8080).

## API Documentation

API documentation is provided using Swagger. You can access the Swagger UI by navigating to:

```
http://localhost:8080/swagger/index.html
```

### Key Endpoints

#### User Authentication

- `POST /auth/login` - Authenticate and obtain a JWT.
- `POST /auth/register` - Register a new user.

#### Tasks

- `GET /tasks` - Retrieve all tasks.
- `POST /tasks` - Create a new task.
- `PUT /tasks/:id` - Update a task.
- `DELETE /tasks/:id` - Delete a task.

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature-name`)
3. Commit your changes (`git commit -m 'Add feature'`)
4. Push to the branch (`git push origin feature-name`)
5. Open a pull request

---

Feel free to reach out for any questions or suggestions. Happy coding!
