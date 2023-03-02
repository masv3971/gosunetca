package gosunetca

import (
	"context"
	"encoding/json"
	"gosunetca/mocks"
	"gosunetca/types"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestSignMarshal(t *testing.T) {
	tts := []struct {
		name string
		have interface{}
		want []byte
	}{
		{
			name: "Reply",
			have: mocks.MockReplySign,
			want: mocks.JSONSignDocumentReply200,
		},
		{
			name: "Request",
			have: mocks.MockRequestSign,
			want: mocks.JSONSignDocumentRequest200,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			var have interface{}
			switch tt.have.(type) {
			case *types.SignRequest:
				have = tt.have.(*types.SignRequest)
			case *types.SignReply:
				have = tt.have.(*types.SignReply)
			default:
				t.Errorf("unknown type %T", tt.have)
			}
			got, err := json.MarshalIndent(have, "", "	")
			assert.NoError(t, err)
			assert.JSONEq(t, string(tt.want), string(got))
		})
	}

}

func TestEndpointSignDocument(t *testing.T) {
	tts := []struct {
		name             string
		serverMethod     string
		serverURL        string
		serverReply      []byte
		serverStatusCode int
		serverToken      string
		clientRequest    *types.SignRequest
		want             interface{}
	}{
		{
			name:             "OK",
			serverMethod:     http.MethodPost,
			serverURL:        "/pkcs11_sign",
			serverReply:      mocks.JSONSignDocumentReply200,
			serverStatusCode: 200,
			serverToken:      "test-token",
			clientRequest:    mocks.MockRequestSign,
			want:             mocks.MockReplySign,
		},
		{
			name:             "BAD_TOKEN",
			serverMethod:     http.MethodPost,
			serverURL:        "/pkcs11_sign",
			serverReply:      mocks.JSONReply401,
			serverStatusCode: 401,
			serverToken:      "bad_token",
			clientRequest:    mocks.MockRequestSign,
			want:             mocks.MockErrorReplyMissingToken,
		},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mux, server, client := mockSetup(t)
			defer server.Close()

			mockGenericEndpointServer(t, mux, tt.serverToken, tt.serverMethod, tt.serverURL, tt.serverReply, tt.serverStatusCode)

			got, _, err := client.Sign.Documents(ctx, tt.clientRequest)
			switch tt.serverStatusCode {
			case 200:
				assert.NoError(t, err)
				if diff := cmp.Diff(got, tt.want); diff != "" {
					t.Errorf("mismatch (-got +want):\n%s", diff)
				}
			case 401:
				assert.Error(t, err)
				assert.Equal(t, tt.want, err)
			default:
				t.Errorf("unknown status code %d", tt.serverStatusCode)
			}

		})
	}
}
