// Code generated by go-swagger; DO NOT EDIT.

package sms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetSMSitemsReader is a Reader for the GetSMSitems structure.
type GetSMSitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSMSitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSMSitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSMSitemsOK creates a GetSMSitemsOK with default headers values
func NewGetSMSitemsOK() *GetSMSitemsOK {
	return &GetSMSitemsOK{}
}

/*GetSMSitemsOK handles this case with default header values.

GetSMSitemsOK get s m sitems o k
*/
type GetSMSitemsOK struct {
	Payload []*models.SM
}

func (o *GetSMSitemsOK) Error() string {
	return fmt.Sprintf("[GET /sms][%d] getSMSitemsOK  %+v", 200, o.Payload)
}

func (o *GetSMSitemsOK) GetPayload() []*models.SM {
	return o.Payload
}

func (o *GetSMSitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}