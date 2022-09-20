package class

import (
	"fmt"
	"reflect"
	"unicode"
)

type DatabaseCreateOptions struct {
	character_set *string
	collate *string
	CHARACTER_SETS []string
	COLLATES []string
}

func NewDatabaseCreateOptions(character_set *string, collate *string) (*DatabaseCreateOptions) {
	x := DatabaseCreateOptions{character_set: character_set, collate: collate}
	
	x.CHARACTER_SETS = []string{"utf8"}
	x.COLLATES = []string{"utf8_general_ci"}
	
	return &x
}

func (this *DatabaseCreateOptions) Validate() []error {
	var errors []error 
	e := reflect.ValueOf(this).Elem()
	
    for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name

		if varName == "character_set" {
			character_set_errs := (*this).validateCharacterSet()

			if character_set_errs != nil {
				errors = append(errors, character_set_errs...)	
			}
		} else if varName == "collate" {
			collate_errs :=  (*this).validateCollate()

			if collate_errs != nil {
				errors = append(errors, collate_errs...)	
			}
		} else {
			if !IsUpper(varName) {
				errors = append(errors, fmt.Errorf("%s field is not being validated for Crendentials", varName))	
			}
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