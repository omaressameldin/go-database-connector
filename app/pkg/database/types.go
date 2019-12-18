package database

// ValidationError Validation error struct
type ValidationError struct {
	Field   string
	Message string
}

// Validator struct to link  validation error with field
type Validator struct {
	Field string
	Error error
}

// Updated struct for values to be updated
type Updated struct {
	Key string
	Val interface{}
}

// Operator enum to determine filter operator
type Operator string

const (
	// Equals Operator
	Equals Operator = "=="

	// NotEquals Operator
	NotEquals string = "!="
	// MoreThan Operator
	MoreThan = ">"
	// LessThan Operator
	LessThan = "<"
	// MoreThanOrEqual Operator
	MoreThanOrEqual = ">="
	// LessThanOrEqual Operator
	LessThanOrEqual = "<="
)

// Filter struct for filtering queries
type Filter struct {
	Operator Operator
	Field    string
	Value    interface{}
}

type User struct {
	ID     string
	Email  string
	Name   string
	Avatar string
}

// Connector database adapter interface
type Connector interface {
	CloseConnection() error
	Create(validators []Validator, key string, data interface{}) error
	Update(validators []Validator, key string, data []Updated) error
	Read(key string, model interface{}) error
	ReadAll(genRefFn func() interface{}, appendFn func(interface{}), filters []Filter) error
	Delete(key string) error
	Authenticate(token string) (*User, error)
}
