package segment

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
	"github.com/tommartensen/jaga/pkg/client/middleware"
	"github.com/tommartensen/jaga/pkg/conversions"
	"github.com/tommartensen/jaga/pkg/logging"

	pb "github.com/tommartensen/jaga/generated/api/v1"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "segment",
		Short: "Finds segments",
		Run:   fetchSegment,
	}
}

func fetchSegment(cmd *cobra.Command, args []string) {
	remoteAddr := "localhost:50051"
	ctx := cmd.Context()
	// Set up a connection to the gRPC server
	conn, err := middleware.CreateConnection(ctx, remoteAddr)
	if err != nil {
		logging.Logger.Fatalf(fmt.Sprintf("failed to open connection to %s:", remoteAddr), "remoteAddr", remoteAddr, "error", err)
	}
	defer conn.Close()

	client := pb.NewSegmentClient(conn)
	id := "anyId"
	res, err := RequestSegment(ctx, client, id)
	if err != nil {
		logging.Logger.Fatalf(fmt.Sprintf("failed to fetch segment '%s'", id), "error", err)
	}

	length := conversions.Length(res.Length)
	logging.Logger.Infof("Segment: %s, %s", res.Name, length.String())
}

func RequestSegment(mainCtx context.Context, client pb.SegmentClient, id string) (*pb.SegmentResponse, error) {
	ctx, cancel := context.WithTimeout(mainCtx, 10*time.Second)
	defer cancel()
	res, err := client.Get(ctx, &pb.SegmentRequest{ID: id})
	base, cap := time.Second, time.Minute
	count := 0
	for backoff := base; err != nil; backoff <<= 1 {
		if count >= 5 {
			return nil, fmt.Errorf("retries exhausted, original error: %v", err)
		}
		if backoff > cap {
			backoff = cap
		}
		jitter := rand.Int63n(int64(backoff))
		sleep := base + time.Duration(jitter)
		logging.Logger.Warnf(fmt.Sprintf("could not get segment for ID %s, retrying in %v", id, sleep), "error", err)
		time.Sleep(sleep)
		ctx, cancel := context.WithTimeout(mainCtx, time.Second)
		defer cancel()
		res, err = client.Get(ctx, &pb.SegmentRequest{ID: id})
		count++
	}
	return res, nil
}
