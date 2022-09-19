package class

import (
	"fmt"
)

type Database struct {
    db_name *string
	character_set *string
	collate *string
	CHARACTER_SETS *[]string
	COLLATES *[]string
}

func NewDatabase(db_name *string, character_set *string, collate *string) (*Database) {
	x := Database{db_name: db_name, character_set: character_set, collate: collate}
	
	ARRAY_CHARACTER_SETS := []string{"utf8"}
	x.CHARACTER_SETS = &ARRAY_CHARACTER_SETS

	ARRAY_COLLATES := []string{"utf8_general_ci"}
	x.COLLATES = &ARRAY_COLLATES
	return &x
}

func (this *Database) Validate() []error {
	var errors []error 

	db_name_errs := this.ValidateDatabaseName()

	if db_name_errs != nil {
		errors = append(errors, db_name_errs...)	
	}

	character_set_errs := this.ValidateCharacterSet()

	if character_set_errs != nil {
		errors = append(errors, character_set_errs...)	
	}

	collate_errs := this.ValidateCollate()

	if collate_errs != nil {
		errors = append(errors, collate_errs...)	
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (this *Database) ValidateDatabaseName() ([]error) {
	var VALID_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return this.ValidateCharacters(VALID_CHARACTERS, (*this).db_name)
}

func (this *Database) ValidateCharacterSet() ([]error) {
	return (*this).contains((*this).CHARACTER_SETS, (*this).character_set)
}

func (this *Database) ValidateCollate() ([]error) {
	return (*this).contains((*this).COLLATES, (*this).collate)
}

func (this *Database) ValidateCharacters(whitelist string, str *string) ([]error) {
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

 func (this *Database) contains(s *[]string, str *string) []error {
	for _, v := range *s {
		if v == *str {
			return nil
		}
	}

	var errors []error 
    errors = append(errors, fmt.Errorf("%s does not contain %s", (*s), (*str)))
	return errors
}