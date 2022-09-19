package class

type HostSession struct {
	host *Host
	creds *Credentials
}

func NewHostSession(host *Host, creds *Credentials) (*HostSession) {
	x := HostSession{host: host, creds: creds}
	return &x
}

func (this *HostSession) Validate() []error {
	var errors []error 

	host_errs := (*(*this).host).Validate()

	if host_errs != nil {
		errors = append(errors, host_errs...)	
	}


	creds_errs :=  (*(*this).creds).Validate()

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

	//CREATE DATABASE IF NOT EXISTS " + db_name + " CHARACTER SET utf8 COLLATE utf8_general_ci
	command := fmt.Printf("CREATE DATABASE IF NOT EXISTS %s", (*database).GetDatabaseName())

	character_set := (*database).GetCharacterSet()
	if character_set != nil {
		command += fmt.Printf(" CHARACTER SET %s", character_set)
	}

	collate := (*database).GetCollate()
	if collate != nil {
		command += fmt.Printf(" COLLATE %s", collate)
	}

	command += ";"

	// execute mysql command

	return nil
}




