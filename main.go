package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type FluoriteConfig struct {
	Root    string   `yaml:"root"`
	Theme   string   `yaml:"theme"`
	Output  string   `yaml:"output"`
	Include []string `yaml:"include"`
	Lang    string   `yaml:"lang"`
}

var configFile = flag.String("c", "", "YAML config file path")
var obsidianRoot = flag.String("i", "", "obsidian notes folder")
var themeName = flag.String("t", "default", "theme name")
var outputFolder = flag.String("o", "output", "output folder")
var lang = flag.String("l", "en", "language")
var upgrade_theme = flag.Bool("upgrade", false, "Upgrade theme, not program")

var includeDirs []string

func main() {
	flag.Parse()
	if *configFile != "" {
		confs := FluoriteConfig{}
		conf, err := os.ReadFile(*configFile)
		IfFatal(err, "Read config error: "+*configFile)
		err = yaml.Unmarshal(conf, &confs)
		IfFatal(err, "Parse yaml config error")
		if *obsidianRoot == "" && confs.Root != "" {
			*obsidianRoot = confs.Root
		}
		if *themeName == "default" && confs.Theme != "" {
			*themeName = confs.Theme
		}
		if *outputFolder == "output" && confs.Output != "" {
			*outputFolder = confs.Output
		}
		if *lang == "en" && confs.Lang != "" {
			*lang = confs.Lang
		}
		includeDirs = confs.Include
	}
	if *upgrade_theme || !FileIsExisted("theme") {
		fmt.Println("Downloading theme...")
		DownloadFile("https://github.com/dicarne/fluorite/releases/latest/download/theme.zip", "theme.zip")
		p := path.Join("theme", "default")
		if FileIsExisted(p) {
			os.RemoveAll(p)
		}
		UnzipFile("theme.zip", "theme")
		fmt.Printf("Upgrade theme success!")
	}
	if *obsidianRoot == "" {
		fmt.Println("Must specify a Obsidian folder")
		return
	}
	generateObsidianValt(*obsidianRoot, *outputFolder, *themeName)
	fmt.Println("Completed! Output is in: " + *outputFolder)
}
