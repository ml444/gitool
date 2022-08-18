package gitlab

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
	"os"
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
