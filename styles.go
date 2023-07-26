package main

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type StyleConfig struct {
	Name string   `yaml:"name"`
	Css  []string `yaml:"css"`
	Js   []string `yaml:"js"`
}

func readStyle(themeName string) StyleConfig {
	themeConfig := path.Join("theme", themeName, "config.yaml")
	confs := StyleConfig{}
	conf, err := os.ReadFile(themeConfig)
	IfFatal(err, "Read theme config error: "+themeConfig)
	err = yaml.Unmarshal(conf, &confs)
	IfFatal(err, "Parse theme config error")
	return confs
}
