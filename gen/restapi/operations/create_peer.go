// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreatePeerHandlerFunc turns a function with the right signature into a create peer handler
type CreatePeerHandlerFunc func(CreatePeerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreatePeerHandlerFunc) Handle(params CreatePeerParams) middleware.Responder {
	return fn(params)
}

// CreatePeerHandler interface for that can handle valid create peer params
type CreatePeerHandler interface {
	Handle(CreatePeerParams) middleware.Responder
}

// NewCreatePeer creates a new http.Handler for the create peer operation
func NewCreatePeer(ctx *middleware.Context, handler CreatePeerHandler) *CreatePeer {
	return &CreatePeer{Context: ctx, Handler: handler}
}

/*CreatePeer swagger:route POST /peers/{PeerName} createPeer

Creates a new peer

*/
type CreatePeer struct {
	Context *middleware.Context
	Handler CreatePeerHandler
}

func (o *CreatePeer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreatePeerParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
