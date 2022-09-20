package class

import (
	"fmt"
	"bytes"
	"os/exec"
)

type HostSession struct {
	host *Host
	credentials *Credentials
}

func NewHostSession(host *Host, credentials *Credentials) (*HostSession) {
	x := HostSession{host: host, credentials: credentials}
	return &x
}

func (this *HostSession) Validate() []error {
	var errors []error 

	host_errs := (*(*this).host).Validate()

	if host_errs != nil {
		errors = append(errors, host_errs...)	
	}


	creds_errs := (*(*this).GetCredentials()).Validate()

	if creds_errs != nil {
		errors = append(errors, creds_errs...)	
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (this *HostSession) Create_Database_If_Not_Exists(database *Database) []error {
	var errors []error 

	host_session_errs := (*this).Validate()
	database_errs := (*database).Validate()

	if host_session_errs != nil {
		errors = append(errors, host_session_errs...)	
	}

	if database_errs != nil {
		errors = append(errors, database_errs...)	
	}

	if len(errors) > 0 {
		return errors
	}

	credentials := (*(*this).GetCredentials())
	username := *(credentials.GetUsername())
	password := *(credentials.GetPassword())

	command := fmt.Sprintf("mysql -u %s -p %s", username, password)
	command += fmt.Sprintf(" -e \"CREATE DATABASE IF NOT EXISTS %s", (*database).GetDatabaseName())

	character_set := (*database).GetCharacterSet()
	if character_set != nil {
		command += fmt.Sprintf(" CHARACTER SET %s", character_set)
	}

	collate := (*database).GetCollate()
	if collate != nil {
		command += fmt.Sprintf(" COLLATE %s", collate)
	}

	command += ";\""

	// execute mysql command  mysql -u root -pmy_password -D DATABASENAME -e "UPDATE `database` SET `field1` = '1' WHERE `id` = 1111;" > output.txt 
	cmd := exec.Command(command)

    var out bytes.Buffer
    cmd.Stdout = &out

    command_err := cmd.Run()

    if command_err != nil {
		errors = append(errors, command_err)	
	}

	shell_ouput := out.String()
    fmt.Printf("translated phrase: %q\n", shell_ouput)
	return nil
}

func (this *HostSession) GetCredentials() *Credentials {
	return (*this).credentials
}
