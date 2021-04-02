package app

import (
	"fmt"
	"io/ioutil"
	"os"

	"path/filepath"

	"github.com/flosch/pongo2"
	log "github.com/sirupsen/logrus"
	"gitlab.com/heb-engineering/templatr/config"
)

var ErrShutdown = fmt.Errorf("application was shutdown gracefully")

type Templater struct {
	Context pongo2.Context
	Config  config.Config
}

// Initialize new templater with desired context
func NewTemplaterFromConfig(cfg config.Config) (Templater, error) {
	var templater Templater
	templater.Context = map[string]interface{}{
		"module":    cfg.Module,
		"repoUrl":   cfg.RepoUrl,
		"namespace": cfg.Namespace,
	}
	templater.Config = cfg

	return templater, nil
}

func (t Templater) Start() error {
	// Application runtime code goes here
	baseTemplateFilepath := filepath.Join("templates", t.Config.TemplateName)
	baseOutputFilepath := filepath.Join(t.Config.OutputDirectory)
	log.Info(baseOutputFilepath)

	c, err := ioutil.ReadDir(baseTemplateFilepath)
	log.Info(c)

	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	cmdFilepath := filepath.Join(baseTemplateFilepath, "cmd", "example", "_main.go")
	cmdOutputPath := filepath.Join(baseOutputFilepath, t.Config.AppName, "cmd", "example", "main.go")
	dat, err := ioutil.ReadFile(cmdFilepath)
	if err != nil {
		return err
	}
	log.Info(string(dat))

	// Compile the template first (i. e. creating the AST)
	tpl, err := pongo2.FromFile(cmdFilepath)
	if err != nil {
		panic(err)
	}

	out, err := tpl.Execute(t.Context)
	if err != nil {
		panic(err)
	}
	fmt.Println(out) // Output: Hello Florian!

	err = os.MkdirAll(filepath.Join(baseOutputFilepath, t.Config.AppName, "cmd", "example"), 0755)
	if err != nil {
		log.Error(err)
	}
	err = ioutil.WriteFile(cmdOutputPath, []byte(out), 0755)
	if err != nil {
		log.Error(err)
	}
	return nil
}

func Shutdown() {
	// Shutdown contexts, listeners, and such
}
