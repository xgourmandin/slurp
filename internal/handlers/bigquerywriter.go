package handlers

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
	"log"
	"os"
	"slurp/internal/core/ports"
)

type BigQueryWriter struct {
	Project          string
	Dataset          string
	Table            string
	AutodetectSchema bool
	TmpFile          string
	tmpWriter        FileWriter
}

func NewBigQueryWriter(project string, dataset string, table string, autodetect bool, tmpfile string) BigQueryWriter {
	return BigQueryWriter{
		Project:          project,
		Dataset:          dataset,
		Table:            table,
		AutodetectSchema: autodetect,
		TmpFile:          tmpfile,
		tmpWriter: FileWriter{
			FileName: tmpfile,
			Format:   "json",
		},
	}
}

func (w BigQueryWriter) StoreApiResult(data interface{}) ports.ApiDataWriter {
	w.tmpWriter.StoreApiResult(data)
	return w
}

func (w BigQueryWriter) Finalize() error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, w.Project)
	if err != nil {
		log.Println(fmt.Errorf("bigquery.NewClient: %v", err))
	}
	defer client.Close()

	tmpFile, err := os.Open(w.TmpFile)
	if err != nil {
		return err
	}
	source := bigquery.NewReaderSource(tmpFile)
	source.SourceFormat = bigquery.JSON
	source.AutoDetect = w.AutodetectSchema
	loader := client.Dataset(w.Dataset).Table(w.Table).LoaderFrom(source)
	loader.WriteDisposition = bigquery.WriteEmpty

	job, err := loader.Run(ctx)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}
	status, err := job.Wait(ctx)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}

	if status.Err() != nil {
		log.Println(fmt.Errorf("job completed with error: %v", status.Err()))
		return status.Err()
	}
	return nil
}
