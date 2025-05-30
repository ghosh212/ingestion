package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"time"

	"tryout.com/ingestion/config"
	"tryout.com/ingestion/model"
	"tryout.com/ingestion/services"
)

func main() {
	fmt.Println("Hello, Go!")

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	ctx := context.Background()

	ingestedData, err := services.FetchLogsFromAPI(config.PublicAPIURL)
	if err != nil {
		log.Fatalf("Error fetching logs: %v", err)
	}

	if len(ingestedData) == 0 {
		log.Println("No log entries fetched...")
		return
	}

	processedData := services.ProcessLogData(ingestedData)

	objectName := fmt.Sprintf("logs/%s_api_logs.json", time.Now().Format("2006-01-02 15:04:05"))

	err = services.SaveLogEntriesToCloud(ctx, config.GCPProjectID, config.GCSBucketName, objectName, config.ServiceAccountKeyPath, convertStructToBytes(processedData))

	if err != nil {
		log.Fatalf("Error uploading logs to Google Cloud : %v", err)
	}

	log.Println("Data ingestion process completed successfully!")
}

func convertStructToBytes(processedData []model.LogEntry) []byte {

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(processedData)
	if err != nil {
		log.Fatalf("Error encoding slice of structs to Byte Array: %v", err)
	}

	return buf.Bytes()
}
