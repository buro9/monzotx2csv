# monzotx2csv

Given a JSON file of monzo transactions from their API, produces a CSV file so that you can import into Excel or Google Docs

# Usage

Authenticate yourself at https://developers.getmondo.co.uk/ and view `GET /transactions` from the API list. Save the JSON locally as `transactions.json`.

```go
go install github.com/buro9/monzotx2csv
monzotx2csv transactions.json > transactions.csv
```
