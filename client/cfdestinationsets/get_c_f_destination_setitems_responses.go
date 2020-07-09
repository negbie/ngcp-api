// Code generated by go-swagger; DO NOT EDIT.

package cfdestinationsets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetCFDestinationSetitemsReader is a Reader for the GetCFDestinationSetitems structure.
type GetCFDestinationSetitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCFDestinationSetitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCFDestinationSetitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCFDestinationSetitemsOK creates a GetCFDestinationSetitemsOK with default headers values
func NewGetCFDestinationSetitemsOK() *GetCFDestinationSetitemsOK {
	return &GetCFDestinationSetitemsOK{}
}

/*GetCFDestinationSetitemsOK handles this case with default header values.

GetCFDestinationSetitemsOK get c f destination setitems o k
*/
type GetCFDestinationSetitemsOK struct {
	Payload []*models.CFDestinationSets2
}

func (o *GetCFDestinationSetitemsOK) Error() string {
	return fmt.Sprintf("[GET /cfdestinationsets][%d] getCFDestinationSetitemsOK  %+v", 200, o.Payload)
}

func (o *GetCFDestinationSetitemsOK) GetPayload() []*models.CFDestinationSets2 {
	return o.Payload
}

func (o *GetCFDestinationSetitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
