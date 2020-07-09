// Code generated by go-swagger; DO NOT EDIT.

package speeddials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetSpeedDialitemsReader is a Reader for the GetSpeedDialitems structure.
type GetSpeedDialitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpeedDialitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpeedDialitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSpeedDialitemsOK creates a GetSpeedDialitemsOK with default headers values
func NewGetSpeedDialitemsOK() *GetSpeedDialitemsOK {
	return &GetSpeedDialitemsOK{}
}

/*GetSpeedDialitemsOK handles this case with default header values.

GetSpeedDialitemsOK get speed dialitems o k
*/
type GetSpeedDialitemsOK struct {
	Payload []*models.SpeedDials
}

func (o *GetSpeedDialitemsOK) Error() string {
	return fmt.Sprintf("[GET /speeddials][%d] getSpeedDialitemsOK  %+v", 200, o.Payload)
}

func (o *GetSpeedDialitemsOK) GetPayload() []*models.SpeedDials {
	return o.Payload
}

func (o *GetSpeedDialitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
