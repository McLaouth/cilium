// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewPostIpamParams creates a new PostIpamParams object
//
// There are no default values defined in the spec.
func NewPostIpamParams() PostIpamParams {

	return PostIpamParams{}
}

// PostIpamParams contains all the bound params for the post ipam operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostIpam
type PostIpamParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: header
	*/
	Expiration *bool
	/*
	  In: query
	*/
	Family *string
	/*
	  In: query
	*/
	Owner *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostIpamParams() beforehand.
func (o *PostIpamParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindExpiration(r.Header[http.CanonicalHeaderKey("expiration")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	qFamily, qhkFamily, _ := qs.GetOK("family")
	if err := o.bindFamily(qFamily, qhkFamily, route.Formats); err != nil {
		res = append(res, err)
	}

	qOwner, qhkOwner, _ := qs.GetOK("owner")
	if err := o.bindOwner(qOwner, qhkOwner, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindExpiration binds and validates parameter Expiration from header.
func (o *PostIpamParams) bindExpiration(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("expiration", "header", "bool", raw)
	}
	o.Expiration = &value

	return nil
}

// bindFamily binds and validates parameter Family from query.
func (o *PostIpamParams) bindFamily(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Family = &raw

	if err := o.validateFamily(formats); err != nil {
		return err
	}

	return nil
}

// validateFamily carries on validations for parameter Family
func (o *PostIpamParams) validateFamily(formats strfmt.Registry) error {

	if err := validate.EnumCase("family", "query", *o.Family, []interface{}{"ipv4", "ipv6"}, true); err != nil {
		return err
	}

	return nil
}

// bindOwner binds and validates parameter Owner from query.
func (o *PostIpamParams) bindOwner(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Owner = &raw

	return nil
}
