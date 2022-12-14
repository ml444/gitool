package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/ml444/gitool/conf"
	"github.com/ml444/gitool/gitlab"
	log "github.com/ml444/glog"
)

func CloneOneRepoBySearch(domain int, searchContent string) {
	switch domain {
	case conf.GitDomainGithub:
		// Pass
	case conf.GitDomainGitlab:
		gitlab.CloneRepoBySearch(searchContent)
	}
}

func CloneAllRepo(domain int, gn int) {
	switch domain {
	case conf.GitDomainGithub:
		// Pass
	case conf.GitDomainGitlab:
		gitlab.CloneAllRepo(gn)
	}
}

func CloneDefaultSelectOption(gitDomain int) {
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
				if s == "" {
					return nil
				}
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
		switch gitDomain {
		case conf.GitDomainGithub:
			// pass
		case conf.GitDomainGitlab:
			gitlab.CloneAllRepo(int(goroutineCount))
		}
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
		switch gitDomain {
		case conf.GitDomainGithub:
			// pass
		case conf.GitDomainGitlab:
			gitlab.CloneRepoBySearch(v)
		}
	default:
		os.Exit(0)
	}
}
