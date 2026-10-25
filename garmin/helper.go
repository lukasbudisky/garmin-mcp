package garmin

import (
	"fmt"
	"log"
	"net/http"
)

func (auth *Auth) addHeader(request *http.Request) {
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *auth.Token))
	request.Header.Add("Content-Type", "application/json")
}

func httpStatusCodeError(name string, response *http.Response) error {
	switch response.StatusCode {
	case http.StatusBadRequest:
		return fmt.Errorf("%s, bad request http_status: %d",
			name, response.StatusCode)
	case http.StatusUnauthorized:
		return fmt.Errorf("%s, missing or invalid authentification http_status: %d",
			name, response.StatusCode)
	case http.StatusForbidden:
		return fmt.Errorf("%s, unathorized http_status: %d",
			name, response.StatusCode)
	case http.StatusInternalServerError:
		return fmt.Errorf("%s, internal server error http_status: %d",
			name, response.StatusCode)
	case http.StatusNotFound:
		return fmt.Errorf("%s, not found http_status: %d",
			name, response.StatusCode)
	case http.StatusOK:
		log.Printf("successfully %s http_status: %d\n",
			name, response.StatusCode)
	default:
		return fmt.Errorf("%s, unexpected error http_status: %d",
			name, response.StatusCode)
	}
	return nil
}
