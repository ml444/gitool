package git

import (
	"errors"
	"fmt"
	"github.com/ml444/gitool/conf"
	"path/filepath"
	"strings"
)

func GetRepoPathBySSHURL(SSHURL string) (string, error) {
	if SSHURL == "" {
		return "", errors.New("this ssh url is null")
	}
	s := strings.TrimPrefix(SSHURL, "git@")
	s = strings.TrimSuffix(s, ".git")
	var sList []string
	list := strings.Split(s, ":")
	if len(list) != 2 {
		return "", fmt.Errorf("this ssh url [%s] is error", SSHURL)
	}
	sList = append(sList, list[0])
	list = strings.Split(list[1], "/")
	if len(list) > 0 {
		sList = append(sList, list...)
	}
	return filepath.Join(sList...), nil
}
func GetRepoPathByHTTPSURL(HTTPSURL string) (string, error) {
	if HTTPSURL == "" {
		return "", errors.New("this ssh url is null")
	}
	s := strings.TrimPrefix(HTTPSURL, "https://")
	s = strings.TrimSuffix(s, ".git")
	var sList []string
	list := strings.Split(s, "/")
	if len(list) <= 1 {
		return "", fmt.Errorf("this https url [%s] is error", HTTPSURL)
	}
	sList = append(sList, list...)
	return filepath.Join(sList...), nil
}
func getRepoLocalPath(httpsUrl string) (string, error) {
	dir, err := GetRepoPathByHTTPSURL(httpsUrl)
	if err != nil {
		return "", err
	}
	return filepath.Join(conf.Get(conf.GitlabLocalBaseDir), dir), nil
}
