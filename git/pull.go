package git

import (
	log "github.com/ml444/glog"
	"os"
	"os/exec"
)

func PullOneRepo(dir string) error {
	var err error
	err = os.Chdir(dir)
	if err != nil {
		return err
	}
	out, err := exec.Command("git", "pull").Output()
	if err != nil {
		return err
	}
	if len(out) != 0 {
		log.Info(string(out))
	}
	return nil
}
