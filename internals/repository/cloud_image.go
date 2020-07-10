package repository

import (
	"context"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gofrs/uuid"
)

// CloudStorageImages is the dependency of adding images into the cloud
type CloudStorageImages struct {
	StorageBucket     *storage.BucketHandle
	StorageBucketName string
}

// NewCloudStorage .
func NewCloudStorage() *CloudStorageImages {
	ctx := context.Background()
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		panic(err)
	}
	bucketName := os.Getenv("BUCKET_NAME")
	bucket := storageClient.Bucket(bucketName)
	return &CloudStorageImages{StorageBucket: bucket, StorageBucketName: bucketName}
}

// UploadFile uploads a file into the cloud
func (c *CloudStorageImages) UploadFile(ctx context.Context, r io.Reader, contentType string) (string, error) {
	if contentType != "image/png" && contentType != "image/jpeg" {
		return "", fmt.Errorf("bad content type: %s", contentType)
	}
	name := uuid.Must(uuid.NewV4()).String()
	w := c.StorageBucket.Object(name).NewWriter(ctx)
	w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	w.ContentType = contentType
	w.CacheControl = "public, max-age=86400"
	if _, err := io.Copy(w, r); err != nil {
		return "", nil
	}
	if err := w.Close(); err != nil {
		return "", err
	}
	const publicURL = "https://storage.googleapis.com/%s/%s"
	return fmt.Sprintf(publicURL, c.StorageBucketName, name), nil
}
