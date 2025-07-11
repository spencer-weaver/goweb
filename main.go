package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// build globals
var BuildSourceDir string = "./frontend/src/"
var BuildTemplateDir string = "./frontend/templates/"
var BuildOutputDir string = "./frontend/public/"
var BuildUpdated bool = false

// clean globals

// test globals
var TestDefaultDir string = BuildOutputDir
var TestDefaultPort string = "0.0.0.0:8080"

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage: goweb <options> -flags")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "build":
		flags := flag.NewFlagSet("build", flag.ExitOnError)
		sourceDir := flags.String("sourceDir", BuildSourceDir, "directory with source template")
		templateDir := flags.String("templateDir", BuildTemplateDir, "directory with template parts")
		outputDir := flags.String("outputDir", BuildOutputDir, "directory to output generated html")
		//updated := flags.Bool("updated", BuildUpdated, "only build updated files")
		flags.Parse(os.Args[2:])
		Build(*sourceDir, *templateDir, *outputDir)

	case "clean":

	case "test":
		flags := flag.NewFlagSet("test", flag.ExitOnError)
		dir := flags.String("dir", TestDefaultDir, "directory to serve")
		port := flags.String("port", TestDefaultPort, "port to run on")
		flags.Parse(os.Args[2:])
		TestServer(*dir, *port)
	}
}

func Build(sourceDir, templateDir, outputDir string) {
	// load templates from templateDir
	baseTmpl := template.New("").Funcs(template.FuncMap{})
	partials, err := filepath.Glob(filepath.Join(templateDir, "*.html"))
	if err != nil {
		log.Fatalf("load base templates: %v", err)
	}
	if len(partials) > 0 {
		baseTmpl, err = baseTmpl.ParseFiles(partials...)
		if err != nil {
			log.Fatalf("parse base templates: %v", err)
		}
	}

	// process each file in sourceDir
	err = filepath.WalkDir(sourceDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".html") {
			return nil
		}

		// clone the base templates
		tmpl, err := baseTmpl.Clone()
		if err != nil {
			return fmt.Errorf("clone templates: %w", err)
		}

		// parse the page file into the template
		tmpl, err = tmpl.ParseFiles(path)
		if err != nil {
			return fmt.Errorf("parse page %s: %w", path, err)
		}

		// create the output file
		outName := d.Name()
		outPath := filepath.Join(outputDir, outName)
		err = os.MkdirAll(filepath.Dir(outPath), 0755)
		if err != nil {
			return fmt.Errorf("mkdir output dir: %w", err)
		}
		outFile, err := os.Create(outPath)
		if err != nil {
			return fmt.Errorf("create output file: %w", err)
		}
		defer outFile.Close()

		// execute template
		data := map[string]any{}
		err = tmpl.ExecuteTemplate(outFile, filepath.Base(path), data)
		if err != nil {
			return fmt.Errorf("execute %s: %w", path, err)
		}

		log.Printf("rendered %s", outName)
		return nil
	})
	if err != nil {
		log.Fatalf("build failed: %v", err)
	}
}

func TestServer(dir, port string) {

	router := gin.Default()
	router.Static("/", dir)

	fmt.Printf("serving %s at http://%s\n", dir, port)
	if err := router.Run(port); err != nil {
		fmt.Println("error starting server:", err)
		os.Exit(1)
	}
}
