package api

type Scheduler interface {
	AddJob(schema ReceiverPostSchema)
}

type IO interface {
	Update()
}
