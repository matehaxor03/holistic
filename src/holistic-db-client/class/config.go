package class

import (
	"fmt"
)

type config struct {
    host string
	user string
	password string
}

func NewConfig(host string, user string, password string) *config {
	x := config{host: host, 
                user: user,
			    password: password}
	return &x
}

func (this config) Validate() []error {
	var errors []error 

	host_errs := this.ValidateHost()

	if host_errs != nil {
		errors = append(errors, host_errs...)	
	}

	user_errs := this.ValidateUser()

	if user_errs != nil {
		errors = append(errors, user_errs...)	
	}

	password_errs := this.ValidatePassword()

	if password_errs != nil {
		errors = append(errors, password_errs...)	
	}
	
	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (this config) ValidateUser() ([]error) {
	var VALID_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var errors []error 

	if this.user == "" {
		errors = append(errors, fmt.Errorf("user cannot have an empty value"))
		return errors
	}

	return this.ValidateCharacters(VALID_CHARACTERS, this.user)
}

 func (this config) ValidatePassword() ([]error) {
	var VALID_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789="
	var errors []error 

	if this.password == "" {
		errors = append(errors, fmt.Errorf("password cannot have an empty value"))
		return errors
	}

	return this.ValidateCharacters(VALID_CHARACTERS, this.password)
}

func (this config) ValidateHost() ([]error) {
	var VALID_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789."
	var errors []error 

	if this.host == "" {
		errors = append(errors, fmt.Errorf("host cannot have an empty value"))
		return errors
	}

	return this.ValidateCharacters(VALID_CHARACTERS, this.host)
}

func (this config) ValidateCharacters(whitelist string, str string) ([]error) {
	var errors []error 
	for _, letter := range str {
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