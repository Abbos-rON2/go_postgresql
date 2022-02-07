package models

type Teacher struct {
	Guid        string  `json:"guid"`
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	BirthDate   string  `json:"birth_date"`
	Address     string  `json:"address"`
	Salary      float64 `json:"salary"`
	BranchID    string  `json:"branch_id"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type CreateTeacher struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	BirthDate   string  `json:"birth_date"`
	Address     string  `json:"address"`
	Salary      float64 `json:"salary"`
	BranchID    string  `json:"branch_id"`
}

type UpdateTeacher struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	BirthDate   string  `json:"birth_date"`
	Address     string  `json:"address"`
	Salary      float64 `json:"salary"`
	BranchID    string  `json:"branch_id"`
}
