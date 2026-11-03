package main

import (
	"fmt"
	"log"
	"os"

	garmin "github.com/lukasbudisky/garmin-api"
	"github.com/lukasbudisky/garmin-mcp/mcp"
	"github.com/spf13/cobra"
)

const (
	name    = "garmin-mcp"
	version = "0.0.1"
)

var (
	address string
	port    int
	apiKey  string
)

var rootCmd = &cobra.Command{
	Use:   "garmin-mcp",
	Short: "Garmin MCP communication between AI assistent and Garmin personal data.",
	Long: `Garmin MCP provides communication between AI assistant and Garmin personal data

Connection between AI assistnant and Garmin helps user identify health
and fitness gaps, spot anomalies, analyse data and suggest improvments for user
daily life and training.
	`,
}

var stdioCmd = &cobra.Command{
	Use:   "stdio",
	Short: "Execute MCP server using STDIO",
	Long: `Executable command for starting the MCP server using STDIO.
	`,
	Run: stdioServe,
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start and serve MCP server",
	Long: `Starts and runs the MCP server, listening on defined address and port.
	`,
	Run: httpServe,
}

func init() {
	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		rootCmd.PersistentFlags().StringVarP(&apiKey, "key", "k", "", "personal Garmin API key")
	}

	if apiKey == "" {
		log.Fatalln("Garmin API key is not defined")
	}

	httpCmd.PersistentFlags().StringVarP(&address, "address", "a", "127.0.0.1", "server address for MCP")
	httpCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "server port for MCP")

	rootCmd.AddCommand(stdioCmd, httpCmd)
}

func stdioServe(cmd *cobra.Command, args []string) {
	token, err := garmin.GetAccessToken(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	err = mcp.StdioServe(&garmin.Auth{Token: token}, name, version, false)
	if err != nil {
		log.Panic(err)
	}
}

func httpServe(cmd *cobra.Command, args []string) {
	token, err := garmin.GetAccessToken(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	httpAddress := fmt.Sprintf("%s:%d", address, port)

	log.Printf("start server on %s\n", httpAddress)

	errCh := make(chan error, 1)

	go func() {
		errCh <- mcp.HttpServe(&garmin.Auth{Token: token}, name, version, httpAddress, false)
		if err != nil {
			log.Panic(err)
		}
	}()

	if err = <-errCh; err != nil {
		log.Fatalf("Http server %s stopped with error: %v", address, err)
	}

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
