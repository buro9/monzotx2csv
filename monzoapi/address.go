package monzoapi

type Address struct {
	ShortFormatted string  `json:"short_formatted"`
	Formatted      string  `json:"formatted"`
	Address        string  `json:"address"`
	City           string  `json:"city"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Postcode       string  `json:"postcode"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	ZoomLevel      int     `json:"zoom_level"`
	Approximate    bool    `json:"approximate"`
}
