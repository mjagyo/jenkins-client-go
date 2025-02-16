package jenkins

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// HostURL - Default Hashicups URL
const HostURL string = "http://localhost:8080"

// Client -
type Client struct {
	HostURL       string
	HTTPClient    *http.Client
	Token         string
	Auth          AuthStruct
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
		// Default Hashicups URL
		HostURL: HostURL,
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

	c.Authenticated = TokenVerify{
		Authenticated: ar.Authenticated,
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	fmt.Printf("--- %v c.doRequest =-======  %v", c, authToken)
	token := c.Token

	if authToken != nil {
		token = *authToken
	}

	fmt.Printf("should run here")

	req.Header.Set("Authorization", token)

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
