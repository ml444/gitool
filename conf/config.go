package conf

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

//func init() {
//	var err error
//	BaseDir, err = os.Getwd()
//	if err != nil {
//		fmt.Println("Error Cloning: " + err.Error())
//		return
//	}
//}

func Get(key string) string {
	v, ok := Config[key]
	if ok {
		return v
	}

	label, ok := envMap[key]
	if !ok {
		panic(fmt.Sprintf("Not found %s", key))
	}

	prompt := promptui.Prompt{
		Label: label,
		Validate: func(s string) error {
			return nil
		},
	}

	v, err := prompt.Run()
	if err != nil {
		println(err.Error())
		panic(err)
	}
	return v
}
