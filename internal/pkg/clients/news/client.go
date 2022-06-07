package news

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/config"
)

const (
	getNewsPath = "/news"

	urlLimitKey  string = "limit="
	urlOffsetKey string = "offset="
)

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

func (c *Client) GetNews(limit int, offset int) (*ResponseNews, error) {
	cfg := config.FromContext(c.ctx).BMSTUNewsConfig

	url := url.URL{
		Scheme:   cfg.Protocol,
		Host:     cfg.SiteAddress,
		Path:     getNewsPath,
		RawQuery: urlLimitKey + strconv.Itoa(limit) + urlOffsetKey + strconv.Itoa(offset), //Добавлять ли сюда ?&limit=200&offset=0 в качестве дополнительного параметра?
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
	var resp_test ResponseNews

	err = json.Unmarshal(bts, &resp)
	if err != nil {
		return nil, err
	}

	if reflect.DeepEqual(resp.Items, resp_test.Items) {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetFullNews(slug string) (*ResponseFullNews, error) {
	cfg := config.FromContext(c.ctx).BMSTUNewsConfig

	url := url.URL{
		Scheme: cfg.Protocol,
		Host:   cfg.SiteAddress,
		Path:   getNewsPath + "/" + slug,
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

	var resp ResponseFullNews

	err = json.Unmarshal(bts, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
