package strategies

import (
	"fmt"
	"github.com/xgourmandin/slurp/configuration"
	"github.com/xgourmandin/slurp/internal/core/ports/strategies"
	"github.com/xgourmandin/slurp/internal/handlers/strategies/writers"
)

func NewWriterStrategy(apiConfig configuration.ApiConfiguration) strategies.WriterStrategy {
	switch apiConfig.OutputConfig.OutputType {
	case "FILE":
		return writers.FileWriter{
			Format:   "json",
			FileName: apiConfig.OutputConfig.FileName,
		}
	case "BUCKET":
		return writers.GcsStorageWriter{
			Format:     "json",
			BucketName: apiConfig.OutputConfig.BucketName,
			FileName:   apiConfig.OutputConfig.FileName,
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
