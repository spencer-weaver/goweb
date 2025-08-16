package html

import (
	"encoding/json"
	"fmt"
	t "html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type HTMLConfig struct {
	Render   RenderConfig `json:"render"`
	GoConfig GoConfig     `json:"goConfig"`
}

type RenderConfig struct {
	IncludeStandard     bool   `json:"includeStandard"`
	IncludeExperimental bool   `json:"IncludeExperimental"`
	IncludeDeprecated   bool   `json:"IncludeDeprecated"`
	BaseDir             string `json:"baseDir"`
	OutputFile          string `json:"outputFile"`
	ElementTemplate     string `json:"elementTemplate"`
}

type GoConfig struct {
	OutputFile     string `json:"outputFile"`
	StructTemplate string `json:"structTemplate"`
}

type MDNData struct {
	HTML struct {
		Elements map[string]ElementData `json:"elements"`
	} `json:"html"`
}

type ElementData struct {
	Compat CompatData `json:"__compat"`
}

type CompatData struct {
	Status Status `json:"status"`
}

type Status struct {
	Experimental  bool `json:"experimental"`
	StandardTrack bool `json:"standard_track"`
	Deprecated    bool `json:"deprecated"`
}

var HtmlConfig *HTMLConfig

func MinifyHTML(input string) string {
	var out strings.Builder
	inTag := false
	inText := false
	var textBuffer strings.Builder

	for i := 0; i < len(input); i++ {
		c := input[i]

		switch c {
		case '<':
			if inText {
				text := strings.TrimSpace(textBuffer.String())
				if len(text) > 0 {
					out.WriteString(text)
				}
				textBuffer.Reset()
				inText = false
			}
			inTag = true
			out.WriteByte(c)

		case '>':
			inTag = false
			out.WriteByte(c)
			inText = true

		default:
			if inTag {
				out.WriteByte(c)
			} else if inText {
				textBuffer.WriteByte(c)
			}
		}
	}

	if inText && textBuffer.Len() > 0 {
		out.WriteString(strings.TrimSpace(textBuffer.String()))
	}

	return strings.ReplaceAll(out.String(), "> <", "><")
}

func RenderHTMLElementTemplates(config *RenderConfig) error {

	// for parsing MDN's list of currently supported elements
	var mdnJSON MDNData

	loaded := make(map[string]bool)

	// loads the general html element template
	elementTemplate, err := os.ReadFile(config.ElementTemplate)
	if err != nil {
		log.Fatalf("template file %s read failed: %v", config.ElementTemplate, err)
		return err
	}

	// opens output file for writing, clearing previous (better back it up)
	f, err := os.OpenFile(config.OutputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("output file %s open failed: %v", config.OutputFile, err)
		return err
	}
	defer f.Close()

	err = filepath.WalkDir(config.BaseDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}

		base := filepath.Base(path)
		elementName := strings.TrimSuffix(base, ".json")

		jsonContent, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", path, err)
		}

		if err := json.Unmarshal(jsonContent, &mdnJSON); err != nil {
			return fmt.Errorf("failed to parse JSON from %s: %w", path, err)
		}

		component := NewComponent()
		component.Node = *NewNode(elementName)

		if mdnJSON.HTML.Elements[elementName].Compat.Status.StandardTrack && config.IncludeStandard {
			component.Render.Status = "standard"
		} else if mdnJSON.HTML.Elements[elementName].Compat.Status.Experimental && config.IncludeExperimental {
			component.Render.Status = "experimental"
		} else if mdnJSON.HTML.Elements[elementName].Compat.Status.Deprecated && config.IncludeDeprecated {
			component.Render.Status = "deprecated"
		}
		// see if self closing
		if IsSelfClosing(component.Name) {
			component.Render.SelfClosing = true
		}
		if loaded[component.Name] {
			fmt.Printf("duplicate element: %s\r\n", component.Name)
			return nil
		}
		if component.Render.Status != "" {
			var buf string = fmt.Sprintf(string(elementTemplate), component.Name, component.Name, component.Name)

			_, err = f.WriteString(buf)
			if err != nil {
				return fmt.Errorf("failed to write template for %s: %w", component.Name, err)
			}
			loaded[component.Name] = true
			fmt.Printf("wrote template for %s\r\n", component.Name)
		} else {
			fmt.Printf("unwanted element by status: %s\r\n", component.Name)
		}

		return nil
	})
	if err != nil {
		fmt.Printf("walkdir error: %v", err)
		return err
	}

	// extra elements not in jsons
	if loaded["svg"] {
		fmt.Printf("duplicate element: svg\r\n")
	} else {
		var buf string = fmt.Sprintf(string(elementTemplate), "svg", "svg", "svg")
		f.WriteString(buf)
	}
	if loaded["path"] {
		fmt.Printf("duplicate element: path\r\n")
		return nil
	} else {
		var buf string = fmt.Sprintf(string(elementTemplate), "path", "path", "path")
		f.WriteString(buf)
	}

	doctype := "{{ define \"doctype-html\" }}<!doctype html>\n{{ end }}\n"
	_, err = f.WriteString(doctype)
	if err != nil {
		return fmt.Errorf("failed to write template for doctype: %w", err)
	}
	document := "{{ define \"document-html\" }}{{ if .InnerHTML }}{{  .InnerHTML  }}{{ end }}{{ end }}\n"
	_, err = f.WriteString(document)
	if err != nil {
		return fmt.Errorf("failed to write template for doctype: %w", err)
	}

	return nil
}

func RenderGoDataTemplates(config *GoConfig, elementsFile string) error {

	var goStructTemplate string = "go-struct-html"

	var elements struct {
		Elements []string
	}

	structTemplate, err := os.ReadFile(config.StructTemplate)
	if err != nil {
		log.Fatalf("template file %s read failed: %v", config.StructTemplate, err)
		return err
	}

	tmpl, err := t.New(goStructTemplate).Funcs(t.FuncMap{"title": cases.Title(language.English).String, "lower": strings.ToLower}).Parse(string(structTemplate))
	if err != nil {
		log.Fatalf("template string %v load failed: %v", structTemplate, err)
		return err
	}

	elementsTemplatesBytes, err := os.ReadFile(elementsFile)
	if err != nil {
		log.Fatalf("template file %s read failed: %v", elementsFile, err)
		return err
	}

	tmpl, err = tmpl.Parse(string(elementsTemplatesBytes))
	if err != nil {
		log.Fatalf("template string %s load failed: %v", string(elementsTemplatesBytes), err)
		return err
	}

	i := 0
	for _, template := range tmpl.Templates() {
		if template.Name() == goStructTemplate {
			continue
		}
		fmt.Printf("loading template: %v\r\n", template.Name())
		trimmed := strings.TrimSuffix(template.Name(), "-html")
		elements.Elements = append(elements.Elements, trimmed)
		i++
	}

	outputFile, err := os.OpenFile(config.OutputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("output file %s open failed: %v", config.OutputFile, err)
		return err
	}

	err = tmpl.ExecuteTemplate(outputFile, goStructTemplate, elements)
	if err != nil {
		log.Fatalf("template %s execution failed: %v", goStructTemplate, err)
		return err
	}

	fmt.Printf("generated %v structures\r\n", i)

	return nil
}
