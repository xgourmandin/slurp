package handlers

import (
	"encoding/json"
	"log"
	"os"
)

type FileWriter struct {
	FileName string
	Format   string
}

func (w FileWriter) StoreApiResult(data interface{}) {
	file, err := getFile(w.FileName)
	if err != nil {
		log.Fatalf("An error occured while appending to file %s : %v", w.FileName, err)
	}
	json, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error during conversion to JSON: %v", err)
	}
	file.Write(json)
	file.WriteString("\n")
}

func getFile(filename string) (*os.File, error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return f, nil
}
