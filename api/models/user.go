package models

type LoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Guid        string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Firstname   string `json:"first_name"`
	Lastname    string `json:"last_name"`
	IsActive    bool   `json:"is_active"`
	BirthDate   string `json:"birth_date"`
	Role        string `json:"role" enums:"super_admin, branch_admin"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateUser struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Firstname   string `json:"first_name"`
	Lastname    string `json:"last_name"`
	Address     string `json:"address"`
	IsActive    bool   `json:"is_active"`
	BirthDate   string `json:"birth_date"`
	Role        string `json:"role" enums:"super_admin, branch_admin"`
}

type UpdateUser struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Firstname   string `json:"first_name"`
	Lastname    string `json:"last_name"`
	IsActive    bool   `json:"is_active"`
	BirthDate   string `json:"birth_date"`
	Role        string `json:"role" enums:"super_admin, branch_admin"`
}

type UpdateUserPassword struct {
	Username    string `json:"username"`
	PasswordOld string `json:"password_old"`
	PasswordNew string `json:"password_new"`
}
