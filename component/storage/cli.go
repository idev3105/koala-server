package storage

import (
	"bytes"
	"context"
	"io"

	"github.com/minio/minio-go/v7"
)

const CHUNK_SIZE = 1024 * 1024

type StorageClient struct {
	minioClient *minio.Client
}

func NewStorageClient(minioClient *minio.Client) (*StorageClient, error) {
	return &StorageClient{minioClient: minioClient}, nil
}

func (c *StorageClient) Save(ctx context.Context, bucketName, objectName string, data []byte) error {
	_, err := c.minioClient.PutObject(ctx, bucketName, objectName, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{})
	if err != nil {
		return nil
	}
	return nil
}

func (c *StorageClient) Get(ctx context.Context, bucketName, objectName string) ([]byte, error) {
	obj, err := c.minioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer obj.Close()

	var data []byte
	buffer := make([]byte, CHUNK_SIZE)

	for {
		n, err := obj.Read(buffer)
		if n > 0 {
			data = append(data, buffer[:n]...)
		}
		if err != nil {
			if err == io.EOF {
				break // End of file reached
			}
			return nil, err // Return error if it's not EOF
		}
	}
	return data, nil
}
