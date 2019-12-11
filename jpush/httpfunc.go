package jpush

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	contentType string
	httpClient  *http.Client
}

func NewClient(contentType string, httpClient *http.Client) Client {
	return Client{
		contentType: contentType,
		httpClient:  httpClient,
	}
}

func (c Client) GetContentType() string {
	return c.contentType
}

func (c Client) GetHttpClient() *http.Client {
	return c.httpClient
}

func (c Client) DoPost(req *http.Request) ([]byte, error) {
	var (
		result []byte
		res    *http.Response
		err    error
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	client := c.GetHttpClient()
	if client == nil {
		client = &http.Client{}
	}

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
	result, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
