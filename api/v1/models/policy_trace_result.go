// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyTraceResult Response to a policy resolution process
//
// swagger:model PolicyTraceResult
type PolicyTraceResult struct {

	// log
	Log string `json:"log,omitempty"`

	// verdict
	Verdict string `json:"verdict,omitempty"`
}

// Validate validates this policy trace result
func (m *PolicyTraceResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this policy trace result based on context it is used
func (m *PolicyTraceResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyTraceResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyTraceResult) UnmarshalBinary(b []byte) error {
	var res PolicyTraceResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
