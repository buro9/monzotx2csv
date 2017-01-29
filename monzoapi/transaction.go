package monzoapi

// Transaction is a row in your Monzo account, it may be zero-value (active card
// check)
type Transaction struct {
	ID                string            `json:"id"`
	Created           Time              `json:"created"`
	Description       string            `json:"description"`
	Amount            int64             `json:"amount"`
	Currency          string            `json:"currency"`
	DeclineReason     string            `json:"decline_reason"`
	Merchant          *Merchant         `json:"merchant,omitempty"`
	Notes             string            `json:"notes,omitempty"`
	Metadata          map[string]string `json:"metadata"`
	AccountBalance    int64             `json:"account_balance"`
	Attachments       *[]Attachment     `json:"attachments,omitempty"`
	Category          string            `json:"category"`
	IsLoad            bool              `json:"is_load,omitempty"`
	Settled           *Time             `json:"settled,omitempty"`
	LocalAmount       int64             `json:"local_amount"`
	LocalCurrency     string            `json:"local_currency"`
	Updated           *Time             `json:"updated"`
	AccountID         string            `json:"account_id"`
	CounterParty      CounterParty      `json:"counterparty"`
	Scheme            string            `json:"scheme"`
	DedupeID          string            `json:"dedupe_id"`
	Originator        bool              `json:"originator,omitempty"`
	IncludeInSpending bool              `json:"include_in_spending,omitempty"`
}
