package git

import (
	"context"
	"os"
	"os/exec"
	"sync"

	log "github.com/ml444/glog"
	"github.com/xanzy/go-gitlab"
)

func PullOneRepo(dir string) error {
	var err error
	err = os.Chdir(dir)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("==> pull repo to: ", dir)
	out, err := exec.Command("git", "pull").Output()
	if err != nil {
		log.Error(err)
		return err
	}
	if len(out) != 0 {
		log.Info(string(out))
	}
	return nil
}

func PullProjects(ctx context.Context, wg *sync.WaitGroup, projCh *chan *gitlab.Project) {
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case project := <-*projCh:
			dir, err := GetRepoLocalPath(project.HTTPURLToRepo)
			if err != nil {
				log.Error(err)
				return
			}
			PullOneRepo(dir)
		case <-ctx.Done():
			log.Warn("cancel pull")
			return
		}
	}
}
