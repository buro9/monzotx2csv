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

	if err := csv.CSV(txs); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
