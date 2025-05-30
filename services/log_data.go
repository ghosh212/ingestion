package services

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func SaveLogEntriesToCloud(ctx context.Context, projectID, bucketName, objectName, saKeyPath string, logs []byte) error {

	if len(logs) == 0 {
		log.Println("No logs to save to Firestore.")
		return nil
	}

	var client *storage.Client
	var err error

	if saKeyPath != "" {
		client, err = storage.NewClient(ctx, option.WithCredentialsFile(saKeyPath))
	} else {
		client, err = storage.NewClient(ctx)
	}

	if err != nil {
		return fmt.Errorf("failed to create GCS client: %w", err)
	}
	defer client.Close()

	wc := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	wc.ContentType = "application/json"

	if _, err := wc.Write(logs); err != nil {
		return fmt.Errorf("failed to write data to GCS: %w", err)
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("failed to close GCS writer: %w", err)
	}

	log.Printf("Data successfully uploaded to gs://%s/%s", bucketName, objectName)
	return nil
}
