package controller

import (
	"context"

	relayproxyv1 "github.com/thomaspoignant/go-feature-flag/cmd/relayproxy/schema/gen/go/v1"
)

type RelayProxyServiceServer struct{}

func (s *RelayProxyServiceServer) AllFlags(ctx context.Context, in *relayproxyv1.AllFlagsRequest) (*relayproxyv1.AllFlagsResponse, error) {
	data := &relayproxyv1.AllFlagsResponse{
		Flags: map[string]*relayproxyv1.FlagState{},
		Valid: false,
	}
	return data, nil
}
