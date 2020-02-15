// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mburtless/wire-commander/gen/models"
)

// GetPeerOKCode is the HTTP code returned for type GetPeerOK
const GetPeerOKCode int = 200

/*GetPeerOK The details of the peer

swagger:response getPeerOK
*/
type GetPeerOK struct {

	/*
	  In: Body
	*/
	Payload *models.Peer `json:"body,omitempty"`
}

// NewGetPeerOK creates GetPeerOK with default headers values
func NewGetPeerOK() *GetPeerOK {

	return &GetPeerOK{}
}

// WithPayload adds the payload to the get peer o k response
func (o *GetPeerOK) WithPayload(payload *models.Peer) *GetPeerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get peer o k response
func (o *GetPeerOK) SetPayload(payload *models.Peer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPeerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
