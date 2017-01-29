package monzoapi

import (
	"fmt"
	"strings"
)

// FormatAmount will take a 3-letter ISO 4127 currency code and an int64 value
// and format it as a string, performing correction of decimal places (scale)
// and additionally offering support for accountancy style numbers.
func FormatAmount(
	currency string,
	amount int64,
	withSymbol *bool,
	accountingStyle *bool,
) (string, error) {

	if amount == 0 {
		return "", nil
	}

	curr, ok := Currencies[currency]
	if !ok {
		return "", fmt.Errorf("currency not found")
	}

	var (
		symbol   string
		template string
	)
	switch {
	case curr.UniqueSymbol != nil:
		symbol = curr.UniqueSymbol.Grapheme
		template = curr.UniqueSymbol.Template
	case curr.Symbol != nil:
		symbol = curr.Symbol.Grapheme
		template = curr.Symbol.Template
	default:
		symbol = curr.Code
		template = "$1 $2"
	}

	// Shift the display value by the currency sub units.
	v := float64(amount)
	switch curr.Scale {
	case 0:
		// do nothing
	case 2:
		v = v / 100
	case 3:
		v = v / 1000
	case 4:
		v = v / 10000
	default:
		return "", fmt.Errorf("unexpected scale for %s", curr.Code)
	}

	if withSymbol != nil && *withSymbol {
		template = strings.Replace(template, "$1", symbol, 1)
	} else {
		template = `$2`
	}

	var out string
	if accountingStyle != nil && *accountingStyle {
		switch {
		case v < 0:
			// negative: £ (%v)
			format := strings.Replace(
				template,
				"$2",
				fmt.Sprintf("(%%.%df)", curr.Scale),
				1,
			)
			out = strings.Replace(fmt.Sprintf(format, v), "-", "", 1)
		case v == 0:
			// zero: £ --
			out = strings.Replace(template, "$2", "--", 1)
		default:
			// positive: £ %v
			format := strings.Replace(
				template,
				"$2",
				fmt.Sprintf("%%.%df", curr.Scale),
				1,
			)
			out = fmt.Sprintf(format, v)
		}
	} else {
		format := strings.Replace(
			template,
			"$2",
			fmt.Sprintf("%%.%df", curr.Scale),
			1,
		)
		out = fmt.Sprintf(format, v)
	}

	return out, nil
}
