package monzoapi

// Currencies is a map of ISO 4217 3-letter currency codes to descriptions of
// how to handle the currency
var Currencies map[string]Currency

// Currency describes a currency, and how to format it
type Currency struct {
	Code         string
	Name         string
	Scale        int
	Symbol       *Symbol
	UniqueSymbol *Symbol
}

// Symbol describes how to print a currency symbol, i.e. £ and where the symbol
// is placed relative to the value, and whether this is an RTL currency.
type Symbol struct {
	Grapheme    string
	Template    string
	RightToLeft bool
}

// setCurrencies is expected to be called by init()
func setCurrencies() {
	Currencies = make(map[string]Currency)

	Currencies["AED"] = Currency{
		Code:  "AED",
		Name:  "UAE Dirham",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    ".د.إ",
			Template:    "$2 $1",
			RightToLeft: true,
		},
	}
	Currencies["AFN"] = Currency{
		Code:  "AFN",
		Name:  "Afghani",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "؋",
			Template:    "$2 $1",
			RightToLeft: true,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "؋",
			Template:    "$2 $1",
			RightToLeft: true,
		},
	}
	Currencies["ALL"] = Currency{
		Code:  "ALL",
		Name:  "Lek",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "L",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "Lek",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["AMD"] = Currency{
		Code:  "AMD",
		Name:  "Armenian Dram",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "դր.",
			Template:    "$2 $1",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "դր.",
			Template:    "$2 $1",
			RightToLeft: false,
		},
	}
	Currencies["ANG"] = Currency{
		Code:  "ANG",
		Name:  "Netherlands Antillean Guilder",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "ƒ",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "NAƒ",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["AOA"] = Currency{
		Code:  "AOA",
		Name:  "Kwanza",
		Scale: 2,
	}
	Currencies["ARS"] = Currency{
		Code:  "ARS",
		Name:  "Argentine Peso",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["AUD"] = Currency{
		Code:  "AUD",
		Name:  "Australian Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "A$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["AWG"] = Currency{
		Code:  "AWG",
		Name:  "Aruban Florin",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "ƒ",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "Afl",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["AZN"] = Currency{
		Code:  "AZN",
		Name:  "Azerbaijanian Manat",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₼",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "₼",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["BAM"] = Currency{
		Code:  "BAM",
		Name:  "Convertible Mark",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "KM",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "KM",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["BBD"] = Currency{
		Code:  "BBD",
		Name:  "Barbados Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["BDT"] = Currency{
		Code:  "BDT",
		Name:  "Taka",
		Scale: 2,
	}
	Currencies["BGN"] = Currency{
		Code:  "BGN",
		Name:  "Bulgarian Lev",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "лв",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "лв",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["BHD"] = Currency{
		Code:  "BHD",
		Name:  "Bahraini Dinar",
		Scale: 3,
		Symbol: &Symbol{
			Grapheme:    ".د.ب",
			Template:    "$2 $1",
			RightToLeft: true,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    ".د.ب",
			Template:    "$2 $1",
			RightToLeft: true,
		},
	}
	Currencies["BIF"] = Currency{
		Code:  "BIF",
		Name:  "Burundi Franc",
		Scale: 0,
	}
	Currencies["BMD"] = Currency{
		Code:  "BMD",
		Name:  "Bermudian Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "BD$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["BND"] = Currency{
		Code:  "BND",
		Name:  "Brunei Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["BOB"] = Currency{
		Code:  "BOB",
		Name:  "Boliviano",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "Bs.",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "Bs.",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["BOV"] = Currency{
		Code:  "BOV",
		Name:  "Mvdol",
		Scale: 2,
	}
	Currencies["BRL"] = Currency{
		Code:  "BRL",
		Name:  "Brazilian Real",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "R$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "R$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["BSD"] = Currency{
		Code:  "BSD",
		Name:  "Bahamian Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["BTN"] = Currency{
		Code:  "BTN",
		Name:  "Ngultrum",
		Scale: 2,
	}
	Currencies["BWP"] = Currency{
		Code:  "BWP",
		Name:  "Pula",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "P",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "P",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["BYN"] = Currency{
		Code:  "BYN",
		Name:  "Belarussian Ruble",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "p.",
			Template:    "$2 $1",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "р.",
			Template:    "$2 $1",
			RightToLeft: false,
		},
	}
	Currencies["BYR"] = Currency{
		Code:  "BYR",
		Name:  "Belarussian Ruble",
		Scale: 0,
		Symbol: &Symbol{
			Grapheme:    "p.",
			Template:    "$2 $1",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "р.",
			Template:    "$2 $1",
			RightToLeft: false,
		},
	}
	Currencies["BZD"] = Currency{
		Code:  "BZD",
		Name:  "Belize Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "BZ$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "BZ$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["CAD"] = Currency{
		Code:  "CAD",
		Name:  "Canadian Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "CA$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["CDF"] = Currency{
		Code:  "CDF",
		Name:  "Congolese Franc",
		Scale: 2,
	}
	Currencies["CHE"] = Currency{
		Code:  "CHE",
		Name:  "WIR Euro",
		Scale: 2,
	}
	Currencies["CHF"] = Currency{
		Code:  "CHF",
		Name:  "Swiss Franc",
		Scale: 2,
	}
	Currencies["CHW"] = Currency{
		Code:  "CHW",
		Name:  "WIR Franc",
		Scale: 2,
	}
	Currencies["CLF"] = Currency{
		Code:  "CLF",
		Name:  "Unidad de Fomento",
		Scale: 4,
	}
	Currencies["CLP"] = Currency{
		Code:  "CLP",
		Name:  "Chilean Peso",
		Scale: 0,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["CNY"] = Currency{
		Code:  "CNY",
		Name:  "Yuan Renminbi",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "元",
			Template:    "$2 $1",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "元",
			Template:    "$2 $1",
			RightToLeft: false,
		},
	}
	Currencies["COP"] = Currency{
		Code:  "COP",
		Name:  "Colombian Peso",
		Scale: 0,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["COU"] = Currency{
		Code:  "COU",
		Name:  "Unidad de Valor Real",
		Scale: 2,
	}
	Currencies["CRC"] = Currency{
		Code:  "CRC",
		Name:  "Cost Rican Colon",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₡",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "₡",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["CUC"] = Currency{
		Code:  "CUC",
		Name:  "Peso Convertible",
		Scale: 2,
	}
	Currencies["CUP"] = Currency{
		Code:  "CUP",
		Name:  "Cuban Peso",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$MN",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "$MN",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["CVE"] = Currency{
		Code:  "CVE",
		Name:  "Cabo Verde Escudo",
		Scale: 2,
	}
	Currencies["CZK"] = Currency{
		Code:  "CZK",
		Name:  "Czech Koruna",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "Kč",
			Template:    "$2 $1",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "Kč",
			Template:    "$2 $1",
			RightToLeft: false,
		},
	}
	Currencies["DJF"] = Currency{
		Code:  "DJF",
		Name:  "Djibouti Franc",
		Scale: 0,
	}
	Currencies["DKK"] = Currency{
		Code:  "DKK",
		Name:  "Danish Krone",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "kr",
			Template:    "$2 $1",
			RightToLeft: false,
		},
	}
	Currencies["DOP"] = Currency{
		Code:  "DOP",
		Name:  "Dominican Peso",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "RD$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "RD$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["DZD"] = Currency{
		Code:  "DZD",
		Name:  "Algerian Dinar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    ".د.ج",
			Template:    "$2 $1",
			RightToLeft: true,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    ".د.ج",
			Template:    "$2 $1",
			RightToLeft: true,
		},
	}
	Currencies["EEK"] = Currency{
		Code:  "EEK",
		Name:  "Estonian Kroon",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "kr",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["EGP"] = Currency{
		Code:  "EGP",
		Name:  "Egyptian Pound",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "£",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    ".ج.م",
			Template:    "$2 $1",
			RightToLeft: true,
		},
	}
	Currencies["ERN"] = Currency{
		Code:  "ERN",
		Name:  "Nakfa",
		Scale: 2,
	}
	Currencies["ETB"] = Currency{
		Code:  "ETB",
		Name:  "Ethiopian Birr",
		Scale: 2,
	}
	Currencies["EUR"] = Currency{
		Code:  "EUR",
		Name:  "Euro",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "€",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "€",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["FJD"] = Currency{
		Code:  "FJD",
		Name:  "Fiji Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "FJ$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["FKP"] = Currency{
		Code:  "FKP",
		Name:  "Falkland Islands Pound",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "£",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["GBP"] = Currency{
		Code:  "GBP",
		Name:  "Pound Sterling",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "£",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "£",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["GEL"] = Currency{
		Code:  "GEL",
		Name:  "Lari",
		Scale: 2,
	}
	Currencies["GGP"] = Currency{
		Code:  "GGP",
		Name:  "Guernsey Pound",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "£",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["GHC"] = Currency{
		Code:  "GHC",
		Name:  "Ghanaian Cedi",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "¢",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "¢",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["GHS"] = Currency{
		Code:  "GHS",
		Name:  "Ghan Cedi",
		Scale: 2,
	}
	Currencies["GIP"] = Currency{
		Code:  "GIP",
		Name:  "Gibraltar Pound",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "£",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["GMD"] = Currency{
		Code:  "GMD",
		Name:  "Dalasi",
		Scale: 2,
	}
	Currencies["GNF"] = Currency{
		Code:  "GNF",
		Name:  "Guine Franc",
		Scale: 0,
	}
	Currencies["GTQ"] = Currency{
		Code:  "GTQ",
		Name:  "Quetzal",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "Q",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "Q",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["GYD"] = Currency{
		Code:  "GYD",
		Name:  "Guyan Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "GY$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["HKD"] = Currency{
		Code:  "HKD",
		Name:  "Hong Kong Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "HK$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["HNL"] = Currency{
		Code:  "HNL",
		Name:  "Lempira",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "L",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "L",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["HRK"] = Currency{
		Code:  "HRK",
		Name:  "Croatian Kuna",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "kn",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "kn",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["HTG"] = Currency{
		Code:  "HTG",
		Name:  "Gourde",
		Scale: 2,
	}
	Currencies["HUF"] = Currency{
		Code:  "HUF",
		Name:  "Forint",
		Scale: 0,
		Symbol: &Symbol{
			Grapheme:    "Ft",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "Ft",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["IDR"] = Currency{
		Code:  "IDR",
		Name:  "Rupiah",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "Rp",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "Rp",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["ILS"] = Currency{
		Code:  "ILS",
		Name:  "New Israeli Sheqel",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₪",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "₪",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["IMP"] = Currency{
		Code:  "IMP",
		Name:  "Manx Pound",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "£",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["INR"] = Currency{
		Code:  "INR",
		Name:  "Indian Rupee",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₹",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "₹",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["IQD"] = Currency{
		Code:  "IQD",
		Name:  "Iraqi Dinar",
		Scale: 3,
		Symbol: &Symbol{
			Grapheme:    ".د.ع",
			Template:    "$2 $1",
			RightToLeft: true,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    ".د.ع",
			Template:    "$2 $1",
			RightToLeft: true,
		},
	}
	Currencies["IRR"] = Currency{
		Code:  "IRR",
		Name:  "Iranian Rial",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "﷼",
			Template:    "$2 $1",
			RightToLeft: true,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    ".ر.ا",
			Template:    "$2 $1",
			RightToLeft: true,
		},
	}
	Currencies["ISK"] = Currency{
		Code:  "ISK",
		Name:  "Iceland Krona",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "kr",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["JEP"] = Currency{
		Code:  "JEP",
		Name:  "Jersey Pound",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "£",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["JMD"] = Currency{
		Code:  "JMD",
		Name:  "Jamaican Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "J$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "J$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["JOD"] = Currency{
		Code:  "JOD",
		Name:  "Jordanian Dinar",
		Scale: 3,
		Symbol: &Symbol{
			Grapheme:    ".د.إ",
			Template:    "$2 $1",
			RightToLeft: true,
		},
	}
	Currencies["JPY"] = Currency{
		Code:  "JPY",
		Name:  "Yen",
		Scale: 0,
		Symbol: &Symbol{
			Grapheme:    "¥",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "¥",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["KES"] = Currency{
		Code:  "KES",
		Name:  "Kenyan Shilling",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "KSh",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "KSh",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["KGS"] = Currency{
		Code:  "KGS",
		Name:  "Som",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "сом",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "сом",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["KHR"] = Currency{
		Code:  "KHR",
		Name:  "Riel",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "៛",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "៛",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["KMF"] = Currency{
		Code:  "KMF",
		Name:  "Comoro Franc",
		Scale: 0,
	}
	Currencies["KPW"] = Currency{
		Code:  "KPW",
		Name:  "North Korean Won",
		Scale: 0,
		Symbol: &Symbol{
			Grapheme:    "₩",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["KRW"] = Currency{
		Code:  "KRW",
		Name:  "Won",
		Scale: 0,
		Symbol: &Symbol{
			Grapheme:    "₩",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "₩",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["KWD"] = Currency{
		Code:  "KWD",
		Name:  "Kuwaiti Dinar",
		Scale: 3,
		Symbol: &Symbol{
			Grapheme:    ".د.ك",
			Template:    "$2 $1",
			RightToLeft: true,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    ".د.ك",
			Template:    "$2 $1",
			RightToLeft: true,
		},
	}
	Currencies["KYD"] = Currency{
		Code:  "KYD",
		Name:  "Cayman Islands Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "CI$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["KZT"] = Currency{
		Code:  "KZT",
		Name:  "Tenge",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₸",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "₸",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["LAK"] = Currency{
		Code:  "LAK",
		Name:  "Kip",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₭",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "₭",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["LBP"] = Currency{
		Code:  "LBP",
		Name:  "Lebanese Pound",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "£",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    ".ل.ل",
			Template:    "$2 $1",
			RightToLeft: true,
		},
	}
	Currencies["LKR"] = Currency{
		Code:  "LKR",
		Name:  "Sri Lank Rupee",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₨",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["LRD"] = Currency{
		Code:  "LRD",
		Name:  "Liberian Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "L$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["LSL"] = Currency{
		Code:  "LSL",
		Name:  "Loti",
		Scale: 2,
	}
	Currencies["LTL"] = Currency{
		Code:  "LTL",
		Name:  "Lithuanian Litas",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "Lt",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "Lt",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["LVL"] = Currency{
		Code:  "LVL",
		Name:  "Latvian Lats",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "Ls",
			Template:    "$2 $1",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "Ls",
			Template:    "$2 $1",
			RightToLeft: false,
		},
	}
	Currencies["LYD"] = Currency{
		Code:  "LYD",
		Name:  "Libyan Dinar",
		Scale: 3,
		Symbol: &Symbol{
			Grapheme:    ".د.ل",
			Template:    "$2 $1",
			RightToLeft: true,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    ".د.ل",
			Template:    "$2 $1",
			RightToLeft: true,
		},
	}
	Currencies["MAD"] = Currency{
		Code:  "MAD",
		Name:  "Moroccan Dirham",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    ".د.م",
			Template:    "$2 $1",
			RightToLeft: true,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    ".د.م",
			Template:    "$2 $1",
			RightToLeft: true,
		},
	}
	Currencies["MDL"] = Currency{
		Code:  "MDL",
		Name:  "Moldovan Leu",
		Scale: 2,
	}
	Currencies["MGA"] = Currency{
		Code:  "MGA",
		Name:  "Malagasy riary",
		Scale: 2,
	}
	Currencies["MKD"] = Currency{
		Code:  "MKD",
		Name:  "Denar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "ден",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "ден",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["MMK"] = Currency{
		Code:  "MMK",
		Name:  "Kyat",
		Scale: 2,
	}
	Currencies["MNT"] = Currency{
		Code:  "MNT",
		Name:  "Tugrik",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₮",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "₮",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["MOP"] = Currency{
		Code:  "MOP",
		Name:  "Pataca",
		Scale: 2,
	}
	Currencies["MRO"] = Currency{
		Code:  "MRO",
		Name:  "Ouguiya",
		Scale: 2,
	}
	Currencies["MUR"] = Currency{
		Code:  "MUR",
		Name:  "Mauritius Rupee",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₨",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["MVR"] = Currency{
		Code:  "MVR",
		Name:  "Rufiyaa",
		Scale: 2,
	}
	Currencies["MWK"] = Currency{
		Code:  "MWK",
		Name:  "Kwacha",
		Scale: 2,
	}
	Currencies["MXN"] = Currency{
		Code:  "MXN",
		Name:  "Mexican Peso",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["MXV"] = Currency{
		Code:  "MXV",
		Name:  "Mexican Unidad de Inversion (UDI)",
		Scale: 2,
	}
	Currencies["MYR"] = Currency{
		Code:  "MYR",
		Name:  "Malaysian Ringgit",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "RM",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "RM",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["MZN"] = Currency{
		Code:  "MZN",
		Name:  "Mozambique Metical",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "MT",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "MT",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["NAD"] = Currency{
		Code:  "NAD",
		Name:  "Namibi Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "N$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["NGN"] = Currency{
		Code:  "NGN",
		Name:  "Naira",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₦",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "₦",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["NIO"] = Currency{
		Code:  "NIO",
		Name:  "Cordob Oro",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "C$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "C$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["NOK"] = Currency{
		Code:  "NOK",
		Name:  "Norwegian Krone",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "kr",
			Template:    "$2 $1",
			RightToLeft: false,
		},
	}
	Currencies["NPR"] = Currency{
		Code:  "NPR",
		Name:  "Nepalese Rupee",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₨",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["NZD"] = Currency{
		Code:  "NZD",
		Name:  "New Zealand Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "NZ$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["OMR"] = Currency{
		Code:  "OMR",
		Name:  "Rial Omani",
		Scale: 3,
		Symbol: &Symbol{
			Grapheme:    "﷼",
			Template:    "$2 $1",
			RightToLeft: true,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    ".ر.ع",
			Template:    "$2 $1",
			RightToLeft: true,
		},
	}
	Currencies["PAB"] = Currency{
		Code:  "PAB",
		Name:  "Balboa",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "B/.",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "B/.",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["PEN"] = Currency{
		Code:  "PEN",
		Name:  "Nuevo Sol",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "S/",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "S/",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["PGK"] = Currency{
		Code:  "PGK",
		Name:  "Kina",
		Scale: 2,
	}
	Currencies["PHP"] = Currency{
		Code:  "PHP",
		Name:  "Philippine Peso",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₱",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "₱",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["PKR"] = Currency{
		Code:  "PKR",
		Name:  "Pakistan Rupee",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₨",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["PLN"] = Currency{
		Code:  "PLN",
		Name:  "Zloty",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "zł",
			Template:    "$2 $1",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "zł",
			Template:    "$2 $1",
			RightToLeft: false,
		},
	}
	Currencies["PYG"] = Currency{
		Code:  "PYG",
		Name:  "Guarani",
		Scale: 0,
		Symbol: &Symbol{
			Grapheme:    "Gs",
			Template:    "$2$1",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "Gs",
			Template:    "$2$1",
			RightToLeft: false,
		},
	}
	Currencies["QAR"] = Currency{
		Code:  "QAR",
		Name:  "Qatari Rial",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "﷼",
			Template:    "$2 $1",
			RightToLeft: true,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    ".ر.ق",
			Template:    "$2 $1",
			RightToLeft: true,
		},
	}
	Currencies["RON"] = Currency{
		Code:  "RON",
		Name:  "New Romanian Leu",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "lei",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "lei",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["RSD"] = Currency{
		Code:  "RSD",
		Name:  "Serbian Dinar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "Дин.",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "Дин.",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["RUB"] = Currency{
		Code:  "RUB",
		Name:  "Russian Ruble",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₽",
			Template:    "$2 $1",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "₽",
			Template:    "$2 $1",
			RightToLeft: false,
		},
	}
	Currencies["RUR"] = Currency{
		Code:  "RUR",
		Name:  "Russian Ruble",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₽",
			Template:    "$2 $1",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "₽",
			Template:    "$2 $1",
			RightToLeft: false,
		},
	}
	Currencies["RWF"] = Currency{
		Code:  "RWF",
		Name:  "Rwand Franc",
		Scale: 0,
	}
	Currencies["SAR"] = Currency{
		Code:  "SAR",
		Name:  "Saudi Riyal",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "﷼",
			Template:    "$2 $1",
			RightToLeft: true,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    ".ر.س",
			Template:    "$2 $1",
			RightToLeft: true,
		},
	}
	Currencies["SBD"] = Currency{
		Code:  "SBD",
		Name:  "Solomon Islands Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "SI$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["SCR"] = Currency{
		Code:  "SCR",
		Name:  "Seychelles Rupee",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₨",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["SDG"] = Currency{
		Code:  "SDG",
		Name:  "Sudanese Pound",
		Scale: 2,
	}
	Currencies["SEK"] = Currency{
		Code:  "SEK",
		Name:  "Swedish Krona",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "kr",
			Template:    "$2 $1",
			RightToLeft: false,
		},
	}
	Currencies["SGD"] = Currency{
		Code:  "SGD",
		Name:  "Singapore Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "S$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["SHP"] = Currency{
		Code:  "SHP",
		Name:  "Saint Helen Pound",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "£",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["SLL"] = Currency{
		Code:  "SLL",
		Name:  "Leone",
		Scale: 2,
	}
	Currencies["SOS"] = Currency{
		Code:  "SOS",
		Name:  "Somali Shilling",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "S",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "S",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["SRD"] = Currency{
		Code:  "SRD",
		Name:  "Surinam Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["SSP"] = Currency{
		Code:  "SSP",
		Name:  "South Sudanese Pound",
		Scale: 2,
	}
	Currencies["STD"] = Currency{
		Code:  "STD",
		Name:  "Dobra",
		Scale: 2,
	}
	Currencies["SVC"] = Currency{
		Code:  "SVC",
		Name:  "El Salvador Colon",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "C",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["SYP"] = Currency{
		Code:  "SYP",
		Name:  "Syrian Pound",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "£",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    ".ل.س",
			Template:    "$2 $1",
			RightToLeft: true,
		},
	}
	Currencies["SZL"] = Currency{
		Code:  "SZL",
		Name:  "Lilangeni",
		Scale: 2,
	}
	Currencies["THB"] = Currency{
		Code:  "THB",
		Name:  "Baht",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "฿",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "฿",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["TJS"] = Currency{
		Code:  "TJS",
		Name:  "Somoni",
		Scale: 2,
	}
	Currencies["TMT"] = Currency{
		Code:  "TMT",
		Name:  "Turkmenistan New Manat",
		Scale: 2,
	}
	Currencies["TND"] = Currency{
		Code:  "TND",
		Name:  "Tunisian Dinar",
		Scale: 3,
		Symbol: &Symbol{
			Grapheme:    ".د.ت",
			Template:    "$2 $1",
			RightToLeft: true,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    ".د.ت",
			Template:    "$2 $1",
			RightToLeft: true,
		},
	}
	Currencies["TOP"] = Currency{
		Code:  "TOP",
		Name:  "Pa’anga",
		Scale: 2,
	}
	Currencies["TRL"] = Currency{
		Code:  "TRL",
		Name:  "Turkish Lira",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₤",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["TRY"] = Currency{
		Code:  "TRY",
		Name:  "Turkish Lira",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₺",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "₺",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["TTD"] = Currency{
		Code:  "TTD",
		Name:  "Trinidad and Tobago Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "TT$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "TT$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["TWD"] = Currency{
		Code:  "TWD",
		Name:  "New Taiwan Dollar",
		Scale: 0,
		Symbol: &Symbol{
			Grapheme:    "NT$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "NT$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["TZS"] = Currency{
		Code:  "TZS",
		Name:  "Tanzanian Shilling",
		Scale: 0,
		Symbol: &Symbol{
			Grapheme:    "TSh",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "TSh",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["UAH"] = Currency{
		Code:  "UAH",
		Name:  "Hryvnia",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "₴",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "₴",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["UGX"] = Currency{
		Code:  "UGX",
		Name:  "Ugand Shilling",
		Scale: 0,
		Symbol: &Symbol{
			Grapheme:    "USh",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "USh",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["USD"] = Currency{
		Code:  "USD",
		Name:  "US Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["USN"] = Currency{
		Code:  "USN",
		Name:  "US Dollar (Next day)",
		Scale: 2,
	}
	Currencies["UYI"] = Currency{
		Code:  "UYI",
		Name:  "Uruguay Peso en Unidades Indexadas (URUIURUI)",
		Scale: 0,
	}
	Currencies["UYU"] = Currency{
		Code:  "UYU",
		Name:  "Peso Uruguayo",
		Scale: 0,
		Symbol: &Symbol{
			Grapheme:    "$U",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "$U",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["UZS"] = Currency{
		Code:  "UZS",
		Name:  "Uzbekistan Sum",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "so’m",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "so’m",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["VEF"] = Currency{
		Code:  "VEF",
		Name:  "Bolivar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "Bs",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "Bs",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["VND"] = Currency{
		Code:  "VND",
		Name:  "Dong",
		Scale: 0,
		Symbol: &Symbol{
			Grapheme:    "₫",
			Template:    "$2 $1",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "₫",
			Template:    "$2 $1",
			RightToLeft: false,
		},
	}
	Currencies["VUV"] = Currency{
		Code:  "VUV",
		Name:  "Vatu",
		Scale: 0,
	}
	Currencies["WST"] = Currency{
		Code:  "WST",
		Name:  "Tala",
		Scale: 2,
	}
	Currencies["XAF"] = Currency{
		Code:  "XAF",
		Name:  "CF Franc BEAC",
		Scale: 0,
	}
	Currencies["XCD"] = Currency{
		Code:  "XCD",
		Name:  "East Caribbean Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "EC$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["XDR"] = Currency{
		Code:  "XDR",
		Name:  "SDR (Special Drawing Right)",
		Scale: 0,
	}
	Currencies["XOF"] = Currency{
		Code:  "XOF",
		Name:  "CF Franc BCEAO",
		Scale: 0,
	}
	Currencies["XPF"] = Currency{
		Code:  "XPF",
		Name:  "CFP Franc",
		Scale: 0,
	}
	Currencies["XSU"] = Currency{
		Code:  "XSU",
		Name:  "Sucre",
		Scale: 0,
	}
	Currencies["XUA"] = Currency{
		Code:  "XUA",
		Name:  "ADB Unit of Account",
		Scale: 0,
	}
	Currencies["YER"] = Currency{
		Code:  "YER",
		Name:  "Yemeni Rial",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "﷼",
			Template:    "$2 $1",
			RightToLeft: true,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    ".ر.ي",
			Template:    "$2 $1",
			RightToLeft: true,
		},
	}
	Currencies["ZAR"] = Currency{
		Code:  "ZAR",
		Name:  "Rand",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "R",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "R",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["ZMW"] = Currency{
		Code:  "ZMW",
		Name:  "Zambian Kwacha",
		Scale: 2,
	}
	Currencies["ZWD"] = Currency{
		Code:  "ZWD",
		Name:  "Zimbabwe Dollar",
		Scale: 2,
		Symbol: &Symbol{
			Grapheme:    "Z$",
			Template:    "$1$2",
			RightToLeft: false,
		},
		UniqueSymbol: &Symbol{
			Grapheme:    "Z$",
			Template:    "$1$2",
			RightToLeft: false,
		},
	}
	Currencies["ZWL"] = Currency{
		Code:  "ZWL",
		Name:  "Zimbabwe Dollar",
		Scale: 2,
	}
}
