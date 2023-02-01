package handlers

import "fmt"

type LogWriter struct {
	Format     string
	BucketName string
	FileName   string
}

func (s LogWriter) StoreApiResult(data interface{}) {
	fmt.Println(data)
}
