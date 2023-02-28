package handlers

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/json"
	"fmt"
	"github.com/xgourmandin/slurp/internal/core/ports"
	"log"
)

var formatToContentType = map[string]string{
	"json": "application/json",
	"xml":  "application/xml",
}

type GcsStorageWriter struct {
	Format     string
	BucketName string
	FileName   string
	Data       []interface{}
}

func (w GcsStorageWriter) StoreApiResult(data interface{}) ports.ApiDataWriter {
	w.Data = append(w.Data, data)
	return w
}

func (w GcsStorageWriter) Finalize() error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	wc := client.Bucket(w.BucketName).Object(w.FileName).NewWriter(ctx)
	if contentType, ok := formatToContentType[w.Format]; ok {
		wc.ContentType = contentType
	} else {
		log.Println(fmt.Sprintf("Output format %s unkown, defaulting to text/plain", w.Format))
		wc.ContentType = "text/plain"
	}

	if binData, err := marshallData(w.Data, w.Format); err != nil {
		if _, err := wc.Write(binData); err != nil {
			return err
		}
	} else {
		return err
	}

	if err := wc.Close(); err != nil {
		log.Printf("unable to close bucket %s, file %s: %v\n", w.BucketName, w.FileName, err)
		return nil
	}
	return nil
}

func marshallData(data []interface{}, format string) ([]byte, error) {
	switch format {
	case "json":
		return json.Marshal(data)
	default:
		return []byte(fmt.Sprint(data)), nil
	}
}
