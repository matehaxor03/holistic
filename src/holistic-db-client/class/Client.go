package class

type Client struct {
	hostSesssion *HostSession
}

func NewClient() (*Client) {
	x := Client{}
	return &x
}

func (this *Client) Login(host *Host, creds *Credentials) (*HostSession, []error) {
	var errors []error 
	var x *HostSession 
	x = NewHostSession(host, creds)

	host_session_errs := (*x).Validate()

	if host_session_errs != nil {
		errors = append(errors, host_session_errs...)	
	}

	if errors != nil {
		return nil, errors
	}
	
	(*this).hostSesssion = x
	return x, nil
}

