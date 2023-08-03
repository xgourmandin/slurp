package strategies

import (
	"fmt"
	"github.com/xgourmandin/slurp/configuration"
	"github.com/xgourmandin/slurp/internal/core/ports/strategies"
	"github.com/xgourmandin/slurp/internal/handlers/strategies/writers"
	"strings"
	"time"
)

func NewWriterStrategy(apiConfig configuration.ApiConfiguration) strategies.WriterStrategy {
	switch apiConfig.OutputConfig.OutputType {
	case "FILE":
		return writers.FileWriter{
			Format:   "json",
			FileName: apiConfig.OutputConfig.FileName,
		}
	case "BUCKET":
		chunked := strings.Split(apiConfig.OutputConfig.FileName, ".")
		filename := strings.Join(chunked[:len(chunked)-1], ".") + "-" + time.Now().Format("20060201150405") + "." + chunked[len(chunked)-1]
		return writers.GcsStorageWriter{
			Format:     "json",
			BucketName: apiConfig.OutputConfig.BucketName,
			FileName:   filename,
		}
	case "BIGQUERY":
		return writers.NewBigQueryWriter(
			apiConfig.OutputConfig.Project,
			apiConfig.OutputConfig.Dataset,
			apiConfig.OutputConfig.Table,
			apiConfig.OutputConfig.Autodetect,
			fmt.Sprintf("/tmp/slurp-%s.json", apiConfig.OutputConfig.Table),
		)
	default:
		return writers.LogWriter{}
	}
}
