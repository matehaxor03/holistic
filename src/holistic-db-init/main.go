package main

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	errors := InitDB("HOLISTIC_DB_ROOT_USERNAME",
		"HOLISTIC_DB_ROOT_PASSWORD",
		"HOLISTIC_DB_USERNAME",
		"HOLISTIC_DB_PASSWORD")
	if errors != nil {
		panic(fmt.Errorf("%s", errors))
	}

}

func InitDB(root_username_env_var string,
	root_password_env_var string,
	username_env_var string,
	password_env_var string) []error {
	var errors []error

	root_db_username, root_db_username_err := validateEnvironmentVariable(root_username_env_var, `^[A-Za-z]+$`)
	if root_db_username_err != nil {
		errors = append(errors, root_db_username_err)
	}

	root_db_password, root_db_password_err := verifyPassword(root_password_env_var)
	if root_db_password_err != nil {
		errors = append(errors, root_db_password_err)
	}

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

	db_connection_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/", root_db_username, root_db_password, db_hostname, db_port_number)
	db, dberr := sql.Open("mysql", db_connection_string)

	if dberr != nil {
		errors = append(errors, dberr)
		defer db.Close()
		return errors
	}

	_, database_creation_err := db.Exec("CREATE DATABASE IF NOT EXISTS " + db_name)
	if database_creation_err != nil {
		fmt.Println("error creating database")
		errors = append(errors, database_creation_err)
		defer db.Close()
		return errors
	}

	_, drop_user_err := db.Exec("DROP USER '" + db_username + "'@'%';")
	if drop_user_err != nil {
		fmt.Println("error dropping user")
		errors = append(errors, drop_user_err)
		defer db.Close()
		return errors
	}

	_, create_user_err := db.Exec("CREATE USER IF NOT EXISTS '" + db_username + "'@'%' IDENTIFIED BY '" + db_password + "'")
	if create_user_err != nil {
		fmt.Println("error creating user")
		errors = append(errors, create_user_err)
		defer db.Close()
		return errors
	}

	_, grant_permissions_err := db.Exec("GRANT ALL ON " + db_name + ".* To '" + db_username + "'@'%'")
	if grant_permissions_err != nil {
		fmt.Println("error granting permissions")
		errors = append(errors, grant_permissions_err)
		defer db.Close()
		return errors
	}

	defer db.Close()
	return nil
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
