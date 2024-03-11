package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"log"
	"net/http"

	"buf.build/gen/go/coobeet/onlineboutique/connectrpc/go/ob/v1/obv1connect"
	obv1 "buf.build/gen/go/coobeet/onlineboutique/protocolbuffers/go/ob/v1"
	"connectrpc.com/connect"
	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
	"github.com/samber/lo"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

//go:embed data/currency_conversion.json
var currencyConversionBytes []byte
var currencyConversion map[string]string

func init() {
	if err := json.Unmarshal(currencyConversionBytes, &currencyConversion); err != nil {
		log.Fatal(err)
	}
}

func main() {
	srv := &currencyServer{}
	mux := http.NewServeMux()
	path, handler := obv1connect.NewCurrencyServiceHandler(srv)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(withCors(mux), &http2.Server{}),
	)
}

type currencyServer struct {
	obv1connect.UnimplementedCurrencyServiceHandler
}

func (currencyServer) ListSupportedCurrencies(
	_ context.Context,
	req *connect.Request[obv1.ListSupportedCurrenciesRequest],
) (*connect.Response[obv1.ListSupportedCurrenciesResponse], error) {
	return connect.NewResponse(&obv1.ListSupportedCurrenciesResponse{
		CurrencyCodes: lo.Keys(currencyConversion),
	}), nil
}

// withCors adds CORS support to a Connect HTTP handler.
func withCors(connectHandler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://buf.build"}, // replace with your domain
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
		MaxAge:         7200, // 2 hours in seconds
	})
	return c.Handler(connectHandler)
}
