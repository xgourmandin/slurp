package writers

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/json"
	"fmt"
	"github.com/xgourmandin/slurp/internal/core/ports/strategies"
	"log"
	"strings"
	"time"
)

var formatToContentType = map[string]string{
	"json": "application/json",
	"xml":  "application/xml",
}

type GcsStorageWriter struct {
	Format     string
	BucketName string
	FileName   string
	writer     *storage.Writer
}

func (w GcsStorageWriter) StoreApiResult(data interface{}) strategies.WriterStrategy {
	if w.writer == nil {
		ctx := context.Background()
		client, err := storage.NewClient(ctx)
		if err != nil {
			log.Printf("%+v", err)
			return w
		}
		fileName := computeFilename(w.FileName)
		wc := client.Bucket(w.BucketName).Object(fileName).NewWriter(ctx)
		if contentType, ok := formatToContentType[w.Format]; ok {
			wc.ContentType = contentType
		} else {
			log.Println(fmt.Sprintf("Output format %s unkown, defaulting to text/plain", w.Format))
			wc.ContentType = "text/plain"
		}
		w.writer = wc
	}
	if binData, err := marshallData(data, w.Format); err == nil {
		if _, err := w.writer.Write(binData); err != nil {
			log.Printf("%+v", err)
		}
		if _, err := w.writer.Write([]byte("\n")); err != nil {
			log.Printf("%+v", err)
		}
	} else {
		log.Printf("%+v", err)
	}

	return w
}

func computeFilename(name string) string {
	if strings.Contains(name, "{extract_time}") {
		return strings.Replace(name, "{extract_time}", time.Now().UTC().Format(time.RFC3339), 1)
	} else {
		chunked := strings.Split(name, ".")
		filename := strings.Join(chunked[:len(chunked)-1], ".") + "-" + time.Now().UTC().Format("20060201150405") + "." + chunked[len(chunked)-1]
		return filename
	}
}

func (w GcsStorageWriter) Finalize() error {
	if err := w.writer.Close(); err != nil {
		log.Printf("unable to close bucket %s, file %s: %v\n", w.BucketName, w.FileName, err)
		return err
	} else {
		return nil
	}
}

func marshallData(data interface{}, format string) ([]byte, error) {
	switch format {
	case "json":
		return json.Marshal(data)
	default:
		return []byte(fmt.Sprint(data)), nil
	}
}
