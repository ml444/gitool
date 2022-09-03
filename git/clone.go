package git

import (
	"context"
	"os"
	"sync"

	"github.com/ml444/gitool/conf"
	log "github.com/ml444/glog"
	"github.com/xanzy/go-gitlab"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func CloneOneProject(project *gitlab.Project) {
	var err error
	repoPath, err := GetRepoLocalPath(project.HTTPURLToRepo)
	if err != nil {
		log.Error(err)
		return
	}
	_, err = git.PlainClone(repoPath, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: conf.Get(conf.GitlabUsername),
			Password: conf.Get(conf.GitlabAccessToken),
		},
		URL:      project.SSHURLToRepo,
		Progress: os.Stdout,
	})

	if err != nil {
		// Error Cloning xxx_server: repository already exists
		log.Errorf("Error Cloning " + project.Name + ": " + err.Error())
	} else {
		log.Info("Successfully cloned repo:", project.Name)
	}
}
func CloneProjects(ctx context.Context, wg *sync.WaitGroup) {
}
