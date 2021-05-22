package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.com/heb-engineering/templatr/app"
	"gitlab.com/heb-engineering/templatr/config"
	"gitlab.com/heb-engineering/templatr/renderer"
)

// options are command-line options that are provided by the user.
type options struct {
	TemplateName    string `short:"t" long:"templateName" description:"Name of template"`
	OutputDirectory string `short:"o" long:"outputDirectory" description:"Set output directory for template generated code" default:"./"`
}

// main runs the command-line parsing and validations. This function will also start the application logic execution.
func main() {
	// Parse command-line arguments
	var opts options
	args, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
	log.Info(opts)
	log.Info(args)

	// Load context from YAML
	viper.SetConfigName("context")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	log.Info(viper.AllSettings())

	// Convert to internal config
	cfg := config.New()
	log.Info(opts.OutputDirectory)
	cfg.Values = viper.AllSettings()
	if len(opts.OutputDirectory) == 0 {
		cfg.OutputDirectory = "./output"
	} else {
		cfg.OutputDirectory = opts.OutputDirectory
	}

	log.Info(cfg.OutputDirectory)
	renderer, err := renderer.NewRendererFromConfig(cfg)

	templater := app.NewApp(renderer)

	err = templater.Start()
	if err != nil {
		// do stuff
		log.Error(err)
		os.Exit(1)
	}
}
