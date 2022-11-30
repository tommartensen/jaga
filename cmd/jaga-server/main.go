package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	pb "github.com/tommartensen/jaga/generated/api/v1"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSegmentServer
}

func (s *server) Get(ctx context.Context, r *pb.SegmentRequest) (*pb.SegmentResponse, error) {
	log.Printf("Received GET key=%v", r.ID)
	return &pb.SegmentResponse{Name: "Hoch die Halde", Time: 23.05, Length: 450, Gradient: 4.0}, nil
}

func healthLivenessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func startHealthchecks() {
	log.Println("Starting Healthcheck server")
	r := mux.NewRouter()
	http.HandleFunc("/health", healthLivenessHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go startHealthchecks()

	s := grpc.NewServer()

	pb.RegisterSegmentServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
