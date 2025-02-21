package jenkins

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) Verify() (*TokenVerify, error) {
	if c.Auth.Username == "" || c.Auth.Token == "" {
		return nil, fmt.Errorf("define username and token")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/whoAmI/api/json", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// Add Basic Authentication
	auth := c.Auth.Username + ":" + c.Auth.Token
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", basicAuth)

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	ar := TokenVerify{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}
