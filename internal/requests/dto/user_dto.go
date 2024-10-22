package requests

// UserDTO digunakan untuk request dan response terkait user
type UserDTO struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
