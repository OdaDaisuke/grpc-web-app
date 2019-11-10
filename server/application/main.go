package application

import (
	"context"

	pb "github.com/OdaDaisuke/grpc-web/server/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
)

type Server struct {
}

func (s *Server) GetMessageStream(e *empty.Empty, params pb.MessageService_GetMessageStreamServer) error {
	return nil
}

func (s *Server) PostMessage(ctx context.Context, params *pb.Message) (*pb.PostMessageResponse, error) {
	return nil, nil
}
