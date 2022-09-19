package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const CLS_COMMAND string = "command"
	const CLS_CLASS string = "class"
	const CLS_IF_EXISTS string = "if_exists"
	const CLS_IF_NOT_EXISTS string = "if_not_exists"
	const CLS_DATABASE_NAME string = "db_name"
	const CLS_CHARACTER_SET string = "character_set"
	const CLS_COLLATE string = "collate"


	var CREATE_COMMAND = "CREATE"
	var COMMANDS = []string{CREATE_COMMAND}
	
	var DATABASE_CLASS = "DATABASE"
	var CLASSES = []string{DATABASE_CLASS}

	//var IF_EXISTS string = "IF EXISTS"
	//var IF_NOT_EXISTS string = "IF NOT EXISTS"

	var CHARACTER_SET_UTF8 = "utf8"
	var CHARACTER_SETS = []string{CHARACTER_SET_UTF8}

	var COLLATE_UTF8_GENERAL_CI = "utf8_general_ci"
	var COLLATES = []string{COLLATE_UTF8_GENERAL_CI}
	
    params, errors := getParams(os.Args[1:])
	if errors != nil {
		fmt.Println(fmt.Errorf("%s", errors))
		os.Exit(1)
	}

	command_value, found := params[CLS_COMMAND] 
	if !found || command_value == "" {
		fmt.Printf("%s is a mandatory field available commands: %s", CLS_COMMAND, strings.Join(COMMANDS,","))
		os.Exit(1)
	}

	command_value = strings.ToUpper(command_value)
	if !contains(COMMANDS, command_value) {
		fmt.Printf("%s is an invalid command available commands: %s", command_value, strings.Join(COMMANDS,",") )
		os.Exit(1)
	}

	class_value, found := params[CLS_CLASS]
	if !found || class_value == ""{
		fmt.Printf("%s is a mandatory field available classes %s", CLS_CLASS,  strings.Join(CLASSES,","))
		os.Exit(1)
	} 

	class_value = strings.ToUpper(class_value)

	_, if_exists_found := params[CLS_IF_EXISTS]
	_, if_not_exists_found := params[CLS_IF_NOT_EXISTS]


	if if_exists_found && if_not_exists_found {
		fmt.Printf("%s and %s cannot be used together", CLS_IF_EXISTS, CLS_IF_NOT_EXISTS)
		os.Exit(1)
	}

	if command_value == CREATE_COMMAND {
		if class_value == DATABASE_CLASS {
			
			db_name_value, found := params[CLS_DATABASE_NAME]
			if !found {
				fmt.Printf("%s is a mandatory field with command: %s %s", CLS_DATABASE_NAME, CREATE_COMMAND, DATABASE_CLASS)
				os.Exit(1)
			} 

			db_name_errs := validateDatabaseName(db_name_value)
			if db_name_errs != nil {
				fmt.Println(fmt.Errorf("%s contains invalid chracters: %s", CLS_DATABASE_NAME, db_name_errs))
				os.Exit(1)
			}

			character_set_value, character_set_found := params[CLS_CHARACTER_SET]
			if character_set_found && !contains(CHARACTER_SETS, character_set_value) { 
				fmt.Printf("%s is an invalid value for %s available values: %s", character_set_value, CLS_CHARACTER_SET, strings.Join(CHARACTER_SETS,",") )
				os.Exit(1)
			} 


			collate_value, colalte_found := params[CLS_COLLATE]
			if colalte_found && !contains(COLLATES, collate_value) { 
				fmt.Printf("%s is an invalid value for %s available values: %s", collate_value, CLS_COLLATE, strings.Join(COLLATES,",") )
				os.Exit(1)
			} 

			






		} else {
			fmt.Printf("class: %s is not supported", class_value)
			os.Exit(1)
		}

	}
	






	os.Exit(0)
}

func validateDatabaseName(db_name string) ([]error) {
	var VALID_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return validateCharacters(VALID_CHARACTERS, db_name)
}

func getParams(params []string) (map[string]string, []error) {
	var errors []error 
	m := make(map[string]string)
	for _, value := range params {
		if !strings.Contains(value, "=") {
			m[value] = ""
			continue
		}

		results := strings.SplitN(value, "=", 2)
		if len(results) != 2 {
			errors = append(errors, fmt.Errorf("invalid param found: %s must be in the format {paramName}={paramValue}", value))
			continue
		}
		m[results[0]] = results[1]
	}

	if len(errors) > 0 {
		return nil, errors
	}
 
	return m, nil
}

func validateCharacters(whitelist string, str string) ([]error) {
	var errors []error 
	for _, letter := range str {
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

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

