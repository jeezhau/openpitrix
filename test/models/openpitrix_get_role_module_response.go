// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixGetRoleModuleResponse openpitrix get role module response
// swagger:model openpitrixGetRoleModuleResponse
type OpenpitrixGetRoleModuleResponse struct {

	// module
	Module *OpenpitrixModule `json:"module,omitempty"`

	// role id
	RoleID string `json:"role_id,omitempty"`
}

// Validate validates this openpitrix get role module response
func (m *OpenpitrixGetRoleModuleResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModule(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixGetRoleModuleResponse) validateModule(formats strfmt.Registry) error {

	if swag.IsZero(m.Module) { // not required
		return nil
	}

	if m.Module != nil {

		if err := m.Module.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("module")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixGetRoleModuleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixGetRoleModuleResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixGetRoleModuleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
