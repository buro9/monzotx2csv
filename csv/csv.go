package csv

import (
	"encoding/csv"
	"os"

	"github.com/buro9/monzotx2csv/monzoapi"
)

// CSV will take a slice of transactions and output as CSV according to the
// fields and options defined
func CSV(txs monzoapi.Transactions, options Options) error {
	options.Validate()

	var rows [][]string
	rows = append(rows, options.Fields)

	for _, tx := range txs.Transactions {
		var row []string
		for _, field := range options.Fields {
			row = append(row, PrintField(tx, field, options))
		}
		rows = append(rows, row)
	}

	w := csv.NewWriter(os.Stdout)
	for _, row := range rows {
		if err := w.Write(row); err != nil {
			return err
		}
	}
	w.Flush()

	return nil
}
