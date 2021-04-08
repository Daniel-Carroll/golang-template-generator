package config

// Config provides a configuration structure used within the Efs2 application.
type Config struct {

	// TemplateName : name of template to be applied
	TemplateName string

	// OutputDirectory : directory where the generated template will go
	OutputDirectory string

	// AppName : name of your application
	AppName string

	// Values :
	Values map[string]interface{}
}

// New will return Config populated with pre-defined defaults.
func New() Config {
	c := Config{}
	c.OutputDirectory = "./output"
	return c
}
