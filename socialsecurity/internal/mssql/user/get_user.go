package mssqlUser

import (
	"context"
	"database/sql"
	"socialsecurity/internal/types"

	"github.com/pkg/errors"
)

func (r *UserRepository) GetUser(ctx context.Context, email string) (*types.User, error) {
	// Query the database to fetch the user by email and hashed password
	query := `
        SELECT 
            [user_id],
            [first_name],
            [second_name],
            [middle_name],
            [date_of_birth],
            [gender],
            [address],
            [phone_number],
            [email],
            [password_hash],
            [snils],
            [passport_series],
            [passport_number],
            [passport_issued_by],
            [passport_issue_date],
            [role]
        FROM [dbo].[Users]
        WHERE [email] = ?
    `

	// Execute the query
	var user types.User
	row := r.db.QueryRowContext(ctx, query, email)

	// Scan the result into the User struct
	err := row.Scan(
		&user.UserID,
		&user.FirstName,
		&user.SecondName,
		&user.MiddleName,
		&user.DateOfBirth,
		&user.Gender,
		&user.Address,
		&user.PhoneNumber,
		&user.Email,
		&user.PasswordHash,
		&user.SNILS,
		&user.PassportSeries,
		&user.PassportNumber,
		&user.PassportIssuedBy,
		&user.PassportIssueDate,
		&user.Role,
	)

	// Handle potential errors
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, errors.Wrap(err, "failed to execute query")
	}

	// Return the user object
	return &user, nil
}
