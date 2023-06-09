// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetMapNameEventsOKCode is the HTTP code returned for type GetMapNameEventsOK
const GetMapNameEventsOKCode int = 200

/*
GetMapNameEventsOK Success

swagger:response getMapNameEventsOK
*/
type GetMapNameEventsOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewGetMapNameEventsOK creates GetMapNameEventsOK with default headers values
func NewGetMapNameEventsOK() *GetMapNameEventsOK {

	return &GetMapNameEventsOK{}
}

// WithPayload adds the payload to the get map name events o k response
func (o *GetMapNameEventsOK) WithPayload(payload io.ReadCloser) *GetMapNameEventsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get map name events o k response
func (o *GetMapNameEventsOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMapNameEventsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetMapNameEventsNotFoundCode is the HTTP code returned for type GetMapNameEventsNotFound
const GetMapNameEventsNotFoundCode int = 404

/*
GetMapNameEventsNotFound Map not found

swagger:response getMapNameEventsNotFound
*/
type GetMapNameEventsNotFound struct {
}

// NewGetMapNameEventsNotFound creates GetMapNameEventsNotFound with default headers values
func NewGetMapNameEventsNotFound() *GetMapNameEventsNotFound {

	return &GetMapNameEventsNotFound{}
}

// WriteResponse to the client
func (o *GetMapNameEventsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
