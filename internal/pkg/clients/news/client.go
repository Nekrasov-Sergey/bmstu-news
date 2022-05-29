package news

import "net/http"

type Client struct {
	client http.Client
}

func (c *Client) GetNews() {

}
