package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/spencer-weaver/goweb/internal/cmd/handler"
	"github.com/spencer-weaver/goweb/internal/template/html"
)

const ConfigFile string = "config.json"

type ConfigWrapper struct {
	Type        string              `json:"__type"`
	BuildConfig handler.BuildConfig `json:"buildConfig"`
	TestConfig  handler.TestConfig  `json:"testConfig"`
	HtmlConfig  html.HTMLConfig     `json:"htmlConfig"`
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage: goweb <options> -flags")
		os.Exit(1)
	}

	f, err := os.OpenFile(ConfigFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("config file %s open failed: %v\r\n", ConfigFile, err)
	}
	defer f.Close()

	jsonContent, err := os.ReadFile(ConfigFile)
	if err != nil {
		fmt.Printf("failed to read %s: %v\r\n", ConfigFile, err)
	}

	config := ConfigWrapper{}
	if err := json.Unmarshal(jsonContent, &config); err != nil {
		fmt.Printf("failed to parse JSON from %s: %v\r\n", ConfigFile, err)
	}

	handler.BuildSettings = &config.BuildConfig
	handler.TestSettings = &config.TestConfig
	handler.HtmlSettings = &config.HtmlConfig

	switch os.Args[1] {
	case "build":
		flags := flag.NewFlagSet("build", flag.ExitOnError)
		flags.String("configDir", config.BuildConfig.ConfigDir, "directory to define page layout in json")
		flags.String("outputDir", config.BuildConfig.OutputDir, "directory to output generated html")
		flags.String("templatesDir", config.BuildConfig.TemplatesDir, "directory with templates")
		//updated := flags.Bool("updated", , "only build updated files")
		flags.Parse(os.Args[2:])
		handler.BuildHandler()

	case "clean":

	case "test":
		flags := flag.NewFlagSet("test", flag.ExitOnError)
		flags.String("dir", config.TestConfig.WatchDir, "directory to serve")
		flags.String("port", config.TestConfig.ServePort, "port to run on")
		flags.Parse(os.Args[2:])
		handler.TestHandler()

	case "html":
		flags := flag.NewFlagSet("html", flag.ExitOnError)
		updateHtml := flags.Bool("updateHtml", false, "update internal html templates")
		updateGo := flags.Bool("updateGo", false, "update internal html templates")
		update := flags.Bool("update", false, "update internal html templates")
		flags.Parse(os.Args[2:])
		if *update {
			*updateHtml = true
			*updateGo = true
		}
		err := handler.HtmlHandler(*updateHtml, *updateGo)
		if err != nil {
			fmt.Print(err)
		}
	}
}
