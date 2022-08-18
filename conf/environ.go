package conf

import "os"

var Config = map[string]string{}
var envMap = map[string]string{
	"GITLAB_BASE_URL":       "Please input GITLAB_BASE_URL",
	"GITLAB_USERNAME":       "Please input GITLAB_USERNAME",
	"GITLAB_ACCESS_TOKEN":   "Please input GITLAB_ACCESS_TOKEN",
	"GITLAB_LOCAL_BASE_DIR": "Please input GITLAB_LOCAL_BASE_DIR",
}

func InitEnv() error {
	for envName := range envMap {
		v := os.Getenv(envName)
		if v != "" {
			Config[envName] = v
		}
	}
	return nil
}
