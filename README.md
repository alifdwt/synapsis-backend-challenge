<h1 align="center">Synapsis Backend Challenge</h1>

> **E-Commerce Backend Challenge**, build built using the [Go programming language](https://golang.org), the [Gin](https://gin-gonic.com) web framework, and the [SQLC](https://github.com/kyleconroy/sqlc) for interacting with a [PostgreSQL](https://www.postgresql.org) database.

## 🚀 Deployments

### [Docker](https://hub.docker.com/repository/docker/alifdwt/synapsis-backend-challenge-api)

```bash
docker pull alifdwt/synapsis-backend-challenge-api
```

### Live Server

```bash
http://a131d68128f1a4eb39fd292b16010ada-1179032583.ap-southeast-2.elb.amazonaws.com
```

## 🧰 Installation

1. Clone the repository

```bash
git clone https://github.com/alifdwt/synapsis-backend-challenge.git
```

2. Install Dependencies

```bash
cd synapsis-backend-challenge
go mod tidy
```

3. Add your .env file

```sh
DB_DRIVER=postgres
DB_SOURCE=postgresql://root:secret@postgres:5432/synapsis_challenge?sslmode=disable
SERVER_ADDRESS=0.0.0.0:8080
TOKEN_SYMETRIC_KEY=12345678901234567890123456789012
ACCESS_TOKEN_DURATION=30m
```

4. Run make

```bash
make postgres
make createdb
make migrateup
make server
```

Optional:

```bash
make test
```

4. Start the server

```bash
go run main.go
```

5. By default, the server is running on port 5000. Access the server at http://localhost:8080

## 📖 Documentation

### Swagger

You can access the documentation at http://a131d68128f1a4eb39fd292b16010ada-1179032583.ap-southeast-2.elb.amazonaws.com/docs/index.html
