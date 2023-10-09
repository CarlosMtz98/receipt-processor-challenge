## Run webservice instructions

Hello 👋 Thank you for taking the time to test and evaluate my api implementation. 

This api service uses Docker which will take care to set up the environment and all dependencies.

## Getting Started
### Run server
To run the server locally run the following command
```bash
make local
```

#### Create new receipt
```bash
curl --location 'localhost:7070/receipts/process' \
--header 'Content-Type: application/json' \
--data '{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2023-10-08",
  "purchaseTime": "15:45",
  "items": [
    {
      "shortDescription": "Mountain Dew",
      "price": "1.85"
    },
    {
      "shortDescription": "Sour Patch Kids",
      "price": "4.21"
    },
    {
      "shortDescription": "Swedish Fish",
      "price": "1.50"
    },
    {
      "shortDescription": "Takis",
      "price": "2.30"
    }
  ],
  "total": "9.86"
}'
```
#### Get receipt points
```bash
curl --location 'localhost:7070/receipts/<receiptId>/points'
```

### Debug server
To debug the server run the following command
```bash
make run
```

### Run tests
This command will run all the available tests inside the project with the coverage they have
```bash
make tests
```

## Troubleshooting
If you got any problems using the make files you can manually execute them using the following commands

### Run server alternative
```bash
docker-compose -f docker-compose.local.yml up --build
```

### Debug server alternative
```bash
go run ./cmd/main.go
```