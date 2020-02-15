package handlers

import (
    middleware "github.com/go-openapi/runtime/middleware"
    "github.com/go-openapi/swag"

	"github.com/mburtless/wire-commander/gen/models"
	"github.com/mburtless/wire-commander/gen/restapi/operations"
)

func NewGetPeer() operations.GetPeerHandler {
    return &getPeer{}
}

type getPeer struct {}

func (d *getPeer) Handle(params operations.GetPeerParams) middleware.Responder {
    peerName := swag.StringValue(&params.PeerName)
    if  peerName == "" {
        peerName = "Foo"
    }

    payload := &models.Peer{
        PeerName: peerName,
        PublicKey: "666",
    }
    return operations.NewGetPeerOK().WithPayload(payload)
}
