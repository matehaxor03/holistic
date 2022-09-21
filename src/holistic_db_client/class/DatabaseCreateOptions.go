package class

import (
	"fmt"
	"reflect"
)

type DatabaseCreateOptions struct {
	character_set *string
	collate *string
	CHARACTER_SETS []string
	COLLATES []string

	validation_functions map[string]func() []error
}

func NewDatabaseCreateOptions(character_set *string, collate *string) (*DatabaseCreateOptions) {
	x := DatabaseCreateOptions{character_set: character_set, collate: collate}
	
	x.CHARACTER_SETS = []string{"utf8"}
	x.COLLATES = []string{"utf8_general_ci"}
	
	x.validation_functions = make(map[string]func() []error)
	x.InitValidationFunctions()
	return &x
}

func (this *DatabaseCreateOptions) InitValidationFunctions() ()  {
	validation_functions := (*this).getValidationFunctions()
	validation_functions["validateCharacterSet"] = (*this).validateCharacterSet
	validation_functions["validateCollate"] = (*this).validateCollate
	validation_functions["validateValidationFunctions"] = (*this).validateValidationFunctions
	validation_functions["validateConstants"] = (*this).validateConstants


	if validation_functions["validateValidationFunctions"] == nil|| 
	   GetFunctionName(validation_functions["validateValidationFunctions"]) != GetFunctionName((*this).validateValidationFunctions) {
		panic(fmt.Errorf("validateValidationFunctions validation method not found potential sql injection without it"))
	}

	if validation_functions["validateConstants"] == nil|| 
	   GetFunctionName(validation_functions["validateConstants"]) != GetFunctionName((*this).validateConstants) {
		panic(fmt.Errorf("validateConstants validation method not found potential sql injection without it"))
	}
}


func (this *DatabaseCreateOptions) validateCharacterSet() ([]error) {
	if (*this).character_set == nil {
		return nil
	}

	return Contains((*this).CHARACTER_SETS, (*this).character_set, "character_set")
}

func (this *DatabaseCreateOptions) validateCollate() ([]error) {
	if (*this).collate == nil {
		return nil
	}

	return Contains((*this).COLLATES, (*this).collate, "collate")
}

func (this *DatabaseCreateOptions) GetCharacterSet() *string {
	return (*this).character_set
}

func (this *DatabaseCreateOptions) GetCollate() *string {
	return (*this).collate
}

func (this *DatabaseCreateOptions) Validate() []error {
	var errors []error 
	var fieldsNotFound []string
	reflected_value := reflect.ValueOf(this)
	refected_element := reflected_value.Elem()
	
    for i := 0; i < refected_element.NumField(); i++ {
		field := refected_element.Type().Field(i)
		validationMethodName := GetValidationMethodNameForFieldName(field.Name)

		method, found_method := (*this).getValidationFunctions()[validationMethodName]
		if !found_method {
			fieldsNotFound = append(fieldsNotFound, field.Name)
		} else {
			relection_errors := method()
			if relection_errors != nil{
				errors = append(errors, relection_errors...)
			}
		}
	}

	method, found_method := (*this).getValidationFunctions()["validateConstants"]
	if !found_method {
		errors = append(errors, fmt.Errorf("validation method: validateConstants not found please add to InitValidationFunctions"))
	} else {
		constant_errors := method()
		if constant_errors != nil{
			errors = append(errors, constant_errors...)
		}
	}

	for _, value := range fieldsNotFound {
		if !IsUpper(value) {
			errors = append(errors, fmt.Errorf("validation method: %s not found for %s please add to InitValidationFunctions", GetValidationMethodNameForFieldName(value), value))	
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (this *DatabaseCreateOptions) getValidationFunctions() map[string]func() []error {
	return (*this).validation_functions
}

func (this *DatabaseCreateOptions) validateConstants()  ([]error) {
	var errors []error 
	VALID_CHARACTERS := GetConstantValueAllowedCharacters()
	reflected_value := reflect.ValueOf(this)
	refected_element := reflected_value.Elem()
	string_fieldValue := ""

	for i := 0; i < refected_element.NumField(); i++ {
		string_fieldValue = ""
		field := refected_element.Type().Field(i)
		fieldName := field.Name
		if !IsUpper(fieldName) {
			continue
		}

		fieldValue := refected_element.FieldByName(fieldName)		
		if fieldValue.Kind().String() == "string" {
			string_fieldValue = fmt.Sprintf("%s", fieldValue)	
		} else if fieldValue.Kind().String() == "slice" {
			var array = fieldValue.Interface().([]string)
			for _, value := range array {
				string_fieldValue += fmt.Sprintf("%s", value)
			}
		} else {
			panic(fmt.Sprintf("please implement validation for constant value %s", fieldName))
		}

		character_errors := ValidateCharacters(VALID_CHARACTERS, &string_fieldValue, fieldName)
		if character_errors != nil {
			errors = append(errors, character_errors...)
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (this *DatabaseCreateOptions) validateValidationFunctions() ([]error) {
	var errors []error 
	current := (*this).getValidationFunctions()
	compare := make(map[string]func() []error)
	found := false

    for current_key, current_value := range current {
		found = false
		for compare_key, compare_value := range compare {
			if GetFunctionName(current_value) == GetFunctionName(compare_value) && 
			   current_key != compare_key {
				found = true
				errors = append(errors, fmt.Errorf("key %s and key %s contain duplicate validation functions %s",  current_key, compare_key, current_value))
				break
			}
		}

		if !found {
			compare[current_key] = current_value
		}
    }

	if len(errors) > 0 {
		return errors
	}

	return nil
}