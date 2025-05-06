# Wallet App

A simple digital wallet API with Golang, Gin, GORM, and PostgreSQL.

## Features
- Withdraw funds
- Check balance
- Validations and error handling

## Requirements
- Golang
- PostgreSQL

## Setup Instructions

1. Clone repository:
```
git clone <repo-url>
```

2. Set up PostgreSQL and create database `wallet_db`.

3. Configure `.env` with your database credentials:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=deddy
DB_PASSWORD=123456789
DB_NAME=wallet_db
```

4. Install dependencies:
```
go mod tidy
```

5. Run the application:
```
go run main.go
```

6. Test API:


-  balance
```
curl --location 'http://localhost:8080/api/balance/2' \
--header 'Cookie: JSESSIONID=4205F2AB221DD365EFD75D1C508D232F'
```

-  Withdraw
```
curl --location 'http://localhost:8080/api/withdraw' \
--header 'Content-Type: application/json' \
--header 'Cookie: JSESSIONID=4205F2AB221DD365EFD75D1C508D232F' \
--data '{
    "id": "1",
    "amount":100
}'
```