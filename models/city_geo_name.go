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

// CityGeoName city geo name
// swagger:model CityGeoName
type CityGeoName struct {

	// CL of city
	// Required: true
	Cl *string `json:"cl"`

	// Code of city
	// Required: true
	Code *string `json:"code"`

	// Parent of city
	// Required: true
	Parent *float64 `json:"parent"`
}

// Validate validates this city geo name
func (m *CityGeoName) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CityGeoName) validateCl(formats strfmt.Registry) error {

	if err := validate.Required("cl", "body", m.Cl); err != nil {
		return err
	}

	return nil
}

func (m *CityGeoName) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *CityGeoName) validateParent(formats strfmt.Registry) error {

	if err := validate.Required("parent", "body", m.Parent); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CityGeoName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CityGeoName) UnmarshalBinary(b []byte) error {
	var res CityGeoName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
