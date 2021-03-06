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

// NewGetPeerParams creates a new GetPeerParams object
// no default values defined in spec.
func NewGetPeerParams() GetPeerParams {

	return GetPeerParams{}
}

// GetPeerParams contains all the bound params for the get peer operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetPeer
type GetPeerParams struct {

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
// To ensure default values, the struct must have been initialized with NewGetPeerParams() beforehand.
func (o *GetPeerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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
func (o *GetPeerParams) bindPeerName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.PeerName = raw

	return nil
}
