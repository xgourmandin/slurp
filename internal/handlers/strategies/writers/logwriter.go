package writers

import (
	"fmt"
	"github.com/xgourmandin/slurp/internal/core/ports/strategies"
)

type LogWriter struct {
	Format     string
	BucketName string
	FileName   string
}

func (s LogWriter) StoreApiResult(data interface{}) strategies.WriterStrategy {
	fmt.Println(data)
	return s
}

func (s LogWriter) Finalize() error {
	return nil
}
