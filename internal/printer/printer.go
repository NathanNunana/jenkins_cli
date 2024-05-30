package printer

import (
	"fmt"
	"strings"

	"github.com/NathanNunana/jenkins_cli/internal/jenkins"
)

type IPrinter interface {
	List() (string, error)
}

type PrintJob struct {
	Jobs []jenkins.Job
}

func (*PrintJob) getStatus(color string) string {
	if strings.ToLower(color) == "green" {
		return "Successful"
	}
	return "Failed"
}

func (p *PrintJob) List() (string, error) {
	var builder strings.Builder
	builder.WriteString("NAME\tURL\tSTATUS\n")
	for _, job := range p.Jobs {
		builder.WriteString(fmt.Sprintf("%s\t%s\t%s\n", job.Name, job.Url, p.getStatus(job.Color)))
	}
	return builder.String(), nil
}
