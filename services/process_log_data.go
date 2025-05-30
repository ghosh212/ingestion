package services

import (
	"log"
	"time"

	"tryout.com/ingestion/model"
)

func ProcessLogData(apiPosts []model.APIPost) []model.LogEntry {
	var processedEntries []model.LogEntry
	currentTime := time.Now()

	for _, post := range apiPosts {
		processedEntry := model.LogEntry{
			SourceID:     post.ID,
			SourceUserID: post.UserID,
			Title:        post.Title,
			Body:         post.Body,
			ProcessedAt:  currentTime,
		}
		processedEntries = append(processedEntries, processedEntry)
	}

	log.Printf("Processed %d logs.\n", len(processedEntries))

	return processedEntries
}
