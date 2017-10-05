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

// Segment segment
// swagger:model segment

type Segment struct {

	// constraints
	Constraints SegmentConstraints `json:"constraints"`

	// description
	// Required: true
	// Min Length: 1
	Description *string `json:"description"`

	// distributions
	Distributions SegmentDistributions `json:"distributions"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`
}

/* polymorph segment constraints false */

/* polymorph segment description false */

/* polymorph segment distributions false */

/* polymorph segment id false */

// Validate validates this segment
func (m *Segment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Segment) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Segment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Segment) UnmarshalBinary(b []byte) error {
	var res Segment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}