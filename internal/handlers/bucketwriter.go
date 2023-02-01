package handlers

type GcsStorageWriter struct {
	Format     string
	BucketName string
	FileName   string
}

func (s GcsStorageWriter) StoreApiResult(data interface{}) {

}
