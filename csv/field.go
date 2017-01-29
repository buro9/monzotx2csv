package csv

import (
	"fmt"
	"strings"

	"github.com/buro9/monzotx2csv/monzoapi"
)

// excelTime represents time in the most Excel friendly manner, subseconds
// only work in `en` locale computers that use `.` delimits so these are not
// used here.
const excelTime string = "2006-01-02 15:04:05"

// PrintField will take a Monzo transaction and extract the field and format it
// as a string in the CSV according to the options
func PrintField(
	tx monzoapi.Transaction,
	field string,
	options Options,
) string {

	switch {
	case field == "created":
		return tx.Created.Format(excelTime)
	case field == "description":
		return tx.Description
	case field == "amount":
		amount, _ := monzoapi.FormatAmount(
			tx.Currency,
			tx.Amount,
			options.CurrenciesWithSymbol,
			options.AccountingStyle,
		)
		return amount
	case field == "currency":
		return tx.Currency
	case field == "decline_reason":
		return tx.DeclineReason
	case field == "merchant":
		if tx.Merchant != nil {
			return tx.Merchant.Name
		}
		if tx.IsLoad {
			return "Monzo Top Up"
		}
		return ""
	case field == "merchant.id":
		if tx.Merchant != nil {
			return tx.Merchant.ID
		}
		return ""
	case field == "merchant.group_id":
		if tx.Merchant != nil {
			return tx.Merchant.GroupID
		}
		return ""
	case field == "merchant.created":
		if tx.Merchant != nil {
			return tx.Merchant.Created.Format(excelTime)
		}
		return ""
	case field == "merchant.name":
		if tx.Merchant != nil {
			return tx.Merchant.Name
		}
		return ""
	case field == "merchant.logo":
		if tx.Merchant != nil {
			return tx.Merchant.Logo
		}
		return ""
	case field == "merchant.emoji":
		if tx.Merchant != nil {
			return tx.Merchant.Emoji
		}
		return ""
	case field == "merchant.category":
		if tx.Merchant != nil {
			return tx.Merchant.Category
		}
		return ""
	case field == "merchant.online":
		if tx.Merchant != nil {
			return fmt.Sprintf(`%t`, tx.Merchant.Online)
		}
		return ""
	case field == "merchant.atm":
		if tx.Merchant != nil {
			return fmt.Sprintf(`%t`, tx.Merchant.ATM)
		}
		return ""
	case field == "merchant.address":
		if tx.Merchant != nil {
			return tx.Merchant.Address.ShortFormatted
		}
		return ""
	case field == "merchant.address.short_formatted":
		if tx.Merchant != nil {
			return tx.Merchant.Address.ShortFormatted
		}
		return ""
	case field == "merchant.address.formatted":
		if tx.Merchant != nil {
			return tx.Merchant.Address.Formatted
		}
		return ""
	case field == "merchant.address.address":
		if tx.Merchant != nil {
			return tx.Merchant.Address.Address
		}
		return ""
	case field == "merchant.address.city":
		if tx.Merchant != nil {
			return tx.Merchant.Address.City
		}
		return ""
	case field == "merchant.address.region":
		if tx.Merchant != nil {
			return tx.Merchant.Address.Region
		}
		return ""
	case field == "merchant.address.country":
		if tx.Merchant != nil {
			return tx.Merchant.Address.Country
		}
		return ""
	case field == "merchant.address.postcode":
		if tx.Merchant != nil {
			return tx.Merchant.Address.Postcode
		}
		return ""
	case field == "merchant.address.latitude":
		if tx.Merchant != nil {
			return fmt.Sprintf("%f", tx.Merchant.Address.Latitude)
		}
		return ""
	case field == "merchant.address.longitude":
		if tx.Merchant != nil {
			return fmt.Sprintf("%f", tx.Merchant.Address.Longitude)
		}
		return ""
	case field == "merchant.address.zoom_level":
		if tx.Merchant != nil {
			return fmt.Sprintf("%d", tx.Merchant.Address.ZoomLevel)
		}
		return ""
	case field == "merchant.address.approximate":
		if tx.Merchant != nil {
			return fmt.Sprintf("%t", tx.Merchant.Address.Approximate)
		}
		return ""
	case field == "merchant.updated":
		if tx.Merchant != nil {
			if !tx.Merchant.Updated.IsZero() {
				return tx.Merchant.Updated.Format(excelTime)
			}
		}
		return ""
	case strings.HasPrefix(field, "merchant.metadata[") && strings.HasSuffix(field, "]"):
		key := field[strings.Index(field, "[")+1 : strings.Index(field, "]")]
		if tx.Merchant != nil {
			if s, ok := tx.Merchant.Metadata[key]; ok {
				return s
			}
		}
		return ""
	case field == "merchant.disable_feedback":
		if tx.Merchant != nil {
			return fmt.Sprintf("%t", tx.Merchant.DisableFeedback)
		}
		return ""
	case field == "notes":
		return tx.Notes
	case strings.HasPrefix(field, "metadata[") && strings.HasSuffix(field, "]"):
		key := field[strings.Index(field, "[")+1 : strings.Index(field, "]")]
		if s, ok := tx.Metadata[key]; ok {
			return s
		}
		return ""
	case field == "account_balance":
		amount, _ := monzoapi.FormatAmount(
			tx.Currency,
			tx.AccountBalance,
			options.CurrenciesWithSymbol,
			options.AccountingStyle,
		)
		return amount
	case field == "attachments":
		if tx.Attachments != nil && len(*tx.Attachments) > 0 {
			// Only returning the URL to the first because spreadsheets don't
			// know what to do with multiple attachments in a single field
			return (*tx.Attachments)[0].URL
		}
		return ""
	case field == "category":
		return tx.Category
	case field == "is_load":
		return fmt.Sprintf("%t", tx.IsLoad)
	case field == "settled":
		if !tx.Settled.IsZero() {
			return tx.Settled.Format(excelTime)
		}
		return ""
	case field == "local_amount":
		amount, _ := monzoapi.FormatAmount(
			tx.LocalCurrency,
			tx.LocalAmount,
			options.CurrenciesWithSymbol,
			options.AccountingStyle,
		)
		return amount
	case field == "local_currency":
		return tx.LocalCurrency
	case field == "updated":
		if !tx.Updated.IsZero() {
			return tx.Updated.Format(excelTime)
		}
		return ""
	case field == "account_id":
		return tx.AccountID
	case field == "scheme":
		return tx.Scheme
	case field == "dedupe_id":
		return tx.DedupeID
	case field == "originator":
		return fmt.Sprintf("%t", tx.Originator)
	case field == "include_in_spending":
		return fmt.Sprintf("%t", tx.IncludeInSpending)
	default:
		return ""
	}
}
