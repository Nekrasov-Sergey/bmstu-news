package news

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/config"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
)

const getNewsPath = "/news"

type Client struct {
	ctx        context.Context
	httpClient *http.Client
}

func New(ctx context.Context) *Client {
	return &Client{
		ctx:        ctx,
		httpClient: http.DefaultClient,
	}
}

// WithTransport Для прокси сервера
func (c *Client) WithTransport(transport *http.Transport) {
	c.httpClient.Transport = transport
}

func (c *Client) GetNews() (*ResponseNews, error) {
	cfg := config.FromContext(c.ctx).BMSTUNewsConfig

	url := url.URL{
		Scheme: cfg.Protocol,
		Host:   cfg.SiteAddress,
		Path:   getNewsPath, //Добавлять ли сюда ?&limit=200&offset=0 в качестве дополнительного параметра?
	}

	log.Info("generated url ", url.String())

	reqToBMSTU, err := http.NewRequest(http.MethodGet, url.String(), bytes.NewBuffer([]byte("")))
	if err != nil {
		return nil, err
	}

	r, err := c.httpClient.Do(reqToBMSTU)
	if err != nil {
		return nil, err
	}

	bts, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var resp ResponseNews

	err = json.Unmarshal(bts, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
