package mcp

import (
	"context"

	"github.com/lukasbudisky/garmin-mcp/garmin"
	gomcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

var garminAuth *garmin.Auth

func StdioServer(auth *garmin.Auth, name string, version string, toolCapabilities bool) error {
	garminAuth = auth

	mcp := gomcp.Implementation{
		Name:    name,
		Version: version,
	}

	server := gomcp.NewServer(&mcp, nil)

	allActivitiesCountTool := gomcp.Tool{
		Name: tool_all_activities_count,
		Description: `get total count of all activities
		
		IMPORTANT:
		- use this tool if you want to verify total count of activities
		  if it match with list of activities which you got
		- default limit is 10
		- max limit is 99 do not exeed it because it will fail
		`,
	}

	activitiesTool := gomcp.Tool{
		Name: tool_activities,
		Description: `get user activities
		
		IMPORTANT:
		- Each response may be truncated if too many activities are requested.
		- Use limit=5 for each request.
		- Start the first request with start=0.
		- After each response, increment start by the number of activities actually returned.
		- Continue requesting pages until an empty page is returned.
		- Never assume that a truncated response contains all requested activities.
		- If you are trying to get download all activities you should always call
		  call CountAllActivities function for verification for count
		`}

	activitiesDateFilter := gomcp.Tool{
		Name: tool_activities_date_filter,
		Description: `filter activities by date range
	
		LIMITATION:
		When this function is called with range dates which are too big it
		can happen that result will be cut because it can be limited by model
		context length.

		ALWAYS:
		You need to always call function for counting activities between specific
		range dates because model context length can be limited and results can be
		cut.

		IMPORTANT:
		For start and end dates use this date format YYYY-MM-DD
		`,
	}

	dailySleepTool := gomcp.Tool{
		Name: tool_daily_sleep,
		Description: `Retrieves daily sleep data for a calendar date
		
		USAGE:
		User get useful information about their sleep including phases of sleep
		like REM, awake, deep sleep.

		WHEN NOT TO USE:
		This call is not suitable for specific date range

		IMPORTANT:
		For start and end dates use this date format YYYY-MM-DD
		`,
	}

	dailyHrvTool := gomcp.Tool{
		Name: tool_daily_hrv,
		Description: `Get HRV data for a specific date
		
		WHEN NOT TO USE:
		This call is not suitable for specific date range

		IMPORTANT:
		For start and end dates use this date format YYYY-MM-DD
		`,
	}

	dailyUserSummaryTool := gomcp.Tool{
		Name: tool_daily_user_summary,
		Description: `Get user summary data for a specific date

		USAGE:
		User summary data includes steps, calories, stress,  Body Battery
		and other wellness metrics for the given date. Call this function 
		when user ask for their steps, calories, stress, Body Battery or daily data.

		WHEN NOT TO USE:
		If user ask for sleep data or HRV data it is better to use other functions more suitable for sleep data

		IMPORTANT:
		For start and end dates use this date format YYYY-MM-DD
		`,
	}

	gomcp.AddTool(server, &allActivitiesCountTool, allActivitiesCountHandler)
	gomcp.AddTool(server, &activitiesTool, activitiesHandler)
	gomcp.AddTool(server, &activitiesDateFilter, activitiesDateFilterHandler)
	gomcp.AddTool(server, &dailySleepTool, dailySleepHandler)
	gomcp.AddTool(server, &dailyHrvTool, dailyHrvHandler)
	gomcp.AddTool(server, &dailyUserSummaryTool, dailyUserSummaryHandler)

	if err := server.Run(context.Background(), &gomcp.StdioTransport{}); err != nil {
		return err
	}

	return nil
}
