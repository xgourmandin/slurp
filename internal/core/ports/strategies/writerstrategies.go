package strategies

type WriterStrategy interface {
	StoreApiResult(data interface{}) WriterStrategy
	Finalize() error
}
