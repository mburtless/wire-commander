// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeletePeerOKCode is the HTTP code returned for type DeletePeerOK
const DeletePeerOKCode int = 200

/*DeletePeerOK Returns empty object

swagger:response deletePeerOK
*/
type DeletePeerOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewDeletePeerOK creates DeletePeerOK with default headers values
func NewDeletePeerOK() *DeletePeerOK {

	return &DeletePeerOK{}
}

// WithPayload adds the payload to the delete peer o k response
func (o *DeletePeerOK) WithPayload(payload interface{}) *DeletePeerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete peer o k response
func (o *DeletePeerOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePeerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
