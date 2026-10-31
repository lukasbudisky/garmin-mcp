package garmin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Activity struct {
	ActivityId   int64  `json:"activityId,omitempty"`
	ActivityName string `json:"activityName,omitempty"`
}

func (auth *Auth) CountAllActivities(paginationLimit int) (int, error) {
	var activities []Activity
	url := fmt.Sprintf("%s://%s%s", protocol, garmin_url, endpoint_search_activities)
	start := 0

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return -1, err
	}
	auth.addHeader(request)
	query := request.URL.Query()

	for {
		var tempActivites []Activity

		query.Set("start", fmt.Sprintf("%d", start))
		query.Set("limit", fmt.Sprintf("%d", paginationLimit))
		request.URL.RawQuery = query.Encode()

		response, err := http.DefaultClient.Do(request)
		if err != nil {
			return -1, err
		}
		defer response.Body.Close()

		if err = httpStatusCodeError("get activities data", response); err != nil {
			return -1, err
		}

		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			return -1, fmt.Errorf("cannot parse activities response data: %v", err)
		}

		if err := json.Unmarshal(responseData, &tempActivites); err != nil {
			return -1, fmt.Errorf("cannot parse activities structure data: %v", err)
		}
		activities = append(activities, tempActivites...)

		if paginationLimit == len(tempActivites) {
			start += paginationLimit
			paginationLimit += paginationLimit
		} else {
			break
		}
	}

	return len(activities), nil

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

func (auth *Auth) ActivitiesDateFilter(start time.Time, end time.Time) (*[]Activity, error) {
	url := fmt.Sprintf("%s://%s%s", protocol, garmin_url, endpoint_search_activities)
	formatLayout := "2006-01-02"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	auth.addHeader(request)

	query := request.URL.Query()
	query.Add("startDate", fmt.Sprintf("%v", start.Format(formatLayout)))
	query.Add("endDate", fmt.Sprintf("%v", end.Format(formatLayout)))
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
