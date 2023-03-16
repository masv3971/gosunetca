package gosunetca

import (
	"context"
	"net/http"

	"github.com/masv3971/gosunetca/types"
)

// SignService is the service for signing
type SignService struct {
	client   *Client
	endpoint string
}

// Documents signs documents
func (s *SignService) Documents(ctx context.Context, body *types.SignRequest) (*types.SignReply, *http.Response, error) {
	if err := Check(body); err != nil {
		return nil, nil, err
	}

	reply := &types.SignReply{}

	resp, err := s.client.call(
		ctx,
		http.MethodPost,
		s.endpoint,
		body,
		reply,
	)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}
