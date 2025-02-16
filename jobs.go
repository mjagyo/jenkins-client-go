package jenkins

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetJobs - Returns list of jobs
func (c *Client) GetJobs() (Jobs, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/json?tree=jobs[name,url,color,description,displayName,fullName,buildable,inQueue]", c.HostURL), nil)
	if err != nil {
		return Jobs{}, err
	}

	// Add Basic Authentication
	auth := c.Auth.Username + ":" + c.Auth.Token
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", basicAuth)

	body, err := c.doRequest(req, nil)
	if err != nil {
		return Jobs{}, err
	}

	var jobs Jobs
	err = json.Unmarshal(body, &jobs)
	if err != nil {
		return Jobs{}, err
	}

	return jobs, nil
}
