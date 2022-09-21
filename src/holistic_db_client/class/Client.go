package class

type Client struct {
	host *Host
	credentials *Credentials
	database *Database
}

func NewClient(host *Host, credentials *Credentials, database *Database) (*Client) {
	x := Client{host: host,
				credentials: credentials,
				database: database}
			    
	return &x
}

func (this *Client) CreateDatabase(database_name *string, database_create_options *DatabaseCreateOptions, options map[string][]string) (*Database, *string, []error) {
	return NewDatabase((*this).GetHost(), (*this).GetCredentials(), database_name, database_create_options, options).Create()
}

func (this *Client) GetHost() *Host {
	return (*this).host
 }

func (this *Client) GetCredentials() *Credentials {
	return (*this).credentials
 }
