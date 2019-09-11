// Code generated by go-swagger; DO NOT EDIT.

package id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// CftimesetsByIDPutReader is a Reader for the CftimesetsByIDPut structure.
type CftimesetsByIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CftimesetsByIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCftimesetsByIDPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCftimesetsByIDPutOK creates a CftimesetsByIDPutOK with default headers values
func NewCftimesetsByIDPutOK() *CftimesetsByIDPutOK {
	return &CftimesetsByIDPutOK{}
}

/*CftimesetsByIDPutOK handles this case with default header values.

CftimesetsByIDPutOK cftimesets by Id put o k
*/
type CftimesetsByIDPutOK struct {
	Payload *models.CFTimeSets
}

func (o *CftimesetsByIDPutOK) Error() string {
	return fmt.Sprintf("[PUT /cftimesets/{id}][%d] cftimesetsByIdPutOK  %+v", 200, o.Payload)
}

func (o *CftimesetsByIDPutOK) GetPayload() *models.CFTimeSets {
	return o.Payload
}

func (o *CftimesetsByIDPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CFTimeSets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}