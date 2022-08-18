package gitlab

import (
	"github.com/ml444/gitool/conf"
	"github.com/ml444/glog"
	"github.com/xanzy/go-gitlab"
)

var Cli *gitlab.Client

func InitClient() error {
	cli, err := gitlab.NewClient(conf.Get(conf.GitlabAccessToken), gitlab.WithBaseURL(conf.Get(conf.GitlabBaseUrl)))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	Cli = cli
	//users, _, err := cli.Users.ListUsers(&gitlab.ListUsersOptions{})
	//if err != nil {
	//	log.Errorf("err:%v", err)
	//	return err
	//}
	//log.Info(users)
	return nil
}
