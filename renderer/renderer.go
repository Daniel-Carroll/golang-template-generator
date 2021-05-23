package renderer

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"path/filepath"

	"github.com/Daniel-Carroll/golang-template-generator/config"
	"github.com/flosch/pongo2"
	log "github.com/sirupsen/logrus"
)

var ErrShutdown = fmt.Errorf("application was shutdown gracefully")

type Renderer struct {
	Context pongo2.Context
	Config  config.Config
}

// Initialize new templater with desired context
func NewRendererFromConfig(cfg config.Config) (Renderer, error) {
	var templater Renderer
	templater.Context = cfg.Values
	templater.Config = cfg

	return templater, nil
}

func (r Renderer) Render() error {
	// Application runtime code goes here
	baseTemplateFilepath := filepath.Join("templates", r.Config.TemplateName)

	err := filepath.WalkDir(baseTemplateFilepath, r.renderFile)

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (r Renderer) renderFile(templatePath string, d os.DirEntry, err error) error {
	outputPath := strings.Replace(templatePath, "templates/"+r.Config.TemplateName, "", 1)
	// Check if path is a directory
	if d.IsDir() {
		tpl, err := pongo2.FromString(outputPath)
		if err != nil {
			return err
		}

		out, err := tpl.Execute(r.Context)
		if err != nil {
			return err
		}

		// Create directory in output
		err = os.MkdirAll(filepath.Join(r.Config.OutputDirectory, out), 0755)
		if err != nil {
			return err
		}

	} else {
		pathTpl, err := pongo2.FromString(outputPath)
		if err != nil {
			return err
		}
		pathOut, err := pathTpl.Execute(r.Context)
		if err != nil {
			return err
		}

		tpl, err := pongo2.FromFile(templatePath)
		if err != nil {
			return err
		}

		out, err := tpl.Execute(r.Context)
		if err != nil {
			return err
		}

		trimmedPath := strings.Replace(pathOut, "/_", "/", 1)

		err = ioutil.WriteFile(filepath.Join(r.Config.OutputDirectory, trimmedPath), []byte(out), 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func trimDirectoryPath(path string) string {
	return strings.Replace(path, "templates/", "", 1)
}
