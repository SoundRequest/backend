package structure

// SignIn handles db for signin event
type SignIn struct {
	Name     string
	Password string
}

// SignUp handles db for signup event
type SignUp struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UpdatePassword handles db for updatepassword event
type UpdatePassword struct {
	Name   string `json:"name" binding:"required"`
	Origin string `json:"origin" binding:"required"`
	New    string `json:"new" binding:"required"`
}

// SendVerifyPasswordCode handles db for verifypasswordcode event
type SendVerifyPasswordCode struct {
	Email string `json:"email" binding:"required"`
}

// PasswordWithCode handles db for verifypasswordcode event
type PasswordWithCode struct {
	Code string `json:"code" binding:"required"`
	New  string `json:"new" binding:"required"`
}
