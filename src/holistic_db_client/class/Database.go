package class

import (
	"fmt"
	"bytes"
	"os/exec"
)

type Database struct {
	host *Host
	credentials *Credentials
    database_name *string
	database_create_options *DatabaseCreateOptions
	extra_options map[string]string
	
	DATA_DEFINITION_STATEMENT_CREATE string
	DATA_DEFINITION_STATEMENTS []string

	LOGIC_OPTION_FIELD string
	LOGIC_OPTION_IF_NOT_EXISTS string
	LOGIC_OPTION_CREATE_OPTIONS []string
}

func NewDatabase(host *Host, credentials *Credentials, database_name *string, database_create_options *DatabaseCreateOptions, extra_options map[string]string) (*Database) {
	x := Database{host: host, credentials: credentials, database_name: database_name, database_create_options: database_create_options, extra_options: extra_options}
	
	x.DATA_DEFINITION_STATEMENT_CREATE = "CREATE"
	x.DATA_DEFINITION_STATEMENTS = []string{x.DATA_DEFINITION_STATEMENT_CREATE}
	
	x.LOGIC_OPTION_FIELD = "logic"
	x.LOGIC_OPTION_IF_NOT_EXISTS = "IF NOT EXISTS"
	x.LOGIC_OPTION_CREATE_OPTIONS = []string{x.LOGIC_OPTION_IF_NOT_EXISTS}
	
	return &x
}

func (this *Database) Create() (*Database, *string, []error)  {
	this, result, errors := (*this).createDatabase()
	if errors != nil {
		return nil, result, errors
	}

	return this, result, nil
}

func (this *Database) Validate() []error {
	var errors []error 

	host_errs := (*this).validateHost()

	if host_errs != nil {
		errors = append(errors, host_errs...)	
	}

	credentials_errs := (*this).validateCredentials()

	if credentials_errs != nil {
		errors = append(errors, credentials_errs...)	
	}

	db_name_errs := this.validateDatabaseName()

	if db_name_errs != nil {
		errors = append(errors, db_name_errs...)	
	}
	
	database_create_options_errs := ((*this).GetDatabaseCreateOptions()).Validate()

	if database_create_options_errs != nil {
		errors = append(errors, database_create_options_errs...)	
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (this *Database) validateHost()  ([]error) {
	var errors []error 
	if (*this).GetHost() == nil {
		errors = append(errors, fmt.Errorf("host is nil"))
		return errors
	}

	return (*((*this).GetHost())).Validate()
}

func (this *Database) validateCredentials()  ([]error) {
	var errors []error 
	if (*this).GetCredentials() == nil {
		errors = append(errors, fmt.Errorf("credentials is nil"))
		return errors
	}

	return (*((*this).GetCredentials())).Validate()
}

func (this *Database) validateDatabaseName() ([]error) {
	var VALID_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return ValidateCharacters(VALID_CHARACTERS, (*this).database_name, "database_name")
}

func (this *Database) GetDatabaseName() *string {
	return (*this).database_name
}

func (this *Database) GetDataDefinitionStatements() []string {
	return (*this).DATA_DEFINITION_STATEMENTS
}

func (this *Database) createDatabase() (*Database, *string, []error) {
	var errors []error 
	crud_sql_command, crud_command_errors := (*this).getCLSCRUDDatabaseCommand((*this).DATA_DEFINITION_STATEMENT_CREATE, (*this).GetExtraOptions())

	if crud_command_errors != nil {
		errors = append(errors, crud_command_errors...)	
	}

	if len(errors) > 0 {
		return nil, nil, errors
	}

	var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd := exec.Command("bash", "-c", *crud_sql_command)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    command_err := cmd.Run()

    if command_err != nil {
		errors = append(errors, command_err)	
	}

	shell_ouput := ""

	if len(errors) > 0 {
		shell_ouput = stderr.String()
		return nil, &shell_ouput, errors
	}

	shell_ouput = stdout.String()
    return this, &shell_ouput, nil
}


func (this *Database) getCLSCRUDDatabaseCommand(command string, options map[string]string) (*string, []error) {
	var errors []error 

	command_errs := Contains((*this).DATA_DEFINITION_STATEMENTS, &command, "command")

	if command_errs != nil {
		errors = append(errors, command_errs...)	
	}

	database_errs := (*this).Validate()

	if database_errs != nil {
		errors = append(errors, database_errs...)	
	}

	logic_option := ""
	if options != nil {
	    logic_option_value, logic_option_exists := options[(*this).LOGIC_OPTION_FIELD]
		if command == (*this).DATA_DEFINITION_STATEMENT_CREATE &&
		   logic_option_exists {
		    logic_option_errors := Contains((*this).LOGIC_OPTION_CREATE_OPTIONS, &logic_option_value, "logic")
			if logic_option_errors != nil {
				errors = append(errors, logic_option_errors...)	
			} else {
				logic_option = logic_option_value
			}
		}
	}

	host_command, host_command_errors := (*(*this).GetHost()).GetCLSCommand()
	if host_command_errors != nil {
		errors = append(errors, host_command_errors...)	
	}

	credentials_command, credentials_command_errors := (*(*this).GetCredentials()).GetCLSCommand()
	if credentials_command_errors != nil {
		errors = append(errors, credentials_command_errors...)	
	}

	if len(errors) > 0 {
		return nil, errors
	}

	sql_command :=  fmt.Sprintf("/usr/local/mysql/bin/mysql %s %s", *host_command, *credentials_command) 
	sql_command += fmt.Sprintf(" -e \"%s DATABASE ", command)
	
	if logic_option != "" {
		sql_command += fmt.Sprintf("%s ", logic_option)
	}
	
	sql_command += fmt.Sprintf("%s ", (*(*this).GetDatabaseName()))
	
	character_set := (*(*this).GetDatabaseCreateOptions()).GetCharacterSet()
	if character_set != nil {
		sql_command += fmt.Sprintf("CHARACTER SET %s ", *character_set)
	}

	collate := (*(*this).GetDatabaseCreateOptions()).GetCollate()
	if collate != nil {
		sql_command += fmt.Sprintf("COLLATE %s", *collate)
	}

	sql_command += ";\""
	return &sql_command, nil
}

func (this *Database) GetHost() *Host {
	return (*this).host
}

func (this *Database) GetCredentials() *Credentials {
	return (*this).credentials
}

func (this *Database) GetDatabaseCreateOptions() *DatabaseCreateOptions {
	return (*this).database_create_options
}

func (this *Database) GetExtraOptions() map[string]string {
	return (*this).extra_options
}