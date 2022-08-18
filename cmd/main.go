package main

import (
	"github.com/ml444/gitool/conf"
	"github.com/ml444/gitool/gitlab"
	log "github.com/ml444/glog"
	"github.com/ml444/glog/config"
)

func main() {
	var err error
	err = log.InitLog(
		config.SetFileDir(".", false),
		config.SetFileName("gitlab", false),
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
	//_, err = gitlab.ListProjects4AllGroups()
	//if err != nil {
	//	log.Errorf("err:%v", err)
	//	return
	//}
	//gitlab.ListAllProjects()
}
