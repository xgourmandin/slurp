package handlers

import (
	"encoding/json"
	"log"
	"os"
	"slurp/internal/core/ports"
)

type FileWriter struct {
	FileName string
	Format   string
}

func (w FileWriter) StoreApiResult(data interface{}) ports.ApiDataWriter {
	file, err := getFile(w.FileName)
	if err != nil {
		log.Fatalf("An error occured while appending to file %s : %v", w.FileName, err)
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error during conversion to JSON: %v", err)
	}
	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatalf("Error writng to file %s: %v", w.FileName, err)
	}
	_, err = file.WriteString("\n")
	if err != nil {
		log.Fatalf("Error writing to file %s: %v", w.FileName, err)
	}
	return w
}

func getFile(filename string) (*os.File, error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return f, nil
}

func (w FileWriter) Finalize() error {
	return nil
}
