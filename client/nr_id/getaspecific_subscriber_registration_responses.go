// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetaspecificSubscriberRegistrationReader is a Reader for the GetaspecificSubscriberRegistration structure.
type GetaspecificSubscriberRegistrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetaspecificSubscriberRegistrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetaspecificSubscriberRegistrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetaspecificSubscriberRegistrationOK creates a GetaspecificSubscriberRegistrationOK with default headers values
func NewGetaspecificSubscriberRegistrationOK() *GetaspecificSubscriberRegistrationOK {
	return &GetaspecificSubscriberRegistrationOK{}
}

/*GetaspecificSubscriberRegistrationOK handles this case with default header values.

GetaspecificSubscriberRegistrationOK getaspecific subscriber registration o k
*/
type GetaspecificSubscriberRegistrationOK struct {
	Payload *models.SubscriberRegistrations
}

func (o *GetaspecificSubscriberRegistrationOK) Error() string {
	return fmt.Sprintf("[GET /subscriberregistrations/{id}][%d] getaspecificSubscriberRegistrationOK  %+v", 200, o.Payload)
}

func (o *GetaspecificSubscriberRegistrationOK) GetPayload() *models.SubscriberRegistrations {
	return o.Payload
}

func (o *GetaspecificSubscriberRegistrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriberRegistrations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
