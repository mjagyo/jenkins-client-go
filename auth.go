package jenkins

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Sign up - Create new user, return user token upon successful creation
func (c *Client) SignUp(auth AuthStruct) (*AuthResponse, error) {
	if auth.Username == "" || auth.Token == "" {
		return nil, fmt.Errorf("define username and password")
	}
	rb, err := json.Marshal(auth)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signup", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

// Verify - Check the token that will be used
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

	fmt.Printf("----- %s body", body)

	ar := TokenVerify{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

// SignIn - Get a new token for user
func (c *Client) GetUserTokenSignIn(auth AuthStruct) (*AuthResponse, error) {
	if auth.Username == "" || auth.Token == "" {
		return nil, fmt.Errorf("define username and password")
	}
	rb, err := json.Marshal(auth)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signin", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, errors.New("Unable to login")
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

// SignOut - Revoke the token for a user
func (c *Client) SignOut(authToken *string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signout", c.HostURL), strings.NewReader(string("")))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return err
	}

	if string(body) != "Signed out user" {
		return errors.New(string(body))
	}

	return nil
}
