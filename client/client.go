package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Client client
type Client struct {
	contentType string
	httpClient  *http.Client
	withoutTime time.Duration
}

const (
	ContentTypeFORM = "application/x-www-form-urlencoded"
	ContentTypeJSON = "application/json"
)

// NewClient new client
func NewClient(httpClient *http.Client, withoutTime time.Duration) Client {
	return Client{
		httpClient:  httpClient,
		withoutTime: withoutTime,
	}
}

// GetHttpClient get http_client
func (c Client) GetHttpClient() *http.Client {
	if c.httpClient == nil {
		return &http.Client{}
	}
	return c.httpClient
}

// GetWithoutTime get without time
func (c Client) GetWithoutTime() time.Duration {
	return c.withoutTime
}

// Do do transport
func (c Client) Do(req *http.Request) ([]byte, error) {
	var (
		res *http.Response
		err error
	)

	ctx, cancel := context.WithTimeout(context.Background(), c.GetWithoutTime())
	defer cancel()

	req = req.WithContext(ctx)

	client := c.GetHttpClient()

	// 出错重试机制
	tryCount := 0
tryAgain:
	res, err = client.Do(req)
	if err != nil {
		fmt.Println("do post err:", err, tryCount)
		select {
		case <-ctx.Done():
			return nil, err
		default:
		}
		tryCount++
		if tryCount < 3 {
			goto tryAgain
		}
		return nil, err
	}

	if res.Body == nil {
		return nil, fmt.Errorf("do post response is nil")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code error: %v", res.StatusCode)
	}

	return ioutil.ReadAll(res.Body)
}
