# CoinMarketCap Currency Price Updater


## Features

Automated price updates from CoinMarketCap API every 5 minutes
RESTful endpoint to retrieve all stored prices
Clean Architecture implementation
MySQL database storage
Configurable settings via environment variables

## Setup

### 1. Clone the Repository
```bash
git clone https://github.com/akshaybt001/CoinMarket.git
cd CoinMarketCap Currency Price
```
### 2. Create a .env File
```
DB_KEY=your_postgresql_connection_string
APIKEY=your_api_key
```
### 3. Install Dependencies
```
go mod tidy
```
### 4. Run the Application
```
go run main.go
```
### Usage
* Connect to the API using Postman on port 8080.

### Project Structure
```bash
├── main.go
├── routes
│   └── routes.go
├── src
│   ├── controller
│   │   └── controller.go
│   ├── models
│   │   └── price.go
│   ├── repository
│   │   ├── repository_interface.go
│   │   └── repository.go
│   ├── service
│   │   └── service.go
├── util
│   └── database.go
├── go.mod
└── go.sum
```

### API Documentation
 #### Request 
```http
GET  http://localhost:8080/getPrice
```
#### Response
```json
{
  "data": [
    {
      "id": 1,
      "currency": "BTC",
      "price": 50000.00,
      "last_updated": "2024-03-15T14:30:00Z"
    },
    {
      "id": 2,
      "currency": "ETH",
      "price": 3000.00,
      "last_updated": "2024-03-15T14:30:00Z"
    }
  ]
}
```
