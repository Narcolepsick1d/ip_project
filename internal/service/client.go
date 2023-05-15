package service

type ClientRepo interface {
	Create(client_user_id int64) error
}
type Client struct {
	repo ClientRepo
}

func (c *Client) Create(client_user_id int64) error {
	return c.repo.Create(client_user_id)
}
