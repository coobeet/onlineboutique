package ports

import (
	"log"
	"net/http"

	"buf.build/gen/go/coobeet/onlineboutique/connectrpc/go/ob/v1/obv1connect"
	"connectrpc.com/connect"
	connectcors "connectrpc.com/cors"
	"connectrpc.com/otelconnect"
	portsv1 "github.com/coobeet/onlineboutique/cmd/currencyservice/ports/v1"
	"github.com/rs/cors"
)

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()
	otelIntercepor, err := otelconnect.NewInterceptor()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle(obv1connect.NewCurrencyServiceHandler(
		portsv1.NewConnectHandler(),
		connect.WithInterceptors(
			otelIntercepor,
		),
	))
	handler := withCors(mux)
	return handler
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
