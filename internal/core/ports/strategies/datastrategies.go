package strategies

type DataStrategy interface {
	ExtractData(body []byte, out chan interface{})
	GetResultSize(response []byte) int
	GetSingleValue(body []byte, path string) *string
}
