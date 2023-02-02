package handlers

import (
	"fmt"
	"slurp/internal/core/ports"
)

type LogWriter struct {
	Format     string
	BucketName string
	FileName   string
}

func (s LogWriter) StoreApiResult(data interface{}) ports.ApiDataWriter {
	fmt.Println(data)
	return s
}

func (s LogWriter) Finalize() error {
	return nil
}
