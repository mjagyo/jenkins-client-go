package jenkins

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) CreateSecret(secretDetails CredentialRequest) error {
	rb, err := json.Marshal(secretDetails)
	if err != nil {
		return err
	}

	formEncoded := "json=" + string(rb)
	fmt.Printf("Raw Response Body: %v \n", formEncoded)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/credentials/store/system/domain/_/createCredentials", c.HostURL), strings.NewReader(formEncoded))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetSecret(secretId string) (*CredentialResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/credentials/store/system/domain/_/credential/%s/api/json", c.HostURL, secretId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	credential := CredentialResponse{}
	err = json.Unmarshal(body, &credential)
	if err != nil {
		return nil, err
	}

	return &credential, nil
}

// CreateOrder - Create new order
func (c *Client) UpdateSecret(secretDetails Credential) error {
	rb, err := json.Marshal(secretDetails)

	if err != nil {
		return err
	}

	formEncoded := "json=" + string(rb)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/manage/credentials/store/system/domain/_/credential/%s/updateSubmit", c.HostURL, secretDetails.ID), strings.NewReader(formEncoded))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteSecret(secretId string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/manage/credentials/store/system/domain/_/credential/%s/doDelete", c.HostURL, secretId), nil)

	if err != nil {
		return err
	}

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}
