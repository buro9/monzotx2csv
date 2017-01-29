package monzoapi

type Attachment struct {
	Created    Time   `json:"created"`
	ExternalID string `json:"external_id"`
	FileType   string `json:"file_type,omitempty"`
	FileURL    string `json:"file_url"`
	ID         string `json:"id"`
	Type       string `json:"type,omitempty"`
	URL        string `json:"url"`
	UserID     string `json:"user_id"`
}
