package utils

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Conf struct {
	GoLandPath string `yaml:"goLandPath"`
	Comment    string `yaml:"comment"`
	Leetcode   struct {
		Path string `yaml:"Path"`
	} `yaml:"leetcode"`
	Codeforces struct {
		Path string `yaml:"Path"`
	} `yaml:"codeforces"`
}

var Config Conf

func InitSettings() Conf {
	yamlFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}
	return Config
}
