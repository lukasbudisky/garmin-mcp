package garmin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Activity struct {
	ActivityId   int64  `json:"activityId,omitempty"`
	ActivityName string `json:"activityName,omitempty"`
}

func (auth *Auth) SearchActivities(start int, limit int) (*[]Activity, error) {
	url := fmt.Sprintf("%s://%s%s", protocol, garmin_url, endpoint_search_activities)

	if start <= 0 {
		start = 0
	}

	if limit <= 0 {
		limit = 5
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	auth.addHeader(request)

	query := request.URL.Query()
	query.Add("start", fmt.Sprintf("%d", start))
	query.Add("limit", fmt.Sprintf("%d", limit))
	request.URL.RawQuery = query.Encode()

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err = httpStatusCodeError("get activities data", response); err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot parse activities response data: %v", err)
	}

	var activities []Activity

	if err := json.Unmarshal(responseData, &activities); err != nil {
		return nil, fmt.Errorf("cannot parse activities structure data: %v", err)
	}

	return &activities, nil
}
