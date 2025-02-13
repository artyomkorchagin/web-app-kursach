package mssqlUser

import (
	"context"
	"fmt"
	"socialsecurity/internal/types"
)

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]types.User, error) {
	// Define the SQL query
	query := `
        SELECT 
            id, 
            first_name, 
            last_name, 
            email, 
            role 
        FROM users`

	// Create a slice to store the users
	var users []types.User

	// Execute the query
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Iterate over the rows and map them to User structs
	for rows.Next() {
		var user types.User
		if err := rows.Scan(&user.UserID, &user.FirstName, &user.SecondName, &user.Email, &user.Role); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		users = append(users, user)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	// If no users are found, return an empty slice
	if len(users) == 0 {
		return []types.User{}, nil
	}

	return users, nil
}
