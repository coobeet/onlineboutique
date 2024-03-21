module github.com/coobeet/onlineboutique/src/currencyservice

go 1.22.1

require (
	buf.build/gen/go/coobeet/onlineboutique/connectrpc/go v1.15.0-20240313160751-8a920b4a57cd.1
	buf.build/gen/go/coobeet/onlineboutique/protocolbuffers/go v1.33.0-20240313160751-8a920b4a57cd.1
	connectrpc.com/connect v1.16.0
	connectrpc.com/cors v0.1.0
	connectrpc.com/otelconnect v0.7.0
	github.com/rs/cors v1.10.1
	github.com/samber/lo v1.39.0
	go.opentelemetry.io/otel v1.24.0
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.24.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.24.0
	go.opentelemetry.io/otel/sdk v1.24.0
	go.opentelemetry.io/otel/sdk/metric v1.24.0
)

require (
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel/metric v1.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
	golang.org/x/exp v0.0.0-20220303212507-bbda1eaf7a17 // indirect
	golang.org/x/sys v0.17.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
