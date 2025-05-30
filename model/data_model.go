package model

import "time"

// APIPost represents the structure of a post fetched from the public API.
type APIPost struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_Id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type LogEntry struct {
	SourceID       int       `json:"source_Id"`
	SourceUserID   int       `json:"source_user_Id"`
	Title          string    `json:"title"`
	Body           string    `json:"body"`
	ProcessedAt    time.Time `json:"processed_at"`
	IngestionBatch string    `json:"ingestion_Batch"` // Identifier for the ingestion run
}

type Config struct {
	PublicAPIURL          string
	GCPProjectID          string
	GCSBucketName         string
	ServiceAccountKeyPath string
}
