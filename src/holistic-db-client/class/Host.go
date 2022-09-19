package class

import (
	"fmt"
)

type Host struct {
    host *string
	port *string
}

func NewHost(host *string, port *string) (*Host) {
	x := Host{host: host, port: port}

	return &x
}

func (this *Host) Validate() []error {
	var errors []error 

	host_errs := this.ValidateHost()

	if host_errs != nil {
		errors = append(errors, host_errs...)	
	}


	port_errs := this.ValidatePort()

	if port_errs != nil {
		errors = append(errors, port_errs...)	
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (this *Host) ValidateHost() ([]error) {
	var VALID_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789."
	var errors []error 

	if (*this).host == nil || *((*this).host) == "" {
		errors = append(errors, fmt.Errorf("host cannot have an empty value"))
		return errors
	}

	return this.ValidateCharacters(VALID_CHARACTERS, (*this).host)
}

func (this *Host) ValidatePort() ([]error) {
	var VALID_CHARACTERS = "1234567890"
	var errors []error 

	if (*this).port == nil || *((*this).port) == "" {
		errors = append(errors, fmt.Errorf("port cannot have an empty value"))
		return errors
	}

	return this.ValidateCharacters(VALID_CHARACTERS, (*this).port)
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