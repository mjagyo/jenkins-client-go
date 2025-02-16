package jenkins

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetJobs - Returns list of jobs
func (c *Client) GetJobs(authToken *string) (*Jobs, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/json?tree=jobs[name,url,color,description,displayName,fullName,buildable,inQueue]", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	var jobs *Jobs
	err = json.Unmarshal(body, &jobs)
	if err != nil {
		return nil, err
	}

	return jobs, nil
}
