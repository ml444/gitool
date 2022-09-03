package gitlab

import (
	"fmt"
	"os"

	log "github.com/ml444/glog"
	"github.com/xanzy/go-gitlab"
)

func GetAllGroups() ([]*gitlab.Group, error) {
	var groupList []*gitlab.Group
	var nextPage int
	log.Info("===> Get all groups")
	for {
		ops := &gitlab.ListGroupsOptions{
			ListOptions: gitlab.ListOptions{
				Page:    nextPage,
				PerPage: 0,
			},
		}
		groups, rsp, err := Cli.Groups.ListGroups(ops)
		if err != nil {
			log.Errorf("err:%v", err)
			return nil, err
		}
		fmt.Printf("%#v \n", rsp)
		groupList = append(groupList, groups...)
		if rsp.CurrentPage == rsp.TotalPages {
			break
		}
		nextPage = rsp.CurrentPage + 1
	}

	return groupList, nil
}

func IterProjects4AllGroups(projCh chan *gitlab.Project) error {
	allGroups, err := GetAllGroups()
	log.Info("===> Iter projects from all groups")
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}
	for _, group := range allGroups {
		err = ListProjectsByGroup(group.ID, projCh)
		if err != nil {
			log.Errorf("err:%v", err)
			return err
		}
	}
	return nil
}

func ListProjectsByGroup(groupId int, projCh chan *gitlab.Project) error {
	var nextPage int
	for {
		opt := &gitlab.ListGroupProjectsOptions{
			ListOptions: gitlab.ListOptions{
				Page:    nextPage,
				PerPage: 0,
			},
			//Owned: gitlab.Bool(false),
			//Membership: gitlab.Bool(true),
		}
		projects, rsp, err := Cli.Groups.ListGroupProjects(groupId, opt)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%#v \n", rsp)
		if len(projects) == 0 {
			fmt.Println("No Projects Found")
			os.Exit(0)
		}
		for _, project := range projects {
			projCh <- project
			fmt.Println(project.SSHURLToRepo)
		}
		if rsp.CurrentPage == rsp.TotalPages {
			break
		}
		nextPage = rsp.CurrentPage + 1
	}
	return nil
}
