package mcp

import (
	"context"
	"fmt"
	"time"

	"github.com/lukasbudisky/garmin-mcp/garmin"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type AllActivityCountInput struct {
	Limit int `json:"limit,omitempty" jsonschema:"pagination maximum number of activities to return, default: 10"`
}

type AllActivityCountOutput struct {
	Total int
}

type ActivityInput struct {
	Start int `json:"start,omitempty" jsonschema:"pagination offset of the first activity to return, default: 0"`
	Limit int `json:"limit,omitempty" jsonschema:"pagination maximum number of activities to return, default: 10"`
}

type ActivityOutput struct {
	Activity *[]garmin.Activity
}

type AllActivityDateFilterInput struct {
	Start string `json:"startDate,omitempty" jsonschema:"define start date where filter start, format: YYYY-MM-DD"`
	End   string `json:"endDate,omitempty" jsonschema:"define end date where filter stop, format: YYYY-MM-DD"`
}

func activitiesHandler(ctx context.Context, req *mcp.CallToolRequest, input ActivityInput) (
	*mcp.CallToolResult, *ActivityOutput, error) {

	activities, err := garminAuth.SearchActivities(input.Start, input.Limit)
	if err != nil {
		return nil, nil, err
	}

	return nil, &ActivityOutput{Activity: activities}, nil
}

func allActivitiesCountHandler(ctx context.Context, req *mcp.CallToolRequest, input AllActivityCountInput) (
	*mcp.CallToolResult, *AllActivityCountOutput, error) {

	activitiesCount, err := garminAuth.CountAllActivities(input.Limit)
	if err != nil {
		return nil, nil, err
	}

	return nil, &AllActivityCountOutput{Total: activitiesCount}, nil
}

func activitiesDateFilterHandler(ctx context.Context, req *mcp.CallToolRequest, input AllActivityDateFilterInput) (
	*mcp.CallToolResult, *ActivityOutput, error) {

	formatLayout := "2006-01-02"

	start, err := time.Parse(formatLayout, input.Start)
	if err != nil {
		return nil, nil, fmt.Errorf("activities date filter handler start time parse error: %v", err)
	}

	end, err := time.Parse(formatLayout, input.End)
	if err != nil {
		return nil, nil, fmt.Errorf("activities date filter handler end time parse error: %v", err)
	}
	activities, err := garminAuth.ActivitiesDateFilter(start, end)
	if err != nil {
		return nil, nil, err
	}

	return nil, &ActivityOutput{Activity: activities}, nil
}
