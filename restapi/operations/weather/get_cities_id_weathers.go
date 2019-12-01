// Code generated by go-swagger; DO NOT EDIT.

package weather

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCitiesIDWeathersHandlerFunc turns a function with the right signature into a get cities ID weathers handler
type GetCitiesIDWeathersHandlerFunc func(GetCitiesIDWeathersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCitiesIDWeathersHandlerFunc) Handle(params GetCitiesIDWeathersParams) middleware.Responder {
	return fn(params)
}

// GetCitiesIDWeathersHandler interface for that can handle valid get cities ID weathers params
type GetCitiesIDWeathersHandler interface {
	Handle(GetCitiesIDWeathersParams) middleware.Responder
}

// NewGetCitiesIDWeathers creates a new http.Handler for the get cities ID weathers operation
func NewGetCitiesIDWeathers(ctx *middleware.Context, handler GetCitiesIDWeathersHandler) *GetCitiesIDWeathers {
	return &GetCitiesIDWeathers{Context: ctx, Handler: handler}
}

/*GetCitiesIDWeathers swagger:route GET /cities/{id}/weathers Weather getCitiesIdWeathers

Show a list of weathers of a single city

Show a list of weathers of a single city

*/
type GetCitiesIDWeathers struct {
	Context *middleware.Context
	Handler GetCitiesIDWeathersHandler
}

func (o *GetCitiesIDWeathers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCitiesIDWeathersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
