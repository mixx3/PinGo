package api

type LogService interface {
	Create(schema *LogPostSchema) error
	GetAll() ([]*LogGetSchema, error)
	Get(id int) (*LogGetSchema, error)
	Delete(id int) error
}

type RequestService interface {
	Create(schema *RequestPostSchema) error
	GetAll() ([]*RequestGetSchema, error)
	Get(id int) (*RequestGetSchema, error)
	Delete(id int) error
}

type ReceiverService interface {
	Create(schema *ReceiverPostSchema) error
	GetAll() ([]*ReceiverGetSchema, error)
	Get(id int) (*ReceiverGetSchema, error)
	Delete(id int) error
}
