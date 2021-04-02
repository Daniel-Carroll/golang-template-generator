package main

import (
	"os"

	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"gitlab.com/heb-engineering/templatr/app"
	"gitlab.com/heb-engineering/templatr/config"
)

// options are command-line options that are provided by the user.
type options struct {
	TemplateName    string `short:"t" long:"templateName" description:"Name of template"`
	OutputDirectory string `short:"o" long:"outputDirectory" description:"Set output directory for template generated code" default:"./"`
	AppName         string `short:"a" long:"appName" description:"Name of Application"`
	RepoUrl         string `short:"r" long:"repoUrl" description:"Repo URL for application"`
	Module          string `short:"m" long:"module" description:"go module name"`
}

// main runs the command-line parsing and validations. This function will also start the application logic execution.
func main() {
	// Parse command-line arguments
	var opts options
	args, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
	log.Info(args)

	// Convert to internal config
	cfg := config.New()
	cfg.AppName = opts.AppName
	cfg.OutputDirectory = opts.OutputDirectory
	cfg.RepoUrl = opts.RepoUrl
	cfg.Module = opts.Module

	templater, err := app.NewTemplaterFromConfig(cfg)

	err = templater.Start()
	if err != nil {
		// do stuff
		log.Error(err)
		os.Exit(1)
	}
}
