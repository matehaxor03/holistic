package class


import (
	"fmt"
	"unicode"
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

func ArrayContainsArray(array []string, second_array []string, label string) []error {
	var errors []error 
	var array_found []string
	
	for _, array_value := range array {
		for _, second_value := range second_array {
			if array_value == second_value {
				array_found = append(array_found, second_value)
			}
		}
	}

	if len(array_found) != len(second_array) {
		errors = append(errors, fmt.Errorf("%s has value '%s' expected to have value in %s", label, second_array, array))
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
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

 func IsUpper(s string) bool {
    for _, r := range s {
        if !unicode.IsUpper(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func IsLower(s string) bool {
    for _, r := range s {
        if !unicode.IsLower(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func GetConstantValueAllowedCharacters() (string) {
	return "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
}

