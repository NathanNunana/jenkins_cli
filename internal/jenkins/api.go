package jenkins

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NathanNunana/jenkins_cli/internal/config"
	"github.com/NathanNunana/jenkins_cli/internal/util"
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

func CreateJob() (Job, error) {
	return Job{}, nil
}
