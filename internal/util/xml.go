package util

import (
	"bytes"
	// "fmt"
	"log"
	// "os"
	"text/template"
)

type JobData struct {
	Name    string
	RepoURL string
	// BranchName      string
	JenkinsFilePath string
	CredentialsId   string
}

var (
// repoURL         string
// branchName      string
// jenkinsFilePath string
// credentialsId   string
)

func ReadXml(name, repoURL, jenkinsfilePath, credentialsId, jobType string) ([]byte, error) {
	var tmpl *template.Template
	var err error

	switch jobType {
	case "multi-branch":
		{
			tmpl, err = template.ParseFiles("./internal/templates/multi-branch/config.xml")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	data := &JobData{
		Name:    name,
		RepoURL: repoURL,
		// BranchName:      branchName,
		JenkinsFilePath: jenkinsfilePath,
		CredentialsId:   credentialsId,
	}
	var xmlBuffer bytes.Buffer
	if err := tmpl.Execute(&xmlBuffer, data); err != nil {
		log.Fatal(err)
	}

	// if err := os.WriteFile("./internal/templates/generated.xml", xmlBuffer.Bytes(), 0644); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("File saved to path")

	return xmlBuffer.Bytes(), nil

}
