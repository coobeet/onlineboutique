package portsv1

import (
	_ "embed"

	"buf.build/gen/go/coobeet/onlineboutique/connectrpc/go/ob/v1/obv1connect"
)

type connectHandler struct {
	obv1connect.UnimplementedCartServiceHandler
}
