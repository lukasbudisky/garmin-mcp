package mcp

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type DailyUserSummaryInput struct {
	Date string `json:"date,omitempty" jsonschema:"date for daily user summary statisticsresults"`
}

type DailyUserSummaryOutput struct {
	DailyUserSummary string
}

func dailyUserSummaryHandler(ctx context.Context, req *mcp.CallToolRequest, input DailyUserSummaryInput) (
	*mcp.CallToolResult,
	*DailyUserSummaryOutput,
	error,
) {
	dailyUserSummary, err := garminAuth.GetDailySleep(input.Date)
	if err != nil {
		return nil, nil, err
	}

	return nil, &DailyUserSummaryOutput{DailyUserSummary: string(*dailyUserSummary)}, nil
}
