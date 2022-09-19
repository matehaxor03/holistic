package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const COMMAND string = "command"
	var COMMANDS = []string{"CREATE"}
	
    params, errors := getParams(os.Args[1:])
	if errors != nil {
		fmt.Println(fmt.Errorf("%s", errors))
		os.Exit(1)
	}

	command_value, found := params[COMMAND] 
	if !found {
		fmt.Printf("%s is a mandatory field", COMMAND)
		os.Exit(1)
	}

	command_value = strings.ToUpper(command_value)
	if !contains(COMMANDS, command_value) {
		fmt.Printf("%s is an invalid command", command_value)
		os.Exit(1)
	}

	os.Exit(0)
}


func getParams(params []string) (map[string]string, []error) {
	var errors []error 
	m := make(map[string]string)
	for _, value := range params {
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

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

