package vivsemoapi

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/pellegrino/vivsemo-api/rpc/vivsemoapi"
)

// NewVivsemoAPIServer returns an instance of the twirp server that can be used as a http.Handler
func NewVivsemoAPIServer() pb.TwirpServer {
	svc := &vivsemoAPIServer{}
	return pb.NewApiServer(svc)
}

type vivsemoAPIServer struct{}

func (s *vivsemoAPIServer) GetAllPhotos(context.Context, *google_protobuf.Empty) (*pb.AllPhotosResponse, error) {
	gofakeit.Seed(0)
	photos := []*pb.Photo{}

	for i := 0; i < gofakeit.Number(1, 100); i++ { // start of the execution block
		p := &pb.Photo{
			Title:       gofakeit.Sentence(3),
			Description: gofakeit.Sentence(10),
			Url:         fmt.Sprintf("https://source.unsplash.com/1600x900/?family&i=%v", gofakeit.Number(1, 4000)),
		}

		photos = append(photos, p)
	}

	return &pb.AllPhotosResponse{Photos: photos}, nil
}
