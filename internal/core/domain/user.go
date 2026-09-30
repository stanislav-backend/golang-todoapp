package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/stanislav-backend/golang-todoapp/internal/core/errors"
)

type User struct {
	ID      int
	Version int

	FullName    string
	PhoneNumber *string
}

func NewUser(id int, version int, full_name string, phone_number *string) User {
	return User{
		ID:          id,
		Version:     version,
		FullName:    full_name,
		PhoneNumber: phone_number,
	}
}

func NewUserUninitialized(FullName string, PhoneNumber *string) User {
	return NewUser(
		UninitializedID,
		UninitializedVersion,
		FullName,
		PhoneNumber,
	)
}

func (u *User) Validate() error {
	fullNameLength := len([]rune(u.FullName))

	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf("Invalid `FullName` len: %d: %w", fullNameLength, core_errors.ErrInvalidArgument)
	}

	if u.PhoneNumber != nil{
		phoneNumberLength := len([]rune(*u.PhoneNumber))
		if phoneNumberLength < 10 || phoneNumberLength > 15 {
			return fmt.Errorf("Invalid `PhoneNumber` len: %d: %w", phoneNumberLength, core_errors.ErrInvalidArgument)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*u.PhoneNumber){
			return fmt.Errorf("Invalid `PhoneNumber` len: %w", core_errors.ErrInvalidArgument)
		}
	}


	return nil
}


type UserPatch struct{
	FullName Nullable[string]
	PhoneNumber Nullable[string]
}

func (p *UserPatch) Validate() error{
	if p.FullName.Set && p.FullName.Value == nil{
		return fmt.Errorf("`FullName` can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error{
	if err := patch.Validate(); err != nil{
		return fmt.Errorf("validate user patch: %w", err)
	}

	tmp := *u

	if patch.FullName.Set{
		tmp.FullName = *patch.FullName.Value
	}

	if patch.PhoneNumber.Set{
		tmp.PhoneNumber = patch.PhoneNumber.Value
	}

	if err := tmp.Validate(); err != nil{
		return fmt.Errorf("validate patch user: %w", err)
	}

	*u = tmp
	return nil
}