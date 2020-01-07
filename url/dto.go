package url

// DTO has the info given by json
type DTO struct {
	ID          uint   `json:"id,string,omitempty"`
	Code        string `json:"code"`
	OriginalURL string `json:"original_url"`
}
