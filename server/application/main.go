package application

import (
	"context"

	pb "github.com/OdaDaisuke/grpc-web/server/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
)

type MainApplication struct {
	Messages []*pb.Message
}

func (s *MainApplication) GetMessageStream(e *empty.Empty, stream pb.MessageService_GetMessageStreamServer) error {
	for _, m := range s.Messages {
		if err := stream.Send(m); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (s *MainApplication) PostMessage(ctx context.Context, message *pb.Message) (*pb.PostMessageResponse, error) {
	s.Messages = append(s.Messages, message)
	return &pb.PostMessageResponse{
		Status: "ok",
	}, nil
}
