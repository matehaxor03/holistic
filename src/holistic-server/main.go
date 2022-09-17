package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
)

func ProcessRequest(w http.ResponseWriter, req *http.Request) {
	//w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(formatRequest(req)))

	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func verifyPassword(password string) error {
	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	var specialCharPresent bool
	const minPassLength = 8
	const maxPassLength = 64
	var passLen int
	var errorString string

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
		return fmt.Errorf(errorString)
	}
	return nil
}

func main() {
	db_username_regex := `^[A-Za-z]+$`
	db_username_regex_matcher := regexp.MustCompile(db_username_regex).MatchString
	db_username := os.Getenv("HOLISTIC_DB_USERNAME")
	if !db_username_regex_matcher(db_username) {
		res := fmt.Sprintf("HOLISTIC_DB_USERNAME environment variable contains invalid characters: %s regex: %s", db_username, db_username_regex)
		panic(res)
	}

	db_password := os.Getenv("HOLISTIC_DB_PASSWORD")
	db_password_err := verifyPassword(db_password)
	if db_password_err != nil {
		res := fmt.Sprintf("HOLISTIC_DB_PASSWORD did not meet the requirements: %s", db_password_err)
		panic(res)
	}

	db_hostname_regex := `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	db_hostname_regex_matcher := regexp.MustCompile(db_hostname_regex).MatchString
	db_hostname := os.Getenv("HOLISTIC_DB_HOSTNAME")
	if !db_hostname_regex_matcher(db_hostname) {
		res := fmt.Sprintf("HOLISTIC_DB_HOSTNAME environment variable contains invalid characters: %s regex: %s", db_hostname, db_hostname_regex)
		panic(res)
	}

	db_port_number_regex := `\d+`
	db_port_number_regex_matcher := regexp.MustCompile(db_port_number_regex).MatchString
	db_port_number := os.Getenv("HOLISTIC_DB_PORT_NUMBER")
	if !db_port_number_regex_matcher(db_port_number) {
		res := fmt.Sprintf("HOLISTIC_DB_PORT_NUMBER environment variable contains invalid characters: %s regex: %s", db_port_number, db_port_number_regex)
		panic(res)
	}

	db_name_regex := `^[A-Za-z]+$`
	db_name_regex_matcher := regexp.MustCompile(db_name_regex).MatchString
	db_name := os.Getenv("HOLISTIC_DB_NAME")
	if !db_name_regex_matcher(db_name) {
		res := fmt.Sprintf("HOLISTIC_DB_NAME environment variable contains invalid characters: %s regex: %s", db_name, db_name_regex)
		panic(res)
	}

	db_connection_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_username, db_password, db_hostname, db_port_number, db_name)
	db, dberr := sql.Open("mysql", db_connection_string)
	if dberr != nil {
		panic(dberr.Error())
	}

	version, version_err := db.Query("SELECT VERSION()")
	if version_err != nil {
		panic(version_err.Error())
	}

	fmt.Println(version)

	defer db.Close()
	fmt.Println("Success!")

	buildHandler := http.FileServer(http.Dir("static"))
	http.Handle("/", buildHandler)

	http.HandleFunc("/api", ProcessRequest)

	err := http.ListenAndServeTLS(":5000", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
