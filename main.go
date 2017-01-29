package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/buro9/monzotx2csv/csv"
	"github.com/buro9/monzotx2csv/monzoapi"
)

func main() {
	// Read the file
	flag.Parse()
	in := flag.Arg(0)
	b, err := ioutil.ReadFile(in)
	if err != nil {
		fmt.Printf("error reading file: %s\n", err.Error())
		os.Exit(1)
	}

	// Parse as JSON
	var txs monzoapi.Transactions
	err = json.Unmarshal(b, &txs)
	if err != nil {
		fmt.Printf("error reading JSON: %s\n", err.Error())
		os.Exit(1)
	}

	// Output the CSV
	options := csv.Options{
		// These are my options, feel free to change them, check out
		// csv/options.go and you'll see that if you don't supply any options we
		// will default them anyway.
		Fields: []string{
			"created",
			"settled",
			"decline_reason",
			"category",
			"merchant",
			"merchant.address.postcode",
			"local_amount",
			"amount",
			"account_balance",
			"notes",
		},
	}

	if err := csv.CSV(txs, options); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
