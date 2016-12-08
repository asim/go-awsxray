# Go AWS X-Ray

Go AWS X-Ray is a library to help simplify usage of AWS X-Ray for Micro

## Usage

```go
// create xray client
xray := awsxray.New(
        // Used as segment name
        awsxray.WithName("go.micro.srv.greeter"),
        // Specify X-Ray Daemon Address
        awsxray.WithDaemon("localhost:2000"),
        // Or X-Ray Client
        awsxray.WithClient(xray.New(awsSession)),
)

// create segment
segment := &awsxray.Segment{
	Id: ...
	TraceId: ...
	// more values
}

xray.Record(segment)
```

### With Wrappers

Wrappers scope the library for simplified use

```go
import "github.com/micro/go-plugins/wrapper/trace/awsxray"

opts := []awsxray.Option{
	// Used as segment name
	awsxray.WithName("go.micro.srv.greeter"),
	// Specify X-Ray Daemon Address
	awsxray.WithDaemon("localhost:2000"),
	// Or X-Ray Client
	awsxray.WithClient(xray.New(awsSession)),
}

service := micro.NewService(
	micro.Name("go.micro.srv.greeter"),
	micro.WrapCall(awsxray.NewCallWrapper(opts...)),
	micro.WrapClient(awsxray.NewClientWrapper(opts...)),
	micro.WrapHandler(awsxray.NewHandlerWrapper(opts...)),
)
```

## Example

<p align="center">
  <img src="awsxray.png" />
</p>
