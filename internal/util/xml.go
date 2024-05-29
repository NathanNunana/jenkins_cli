package util

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"
)

type JobData struct {
	RepoURL         string
	BranchName      string
	JenkinsFilePath string
}

var (
	repoURL         string
	branchName      string
	jenkinsFilePath string
)

func ReadXml() {
	tmpl, err := template.ParseFiles("../template/config.xml")
	if err != nil {
		log.Fatal(err)
	}
	data := &JobData{
		RepoURL:         repoURL,
		BranchName:      branchName,
		JenkinsFilePath: jenkinsFilePath,
	}
	var xmlBuffer bytes.Buffer
	if err := tmpl.Execute(&xmlBuffer, data); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("../template/generated.xml", xmlBuffer.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("File saved to path")
}
