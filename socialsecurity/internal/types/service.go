package types

type Service struct {
	ServiceID   string `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Provider    string `json:"provider"`
}
