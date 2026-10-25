package mcp

import (
	"context"

	"github.com/lukasbudisky/garmin-mcp/garmin"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ActivityInput struct {
	Start int `json:"start,omitempty" jsonschema:"pagination offset of the first activity to return, default: 0"`
	Limit int `json:"limit,omitempty" jsonschema:"pagination maximum number of activities to return, default: 5"`
}

type ActivityOutput struct {
	Activity *[]garmin.Activity
}

func activitiesHandler(ctx context.Context, req *mcp.CallToolRequest, input ActivityInput) (
	*mcp.CallToolResult,
	*ActivityOutput,
	error,
) {
	activities, err := garminAuth.SearchActivities(input.Start, input.Limit)
	if err != nil {
		return nil, nil, err
	}

	return nil, &ActivityOutput{Activity: activities}, nil
}
