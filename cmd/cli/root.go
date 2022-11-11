package cli

import (
	"log"

	"github.com/edjroz/skii/cmd/server"
	"github.com/edjroz/skii/graph/parser"
	"github.com/spf13/cobra"
)

var path, port string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "skii",
	Short: "Skii provides a way for any skier to check which routes they can take",
	Long:  `Skii is a compute engine that can retrieve all available paths for a skiier from a given point based on their difficulty as measured descending (black|red|blue)`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringVar(&path, "path", "", "path to graphviz dot file")
	startCmd.Flags().StringVar(&port, "port", "50051", "server port")

}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start ",
	Short: "starts skii daemon",
	Long:  `starts gRPC server, with provided graph`,
	Run: func(cmd *cobra.Command, args []string) {
		if path == "" {
			log.Fatalf("path must not be empty")
		}

		g, err := parser.ReadFile(path)
		if err != nil {
			log.Fatalf("could not parse file: %v", err)
		}

		server.Start(g, port) // Get start from server?
	},
}
