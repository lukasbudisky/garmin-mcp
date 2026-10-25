package mcp

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type DailySleepInput struct {
	Date string `json:"date,omitempty" jsonschema:"date for daily sleep results"`
}

type DailySleepOutput struct {
	DailySleep string
}

func dailySleepHandler(ctx context.Context, req *mcp.CallToolRequest, input DailySleepInput) (
	*mcp.CallToolResult,
	*DailySleepOutput,
	error,
) {
	dailySleep, err := garminAuth.GetDailySleep(input.Date)
	if err != nil {
		return nil, nil, err
	}

	return nil, &DailySleepOutput{DailySleep: string(*dailySleep)}, nil
}
