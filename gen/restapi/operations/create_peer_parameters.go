// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCreatePeerParams creates a new CreatePeerParams object
// no default values defined in spec.
func NewCreatePeerParams() CreatePeerParams {

	return CreatePeerParams{}
}

// CreatePeerParams contains all the bound params for the create peer operation
// typically these are obtained from a http.Request
//
// swagger:parameters CreatePeer
type CreatePeerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Name of peer
	  Required: true
	  In: path
	*/
	PeerName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreatePeerParams() beforehand.
func (o *CreatePeerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rPeerName, rhkPeerName, _ := route.Params.GetOK("PeerName")
	if err := o.bindPeerName(rPeerName, rhkPeerName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPeerName binds and validates parameter PeerName from path.
func (o *CreatePeerParams) bindPeerName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.PeerName = raw

	return nil
}
