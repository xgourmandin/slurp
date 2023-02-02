package ports

type ApiDataWriter interface {
	StoreApiResult(data interface{}) ApiDataWriter
	Finalize() error
}
