package class

import (
	"fmt"
)

type Credentials struct {
	user *string
	password *string
}

func NewCredentials(user *string, password *string) (*Credentials) {
	x := Credentials{user: user,
			    password: password}

	return &x
}

func (this *Credentials) Validate() []error {
	var errors []error 

	user_errs := (*this).ValidateUser()

	if user_errs != nil {
		errors = append(errors, user_errs...)	
	}

	password_errs := (*this).ValidatePassword()

	if password_errs != nil {
		errors = append(errors, password_errs...)	
	}
	
	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (this *Credentials) ValidateUser() ([]error) {
	var VALID_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return this.ValidateCharacters(VALID_CHARACTERS, (*this).user)
}

 func (this *Credentials) ValidatePassword() ([]error) {
	var VALID_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789="
	return this.ValidateCharacters(VALID_CHARACTERS, this.password)
}

func (this *Credentials) ValidateCharacters(whitelist string, str *string) ([]error) {
	var errors []error 
	if str == nil {
		errors = append(errors, fmt.Errorf("string is nil"))
		return errors
	}

	if *str == "" {
		errors = append(errors, fmt.Errorf("string is empty"))
		return errors
	}

	for _, letter := range *str {
		found := false

		for _, check := range whitelist {
			if check == letter {
				found = true
				break
			}
		}

		if !found {
			errors = append(errors, fmt.Errorf("invalid letter detected %s", string(letter)))
		}
	}
	
	if len(errors) > 0 {
		return errors
	}

	return nil
 }