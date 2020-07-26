package shortapi

import (
	"context"

	"github.com/short-d/app/fw/rpc"
	"github.com/short-d/short/backend/app/adapter/grpcapi/proto"
	"github.com/short-d/short/frontend/serve/entity"
)

// GRPC helps retrieve meta tags for a given short link using a gRPC client.
type GRPC struct {
	client proto.MetaTagServiceClient
}

// GetOpenGraphTags returns opengraph tags given an alias.
func (g *GRPC) GetOpenGraphTags(alias string) (entity.OpenGraphTags, error) {
	openGraphTags, err := g.client.GetOpenGraphTags(context.Background(), &proto.GetOpenGraphTagsRequest{Alias: alias})
	if err != nil {
		return entity.OpenGraphTags{}, err
	}

	return entity.OpenGraphTags{
		Title:       openGraphTags.GetTitle(),
		Description: openGraphTags.GetDescription(),
		ImageURL:    openGraphTags.GetImageUrl(),
	}, nil
}

// GetTwitterTags returns twitter tags given an alias.
func (g *GRPC) GetTwitterTags(alias string) (entity.TwitterTags, error) {
	twitterTags, err := g.client.GetTwitterTags(context.Background(), &proto.GetTwitterTagsRequest{Alias: alias})
	if err != nil {
		return entity.TwitterTags{}, err
	}

	return entity.TwitterTags{
		Title:       twitterTags.GetTitle(),
		Description: twitterTags.GetDescription(),
		ImageURL:    twitterTags.GetImageUrl(),
	}, nil
}

// NewGRPC initializes GRPC
func NewGRPC(hostname string, port int) (GRPC, error) {
	conn, err := rpc.NewClientConnBuilder(hostname, port).Build()
	if err != nil {
		return GRPC{}, err
	}
	client := proto.NewMetaTagServiceClient(conn)
	return GRPC{client: client}, nil
}