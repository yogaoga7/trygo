package dto

type UserResponse struct {
	ID    uint   `json:"ID"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}
