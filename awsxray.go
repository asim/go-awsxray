package awsxray

import (
	"encoding/json"
	"net"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/xray"
	"golang.org/x/net/context"
)

type contextSegmentKey struct{}

type AWSXRay struct {
	Options Options
}

var (
	TraceHeader = "X-Amzn-Trace-Id"
)

func (x *AWSXRay) Record(s *Segment) error {
	// marshal
	s.RLock()
	b, err := json.Marshal(s)
	if err != nil {
		s.RUnlock()
		return err
	}
	s.RUnlock()

	// Use XRay Client if available
	if x.Options.Client != nil {
		_, err := x.Options.Client.PutTraceSegments(&xray.PutTraceSegmentsInput{
			TraceSegmentDocuments: []*string{
				aws.String("TraceSegmentDocument"),
				aws.String(string(b)),
			},
		})
		return err
	}

	// Use Daemon
	c, err := net.Dial("udp", x.Options.Daemon)
	if err != nil {
		return err
	}

	header := append([]byte(`{"format": "json", "version": 1}`), byte('\n'))
	_, err = c.Write(append(header, b...))
	return err
}

func New(opts ...Option) *AWSXRay {
	options := Options{
		Daemon: "localhost:2000",
	}

	for _, o := range opts {
		o(&options)
	}

	return &AWSXRay{
		Options: options,
	}
}

func FromContext(ctx context.Context) (*Segment, bool) {
	s, ok := ctx.Value(contextSegmentKey{}).(*Segment)
	return s, ok
}

func NewContext(ctx context.Context, s *Segment) context.Context {
	return context.WithValue(ctx, contextSegmentKey{}, s)
}
