// Code generated by go-swagger; DO NOT EDIT.

package cfbnumbersets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetCFBNumberSetitemsReader is a Reader for the GetCFBNumberSetitems structure.
type GetCFBNumberSetitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCFBNumberSetitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCFBNumberSetitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCFBNumberSetitemsOK creates a GetCFBNumberSetitemsOK with default headers values
func NewGetCFBNumberSetitemsOK() *GetCFBNumberSetitemsOK {
	return &GetCFBNumberSetitemsOK{}
}

/*GetCFBNumberSetitemsOK handles this case with default header values.

GetCFBNumberSetitemsOK get c f b number setitems o k
*/
type GetCFBNumberSetitemsOK struct {
	Payload []*models.CFBNumberSets1
}

func (o *GetCFBNumberSetitemsOK) Error() string {
	return fmt.Sprintf("[GET /cfbnumbersets][%d] getCFBNumberSetitemsOK  %+v", 200, o.Payload)
}

func (o *GetCFBNumberSetitemsOK) GetPayload() []*models.CFBNumberSets1 {
	return o.Payload
}

func (o *GetCFBNumberSetitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
