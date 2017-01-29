package csv

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/buro9/monzotx2csv/monzoapi"
)

// excelTime represents time in the most Excel friendly manner, subseconds
// only work in `en` locale computers that use `.` delimits so these are not
// used here.
const excelTime string = "2006-01-02 15:04:05"

func CSV(txs monzoapi.Transactions) error {
	// Output as CSV to os.Stdout
	var rows [][]string

	for _, tx := range txs.Transactions {
		row := []string{
			tx.Created.Format(excelTime),
			fmt.Sprintf("%d", tx.Amount),
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
