# Digital Wallet

A Simple Wallet App .

## Features
- Check balance
- Withdraw funds

## Requirements
- Golang 1.20
- PostgreSQL latest

## Setup Instructions

1. Clone repository:
```
git clone https://github.com/fdeddys/wallet-app.git
```

2. Set up PostgreSQL and create database `wallet_db`. user `deddy` and password `123456789`

3. Configure `.env.development` with:
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
start.sh
```

6. Test in POSTMAN


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
    "amount":1000
}'
```