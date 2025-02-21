package jenkins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func (c *Client) GetJobs(authToken *string) (*Jobs, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/json?tree=jobs[name,url,color,description,displayName,fullName,buildable,inQueue]", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	fmt.Print("11")
	body, err := c.doRequest(req, authToken)
	if err != nil {
		return nil, err
	}

	jobs := Jobs{}
	err = json.Unmarshal(body, &jobs)
	if err != nil {
		return nil, err
	}

	return &jobs, nil
}

func (c *Client) CreateJob(jobName, filePath string) error {
	fmt.Printf("Creating job %s\n", jobName)
	configData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// print configData
	fmt.Printf("configData: %s\n", configData)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/createItem?name=%s", c.HostURL, jobName), bytes.NewBuffer(configData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/xml")

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateJob(jobName, filePath string) error {
	configData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/job/%s/config.xml", c.HostURL, jobName), bytes.NewBuffer(configData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/xml")

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteJob(jobName string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/job/%s/doDelete", c.HostURL, jobName), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}
