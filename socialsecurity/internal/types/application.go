package types

import "time"

type Application struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateApplicationRequest struct {
	UserID      string `json:"user_id"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type UpdateApplicationStatusRequest struct {
	Status string `json:"status"`
}

type EditApplicatonRequest struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

const (
	StatusPending   = "pending"
	StatusApproved  = "approved"
	StatusRejected  = "rejected"
	StatusCancelled = "cancelled"
)
