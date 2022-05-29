package news

import (
	"context"
	"fmt"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/config"
	"net/http"
	"net/url"
)

const getNewsPath = "/news"

type Client struct {
	ctx        context.Context
	httpclient *http.Client
}

func New(ctx context.Context) *Client {
	return &Client{
		ctx:        ctx,
		httpclient: http.DefaultClient,
	}
}

// WithTransport Для прокси сервера
func (c *Client) WithTransport(transport *http.Transport) {
	c.httpclient.Transport = transport
}

func (c *Client) GetNews(req RequestGetNews) (*RequestGetNews, error) {
	cfg := config.FromContext(c.ctx).BMSTUNewsConfig

	url := url.URL{
		Scheme: cfg.Protocol,
		Host:   cfg.SiteAddress,
		Path:   getNewsPath, //Добавлять ли сюда ?&limit=200&offset=0 в качестве дополнительного параметра?
	}

	fmt.Println(url.String())

	/*req := http.NewRequest(http.MethodGet)
	c.httpclient.Do()*/

	return nil, nil
}
