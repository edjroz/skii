package cli

import (
	"log"
	"regexp"

	"github.com/spf13/cobra"

	"context"
	"time"

	pb "github.com/edjroz/skii/types/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var color, start, addr string

// rootCmd represents the base command when called without any subcommands
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client to interact with gRPC server",
	Long:  `Skii is a compute engine that can retrieve all available paths for a skiier from a given point based on their difficulty as measured descending (black|red|blue)`,
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.AddCommand(queryCmd)

	queryCmd.Flags().StringVar(&color, "color", "blue", "color difficulty. defaults to blue(easiest)")
	queryCmd.Flags().StringVar(&start, "start", "", "starting point. ")
	queryCmd.Flags().StringVar(&addr, "addr", "localhost:50051", "address to connect to")
}

// startCmd represents the start command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "query which paths the skie can take",
	Long:  `calls gRPC server, starting point and difficulty. (start, diff) -> [[route1]...[routeN]]`,
	Run: func(cmd *cobra.Command, args []string) {
		if start == "" {
			log.Fatalf("must not have empty string for start flag")
		}
		if match, _ := regexp.MatchString("(black|red|blue)", color); !match {
			log.Fatalf("Requested color does not match ('black|red|blue')")
		}

		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewGreeterClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.Routes(ctx, &pb.RouteRequest{Start: start, Color: color})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Routes: %v", r)
	},
}
