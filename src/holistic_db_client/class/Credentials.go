package class

import (
	"fmt"
	"reflect"
)

type Credentials struct {
	username *string
	password *string
}

func NewCredentials(username *string, password *string) (*Credentials) {
	x := Credentials{username: username,
			    password: password}

	return &x
}

func (this *Credentials) Validate() []error {
	var errors []error 

	e := reflect.ValueOf(this).Elem()
	
    for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		
		if varName == "username" {
			user_errs := (*this).ValidateUsername()
			if user_errs != nil {
				errors = append(errors, user_errs...)	
			}
		} else if varName == "password" {
			password_errs := (*this).ValidatePassword()

			if password_errs != nil {
				errors = append(errors, password_errs...)	
			}
		} else {
			errors = append(errors, fmt.Errorf("%s field is not being validated for Crendentials", varName))	
		}
	}
		
	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (this *Credentials) ValidateUsername() ([]error) {
	var VALID_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return ValidateCharacters(VALID_CHARACTERS, (*this).GetUsername(), "username")
}

 func (this *Credentials) ValidatePassword() ([]error) {
	var VALID_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789="
	return ValidateCharacters(VALID_CHARACTERS, (*this).GetPassword(), "password")
}

 func (this *Credentials) GetUsername() *string {
	return (*this).username
 }

 func (this *Credentials) GetPassword() *string {
	return (*this).password
 }

 func (this *Credentials) GetCLSCommand() (*string, []error) {
	errors := (*this).Validate()
	if len(errors) > 0 {
		return nil, errors
	}

	command := fmt.Sprintf("--user=%s --password=%s ", (*(*this).GetUsername()), (*(*this).GetPassword()))

	return &command, nil
 }

