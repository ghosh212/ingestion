package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"tryout.com/ingestion/model"
)

func FetchLogsFromAPI(apiEndpoint string) ([]model.APIPost, error) {

	log.Printf("Fetching logs from API: %s\n", apiEndpoint)

	var result []model.APIPost

	resp, err := http.Get(apiEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error in response status code: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling API response: %v", err)
	}

	log.Printf("Successfully fetched %d logs from API.\n", len(result))
	return result, nil
}
