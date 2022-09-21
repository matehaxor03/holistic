package class


import (
	"fmt"
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
	for _, array_value := range array {
		for _, second_value := range second_array {
			if array_value == second_value {
				return nil
			}
		}
	}

	var errors []error 
    errors = append(errors, fmt.Errorf("%s has value '%s' expected to have value in %s", label, second_array, array))
	return errors
}