package renderer

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"path/filepath"

	"github.com/flosch/pongo2"
	log "github.com/sirupsen/logrus"
	"gitlab.com/heb-engineering/templatr/config"
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
	baseOutputFilepath := filepath.Join(r.Config.OutputDirectory)
	log.Info(baseOutputFilepath)

	c, err := ioutil.ReadDir(baseTemplateFilepath)
	log.Info(c)

	log.Info("Printing directory contents")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	err = filepath.WalkDir(baseTemplateFilepath, r.renderFile)

	cmdFilepath := filepath.Join(baseTemplateFilepath, "go", "cmd", "example", "_main.go")
	cmdOutputPath := filepath.Join(baseOutputFilepath, r.Config.AppName, "cmd", "example", "main.go")
	//dat, err := ioutil.ReadFile(cmdFilepath)
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info(cmdFilepath)
	log.Info(cmdOutputPath)

	return nil
}

func (r Renderer) renderFile(path string, d os.DirEntry, err error) error {
	if d.IsDir() {
		log.Info("Is Directory")
		tpl, err := pongo2.FromString(path)
		if err != nil {
			return err
		}

		out, err := tpl.Execute(r.Context)
		if err != nil {
			return err
		}
		log.Info("Output:", out)

		err = os.MkdirAll(filepath.Join(r.Config.OutputDirectory, out), 0755)
		if err != nil {
			return err
		}

	} else {
		tpl, err := pongo2.FromFile(path)
		if err != nil {
			return err
		}

		out, err := tpl.Execute(r.Context)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(filepath.Join(r.Config.OutputDirectory, path), []byte(out), 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// Creates a directory at the specified path, returns error if applicable
func createDirectory(path string, permissions fs.FileMode) error {
	err := os.MkdirAll(path, permissions)
	if err != nil {
		return err
	}

	return nil
}
