package monzoapi

// Merchant describes a retailer
type Merchant struct {
	ID              string            `json:"id"`
	GroupID         string            `json:"group_id"`
	Created         Time              `json:"created"`
	Name            string            `json:"name"`
	Logo            string            `json:"logo"`
	Emoji           string            `json:"emoji"`
	Category        string            `json:"category"`
	Online          bool              `json:"online,omitempty"`
	ATM             bool              `json:"atm,omitempty"`
	Address         Address           `json:"address"`
	Updated         Time              `json:"updated"`
	Metadata        map[string]string `json:"metadata"`
	DisableFeedback bool              `json:"disable_feedback,omitempty"`
}
