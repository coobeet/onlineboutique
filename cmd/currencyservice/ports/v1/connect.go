package portsv1

import (
	"context"
	_ "embed"
	"encoding/json"
	"log"
	"slices"

	"buf.build/gen/go/coobeet/onlineboutique/connectrpc/go/ob/v1/obv1connect"
	obv1 "buf.build/gen/go/coobeet/onlineboutique/protocolbuffers/go/ob/v1"
	"connectrpc.com/connect"
	"github.com/samber/lo"
)

const CurrencyServicePath = "/" + obv1connect.CurrencyServiceName + "/"

//go:embed data/currency_conversion.json
var currencyConversionBytes []byte

var currencyConversion map[string]string

func init() {
	err := json.Unmarshal(currencyConversionBytes, &currencyConversion)
	if err != nil {
		log.Fatal(err)
	}
}

type connectHandler struct {
	obv1connect.UnimplementedCurrencyServiceHandler
}

func NewConnectHandler() obv1connect.CurrencyServiceHandler {
	return connectHandler{}
}

func (connectHandler) ListSupportedCurrencies(
	_ context.Context,
	req *connect.Request[obv1.ListSupportedCurrenciesRequest],
) (*connect.Response[obv1.ListSupportedCurrenciesResponse], error) {
	codes := lo.Keys(currencyConversion)
	slices.Sort(codes)
	return connect.NewResponse(&obv1.ListSupportedCurrenciesResponse{
		CurrencyCodes: codes,
	}), nil
}
