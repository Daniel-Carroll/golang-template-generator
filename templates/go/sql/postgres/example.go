package postgres

import (
	log "github.com/sirupsen/logrus"
	"poop-module/domain"
)

// Ensure that ExampleService is implemented
var _ domain.ExampleService = &ExampleService{}

// ExampleService : implementation
type ExampleService struct {
	db *DB
}

// NewExampleService : creates a new instance of the postgres RouteService implementation
func NewExampleService(db *DB) *ExampleService {
	return &ExampleService{db: db}
}

// Examples :
func (s *ExampleService) Examples() (*[]domain.Example, error) {
	var Examples []domain.Example
	err := s.db.Select(&Examples,
		`SELECT example_jd
			FROM example
			ORDER BY example_id ASC`)
	if err != nil {
		log.Fatal(err)
	}

	return &Examples, nil
}

// Example : return one Example record by corp id
func (s *ExampleService) Example(exampleId int) (*domain.Example, error) {
	var Example domain.Example
	err := s.db.Get(&Example, "SELECT example_id from example where example_id = $1", exampleId)
	if err != nil {
		log.Error("Failed to query settings for Example with ID: ", exampleId)
		return nil, err
	}

	return &Example, nil
}

// CreateExample : Create a new Example record
func (s *ExampleService) CreateExample(Example domain.Example) (*domain.Example, error) {
	sqlStatement := `INSERT INTO example(
		example_id) 
		VALUES (:example)`
	_, err := s.db.NamedExec(sqlStatement, &Example)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &Example, nil
}

// UpdateExample : Update an existing Example record
func (s *ExampleService) UpdateExample(Example domain.Example) (*domain.Example, error) {
	// variable will store newly created row id
	_, err := s.db.NamedExec(`UPDATE example SET
		example_id = :example_id
		WHERE example_id = :example_id`, &Example)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &Example, nil
}

// DeleteExample : Delete an existing Example record
func (s *ExampleService) DeleteExample(exampleId int) error {
	_, err := s.db.Exec(
		`DELETE from example
		 WHERE	example_id = $1`,
		&exampleId)
	if err != nil {
		log.Printf("Failed to update delete windows from zone with id: %v", &exampleId)
		return err
	}
	return nil
}
