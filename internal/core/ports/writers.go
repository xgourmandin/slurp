package ports

type ApiDataWriter interface {
	StoreApiResult(data interface{})
}
