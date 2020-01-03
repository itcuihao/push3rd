package hwpush

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/itcuihao/gopush/client"
)

func GetToken(clientId, secret string) (string, error) {

	var (
		req     *http.Request
		err     error
		resByte []byte
	)

	param := make(url.Values)
	param.Add("grant_type", GrantType)
	param.Add("client_id", clientId)
	param.Add("client_secret", secret)

	req, err = http.NewRequest("POST", TokenUrl, strings.NewReader(param.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", client.ContentTypeFORM)

	c := client.NewClient(&http.Client{}, 10*time.Second)
	resByte, err = c.Do(req)
	if err != nil {
		return "", err
	}
	fmt.Println(string(resByte))
	res := &TokenRes{}
	err = json.Unmarshal(resByte, res)
	if err != nil {
		return "", err
	}
	return res.Access_token, nil
}
