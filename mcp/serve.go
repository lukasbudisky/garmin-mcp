package mcp

import (
	"context"
	"net/http"

	garmin "github.com/lukasbudisky/garmin-api"
	gomcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

var garminAuth *garmin.Auth

func server(name string, version string) *gomcp.Server {
	mcp := gomcp.Implementation{
		Name:    name,
		Version: version,
	}

	return gomcp.NewServer(&mcp, nil)
}

func HttpServe(auth *garmin.Auth, name string, version string, address string, toolCapabilities bool) error {
	garminAuth = auth

	server := server(name, version)

	toolDefinition(server)

	if err := http.ListenAndServe(address, gomcp.NewStreamableHTTPHandler(func(r *http.Request) *gomcp.Server {
		return server
	}, nil)); err != nil {
		return err
	}

	return nil
}

func StdioServe(auth *garmin.Auth, name string, version string, toolCapabilities bool) error {
	garminAuth = auth

	server := server(name, version)

	toolDefinition(server)

	if err := server.Run(context.Background(), &gomcp.StdioTransport{}); err != nil {
		return err
	}

	return nil
}
