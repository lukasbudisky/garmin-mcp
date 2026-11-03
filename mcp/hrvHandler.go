package mcp

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type DailyHrvInput struct {
	Date string `json:"date,omitempty" jsonschema:"date for daily HRV statistics results"`
}

type DailyHrvOutput struct {
	DailyHrv string
}

func dailyHrvHandler(ctx context.Context, req *mcp.CallToolRequest, input DailySleepInput) (
	*mcp.CallToolResult,
	*DailyHrvOutput,
	error,
) {
	dailyHrv, err := garminAuth.GetDailyHrv(input.Date)
	if err != nil {
		return nil, nil, err
	}

	return nil, &DailyHrvOutput{DailyHrv: string(*dailyHrv)}, nil
}
