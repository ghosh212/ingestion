package config

import (
	"encoding/json"
	"log"
	"os"

	"tryout.com/ingestion/model"
)

// loads config from json value
func LoadConfig() (*model.Config, error) {

	configFile, err := os.ReadFile("configs.json")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var config model.Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return &model.Config{
		PublicAPIURL:          config.PublicAPIURL,
		GCPProjectID:          config.GCPProjectID,
		GCSBucketName:         config.GCSBucketName,
		ServiceAccountKeyPath: config.ServiceAccountKeyPath,
	}, nil
}
