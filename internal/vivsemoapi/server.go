package vivsemoapi

import (
	"context"

	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/pellegrino/vivsemo-api/rpc/vivsemoapi"
)

// NewVivsemoApiServer returns an instance of the twirp server that can be used as a http.Handler
func NewVivsemoApiServer() pb.TwirpServer {
	svc := &vivsemoApiServer{}
	return pb.NewApiServer(svc)
}

type vivsemoApiServer struct{}

func (s *vivsemoApiServer) GetAllPhotos(context.Context, *google_protobuf.Empty) (*pb.AllPhotosResponse, error) {
	return &pb.AllPhotosResponse{
		Photos: []*pb.Photo{{
			Title:       "new photo",
			Description: "description of a new photo",
			Url:         "https://source.unsplash.com/collection/190727/1600x900",
		}, {
			Title:       "a new random photo",
			Description: "description of a random family photo",
			Url:         "https://source.unsplash.com/1600x900/?family",
		},
		},
	}, nil
}
