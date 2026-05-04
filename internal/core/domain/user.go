package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/PopovMarko/todo_app/internal/core/errors"
)

type User struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

// Response constructor
func NewUser(
	id int,
	version int,
	fullName string,
	phoneNumber *string,
) User {
	return User{
		ID:          id,
		Version:     version,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}

// Request constructor
func NewUserUninitialized(fullName string, PhoneNumber *string) User {
	return User{
		ID:          UninitializedID,
		Version:     UninitializedVersion,
		FullName:    fullName,
		PhoneNumber: PhoneNumber,
	}
}

func (u *User) Validate() error {
	fullNameLength := len([]rune(u.FullName))
	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf("invalid full name length %d: %w",
			fullNameLength,
			core_errors.ErrInvalidArgument,
		)
	}
	if u.PhoneNumber != nil {
		phoneNumberLength := len([]rune(*u.PhoneNumber))
		if phoneNumberLength < 10 || phoneNumberLength > 15 {
			return fmt.Errorf("invalid phone number length %d: %w",
				phoneNumberLength,
				core_errors.ErrInvalidArgument)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.Match([]byte(*u.PhoneNumber)) {
			return fmt.Errorf("invalid phone number format %s: %w",
				*u.PhoneNumber,
				core_errors.ErrInvalidArgument,
			)
		}
	}
	return nil
}

type UserPatch struct {
	FullName    Nullable[string]
	PhoneNumber Nullable[string]
}

func (p *UserPatch) Validate() error {
	if p.FullName.Set && p.FullName.Value == nil {
		return fmt.Errorf("service validate fullName can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate user patch: %w", err)
	}

	tempU := *u

	if patch.FullName.Set {
		tempU.FullName = *patch.FullName.Value
	}
	if patch.PhoneNumber.Set {
		tempU.PhoneNumber = patch.PhoneNumber.Value
	}

	if err := tempU.Validate(); err != nil {
		return fmt.Errorf("Apply patch: %w", err)
	}
	*u = tempU

	return nil
}
