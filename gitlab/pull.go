package gitlab

import (
	"context"
	"sync"

	"github.com/manifoldco/promptui"
	"github.com/ml444/gitool/git"
	log "github.com/ml444/glog"
	"github.com/xanzy/go-gitlab"
)

func PullAllRepo(gn int) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	projCh := make(chan *gitlab.Project, 20)
	err := IterProjects4AllGroups(projCh)
	if err != nil {
		log.Error(err)
		return
	}
	var wg sync.WaitGroup
	for i := 0; i < gn; i++ {
		go git.PullProjects(ctx, &wg, projCh)
	}
	wg.Wait()
}

func PullOneRepoBySearch(repoName string) {
	searchList, err := SearchProjectsByName(repoName)
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}
	var project *gitlab.Project
	if len(searchList) == 1 {
		project = searchList[0]
		goto PULL
	} else {
		var selectList []string
		projectMap := map[string]*gitlab.Project{}
		for _, project := range searchList {
			key := project.PathWithNamespace
			selectList = append(selectList, key)
			projectMap[key] = project
		}
		prompt := promptui.Select{
			Label: "Select repo",
			Items: selectList,
		}

		_, result, err := prompt.Run()
		if err != nil {
			log.Errorf("Prompt failed %v\n", err)
			return
		}
		project = projectMap[result]
	}
PULL:
	//dir, err := git.GetRepoPathByHTTPSURL(project.HTTPURLToRepo)
	dir, err := git.GetRepoLocalPath(project.HTTPURLToRepo)
	if err != nil {
		log.Error(err)
		return
	}
	err = git.PullOneRepo(dir)
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("pull %s completed", dir)
}
