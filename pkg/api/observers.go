package api

type Observer interface {
	Notify(schedulers []Scheduler, event RequestPostSchema) error
}

type Scheduler interface {
	AddJob(schema *RequestPostSchema) error
}

type IO interface {
	Update(schema *RequestGetSchema) error
}
