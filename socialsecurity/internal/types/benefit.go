package types

type Benefit struct {
	BenefitID   string `json:"user_id"`
	Name        string `json:"name"`
	Description int    `json:"description"`
	Amount      int    `json:"provider"`
}
