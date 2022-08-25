package gitlab

import (
	"fmt"
	"os"

	"github.com/ml444/glog"
	"github.com/xanzy/go-gitlab"
)

func ListAllProjects() {
	opt := &gitlab.ListProjectsOptions{
		Owned: gitlab.Bool(false),
		//Membership: gitlab.Bool(true),
	}
	projects, rsp, err := Cli.Projects.ListProjects(opt)

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
	}
	//l := list.NewWriter()
	//for _, project := range projects {
	//	l.AppendItem(project.Name)
	//}
	//l.SetStyle(list.StyleBulletCircle)
	//fmt.Println("\n")
	//consoleLog("List all your projects", l.Render(), "")
}

func GetProject(pid int) (*gitlab.Project, error) {
	opt := gitlab.GetProjectOptions{}
	project, _, err := Cli.Projects.GetProject(pid, &opt)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return project, nil
}

func SearchProjectsByName(name string) ([]*gitlab.Project, error) {
	opt := gitlab.ListProjectsOptions{
		ListOptions:              gitlab.ListOptions{},
		Archived:                 nil,
		IDAfter:                  nil,
		IDBefore:                 nil,
		LastActivityAfter:        nil,
		LastActivityBefore:       nil,
		Membership:               nil,
		MinAccessLevel:           nil,
		OrderBy:                  nil,
		Owned:                    nil,
		RepositoryChecksumFailed: nil,
		RepositoryStorage:        nil,
		Search:                   &name,
		SearchNamespaces:         nil,
		Simple:                   nil,
		Sort:                     nil,
		Starred:                  nil,
		Statistics:               nil,
		Topic:                    nil,
		Visibility:               nil,
	}
	projects, _, err := Cli.Projects.ListProjects(&opt)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}
	//fmt.Println(rsp)
	return projects, nil
}
