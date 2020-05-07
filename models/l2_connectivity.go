// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// L2Connectivity l2 connectivity
//
// swagger:model l2-connectivity
type L2Connectivity struct {

	// outgoing ip address
	OutgoingIPAddress string `json:"outgoing_ip_address,omitempty"`

	// outgoing nic
	OutgoingNic string `json:"outgoing_nic,omitempty"`

	// remote ip address
	RemoteIPAddress string `json:"remote_ip_address,omitempty"`

	// remote mac
	RemoteMac string `json:"remote_mac,omitempty"`

	// successful
	Successful bool `json:"successful,omitempty"`
}

// Validate validates this l2 connectivity
func (m *L2Connectivity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *L2Connectivity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *L2Connectivity) UnmarshalBinary(b []byte) error {
	var res L2Connectivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
