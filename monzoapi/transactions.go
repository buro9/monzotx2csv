package monzoapi

// Transactions is the outer JSON object from the Monzo API GET /transactions
type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}
