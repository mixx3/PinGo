package api

type LogRepository interface {
	Add(schema *LogPostSchema) error
	Get(id int) (*LogGetSchema, error)
	Delete(id int) error
}

type ReceiverRepository interface {
	Add(schema *LogPostSchema) error
	Get(id int) (*LogGetSchema, error)
	Delete(id int) error
}

type RequestRepository interface {
	Add(schema *LogPostSchema) error
	Get(id int) (*LogGetSchema, error)
	Delete(id int) error
}
