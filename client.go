package jenkins

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"
)

const HostURL string = "http://localhost:8080"

// Client -
type Client struct {
	HostURL       string
	HTTPClient    *http.Client
	Token         string
	Auth          AuthStruct
	Base64Token   string
	Authenticated TokenVerify
}

type TokenVerify struct {
	Authenticated bool `json:"authenticated"`
}

// AuthStruct -
type AuthStruct struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

// AuthResponse -
type AuthResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

// NewClient -
func NewClient(host, username, token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// If username or token not provided, return empty client
	if username == nil || token == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Username: *username,
		Token:    *token,
	}

	ar, err := c.Verify()
	if err != nil {
		return nil, err
	}

	credentials := c.Auth.Username + ":" + c.Auth.Token

	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(credentials))

	c.Authenticated = TokenVerify{
		Authenticated: ar.Authenticated,
	}

	c.Base64Token = encodedCredentials

	return &c, nil
}
func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Base64Token

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("Authorization", "Basic "+token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
