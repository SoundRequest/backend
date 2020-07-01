package structure

type SignIn struct {
	Name     string
	Password string
}

type SignUp struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
