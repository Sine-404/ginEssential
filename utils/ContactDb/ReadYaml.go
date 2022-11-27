package ContactDb

import (
	"WebProject/ginEssential/model"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
)

// 读取yaml
func ReadYamlConfig() (yamlConfig *model.Config) {

	file, err := os.Open("./config.yaml")
	if err != nil {
		log.Println(err)
		return
	}

	yamls, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
	}

	var config model.Config

	yaml.Unmarshal(yamls, &config)

	return &config

}
