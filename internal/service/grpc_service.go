package service

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/tanush-128/openzo_backend/image/config"
	"github.com/tanush-128/openzo_backend/image/internal/pb"

	"google.golang.org/grpc"
)

type Server struct {
	pb.ImageServiceServer
}

func GrpcServer(
	cfg *config.Config,
	server *Server,

) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCPort))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening at %v", lis.Addr())
	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterImageServiceServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *Server) UploadImage(ctx context.Context, req *pb.ImageMessage) (*pb.ImageURL, error) {
	url, err := SaveFile(req.GetImageData())
	if err != nil {
		return nil, err
	}
	// Implement your business logic here
	return &pb.ImageURL{
		Url: url,
	}, nil

}

func (s *Server) UploadMultipleImage(stream pb.ImageService_UploadMultipleImageServer) error {

	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		url, err := SaveFile(req.GetImageData())
		if err != nil {
			return err
		}

		err = stream.Send(&pb.ImageURL{
			Url: url,
		})
		if err != nil {
			return err
		}

	}

}
