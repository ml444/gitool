package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ml444/gitool/cmd"
	"github.com/ml444/gitool/conf"
	"github.com/ml444/gitool/gitlab"
	log "github.com/ml444/glog"
)

func main() {
	var err error
	var cmdStr string
	gn := flag.Int("C", 1, "goroutine concurrent count")
	search := flag.String("s", "", "search repo name")
	all := flag.Bool("all", false, "operate all repo")
	if argsLen := len(os.Args); argsLen == 1 {
		flag.Usage()
		return
	} else {
		cmdStr = os.Args[1]
		if argsLen > 2 {
			err = flag.CommandLine.Parse(os.Args[2:])
			if err != nil {
				log.Errorf("err:%v", err)
				return
			}
		}
	}

	_ = conf.InitEnv()
	err = gitlab.InitClient()
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}
	switch cmdStr {
	case "clone":
		if *all {
			fmt.Println("running clone all repo")
			fmt.Println("===>", *gn)
			cmd.CloneAllRepo(*gn)
		} else if *search != "" {
			cmd.SearchRepo(*search)
		} else {
			cmd.DefaultSelectOption()
		}
	default:
		flag.Usage()
	}

	//_, err = gitlab.ListProjects4AllGroups()
	//if err != nil {
	//	log.Errorf("err:%v", err)
	//	return
	//}

}
