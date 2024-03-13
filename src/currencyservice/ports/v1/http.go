package portsv1

import (
	"log"
	"net/http"

	"buf.build/gen/go/coobeet/onlineboutique/connectrpc/go/ob/v1/obv1connect"
	"connectrpc.com/connect"
	connectcors "connectrpc.com/cors"
	"connectrpc.com/otelconnect"
	"github.com/rs/cors"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func NewHttpHandler() http.Handler {
	otelIntercepor, err := otelconnect.NewInterceptor()
	if err != nil {
		log.Fatal(err)
	}
	_, handler := obv1connect.NewCurrencyServiceHandler(
		NewConnectHandler(),
		connect.WithInterceptors(
			otelIntercepor,
		),
	)
	// Add HTTP instrumentation for the whole server.
	return otelhttp.NewHandler(withCors(handler), "/")
}

func withCors(handler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://buf.build"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
		MaxAge:         7200,
	})
	return c.Handler(handler)
}