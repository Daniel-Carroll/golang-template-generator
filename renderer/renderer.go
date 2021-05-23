package renderer

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"

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
	log.Info("Template Name: ", r.Config.TemplateName)
	baseTemplateFilepath := filepath.Join("templates", r.Config.TemplateName)
	baseOutputFilepath := filepath.Join(r.Config.OutputDirectory)
	log.Info("Base Output", baseOutputFilepath)

	err := filepath.WalkDir(baseTemplateFilepath, r.renderFile)

	//dat, err := ioutil.ReadFile(cmdFilepath)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (r Renderer) renderFile(path string, d os.DirEntry, err error) error {

	// Check if path is a directory
	if d.IsDir() {

		tpl, err := pongo2.FromString(path)
		if err != nil {
			return err
		}

		out, err := tpl.Execute(r.Context)
		if err != nil {
			return err
		}

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

		trimmedPath := trimOutputPath(path)

		log.Info("Untrimmed Path")
		log.Info(path)

		log.Info("Trimmed Path")
		log.Info(trimmedPath)
		err = ioutil.WriteFile(filepath.Join(r.Config.OutputDirectory, trimmedPath), []byte(out), 0755)
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

// Remove the template part of the path when creating directory
func trimOutputPath(path string) string {
	return strings.Replace(path, "/_", "/", 1)
}
