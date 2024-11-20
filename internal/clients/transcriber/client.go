package transcription

import (
	"github.com/gomscourse/common/pkg/closer"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New(host string) (TranscriptionServiceClient, error) {
	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to transcriber service")
	}

	closer.Add(conn.Close)

	return NewTranscriptionServiceClient(conn), nil
}
