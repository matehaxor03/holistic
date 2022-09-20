package class

import (
	"fmt"
)

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

	host_errs := (*this).ValidateHostname()

	if host_errs != nil {
		errors = append(errors, host_errs...)	
	}


	port_errs :=  (*this).ValidatePort()

	if port_errs != nil {
		errors = append(errors, port_errs...)	
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (this *Host) ValidateHostname() ([]error) {
	var VALID_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789."
	var errors []error 

	if (*this).host_name == nil || *((*this).host_name) == "" {
		errors = append(errors, fmt.Errorf("host_name cannot have an empty value"))
		return errors
	}

	return this.ValidateCharacters(VALID_CHARACTERS, (*this).host_name)
}

func (this *Host) ValidatePort() ([]error) {
	var VALID_CHARACTERS = "1234567890"
	var errors []error 

	if (*this).port_number == nil || *((*this).port_number) == "" {
		errors = append(errors, fmt.Errorf("port_number cannot have an empty value"))
		return errors
	}

	return this.ValidateCharacters(VALID_CHARACTERS, (*this).port_number)
}

func (this *Host) ValidateCharacters(whitelist string, str *string) ([]error) {
	var errors []error 
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

 func (this *Host) GetHostName() (*string) {
	return (*this).host_name
 }

 func (this *Host) GetPortNumber() (*string) {
	return (*this).port_number
 }