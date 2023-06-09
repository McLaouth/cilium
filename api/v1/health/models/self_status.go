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

// SelfStatus Description of the cilium-health node
//
// swagger:model SelfStatus
type SelfStatus struct {

	// Name associated with this node
	Name string `json:"name,omitempty"`
}

// Validate validates this self status
func (m *SelfStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this self status based on context it is used
func (m *SelfStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SelfStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SelfStatus) UnmarshalBinary(b []byte) error {
	var res SelfStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
