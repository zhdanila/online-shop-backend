# Online Store Project

This project is a simple implementation of an online store using Go. It includes the following:

- **Database design** for key entities like sellers, products, customers, and orders.
- **HTTP API** with basic authentication to manage CRUD operations for entities.
- **Docker** and **Docker Compose** to run the Go application and PostgreSQL database together.

---

## Technologies

- **Go (Golang)**: A statically typed, compiled language used for building the backend server.
- **PostgreSQL**: A powerful relational database used for storing application data.
- **Docker**: A platform for developing, shipping, and running applications in containers, ensuring consistency across
  environments.
- **Docker Compose**: A tool for defining and running multi-container applications, used here for running the Go server
  and PostgreSQL database.

## Dependencies

- **github.com/go-playground/validator/v10**: For validating struct fields in Go.
- **github.com/jmoiron/sqlx**: Enhances `database/sql` with additional features like named queries.
- **github.com/lib/pq**: PostgreSQL driver for Go.
- **github.com/spf13/viper**: Configuration management for reading environment variables and configuration files.

## How to Launch

### 1. **Set Environment Variables**

Before running the service, make sure to create a .env file in the root directory of the project. You can use the
example file env.default to set the necessary environment variables. Copy the contents of env.default into a new .env
file, and update the variables as needed for your environment.

### 2. **Start the Application**

To start the Go application, run the following command:

```bash
make up
```

This command will run the Go HTTP server:

### 3. **Migrate Database Up**

To apply database migrations (moving the database schema forward), use the following command:

```bash
make migrate-up
```

Make sure the environment variables (`DB_USERNAME`, `DB_PASSWORD`, `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_SSLMODE`) are
correctly set in the `.env` file.

### 4. **Reset Migrations (Migrate Down)**

To revert all applied migrations and reset the database, use the following command:

```bash
make migrate-down
```

### 5. **Build Docker Containers**

To build and run the application and database with Docker Compose, use the following command:

```bash
make docker-build
```

This command will build and start the application and PostgreSQL database in their respective containers, and the
application will be available at `http://localhost:8080`.

---

All the available endpoints are listed in the online-shop.postman_collection.json file.

---
Here's an improved version:

---

The second part of the task, related to concatenation, can be found in the `cmd/concatenation/main.go` file.

To execute the benchmarks for this part of the task, run the following command:

```bash
make bench
```