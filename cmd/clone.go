package cmd

import (
	"context"
	"github.com/ml444/gitool/git"
	gl "github.com/ml444/gitool/gitlab"
	log "github.com/ml444/glog"
	"github.com/xanzy/go-gitlab"
)

func CloneAllRepo(gn int) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	projCh := make(chan *gitlab.Project, 20)
	err := gl.IterProjects4AllGroups(projCh)
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}
	for i := 0; i < gn; i++ {
		git.CloneProjects(ctx, projCh)
	}
}
