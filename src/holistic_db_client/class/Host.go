package class

import (
	"fmt"
	"reflect"
)

func Contains(array []string, str *string, label string) []error {
	for _, array_value := range array {
		if array_value == *str {
			return nil
		}
	}

	var errors []error 
    errors = append(errors, fmt.Errorf("%s has value '%s' expected to have value in %s", label, (*str) , array))
	return errors
}

func ValidateCharacters(whitelist string, str *string, label string) ([]error) {
	var errors []error 

	if str == nil {
		errors = append(errors, fmt.Errorf("%s is nil", label))
		return errors
	}

	if *str == "" {
		errors = append(errors, fmt.Errorf("%s is empty", label))
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
			errors = append(errors, fmt.Errorf("invalid letter %s for %s please use %s", string(letter), label, whitelist))
		}
	}
	
	if len(errors) > 0 {
		return errors
	}

	return nil
 }

type Host struct {
    host_name *string
	port_number *string
}

func NewHost(host_name *string, port_number *string) (*Host) {
	x := Host{host_name: host_name, port_number: port_number}

	return &x
}

func (this *Host) Validate() []error {
	var errors []error 
	e := reflect.ValueOf(this).Elem()
	
    for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name

		if varName == "host_name" {
			host_errs := (*this).validateHostname()

			if host_errs != nil {
				errors = append(errors, host_errs...)	
			}
		} else if varName == "port_number" {
			port_errs :=  (*this).validatePort()

			if port_errs != nil {
				errors = append(errors, port_errs...)	
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



func (this *Host) validateHostname() ([]error) {
	var VALID_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789."
	return ValidateCharacters(VALID_CHARACTERS, (*this).GetHostName(), "host_name")
}

func (this *Host) validatePort() ([]error) {
	var VALID_CHARACTERS = "1234567890"
	return ValidateCharacters(VALID_CHARACTERS, (*this).GetPortNumber(), "port")
}

 func (this *Host) GetHostName() (*string) {
	return (*this).host_name
 }

 func (this *Host) GetPortNumber() (*string) {
	return (*this).port_number
 }

 func (this *Host) GetCLSCommand() (*string, []error) {
	errors := (*this).Validate()
	if len(errors) > 0 {
		return nil, errors
	}

	command := fmt.Sprintf("--host=%s --port=%s --protocol=TCP ", (*(*this).GetHostName()), (*(*this).GetPortNumber()))

	return &command, nil
 }