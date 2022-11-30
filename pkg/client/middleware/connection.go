package middleware

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/tommartensen/jaga/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateConnection(mainCtx context.Context, remoteAddr string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(mainCtx, time.Second)
	defer cancel()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock()}
	conn, err := grpc.DialContext(ctx, remoteAddr, opts...)
	base, cap := time.Second, time.Minute
	for backoff := base; err != nil; backoff <<= 1 {
		if backoff > cap {
			backoff = cap
		}
		jitter := rand.Int63n(int64(backoff))
		sleep := base + time.Duration(jitter)
		logging.Logger.Infow(
			fmt.Sprintf("could not open connection to %s, retrying in %v", remoteAddr, sleep),
			"remoteAddr", remoteAddr,
			"error", err,
		)
		time.Sleep(sleep)
		ctx, cancel = context.WithTimeout(mainCtx, time.Second)
		defer cancel()
		conn, err = grpc.DialContext(ctx, remoteAddr, opts...)
	}
	return conn, err
}
