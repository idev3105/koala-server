package storage

import (
	"context"
	"os"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func TestSave(t *testing.T) {
	minioClient, err := minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("VyFSnAYniNzObtDYBEZm", "fuEbxCsFXa3PxFkF2vi9wFcFjui07ETgbIp93dRH", ""),
		Secure: false,
	})
	if err != nil {
		t.Fatal(err)
	}
	cli, err := NewStorageClient(minioClient)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.TODO()
	data, err := os.ReadFile("./test.txt")
	if err != nil {
		t.Fatal(err)
	}
	err = cli.Save(ctx, "test", "test.txt", data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	minioClient, err := minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("VyFSnAYniNzObtDYBEZm", "fuEbxCsFXa3PxFkF2vi9wFcFjui07ETgbIp93dRH", ""),
		Secure: false,
	})
	if err != nil {
		t.Fatal(err)
	}
	cli, err := NewStorageClient(minioClient)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.TODO()
	data, err := cli.Get(ctx, "test", "test.txt")
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("data len must be greater than 0")
	}
}
