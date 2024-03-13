module github.com/coobeet/onlineboutique/src/cartservice

go 1.22.1

require (
	buf.build/gen/go/coobeet/onlineboutique/connectrpc/go v1.15.0-20240313160751-8a920b4a57cd.1
	connectrpc.com/connect v1.15.0
	connectrpc.com/cors v0.1.0
	connectrpc.com/otelconnect v0.7.0
	github.com/rs/cors v1.10.1
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.49.0
	go.opentelemetry.io/otel v1.24.0
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.24.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.24.0
	go.opentelemetry.io/otel/sdk v1.24.0
	go.opentelemetry.io/otel/sdk/metric v1.24.0
)

require (
	buf.build/gen/go/coobeet/onlineboutique/protocolbuffers/go v1.33.0-20240313160751-8a920b4a57cd.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel/metric v1.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
