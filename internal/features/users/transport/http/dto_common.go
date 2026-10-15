package users_transport_http

import "github.com/stanislav-backend/golang-todoapp/internal/core/domain"

type UserDTOResponse struct {
	ID          int     `json:"id"          example:"10"`
	Version     int     `json:"version"     example:"1"`
	FullName    string  `json:"full_name"   example:"Stas Timoshenko"`
	PhoneNumber *string `json:"phone_number" example:"+79180905602"`
}

func UserDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}


func UsersDTOFromDomains(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))

	for i, user := range users{
		usersDTO[i] = UserDTOFromDomain(user)
	}

	return usersDTO
}