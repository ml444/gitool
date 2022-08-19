package main

import (
	"flag"
	"fmt"
	"github.com/ml444/gitool/cmd"
	"github.com/ml444/gitool/conf"
	"github.com/ml444/gitool/gitlab"
	log "github.com/ml444/glog"
	"github.com/ml444/glog/config"
	"os"
)

func main() {
	var err error
	var cmdStr string
	gn := flag.Int("gn", 1, "concurrent number")
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

	err = log.InitLog(
		config.SetFileDir(".", false),
		config.SetFileName("gitool", false),
	)
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}
	_ = conf.InitEnv()
	err = gitlab.InitClient()
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}
	switch cmdStr {
	case "clone":
		fmt.Println("running clone all repo")
		fmt.Println("===>", *gn)
		cmd.CloneAllRepo(*gn)
	default:
		flag.Usage()
	}

	//_, err = gitlab.ListProjects4AllGroups()
	//if err != nil {
	//	log.Errorf("err:%v", err)
	//	return
	//}

}
