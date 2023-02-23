package api

type LogRepository interface {
	Add(schema *LogPostSchema) error
	Get(id int) (*LogGetSchema, error)
	GetAll() ([]*LogGetSchema, error)
	Delete(id int) error
}

type ReceiverRepository interface {
	Add(schema *ReceiverPostSchema) error
	Get(id int) (*ReceiverGetSchema, error)
	GetAll() ([]*ReceiverGetSchema, error)
	Delete(id int) error
}

type RequestRepository interface {
	Add(schema *RequestPostSchema) error
	Get(id int) (*RequestGetSchema, error)
	GetAll() ([]*RequestGetSchema, error)
	Delete(id int) error
}
