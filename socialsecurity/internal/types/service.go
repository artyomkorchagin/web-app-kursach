package types

type Service struct {
	ServiceID   string `json:"service_id"`
	Name        string `json:"service_name"`
	Description string `json:"description"`
	Provider    string `json:"provider"`
}
