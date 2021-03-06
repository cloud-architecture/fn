// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RoutesWrapper routes wrapper
// swagger:model RoutesWrapper
type RoutesWrapper struct {

	// error
	Error *ErrorBody `json:"error,omitempty"`

	// cursor to send with subsequent request to receive the next page, if non-empty
	// Read Only: true
	NextCursor string `json:"next_cursor,omitempty"`

	// routes
	// Required: true
	Routes RoutesWrapperRoutes `json:"routes"`
}

// Validate validates this routes wrapper
func (m *RoutesWrapper) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRoutes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoutesWrapper) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {

		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *RoutesWrapper) validateRoutes(formats strfmt.Registry) error {

	if err := validate.Required("routes", "body", m.Routes); err != nil {
		return err
	}

	if err := m.Routes.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("routes")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RoutesWrapper) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoutesWrapper) UnmarshalBinary(b []byte) error {
	var res RoutesWrapper
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
