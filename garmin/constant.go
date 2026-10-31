package garmin

const (
	// ----
	// Base
	// ----
	protocol   = "https"
	garmin_url = "personal.connectapi.garmin.com"

	// ----------------
	// Garmin Endpoints
	// ----------------

	// Api Key Endpoint
	endpoint_get_api_key = "/personalapitoken-service/personal-api/personal-api-key/token"

	// Search activities endpoint
	endpoint_search_activities = "/activitylist-service/activities/search/activities"

	// Daily sleep endpoint
	endpoint_daily_sleep = "/sleep-service/sleep/dailySleep"

	// Daily HRV endpoint
	endpoint_daily_hrv = "/hrv-service/hrv"

	// Daily user summary endpoint
	endpoint_daily_user_summary = "/usersummary-service/usersummary/daily"
)
