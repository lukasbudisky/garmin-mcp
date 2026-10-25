package garmin

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Auth struct {
	Token *string
}

func GetAccessToken(apiKey string) (*string, error) {
	url := fmt.Sprintf("%s://%s%s", protocol, garmin_url, endpoint_get_api_key)

	log.Println("garmin authentification")
	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create authentification request: %v", err)
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("ApiKey %s", apiKey))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("cannot call authentification request: %v", err)
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse authentification data: %v", err)
	}

	token := string(responseData)

	return &token, nil
}
