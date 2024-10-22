package transform

import (
	"trygo/internal/models"
	"trygo/internal/requests/dto"
)

func UserToResponse(user *models.User) dto.UserResponse {
	return dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
