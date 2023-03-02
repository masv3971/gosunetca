package gosunetca

import (
	"context"
	"gosunetca/types"
	"net/http"
)

type endpointsSign struct {
	client   *Client
	endpoint string
}

func (s *endpointsSign) Documents(ctx context.Context, body *types.SignRequest) (*types.SignReply, *http.Response, error) {
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
