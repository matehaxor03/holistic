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

func (this *HostSession) CreateDatabase(database *Database) []error {
	var errors []error 

	host_session_errs := *this.Validate()

	if host_session_errs != nil {
		errors = append(errors, host_session_errs...)	
	}

	database_errs :=  (*(*this).database).Validate()

	if database_errs != nil {
		errors = append(errors, database_errs...)	
	}

	if len(errors) > 0 {
		return errors
	}

	// execute mysql command

	return nil
}




