// Code generated by go-swagger; DO NOT EDIT.

package lnpnumbers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetLnpNumberitemsReader is a Reader for the GetLnpNumberitems structure.
type GetLnpNumberitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLnpNumberitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLnpNumberitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLnpNumberitemsOK creates a GetLnpNumberitemsOK with default headers values
func NewGetLnpNumberitemsOK() *GetLnpNumberitemsOK {
	return &GetLnpNumberitemsOK{}
}

/*GetLnpNumberitemsOK handles this case with default header values.

GetLnpNumberitemsOK get lnp numberitems o k
*/
type GetLnpNumberitemsOK struct {
	Payload []*models.LnpNumbers
}

func (o *GetLnpNumberitemsOK) Error() string {
	return fmt.Sprintf("[GET /lnpnumbers][%d] getLnpNumberitemsOK  %+v", 200, o.Payload)
}

func (o *GetLnpNumberitemsOK) GetPayload() []*models.LnpNumbers {
	return o.Payload
}

func (o *GetLnpNumberitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
