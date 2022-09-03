package gitlab

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/ml444/gitool/git"
	log "github.com/ml444/glog"
	"github.com/xanzy/go-gitlab"
)

func CloneAllRepo(gn int) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	projCh := make(chan *gitlab.Project, 20)
	var isDone bool
	go func() {
		err := IterProjects4AllGroups(&projCh, &isDone)
		if err != nil {
			log.Errorf("err:%v", err)
			return
		}
		log.Info("publisher is close.")
	}()

	var wg sync.WaitGroup
	for i := 0; i < gn; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			defer wg.Done()
			for {
				select {
				case project := <-projCh:
					git.CloneOneProject(project)
				case <-ctx.Done():
					log.Warn("cancel clone repo")
					return
				default:
					if isDone {
						log.Info("publisher is close. so break this goroutine:", i)
						break
					}
					time.Sleep(100 * time.Millisecond)
				}
			}

		}(i)
	}
	log.Info("waiting....")
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
