package users_transport_http

import "github.com/PopovMarko/todo_app/internal/core/domain"

// DTO for get user from service layer and send to http
type UserDTOResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

// Helper func to connect domain and transport without
// importing each other
func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTOResponse {
	usersDTOs := make([]UserDTOResponse, len(users))
	for i, user := range users {
		userDTO := userDTOFromDomain(user)
		usersDTOs[i] = userDTO
	}
	return usersDTOs
}
