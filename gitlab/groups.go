package gitlab

import (
	"fmt"
	"github.com/ml444/gitool/git"
	log "github.com/ml444/glog"
	"github.com/xanzy/go-gitlab"
	"os"
)

func GetAllGroups() ([]*gitlab.Group, error) {
	var groupList []*gitlab.Group
	var nextPage int
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

func ListProjects4AllGroups() (interface{}, error) {
	allGroups, err := GetAllGroups()
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}
	for _, group := range allGroups {
		err = ListProjectsByGroup(group.ID)
		if err != nil {
			log.Errorf("err:%v", err)
			return nil, err
		}
	}
	return nil, nil
}

func ListProjectsByGroup(groupId int) error {
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
			fmt.Println(project.SSHURLToRepo)
			git.CloneProject(project)
		}
		if rsp.CurrentPage == rsp.TotalPages {
			break
		}
		nextPage = rsp.CurrentPage + 1
	}

	//l := list.NewWriter()
	//for _, project := range projects {
	//	l.AppendItem(project.Name)
	//}
	//l.SetStyle(list.StyleBulletCircle)
	//fmt.Println("\n")
	//consoleLog("List all your projects", l.Render(), "")
	return nil
}
