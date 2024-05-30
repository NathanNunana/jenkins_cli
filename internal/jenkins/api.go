package jenkins

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NathanNunana/jcli/internal/config"
	"github.com/NathanNunana/jcli/internal/util"
)

func GetJobs() ([]Job, error) {
	path, err := util.GetPath()
	if err != nil {
		return nil, err
	}

	cfg, err := config.GetConfig(path)
	if err != nil {
		return nil, err
	}

	// api
	url := fmt.Sprintf("%s/api/json", cfg.JenkinsURL)
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// authenticate
	req.SetBasicAuth(cfg.Username, cfg.ApiToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to load jobs, %v", resp.StatusCode)
	}

	var jenkinsResponse JobResponse

	if err := json.NewDecoder(resp.Body).Decode(&jenkinsResponse); err != nil {
		return nil, err
	}
	return jenkinsResponse.Jobs, nil
}

func CreateJob(name string) (Job, error) {
	path, err := util.GetPath()
	if err != nil {
		return Job{}, err
	}
	cfg, err := config.GetConfig(path)
	if err != nil {
		return Job{}, err
	}
	util.ReadXml()
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/createItem?name=%s", cfg.JenkinsURL, name), nil)
	if err != nil {
		return Job{}, nil
	}
	req.SetBasicAuth(cfg.Username, cfg.ApiToken)
	req.Header.Set("Content-Type", "application/xml")
	return Job{}, nil
}
