package class

type Client struct {
    host *Host
	credentials *Credentials
	database *Database
}

func NewClient(host *Host, credentials *Credentials, database *Database) (*Client) {
	x := Client{host: host, credentials: credentials, database: database}
	return &x
}

