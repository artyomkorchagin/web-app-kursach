package application

import (
	"context"
	"fmt"
	"socialsecurity/internal/types"
	"time"

	"github.com/google/uuid"
)

func (s *Service) CreateApplication(ctx context.Context, req types.CreateApplicationRequest) (types.Application, error) {
	// Generate a new UUID for the application ID
	applicationID := uuid.New().String()

	// Set default values for the application
	application := types.Application{
		ID:              applicationID,
		UserID:          req.UserID,
		BenefitID:       req.BenefitID,
		ServiceID:       req.ServiceID,
		Date:            time.Now().Format("2006-01-02"), // Format the date as YYYY-MM-DD
		Status:          types.StatusPending,             // Default status to "pending"
		Description:     req.Description,
		ApprovalDate:    "", // Nullable field, defaults to empty string
		RejectionReason: "", // Nullable field, defaults to empty string
	}

	// Add the application to the database
	if err := s.repo.AddApplication(ctx, &application); err != nil {
		return types.Application{}, fmt.Errorf("failed to add application: %w", err)
	}

	return application, nil
}
