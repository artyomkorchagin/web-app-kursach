package types

type User struct {
	UserID            string `json:"user_id"`
	FirstName         string `json:"first_name"`
	SecondName        string `json:"second_name"`
	MiddleName        string `json:"middle_name"`
	DateOfBirth       string `json:"date_of_birth"`
	Gender            string `json:"gender"`
	Address           string `json:"address"`
	PhoneNumber       string `json:"phone_number"`
	Email             string `json:"email"`
	PasswordHash      string `json:"-"`
	SNILS             string `json:"snils"`
	PassportSeries    string `json:"passport_series"`
	PassportNumber    string `json:"passport_number"`
	PassportIssuedBy  string `json:"passport_issued_by"`
	PassportIssueDate string `json:"passport_issue_date"`
	Role              string `json:"role"`
}

type CreateUserRequest struct {
	FirstName         string `form:"first_name" binding:"required"`
	SecondName        string `form:"second_name" binding:"required"`
	MiddleName        string `form:"middle_name"`
	DateOfBirth       string `form:"date_of_birth" binding:"required"`
	Gender            string `form:"gender" binding:"required"`
	Address           string `form:"address"`
	PhoneNumber       string `form:"phone_number" binding:"required"`
	Email             string `form:"email" binding:"required"`
	Password          string `form:"password" binding:"required"`
	SNILS             string `form:"snils" binding:"omitempty"`
	PassportSeries    string `form:"passport_series" binding:"omitempty"`
	PassportNumber    string `form:"passport_number" binding:"omitempty"`
	PassportIssuedBy  string `form:"passport_issued_by" binding:"omitempty"`
	PassportIssueDate string `form:"passport_issue_date" binding:"omitempty"`
	Role              string `form:"role" binding:"omitempty"`
}

func NewUser(req CreateUserRequest) *User {

	user := &User{
		FirstName:         req.FirstName,
		SecondName:        req.SecondName,
		MiddleName:        req.MiddleName,
		DateOfBirth:       req.DateOfBirth,
		Gender:            req.Gender,
		Address:           req.Address,
		PhoneNumber:       req.PhoneNumber,
		Email:             req.Email,
		PasswordHash:      req.Password,
		SNILS:             req.SNILS,
		PassportSeries:    req.PassportSeries,
		PassportNumber:    req.PassportNumber,
		PassportIssuedBy:  req.PassportIssuedBy,
		PassportIssueDate: req.PassportIssueDate,
		Role:              req.Role,
	}

	return user
}
