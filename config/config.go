package config

// Config provides a configuration structure used within the Efs2 application.
type Config struct {

	// TemplateName : name of template to be applied
	TemplateName string

	// OutputDirectory : directory where the generated template will go
	OutputDirectory string

	// AppName : name of your application
	AppName string

	// RepoUrl : url where your repo will reside
	RepoUrl string

	// Module : golang module name
	Module string

	// Namespace : Kubernetes namespace where project will live
	Namespace string
}

// New will return Config populated with pre-defined defaults.
func New() Config {
	c := Config{}
	c.OutputDirectory = "./"
	c.Module = "gitlab.com/heb-engineering"
	return c
}
