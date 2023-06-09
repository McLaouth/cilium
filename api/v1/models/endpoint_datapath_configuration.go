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

// EndpointDatapathConfiguration Datapath configuration to be used for the endpoint
//
// swagger:model EndpointDatapathConfiguration
type EndpointDatapathConfiguration struct {

	// Disable source IP verification for the endpoint.
	//
	DisableSipVerification bool `json:"disable-sip-verification,omitempty"`

	// Indicates that IPAM is done external to Cilium. This will prevent the IP from being released and re-allocation of the IP address is skipped on restore.
	//
	ExternalIpam bool `json:"external-ipam,omitempty"`

	// Installs a route in the Linux routing table pointing to the device of the endpoint's interface.
	//
	InstallEndpointRoute bool `json:"install-endpoint-route,omitempty"`

	// Enable ARP passthrough mode
	RequireArpPassthrough bool `json:"require-arp-passthrough,omitempty"`

	// Endpoint requires a host-facing egress program to be attached to implement ingress policy and reverse NAT.
	//
	RequireEgressProg bool `json:"require-egress-prog,omitempty"`

	// Endpoint requires BPF routing to be enabled, when disabled, routing is delegated to Linux routing.
	//
	RequireRouting *bool `json:"require-routing,omitempty"`
}

// Validate validates this endpoint datapath configuration
func (m *EndpointDatapathConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this endpoint datapath configuration based on context it is used
func (m *EndpointDatapathConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EndpointDatapathConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointDatapathConfiguration) UnmarshalBinary(b []byte) error {
	var res EndpointDatapathConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
