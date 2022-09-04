package gitlab

import (
	"context"
	"fmt"
	"sync"

	"github.com/manifoldco/promptui"
	"github.com/ml444/gitool/git"
	log "github.com/ml444/glog"
	"github.com/xanzy/go-gitlab"
)

func CloneAllRepo(gn int) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	projCh := make(chan *gitlab.Project, 20)
	err := IterProjects4AllGroups(projCh)
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}
	var wg sync.WaitGroup
	for i := 0; i < gn; i++ {
		go git.CloneProjects(ctx, &wg, projCh)
	}
	wg.Wait()
}

func CloneRepoBySearch(search string) {
	searchList, err := SearchProjectsByName(search)
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}
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
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	project := projectMap[result]
	git.CloneOneProject(project)
}
