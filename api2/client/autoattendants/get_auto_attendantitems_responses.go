// Code generated by go-swagger; DO NOT EDIT.

package autoattendants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetAutoAttendantitemsReader is a Reader for the GetAutoAttendantitems structure.
type GetAutoAttendantitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAutoAttendantitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAutoAttendantitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAutoAttendantitemsOK creates a GetAutoAttendantitemsOK with default headers values
func NewGetAutoAttendantitemsOK() *GetAutoAttendantitemsOK {
	return &GetAutoAttendantitemsOK{}
}

/*GetAutoAttendantitemsOK handles this case with default header values.

GetAutoAttendantitemsOK get auto attendantitems o k
*/
type GetAutoAttendantitemsOK struct {
	Payload []*models.AutoAttendants
}

func (o *GetAutoAttendantitemsOK) Error() string {
	return fmt.Sprintf("[GET /autoattendants][%d] getAutoAttendantitemsOK  %+v", 200, o.Payload)
}

func (o *GetAutoAttendantitemsOK) GetPayload() []*models.AutoAttendants {
	return o.Payload
}

func (o *GetAutoAttendantitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}