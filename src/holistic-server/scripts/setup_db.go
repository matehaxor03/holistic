package scripts

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func verifyPassword(password_env_var string) (string, error) {
	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	var specialCharPresent bool
	const minPassLength = 8
	const maxPassLength = 64
	var passLen int
	var errorString string

	password := os.Getenv(password_env_var)

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
			passLen++
		case unicode.IsUpper(ch):
			uppercasePresent = true
			passLen++
		case unicode.IsLower(ch):
			lowercasePresent = true
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
			passLen++
		}
	}

	appendError := func(err string) {
		if len(strings.TrimSpace(errorString)) != 0 {
			errorString += ", " + err
		} else {
			errorString = err
		}
	}

	if !lowercasePresent {
		appendError("lowercase letter missing")
	}
	if !uppercasePresent {
		appendError("uppercase letter missing")
	}
	if !numberPresent {
		appendError("at least one numeric character required")
	}
	if !specialCharPresent {
		appendError("special character missing")
	}
	if !(minPassLength <= passLen && passLen <= maxPassLength) {
		appendError(fmt.Sprintf("password length must be between %d to %d characters long", minPassLength, maxPassLength))
	}

	if len(errorString) != 0 {
		return "", fmt.Errorf("%s %s", password_env_var, errorString)
	}
	return password, nil
}

func validateEnvironmentVariable(environmentVariableName string, regex string) (string, error) {
	regex_matcher, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}

	value := os.Getenv(environmentVariableName)
	if !regex_matcher.MatchString(value) {
		return "", fmt.Errorf("%s environment variable contains invalid characters: %s regex: %s", environmentVariableName, value, regex)
	}

	return value, nil
}

func InitDB(username_env_var string, password_env_var string) []error {
	var errors []error

	db_username, db_username_err := validateEnvironmentVariable(username_env_var, `^[A-Za-z]+$`)
	if db_username_err != nil {
		errors = append(errors, db_username_err)
	}

	db_password, db_password_err := verifyPassword(password_env_var)
	if db_password_err != nil {
		errors = append(errors, db_password_err)
	}

	db_hostname, db_hostname_err := validateEnvironmentVariable("HOLISTIC_DB_HOSTNAME", `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if db_hostname_err != nil {
		errors = append(errors, db_hostname_err)
	}

	db_port_number, db_port_number_err := validateEnvironmentVariable("HOLISTIC_DB_PORT_NUMBER", `\d+`)
	if db_port_number_err != nil {
		errors = append(errors, db_port_number_err)
	}

	db_name, db_name_err := validateEnvironmentVariable("HOLISTIC_DB_NAME", `^[A-Za-z]+$`)
	if db_name_err != nil {
		errors = append(errors, db_name_err)
	}

	if len(errors) > 0 {
		return errors
	}

	db_connection_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_username, db_password, db_hostname, db_port_number, db_name)
	db, dberr := sql.Open("mysql", db_connection_string)

	if dberr != nil {
		errors = append(errors, dberr)
		defer db.Close()
		return errors
	}

	_, version_err := db.Query("SELECT VERSION()")
	if version_err != nil {
		errors = append(errors, version_err)
		defer db.Close()
		return errors
	}

	defer db.Close()
	return nil
}
