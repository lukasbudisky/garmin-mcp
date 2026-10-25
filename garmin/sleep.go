package garmin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (auth *Auth) GetDailySleep(date string) (*json.RawMessage, error) {
	url := fmt.Sprintf("%s://%s%s", protocol, garmin_url, endpoint_daily_sleep)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	parseLayout := "2006-01-02"
	parsedTime, err := time.Parse(parseLayout, date)
	if err != nil {
		return nil, fmt.Errorf("cannot parse time and date: %v", err)
	}

	query := request.URL.Query()
	query.Add("date", parsedTime.Format(parseLayout))
	request.URL.RawQuery = query.Encode()

	auth.addHeader(request)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err = httpStatusCodeError("get daily sleep data", response); err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot parse sleep data: %v", err)
	}

	return (*json.RawMessage)(&responseData), nil
}
