package data

type Client struct {
	clientID string
	isinClub bool
	table    int
}

func NewClient(clientID string) *Client {
	return &Client{
		clientID: clientID,
		isinClub: true,
		table:    0}
}

func (c *Client) Came() {
	c.isinClub = true
}

func (c *Client) TakeTable(tableID int) {
	c.table = tableID
}

func (c *Client) Gone() {
	c.table = 0
	c.isinClub = false
}

func (c *Client) GetIsInClub() bool {
	return c.isinClub
}

func (c *Client) GetTable() int {
	return c.table
}

func (c *Client) GetClient() string {
	return c.clientID
}
