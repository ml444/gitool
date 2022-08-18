package git

import (
	"fmt"
	"github.com/ml444/gitool/conf"
	"github.com/xanzy/go-gitlab"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"os"
)

func CloneProject(project *gitlab.Project) {
	var err error
	_, err = git.PlainClone(conf.Get(conf.GitlabLocalBaseDir)+"/"+project.PathWithNamespace, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: conf.GitlabUsername,
			Password: conf.Get(conf.GitlabAccessToken),
		},
		URL:      project.SSHURLToRepo,
		Progress: os.Stdout,
	})

	if err != nil {
		// Error Cloning xxx_server: repository already exists
		fmt.Println("Error Cloning " + project.Name + ": " + err.Error())
	}
}
