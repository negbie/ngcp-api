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

// SubscriberprofilesetsGetReader is a Reader for the SubscriberprofilesetsGet structure.
type SubscriberprofilesetsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriberprofilesetsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubscriberprofilesetsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubscriberprofilesetsGetOK creates a SubscriberprofilesetsGetOK with default headers values
func NewSubscriberprofilesetsGetOK() *SubscriberprofilesetsGetOK {
	return &SubscriberprofilesetsGetOK{}
}

/*SubscriberprofilesetsGetOK handles this case with default header values.

SubscriberprofilesetsGetOK subscriberprofilesets get o k
*/
type SubscriberprofilesetsGetOK struct {
	Payload []*models.SubscriberProfileSet
}

func (o *SubscriberprofilesetsGetOK) Error() string {
	return fmt.Sprintf("[GET /subscriberprofilesets][%d] subscriberprofilesetsGetOK  %+v", 200, o.Payload)
}

func (o *SubscriberprofilesetsGetOK) GetPayload() []*models.SubscriberProfileSet {
	return o.Payload
}

func (o *SubscriberprofilesetsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}