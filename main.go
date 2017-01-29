package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

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

	// Output as CSV
	for _, tx := range txs.Transactions {
		if err := tx2csv(tx); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}

// excelTime represents time in the most Excel friendly manner, subseconds
// only work in `en` locale computers that use `.` delimits so these are not
// used here.
const excelTime string = "2006-01-02 15:04:05"

func tx2csv(tx monzoapi.Transaction) error {
	row := []string{
		tx.Created.Format(excelTime),
		fmt.Sprintf("%d", tx.Amount),
	}

	w := csv.NewWriter(os.Stdout)
	if err := w.Write(row); err != nil {
		return err
	}
	w.Flush()

	return nil
}
