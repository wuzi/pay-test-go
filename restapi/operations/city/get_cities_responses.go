// Code generated by go-swagger; DO NOT EDIT.

package city

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Wuzi/pay-test-go/models"
)

// GetCitiesOKCode is the HTTP code returned for type GetCitiesOK
const GetCitiesOKCode int = 200

/*GetCitiesOK Return an array with all cities

swagger:response getCitiesOK
*/
type GetCitiesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.City `json:"body,omitempty"`
}

// NewGetCitiesOK creates GetCitiesOK with default headers values
func NewGetCitiesOK() *GetCitiesOK {

	return &GetCitiesOK{}
}

// WithPayload adds the payload to the get cities o k response
func (o *GetCitiesOK) WithPayload(payload []*models.City) *GetCitiesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get cities o k response
func (o *GetCitiesOK) SetPayload(payload []*models.City) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCitiesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.City, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
