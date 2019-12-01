// Code generated by go-swagger; DO NOT EDIT.

package city

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "github.com/Wuzi/pay-test-go/models"
)

// GetCitiesIDHandlerFunc turns a function with the right signature into a get cities ID handler
type GetCitiesIDHandlerFunc func(GetCitiesIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCitiesIDHandlerFunc) Handle(params GetCitiesIDParams) middleware.Responder {
	return fn(params)
}

// GetCitiesIDHandler interface for that can handle valid get cities ID params
type GetCitiesIDHandler interface {
	Handle(GetCitiesIDParams) middleware.Responder
}

// NewGetCitiesID creates a new http.Handler for the get cities ID operation
func NewGetCitiesID(ctx *middleware.Context, handler GetCitiesIDHandler) *GetCitiesID {
	return &GetCitiesID{Context: ctx, Handler: handler}
}

/*GetCitiesID swagger:route GET /cities/{id} City getCitiesId

Show a single city by id

Show a single city by id

*/
type GetCitiesID struct {
	Context *middleware.Context
	Handler GetCitiesIDHandler
}

func (o *GetCitiesID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCitiesIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetCitiesIDOKBody get cities ID o k body
// swagger:model GetCitiesIDOKBody
type GetCitiesIDOKBody struct {
	models.City

	// weather
	Weather *models.Weather `json:"weather,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetCitiesIDOKBody) UnmarshalJSON(raw []byte) error {
	// GetCitiesIDOKBodyAO0
	var getCitiesIDOKBodyAO0 models.City
	if err := swag.ReadJSON(raw, &getCitiesIDOKBodyAO0); err != nil {
		return err
	}
	o.City = getCitiesIDOKBodyAO0

	// GetCitiesIDOKBodyAO1
	var dataGetCitiesIDOKBodyAO1 struct {
		Weather *models.Weather `json:"weather,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetCitiesIDOKBodyAO1); err != nil {
		return err
	}

	o.Weather = dataGetCitiesIDOKBodyAO1.Weather

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetCitiesIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getCitiesIDOKBodyAO0, err := swag.WriteJSON(o.City)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getCitiesIDOKBodyAO0)

	var dataGetCitiesIDOKBodyAO1 struct {
		Weather *models.Weather `json:"weather,omitempty"`
	}

	dataGetCitiesIDOKBodyAO1.Weather = o.Weather

	jsonDataGetCitiesIDOKBodyAO1, errGetCitiesIDOKBodyAO1 := swag.WriteJSON(dataGetCitiesIDOKBodyAO1)
	if errGetCitiesIDOKBodyAO1 != nil {
		return nil, errGetCitiesIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetCitiesIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get cities ID o k body
func (o *GetCitiesIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.City
	if err := o.City.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWeather(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCitiesIDOKBody) validateWeather(formats strfmt.Registry) error {

	if swag.IsZero(o.Weather) { // not required
		return nil
	}

	if o.Weather != nil {
		if err := o.Weather.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getCitiesIdOK" + "." + "weather")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCitiesIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCitiesIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetCitiesIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}