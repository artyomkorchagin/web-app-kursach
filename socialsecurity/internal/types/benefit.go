package types

type Benefit struct {
	BenefitID   string `json:"benefit_id"`
	Name        string `json:"benefit_name"`
	Description string `json:"description"`
	Amount      int    `json:"provider"`
	Frequency   int    `json:"frequency"`
}
