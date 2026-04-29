package users_postgres_repository

import "github.com/PopovMarko/todo_app/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func usersDomainsFromModels(models []UserModel) []domain.User {
	domainsUsers := make([]domain.User, len(models))
	for i, model := range models {
		domainsUsers[i] = domain.NewUser(
			model.ID,
			model.Version,
			model.FullName,
			model.PhoneNumber,
		)
	}

	return domainsUsers
}

func userDomainFromModel(model UserModel) domain.User {
	userDomain := domain.NewUser(
		model.ID,
		model.Version,
		model.FullName,
		model.PhoneNumber,
	)

	return userDomain
}
