package csv

// Options is the set of options that will be used to produce the CSV output
type Options struct {
	// List of fields that should be in the CSV
	//
	// i.e.
	// fields := []string{
	// 	"created",
	// 	"settled",
	// 	"decline_reason",
	// 	"category",
	// 	"merchant",
	// 	"merchant.address.postcode",
	// 	"local_amount",
	// 	"amount",
	// 	"account_balance",
	// 	"notes",
	// }
	//
	// The field names conform to the origin JSON, metadata should be expressed
	// as: metadata[is_topup] with no quotes on the metadata key.
	Fields []string

	// CurrenciesWithSymbol = true means to print the £ or $ symbol alongside
	// the value when printing the amount.
	//
	// Recommended value = true, if not defined defaults to true
	CurrenciesWithSymbol *bool

	// AccountingStyle = true means to print negatives currency values as ()
	// and zeros as -- which conforms to common accounting practise.
	//
	// Recommended value = true, if not defined defaults to true
	AccountingStyle *bool
}

// Validate will check the options provided and if absent will set default
// values
func (m *Options) Validate() {
	if len(m.Fields) == 0 {
		m.Fields = []string{
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
		}
	}

	if m.CurrenciesWithSymbol == nil {
		withSymbol := true
		m.CurrenciesWithSymbol = &withSymbol
	}

	if m.AccountingStyle == nil {
		accountingStyle := true
		m.AccountingStyle = &accountingStyle
	}
}
