package class

type DatabaseCreateOptions struct {
	character_set *string
	collate *string
	CHARACTER_SETS []string
	COLLATES []string
}

func NewDatabaseCreateOptions(character_set *string, collate *string) (*DatabaseCreateOptions, []error) {
	x := DatabaseCreateOptions{character_set: character_set, collate: collate}
	
	x.CHARACTER_SETS = []string{"utf8"}
	x.COLLATES = []string{"utf8_general_ci"}
	
	errors := x.Validate()
	if errors != nil {
		return nil, errors
	}

	return &x, nil
}

func (this *DatabaseCreateOptions) Validate() []error {
	var errors []error 

	character_set_errs := (*this).validateCharacterSet()

	if character_set_errs != nil {
		errors = append(errors, character_set_errs...)	
	}

	collate_errs :=  (*this).validateCollate()

	if collate_errs != nil {
		errors = append(errors, collate_errs...)	
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
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