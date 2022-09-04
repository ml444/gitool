package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ml444/gitool/cmd"
	"github.com/ml444/gitool/conf"
	"github.com/ml444/gitool/gitlab"
	log "github.com/ml444/glog"
)

func main() {
	var err error
	var cmdStr string
	var domain int
	gn := flag.Int("C", 1, "goroutine concurrent count")
	tPtr := flag.Int("T", 0, "<DomainType: 1-github,2-gitlab>")
	all := flag.Bool("all", false, "operate all repo")
	search := flag.String("s", "", "search repo name")
	domain = *tPtr
	if domain == 0 {
		switch strings.ToLower(conf.GetOrDefault(conf.GitDefaultDomain, "gitlab")) {
		case "github":
			domain = conf.GitDomainGithub
		case "gitlab":
			domain = conf.GitDomainGitlab
		}

	}
	if argsLen := len(os.Args); argsLen == 1 {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s <cmd:clone|pull> [options]:\n", os.Args[0])
		flag.PrintDefaults()
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
		log.Info("readying clone")
		if *all {
			log.Infof("running clone all repo, concurrent number: %d", *gn)
			cmd.CloneAllRepo(domain, *gn)
		} else if *search != "" {
			cmd.CloneOneRepoBySearch(domain, *search)
		} else {
			cmd.DefaultSelectOption(domain)
		}
	case "pull":
		log.Info("readying pull")
		if *all {
			log.Infof("running pull all repo, concurrent number: %d", *gn)
			cmd.PullAllRepo(domain, *gn)
		} else if *search != "" {
			cmd.PullOneRepoBySearch(domain, *search)
		} else {
			cmd.PullDefaultSelectOption(domain)
		}
	default:
		flag.Usage()
	}

	log.Info("process complete")
}
	
