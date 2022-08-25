package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/ml444/gitool/git"
	gl "github.com/ml444/gitool/gitlab"
	log "github.com/ml444/glog"
	"github.com/xanzy/go-gitlab"
	"os"
	"strconv"
)

func CloneOneRepo(groupName string, repoName string) {

	//git.CloneOneProject(project)
}

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

func SearchRepo(search string) {
	searchList, err := gl.SearchProjectsByName(search)
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

func DefaultSelectOption() {
	prompt := promptui.Select{
		Label: "Select repo",
		Items: []string{"cloneAllRepo", "searchAndSelectRepo", "exit"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	switch result {
	case "cloneAllRepo":
		var goroutineCount int64
		prompt := promptui.Prompt{
			Label: "How much concurrency do you want <int>",
			Validate: func(s string) error {
				goroutineCount, err = strconv.ParseInt(s, 10, 64)
				if err != nil {
					log.Errorf("err:%v", err)
					return err
				}
				return nil
			},
		}
		_, err := prompt.Run()
		if err != nil {
			log.Errorf("err:%v", err)
			return
		}
		CloneAllRepo(int(goroutineCount))
	case "searchAndSelectRepo":
		prompt := promptui.Prompt{
			Label: "Input repo name",
			Validate: func(s string) error {
				if s == "" {
					return errors.New("must not null")
				}
				return nil
			},
		}
		v, err := prompt.Run()
		if err != nil {
			log.Errorf("err:%v", err)
			return
		}
		SearchRepo(v)
	default:
		os.Exit(0)
	}
}
