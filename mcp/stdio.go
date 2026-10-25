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

	gomcp.AddTool(server, &gomcp.Tool{Name: tool_activities,
		Description: `get user activities
		
		IMPORTANT:
		- Each response may be truncated if too many activities are requested.
		- Use limit=5 for each request.
		- Start the first request with start=0.
		- After each response, increment start by the number of activities actually returned.
		- Continue requesting pages until an empty page is returned.
		- Never assume that a truncated response contains all requested activities.
		`},
		activitiesHandler)

	dailySleepTool := gomcp.Tool{
		Name:        tool_daily_sleep,
		Description: "get user daily sleep",
	}

	gomcp.AddTool(server, &dailySleepTool, dailySleepHandler)

	if err := server.Run(context.Background(), &gomcp.StdioTransport{}); err != nil {
		return err
	}

	return nil
}
