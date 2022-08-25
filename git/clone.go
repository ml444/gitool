package git

import (
	"context"
	"fmt"
	"github.com/ml444/gitool/conf"
	log "github.com/ml444/glog"
	"github.com/xanzy/go-gitlab"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"os"
)

func CloneOneProject(project *gitlab.Project) {
	var err error
	_, err = git.PlainClone(getRepoLocalPath(project.PathWithNamespace), false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: conf.Get(conf.GitlabUsername),
			Password: conf.Get(conf.GitlabAccessToken),
		},
		URL:      project.SSHURLToRepo,
		Progress: os.Stdout,
	})

	if err != nil {
		// Error Cloning xxx_server: repository already exists
		fmt.Println("Error Cloning " + project.Name + ": " + err.Error())
	} else {
		fmt.Println("Successfully cloned repo:", project.Name)
	}
}
func CloneProjects(ctx context.Context, prjCh <-chan *gitlab.Project) {
	for {
		select {
		case project := <-prjCh:
			var err error
			_, err = git.PlainClone(getRepoLocalPath(project.PathWithNamespace), false, &git.CloneOptions{
				Auth: &http.BasicAuth{
					Username: conf.Get(conf.GitlabUsername),
					Password: conf.Get(conf.GitlabAccessToken),
				},
				URL:      project.SSHURLToRepo,
				Progress: os.Stdout,
			})

			if err != nil {
				// Error Cloning xxx_server: repository already exists
				fmt.Println("Error Cloning " + project.Name + ": " + err.Error())
			}
		case <-ctx.Done():
			log.Warnf("cancel clone repo")
			return
		}
	}
}

func getRepoLocalPath(pathWithNamespace string) string {
	return conf.Get(conf.GitlabLocalBaseDir) + "/" + pathWithNamespace
}
