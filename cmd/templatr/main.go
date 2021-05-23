package main

import (
	"os"

	"github.com/Daniel-Carroll/golang-template-generator/app"
	"github.com/Daniel-Carroll/golang-template-generator/config"
	"github.com/Daniel-Carroll/golang-template-generator/renderer"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// options are command-line options that are provided by the user.
type options struct {
	TemplateName    string `short:"t" long:"templateName" description:"Name of template"`
	OutputDirectory string `short:"o" long:"outputDirectory" description:"Set output directory for template generated code" default:"./output"`
}

// main runs the command-line parsing and validations. This function will also start the application logic execution.
func main() {
	// Parse command-line arguments
	var opts options
	_, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}

	contextName := opts.TemplateName + "-context"
	// Load context from YAML
	viper.SetConfigName(contextName)
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		log.Info("Context with name ", contextName, " does not exist, defaulting to context.json")
		viper.SetConfigName("context")
		err = viper.ReadInConfig()
		if err != nil {
			log.Fatal("You fool, there is no valid context file. Please add a valid context.json and try again.")
		}
	}

	if len(opts.TemplateName) == 0 {
		log.Info("Please Specify Template Name")
		os.Exit(1)
	}
	// Convert to internal config
	cfg := config.New()
	cfg.Values = viper.AllSettings()
	cfg.OutputDirectory = opts.OutputDirectory
	cfg.TemplateName = opts.TemplateName

	renderer, err := renderer.NewRendererFromConfig(cfg)

	templater := app.NewApp(renderer)
	defer templater.Shutdown()

	err = templater.Start()
	if err != nil {
		// do stuff
		log.Error(err)
		os.Exit(1)
	}

	log.Info("Template Successfully Generated at: ", renderer.Config.TemplateName)
}
