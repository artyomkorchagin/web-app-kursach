package types

type Application struct {
	ID              string `json:"application_id"`
	UserID          string `json:"user_id"`
	BenefitID       string `json:"benefit_id"`
	ServiceID       string `json:"service_id"`
	Date            string `json:"application_date"`
	Status          string `json:"status"`
	ApprovalDate    string `json:"approval_date"`
	RejectionReason string `json:"rejection_reason"`
	Description     string `json:"description"`
}

type CreateApplicationRequest struct {
	UserID      string `json:"user_id"`
	BenefitID   string `json:"benefit_id"`
	ServiceID   string `json:"service_id"`
	Description string `json:"description"`
}

type UpdateApplicationStatusRequest struct {
	Status string `json:"status"`
}

type EditApplicatonRequest struct {
	Description string `json:"description"`
	BenefitID   string `json:"benefit_id"`
	ServiceID   string `json:"service_id"`
}

const (
	StatusPending   = "pending"
	StatusApproved  = "approved"
	StatusRejected  = "rejected"
	StatusCancelled = "cancelled"
)
