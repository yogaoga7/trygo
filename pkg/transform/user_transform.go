package transform

import (
	"trygo/internal/models"
	requests "trygo/internal/requests/dto"
)

// TransformToUser mengubah DTO (Data Transfer Object) dari request menjadi model User
func TransformToUser(userDTO requests.UserDTO) models.User {
	return models.User{
		Name:  userDTO.Name,
		Email: userDTO.Email,
	}
}

// TransformFromUser mengubah model User menjadi format response DTO
func TransformFromUser(user models.User) requests.UserDTO {
	return requests.UserDTO{
		Name:  user.Name,
		Email: user.Email,
	}
}
