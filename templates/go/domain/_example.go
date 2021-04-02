package domain

// Example : a simple example struct
type Example struct {
	ExampleID int `json:"exampleId" db:"example_id"`
}

// EFCService : interface to be implemented
type ExampleService interface {
	Examples() (*[]Example, error)
	Example(exampleId int) (*Example, error)
	CreateExample(example Example) (*Example, error)
	UpdateExample(example Example) (*Example, error)
	DeleteExample(exampleId int) error
}
