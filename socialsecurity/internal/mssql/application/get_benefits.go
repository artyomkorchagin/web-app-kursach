package mssqlApplication

import (
	"context"
	"fmt"
	"socialsecurity/internal/types"
)

func (r *ApplicationRepository) GetBenefits(ctx context.Context) ([]types.Benefit, error) {
	// Define the SQL query
	query := `
        SELECT 
            benefit_id, 
            name, 
            description, 
            amount 
        FROM benefits`

	// Create a slice to store the benefits
	var benefits []types.Benefit

	// Execute the query
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Iterate over the rows and map them to Benefit structs
	for rows.Next() {
		var benefit types.Benefit
		if err := rows.Scan(&benefit.BenefitID, &benefit.Name, &benefit.Description, &benefit.Amount); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		benefits = append(benefits, benefit)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	// If no benefits are found, return an empty slice
	if len(benefits) == 0 {
		return []types.Benefit{}, nil
	}

	return benefits, nil
}
