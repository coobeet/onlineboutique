package portsv1

import (
	_ "embed"

	"buf.build/gen/go/coobeet/onlineboutique/connectrpc/go/ob/v1/obv1connect"
)

const CurrencyServicePath = "/" + obv1connect.CartServiceName + "/"

type connectHandler struct {
	obv1connect.UnimplementedCartServiceHandler
}

func NewConnectHandler() obv1connect.CartServiceHandler {
	return &connectHandler{}
}
