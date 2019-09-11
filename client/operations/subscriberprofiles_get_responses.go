// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/negbie/ngcp-api/models"
)

// SubscriberprofilesGetReader is a Reader for the SubscriberprofilesGet structure.
type SubscriberprofilesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriberprofilesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscriberprofilesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubscriberprofilesGetOK creates a SubscriberprofilesGetOK with default headers values
func NewSubscriberprofilesGetOK() *SubscriberprofilesGetOK {
	return &SubscriberprofilesGetOK{}
}

/*SubscriberprofilesGetOK handles this case with default header values.

SubscriberprofilesGetOK subscriberprofiles get o k
*/
type SubscriberprofilesGetOK struct {
	Payload []*models.SubscriberProfile
}

func (o *SubscriberprofilesGetOK) Error() string {
	return fmt.Sprintf("[GET /subscriberprofiles][%d] subscriberprofilesGetOK  %+v", 200, o.Payload)
}

func (o *SubscriberprofilesGetOK) GetPayload() []*models.SubscriberProfile {
	return o.Payload
}

func (o *SubscriberprofilesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}