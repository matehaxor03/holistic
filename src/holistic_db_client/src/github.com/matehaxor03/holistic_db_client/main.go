package main

import (
	"fmt"
	"os"
	"strings"
	class "github.com/matehaxor03/holistic_db_client/class"
)

func main() {
	var errors []error 
	const CLS_USER string = "username"
	const CLS_PASSWORD string = "password"
	const CLS_HOST string = "host_name"
	const CLS_PORT string = "port_number"
	const CLS_COMMAND string = "command"
	const CLS_CLASS string = "class"
	const CLS_IF_EXISTS string = "if_exists"
	const CLS_IF_NOT_EXISTS string = "if_not_exists"
	const CLS_DATABASE_NAME string = "database_name"
	const CLS_CHARACTER_SET string = "character_set"
	const CLS_COLLATE string = "collate"



	var CREATE_COMMAND = "CREATE"
	
	var DATABASE_CLASS = "DATABASE"

	//var IF_EXISTS string = "IF EXISTS"
	//var IF_NOT_EXISTS string = "IF NOT EXISTS"
	
    params, errors := getParams(os.Args[1:])
	if errors != nil {
		fmt.Println(fmt.Errorf("%s", errors))
		os.Exit(1)
	}

	host_value, _ := params[CLS_HOST] 
	port_value, _ := params[CLS_PORT] 

	user_value, _ := params[CLS_USER] 
	password_value, _ := params[CLS_PASSWORD] 

	database_name, _ := params[CLS_DATABASE_NAME]
	character_set, _ := params[CLS_CHARACTER_SET]
	collate, _ := params[CLS_COLLATE]

	command_pt, command_found := params[CLS_COMMAND] 
	class_pt, class_found := params[CLS_CLASS]
	
	command_value :=  strings.ToUpper(*command_pt)
	class_value := strings.ToUpper(*class_pt)

	_, if_exists := params[CLS_IF_EXISTS]
	_, if_not_exists := params[CLS_IF_NOT_EXISTS]
	
	if if_exists && if_not_exists {
		errors = append(errors, fmt.Errorf("%s and %s cannot be used together", CLS_IF_EXISTS, CLS_IF_NOT_EXISTS))
	}

	if !command_found {
		errors = append(errors, fmt.Errorf("%s is a mandatory field e.g %s=", CLS_COMMAND, CLS_COMMAND))
	}

	if !class_found {
		errors = append(errors, fmt.Errorf("%s is a mandatory field e.g %s=", CLS_CLASS, CLS_CLASS))
	}

	if errors != nil {
		for _, e := range errors {
			fmt.Println(e)
		}
		os.Exit(1)
	}

	options := make(map[string][]string)
	if if_not_exists {
		options["LOGIC"] = []string{"IF","NOT","EXISTS"}
	}

	if if_exists {
		options["LOGIC"] = []string{"IF", "EXISTS"}
	}

	host := class.NewHost(host_value, port_value)
	credentials :=  class.NewCredentials(user_value, password_value)
	client := class.NewClient(host, credentials, nil)

	if command_value == CREATE_COMMAND {
		if class_value == DATABASE_CLASS {

			database_create_options := class.NewDatabaseCreateOptions(character_set, collate)
			_, shell_output, database_errors := client.CreateDatabase(database_name, database_create_options, options)
			
			if database_errors != nil {
				for _, e := range database_errors {
					fmt.Println(e)
				}

				if shell_output != nil {
					fmt.Println(*shell_output)
				}
				os.Exit(1)
			}
		} else {
			fmt.Printf("class: %s is not supported", class_value)
			os.Exit(1)
		}
	} else {
		fmt.Printf("command: %s is not supported", command_value)
		os.Exit(1)
	}
	
	os.Exit(0)
}

func CreateDatabase(host *class.Host, credentials *class.Credentials, database_name *string, database_create_options *class.DatabaseCreateOptions, options map[string][]string) (*class.Database, *string, []error) {
	return class.NewDatabase(host, credentials, database_name, database_create_options, options).Create()
}

func getParams(params []string) (map[string]*string, []error) {
	var errors []error 
	m := make(map[string]*string)
	for _, value := range params {
		if !strings.Contains(value, "=") {
			m[value] = nil
			continue
		}

		results := strings.SplitN(value, "=", 2)
		if len(results) != 2 {
			errors = append(errors, fmt.Errorf("invalid param found: %s must be in the format {paramName}={paramValue}", value))
			continue
		}
		m[results[0]] = &results[1]
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

